package main

import (
	"log"
	"os"
)

type application struct {
	appName  string
	server   server
	debug    bool
	errLog   *log.Logger
	inforLog *log.Logger
}

type server struct {
	host string
	port string
	url  string
}

func main() {
	server := server{
		host: "localhost",
		port: "8000",
		url:  "http://localhost:8000",
	}

	app := &application{
		server:   server,
		appName:  "test1",
		debug:    true,
		inforLog: log.New(os.Stdout, "INFOR\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:   log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	if err := app.listenAndServer(); err != nil {
		log.Fatal(err)
	}

}
