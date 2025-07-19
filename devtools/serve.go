package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		h.ServeHTTP(rw, r)

		status := "SUCCESS"
		if rw.statusCode >= 400 {
			status = "FAILED"
		}

		log.Printf("[ACCESS] %s %s %s - Status: %d (%s)",
			r.Method, r.URL.Path, r.RemoteAddr, rw.statusCode, status)

		if rw.statusCode >= 400 {
			log.Printf("[ERROR] Request failed - Code: %d, Reason: %s",
				rw.statusCode, http.StatusText(rw.statusCode))

			// 404エラーの場合、実際のファイルシステムパスを記録
			if rw.statusCode == 404 {
				webDir := "../web/"
				requestedFile := filepath.Join(webDir, r.URL.Path)
				absolutePath, _ := filepath.Abs(requestedFile)

				if _, err := os.Stat(requestedFile); os.IsNotExist(err) {
					log.Printf("[FILE_NOT_FOUND] Attempted to access: %s (absolute: %s)",
						requestedFile, absolutePath)

					// ディレクトリアクセスの場合はindex.htmlを確認
					if r.URL.Path[len(r.URL.Path)-1] == '/' || filepath.Ext(r.URL.Path) == "" {
						indexFile := filepath.Join(requestedFile, "index.html")
						if _, err := os.Stat(indexFile); err == nil {
							log.Printf("[DIRECTORY_ACCESS] Directory access detected, index.html exists: %s", indexFile)
						} else {
							log.Printf("[DIRECTORY_ACCESS] Directory access detected, but no index.html found: %s", indexFile)
						}
					}
				}
			}
		}
	})
}

func main() {
	webDir := "./web/"
	absWebDir, err := filepath.Abs(webDir)
	if err != nil {
		log.Fatal("Failed to get absolute path for web directory:", err)
	}

	log.Printf("[DEV] Web directory: %s (absolute: %s)", webDir, absWebDir)

	// webディレクトリの存在確認
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		log.Fatal("Web directory does not exist:", absWebDir)
	}

	// index.htmlの存在確認
	indexPath := filepath.Join(webDir, "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		log.Printf("[WARNING] index.html not found at: %s", indexPath)
	} else {
		log.Printf("[INFO] index.html found at: %s", indexPath)
	}

	fileServer := http.FileServer(http.Dir(webDir))
	http.Handle("/", loggingHandler(fileServer))
	log.Println("[DEV] Serving web/ directory on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
