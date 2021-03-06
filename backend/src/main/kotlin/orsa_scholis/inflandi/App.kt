/*
 * This Kotlin source file was generated by the Gradle 'init' task.
 */
package orsa_scholis.inflandi

import orsa_scholis.inflandi.server.simpleServerExample

class App {
    val greeting: String
        get() {
            return "Hello world."
        }
}

fun main(args: Array<String>) {
    println(App().greeting)

    println("Server is starting...")
    simpleServerExample()
}
