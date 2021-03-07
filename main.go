package main

import (
	"net/http"

	"github.com/philippta/web-frontend-demo/html"
)

func main() {
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/profile/show", profileShow)
	http.HandleFunc("/profile/edit", profileEdit)
	http.ListenAndServe(":8080", nil)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	p := html.DashboardParams{
		Title:   "Dashboard",
		Message: "Hello from dashboard",
	}
	html.Dashboard(w, p)
}

func profileShow(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileShowParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}
	html.ProfileShow(w, p)
}

func profileEdit(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileEditParams{
		Title:   "Profile Edit",
		Message: "Hello from profile edit",
	}
	html.ProfileEdit(w, p)
}
