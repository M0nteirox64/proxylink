package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Message struct {
	Webhook string `json:"webhook"`
	Content interface{} `json:"content"`
}

var Msg *Message

func main() {
	r := gin.Default()
	client := resty.New()
	
	r.GET("/", func (c *gin.Context) {
		c.String(http.StatusOK, "[ :) ] Bem vindo à proxylink")
		c.String(http.StatusOK, "[ i ] O objetivo da proxylink é fazer com que possas interagir com uma webhook do Discord.")
		c.String(http.StatusOK, "[ USAGE ] Usa '/post' para enviar uma requisição POST.")
	})


	r.POST("/post", func (c *gin.Context) {
		var newMessage Message
		if err := c.ShouldBindJSON(&newMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid JSON.",
			})
			return
		}

		Msg = &newMessage
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"content": Msg.Content,
			}).
			Post(Msg.Webhook)

		if err != nil {
			fmt.Println(err)
		}

		status := resp.StatusCode()

		switch status {
		case 200:
			fmt.Println("[OK] 200")
			c.String(http.StatusOK, "[OK] 200 | -> Sucessfull request")
		case 401:
			fmt.Println("[!] 401")
			c.String(http.StatusOK, "[!] 401 | -> Unauthorised")
		case 404:
			fmt.Println("[!] 4041")
			c.String(http.StatusOK, "[!] 404 | -> Not found")

		}
	})
	r.Run()	
}	
