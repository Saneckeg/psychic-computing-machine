package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTasksByUserID(userID int) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *TaskService) UpdateTaskByID(id uint, task interface{}) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) (Task, error) {
	return s.repo.DeleteTaskByID(id)
}
