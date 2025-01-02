package api

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	address string
	// db      *sql.DB
}

// func NewAPIServer(address string, db *sql.DB) *APIServer {
func NewAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
		// db:      db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()

	return router.Run()
}
