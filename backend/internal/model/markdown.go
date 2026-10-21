package model

type Markdown struct {
	Content string `json:"content" binding:"required"`
}

type ConversionResponse struct {
	HTML string `json:"html"`
}
