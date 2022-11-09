package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type FindOneOptions = options.FindOneOptions
type Collection = mongo.Collection
type BD = bson.D
type BM = bson.M
type BE = bson.E
type Any = interface{}

type MongoClient struct {
	Uri    string
	Client *mongo.Client
}

type MGError struct {
	string
	int
}

func MgoError(msg string, code int) *MGError {
	return &MGError{msg, code}
}
func (me *MGError) Error() string {
	return me.string
}

func (mc *MongoClient) GetClient(uri string) (*mongo.Client, error) {
	flag := false
	if mc.Client != nil {
		flag = mc.Client.Ping(context.Background(), readpref.Primary()) == nil
	}
	if !flag {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		mc.Client = client
		return client, err
	}
	return mc.Client, nil
}

func (mc *MongoClient) GetCollection(db, coll string) *Collection {
	if client, err := mc.GetClient(mc.Uri); err == nil {
		return client.Database(db).Collection(coll)
	}
	return nil
}
