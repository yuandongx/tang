package db

import "context"

func (mc *MongoClient) SaveD(db, coll string, data BD) error {
	return mc.Save(db, coll, data)
}
func (mc *MongoClient) SaveM(db, coll string, data BM) error {
	return mc.Save(db, coll, data)
}
func (mc *MongoClient) Save(db, coll string, data Any) error {
	if collection := mc.GetCollection(db, coll); collection != nil {
		collection.InsertOne(context.Background(), data)
	}
	return MgoError("Can not get `Collection`.", 2)
}
func (mc *MongoClient) SaveMany(db, coll string, data []Any) error {
	if collection := mc.GetCollection(db, coll); collection != nil {
		collection.InsertOne(context.Background(), data)
	}
	return MgoError("Can not get `Collection`.", 2)
}
