package handlers

import (
	"context"

	"Test.go/internal/models"
	"Test.go/internal/taskService"
	"Test.go/internal/web/tasks"
)

type taskHandler struct {
	Service *taskService.TaskService
}

func (h *taskHandler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	Alltasks, err := h.Service.GetTasksByUserID(request.UserId)
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasksByUserID200JSONResponse{}
	for _, tsk := range Alltasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: tsk.IsDone,
			UserId: &tsk.User_ID,
		}
		response = append(response, task)
	}
	return response, nil
}

func NewTaskHandler(service *taskService.TaskService) *taskHandler {
	return &taskHandler{
		Service: service,
	}
}

func (h *taskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: tsk.IsDone,
			UserId: &tsk.User_ID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *taskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := models.Task{
		Task:    *taskRequest.Task,
		IsDone:  taskRequest.IsDone,
		User_ID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: createdTask.IsDone,
		UserId: &createdTask.User_ID,
	}

	return response, nil
}

func (h *taskHandler) UpdateTaskByID(ctx context.Context, request tasks.UpdateTaskByIDRequestObject) (tasks.UpdateTaskByIDResponseObject, error) {
	taskRequest := request.Body

	instanceUpdateTask := models.Task{
		Task:    *taskRequest.Task,
		IsDone:  taskRequest.IsDone,
		User_ID: *taskRequest.UserId,
	}
	updatedTask, err := h.Service.UpdateTaskByID(request.TaskId, instanceUpdateTask)
	if err != nil {
		return nil, err
	}
	response := tasks.UpdateTaskByID200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: updatedTask.IsDone,
		UserId: &updatedTask.User_ID,
	}

	return response, nil
}

func (h *taskHandler) DeleteTaskByID(ctx context.Context, request tasks.DeleteTaskByIDRequestObject) (tasks.DeleteTaskByIDResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.TaskId)
	if err != nil {
		return tasks.DeleteTaskByID204Response{}, err
	}
	return tasks.DeleteTaskByID204Response{}, nil
}
