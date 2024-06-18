package api

type PositionMessage struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type RegisterPayload struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"userType" binding:"required"`
}

type BatchRegisterPayload struct {
	Payloads []RegisterPayload `json:"data"`
}
