package main

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Request struct {
	Message string `json:"message"`
}

func main() {

	r := gin.Default()

	r.POST("/api/chat", func(ctx *gin.Context) {

		var req map[string]interface{}

		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		client := resty.New()
		client.SetBaseURL("http://localhost:11434/api")

		exec.Command("./ollama_serve.sh")

		resp, err := client.R().SetBody(req).Post("/chat")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		exec.Command("./ollama_kill.sh")

		var respJson map[string]interface{}

		err = json.Unmarshal(resp.Body(), respJson)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, respJson)
		return
	})

	r.Run(":8909")

}
