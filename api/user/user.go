package user

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"devops/pkgs/comm"
	"devops/pkgs/db"
	"devops/pkgs/req"

	"github.com/gofiber/fiber/v3"
)

func Init(api fiber.Router) {
	r := api.Group("/users")

	r.Get("/", GetUsers)
	r.Get("/:id", GetUser)
	r.Post("/insert", InsertUser)
	r.Post("/update/:id", UpdateUser)
	r.Post("/delete/:id", DeleteUser)
}

func GetUsers(c fiber.Ctx) error {
	var users []User
	sql := fmt.Sprintf("SELECT * FROM %s", TABLE_USERS)
	if err := db.New().Select(&users, sql); err != nil {
		return c.JSON(req.
			NewResp([]User{}).
			SetResult(false).
			SetMsg("select user err: %v", err),
		)
	}
	return c.JSON(req.NewResp(users))
}

func GetUser(c fiber.Ctx) error {
	var user User
	id, err := strconv.ParseInt(c.Params("id"), 10, 0)
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("user id %s err: %v", c.Params("id"), err),
		)
	}

	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", TABLE_USERS, id)
	if err := db.New().Get(&user, sql); err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("get user %d err: %v", id, err),
		)
	}
	return c.JSON(req.NewResp(user))
}

func InsertUser(c fiber.Ctx) error {
	u := new(User)
	if err := c.Bind().JSON(u); err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("parse user err: %v", err),
		)
	}
	keys := comm.StructKeys(u, comm.InsertTag)
	values := make([]string, 0, len(keys))
	for _, key := range keys {
		values = append(values, fmt.Sprintf(":%s", key))
	}

	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		TABLE_USERS,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	result, err := db.New().NamedExec(sql, u)
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("insert user err: %v", err),
		)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("get row affected count err: %v", err),
		)
	}
	return c.JSON(req.NewResp(count))
}

func UpdateUser(c fiber.Ctx) error {
	u := new(User)
	if err := c.Bind().JSON(u); err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("parse user err: %v", err),
		)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 0)
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("user id %s err: %v", c.Params("id"), err),
		)
	}

	keys := comm.StructKeys(u, comm.UpdateTag)
	fields := make([]string, 0, len(keys))
	for _, key := range keys {
		fields = append(fields, fmt.Sprintf("%s = :%s", key, key))
	}

	u.UpdatedAt = time.Now()
	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d", TABLE_USERS, strings.Join(fields, ", "), id)
	result, err := db.New().NamedExec(sql, u)
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("update user err: %v", err),
		)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("get row affected count err: %v", err),
		)
	}
	return c.JSON(req.NewResp(count))
}

func DeleteUser(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 0)
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("user id %s err: %v", c.Params("id"), err),
		)
	}

	sql := fmt.Sprintf("UPDATE %s SET deleted_at = $1 WHERE id = $2", TABLE_USERS)
	result, err := db.New().Exec(sql, time.Now(), id)
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("update user err: %v", err),
		)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return c.JSON(req.
			NewResp(nil).
			SetResult(false).
			SetMsg("get row affected count err: %v", err),
		)
	}
	return c.JSON(req.NewResp(count))
}
