package interceptor

import (
	"fmt"
)

const (
	userToken  = "user_token"
	adminToken = "admin_token"
)

func CheckToken(token string) error {
	if token == userToken || token == adminToken {
		return nil
	}
	return fmt.Errorf("пользователь не авторизован")
}

func CheckAdminToken(token string) error {
	if token == adminToken {
		return nil
	}
	if token == userToken {
		return fmt.Errorf("пользователь не имеет доступа")
	}

	return fmt.Errorf("пользователь не авторизован")
}
