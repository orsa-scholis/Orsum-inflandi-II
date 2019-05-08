package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/google/logger"
	"net"
	"strconv"
	"strings"
)

type clientState int

const (
	connecting clientState = iota
	inLobby
	inGame
	playingGame
	disconnected
)

func (state clientState) String() string {
	states := [...]string{
		"Connecting",
		"InLobby",
		"InGame",
		"Disconnected"}

	if state < connecting || state > disconnected {
		return "Unknown ClientState"
	}

	return states[state]
}

type message struct {
	receive    bool
	domain     string
	command    string
	param      string
	attachment string
}

func (m *message) String() string {
	return fmt.Sprintf("%s:%s:%s", m.domain, m.command, m.param)
}

func parseMessage(raw string) (m message, err error) {
	messageParts := strings.Split(raw, ":")
	if len(messageParts) == 3 {
		m = message{
			receive: true,
			domain:  messageParts[0],
			command: messageParts[1],
			param:   messageParts[2],
		}
	}
	if len(messageParts) == 2 {
		m = message{
			receive: true,
			domain:  messageParts[0],
			command: messageParts[1],
			param:   "",
		}
	}
	if len(messageParts) < 2 || len(messageParts) > 3 {
		err = errors.New("message length invalid")
	}

	return
}

type messageValidationError struct {
	err            string
	returnToClient bool
}

func (messageErr *messageValidationError) Error() string {
	return messageErr.err
}

func validateMessageDomainCommand(received message, expected message) (errs []error) {
	if !received.receive {
		errs = append(errs, &messageValidationError{err: "the message needs to be received"})
	}

	if expected.domain == "*" {
		return
	}

	if received.domain != expected.domain {
		errs = append(errs, &messageValidationError{err: fmt.Sprintf("the domain '%v' does not match expected '%v'", received.domain, expected.domain)})
	}
	if received.command != expected.command {
		errs = append(errs, &messageValidationError{err: fmt.Sprintf("the command '%v' does not match expected '%v'", received.command, expected.command)})
	}
	return errs
}

func validateMessageParam(received message, expected message) error {
	if expected.param[0] == '*' && len(expected.param) > 1 {
		rules := expected.param[2:]
		ruleList := strings.Split(rules, ",")
		if len(ruleList) > 0 {
			ruleErrors := ""
			for _, rule := range ruleList {
				ruleParts := strings.Split(rule, ";")

				var ruleError error
				if len(ruleParts) > 1 {
					logger.Infof("executing param checker '%s' with attr '%v' and param '%s'", ruleParts[0], ruleParts[1], received.param)
					ruleError = validatorsWithParam[ruleParts[0]](received.param, ruleParts[1])
				} else {
					logger.Infof("executing param checker '%s' with param '%s'", ruleParts[0], received.param)
					ruleError = validators[ruleParts[0]](received.param)
				}

				if ruleError != nil {
					logger.Infof("the param '%s' does not match rule '%s'", received.param, rule)
					ruleErrors += rule + ","
				}
			}

			if ruleErrors != "" {
				fullErrorMessage := fmt.Sprintf("%v", ruleErrors)
				return &messageValidationError{err: fullErrorMessage[:len(fullErrorMessage)-1], returnToClient: true}
			}
		}
	}

	return nil
}

type client struct {
	name        string
	conn        net.Conn
	stoppedChan chan bool
	receiveChan chan message
	sendChan    chan message
	state       clientState
	server      *server
	currentGame *game
}

func initClient(name string, conn net.Conn, server *server) (c *client) {
	c = &client{
		name:        name,
		conn:        conn,
		stoppedChan: make(chan bool, 1),
		receiveChan: make(chan message),
		sendChan:    make(chan message, 10),
		state:       connecting,
		server:      server,
	}

	go c.receiver()
	go c.sender()

	return
}

func (c *client) receiver() {
	for {
		raw, _, err := bufio.NewReader(c.conn).ReadLine()

		if err != nil {
			select {
			case <-c.stoppedChan:
				logger.Info("server is stopped, can't accept anymore messages")
				return
			default:
				logger.Warning("can't accept message: ", err)
				return
			}
		}

		parsed, err := parseMessage(string(raw))
		if err != nil {
			c.sendChan <- message{
				receive: false,
				domain:  "error",
				command: "message",
				param:   err.Error(),
			}
		}
		logger.Info(fmt.Sprintf("Client '%s' received message: '%+v'", c.name, parsed))
		c.receiveChan <- parsed
	}
}

func (c *client) sender() {
	for {
		sendMessage := <-c.sendChan
		_, err := c.conn.Write([]byte(sendMessage.String() + "\n"))

		if err != nil {
			select {
			case <-c.stoppedChan:
				logger.Info("server is stopped, can't accept anymore messages")
				return
			default:
				logger.Warning("can't write message: ", err)
				continue
			}
		}

		logger.Infof(fmt.Sprintf("Client '%s' sent message: '%+v'", c.name, sendMessage))
	}
}

