package helpers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func ImageUploader(w http.ResponseWriter, r *http.Request) (string, error) {
	// Where it stores uploaded files
	uploadsDir := "./static/uploads/"
	// Check if the uploads directory exists
	_, errs := os.Stat(uploadsDir)
	if os.IsNotExist(errs) {
		os.MkdirAll(uploadsDir, os.ModePerm)
	}

	if errs != nil {
		return "", errs
	}

	// Get the upload image
	file, handler, err := r.FormFile("image")
	if err != nil {
		return "", err

	}
	supportedExtensions := []string{".jpeg", ".png", ".gif", ".jpg"}
	extensionNotSupported := true
	for _, supportedExt := range supportedExtensions {
		if filepath.Ext(handler.Filename) == supportedExt {
			extensionNotSupported = false
			break
		}
	}
	if extensionNotSupported {
		return "", errors.New("Unsupported image extension")
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			// Check the file size.
			if fileHeader.Size > (20 * 1024 * 1024) {
				return "", errors.New("File size exceeds the limit.")
			}
		}
	}
	// Generate a unique filename for each upload
	fileName := time.Now().Format("20060102150405") + filepath.Ext(handler.Filename)
	filepath := filepath.Join(uploadsDir, fileName)
	// Create the file in the 'uploads' directory
	f, err := os.Create(filepath)

	if err != nil {
		ErrorThrower(err, "An error occurred while uploading the image", http.StatusInternalServerError, w, r)
		return "", err
	}

	defer f.Close()

	// Copy the uploaded file's content to the newly created file
	_, copyErr := io.Copy(f, file)

	if copyErr != nil {
		ErrorThrower(err, "An error occurred while uploading the image", 500, w, r)
		return "", copyErr
	}
	return fileName, err

}
