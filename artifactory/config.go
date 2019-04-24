package artifactory

import "fmt"

type (
	// Config for Artifactory
	Config struct {
		APIKey   string
		Debug    bool
		Password string
		URL      string
		Username string
	}
)

// Validate the Config
func (c Config) Validate() error {
	if len(c.Username) == 0 && len(c.APIKey) == 0 {
		return fmt.Errorf("No username or ApiKey provided")
	}

	if len(c.Password) == 0 && len(c.APIKey) == 0 {
		return fmt.Errorf("No ApiKey or Password provided")
	}

	if len(c.URL) == 0 {
		return fmt.Errorf("No url provided")
	}

	return nil
}
