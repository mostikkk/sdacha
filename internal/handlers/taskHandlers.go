package handlers

import (
	"context"
	"fmt"
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

// GetTasksByUserID реализует tasks.StrictServerInterface.
func (h *TaskHandler) GetTasksUserId(_ context.Context, req tasks.GetTasksUserIdRequestObject) (tasks.GetTasksUserIdResponseObject, error) {
	// Извлекаем userID из запроса
	userID := req.UserId

	// Получаем задачи для указанного пользователя из сервиса
	userTasks, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasksUserId200JSONResponse

	// Перебираем задачи и заполняем ответ
	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.UserID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		// Используем append для добавления элемента в слайс
		response = append(response, task)
	}

	// Возвращаем сформированный ответ
	return response, nil
}

// PostTasks реализует tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Получаем тело запроса
	taskRequest := request.Body

	// Создаем задачу для добавления в сервис
	taskToCreate := taskService.Task{
		Task:   taskRequest.Task, // Передаем строку Task без разыменовывания
		IsDone: *taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}

	// Создаем задачу через сервис
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// Формируем ответ с созданной задачей
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.UserID,
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
	id := *request.Body.UserId // Разыменовываем указатель, чтобы получить ID

	taskRequest := request.Body

	// Проверяем, существует ли хотя бы одна задача с таким ID (можно добавить дополнительную проверку)
	existingTasks, err := h.Service.GetTasksByUserID(uint(id))
	if err != nil || len(existingTasks) == 0 {
		// Если задачи не найдены, возвращаем ошибку
		return nil, fmt.Errorf("task with id %d not found", id)
	}

	// Берем первую задачу из списка задач
	existingTask := existingTasks[0]

	// Обновляем задачу на основе полученного запроса
	taskToUpdate := taskService.Task{
		Task:   taskRequest.Task,    // Обновляем поле Task
		IsDone: *taskRequest.IsDone, // Обновляем статус IsDone
		UserID: existingTask.UserID, // Пользователь ID остается тем же
	}

	// Вызываем метод сервиса для обновления задачи
	updatedTask, err := h.Service.UpdateTask(uint(id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	// Формируем ответ с обновленной задачей
	response := tasks.PutTasks200JSONResponse{
		Id:     &updatedTask.UserID, // Возвращаем ID задачи
		Task:   &updatedTask.Task,   // Обновленное поле задачи
		IsDone: &updatedTask.IsDone, // Статус задачи
		UserId: &updatedTask.UserID, // ID пользователя
	}

	// Возвращаем ответ с обновленной задачей
	return response, nil
}
