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

	tmpl := template.New("").Funcs(template.FuncMap{})

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking path %q: %v", path, err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".htmx" {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			templateName := filepath.ToSlash(relPath)

			_, err = tmpl.ParseFiles(path)
			if err != nil {
				log.Printf("Error parsing template %q: %v", path, err)
				return err
			}
			log.Printf("Loaded template: %s", templateName) // Log loaded templates for debugging
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	return tmpl
}
