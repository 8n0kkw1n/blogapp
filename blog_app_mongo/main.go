package main

import (
	_ "hwgo/5/blog_app_mongo/routers"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	if err := logs.SetLogger("file", `{"filename":"test.log"}`); err != nil {
		log.Print(err)
	}

	beego.Run("localhost:" + os.Getenv("httpport"))
}
