package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// a document example:
// bson.D{{"name", "pi"}, {"value", 3.14159}}
// bson.D{{"name", "e"}, {"value", 2.71828}}
// they mean
// {
// 	"name": "pi",
// 	"value": 3.14159
// }
// {
// 	"name": "e",
// 	"value": 2.71828
// }

func ConnectToDB(mongodbURL string) *mongo.Client {
	// set client options
	clientOptions := options.Client().ApplyURI(mongodbURL)
	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected to MongoDB.")

	return client

	// how to use ConnectToDB:
	// client := ConnectToDB("mongodb://localhost:27017")
	// defer DisconnectFromDB(client)
}

func DisconnectFromDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func GetCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)

	// how to use GetCollection:
	// collection := GetCollection(client, "test", "test")
}

func InsertOneDocument(collection *mongo.Collection, document interface{}) {
	insertResult, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// how to use InsertOneDocument:
	// InsertOneDocument(collection, bson.D{{"name", "pi"}, {"value", 3.14159}})
}

func InsertManyDocuments(collection *mongo.Collection, documents []interface{}) {
	insertManyResult, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// how to use InsertManyDocuments:
	// InsertManyDocuments(collection, []interface{}{
	// 	bson.D{{"name", "pi"}, {"value", 3.14159}},
	// 	bson.D{{"name", "e"}, {"value", 2.71828}},
	// })
}

func FindOneDocument(collection *mongo.Collection, filter bson.D) *mongo.SingleResult {
	return collection.FindOne(context.TODO(), filter)

	// how to use FindOneDocument:
	// singleResult := FindOneDocument(collection, bson.D{{"name", "pi"}})

	// how to use the result:
	// var result bson.D
	// err := singleResult.Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

}

func FindManyDocuments(collection *mongo.Collection, filter bson.D) (*mongo.Cursor, error) {
	return collection.Find(context.TODO(), filter)

	// how to use FindManyDocuments:
	// cursor, err := FindManyDocuments(collection, bson.D{{"name", "pi"}})

	// how to use the cursor:
	// defer cursor.Close(context.TODO())
	// for cursor.Next(context.TODO()) {
	// 	var result bson.D
	// 	err := cursor.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(result)
	// }
}

func UpdateOneDocument(collection *mongo.Collection, filter bson.D, update bson.D) {
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated the document: ", updateResult.UpsertedID)

	// how to use UpdateOneDocument:
	// UpdateOneDocument(collection, bson.D{{"name", "pi"}}, bson.D{{"$set", bson.D{{"value", 3.14159265359}}}})

	// what does $set mean?
	// it means to update the value of the field "value" to 3.14159265359
	// if the field "value" does not exist, it will be created
	// if the field "value" exists, it will be updated
	// what else parameters can be used?
	// $inc, $mul, $rename, $setOnInsert, $unset, $min, $max, $currentDate
	// $inc: increment the value of the field by a specified amount
	// $mul: multiply the value of the field by a specified amount
	// $rename: rename a field
	// $setOnInsert: set the value of a field if an update results in an insert of a document
	// $unset: remove the field
	// $min: only update the field if the specified value is less than the existing field value
	// $max: only update the field if the specified value is greater than the existing field value
	// $currentDate: set the value of a field to current date
}

func UpdateManyDocuments(collection *mongo.Collection, filter bson.D, update bson.D) {
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated multiple documents: ", updateResult.UpsertedID)

	// how to use UpdateManyDocuments:
	// UpdateManyDocuments(collection, bson.D{{"name", "pi"}}, bson.D{{"$set", bson.D{{"value", 3.14159265359}}}})
}

func DeleteOneDocument(collection *mongo.Collection, filter bson.D) {
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted the document: ", deleteResult.DeletedCount)

	// how to use DeleteOneDocument:
	// DeleteOneDocument(collection, bson.D{{"name", "pi"}})
}

func DeleteManyDocuments(collection *mongo.Collection, filter bson.D) {
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted multiple documents: ", deleteResult.DeletedCount)

	// how to use DeleteManyDocuments:
	// DeleteManyDocuments(collection, bson.D{{"name", "pi"}})
}

func DropCollection(collection *mongo.Collection) {
	err := collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collection dropped.")

	// how to use DropCollection:
	// DropCollection(collection)
}

func DropDatabase(client *mongo.Client, dbName string) {
	err := client.Database(dbName).Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database dropped.")

	// how to use DropDatabase:
	// DropDatabase(client, "test")
}

func ListDatabases(client *mongo.Client) {
	databases, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Databases: ", databases)

	// how to use ListDatabases:
	// ListDatabases(client)
}

func ListCollections(client *mongo.Client, dbName string) {
	collections, err := client.Database(dbName).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collections: ", collections)

	// how to use ListCollections:
	// ListCollections(client, "test")
}

func ListDocuments(collection *mongo.Collection, filter bson.D) {
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result bson.D
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// how to use ListDocuments:
	// ListDocuments(collection, bson.D{})
}
