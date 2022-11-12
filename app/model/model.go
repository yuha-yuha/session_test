package model

func IsLoggedIn(userId int) bool {
	if userId != 0 {
		return false
	} else {
		return true
	}
}
