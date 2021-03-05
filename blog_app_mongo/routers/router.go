package routers

import (
	"context"
	"hwgo/5/blog_app_mongo/controllers"
	"log"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "test"

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongo-db connected")

	if err = db.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	beego.Router("/", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	})

	beego.Router("/write", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	}, "GET:Write")

	beego.Router("/create", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	})

	beego.Router("/post", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	}, "GET:Open")

	beego.Router("/edit", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	}, "GET:Edit")

	beego.Router("/blogedit/:blog_title", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	}, "POST:EditPost")

	beego.Router("/delete/:blog_title", &controllers.BlogsController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	}, "POST:DeletePost")

}
