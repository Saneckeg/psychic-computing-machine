package handlers

import (
	"SecondProject/internal/taskService"
	"SecondProject/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func TaskNewHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) DeleteTaskId(ctx context.Context, request tasks.DeleteTaskIdRequestObject) (tasks.DeleteTaskIdResponseObject, error) {
	id := request.Id

	_, err := h.Service.DeleteTaskByID(uint(id))

	if err != nil {
		return nil, err
	}
	return nil, err
}

func (h *TaskHandler) PatchTaskId(ctx context.Context, request tasks.PatchTaskIdRequestObject) (tasks.PatchTaskIdResponseObject, error) {
	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), request.Body)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTaskId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) GetTask(_ context.Context, _ tasks.GetTaskRequestObject) (tasks.GetTaskResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTask200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTask(_ context.Context, request tasks.PostTaskRequestObject) (tasks.PostTaskResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTask201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}
