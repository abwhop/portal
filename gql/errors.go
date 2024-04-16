package gql

type ErrorRespond struct {
	Errors []struct {
		Message   string `json:"message"`
		Category  string `json:"category"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
	} `json:"errors"`
}
