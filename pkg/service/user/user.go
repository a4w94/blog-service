package user

import "regexp"

// 檢查用戶email是否非法
func checkEmail(email string) bool {
	const pattern = "^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
