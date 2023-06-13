/*
@File   : update.go
@Author : pan
@Time   : 2023-06-13 15:46:48
*/
package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbConfig struct {
	Host     string
	Port     string
	Password string
	Username string
	DB       string
}

func (m *MongodbConfig) MongodbConn(database string, setobj string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v@%v:%v", m.Username, m.Password, m.Host, m.Port))
	fmt.Println(clientOptions)
	session, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = session.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := session.Database(database).Collection(setobj)
	return collection, nil
}

type StudentInfo struct {
	Name  string
	Age   int
	Score int
}

type FoodInfo struct {
	Sweet  float64
	Spices float64
	Salty  float64
}

// 暂没验证
func UpdateOne(ctx context.Context, collection *mongo.Collection) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "Name", Value: "pan"}}                   // 可多个对象
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "Score", Value: 100}}}} // 可多个对象
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// 暂没验证
func UpdateMany(ctx context.Context, collection *mongo.Collection) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "Name", Value: "pan"}}                   // 可多个对象
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "Score", Value: 100}}}} // 可多个对象
	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func main() {
	m := &MongodbConfig{
		Host:     "ip",
		Port:     "port",
		Username: "user",
		Password: "pass",
		DB:       "",
	}
	session, err := m.MongodbConn("database", "setobj")
	if err != nil {
		fmt.Printf("mongodb conn err:%v", err)
	} else {
		fmt.Printf("mongodb conn successful:%v", session)
	}
}
