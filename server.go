package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func connectionReader(conn *websocket.Conn) {
	/*for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
	*/for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Client:" + string(message))
		var ServerMessage string
		fmt.Println("Enter Message to Client")
		in := bufio.NewReader(os.Stdin)
		ServerMessage, err1 := in.ReadString('\n')
		if err := conn.WriteMessage(messageType, []byte(ServerMessage)); err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Sent!!")
		if err1 != nil {
			fmt.Println(ServerMessage)
		}

	}
}

/*func homepage(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
}*/
func websocketpage(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected")

	s := "Hii Client!"

	err = ws.WriteMessage(1, []byte(s))
	if err != nil {
		log.Println(err)
	}

	connectionReader(ws)
}

func setuproutes() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/ws", websocketpage)
}

func main() {
	fmt.Println("hello")
	setuproutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
