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
	return nil, MgoError("Can not get collection.", CODE_FIND)
}

func (mc *MongoClient) FindMany(db, coll string, filter Any, opts ...*FindOptions) ([]BM, error) {
	if coll := mc.GetCollection(db, coll); coll != nil {
		var rs []BM
		if cur, err := coll.Find(context.Background(), filter, opts...); err == nil {
			err2 := cur.All(context.Background(), &rs)
			return rs, err2
		} else {
			return nil, err
		}
	}
	return nil, MgoError("Can not get collection.", CODE_FIND)
}
