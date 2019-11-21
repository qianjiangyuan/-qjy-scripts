package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ModelBase struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Ctime   time.Time          `json:"ctime,omitempty" bson:"ctime,omitempty"`
	Mtime   time.Time          `json:"mtime,omitempty" bson:"mtime,omitempty"`
	Deleted bool               `json:"_deleted,omitempty" bson:"_deleted"`
	CuserID primitive.ObjectID `json:"cuserID" bson:"cuserID"`
}

type File struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	IsDir  bool   `json:isdir`
}
