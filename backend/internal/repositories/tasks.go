package repositories

import (
	"backend/internal/models"
	"database/sql"
	"fmt"
)

type TasksRepository struct {
	DB *sql.DB
}

func NewTasksRepository(db *sql.DB) *TasksRepository {
	return &TasksRepository{DB: db}
}

// 일별 할일 조회
func (r *TasksRepository) GetTasksByDate(date string) ([]models.Tasks, error) { 
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE DATE(created_at) = ?", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Tasks
	for rows.Next() {
		var task models.Tasks 
		if err := rows.Scan(&task.ID, &task.Task, &task.IsChecked, &task.IsDeleted, &task.CreatedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// 할일 추가
func (r *TasksRepository) SaveTasks(tasks models.Tasks) error {
	// 트랜잭션을 사용하여 여러 쿼리 처리가 필요한 경우 안전하게 처리
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	// 삽입 쿼리 실행
	result, err := tx.Exec("INSERT INTO tasks (task) values (?)", tasks.Task)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute query: %v", err)
	}

	// 삽입된 레코드의 ID를 가져오기
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get last insert id: %v", err)
	}

	// 커밋
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	// 삽입된 ID를 models.Tasks에 추가
	tasks.ID = int(lastInsertID)
 
	return nil
}

// 할일 삭제

// 할일 수정

// 할일 체크 / 체크해제

// 날짜 이동

// 할일 체크 유무 카운팅
