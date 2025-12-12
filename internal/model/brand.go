package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Brand struct {
	Id      bson.ObjectID `bson:"_id, omitempty"`
	Name    string        `bson:"brand_name"`
	Devices []Device      `bson:"devices"`
}
