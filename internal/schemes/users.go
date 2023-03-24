package schemes

type ReqSignUp struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password1 string `json:"password1"`
}

type ReqSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResSignIn struct {
	AccessToken string `json:"access_token"`
}
