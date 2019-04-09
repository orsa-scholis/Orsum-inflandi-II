import * as React from 'react';
import * as PropTypes from 'prop-types';
import { AppBar, Grid, IconButton, Toolbar, Typography, withStyles } from '@material-ui/core';
import PlusIcon from '@material-ui/icons/Add';
import LobbyScreenStyles from './LobbyScreenStyles';
import { withSnackbar, withSnackbarProps } from 'notistack';
import GameList from '../../components/GameList/GameList';
import Game from '../../models/Game/Game';

interface LobbyScreenProps extends withSnackbarProps {
  classes: any;
}

class LobbyScreen extends React.Component<LobbyScreenProps> {
  static propTypes = {
    classes: PropTypes.object
  };

  render() {
    const { classes, enqueueSnackbar } = this.props;

    const myGameList = [
      new Game('First game'),
      new Game('Second game'),
    ];

    return (
      <Grid container spacing={16}>
        <Grid item xs={12}>
          <AppBar position='static' color='default'>
            <Toolbar>
              <Typography variant='h6' color='textPrimary'>
                Lobby
              </Typography>
              <div className={classes.grow}/>
              <IconButton color='inherit' onClick={() => enqueueSnackbar('I am adding a game', { variant: 'success' })}>
                <PlusIcon/>
              </IconButton>
            </Toolbar>
          </AppBar>
        </Grid>
        <Grid item xs={12}>
          <GameList games={myGameList}/>
        </Grid>
      </Grid>
    );
  }
}

export default withSnackbar(withStyles(LobbyScreenStyles)(LobbyScreen));
