package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"lp3/internal/data"
	"lp3/internal/mail"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
)

type config struct {
	db   string
	port int
	env  string
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	cfg       config
	logger    *log.Logger
	model     *data.Model
	mailer    *mail.Mailer
	scheduler *gocron.Scheduler
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8000, "Port number")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | production)")

	flag.StringVar(&cfg.db, "db", "postgres://usermessages_k0v6_user:E7FQKLQ0MhPxlg4Syn8DnWWsMLwHQvVJ@dpg-cl7lsliuuipc73eij48g-a.singapore-postgres.render.com/usermessages_k0v6", "Database Connection String")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "smtp.gmail.com", "SMTP Host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP Port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "ed7f23147046f3", "SMTP Username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "nvmu vdjw vmfz dpla", "SMTP Password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "aamerasim45@gmail.com", "SMTP Sender")

	// Open Database Connection
	db, err := OpenDB(cfg.db)
	if err != nil {
		fmt.Printf("Could not initialize database: %v\n", err)
		return
	}
	defer db.Close()

	// Initialize Application struct
	app := &application{}
	app.cfg = cfg
	app.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app.model = data.NewModel(db)
	app.mailer = mail.InitMailer(cfg.smtp.sender, cfg.smtp.password, cfg.smtp.host)
	app.scheduler = gocron.NewScheduler(time.UTC)

	// Initialize Router
	app.logger.Println("Application succesfully Initialized")
	router := app.route()

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	// Run Cron Job
	go app.runScheduler()

	// Start Server
	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Fatalf("Unable to run server: %v\n", err)
	}
}

func OpenDB(db_dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", db_dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
