package models

import "time"

// Task represents a row in the tasks table
type Tasks struct {
    ID         int       `json:"id"`                          // id 컬럼
    Task       string    `json:"task"`                        // task 컬럼
    IsChecked  bool      `json:"is_checked"`                  // is_checked 컬럼
    IsDeleted  bool      `json:"is_deleted"`                  // is_deleted 컬럼
    CreatedAt  time.Time `json:"created_at"`                  // created_at 컬럼
}
