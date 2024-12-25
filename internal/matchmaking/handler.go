package matchmaking

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"encoding/json"
)

var redisClient *redis.Client

func InitRedisClient(){
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis address
	})
}

// AddPlayerToQueue adds a player to the matchmaking queue

func AddPlayerToQueue(w http.ResponseWriter,r *http.Request){
	var request struct{
		PlayerId string `json:"playerId"`
		Score float64 `json:"score"`
	}

	    // Decode the request
   err:= json.NewDecoder(r.Body).Decode(&request)
   if err!=nil{
	http.Error(w,"Invalid Request",http.StatusBadRequest)
	return
   }

       // Add player to the Redis queue
  if(redisClient!=nil){
	http.Error(w,"Redis client is not intialised",http.StatusInternalServerError);
  }

  err = addToQueue(redisClient,request.PlayerId,request.Score)
  if err != nil{
	http.Error(w,"failed to add player to queue",http.StatusInternalServerError)
	return
  }
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintf("Player %s added to queue with score %f", request.PlayerId, request.Score)))
 
}

// AddToQueue adds the player to the Redis sorted set


func addToQueue(client *redis.Client,playerId string, score float64) error{
	ctx := context.Background()
	return client.ZAdd(ctx,"matchmaking_queue",&redis.Z{
		Score: score,
		Member: playerId,
	}).Err()
}