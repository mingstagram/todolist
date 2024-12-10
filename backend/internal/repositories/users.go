package repositories

import (
	"backend/internal/common"
	"backend/internal/models"
	"database/sql"
)

type UsersRepository struct {
	DB *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

// 로그인 처리
func (r *UsersRepository) FindUserByEmailAndPassword(email, password string) (models.Users, error) {
	var user models.Users
	query := "SELECT id, email, name FROM users WHERE email = ? AND password = ?"
	err := r.DB.QueryRow(query, email, password).Scan(&user.ID, &user.Email, &user.Name)

	if err == sql.ErrNoRows {
		return user, common.ErrInvalidEmailOrPassword
	} else if err != nil {
		return user, err
	}

	return user, nil
}

// 회원가입
func (r *UsersRepository) SaveUsers(user models.Users) error {
	// 트랜잭션 처리 
	// tx, err := r.DB.Begin()
	// if err!= nil {
	// 	return fmt.Errorf("failed to start transaction: %v", err)
    // }
	// // result, err := tx.Exec("INSERT INTO users (email, name, password) VALUES (?, ?, ?)", users.Email, users.Name, users.Password)
	// _, err = tx.Exec("INSERT INTO users (email, name, password) VALUES (?, ?, ?)", users.Email, users.Name, users.Password)
	// if err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("failed to execute query: %v", err)
	// }
 
	// err = tx.Commit()
	// if err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("failed to commit transaction: %v", err)
	// }

	// return nil 

	_, err := r.DB.Exec("INSERT INTO users (email, name, password) VALUES (?, ?, ?)", user.Email, user.Name, user.Password)
	return err
}

// sns회원가입

