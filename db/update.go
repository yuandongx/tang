package db

import "context"

func (mc *MongoClient) UpdateOne(db, coll string, filter Any, update Any, opts ...*UpdateOptions) error {
	if coll := mc.GetCollection(db, coll); coll == nil {
		_, err := coll.UpdateOne(context.Background(), filter, update, opts...)
		return err
	}
	return MgoError("Can not get `Collection`.", 4)
}
func (mc *MongoClient) UpdateMany(db, coll string, filter Any, update Any, opts ...*UpdateOptions) error {
	if coll := mc.GetCollection(db, coll); coll == nil {
		_, err := coll.UpdateMany(context.Background(), filter, update, opts...)
		return err
	}
	return MgoError("Can not get `Collection`.", 4)
}
