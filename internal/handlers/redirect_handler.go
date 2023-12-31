package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/scetle/url-shortener/internal/database"
	"github.com/scetle/url-shortener/internal/models"
	"gorm.io/gorm"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
  db, err := database.NewDB()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  shortURL := "localhost:8080/" + r.URL.Path[1:]
  var url models.URL
  result := db.DB.Where("short_url = ?", shortURL).First(&url)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    http.NotFound(w, r)
    return
  }
  if strings.HasPrefix(url.OriginalURL, "https://") || strings.HasPrefix(url.OriginalURL, "http://") {
  http.Redirect(w, r, url.OriginalURL, http.StatusFound)
} else {
  http.Redirect(w, r, "//" + url.OriginalURL, http.StatusFound)
  }
}
