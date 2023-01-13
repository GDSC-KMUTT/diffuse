package payload

type CreateResponse struct {
	Name string `json:"name"`
	Pin  string `json:"pin"`
}

type CreateRequest struct {
	Name string `json:"name"`
}

type JoinRequest struct {
	Name string `json:"name"`
	Pin  string `json:"pin"`
}

type JoinResponse struct {
	Name         string `json:"name"`
	OpponentName string `json:"opponent_name"`
}
