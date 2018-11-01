package template

const GinMainFile = `package backend
import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"log"
)

var (
	version string
	buildAt    string // time value should be passed in when building
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load config here

	// init database here
	
	// init router
	router := gin.Default()

	//router.Use(gin.Recovery())

	api.Register(router)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("version: %v, build at: %v\n", version, buildAt)
	log.Printf("listen: %s\n", srv.Addr)
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}`

const GinApiRouterFile = `package api
import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// API v1 Router
	routerV1 := router.Group("/v1")
	routerV1.GET("/ping", v1.Ping)
}`

const GinPingFile = `package v1

import (
	"github.com/gin-gonic/gin"
)

type pingResp struct {
	ErrorCode int    `+ "`json:\"error_code\"`"+ "\n"+
	`Data      string `+ "`json:\"data\"`"+ "\n"+
`}

func Ping(c *gin.Context) {
	var (
		resp pingResp
	)
	resp.Data = "pong"
	c.JSON(http.StatusOK, resp)
}`
