package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/golang-toko-online/routers"
)

func main() {

	r := gin.Default()
	routers.InitRouter(r)
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
