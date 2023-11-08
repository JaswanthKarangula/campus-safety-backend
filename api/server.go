package api

import (
	db "dronesaefty-backend/db/sqlc"
	"dronesaefty-backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Store
	//tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	//if err != nil {
	//	return nil, fmt.Errorf("cannot create token maker: %w", err)
	//}

	server := &Server{
		config: config,
		store:  store,
		//tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(cors.Default())

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return true //origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/h", HealthCheck)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func parseError(err error, ctx *gin.Context) {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation":
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
