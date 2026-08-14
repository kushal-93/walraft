package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewHttpServer(port string) *http.Server {
	httpsrv := newHttpServer()
	r := chi.NewRouter()

	r.Post("/", httpsrv.handleProduce)
	r.Get("/", httpsrv.handleConsume)

	return &http.Server{
		Addr:    port,
		Handler: r,
	}
}

type httpServer struct {
	Log *Log
}

func newHttpServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}
