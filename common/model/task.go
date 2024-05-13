package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
const TaskCollection = "task"

//
type Task struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty"`
	Title		string				`bson:"title"`
	Description	string        		`bson:"description"`
	Price		int32				`bson:"price"`
	Quantity	int32				`bson:"quantity"`
}