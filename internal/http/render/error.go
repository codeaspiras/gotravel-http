package render

import (
	"net/http"

	"github.com/codeaspiras/gotravel/stdio"
)

func Error(io stdio.IO, w http.ResponseWriter, err error, statuses ...int) {
	statusCode := http.StatusInternalServerError
	if len(statuses) > 0 {
		statusCode = statuses[0]
	}

	JSON(io, w, statusCode, map[string]interface{}{
		"status": statusCode,
		"error":  err.Error(),
	})
}
