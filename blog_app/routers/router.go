package routers

import (
	"database/sql"
	"log"

	"hwgo/5/blog_app/controllers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

const (
	dsnBlogs = "gbblog:123456@tcp(localhost:3306)/gb_blog?charset=utf8"
)

func init() {
	db, err := sql.Open("mysql", dsnBlogs)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("db connected")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Router("/", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/post", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	}, "GET:Open")

	beego.Router("/write", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	}, "GET:Write")

	beego.Router("/create", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/edit", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	}, "GET:Edit")

	beego.Router("/edit/:id", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	}, "POST:EditPost")

	beego.Router("/delete/:id", &controllers.BlogsController{
		Controller: beego.Controller{},
		Db:         db,
	}, "POST:DeletePost")
}
