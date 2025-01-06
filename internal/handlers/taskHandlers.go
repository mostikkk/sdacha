package handlers

import (
	"context"
	"fmt"
	"log"
	"obuch/internal/taskService"
	"obuch/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

// NewTaskHandler создает новый TaskHandler с привязанным сервисом задач
func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// GetTasksUserId реализует tasks.StrictServerInterface.
func (h *TaskHandler) GetTasksUserId(_ context.Context, req tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	userID := req.UserId

	userTasks, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		log.Printf("Error fetching tasks for user %d: %v", userID, err)
		return tasks.GetTasksUserId404Response{}, nil
	}

	var response tasks.GetTasksUserId200JSONResponse
	for _, tsk := range userTasks {
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
	taskToCreate := taskService.Task{
		Task:   taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
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

// DeleteTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasks(_ context.Context, req tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	taskID := req.Body.Id

	err := h.Service.DeleteTask(uint(taskID))
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", taskID, err)
		return tasks.DeleteTasks404Response{}, nil
	}

	return tasks.DeleteTasks204Response{}, nil
}

func (h *TaskHandler) PutTasks(_ context.Context, req tasks.PutTasksRequestObject) (tasks.PutTasksResponseObject, error) {
	taskRequest := req.Body
	taskID := taskRequest.Id // Убрали разыменование, так как это значение

	// Получение задач по userID (предположим, taskID соответствует userID)
	existingTasks, err := h.Service.GetTasksByUserID(uint(taskID))
	if err != nil || len(existingTasks) == 0 {
		log.Printf("Task with ID %d not found: %v", taskID, err)
		return tasks.PutTasks404Response{}, nil
	}

	// Предположим, что обновляем первую задачу из списка
	existingTask := existingTasks[0]

	// Обновление задачи
	updatedTask := taskService.Task{
		ID:     existingTask.ID,     // Используем ID из существующей задачи
		Task:   taskRequest.Task,    // Новая задача из запроса
		IsDone: *taskRequest.IsDone, // Новый статус
		UserID: existingTask.UserID, // Сохраняем исходный UserID
	}

	// Вызываем сервис обновления
	resultTask, err := h.Service.UpdateTask(uint(existingTask.ID), updatedTask)
	if err != nil {
		log.Printf("Error updating task with ID %d: %v", taskID, err)
		return nil, fmt.Errorf("failed to update task")
	}

	// Формируем ответ
	response := tasks.PutTasks200JSONResponse{
		Id:     &resultTask.ID,
		Task:   &resultTask.Task,
		IsDone: &resultTask.IsDone,
		UserId: &resultTask.UserID,
	}

	return response, nil
}
