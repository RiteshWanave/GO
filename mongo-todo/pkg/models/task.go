package models

import (
	"context"
	"time"

	"github.com/RiteshWanave/GO/mongo-todo/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type Task struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	TaskName string
	CreatedAt time.Time
	CompletedAt time.Time
	Done bool
}

var db *mongo.Client
var collection *mongo.Collection

func init () {
	config.Connect()
	db = config.GetDB()
	collection = db.Database("todo").Collection("tasks")
}

func Add (taskName string) string {
	if taskName == "" {
		println("Task name cannot be empty")
		return ""
	}
	task := Task {
		TaskName: taskName,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	res, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		println("Error inserting task: ", err.Error())
		return ""
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	println("Inserted task with id: ", id)
	return id
}

func Complete (id string) {
	if id == "" {
		return
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		println("Error converting string to object id: ", err.Error())
		return
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"done": true, "completedAt": time.Now()}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		println("Error updating task: ", err.Error())
		return
	}
	println("Task updated successfully")
}

func List () []Task {
	var tasks []Task
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		println("Error getting tasks: ", err.Error())
		return tasks
	}
	defer cur.Close(context.Background())
	// for cur.Next(context.Background()) {
	// 	var task Task
	// 	err := cur.Decode(&task)
	// 	if err != nil {
	// 		println("Error decoding task: ", err.Error())
	// 		return tasks
	// 	}
	// 	tasks = append(tasks, task)
	// }
	// if err := cur.Err(); err != nil {
	// 	println("Error getting tasks: ", err.Error())
	// 	return tasks
	// }
	// return tasks
	cur.All(context.Background(), &tasks)
	return tasks
}