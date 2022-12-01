package db

import (
	"testing"
)

func TestSave(t *testing.T) {
	mc := MongoClient{Uri: "mongodb://root:example@106.75.63.248:27017/"}
	data := make(map[string]Any)
	data["a"] = 1
	data["b"] = "abcd"
	mc.Save("test", "test", data)
}

func TestFindOne(t *testing.T) {
	mc := MongoClient{Uri: "mongodb://root:example@106.75.63.248:27017/"}
	rs, _ := mc.FindOne("test", "test", BD{{Key: "a", Value: 1}})
	t.Log(rs)
	names := mc.ListCollectionNames("tang")
	t.Log(names)
}
