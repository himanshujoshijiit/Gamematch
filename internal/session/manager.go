package session

import (
	"fmt"
)
 
// Session holds details about a multiplayer session

type Session struct {
    ID        string   // Unique identifier for the session
    PlayerIDs []string // List of player IDs in the session
}



// ActiveSessions keeps track of active game sessions

var ActiveSessions = make(map[string]*Session)

// CreateSession creates a new multiplayer session

func CreateSession(player1,player2 string) *Session{
	sessionId := fmt.Sprintf("session-d%",len(ActiveSessions)+1)

	session := &Session{
		ID: sessionId,
		PlayerIDs: []string{"player1", "player2"},
	}

	ActiveSessions[sessionId] = session
	return session
}

// HandleWebSocket will handle the real-time communication for multiplayer sessions

// func HandleWebSocket(w http.ResponseWriter,r *http.request){
// 	    // Upgrade HTTP connection to WebSocket and handle communication

// }