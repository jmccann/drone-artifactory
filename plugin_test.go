package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	"github.com/jmccann/drone-artifactory/fixtures"
)

func TestPlugin(t *testing.T) {
	g := Goblin(t)
	p := plugin
	server := httptest.NewServer(fixtures.Handler())

	g.Describe("Exec", func() {
		g.Before(func() {
			p = plugin
			// Use fake server url
			plugin.Config.URL = server.URL

		})
		g.After(func() {
			server.Close()
		})

		g.It("Should upload files and directories", func() {
			err := p.Exec()
			uploaded := len(p.Config.Sources)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
			g.Assert(uploaded == 2).IsTrue(fmt.Sprintf("Should have uploaded 2 files instead of %d files", uploaded))
		})

		g.It("Should upload a file", func() {
			// Set to a single file
			p.Config.Sources = []string{"main.go"}
			err := p.Exec()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
		})

		g.It("Should upload glob of files", func() {
			// Set to a single file
			p.Config.Sources = []string{"*.go"}

			err := p.Exec()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
		})
	})

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
			config.APIKey = ""
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
