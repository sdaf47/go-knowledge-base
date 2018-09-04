package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/sdaf47/go-knowledge-base/shortlink/linkshort/transport"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	service := transport.NewGRPCClient(conn)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	formHandler := func(c *gin.Context) {
		link := c.PostForm("link")
		var code string
		var err error

		if link != "" {
			code, err = service.Encode(link)
			if err != nil {
				return
			}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"link":       link,
			"short_link": "/_/"+code,
		})
	}

	router.GET("/", formHandler)
	router.POST("/", formHandler)

	router.GET("/_/:code", func(c *gin.Context) {
		code := c.Param("code")

		link, err := service.Decode(code)
		if err != nil {
			return
		}

		if link != "" {
			c.Redirect(http.StatusMovedPermanently, link)
			return
		}
		// error
		return
	})

	router.Run(":8000")
}
