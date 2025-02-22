package Tests

import (
	"AnthropicIntegration/ProjectManager/Tests/MockRepositories"
	"AnthropicIntegration/ProjectManager/ViewModel"
	"testing"
	_ "testing"
)

func TestUnknownProjectImageLinksWithNilValue(t *testing.T) {
	// Given: A mock repository and view model which fetch the projects based on the type.
	mockRepository := MockRepositories.NewProjectMockRepository(MockRepositories.WithNilUnknownProjects)
	vm := ViewModel.NewProjectViewModel(mockRepository)

	// When: The unknown links from different domains are fetched
	unknownDomains, _ := vm.FetchImageLinksFromDifferentDomain()

	// Then: The count should be 1 since there is only nil based image url mocked.
	unknownDomainsCount := len(unknownDomains)
	expectedUnknownDomainsCount := 1
	if unknownDomainsCount != expectedUnknownDomainsCount {
		t.Error("Wrong number of domains fetched from different URL")
	}
}

func TestUnknownProjectImageLinksWith3UnknownLinks(t *testing.T) {
	// Given: A mock repository and view model which fetch the projects based on the type.
	mockRepository := MockRepositories.NewProjectMockRepository(MockRepositories.WithThreeUnknownProjects)
	vm := ViewModel.NewProjectViewModel(mockRepository)

	// When: The unknown links from different domains are fetched
	unknownDomains, _ := vm.FetchImageLinksFromDifferentDomain()
	unknownDomainsCount := len(unknownDomains)

	// Then: The count should be 3 since there is only nil based image url mocked.
	expectedUnknownDomainsCount := 3

	if unknownDomainsCount != expectedUnknownDomainsCount {
		t.Error("Wrong number of domains fetched from different URL")
	}
}

func TestUnknownProjectImageLinksWithZeroUnknownLinks(t *testing.T) {
	// Given: A mock repository and view model which fetch the projects based on the type.
	mockRepository := MockRepositories.NewProjectMockRepository(MockRepositories.WithNoUnknownProjects)
	vm := ViewModel.NewProjectViewModel(mockRepository)

	// When: The unknown links from different domains are fetched
	unknownDomains, _ := vm.FetchImageLinksFromDifferentDomain()
	unknownDomainsCount := len(unknownDomains)

	// Then: The count should be 0 since there is only nil based image url mocked.
	expectedUnknownDomainsCount := 0

	if unknownDomainsCount != expectedUnknownDomainsCount {
		t.Error("Wrong number of domains fetched from different URL")
	}
}

// TODO: Can we do a parametrized tests ?
