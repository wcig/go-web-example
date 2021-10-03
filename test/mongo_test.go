package test

import (
	"context"
	"fmt"
	"go-app/library/boot"
	mongox "go-app/library/mongo"
	"go-app/library/util/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initMongo() {
	boot.Register(&mongox.MongoStarter{})
	testBootRun()
}

func TestMongo(t *testing.T) {
	initMongo()

	err := mongox.GetClient().Ping(context.Background(), readpref.Primary())
	assert.Nil(t, err)
}

type LogRecord struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Msg      string             `json:"msg" bson:"msg"`
	CreateAt int64              `json:"create_at" bson:"create_at"`
}

func TestFirst(t *testing.T) {
	initMongo()

	var (
		dbName   = "test"
		collName = "log_record"
		coll     = mongox.GetClient().Database(dbName).Collection(collName)
	)

	// insert
	ir := &LogRecord{
		Id:       primitive.NewObjectID(),
		Msg:      "ok",
		CreateAt: time.Now().UnixNano() / 1e6,
	}
	insertResult, err := coll.InsertOne(context.Background(), ir)
	assert.Nil(t, err)
	fmt.Println("insert result", insertResult)
	fmt.Println("insert record:", json.ToJson(ir))

	// select
	var sr LogRecord
	filter := bson.M{"_id": ir.Id}
	err = coll.FindOne(context.Background(), filter).Decode(&sr)
	assert.Nil(t, err)
	fmt.Println("select record:", json.ToJson(sr))

	// update
	update := bson.M{"$set": bson.M{"msg": "hello"}}
	updateResult, err := coll.UpdateOne(context.Background(), filter, update)
	assert.Nil(t, err)
	fmt.Println("update result:", updateResult)

	// delete
	deleteResult, err := coll.DeleteOne(context.Background(), filter)
	assert.Nil(t, err)
	fmt.Println("delete result:", deleteResult)
}
