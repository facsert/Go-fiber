package node

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	
	"panel/utils/database"
)

const (
	TABLE_NODE = "node"
)

func Init(api fiber.Router) {
    router := api.Group("/nodes")
	router.Get("/", GetNodes)
	router.Get("/:id", GetNode)
	router.Post("/", CreateNode)
	router.Put("/", UpdateNode)
	router.Delete("/:id", DeleteNode)
}

// @tags     Node
// @summary  Get nodes
// @Router   /nodes  [get]
func GetNodes(c fiber.Ctx) error {
	log.Info("Get all nodes")
	var nodes []map[string]any
	// database.DB.Table(TABLE_NODE).Find(&nodes)
	database.DB.Raw("SELECT * FROM " + TABLE_NODE).Scan(&nodes)
	return c.JSON(nodes)
}

// @tags     Node
// @summary  Get node by id
// @Param    id path string true "node id"
// @Router   /nodes/{id}  [get]
func GetNode(c fiber.Ctx) error {
	log.Info("Get node " + c.Params("id"))
	var node map[string]any
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TABLE_NODE)
	// database.DB.Table(TABLE_NODE).Where("id = ?", c.Params("id")).Find(&node)
	database.DB.Raw(sql, c.Params("id")).Scan(&node)
	return c.JSON(node)
}

func CreateNode(c fiber.Ctx) error {
	return nil
}

func UpdateNode(c fiber.Ctx) error {
	return nil
}

func DeleteNode(c fiber.Ctx) error {
	return nil
}