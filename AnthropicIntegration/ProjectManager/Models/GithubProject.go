package Models

type GithubProject struct {
	CommitCounts   int
	BranchCounts   int
	StarCounts     int
	GithubURL      string
	RepositoryName string
	// in case we don't have any description check if that's empty string or null
	RepositoryDescription string
	CreatedAt             string
	lastUpdatedAt         string
}
