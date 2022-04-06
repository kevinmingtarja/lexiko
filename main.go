package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/miekg/dns"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	db, err := setupDatabase()
	if err != nil {
		return err
	}

	dnsSrv := dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	dnsSrv.Handler = &handler{db: db}

	if err := dnsSrv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}

	//httpSrv := newHTTPServer(db)

	return nil
}
