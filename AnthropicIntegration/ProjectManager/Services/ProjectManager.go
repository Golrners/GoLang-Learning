package Services

import (
	"AnthropicIntegration/ProjectManager/Models"
	_ "AnthropicIntegration/ProjectManager/Models"
	"AnthropicIntegration/ProjectManager/Services/NetworkManager"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ProjectServiceRepository interface {
	AddNewProject(url string, project Models.Project)
	FetchProjectFromID(url string) (*Models.Project, error)
	FetchProjects() ([]Models.Project, error)
}

type ProjectManager struct {
	baseUrl         string
	fetchEndPoint   string
	storeEndPoint   string
	projectEndpoint string
	networkManager  NetworkManager.BaseNetworkManager
}

// NewProjectManager Similar to the implementation of the initializer.
func NewProjectManager(baseUrl string,
	fetchEndPoint string,
	storeEndPoint string,
	networkManager NetworkManager.BaseNetworkManager) *ProjectManager {
	return &ProjectManager{
		baseUrl:        baseUrl,
		fetchEndPoint:  fetchEndPoint,
		storeEndPoint:  storeEndPoint,
		networkManager: networkManager,
	}
}

// AddNewProject A method required to update a new project.
// TODO: Verify the working of the method.
func (pm *ProjectManager) AddNewProject(url string, project Models.Project) {
	data, jsonErr := json.Marshal(project)
	if jsonErr != nil {
		log.Fatal(`Could not add new project to JSON format`)
		return
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
		return
	} else {
		// Add the custom headers to the request.
		request.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(request)
		defer response.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	defer request.Body.Close()
}

// FetchProjectFromID a method to fetch the project from the identifier.
func (pm *ProjectManager) FetchProjectFromID(url string) (*Models.Project, error) {
	body, err := pm.networkManager.FetchResponseBody(url)
	if err != nil {
		return nil, err
	}
	var project Models.Project
	err = json.NewDecoder(body).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (pm *ProjectManager) FetchProjects() ([]Models.Project, error) {
	url := pm.baseUrl + pm.fetchEndPoint
	body, err := pm.networkManager.FetchResponseBody(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		decoder := json.NewDecoder(body)
		var decodedProjects Models.Projects
		err := decoder.Decode(&decodedProjects)
		if err != nil {
			return nil, err
		} else {
			return decodedProjects.Projects, nil
		}
	}
}
