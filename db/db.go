package db

import (
	"context"
	"time"
)

type DBConfig struct {
	DBName  string
	Timeout time.Duration
	DBHost  string
	DBPort  string
}

type Store interface {
	Insert(ctx context.Context, collection string, data any) error
	Update(ctx context.Context, collection string, filter any, update any) error
	//FindAll(context.Context, string, any, any) error
	//FindOne(context.Context, string, any, any, any) error
	//StartSession() (Session, error)
}
