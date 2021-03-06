package orsa_scholis.inflandi.server

import orsa_scholis.inflandi.data
import java.io.IOException
import java.net.InetSocketAddress
import java.nio.ByteBuffer
import java.nio.channels.SelectionKey
import java.nio.channels.Selector
import java.nio.channels.ServerSocketChannel
import java.nio.channels.SocketChannel

/**
 * @author Crunchify
 * @author pawegio
 */
@Throws(IOException::class)
fun simpleServerExample() {
  val selector = Selector.open()

  val channel = ServerSocketChannel.open()
  val address = InetSocketAddress("localhost", 4560)
  channel.apply {
    bind(address)
    configureBlocking(false)
    register(selector, channel.validOps())
  }

  while (true) {
    simpleServerLoop(selector, channel)
  }
}

fun simpleServerLoop(selector: Selector, channel: ServerSocketChannel) {
  log("Waiting")
  selector.select()

  val keys = selector.selectedKeys()
  val iterator = keys.iterator()
  while (iterator.hasNext()) {
    serverHasNext(iterator, channel, selector)
  }
}

private fun serverHasNext(
  iterator: MutableIterator<SelectionKey>,
  channel: ServerSocketChannel,
  selector: Selector
) {
  val key = iterator.next()
  if (key.isAcceptable) {
    channel.accept().apply {
      configureBlocking(false)
      register(selector, SelectionKey.OP_READ)
      log("Connection accepted: $localAddress")
    }
  } else if (key.isReadable) {
    val client = key.channel() as SocketChannel
    val buffer = ByteBuffer.allocate(256)
    client.read(buffer)
    val result = buffer.data()
    log("Message received: $result")

    if (result == "Close") {
      client.close()
      log("Connection closed")
    }
  }
  iterator.remove()
}

fun log(message: String) {
  println(message)
}
