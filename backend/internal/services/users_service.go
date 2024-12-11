package services

import (
	"backend/internal/common"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
)

type UsersService struct {
	UsersRepo *repositories.UsersRepository
}

func NewUsersService(repo *repositories.UsersRepository) *UsersService {
    return &UsersService{UsersRepo: repo}
}

// 로그인 처리
func (s *UsersService) Login(email, password string) (models.Users, error) {
	// 이메일 또는 비밀번호 유효성 검증 추가 가능
	if email == "" || password == "" {
		return models.Users{}, common.ErrEmptyEmailOrPassword 
	} 
	return s.UsersRepo.FindUserByEmailAndPassword(email, password)
}

// 회원가입
func (s *UsersService) SaveUsers(user models.Users) error { 
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