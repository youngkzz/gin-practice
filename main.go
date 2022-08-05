package main

import (
	"github.com/gin-gonic/gin"

	"github.com/weblogin/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run()
}
