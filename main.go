package main

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Message string `json:"message"`
}

func main() {

	r := gin.Default()

	r.POST("/api/chat", func(ctx *gin.Context) {

		var req Request

		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if req.Message == "" {
			ctx.JSON(http.StatusUnprocessableEntity, "message is required")
			return
		}

		cmd := exec.Command("./chat_compilation.sh", "d")

		out, err := cmd.CombinedOutput()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		var resp map[string]interface{}

		err = json.Unmarshal(out, &resp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
		return
	})

	r.Run(":8909")

}
