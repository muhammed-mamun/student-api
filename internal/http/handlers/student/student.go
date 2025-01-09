package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/muhammed-mamun/student-api/internal/types"
	"github.com/muhammed-mamun/student-api/internal/utils/responses"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		slog.Info("creating a student")

		responses.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
