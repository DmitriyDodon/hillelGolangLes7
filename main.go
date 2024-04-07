package main

import (
	"flag"
	serv "les7/server"
	log "github.com/sirupsen/logrus"
	"slices"
)

const serverPostDefault = 3421
const typeServerToRunDefult = "http"

func main() {

	serverPort := flag.Int("port", serverPostDefault, "Server port")
	typeServerToRun := flag.String("typeServerToRun", typeServerToRunDefult, "Server type")

	flag.Parse()

	var server serv.Server

	validServerTypes := []string{"http", "tcp"}

	if !slices.Contains(validServerTypes, *typeServerToRun) {
		log.Fatal("Wrong server type!")
	}

	if *typeServerToRun == "http" {
		server = serv.NewHttpServer(*serverPort)
	}
	if *typeServerToRun == "tcp" {
		server = serv.NewTcpServer(*serverPort)
	}

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
