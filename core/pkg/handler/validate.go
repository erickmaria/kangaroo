package handler

func Validate(err error) bool {
	if err != nil {
		return true
	}

	return false
}
