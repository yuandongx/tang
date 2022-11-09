package db

import (
	"context"
)

func (mc *MongoClient) FindOne(db, coll string, filter Any, options ...*FindOneOptions) (rs BM, er error) {
	if coll := mc.GetCollection(db, coll); coll != nil {
		rs = make(BM)
		er = coll.FindOne(context.Background(), filter, options...).Decode(rs)
		return
	}
	return nil, MgoError("Can not get collection.", 3)
}
