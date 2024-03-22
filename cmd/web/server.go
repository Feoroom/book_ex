package web

import (
	"book_ex/internal/models"
	"html/template"
	"log"
)

type Database map[string]int

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Items         *models.Model
	TemplateCache map[string]*template.Template
}
