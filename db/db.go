package db

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"bit-project/gateway/config"
	"bit-project/gateway/ent"
	"bit-project/gateway/ent/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	client *ent.Client
)

// Open create a new ent.Client
func Open(conf *config.Config) *ent.Client {
	once.Do(func() { // atomic, does not allow repeating
		client = connect(conf)
	})
	return client
}

func AutoMigrate() {
	log.WithField("component", "db").
		WithField("method", "AutoMigrate").
		Info("run auto migration")

	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func connectionURL(cfg *config.Config) (string, error) {
	switch cfg.DBAdapter {
	case "mysql":
		//v := &url.Values{}
		return "", errors.New("Not support mysql")
	case "postgres":
		var host string
		var port string
		var err error
		if strings.Contains(cfg.DBHost, ":") {
			host, port, err = net.SplitHostPort(cfg.DBHost)
			if err != nil {
				return "", err
			}
		} else {
			host = cfg.DBHost
			port = cfg.DBPort
		}
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			host, port, cfg.DBUsername, cfg.DBPassword, cfg.DBName, "disable", cfg.TimeZone,
		), nil
	}
	return "", errors.New("unknown adapter")
}

func connect(conf *config.Config) *ent.Client {
	var (
		dbUrl  string
		err    error
		client *ent.Client
	)

	logger := log.WithFields(log.Fields{
		"component": "db",
		"methond":   "connect",
	})

	if dbUrl, err = connectionURL(conf); err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		driver, err := sql.Open(dialect.Postgres, dbUrl)
		if err != nil {
			logger.Errorf("failed opening connection to postgres: %v", err)
			logger.Info("DB is not ready yet. Check again after 1 sec.")
			time.Sleep(1 * time.Second)
			continue
		}
		err = driver.DB().Ping()
		if err != nil {
			logger.Errorf("failed opening connection to postgres: %v", err)
			logger.Info("DB is not ready yet. Check again after 1 sec.")
			time.Sleep(1 * time.Second)
			continue
		}
		options := []ent.Option{ent.Driver(driver)}
		client = ent.NewClient(options...)
		logger.Infof("connect client - %v", client)
		logger.Info("Successfully connect to new database")
		break
	}

	return client
}
