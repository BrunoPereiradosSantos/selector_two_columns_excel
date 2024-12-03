package handlers

func HandleErrors(message string, err error, status int) (string, error, int) {
	if err != nil {
		return message, err, status
	}

	return "", nil, 0
}
