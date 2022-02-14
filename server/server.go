package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mgutz/ansi"
	"github.com/rbg-tum/tum-live-monitor/monitor"
	"github.com/rbg-tum/tum-live-monitor/web"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

var (
	// logging stuff:
	getf     = ansi.ColorFunc("green+h:white+h")
	postf    = ansi.ColorFunc("blue+b:white+h")
	defaultf = ansi.ColorFunc("black+h:white+h")

	http2xx = ansi.ColorFunc("green+h:white+h")
	http1xx = ansi.ColorFunc("blue+b:white+h")
	http3xx = ansi.ColorFunc("yellow+h:white+h")
	http4xx = ansi.ColorFunc("red+b:white+h")
	http5xx = ansi.ColorFunc("magenta+h:white+h")
)

type Server struct {
	m *monitor.Monitor
}

func NewServer(m *monitor.Monitor) *Server {
	return &Server{m: m}
}

func (s *Server) Run() {
	log.Info("Starting server")
	go s.m.Run()
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(cors.New(cors.Config{AllowAllOrigins: true, AllowMethods: []string{"GET", "POST", "OPTIONS"}}))
	web.Configure(router)
	api := router.Group("/api")
	api.Use(logger)
	api.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the TUM Live Monitor API",
		})
	})
	port := viper.GetInt("port")
	log.Infof("Listening on http://127.0.0.1:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	log.Fatal(err)
}
