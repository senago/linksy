package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"

	"go.uber.org/zap"

	"github.com/senago/linksy/internal/api"
)

const (
	CONFIG_PATH     = "config.yaml"
	DEFAULT_ADDRESS = "0.0.0.0"
	DEFAULT_PORT    = "8080"
)

func main() {
	// -------------------- Set up viper -------------------- //

	viper.SetConfigFile(CONFIG_PATH)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read the config file: %s\n", err)
	}

	viper.SetDefault("service.bind.address", DEFAULT_ADDRESS)
	viper.SetDefault("service.bind.port", DEFAULT_PORT)

	// -------------------- Set up logging -------------------- //

	zlogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to set up the logger: %s\n", err)
	}
	defer zlogger.Sync()

	logger := zlogger.Sugar()

	// -------------------- Set up database -------------------- //

	log.Println(viper.GetString("db.connection_string"))

	dbPool, err := pgxpool.Connect(context.Background(), viper.GetString("db.connection_string"))
	if err != nil {
		log.Fatalf("unable to connect to the database: %s", err)
	}
	defer dbPool.Close()

	// -------------------- Set up service -------------------- //

	svc, err := api.NewAPIService(logger, dbPool)
	if err != nil {
		log.Fatalf("error creating service instance: %s", err)
	}

	go svc.Serve(viper.GetString("service.bind.address") + ":" + viper.GetString("service.bind.port"))

	// -------------------- Listen for INT signal -------------------- //

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(viper.GetInt("service.shutdown_timeout")),
	)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
