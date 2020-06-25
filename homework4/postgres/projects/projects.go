package projects

import rsrc "github.com/AndreyKlimchuk/golang-learning/homework4/resources"

type Query struct {
}

func (q Query) Create(name string, description string) (rsrc.ProjectExpanded, error) {
	return rsrc.ProjectExpanded{}, nil
}

func (q Query) Get(projectId rsrc.Id) (rsrc.Project, error) {
	return rsrc.Project{}, nil
}

func (q Query) GetExpanded(projectId rsrc.Id) (rsrc.ProjectExpanded, error) {
	return rsrc.ProjectExpanded{}, nil
}

func (q Query) GetMultiple() ([]rsrc.Project, error) {
	return []rsrc.Project{}, nil
}

func (q Query) Update(projectId rsrc.Id, name string, description string) error {
	return nil
}

func (q Query) Delete(projectId rsrc.Id) error {
	return nil
}
