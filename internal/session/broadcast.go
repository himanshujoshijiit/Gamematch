package session

import (
	"github.com/gorilla/websocket",
	"fmt",
	"net/http"
)

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request){
	conn, err := websocket.Upgrade(w,r, null)
	if err!=nil{
		log.println("failed to upgrade to websocket",err)
		return
	}
	defer.conn.close()

	for {
		_,msg, err:= conn.ReadMessage()
		if err != nil{
			if err!=nil{
				log.println("Error in reading message",err)
				return
			}
			log.println("received message",string(msg))

			//Handle incoming message (e.g., game move or chat message)
		}
	}
}