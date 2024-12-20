package matchmaking

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/go-redis/redis/v8"
)

func StartMatchingWorker(){
	ctx := context.BackGround

	for{
		players, err := redisClient.ZRangeWithScores(ctx, "matchmaking_queue",0,1).Result()
        if err!=nil{
			log.Println("Error fetching player from queue", err)
			time.Sleep(2 * time.Second) // //sleep and entry
			continue
		}
		if len(players)<2{
			log.Println("Not enough player ,waiting..")
			time.Sleep(2 * time.Second) // Example;
			continue
		}

		// Match players (for example, match first 2 players)
        matchPlayers(players[0],players[1]);

		rangeBy := &redis.ZRangeBy{
			Min: "0",    // Define the minimum score
			Max: "1000", // Define the maximum score
		}

		 

		// Remove matched players from queue
       // redisClient.ZRangeByScore(ctx, "matchmaking_queue","+inf","-inf")

		//sleep for a while before checking again

		time.Sleep(2 * time.Second)


	}

}

// MatchPlayers will handle the logic to start a match between two players


func matchPlayers(player1, player2 redis.Z){
	fmt.Println("Matching players:",player1.Member,"and",player2.Member)
	    // Create a new session, start game, etc.

}