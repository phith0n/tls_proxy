package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	localAddr := flag.String("l", "", "Local address to listen on (format: IP:Port)")
	remoteAddr := flag.String("r", "", "Remote upstream address (format: IP:Port)")
	certFile := flag.String("c", "", "Path to the TLS certificate file")
	keyFile := flag.String("k", "", "Path to the TLS private key file")
	flag.Parse()

	if *localAddr == "" || *remoteAddr == "" || *certFile == "" || *keyFile == "" {
		fmt.Println("lack of required options")
		flag.Usage()
		return
	}

	run(*localAddr, *remoteAddr, *certFile, *keyFile)
}

func run(localAddr, remoteAddr, certPath, keyPath string) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp", localAddr, config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer listener.Close()
	log.Println("tls proxy started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		log.Println("new connection from", conn.RemoteAddr())
		go handleConnection(conn, remoteAddr)
	}
}

func handleConnection(src net.Conn, remoteAddr string) {
	defer src.Close()

	dest, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	defer dest.Close()

	go io.Copy(dest, src)
	io.Copy(src, dest)
}
