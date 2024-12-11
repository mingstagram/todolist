package services

import (
	"backend/internal/common"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
)

type AuthService struct {
	UsersRepo *repositories.UsersRepository
}

func NewAuthService(repo *repositories.UsersRepository) *AuthService {
	return &AuthService{UsersRepo: repo}
}

func (s *AuthService) Login(email, password string) (interface{}, error) {
    user, err := s.UsersRepo.FindUserByEmail(email)
    if err != nil {
        return nil, common.ErrUserNotFound
    }

    if !utils.ComparePassword(user.Password, password) {
        return nil, common.ErrInvalidEmailOrPassword
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return nil, err
    }

    // Define the response structure
    var responseData struct {
        Token  string `json:"token"`
        UserID int    `json:"userId"`
    }

    responseData.Token = token
    responseData.UserID = user.ID

    return responseData, nil
}

// 회원가입
func (s *AuthService) SaveUsers(user models.Users) error { 
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return common.ErrMissingRequiredFields 
	}
	
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	
	return s.UsersRepo.SaveUsers(user)
}