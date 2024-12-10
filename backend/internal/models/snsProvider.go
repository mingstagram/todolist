package models

type SnsProvider string

const (
	SnsFacebook SnsProvider = "facebook"
	SnsKakao    SnsProvider = "kakao"
	SnsGoogle   SnsProvider = "google"
)

func (s SnsProvider) IsValid() bool {
	switch s {
	case SnsFacebook, SnsKakao, SnsGoogle:
		return true
	}
	return false
}