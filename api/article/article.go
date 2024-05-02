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
	// Tags        string `json:"tags"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func GetArticles(c *fiber.Ctx) error {
	log.Info("GetArticles")
	var articles []Article
	database.DB.Raw("SELECT * FROM article").Scan(&articles)
	return c.JSON(articles)
}

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