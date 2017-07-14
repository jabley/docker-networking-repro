package main_test

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestConnectingToAServer(t *testing.T) {

	server := newServer()
	l := newListener()

	errorChan := make(chan error)

	go func() {
		errorChan <- server.Serve(l)
	}()

	client := http.Client{}
	res, err := client.Get("http://" + l.Addr().String())

	if err != nil {
		t.Fatalf("Got an error: %v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("Got an error: %v", err)
	}

	if string(body) != "Ohai!" {
		t.Fatalf("Got body: %v", body)
	}

	server.Shutdown(context.Background())

	select {
	case err := <-errorChan:
		if err != http.ErrServerClosed {
			t.Fatalf("Got an error: %v", err)
		}
	}
}

func newListener() net.Listener {
	l, err := net.Listen("tcp", "")
	if err != nil {
		panic(err)
	}
	return l
}

func newServer() *http.Server {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Ohai!"))
	})

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      serveMux,
	}

	return server
}
