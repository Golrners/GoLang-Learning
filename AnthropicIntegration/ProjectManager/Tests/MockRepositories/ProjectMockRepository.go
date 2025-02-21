package MockRepositories

import "AnthropicIntegration/ProjectManager/Models"

type ProjectMockRepository struct {
	generationType MockGenerationType
}

type MockGenerationType int

// WithNilUnknownProjects An enum equivalent in golang.
const (
	WithNilUnknownProjects MockGenerationType = iota
	WithThreeUnknownProjects
	WithNoUnknownProjects
)

// NewProjectMockRepository is an initializer for the struct.
func NewProjectMockRepository(generationType MockGenerationType) *ProjectMockRepository {
	return &ProjectMockRepository{
		generationType: generationType,
	}
}

// AddNewProject A mock method which dictates adding the new project.
// TODO : - Needs the complete implementation.
func (pm *ProjectMockRepository) AddNewProject(_ string, _ Models.Project) {}

func (pm *ProjectMockRepository) FetchProjectFromID(_ string) (*Models.Project, error) {
	return &Models.Project{
		Github:      "sampleLink 2",
		TagLine:     "sampleTagLine 2",
		ProjectName: "sampleProjectName 2",
		ImagesURL: []string{
			"https://sareenv.com/img1",
			"https://sareenv.com/img2",
		},
		ID: "32SKOPJE",
	}, nil
}

func (pm *ProjectMockRepository) GenerateMockWithNilImageUrls() []Models.Project {
	return []Models.Project{
		{
			Github:      "sampleLink 2",
			TagLine:     "sampleTagLine 2",
			ProjectName: "sampleProjectName 2",
			ImagesURL:   nil,
			ID:          "32SKOPJE",
		},
	}
}

func (pm *ProjectMockRepository) GenerateMockWithZeroUnknownProjects() []Models.Project {
	return []Models.Project{
		{
			Github:      "sampleLink",
			TagLine:     "sampleTagLine",
			ProjectName: "sampleProjectName",
			ImagesURL: []string{"" +
				"https://sareenv.com/img1",
				"https://sareenv.com/img2",
			},
			ID: "2LOEWI",
		},
	}
}

func (pm *ProjectMockRepository) GenerateMockWith3UnknownProjects() []Models.Project {
	return []Models.Project{
		{
			Github:      "sampleLink",
			TagLine:     "sampleTagLine",
			ProjectName: "sampleProjectName",
			ImagesURL: []string{"" +
				"https://sareenv.com/img1",
				"https://sareenv.com/img2",
				"https://bell.co.uk/api/v1/product",
				"https://riging.com/api/v1/product",
			},
			ID: "2LOEWI",
		},

		{
			Github:      "sampleLink2",
			TagLine:     "sampleTagLine2",
			ProjectName: "sampleProjectName",
			ImagesURL: []string{
				"https://sareenv.com/img1",
				"https://wikipedia.org",
			},
			ID: "COMPOJURE",
		},
	}
}

func (pm *ProjectMockRepository) GenerateMockProjects() []Models.Project {
	switch pm.generationType {
	case WithNilUnknownProjects:
		return pm.GenerateMockWithNilImageUrls()
	case WithThreeUnknownProjects:
		return pm.GenerateMockWith3UnknownProjects()
	case WithNoUnknownProjects:
		return pm.GenerateMockWithZeroUnknownProjects()
	}
	return nil
}

func (pm *ProjectMockRepository) FetchProjects() ([]Models.Project, error) {
	return pm.GenerateMockProjects(), nil
}
