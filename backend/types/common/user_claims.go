package common

type UserClaims struct {
	UserId string `json:"user_id"`
}

func (v *UserClaims) Valid() error {
	return nil
}
