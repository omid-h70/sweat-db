package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoHandler struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoHandler(cfg DBConfig) (NoSQL, error) {

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	//"%s://%s:%s@mongodb-primary,mongodb-secondary,mongodb-arbiter/?replicaSet=replicaset",
	uri := fmt.Sprintf(
		"%s://%s:%s",
		"mongodb",
		cfg.DBHost,
		cfg.DBPort,
	)
	opt := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Print("mongo connection failed")
		return mongoHandler{}, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print("pinging mongo failed")
		return mongoHandler{}, err
	}

	return mongoHandler{
		client: client,
		db:     client.Database(cfg.DBName),
	}, nil
}

func (m mongoHandler) Insert(ctx context.Context, collection string, data any) error {
	if _, err := m.db.Collection(collection).InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}

/*
func (a AccountNoSQL) UpdateBalance(ctx context.Context, ID domain.AccountID, balance domain.Money) error {
	var (
		query  = bson.M{"id": ID}
		update = bson.M{"$set": bson.M{"balance": balance}}
	)

	if err := a.db.Update(ctx, a.collectionName, query, update); err != nil {
		switch err {
		case mongo.ErrNilDocument:
			return errors.Wrap(domain.ErrAccountNotFound, "error updating account balance")
		default:
			return errors.Wrap(err, "error updating account balance")
		}
	}

	return nil
}
*/

func (m mongoHandler) Update(ctx context.Context, collection string, filter any, update any) error {
	if _, err := m.db.Collection(collection).UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}
