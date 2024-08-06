package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Staring the package")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://apadmin:Ze9h4WyuMoK4NQBWHrrqHhX@10.169.177.146:27020,10.169.177.147:27020,10.169.177.57:27020/traiUcc?authSource=admin&replicaSet=truconn"))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	d, _ := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Printf("Name of db %v\n", d)
	db := client.Database("traiUcc")
	collection := db.Collection("ad-tokens")
	cursor, err := collection.Find(ctx, bson.M{"Entity_Class": "Rjio Employee"})
	if err != nil {
		log.Fatal(err)
	}
	var entites []bson.M
	if err = cursor.All(ctx, &entites); err != nil {
		log.Fatal(err)
	}
	s, _ := json.Marshal(*cursor)
	fmt.Println(entites, string(s))

}
