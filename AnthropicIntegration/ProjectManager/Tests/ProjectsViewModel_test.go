package Tests

import (
	"AnthropicIntegration/ProjectManager/Tests/MockRepositories"
	"AnthropicIntegration/ProjectManager/ViewModel"
	"testing"
	_ "testing"
)

// Given When Then.
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

// TODO: Can we do a parametrized tests ?
