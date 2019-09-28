package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

//InitAPI ...
func InitAPI(path string) error {
	return loadProjects(path)
}

//ProjectsEndpoint ...
func ProjectsEndpoint(w http.ResponseWriter, r *http.Request) {
	p := getProjects()

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		HandleErrorAsJSON(w, Error{
			Status: http.StatusInternalServerError,
			Msg:    "sending projects response as stream went wrong",
		})
		return
	}
}

//ProjectEndpoint ...
func ProjectEndpoint(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	p, err := getProject(name)
	if err != nil {
		HandleErrorAsJSON(w, Error{
			Status: http.StatusNotFound,
			Msg:    "no project with name: " + name + " was found",
		})
		return
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		HandleErrorAsJSON(w, Error{
			Status: http.StatusInternalServerError,
			Msg:    "sending projects response as stream went wrong",
		})
		return
	}
}
