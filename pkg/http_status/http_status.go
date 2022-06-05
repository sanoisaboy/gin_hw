package http_status

func Get_http_status(len int) (int, error) {
	if len != 0 {
		return 200, nil
	}

	return 403, nil
}

func Create_http_status(student_name string, id string, point string) (int, error) {
	if id == "" || student_name == "" || point == "" {
		return 403, nil
	}

	return 201, nil
}

func Update_http_status(id string, point string) (int, error) {
	if id == "" || point == "" {
		return 400, nil
	}
	return 200, nil
}

func Delete_http_status(id string) (int, error) {
	if id == "" {
		return 400, nil
	}

	return 204, nil
}
