package api

type TokenResponse struct {
	Token string `json:"token"`
}

type ResponseWSMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}
