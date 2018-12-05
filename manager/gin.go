package manager

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogather/safemap"
	"net/http"
)

var (
	ClientList *safemap.SafeMap = safemap.New()
)

func Run(managerPort int) {
	r := gin.Default()

	r.GET("/clients", func(c *gin.Context) {
		c.JSON(http.StatusOK, ClientList.GetMap())
		return
	})

	err := r.Run(fmt.Sprintf(":%d", managerPort))
	if err != nil {
		fmt.Println("start manage server failed, err:", err.Error())
	}
}
