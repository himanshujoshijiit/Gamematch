 package main 

 import(
	"log",
	"net/http",
	"go-matchmaking/internal/matchmaking",
	"go-matchmaking/internal/session"
 )

 func main(){
	// Initialize Redis connection (if needed)
    // Initialize WebSocket connections
    // Start matchmaking worker

	http.HandleFunc("/matchmaking/add",matchmaking.AddPlayerToQueue);
	http.HandleFunc("/ws",session.HandleWebSocket)

	log.println("Server is running at 8080..")

	if err := http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatalf("Server is failed to start %v",err)
	}
 }