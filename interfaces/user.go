package interfaces

type User struct {
	USERNAME string `json:"username"`
	NAME     string `json:"name"`
	MOBILE   string `json:"mobile"`
	PASSWORD string `json:"password"`
	EMAIL    string `json:"email"`
}
