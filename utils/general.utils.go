package utils

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadTemplates(root string) *template.Template {
	tmpl := template.New("")

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ⛔ Skip directories
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".html" {
			_, err := tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	return tmpl
}
