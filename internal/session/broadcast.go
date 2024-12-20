package session

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request){
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if err!=nil{
		//Println("failed to upgrade to websocket",err)
		return
	}
	defer conn.Close()

	for {
		_,msg, err:= conn.ReadMessage()
		if err != nil{
			if err!=nil{
				//Println("Error in reading message",err)
				return
			}
			//Println("received message",string(msg))

			//Handle incoming message (e.g., game move or chat message)
		}
	}
}