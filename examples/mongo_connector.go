package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type User struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

func main() {
	// 1. Establish a connection to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// 2. Access a collection
	collection := client.Database("dbtest").Collection("users")

	// 3. Query documents
	filter := bson.M{"username": "john_doe"}
	var user User
	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found user:", user)

	// 4. Perform atomic updates
	update := bson.M{"$set": bson.M{"email": "john.doe@example.com"}}
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User email updated.")

	// 5. Implement connection pooling
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetMaxPoolSize(20)
	pooledClient, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := pooledClient.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// ... Perform further operations with the pooledClient

	// 6. Add caching with Redis
	// Example using the `go-redis` package
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Set if Redis requires authentication
		DB:       0,  // Select the appropriate database
	})
	defer redisClient.Close()

	// Query MongoDB and cache the result
	key := "user:john_doe"
	val, err := redisClient.Get(context.TODO(), key).Result()
	if err != nil {
		if err == redis.Nil {
			// Cache miss, fetch data from MongoDB
			if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
				log.Fatal(err)
			}
			// Store the data in Redis cache
			redisClient.Set(context.TODO(), key, user, time.Minute)
			fmt.Println("Fetched user from MongoDB.")
		} else {
			log.Fatal(err)
		}
	} else {
		// Cache hit, use the cached value
		fmt.Println("Found user in Redis cache:", val)
	}

	// ... Additional code for testing, error handling, etc.
}
