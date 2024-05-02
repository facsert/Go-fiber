package router

import (
	"sync"
	"github.com/gofiber/fiber/v2"

	"fibert/api/article"
)

var wg sync.WaitGroup
func InitRouter(app *fiber.App) {

	api := app.Group("/api/v1")
    wg.Add(1)
	// article
    InitArticleRouter(api)

	wg.Wait()
}

func InitArticleRouter(api fiber.Router) {
	router := api.Group("/articles")
	router.Get("/", article.GetArticles)
	router.Get("/:id", article.GetArticle)
	router.Post("/:id", article.CreateArticle)
	router.Put("/:id", article.UpdateArticle)
	router.Delete("/:id", article.DeleteArticle)
	wg.Done()
}