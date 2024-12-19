package session

import "fmt"

// Session holds details about a multiplayer session

type session struct{
	ID string
	Player []string
}

// ActiveSessions keeps track of active game sessions

var ActiveSessions = make(map[string]*Session)

// CreateSession creates a new multiplayer session

func CreateSession(player1,player2 string) *session{
	sessionId := fmt.Sprintf("session-d%",len(ActiveSessions)+1)

	session := &session{
		Id: sessionId,
		Players :[]string{player1,player2},
	}

	ActiveSessions[sessionId] = session
	return session
}

// HandleWebSocket will handle the real-time communication for multiplayer sessions

func HandleWebSocket(w http.ResponseWriter,r *http.request){
	    // Upgrade HTTP connection to WebSocket and handle communication

}