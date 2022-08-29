package apiv1

import (
	"github.com/geraldie900/todo-app/app/helper"
	"github.com/geraldie900/todo-app/app/model"
	"github.com/geraldie900/todo-app/app/service/servicev1"
	"github.com/geraldie900/todo-app/config/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func TodoHandler(app fiber.Router) {
	todoHandler := app.Group("/todo")

	todoHandler.Post("", func(c *fiber.Ctx) error {
		requestBody := new(model.Todo)

		err := helper.BodyParseAndValidateStruct(c, requestBody)
		if err != nil {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, err.Error()))
		}

		response := servicev1.CreateTodoService(requestBody)
		return c.Status(response.StatusCode).JSON(response.ResponseData)
	})

	todoHandler.Get("", func(c *fiber.Ctx) error {
		qp := model.HTTPQueryParameter{}
		todoQP := model.TodoHTTPQueryParameter{}

		if err := c.QueryParser(&qp); err != nil {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, err.Error()))
		}

		if err := c.QueryParser(&todoQP); err != nil {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, err.Error()))
		}

		response := servicev1.GetTodosService(qp, todoQP)
		return c.Status(response.StatusCode).JSON(response.ResponseData)
	})

	todoHandler.Get("/:id", func(c *fiber.Ctx) error {
		todoID := c.Params("id")
		if todoID == "" {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, nil))
		}
		todoIDInt, _ := strconv.Atoi(todoID)
		response := servicev1.GetTodoService(todoIDInt)
		return c.Status(response.StatusCode).JSON(response.ResponseData)
	})

	todoHandler.Put("/:id", func(c *fiber.Ctx) error {
		requestBody := new(model.Todo)

		err := helper.BodyParseAndValidateStruct(c, requestBody)
		if err != nil {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, err.Error()))
		}

		todoID := c.Params("id")
		if todoID == "" {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, nil))
		}

		todoIDInt, _ := strconv.Atoi(todoID)
		requestBody.ID = todoIDInt
		response := servicev1.UpdateTodoService(requestBody)
		return c.Status(response.StatusCode).JSON(response.ResponseData)
	})

	todoHandler.Delete("/:id", func(c *fiber.Ctx) error {
		todoID := c.Params("id")
		if todoID == "" {
			return c.Status(utils.HTTPBadRequest).JSON(utils.GenerateJsonResponse(false, "", nil, nil))
		}

		todo := new(model.Todo)
		todoIDInt, _ := strconv.Atoi(todoID)
		todo.ID = todoIDInt
		response := servicev1.DeleteTodoService(todo)

		return c.Status(response.StatusCode).JSON(response.ResponseData)
	})
}
