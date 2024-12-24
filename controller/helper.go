package controller

import (
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func GetFileFromForm(e echo.Context, response map[string]interface{}) (multipart.File, *multipart.FileHeader, string, string, int) {
	fileHeader, err := e.FormFile("image")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			response["ErrorField"] = []map[string]string{{"image": "required"}}
			return nil, nil, "", "", http.StatusBadRequest
		}
		response["error"] = err.Error()
		return nil, nil, "", "", http.StatusInternalServerError
	}

	fileType := fileHeader.Header.Get("Content-Type")

	ext := filepath.Ext(fileHeader.Filename)

	validMimeTypes := map[string]string{
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".svg":  "image/svg+xml",
	}

	validMime, ok := validMimeTypes[ext]
	if !ok || validMime != fileType {
		response["ErrorField"] = []map[string]string{{"image": "Invalid format ektensi image  must .jpeg, .jpg, .png, .gif, .svg"}}
		return nil, nil, "", "", http.StatusBadRequest
	}

	file, err := fileHeader.Open()
	if err != nil {
		response["error"] = err.Error()
		return nil, nil, "", "", http.StatusInternalServerError
	}
	defer file.Close()

	return file, fileHeader, ext, fileType, http.StatusOK
}
