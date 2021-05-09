package render

import (
	"goBooking/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	resSession, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(resSession.Context(), "flash", "123")
	res := AddDefaultData(&td, resSession)

	if res.Flash != "123" {
		t.Error("Flash value of 123 not found")
	}
}

func TestRenderTemplate(t *testing.T) {
	relativeTempPath = "../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplate(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that doesn't exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	// set the variable context, needed everytime we read or write from the session
	ctx := r.Context()

	// put the session data inside req context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Sessionn"))

	// update request to have context
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	relativeTempPath = "../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
