package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.sagyzdop.com/internal/models"
	"snippetbox.sagyzdop.com/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Call the newTemplateData() helper to get a templateData struct containing
	// the 'default' data (which for now is just the current year), and add the
	// snippets slice to it.
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// Pass the data to the render() helper as normal.
	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

    flash := app.sessionManager.PopString(r.Context(), "flash")

	// And do the same thing again here...
	data := app.newTemplateData(r)
	data.Snippet = snippet
    data.Flash = flash
	
    app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

// Change the signature of the snippetCreate handler so it is defined as a method
// against *application.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
    data.Form = snippetCreateForm{
        Expires: 365,
    }
	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

type snippetCreateForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Expires int    `form:"expires"`
	validator.Validator `form:"-"`
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    var form snippetCreateForm

    err := app.decodePostForm(r, &form)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
    form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
    form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
        return
    }

    
    id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
    if err != nil {
        app.serverError(w, r, err)
        return
    }
    app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
    
    http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}