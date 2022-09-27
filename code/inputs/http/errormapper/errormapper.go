package errormapper

import (
	"github.com/GeoguessrTools/calico-cut-pants/core"
	"net/http"
	"strings"
)

// GetResponseCodeFromError generates an appropriate HTTP response code give the core error. If the error is not a core
// error or this function doesn't know how to categorize the error, it returns http.StatusInternalServerError.
func GetResponseCodeFromError(suppliedErr error) int {
	coreErr, ok := suppliedErr.(core.Error)
	if !ok {
		return http.StatusInternalServerError
	}

	if strings.HasPrefix(coreErr.Code, core.UserErrPrefix) {
		return http.StatusBadRequest
	}

	if strings.HasPrefix(coreErr.Code, core.UnknownErrorPrefix) {
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
