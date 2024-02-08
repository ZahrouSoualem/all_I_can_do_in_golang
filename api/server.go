package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tester/db/sqlc"
	"github.com/tester/token"
	"github.com/tester/util"
)

type Server struct {
	store      db.Store
	config     util.Config
	address    string
	tokenMaker *token.JWTMaker
	router     *gin.Engine
}

func NewServer(store db.Store, address string, config util.Config) (*Server, error) {

	maker, err := token.NewJwtMaker(config.SECURITY_KEY)

	if err != nil {
		return nil, err
	}

	router := gin.Default()
	server := &Server{
		address:    address,
		store:      store,
		tokenMaker: maker,
		config:     config,
	}

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginuser)

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRouter.GET("/user/:id", server.getUser)
	authRouter.GET("/user/name", server.getUserByname)
	authRouter.GET("/users", server.getUsers)
	authRouter.DELETE("/deluser/:id", server.deleteUser)
	authRouter.PUT("/upuser", server.updateUser)
	authRouter.POST("/renewtoken", server.renewtoken)

	server.router = router

	return server, nil

}

func (s *Server) Start() error {
	return s.router.Run(s.address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
