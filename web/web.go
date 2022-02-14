package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//go:embed build
var f embed.FS

var httpfs http.FileSystem

func Configure(r *gin.Engine) {
	httpfs = http.FS(f)

	r.GET("/", index)
	r.GET("/dist/:f", static)
}

func static(ctx *gin.Context) {
	log.Println(ctx.Request.RequestURI)
	err := fileFromFs(ctx, ctx.Request.RequestURI)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	}
}

func index(ctx *gin.Context) {
	err := fileFromFs(ctx, "/index.html")
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	}
}

func fileFromFs(ctx *gin.Context, path string) error {
	indexf, err := httpfs.Open("build" + path)
	if err != nil {
		return err
	}
	defer indexf.Close()
	d, err := ioutil.ReadAll(indexf)
	if err != nil {
		return err
	}
	mimetype := "text/plain"
	if strings.HasSuffix(path, ".html") {
		mimetype = "text/html"
	} else if strings.HasSuffix(path, ".css") {
		mimetype = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		mimetype = "application/javascript"
	} else if strings.HasSuffix(path, ".svg") {
		mimetype = "image/svg+xml"
	} else {
		mimetype = http.DetectContentType(d)
	}
	ctx.Data(http.StatusOK, mimetype, d)
	return nil
}
