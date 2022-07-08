package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName  = "field_test"
	colName = "user"
)

type mongoClient struct {
	cli *mongo.Client
}

func newClient(ctx context.Context) (*mongoClient, error) {
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &mongoClient{client}, nil
}

func (m *mongoClient) GetUser(ctx context.Context, id string) (User, error) {
	col := m.cli.Database(dbName).Collection(colName)
	u := User{}
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (m *mongoClient) UpdateUser(ctx context.Context, u User) error {
	col := m.cli.Database(dbName).Collection(colName)
	opt := options.Update().SetUpsert(true)
	_, err := col.UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": &u,
	}, opt)
	if err != nil {
		return err
	}
	return nil

}
