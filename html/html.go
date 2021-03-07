package html

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

var (
	dashboard   = parse("dashboard.html")
	profileShow = parse("profile/show.html")
	profileEdit = parse("profile/edit.html")
)

type DashboardParams struct {
	Title   string
	Message string
}

func Dashboard(w io.Writer, p DashboardParams) error {
	return dashboard.Execute(w, p)
}

type ProfileShowParams struct {
	Title   string
	Message string
}

func ProfileShow(w io.Writer, p ProfileShowParams) error {
	return profileShow.Execute(w, p)
}

type ProfileEditParams struct {
	Title   string
	Message string
}

func ProfileEdit(w io.Writer, p ProfileEditParams) error {
	return profileEdit.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
