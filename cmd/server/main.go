 package main 

 import(
	"log"
	"net/http"
	"Gamematch/internal/matchmaking"
	"Gamematch/internal/session"
 )

 func main(){
	// Initialize Redis connection (if needed)
    // Initialize WebSocket connections
    // Start matchmaking worker

	http.HandleFunc("/matchmaking/add",matchmaking.AddPlayerToQueue);
	http.HandleFunc("/ws",session.HandleWebSocket)

	log.Println("Server is running at 8080..")

	if err := http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatalf("Server is failed to start %v",err)
	}
 }