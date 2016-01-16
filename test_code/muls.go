package main

import "muls/transpports"

func main() {

	server := transpports.NewUdpServer("127.0.0.1", 5678)

	server.Run()

}
