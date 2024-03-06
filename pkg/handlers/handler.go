package handlers

import (
	"net/http"
	"web/pkg/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl.html")
}

func About(w http.ResponseWriter, r *http.Request) {

}
