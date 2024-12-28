package handlers

import (
	"context"
	"obuch/internal/taskService" // Импортируем наш сервис
	"obuch/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

// NewHandler создает новый Handler с привязанным сервисом задач
func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// GetTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получаем все задачи из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Формируем ответ для gRPC
	response := tasks.GetTasks200JSONResponse{}

	// Перебираем все задачи и заполняем ответ
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	// Возвращаем респонсивный ответ
	return response, nil
}

// PostTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Получаем тело запроса
	taskRequest := request.Body

	// Создаем задачу для добавления в сервис
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}

	// Создаем задачу через сервис
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// Формируем ответ с созданной задачей
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}

	// Возвращаем ответ с задачей
	return response, nil
}

// DeleteTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasks(ctx context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	// Получаем ID задачи для удаления
	id := request.Body.Id

	// Вызываем метод для удаления задачи
	err := h.Service.DeleteTask(uint(id))
	if err != nil {
		return nil, err
	}

	// Формируем ответ о успешном удалении
	response := tasks.DeleteTasks204Response{}

	// Возвращаем успешный ответ
	return response, nil
}

// PutTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) PutTasks(_ context.Context, request tasks.PutTasksRequestObject) (tasks.PutTasksResponseObject, error) {
	// Получаем ID задачи из запроса
	id := *request.Body.Id // Разыменовываем указатель

	taskRequest := request.Body

	// Создаем задачу для обновления
	taskToUpdate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}

	// Вызываем метод сервиса для обновления задачи
	updatedTask, err := h.Service.UpdateTask(id, taskToUpdate) // Передаем id как uint
	if err != nil {
		return nil, err
	}

	// Формируем ответ с обновленной задачей
	response := tasks.PutTasks200JSONResponse{
		Id:     &updatedTask.ID, // Возвращаем ID, так как это поле может быть на выходе
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
		UserId: &updatedTask.UserID,
	}

	// Возвращаем ответ с обновленной задачей
	return response, nil
}
