package main

import (
	_ "hwgo/5/blog_app/routers"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run("localhost:" + os.Getenv("httpport"))
}
