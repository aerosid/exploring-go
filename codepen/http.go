package codepen

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Client struct {
	client *http.Client
	url    string
	logger *log.Logger
}

func NewClient() *Client {
	httpClient := &http.Client{}
	url := "http://www.ifonepay.com"
	logger := log.New(os.Stdout, "", log.Lshortfile)
	c := &Client{client: httpClient, url: url, logger: logger}
	return c
}

func (c *Client) Go() (status int, headers map[string][]string, body string, err error) {
	status = 500
	headers = nil
	body = ""
	err = nil

	resp, err := http.Get(c.url)
	if err == nil {
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		status = resp.StatusCode
		body = string(b)
		headers = resp.Header
	} else {
		c.logger.Println(err.Error())
	}

	return status, headers, body, err
}

func (c *Client) SetUrl(url string) *Client {
	c.url = url
	return c
}

type DefaultHandler struct {
	message string
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{message: "Hello World!"}
}

func (h *DefaultHandler) SetMessage(message string) *DefaultHandler {
	h.message = message
	return h
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, h.message)
}

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func NewServer() *Server {
	addr := ":8000"
	handler := NewDefaultHandler()
	httpServer := &http.Server{Addr: addr, Handler: handler}
	logger := log.New(os.Stdout, "", log.Lshortfile)
	srv := &Server{httpServer: httpServer, logger: logger}
	return srv
}

func (s *Server) SetHandler(handler http.Handler) *Server {
	return s
}

func (s *Server) SetHost(host string) *Server {
	return s
}

func (s *Server) SetPort(port string) *Server {
	return s
}

func (s *Server) Start(blockingStart bool) *Server {
	start := func() { s.httpServer.ListenAndServe() }
	onInterrupt := func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigint
		s.logger.Println(" ..." + sig.String() + " signal received")
		s.Stop()
	}
	s.logger.Println(" ...listening on " + s.httpServer.Addr)
	if blockingStart {
		go onInterrupt()
		start()
	} else {
		go start()
	}

	return s
}

func (s *Server) Stop() {
	s.logger.Println(" ...stopping")
	s.httpServer.Shutdown(context.Background())
}
