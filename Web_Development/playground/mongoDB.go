package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Suggest struct {
	Fname string `bson:"fname"`
	Lname string `bson:"lname"`
	Email string `bson:"email"`
	Sugg  string `bson:"sugg"`
}

func main() {
	//Establish a connection to MongoDB
	client, err := connection()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Do your MongoDb operations using the client

	//define a new suggestion
	newSugg := Suggest{
		Fname: "Tim",
		Lname: "Grissette",
		Email: "Chreiskop@gamil.com",
		Sugg:  "jfdoajfkaj adjfklajfdlas jdfakl'afj ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss",
	}

	deleteOneSubmission(client, "Suggestions", "submissions", newSugg)
	if err != nil {
		log.Fatal(err)
	}

	//Close the connection when youre done
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to MongoDV closed.")

}

func connection() (*mongo.Client, error) {
	// MongoDB connection string
	connectionString := "mongodb+srv://umbra:password1995@umbramortis.m70inkz.mongodb.net/?retryWrites=true&w=majority"

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

func insertOneSubmission(client *mongo.Client, database, collection string, newSugg Suggest) error {
	//Access a database and a collection
	db := client.Database(database)
	coll := db.Collection(collection)

	//insert one submission
	_, err := coll.InsertOne(context.TODO(), newSugg)
	if err != nil {
		return err
	}

	fmt.Println("Document inserted successfully")
	return nil
}

func findOneSubmission(client *mongo.Client, database, collection string, newSugg Suggest) error {
	//Access a database and a collection
	db := client.Database(database)
	coll := db.Collection(collection)

	//define a filter (in tis)
	filter := bson.M{"fname": "Same"}

	//Create a variable to store the result
	var result Suggest

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No matching documetn found.")
		return nil
	} else if err != nil {
		log.Println("Error finding document:", err)
		return err
	} else {
		fmt.Println("Found document:")
		fmt.Println("First Name:", result.Fname)
		fmt.Println("Last Name:", result.Lname)
		fmt.Println("Email:", result.Email)
		fmt.Println("Suggestions", result.Sugg)
	}
	return err
}

func updateOneSubmission(client *mongo.Client, database, collection string, newSugg Suggest) error {
	//Access a database and a collection
	db := client.Database(database)
	coll := db.Collection(collection)

	//define a filter (in tis)
	filter := bson.M{"fname": "Same"}
	update := bson.M{"$set": bson.M{"email": "newemail@example.com"}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
	} else {
		fmt.Println("No matching document found")
	}

	return nil
}

func deleteOneSubmission(client *mongo.Client, database, collection string, newSugg Suggest) error {
	//Access a database and a collection
	db := client.Database(database)
	coll := db.Collection(collection)

	//insert one submission
	_, err := coll.DeleteOne(context.TODO(), newSugg)
	if err != nil {
		return err
	}

	fmt.Println("Document deleted successfully")
	return nil
}
