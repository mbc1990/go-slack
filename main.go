package main

import "fmt"
import "encoding/json"
import "os"

import "github.com/rgamba/evtwebsocket"

type Configuration struct {
	Token string
}

var Conf Configuration = Configuration{}

func initConf() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func startClient() {
	conn := evtwebsocket.Conn{
		// Fires when the connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			fmt.Println("Connected!")
		},
		// Fires when a new message arrives from the server
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			fmt.Printf("New message: %s\n", msg)
		},
		// Fires when an error occurs and connection is closed
		OnError: func(err error) {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		},
		// Ping interval in secs (optional)
		PingIntervalSecs: 5,
		// Ping message to send (optional)
		PingMsg: []byte("PING"),
	}

	err := conn.Dial("ws://echo.websocket.org", "")
	if err != nil {
		panic(err)
	}
}

func main() {
	initConf()
	go startClient()

	// Block forever
	select {}
}
