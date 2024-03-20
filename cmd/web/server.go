package web

import (
	"book_ex/internal/models"
	"log"
)

type Database map[string]int

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Items    *models.ItemModel
}
