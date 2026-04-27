package dto

type UrlResponse struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
}
