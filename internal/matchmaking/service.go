package matchmaking

import (
	"context",
	"fmt",
	"log",
	"time",
	"github.com/go-redis/redis/v8"
)

func StartMatchingWorker(){
	ctx := context.BackGround()

	for{
		players, err := redisClient.ZRangeWithScores(ctx, "matchmaking_queue",0,1).Result()
        if err!=nil{
			log.println("Error fetching player from queue", err)
			time.sleep(1*time.second)//sleep and entry
			continue
		}
		if len(player)<2{
			log.println("Not enough player ,waiting..")
			time.sleep(1*time.second);
			continue
		}

		// Match players (for example, match first 2 players)
        matchPlayers(players[0],players[1]);

		// Remove matched players from queue
        redisClient.ZRangeByScores(ctx, "matchmaking_queue","+inf","-inf")

		//sleep for a while before checking again

		time.sleep(1*time.second)

	}

}

// MatchPlayers will handle the logic to start a match between two players


func matchPlayers(player1, player2 redis.Z){
	fmt.println("Matching players:",player1.Member,"and",player2.Member)
	    // Create a new session, start game, etc.

}