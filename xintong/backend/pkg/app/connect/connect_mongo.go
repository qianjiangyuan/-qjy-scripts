package connect

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoConnector struct {
	*mongo.Database
}

var mongoConn *MongoConnector

func MongoConnect(conf *viper.Viper) *MongoConnector {
	mongoConf := conf.GetStringMapString("mongodb")
	host:=os.Getenv(mongoConf["host"])
	port:=os.Getenv(mongoConf["port"])
	uri:="mongodb://"+host+":"+port
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	fmt.Println("mongodb connect ok ...")

	mongoConn = &MongoConnector{
		client.Database(mongoConf["db"]),
	}
	return mongoConn
}

// CURD
// ADD ONE
func (c *MongoConnector) Add(collection string, doc interface{}, opts *options.InsertOneOptions) (interface{}, error) {
	ctx := context.Background()
	col := c.Collection(collection)

	rs, err := col.InsertOne(ctx, doc, opts)
	return rs, err
}

// GET ONE
func (c *MongoConnector) Get(collection string, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {
	ctx := context.Background()
	col := c.Collection(collection)

	doc := col.FindOne(ctx, filter, opts)
	return doc
}

// DEL ONE, same with Update
func (c *MongoConnector) Del(collection string, filter interface{}, updater interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx := context.Background()
	col := c.Collection(collection)

	rs, err := col.UpdateOne(ctx, filter, updater, opts)
	return rs, err
}

// UPDATE ONE
func (c *MongoConnector) Update(collection string, filter interface{}, updater interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx := context.Background()
	col := c.Collection(collection)

	rs, err := col.UpdateOne(ctx, filter, updater, opts)
	return rs, err
}

// BATCH WITH PAGE CONTROL
func (c *MongoConnector) Batch(collection string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	ctx := context.Background()
	col := c.Collection(collection)

	docs, err := col.Find(ctx, filter, opts)
	return docs, err
}

// COUNT
func (c *MongoConnector) Count(collection string, filter interface{}, opts *options.CountOptions) (int64, error) {
	ctx := context.Background()
	col := c.Collection(collection)

	count, err := col.CountDocuments(ctx, filter, opts)
	return count, err
}
