package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Aakash-0003/Go-React-Training/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func CreateConnection() {

	/* connect with client instance
	local mongo instance */
	clientNew, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	Client = clientNew
	ErrorHandling(err)

	/* connecting with timeout using context pkg of golang
	the context will stop after giving time */
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Client.Connect(ctx)
	ErrorHandling(err)
	log.Println("connection successfull")
	collection := Client.Database("HRManagement").Collection("users")
	fmt.Println("connection instance is ready : ", collection)

}
func AddUser(user *models.User) *mongo.InsertOneResult {
	collection := Client.Database("HRManagement").Collection("users")
	result, _ := collection.InsertOne(context.Background(), user)
	fmt.Println("inserted: ", result)
	return result
}

func FindUserByMail(data *string, user *models.User) *models.User {
	collection := Client.Database("HRManagement").Collection("users")
	collection.FindOne(context.Background(), bson.M{"email": data}).Decode(&user)
	fmt.Println("found ", user)
	return user
}
func UpdateRole(data *string, role *string, user *models.User) *mongo.UpdateResult {
	collection := Client.Database("HRManagement").Collection("users")
	result, err := collection.UpdateOne(context.Background(), bson.M{"email": data}, bson.M{"$set": bson.M{"role": role}})
	ErrorHandling(err)
	return result
}
func DeleteUser(data *string, user *models.User) *mongo.DeleteResult {
	collection := Client.Database("HRManagement").Collection("users")
	result, err := collection.DeleteOne(context.Background(), bson.M{"email": data})
	ErrorHandling(err)
	return result
}
func AddClockIn(attendance *models.Attendance) *mongo.InsertOneResult {
	collection := Client.Database("HRManagement").Collection("attendance")
	result, _ := collection.InsertOne(context.Background(), attendance)
	fmt.Println("inserted: ", result)
	return result

}

/* func UserData() */
func ErrorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
