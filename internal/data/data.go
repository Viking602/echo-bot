package data

import (
	"context"
	"echo/internal/data/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"os"
)

var (
	ProviderSet = wire.NewSet(NewData, NewBotRepo)
)

type Data struct {
	db *ent.Client
}

func NewData() (*Data, error) {
	client, err := ent.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return &Data{
		db: client,
	}, nil
}
