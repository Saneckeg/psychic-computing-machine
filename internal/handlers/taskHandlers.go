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

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	_, err := h.Service.DeleteTaskByID(uint(id))

	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), request.Body)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
		UserId: &updatedTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) GetTasks(_ context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	var allTasks []taskService.Task
	var err error

	if request.Params.UserID != nil {
		allTasks, err = h.Service.GetTasksByUserID(*request.Params.UserID)
	} else {
		allTasks, err = h.Service.GetAllTasks()
	}

	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}
