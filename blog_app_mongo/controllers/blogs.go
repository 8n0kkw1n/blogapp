package controllers

import (
	"context"
	"hwgo/5/blog_app_mongo/models"
	"log"
	"time"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Explorer - object db connect
type Explorer struct {
	Db     *mongo.Client
	DbName string
}

// BlogsController - operations about object beego and mongodb
type BlogsController struct {
	beego.Controller
	Explorer Explorer
}

// @Title GetAll
// @Description output of all objects
// @Success 200 {html}
// @Failure 500 :object is empty
// @router / [get]
func (c *BlogsController) Get() {
	posts, err := c.Explorer.getAllPosts()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, err := c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.Data["Posts"] = posts
	c.TplName = "index.tpl"
}

// @Title GetAllPosts
// @Description output all posts from the database
// @Success {object} models.Object
// @Failure error :empty object
func (e Explorer) getAllPosts() ([]models.Posts, error) {
	c := e.Db.Database(e.DbName).Collection("blogs")

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	posts := make([]models.Posts, 0, 1)
	if err := cur.All(context.Background(), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

// @Title Write
// @Description opens post entry form
// @Success 200 {html}
// @router /write [get]
func (c *BlogsController) Write() {
	c.TplName = "write.tpl"
}

// @Title CreatePosts
// @Description creating a new post
// @Param  Title      path   string true      "Заголовок блога"
// @Param  Content      path   string true      "Контент блога"
// @Param  Date      path   string true      "Дата сейчас"
// @Success 302 {html}
// @Failure 500
// @router /create [post]
func (c *BlogsController) Post() {
	posts := models.Posts{}
	if err := c.ParseForm(&posts); err != nil {
		beego.Error("Could parse form", err)
	} else {
		c.Data["Posts"] = posts
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	post := models.Posts{
		Title:   posts.Title,
		Content: posts.Content,
		Date:    date,
	}

	if err := c.Explorer.createPost(post); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, err := c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.Redirect("/", 302)
}

// @Title createPost
// @Description creating a post and adding it to the database
// @Param  post      path   models.Posts true      "Структура с полями блога"
// @Success {object} models.Object
// @Failure error :empty object
func (e Explorer) createPost(post models.Posts) error {
	c := e.Db.Database(e.DbName).Collection("blogs")
	insertOneRes, err := c.InsertOne(context.Background(), post)
	if err != nil {
		log.Println(insertOneRes, err)
	}
	return err
}

// @Title Open
// @Description open post
// @Param  blogTitle      path   string true      "ID блога по title"
// @Success {object} models.Object
// @Failure 500, 404
// @router /post [get]
func (c *BlogsController) Open() {
	blogTitle := c.Ctx.Request.URL.Query().Get("blog_title")
	if len(blogTitle) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, err := c.Ctx.ResponseWriter.Write([]byte(`empty blog_title`))
		if err != nil {
			log.Println(err)
		}
		return
	}

	post, err := c.Explorer.getPost(blogTitle)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		_, err := c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.Data["Post"] = post
	c.TplName = "post.tpl"
}

// @Title Edit
// @Description getting post form
// @Param  blogTitle      path   string true      "ID строка блога по title"
// @Success 200 {html}
// @Failure 500, 404
// @router /edit [get]
func (c *BlogsController) Edit() {
	blogTitle := c.Ctx.Request.URL.Query().Get("blog_title")
	if len(blogTitle) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, err := c.Ctx.ResponseWriter.Write([]byte(`empty blog_title`))
		if err != nil {
			log.Println(err)
		}
		return
	}

	post, err := c.Explorer.getPost(blogTitle)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		_, err := c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.Data["Post"] = post
	c.TplName = "edit.tpl"
}

// @Title getPost
// @Description getting one post by id to the database
// @Param  blogTitle     path   string true      "ID строка блога по title"
// @Success {object} models.Object
// @Failure error
func (e Explorer) getPost(blogTitle string) (models.Posts, error) {
	c := e.Db.Database(e.DbName).Collection("blogs")

	filter := bson.D{{Key: "title", Value: blogTitle}}

	res := c.FindOne(context.Background(), filter)

	post := new(models.Posts)
	if err := res.Decode(post); err != nil {
		return models.Posts{}, err
	}
	return *post, nil
}

// @Title EditPost
// @Description getting post form
// @Param  blogTitle      path   string true      "ID строка блога по title"
// @Param  Title      path   string true      "Заголовок блога"
// @Param  Сontent      path   string true      "Контент блога"
// @Param  Date      path   string true      "Дата сейчас"
// @Success 302 {html}
// @Failure 500, 404
// @router /blogedit/:blog_title [post]
func (c *BlogsController) EditPost() {
	blogTitle := c.Ctx.Input.Param(":blog_title")

	if len(blogTitle) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, err := c.Ctx.ResponseWriter.Write([]byte(`empty blog_title`))
		if err != nil {
			log.Println(err)
		}
		return
	}

	posts := models.Posts{}
	if err := c.ParseForm(&posts); err != nil {
		beego.Error("Could parse form", err)
	} else {
		c.Data["Posts"] = posts
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	post := models.Posts{
		Title:   posts.Title,
		Content: posts.Content,
		Date:    date,
	}

	if err := c.Explorer.updatePost(&post, blogTitle); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		_, err := c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.Redirect("/", 302)
}

// @Title updatePost
// @Description updating post by title in database
// @Param  blogTitle      path   string true      "ID строка блога по title"
// @Param  post      path   models.Object true      "Структура с полями блога"
// @Success {object} *UpdateResult
// @Failure error
func (e Explorer) updatePost(post *models.Posts, blogTitle string) error {
	filter := bson.D{{Key: "title", Value: blogTitle}}

	update := createUpdates(*post)

	c := e.Db.Database(e.DbName).Collection("blogs")
	updateRes, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(updateRes, err)
	}
	return err
}

// @Title createUpdates
// @Description take from the base what we will update
// @Param  post      path   models.Object true      "Структура с полями блога"
// @Success {object} bson.D type D []E
// @Failure error
func createUpdates(post models.Posts) bson.D {
	update := bson.D{}
	if len(post.Title) != 0 {
		update = append(update, bson.E{Key: "title", Value: post.Title})
	}

	if len(post.Content) != 0 {
		update = append(update, bson.E{Key: "content", Value: post.Content})
	}

	if len(post.Date) != 0 {
		update = append(update, bson.E{Key: "date", Value: post.Date})
	}

	return bson.D{{"$set", update}}

}

// @Title DeletePost
// @Description deletion post
// @Param  blogTitle      path   string true      "ID строка блога по title"
// @Success 302 {html}
// @Failure 500, 404
// @router /delete/:blog_title [post]
func (c *BlogsController) DeletePost() {
	blogTitle := c.Ctx.Input.Param(":blog_title")

	if len(blogTitle) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, err := c.Ctx.ResponseWriter.Write([]byte(`empty blog_title`))
		if err != nil {
			log.Println(err)
		}
		return
	}

	if err := c.Explorer.deletePost(blogTitle); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		_, err := c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.Redirect("/", 302)
}

// @Title deletePost
// @Description deletion post of database
// @Param  blogTitle      path   string true      "ID строка блога по title"
// @Success {object} *DeleteResult
// @Failure error
func (e Explorer) deletePost(blogTitle string) error {
	filter := bson.D{{Key: "title", Value: blogTitle}}

	c := e.Db.Database(e.DbName).Collection("blogs")
	delRes, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(delRes, err)
	}
	return err
}
