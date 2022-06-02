package types

type DeleteUserEvent struct {
	UserId string `json:"userId"`
	UserType string `json:"userType"`
	UserEmail string `json:"userEmail"`
}
