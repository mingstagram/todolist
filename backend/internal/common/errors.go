package common

import "errors"

// 에러 메시지 정의
var (
    ErrInvalidEmailOrPassword  = errors.New("invalid email or password")
    ErrEmptyEmailOrPassword    = errors.New("email and password cannot be empty")
    ErrEmptyTask    = errors.New("task cannot be empty")
    ErrMissingRequiredFields   = errors.New("all fields are required")
    ErrFailedToSaveTask       = errors.New("failed to save task")
)