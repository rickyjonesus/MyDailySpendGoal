package profile

import (
	mongodb "MyDailySpendGoal/MongoDB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOne(emailAddress string) Profile {
	log.Println("RepoGetOne: " + emailAddress)
	client := mongodb.GetConnection()

	collection := client.Database("DailySpend").Collection("Profile")

	var filter2 = bson.D{{
		"emailAddress",
		bson.D{{
			"$in",
			bson.A{emailAddress},
		}},
	}}

	var result Profile

	var err = collection.FindOne(context.TODO(), filter2).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetAll() []*Profile {

	client := mongodb.GetConnection()

	var results []*Profile

	collection := client.Database("DailySpend").Collection("Profile")

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

func Add(profileToAdd Profile) int {

	log.Println("ProfileRepo.Add")

	client := mongodb.GetConnection()

	collection := client.Database("DailySpend").Collection("Profile")

	_, err := collection.InsertOne(context.TODO(), profileToAdd)
	if err != nil {
		log.Fatal(err)
	}

	client.Disconnect(context.TODO())
	return 5

}
