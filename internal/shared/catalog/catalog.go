package catalog

import (
	"net/http"

	"chillhub/internal/shared/errors"
	"chillhub/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type Catalog struct {
	Status  int
	Message string
}

func (c Catalog) Err(err error, msg ...string) *errors.AppError {
	if err == nil {
		return nil
	}

	message := c.Message
	if len(msg) > 0 && msg[0] != "" {
		message = msg[0]
	}

	return &errors.AppError{
		Status:  c.Status,
		Message: message,
		Err:     err,
	}
}

func (c Catalog) Success(cxt *gin.Context, data any, msg ...string) {
	message := c.Message
	if len(msg) > 0 && msg[0] != "" {
		message = msg[0]
	}

	response.Send(cxt, c.Status, message, data)
}

func (c Catalog) Fail(cxt *gin.Context, err error, msg ...string) {
	message := c.Message
	if len(msg) > 0 && msg[0] != "" {
		message = msg[0]
	}

	var data any
	if err != nil {
		data = err.Error()
	}

	response.Send(cxt, c.Status, message, data)
}

var (
	OK      = Catalog{http.StatusOK, "common.success"}
	Created = Catalog{http.StatusCreated, "common.created"}

	BadRequest   = Catalog{http.StatusBadRequest, "common.client.bad_request"}
	Unauthorized = Catalog{http.StatusUnauthorized, "common.client.unauthorized"}
	Forbidden    = Catalog{http.StatusForbidden, "common.client.forbidden"}
	NotFound     = Catalog{http.StatusNotFound, "common.client.not_found"}
	Conflict     = Catalog{http.StatusConflict, "common.client.conflict"}

	Internal = Catalog{http.StatusInternalServerError, "common.server.internal_error"}
)
