package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	grpcserver "github.com/polapolo/omdbapp/app/grpc"
	httpserver "github.com/polapolo/omdbapp/app/http"
)

func main() {
	var wait time.Duration

	// init log
	initLog()

	// init db
	db := initDB()
	defer db.Close()

	// HTTP Server
	httpServer := httpserver.InitHTTPServer(db)
	go func() {
		log.Println("[omdbapp][HTTP] Served on " + httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// GRPC Server
	grpcServer := grpcserver.InitGRPCServer(db)
	go func() {
		lis, err := net.Listen("tcp", `:9999`)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Println("[omdbapp][GRPC] Served on " + ":9999")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	httpServer.Shutdown(ctx)
	grpcServer.GracefulStop()
	// Optionally, you could run httpServer.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func initLog() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func initDB() *sqlx.DB {
	// create database if not exists
	db1, err := sqlx.Open("mysql", "root:supersecretpassword@tcp(db:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	db1.Exec(`CREATE DATABASE IF NOT EXISTS omdb;`)
	db1.Close()

	// connect to database
	db, err := sqlx.Connect("mysql", "root:supersecretpassword@tcp(db:3306)/omdb")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// migrate database
	migrateDB(db)

	return db
}

func migrateDB(db *sqlx.DB) {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///omdbapp/db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	m.Up()
}
