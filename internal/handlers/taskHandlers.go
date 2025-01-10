package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"obuch/internal/taskService"
	"obuch/internal/web/tasks"

	"gorm.io/gorm"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

// NewTaskHandler создает новый TaskHandler с привязанным сервисом задач.
func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// GetTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		log.Printf("Error fetching all tasks: %v", err)
		return nil, fmt.Errorf("failed to fetch tasks")
	}

	var response tasks.GetTasks200JSONResponse
	for _, tsk := range allTasks {
		response = append(response, tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		})
	}

	return response, nil
}

// PostTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(_ context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := req.Body
	if taskRequest.IsDone == nil {
		return nil, fmt.Errorf("is_done field is required")
	}

	taskToCreate := taskService.Task{
		Task:   taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}

	createdTask, err := h.Service.PostTask(taskToCreate)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, fmt.Errorf("failed to create task")
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}

	return response, nil
}

// GetTasksTaskId реализует tasks.StrictServerInterface.
func (h *TaskHandler) GetTasksTaskId(_ context.Context, req tasks.GetTasksTaskIdRequestObject) (tasks.GetTasksTaskIdResponseObject, error) {
	// Получение ID задачи из запроса
	taskID := req.TaskId

	// Запрос задачи через сервис
	task, err := h.Service.GetTaskByID(uint(taskID))
	if err != nil {
		log.Printf("Error fetching task with ID %d: %v", taskID, err)

		// Проверка на ошибку "не найдено"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tasks.GetTasksTaskId404Response{}, nil
		}

		return nil, err
	}

	// Формирование успешного ответа
	response := tasks.GetTasksTaskId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
		UserId: &task.UserID,
	}

	return response, nil
}

// DeleteTasksTaskId реализует tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasksTaskId(_ context.Context, req tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	taskID := req.TaskId

	err := h.Service.DeleteTask(uint(taskID))
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", taskID, err)
		return tasks.DeleteTasksTaskId404Response{}, nil
	}

	return tasks.DeleteTasksTaskId204Response{}, nil
}

// PatchTasksTaskId реализует tasks.StrictServerInterface.
func (h *TaskHandler) PatchTasksTaskId(_ context.Context, req tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	taskID := req.TaskId
	taskRequest := req.Body

	updatedTask := taskService.Task{
		ID:     taskID,
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}

	resultTask, err := h.Service.PathTask(uint(taskID), updatedTask)
	if err != nil {
		log.Printf("Error updating task with ID %d: %v", taskID, err)
		return tasks.PatchTasksTaskId404Response{}, nil
	}

	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &resultTask.ID,
		Task:   &resultTask.Task,
		IsDone: &resultTask.IsDone,
		UserId: &resultTask.UserID,
	}

	return response, nil
}

// GetTasksUserUserId реализует tasks.StrictServerInterface.
func (h *TaskHandler) GetTasksUserUserId(_ context.Context, req tasks.GetTasksUserUserIdRequestObject) (tasks.GetTasksUserUserIdResponseObject, error) {
	userID := req.UserId

	userTasks, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		log.Printf("Error fetching tasks for user %d: %v", userID, err)
		return nil, fmt.Errorf("failed to fetch user tasks")
	}

	var response tasks.GetTasksUserUserId200JSONResponse
	for _, tsk := range userTasks {
		response = append(response, tasks.Task{
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		})
	}

	return response, nil
}