func (c *client) start() {
	for {
		cMessage := <-c.receiveChan

		messageHandled, messageError := c.checkMessageHandlersIfFit(cMessage)

		if !messageHandled {
			logger.Warningf("unhandled message: '%s', \nerrors: '%v+'\nstate: '%s'", cMessage.String(), messageError, c.state)
			c.sendChan <- message{
				receive: false,
				domain:  "error",
				command: cMessage.command,
				param:   messageError,
			}
		}
	}
}

func (c *client) checkMessageHandlersIfFit(cMessage message) (messageHandled bool, messageError string) {
	stateFunctions := clientStateMessageHandlers[c.state]
	for expMessage, messageFunc := range stateFunctions {
		errs := validateMessageDomainCommand(cMessage, expMessage)
		if len(errs) > 0 {
			logger.Infof("%+v", errs)
			continue
		}
		if (cMessage.param == "" || len(cMessage.param) == 0) && expMessage.param == "*" {
			continue
		}

		paramErr := validateMessageParam(cMessage, expMessage)
		errs = append(errs, paramErr)
		if paramErr == nil {
			messageFunc(c, cMessage)
			logger.Infof("Client '%s' handled message: '%+v', with messageHandler: '%s'", c.name, cMessage, expMessage.String())

			messageHandled = true
			break
		} else {
			if err, ok := paramErr.(*messageValidationError); ok && err.returnToClient {
				messageError = err.Error()
			}
		}
	}
	return
}

func connectionConnectHandler(c *client, recMessage message) {
	c.name = recMessage.param
	c.state = inLobby

	c.sendChan <- message{
		receive: false,
		domain:  "success",
		command: "accepted",
		param:   "",
	}
}

func infoRequestGamesHandler(c *client, _ message) {
	games := c.server.getGamesAsString()

	c.sendChan <- message{
		receive: false,
		domain:  "success",
		command: "requested",
		param:   games,
	}
}

func serverNewGameHandler(c *client, recMessage message) {
	success := c.server.openGame(recMessage.param, c)

	if success {
		c.sendChan <- message{
			receive: false,
			domain:  "success",
			command: "created",
			param:   recMessage.param,
		}

		go c.server.broadcastMessage(message{
			receive: false,
			domain:  "subscription",
			command: "gameAdded",
			param:   c.server.getGamesAsString(),
		}, c)
	} else {
		c.sendChan <- message{
			receive: false,
			domain:  "error",
			command: "newGame",
			param:   "game name not unique",
		}
	}
}

func gameJoinHandler(c *client, recMessage message) {
	gameId, _ := strconv.Atoi(recMessage.param)
	clientState, game := c.server.joinGame(gameId, c)

	if clientState == inLobby {
		c.sendChan <- message{
			receive: false,
			domain:  "error",
			command: "join",
			param:   "game full or non existent",
		}

		return
	}

	c.state = clientState
	c.currentGame = &game

	var successMessageTemplate = message{
		receive: false,
		domain:  "success",
		command: "joined",
	}

	if clientState == inGame {
		successMessageTemplate.param = "1"
		c.sendChan <- successMessageTemplate
	} else if clientState == playingGame {
		successMessageTemplate.param = "2"
		c.sendChan <- successMessageTemplate
	}
}

func gameSetStoneHandler(c *client, recMessage message) {
	rowNr, _ := strconv.Atoi(recMessage.param)
	success := c.currentGame.setStone(c, rowNr)

	if success {
		c.sendChan <- message{
			receive: false,
			domain:  "success",
			command: "setStone",
			param:   "",
		}
	} else {
		c.sendChan <- message{
			receive: false,
			domain:  "error",
			command: "setStone",
			param:   "",
		}
	}
}

var clientStateMessageHandlers = map[clientState]map[message]func(*client, message){
	connecting: {
		connectionConnect: connectionConnectHandler,
	},
	inLobby: {
		infoRequestGames: infoRequestGamesHandler,
		serverNewGame:    serverNewGameHandler,
		gameJoin:         gameJoinHandler,
	},
	inGame: {
		gameSetStone: gameSetStoneHandler,
	},
	playingGame: {},
}

var connectionConnect = message{
	receive: true,
	domain:  "connection",
	command: "connect",
	param:   "*",
}

var infoRequestGames = message{
	receive: true,
	domain:  "info",
	command: "requestGames",
	param:   "*",
}

var serverNewGame = message{
	receive: true,
	domain:  "server",
	command: "newGame",
	param:   "*|required,gt;2,lt;6",
}

var gameJoin = message{
	receive: true,
	domain:  "game",
	command: "join",
	param:   "*|required,int",
}

var gameSetStone = message{
	receive: true,
	domain:  "game",
	command: "setStone",
	param:   "*|required,int,gt:0,lt:8",
}
