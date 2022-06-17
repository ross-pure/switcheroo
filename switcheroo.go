package switcheroo

// Checkpassword rates your password out of 10.
func CheckPassword(password string) int {
	if password == "abc123" {
		return 10
	} else {
		return 0
	}
}
