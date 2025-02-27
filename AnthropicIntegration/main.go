package main

import (
	"AnthropicIntegration/ProjectManager/Services"
	"AnthropicIntegration/ProjectManager/Services/NetworkManager"
	"AnthropicIntegration/ProjectManager/ViewModel"
	"fmt"
	_ "strings"
)

func FetchProjectsAndLog() {
	fmt.Println("Fetching Projects With Unknown Images...")
	projectsViewModel := ViewModel.NewProjectViewModel(
		Services.NewProjectManager(
			"https://api.sareenv.com/api",
			"/v1/projects",
			"/v1/projects/upload",
			NetworkManager.BaseNetworkManager{}),
	)
	unknownDomains, err := projectsViewModel.FetchImageLinksFromDifferentDomain()
	if err != nil {
		// similar to the fatal error in swift.
		panic(err)
	} else {
		for _, unknownDomain := range unknownDomains {
			fmt.Println(unknownDomain)
		}
	}
}

// Download a json object from the API.
func main() {
	FetchProjectsAndLog()
}
