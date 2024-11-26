package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"goftp.io/server"
)

func main() {

	flag.Parse()

	if len(flag.Args()) < 4 {
		fmt.Println("usage: xjftpd [port] [username] [password] [directory]")
		fmt.Println("example: xjftpd 21 user pass /home/user/ftp")
		os.Exit(0)
	}

	PORT, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalf("Invalid port number: %v\n", err)
	}
	USERNAME := flag.Arg(1)
	PASSWORD := flag.Arg(2)
	DIR := flag.Arg(3)

	factory := &server.FileDriverFactory{
		RootPath: DIR,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory: factory,
		Port:    PORT,
		Auth:    &server.SimpleAuth{Name: USERNAME, Password: PASSWORD},
	}

	ftpServer := server.NewServer(opts)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting FTP server:", err)
	}
}
