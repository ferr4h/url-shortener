package url

type CreateUrlRequest struct {
	Url string `json:"url" validate:"required,url"`
}
