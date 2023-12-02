package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Migrate() {
	var result, err = Collection.UpdateMany(context.TODO(), bson.D{}, bson.D{
		{"$currentDate", bson.D{{"timestamp", bson.D{{"$type", "date"}}}}},
	})

	if err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}

	fmt.Println(result)
}
