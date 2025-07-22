package handler

import (
	"encurtador-url/internal/crypto"
	"encurtador-url/internal/utils"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       sync.Mutex
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortId := r.URL.Path[1:]

	mu.Lock()
	encryptedUrl, exists := urlStore[shortId]
	mu.Unlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	decryptedUrl := crypto.Decrypt(encryptedUrl)
	http.Redirect(w, r, decryptedUrl, http.StatusFound)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.URL.Query().Get("url")
	if originalUrl == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	if !(strings.HasPrefix(originalUrl, "http://") || strings.HasPrefix(originalUrl, "https://")) {
		http.Error(w, "URL must start with http:// or https://", http.StatusBadRequest)
		return
	}

	encryptedUrl := crypto.Encrypt(originalUrl)
	shortId := utils.GenerateShortId()

	mu.Lock()
	urlStore[shortId] = encryptedUrl
	mu.Unlock()

	shortUrl := fmt.Sprintf("http://localhost:8080/%s", shortId)
	fmt.Fprintf(w, "A URL encurtada deste url original é %s", shortUrl)
}
