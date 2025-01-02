package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type TasksRepository struct {
	DB *sql.DB
}

func NewTasksRepository(db *sql.DB) *TasksRepository {
	return &TasksRepository{DB: db}
}

// 일별 할일 조회
func (r *TasksRepository) GetTasksByDate(date string, userId string) ([]models.Tasks, error) { 
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE DATE(created_at) = ? AND user_id = ? AND !is_deleted", date, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Tasks
	for rows.Next() {
		var task models.Tasks  
		var userId sql.NullInt64  // NULL 처리용 변수
		if err := rows.Scan(&task.ID, &task.Task, &task.IsChecked, &task.IsDeleted, &task.CreatedAt, &userId); err != nil {
			return nil, err
		}

		if userId.Valid {
			task.UserId = int(userId.Int64)
		} else {
			task.UserId = 0  // NULL인 경우 기본값 설정
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
		return errors.New("failed to start transaction")
	} 
	// 삽입 쿼리 실행
	result, err := tx.Exec("INSERT INTO tasks (task, created_at, user_id) values (?, ?, ?)", tasks.Task, tasks.CreatedAt, tasks.UserId)
	if err != nil {
		tx.Rollback()
		return errors.New("failed to execute query")
	} 

	// 삽입된 레코드의 ID를 가져오기
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return errors.New("failed to get last insert id") 
	}

	// 커밋
	err = tx.Commit()
	if err != nil {
		tx.Rollback() 
		return errors.New("failed to commit transaction") 
	}

	// 삽입된 ID를 models.Tasks에 추가
	tasks.ID = int(lastInsertID)
 
	return nil
}

// 할일 체크 유무 카운팅
func (r *TasksRepository) CountTasks(date time.Time, userId string) (int, error) {
	rows, err := r.DB.Query("SELECT COUNT(*) FROM tasks WHERE !is_checked AND !is_deleted AND DATE(created_at) = ? AND user_id = ?", date, userId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
        if err := rows.Scan(&count); err!= nil {
            return 0, err
        }  
    }
	return count, nil
}

// 할일 체크 / 체크해제
func (r *TasksRepository) UpdateChecked(checked bool, id int) error {
	_, err := r.DB.Exec("UPDATE tasks SET is_checked = ? WHERE id = ?", checked, id)
	if err != nil {
		return fmt.Errorf("failed to execute query for id %d: %v", id, err)
	}
	return nil
}

// 할일 삭제
func (r *TasksRepository) DeleteTasks(id int) error {
	_, err := r.DB.Exec("UPDATE tasks SET is_deleted = true WHERE id =?", id)
	if err != nil {
		return fmt.Errorf("failed to execute query for id %d: %v", id, err)
	}
	return nil
}


// 할일 수정 
