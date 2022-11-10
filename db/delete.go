package db

import "context"

func (mc *MongoClient) DeleteOne(db, coll string, filter Any, opts ...*DeleteOptions) (int64, error) {
	if c := mc.GetCollection(db, coll); c != nil {
		rs, err := c.DeleteOne(context.Background(), filter, opts...)
		return rs.DeletedCount, err
	}
	return 0, MgoError("Can not get `Collection`.", 5)
}

func (mc *MongoClient) DeleteMany(db, coll string, filter Any, opts ...*DeleteOptions) (int64, error) {
	if c := mc.GetCollection(db, coll); c != nil {
		rs, err := c.DeleteMany(context.Background(), filter, opts...)
		return rs.DeletedCount, err
	}
	return 0, MgoError("Can not get `Collection`.", 5)
}
