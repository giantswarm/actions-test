package project

var (
	description = "Minimal binary for testing GitHub Actions?"
	gitSHA      = "n/a"
	name        = "actions-test"
	source      = "https://github.com/giantswarm/actions-test"
	version     = "0.1.1-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
