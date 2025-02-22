package ViewModel

import (
	"AnthropicIntegration/ProjectManager/Services"
	"strings"
)

type ProjectViewModel struct {
	repository Services.ProjectServiceRepository
}

func NewProjectViewModel(repository Services.ProjectServiceRepository) ProjectViewModel {
	return ProjectViewModel{
		repository: repository,
	}
}

func (pm *ProjectViewModel) FetchRepositoryLinks() ([]string, error) {
	projects, err := pm.repository.FetchProjects()
	if err != nil {
		return nil, err
	}
	projectLinks := make([]string, len(projects))
	for _, project := range projects {
		projectLinks = append(projectLinks, project.Github)
	}
	return projectLinks, nil
}

func (pm *ProjectViewModel) FetchImageLinksFromDifferentDomain() ([]string, error) {
	projects, err := pm.repository.FetchProjects()
	if err != nil {
		return nil, err
	} else {
		var unknownDomains []string = []string{}
		for _, project := range projects {
			if project.ImagesURL == nil {
				unknownDomains = append(unknownDomains, "unknown missing domain")
				continue
			}
			for _, imageUrl := range project.ImagesURL {
				hostDomainSubstring := "sareenv"
				if strings.Contains(imageUrl, hostDomainSubstring) == false {
					unknownDomains = append(unknownDomains, imageUrl)
				}
			}
		}
		return unknownDomains, nil
	}
}
