package artifactory

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestPlugin(t *testing.T) {
	g := Goblin(t)

	g.Describe("Validate Config", func() {
		config := c

		g.BeforeEach(func() {
			config = c
		})

		g.It("should validate input", func() {
			err := config.Validate()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to validate input: %s", err))
		})
		g.It("should fail on password/api key", func() {
			config.Password = ""
			config.APIKey = ""
			err := config.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No ApiKey or Password provided"))
		})
		g.It("should fail on url", func() {
			config.URL = ""
			err := config.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No url provided"))
		})
		g.It("should fail on username", func() {
			config.Username = ""
			config.APIKey = ""
			err := config.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No username or ApiKey provided"))
		})
	})
}

var (
	c = Config{
		Password: "supersecret",
		APIKey:   "apikeyofartifactory",
		URL:      "http://company.com/artifactory",
		Username: "johndoe",
	}
)
