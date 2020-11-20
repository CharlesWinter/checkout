package checkout

// Repository defines a type capable of affording basic checkout capability
type Repository struct {
}

// RepositoryConfig is the config struct for the repostitory
type RepositoryConfig struct {
}

// New returns a new Repository
func New(cfg RepositoryConfig) *Repository {
	return &Repository{}
}
