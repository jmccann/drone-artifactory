package main

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestPlugin(t *testing.T) {
	g := Goblin(t)
	p := plugin

	g.Describe("Validate Input", func() {
		config := p.Config

		g.BeforeEach(func() {
			config = p.Config
		})

		g.It("- should validate input", func() {
			err := validateInput(config)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to validate input: %s", err))
		})
		g.It("- should fail on sources", func() {
			config.Sources = []string{}
			err := validateInput(config)
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No sources provided"))
		})
		g.It("- should fail on password/api key", func() {
			config.Password = ""
			config.APIKey = ""
			err := validateInput(config)
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No ApiKey or Password provided"))
		})
		g.It("- should fail on path", func() {
			config.Path = ""
			err := validateInput(config)
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No path provided"))
		})
		g.It("- should fail on url", func() {
			config.URL = ""
			err := validateInput(config)
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No url provided"))
		})
		g.It("- should fail on username", func() {
			config.Username = ""
			err := validateInput(config)
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No username provided"))
		})
	})
}

var (
	c = Config{
		DryRun:   true,
		Path:     "thekey/with/path",
		Password: "supersecret",
		APIKey:   "apikeyofartifactory",
		Sources:  []string{"main.go", "fixtures/*"},
		URL:      "http://company.com",
		Username: "johndoe",
	}
	plugin = Plugin{
		Config: c,
	}
)
