package services

import (
	"backend/internal/common"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
 
// 카카오 로그인 처리
func (s *AuthService) GetKakaoAccessToken(code string) (string, error) {
	// 카카오 인증 API 요청
	redirectURI := url.QueryEscape("http://localhost:3000/kakao/callback")
	clientID := "f935795848dcfc4ebec4716acfc63bfa"
	reqBodyData := "grant_type=authorization_code&client_id=" + clientID +
		"&redirect_uri=" + redirectURI + "&code=" + code

	url := "https://kauth.kakao.com/oauth/token"
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBodyData))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse map[string]interface{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	// access_token 추출
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return accessToken, nil
}

func (s *AuthService) GetKakaoUserInfo(accessToken string) (map[string]interface{}, error) {
	userInfoURL := "https://kapi.kakao.com/v2/user/me"
	
	// 요청할 property_keys를 JSON 배열로 정의
	propertyKeys := []string{
		"kakao_account.email",
		"kakao_account.profile",
		"kakao_account.name",
		"kakao_account.age_range",
		"kakao_account.birthday",
		"kakao_account.gender",
	}

	// property_keys를 JSON 형식으로 변환
	propertyKeysJSON, err := json.Marshal(propertyKeys)
	if err != nil {
		return nil, err
	}

	// URL 인코딩
	propertyKeysEncoded := url.QueryEscape(string(propertyKeysJSON))

	// 본문 데이터를 설정 (property_keys 쿼리 파라미터를 포함)
	reqBodyData := "property_keys=" + propertyKeysEncoded
	req, err := http.NewRequest("POST", userInfoURL, strings.NewReader(reqBodyData))
	if err != nil {
		return nil, err
	}

	// 요청 헤더 설정
	req.Header.Add("Authorization", "Bearer " + accessToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	// HTTP 클라이언트로 요청 보내기
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// JSON 응답 파싱
	var userInfoResponse map[string]interface{}
	err = json.Unmarshal(body, &userInfoResponse)
	if err != nil {
		return nil, err
	}

	// 반환된 사용자 정보 출력
	fmt.Println("userInfo: ", userInfoResponse)

	return userInfoResponse, nil
}