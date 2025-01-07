package handlers

import (
	"context"

	"Test.go/internal/taskService"
	"Test.go/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
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
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: createdTask.IsDone,
	}

	return response, nil
}

func (h *Handler) UpdateTaskByID(ctx context.Context, request tasks.UpdateTaskByIDRequestObject) (tasks.UpdateTaskByIDResponseObject, error) {
	taskRequest := request.Body

	instanceUpdateTask := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}
	updatedTask, err := h.Service.UpdateTaskByID(request.TaskId, instanceUpdateTask)
	if err != nil {
		return nil, err
	}
	response := tasks.UpdateTaskByID200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: updatedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTaskByID(ctx context.Context, request tasks.DeleteTaskByIDRequestObject) (tasks.DeleteTaskByIDResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.TaskId)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTaskByID204Response{}, nil
}
