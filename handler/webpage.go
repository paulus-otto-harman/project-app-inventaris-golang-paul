package handler

import (
	"fmt"
	"net/http"
	"os"
	"project/service"
	"strings"
)

type WebPageHandler struct {
	WebPageService service.WebPageService
}

func InitWebPageHandler(webPageService service.WebPageService) WebPageHandler {
	return WebPageHandler{WebPageService: webPageService}
}

func (handle *WebPageHandler) Index(w http.ResponseWriter, r *http.Request) {
	handle.WebPageService.Render(w, "index.html", "Home")
}

func (handle *WebPageHandler) Static(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	var data []byte
	data, _ = os.ReadFile(fmt.Sprintf("view/%s", path))
	w.Header().Set("Content-Type", "text/css")
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "application/javascript")
	}

	w.Write(data)
}
