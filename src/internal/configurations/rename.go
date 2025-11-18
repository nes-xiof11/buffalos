package configurations

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReplaceProjectName(root, oldName, newName string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		if !isTextFile(path) {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		updated := strings.ReplaceAll(string(content), oldName, newName)

		if string(content) != updated {
			err := os.WriteFile(path, []byte(updated), 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func isTextFile(path string) bool {
	return strings.HasSuffix(path, ".go") ||
		strings.HasSuffix(path, ".mod") ||
		strings.HasSuffix(path, ".sum") ||
		strings.HasSuffix(path, ".sql") ||
		strings.HasSuffix(path, ".txt") ||
		strings.HasSuffix(path, ".md")
}

func Rename(args []string) {
	if len(args) < 3 {
		log.Fatal("Missing project name. Use: ./app --rename newname")
	}

	oldName := "buffalos"
	newName := args[2]

	if err := ReplaceProjectName(".", oldName, newName); err != nil {
		log.Fatalf("Rename failed: %v", err)
	}

	fmt.Println("Project renamed successfully!")
	return
}
