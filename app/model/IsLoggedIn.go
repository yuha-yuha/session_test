package model

func IsLoggedIn(Id interface{}) (bool, int) {

	Idint, ok := Id.(int)

	if Idint == 0 {
		return false, 0
	}

	return ok, Idint

}
