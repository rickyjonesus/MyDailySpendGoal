package transaction


import (
	mongodb "MyDailySpendGoal/MongoDB"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDailySpend(date time.Date) decimal {

	client := mongodb.GetConnection()

	collection := client.Database("Transactions").Collection("Profile")


	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
   		var result bson.M
   		err := cur.Decode(&result)
   		if err != nil { log.Fatal(err) }
		// do something with result....
	}
	if err := cur.Err(); err != nil {
  		log.Fatal(err)
	}


	return 25;


}