package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"
	"time"
)

type TasksService struct {
	TasksRepo *repositories.TasksRepository
}

func NewTasksService(repo *repositories.TasksRepository) *TasksService {
	return &TasksService{TasksRepo: repo}
}

// 오늘의 할일 조회
func (s *TasksService) GetTodayTasks() ([]models.Tasks, error) { 
	today := time.Now().Format("2006-01-02")
	return s.TasksRepo.GetTasksByDate(today)
}

// 특정 날짜의 할일 조회
func (s *TasksService) GetTasksForDate(date time.Time) ([]models.Tasks, error) {
	dateString := date.Format("2006-01-02") 
	return s.TasksRepo.GetTasksByDate(dateString)
}

// 할일 추가
func(s *TasksService) SaveTasks(tasks models.Tasks) error {
	if tasks.Task == "" {
		return fmt.Errorf("task cannot be empty")
	}

	err := s.TasksRepo.SaveTasks(tasks)
	if err != nil {
		return fmt.Errorf("failed to save task: %v", err)
	}

	return nil
}

// 할일 체크 유무 카운팅
func (s *TasksService) CountTasks(date time.Time) (int, error) {
	return s.TasksRepo.CountTasks(date)
}

// 할일 체크 / 체크해제
func (s *TasksService) UpdateChecked(checked bool, id int) error {
	return s.TasksRepo.UpdateChecked(checked, id)
}

// 할일 삭제 
func (s *TasksService) DeleteTasks(id int) error {
	return s.TasksRepo.DeleteTasks(id)
}