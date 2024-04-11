package interceptor

import (
	"fmt"
	"net/http"
)

const (
	userToken  = "user_token"
	adminToken = "admin_token"
)

func CheckToken(token string) (int, error) {
	if token == userToken || token == adminToken {
		return http.StatusOK, nil
	}
	return http.StatusUnauthorized, fmt.Errorf("пользователь не авторизован")
}

func CheckAdminToken(token string) (int, error) {
	if token == adminToken {
		return http.StatusOK, nil
	}
	if token == userToken {
		return http.StatusForbidden, fmt.Errorf("пользователь не имеет доступа")
	}

	return http.StatusUnauthorized, fmt.Errorf("пользователь не авторизован")
}
