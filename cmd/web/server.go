package web

import (
	"book_ex/internal/models"
	"github.com/go-playground/form/v4"
	"html/template"
	"log"
)

type Database map[string]int

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Items         *models.ItemModel
	Reviews       *models.ReviewModel
	TemplateCache map[string]*template.Template
	FormDecoder   *form.Decoder
}
