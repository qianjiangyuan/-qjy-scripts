package dao

import (
	"context"
	"fmt"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseResult interface {
	PushData(*mongo.Cursor)
	GetModeName() string
	SetData(singleReult *mongo.SingleResult)
}
type Dao struct {
}

func (dao Dao) Insert(m interface{}, baseResult BaseResult) (interface{}, error) {
	collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())
	_, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (dao Dao) SelectById(id string, baseResult BaseResult) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{
		{"_id", _id},
	}
	projection := bson.D{
		// useless stub for now
		{"runner", 0},
	}
	fmt.Printf("id=" + id + "\n")
	collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())
	singleResult := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection))
	baseResult.SetData(singleResult)
	return err
}

func (dao Dao) Update(m interface{}, id primitive.ObjectID, baseResult BaseResult) (interface{}, error) {
	collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())
	updateStr := bson.M{"$set": m}
	filter := bson.D{
		{"_id", id},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, updateStr)
	if err != nil {
		return m, err
	}
	return m, err
}
func (dao Dao) Deletes(ids []primitive.ObjectID, baseResult BaseResult) error {
	//collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())
	//filter := bson.D{{"_id", bson.D{{"$in", ids}}}}
	//_, err := collection.DeleteMany(context.TODO(), filter)
	for _,id := range ids {
		collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())
		updateStr := bson.M{"$set":bson.M{"_deleted":true} }
		filter := bson.D{
			{"_id", id},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, updateStr)
		if err != nil {
			return  err
		}
	}

	return nil

}

func (dsDao Dao) GetSelectPageList(parasMap map[string]string, pageNum int64, pageSize int64, baseResult BaseResult) (int64,error) {
	//var query primitive.M
	name, _ := parasMap["name"]

	filter := bson.M{
		"_deleted": false,
		"name":     bson.M{"$regex": name},
	}

	projection := bson.D{
		{"runner", 0},
	}
	opt := options.Find()
	opt.SetSkip((pageNum - 1) * pageSize)
	opt.SetLimit(pageSize)
	opt.SetProjection(projection)
	opt.SetSort(bson.D{
		{"ctime", -1},
	})

	collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())

	cur, err := collection.Find(context.TODO(), filter, opt)

	if err != nil {
		return 0,err
	}
	for cur.Next(context.TODO()) {
		baseResult.PushData(cur)
	}
	total ,err:= collection.CountDocuments(context.TODO(), filter)

	return total,nil
}
