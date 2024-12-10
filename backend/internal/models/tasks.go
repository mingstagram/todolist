package models

import "time"
 
type Tasks struct {
    ID         int       `json:"id"`                        
    UserId int       `json:"user_id"` 
    Task       string    `json:"task"`                        
    IsChecked  bool      `json:"is_checked"`                  
    IsDeleted  bool      `json:"is_deleted"`                 
    CreatedAt  time.Time `json:"created_at"`             
}
