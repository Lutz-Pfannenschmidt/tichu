package websocket

import (
	"fmt"
	"net"
)

type WebsocketServer struct {
	host string
	port string
}

func (ws *WebsocketServer) runBroadcaster() {
	broadcaster, err := net.Listen("tcp", ws.host+":"+ws.port+"/brc")
	if err != nil {
		panic("Error starting the broadcasting socket server")
	}
	defer broadcaster.Close()

	fmt.Println("Broadcasting socket server Running...")
	fmt.Println("Listening on " + ws.host + ":" + ws.port + "/brc")

	for {
		connection, err := broadcaster.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("client connected")
		go runBroadcast(connection)
	}
}

func (ws *WebsocketServer) runReceiver() {
	receiver, err := net.Listen("tcp", ws.host+":"+ws.port+"/rec")
	if err != nil {
		panic("Error starting the receiving socket server")
	}
	defer receiver.Close()

	fmt.Println("Receiving socket server Running...")
	fmt.Println("Listening on " + ws.host + ":" + ws.port + "/rec")

	for {
		connection, err := receiver.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("client connected")
		go runReceiver(connection)
	}
}

func (ws *WebsocketServer) start() {
	go ws.runBroadcaster()
	go ws.runReceiver()
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}
