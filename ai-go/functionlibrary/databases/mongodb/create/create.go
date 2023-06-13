/*
@File   : create.go
@Author : pan
@Time   : 2023-06-13 15:30:31
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

func (m *MongodbConfig) MongodbConn() (*mongo.Client, error) {
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
	return session, nil
}

func DBConn(client *mongo.Client, database string, setobj string) *mongo.Collection {
	collection := client.Database(database).Collection(setobj)
	return collection
}

func main() {
	m := &MongodbConfig{
		Host:     "ip",
		Port:     "port",
		Username: "user",
		Password: "pass",
		DB:       "",
	}
	session, err := m.MongodbConn()
	if err != nil {
		fmt.Printf("mongodb conn err:%v", err)
	} else {
		fmt.Printf("mongodb conn successful:%v", session)
	}
	collection := DBConn(session, "database", "setobj")
	fmt.Println(collection)
}
