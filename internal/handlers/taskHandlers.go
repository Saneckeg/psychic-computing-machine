package handlers

import (
	"SecondProject/internal/taskService"
	"SecondProject/internal/web/tasks"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteTaskId(ctx context.Context, request tasks.DeleteTaskIdRequestObject) (tasks.DeleteTaskIdResponseObject, error) {
	id := request.Id

	_, err := h.Service.DeleteTaskByID(uint(id))

	if err != nil {
		return nil, err
	}
	return nil, err
}

func (h *Handler) PatchTaskId(ctx context.Context, request tasks.PatchTaskIdRequestObject) (tasks.PatchTaskIdResponseObject, error) {
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

func (h *Handler) GetTask(_ context.Context, _ tasks.GetTaskRequestObject) (tasks.GetTaskResponseObject, error) {
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

func (h *Handler) PostTask(_ context.Context, request tasks.PostTaskRequestObject) (tasks.PostTaskResponseObject, error) {
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

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) PatchMessages(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), updates)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(updatedTask)
}

//func (h *Handler) DeleteMessages(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	idStr := vars["id"]
//
//	id, err := strconv.ParseUint(idStr, 10, 32)
//	if err != nil {
//		http.Error(w, "Invalid ID", http.StatusBadRequest)
//		return
//	}
//
//	_, err = h.Service.DeleteTaskByID(uint(id))
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprintln(w, "Message deleted")
//}
