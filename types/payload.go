package types

type DecodedJwt struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	Id    string   `json:"id"`
	Iat   int64    `json:"iat"`
	Exp   int64    `json:"exp"`
}

type Payload struct {
	User DecodedJwt
}
