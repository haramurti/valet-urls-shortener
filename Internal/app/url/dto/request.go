package dto

type CreateUrlRequest struct {
	OriginalUrl     string `json:"original_url" validate:"required,url"`
	CustomShortCode string `json:"short_code" validate:"required,alphanum,min=4,max=255"`
}

//premature code, not yet implemented
// type UpdateUrlRequest struct {
// 	OriginalUrl string `json:"original_url" validate:"required,url"`
// 	ShortCode   string `json:"short_code" validate:"required,alphanum,min=4,max=30"`
// }

//making still need to convert databse from string to database.
