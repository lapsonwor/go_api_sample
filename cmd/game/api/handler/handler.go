package handler

import (
	"fmt"
	"lapson_go_api_sample/cmd/game/controller"
	"lapson_go_api_sample/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var handlerLogger *logrus.Entry = logger.GetLogger("handler")

type Handler struct {
	controller        *controller.Controller
	advUsageLogBuffer map[string]chan []byte
}

func New(controller *controller.Controller) *Handler {
	return &Handler{
		controller:        controller,
		advUsageLogBuffer: map[string]chan []byte{},
	}
}

type Hello struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Result  string `json:"result"`
}

func (h *Handler) Hello(c *gin.Context) {
	response := Hello{
		Message: "Hello World",
		Status:  "200",
		Result:  "Success",
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) verifyJWT(c *gin.Context) {

}

func (h *Handler) ErrorResponse(w http.ResponseWriter, r *http.Request, httpStatus int) {
	w.WriteHeader(httpStatus)
	fmt.Fprint(w, http.StatusText(httpStatus))
}

// Not Found 404
func (h *Handler) StatusNotFound(w http.ResponseWriter, r *http.Request) {
	h.ErrorResponse(w, r, http.StatusNotFound)
}

// Bad Request 400
func (h *Handler) StatusBadRequest(w http.ResponseWriter, r *http.Request) {
	h.ErrorResponse(w, r, http.StatusBadRequest)
}

// Internal Server Error 500
func (h *Handler) StatusInternalServerError(w http.ResponseWriter, r *http.Request) {
	h.ErrorResponse(w, r, http.StatusInternalServerError)
}

// Temporary Redirect 307
func (h *Handler) StatusTemporaryRedirect(w http.ResponseWriter, r *http.Request) {
	h.ErrorResponse(w, r, http.StatusTemporaryRedirect)
}
