package main

import (
	"log"

	"github.com/Euclid0192/social-blog-golang/internal/db"
	"github.com/Euclid0192/social-blog-golang/internal/env"
	"github.com/Euclid0192/social-blog-golang/internal/store"
	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		/// DB config
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	/// Storage (Repository Pattern)
	store := store.NewStorage(db)

	/// Main app
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
