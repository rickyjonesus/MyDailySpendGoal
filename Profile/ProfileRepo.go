package Profile

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetAll() []*Profile {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		return []*Profile{{FirstName: "Error"}}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return []*Profile{{FirstName: "Error"}}
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return []*Profile{{FirstName: "Error"}}
	}
	var results []*Profile

	collection := client.Database("DailySpend").Collection("Profile")

	//filter := bson.D{{"firstName", "Ricky"}}
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Profile
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	return results
}
