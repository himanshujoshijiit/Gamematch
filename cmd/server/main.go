 package main 

 import(
	"log"
	"net/http"
	"Gamematch/internal/session"
	"Gamematch/internal/matchmaking"
 )
var redisClient *redis.Client

//Intialize the Redis  client
func init(){
	redisClient = redis.NewClient(@redis.Options{
      Addr:"localhost:6379",
	})

	//check for redis connection

	ctx := context.Background()
	_,err :=redisClient.Ping(ctx).Result()
	if err!=nil{
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

}

 func main(){
	// Initialize Redis connection (if needed)
    // Initialize WebSocket connections
    // Start matchmaking worker

	//health check route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this is healthcheck")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	//matchmaking add user
	http.HandleFunc("/matchmaking/add",func(w http.ResponseWriter,r* http.Request){
		log.Println("add user for matchmaking")
		matchmaking.AddPlayerToQueue(w,r)
	})
    // root level endpoints
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		log.Println("websocket received");
		session.HandleWebSocket(w,r) // what is the meaning of this
	})
	

	//http.HandleFunc("/matchmaking/add",matchmaking.AddPlayerToQueue);
	http.HandleFunc("/ws",session.HandleWebSocket)

	log.Println("Server is running at 8080..")

	if err := http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatalf("Server is failed to start %v",err)
	}
 }