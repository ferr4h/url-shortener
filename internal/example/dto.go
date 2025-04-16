package example

type ExampleRequest struct {
	Line string `json:"Line" validate:"ip"`
}

type ExampleResponse struct {
	Line string
}
