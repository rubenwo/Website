package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//Project ...
type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	PreviewImage string   `json:"preview_image"`
	Images       []string `json:"images"`
	Videos       []string `json:"videos"`
}

var projects map[string]Project

func getProjects() []Project {
	var p []Project

	for _, value := range projects {
		p = append(p, value)
	}

	return p
}

func getProject(name string) (Project, error) {

	p, ok := projects[name]
	if !ok {
		return Project{}, errors.New("project doesn't exist")
	}
	return p, nil
}

func loadProjects(path string) error {
	projects = make(map[string]Project)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var p []Project

	err = json.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	for _, value := range p {
		projects[value.Name] = value
	}

	return nil
}
