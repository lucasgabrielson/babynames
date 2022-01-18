package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"./routers/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(":5000"); err != nil {
		fmt.Println("startup service failed, err:#%v\n", err)
	}
}
