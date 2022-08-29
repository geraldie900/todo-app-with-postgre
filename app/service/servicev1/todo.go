package servicev1

import (
	"encoding/json"
	"errors"
	"github.com/geraldie900/todo-app/app/model"
	"github.com/geraldie900/todo-app/config/database/postgresql"
	"github.com/geraldie900/todo-app/config/logger"
	"github.com/geraldie900/todo-app/config/utils"
	"gorm.io/gorm"
)

func CreateTodoService(todo *model.Todo) utils.Response {
	logging := logger.Logger{
		FunctionName: "CreateTodoService",
	}

	response := utils.Response{
		StatusCode: 200,
	}

	todo.Timestamp, _ = utils.TimestampNow(false, "")
	tx := postgresql.StartTransaction()
	err := tx.CreateData(todo)
	if err != nil {
		logging.LogError("create todo", err)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, err.Error())
		return response
	}
	err = tx.CommitTransaction()
	if err != nil {
		logging.LogError("commit create todo", err)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, err.Error())
		return response
	}

	response.ResponseData = utils.GenerateJsonResponse(true, "", todo, nil)

	return response
}

func GetTodosService(qp model.HTTPQueryParameter, todoQP model.TodoHTTPQueryParameter) utils.Response {
	logging := logger.Logger{
		FunctionName: "GetTodosService",
	}

	response := utils.Response{
		StatusCode: 200,
	}

	tx := postgresql.Start()

	todos := &[]model.Todo{}
	whereField := []string{}
	whereValue := []interface{}{}
	orderBy := []string{}

	var todoQPMap map[string]interface{}
	todoQPByte, _ := json.Marshal(todoQP)
	json.Unmarshal(todoQPByte, &todoQPMap)

	for k, v := range todoQPMap {
		whereField = append(whereField, k)
		whereValue = append(whereValue, v)
	}

	tx.WhereField = whereField
	tx.WhereValue = whereValue

	if qp.Limit != 0 {
		tx.Limit = qp.Limit
	}

	if qp.Offset != 0 {
		tx.Offset = qp.Offset
	}

	if len(qp.OrderBy) > 0 {
		if qp.Desc == "true" {
			qp.OrderBy += " desc"
		}
		orderBy = append(orderBy, qp.OrderBy)
		tx.OrderBy = orderBy
	}

	queryResult := tx.GetList(todos)
	if queryResult.Error != nil {
		logging.LogError("get todo", queryResult.Error)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, queryResult.Error.Error())
		return response
	}

	response.ResponseData = utils.GenerateJsonResponse(true, "", todos, nil)

	return response
}

func GetTodoService(todoID int) utils.Response {
	logging := logger.Logger{
		FunctionName: "GetTodoService",
	}

	response := utils.Response{
		StatusCode: 200,
	}

	tx := postgresql.Start()

	todo := new(model.Todo)
	todo.ID = todoID
	queryResult := tx.Get(todo)
	if queryResult.Error != nil {
		logging.LogError("get todo", queryResult.Error)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, queryResult.Error.Error())
		if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
			response.ResponseData = utils.GenerateJsonResponse(true, "", nil, queryResult.Error.Error())
		}
		return response
	}

	response.ResponseData = utils.GenerateJsonResponse(true, "", todo, nil)

	return response
}

func UpdateTodoService(todo *model.Todo) utils.Response {
	logging := logger.Logger{
		FunctionName: "UpdateTodoService",
	}

	response := utils.Response{
		StatusCode: 200,
	}

	tx := postgresql.StartTransaction()
	err := tx.UpdateData(todo)
	if err != nil {
		logging.LogError("update todo", err)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, err.Error())
		return response
	}
	err = tx.CommitTransaction()
	if err != nil {
		logging.LogError("commit update todo", err)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, err.Error())
		return response
	}

	response.ResponseData = utils.GenerateJsonResponse(true, "", todo, nil)

	return response
}

func DeleteTodoService(todo *model.Todo) utils.Response {
	logging := logger.Logger{
		FunctionName: "DeleteTodoService",
	}

	response := utils.Response{
		StatusCode: 200,
	}

	tx := postgresql.StartTransaction()
	err := tx.DeleteData(todo)
	if err != nil {
		logging.LogError("delete todo", err)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, err.Error())
		return response
	}
	err = tx.CommitTransaction()
	if err != nil {
		logging.LogError("commit delete todo", err)
		response.ResponseData = utils.GenerateJsonResponse(false, "", nil, err.Error())
		return response
	}

	response.ResponseData = utils.GenerateJsonResponse(true, "", todo, nil)

	return response
}
