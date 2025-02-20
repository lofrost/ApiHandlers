// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint   `json:"id,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
	Task   *string `json:"task,omitempty"`
	UserId *uint   `json:"user_id,omitempty"`
}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// UpdateTaskByIDJSONRequestBody defines body for UpdateTaskByID for application/json ContentType.
type UpdateTaskByIDJSONRequestBody = Task

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
	// Delete a task by ID
	// (DELETE /tasks/{task_id})
	DeleteTaskByID(ctx echo.Context, taskId uint) error
	// Update a task by ID
	// (PATCH /tasks/{task_id})
	UpdateTaskByID(ctx echo.Context, taskId uint) error
	// Get all user tasks by ID
	// (GET /user/{user_id}/tasks)
	GetTasksByUserID(ctx echo.Context, userId uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// DeleteTaskByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTaskByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "task_id" -------------
	var taskId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "task_id", runtime.ParamLocationPath, ctx.Param("task_id"), &taskId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter task_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTaskByID(ctx, taskId)
	return err
}

// UpdateTaskByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTaskByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "task_id" -------------
	var taskId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "task_id", runtime.ParamLocationPath, ctx.Param("task_id"), &taskId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter task_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTaskByID(ctx, taskId)
	return err
}

// GetTasksByUserID converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasksByUserID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "user_id" -------------
	var userId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, ctx.Param("user_id"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasksByUserID(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)
	router.DELETE(baseURL+"/tasks/:task_id", wrapper.DeleteTaskByID)
	router.PATCH(baseURL+"/tasks/:task_id", wrapper.UpdateTaskByID)
	router.GET(baseURL+"/user/:user_id/tasks", wrapper.GetTasksByUserID)

}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks201JSONResponse Task

func (response PostTasks201JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTaskByIDRequestObject struct {
	TaskId uint `json:"task_id"`
}

type DeleteTaskByIDResponseObject interface {
	VisitDeleteTaskByIDResponse(w http.ResponseWriter) error
}

type DeleteTaskByID204Response struct {
}

func (response DeleteTaskByID204Response) VisitDeleteTaskByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type UpdateTaskByIDRequestObject struct {
	TaskId uint `json:"task_id"`
	Body   *UpdateTaskByIDJSONRequestBody
}

type UpdateTaskByIDResponseObject interface {
	VisitUpdateTaskByIDResponse(w http.ResponseWriter) error
}

type UpdateTaskByID200JSONResponse Task

func (response UpdateTaskByID200JSONResponse) VisitUpdateTaskByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTasksByUserIDRequestObject struct {
	UserId uint `json:"user_id"`
}

type GetTasksByUserIDResponseObject interface {
	VisitGetTasksByUserIDResponse(w http.ResponseWriter) error
}

type GetTasksByUserID200JSONResponse []Task

func (response GetTasksByUserID200JSONResponse) VisitGetTasksByUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
	// Delete a task by ID
	// (DELETE /tasks/{task_id})
	DeleteTaskByID(ctx context.Context, request DeleteTaskByIDRequestObject) (DeleteTaskByIDResponseObject, error)
	// Update a task by ID
	// (PATCH /tasks/{task_id})
	UpdateTaskByID(ctx context.Context, request UpdateTaskByIDRequestObject) (UpdateTaskByIDResponseObject, error)
	// Get all user tasks by ID
	// (GET /user/{user_id}/tasks)
	GetTasksByUserID(ctx context.Context, request GetTasksByUserIDRequestObject) (GetTasksByUserIDResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context) error {
	var request GetTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(ctx echo.Context) error {
	var request PostTasksRequestObject

	var body PostTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx.Request().Context(), request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		return validResponse.VisitPostTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTaskByID operation middleware
func (sh *strictHandler) DeleteTaskByID(ctx echo.Context, taskId uint) error {
	var request DeleteTaskByIDRequestObject

	request.TaskId = taskId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTaskByID(ctx.Request().Context(), request.(DeleteTaskByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTaskByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskByIDResponseObject); ok {
		return validResponse.VisitDeleteTaskByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateTaskByID operation middleware
func (sh *strictHandler) UpdateTaskByID(ctx echo.Context, taskId uint) error {
	var request UpdateTaskByIDRequestObject

	request.TaskId = taskId

	var body UpdateTaskByIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateTaskByID(ctx.Request().Context(), request.(UpdateTaskByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateTaskByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateTaskByIDResponseObject); ok {
		return validResponse.VisitUpdateTaskByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTasksByUserID operation middleware
func (sh *strictHandler) GetTasksByUserID(ctx echo.Context, userId uint) error {
	var request GetTasksByUserIDRequestObject

	request.UserId = userId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasksByUserID(ctx.Request().Context(), request.(GetTasksByUserIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasksByUserID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksByUserIDResponseObject); ok {
		return validResponse.VisitGetTasksByUserIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
