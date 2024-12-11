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

func (r *UsersRepository) FindUserByEmail(email string) (models.Users, error) {
	var user models.Users
    query := "SELECT id, email, name, password FROM users WHERE email =?"
    err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Password)

    if err == sql.ErrNoRows {
        return user, common.ErrUserNotFound
    } else if err!= nil {
        return user, err
    }

    return user, nil
}

// 회원가입
func (r *UsersRepository) SaveUsers(user models.Users) error {   
	query := "INSERT INTO users (email, name, password) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, user.Email, user.Name, user.Password)
	return err
}

// sns회원가입

