package pkg

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"sync"
	"time"
)

var mutex = sync.Mutex{}

type Postgres struct {
	Db *pg.DB
}

var db *Postgres

func InitializeDb() (*Postgres, error) {
	if db != nil {
		return db, nil
	}
	mutex.Lock()
	defer mutex.Unlock()
	var con Postgres
	con.Db = pg.Connect(&pg.Options{
		Addr: fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
		User: config.Db.Username,
		Password: config.Db.Password,
		Database: config.Db.Database,
		IdleTimeout: 20 * time.Second,
		MaxRetries: 3,
		PoolSize: 5,
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			fmt.Println("connected to postgres db")
			return nil
		},
	})
	db = &con
	return db, nil
}

func (p *Postgres) Close()  {
	p.Db.Close()
}

func (p *Postgres) Select(query string) (interface{}, error) {
	type Beans struct {
		Id int
		SupplierId string
		DriverId string
		UpdatedAt string
		CreatedAt string
	}
	var res []Beans
	_, err := p.Db.Query(&res, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}