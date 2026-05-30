package static

import (
	"log"
	"os"
	"path/filepath"
)

// ResolveCatalogImagesDir finds frontend/public/images (seed catalog paths like /images/products/...).
func ResolveCatalogImagesDir() string {
	candidates := []string{
		"../../frontend/public/images",
		"../frontend/public/images",
		"frontend/public/images",
	}
	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(wd, "../../frontend/public/images"),
			filepath.Join(wd, "../../../frontend/public/images"),
		)
	}
	for _, dir := range candidates {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			abs, _ := filepath.Abs(dir)
			return abs
		}
	}
	return ""
}

// LogCatalogImagesReady logs whether catalog static files are available (caller mounts r.Static).
func LogCatalogImagesReady(dir string) {
	if dir == "" {
		log.Printf("[static] catalog images dir not found; /images/* will 404 (upload /uploads/* still works)")
		return
	}
	log.Printf("[static] serving /images from %s", dir)
}
