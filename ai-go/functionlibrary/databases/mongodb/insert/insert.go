/*
@File   : insert.go
@Author : pan
@Time   : 2023-06-13 15:35:52
*/
package main

import (
	"context"
	"fmt"

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

func InsertOne(ctx context.Context, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	doc := StudentInfo{Name: "pan", Age: 19, Score: 99}
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func InsertMany(ctx context.Context, collection *mongo.Collection) (*mongo.InsertManyResult, error) {
	docs := []interface{}{StudentInfo{Name: "pan", Age: 19, Score: 99}, FoodInfo{Sweet: 90.1, Spices: 90.8, Salty: 1000.3}}
	result, err := collection.InsertMany(ctx, docs)
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
