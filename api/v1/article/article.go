package article

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lib/pq"
	
	"fibert/lib/database"
)

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Tags        pq.StringArray `json:"tags" gorm:"type: varchar(50)[]"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func Init(api fiber.Router) {
	router := api.Group("/articles")
	router.Get("/", GetArticles)
	router.Get("/:id", GetArticle)
	router.Post("/:id", CreateArticle)
	router.Put("/:id", UpdateArticle)
	router.Delete("/:id", DeleteArticle)
}

// @tags     Article
// @summary  Get articles
// @Router   /articles  [get]
func GetArticles(c *fiber.Ctx) error {
	log.Info("GetArticles")
	var articles []Article
	database.DB.Raw("SELECT * FROM article").Scan(&articles)
	return c.JSON(articles)
}

// @Tags    Article
// @summary Get article by id
// @Param   id path string true "article id"
// @Router  /articles/{id}  [get]
func GetArticle(c *fiber.Ctx) error {
	log.Info("GetArticle")
	var article Article
	database.DB.Raw("SELECT * FROM article WHERE id = ?", c.Params("id")).Scan(&article)
	return c.JSON(article)
}

func CreateArticle(c *fiber.Ctx) error {
	log.Info("CreateArticle")
	return nil
}

func UpdateArticle(c *fiber.Ctx) error {
	log.Info("UpdateArticles")
	return nil
}

func DeleteArticle(c *fiber.Ctx) error {
	log.Info("DeleteArticles")
	return nil
}