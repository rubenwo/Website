package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBConn ...
type DBConn struct {
	client      *mongo.Client
	collections map[string]*mongo.Collection
}

//InitializeConnection ...
func InitializeConnection(maxRetry int, addr string) (*DBConn, error) {
	var dbc *DBConn
	var err error
	for i := 0; i < maxRetry; i++ {
		dbc, err = connect(addr)
		if err == nil {
			return dbc, nil
		}
	}
	return nil, errors.New("couldn't get a connection to the database")
}

func connect(addr string) (*DBConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &DBConn{
		client:      client,
		collections: make(map[string]*mongo.Collection),
	}, nil
}
