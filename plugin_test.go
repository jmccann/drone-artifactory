package main

import (
	"fmt"
	// "io"
	"testing"
	. "github.com/franela/goblin"

	"github.com/jmccann/drone-artifactory/fixtures"
	"net/http/httptest"
)

func TestPlugin(t *testing.T) {
	g := Goblin(t)
	server := httptest.NewServer(fixtures.Handler())

	g.Describe("Exec", func() {
		g.Before(func() {
			// Use fake server url
			plugin.Config.Url = server.URL

		})
		g.After(func() {
			server.Close()
		})

		g.It("Should upload files and directories", func() {
			err := plugin.Exec()
			uploaded := len(plugin.Config.Sources)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
			g.Assert(uploaded == 2).IsTrue(fmt.Sprintf("Should have uploaded 2 files instead of %d files", uploaded))
		})

		g.It("Should upload a file", func() {
			// Set to a single file
			plugin.Config.Sources = []string{"main.go"}

			err := plugin.Exec()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
			//g.Assert(uploaded == 1).IsTrue(fmt.Sprintf("Should have uploaded 1 file instead of %d files", uploaded))
		})

		g.It("Should upload glob of files", func() {
			// Set to a single file
			plugin.Config.Sources = []string{"*.go"}

			err := plugin.Exec()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
			//g.Assert(uploaded == 3).IsTrue(fmt.Sprintf("Should have uploaded 3 files instead of %d files", uploaded))
		})
	})
}

var (
	c = Config{
		DryRun: true,
		Path: "thekey/with/path",
		Password: "supersecret",
		ApiKey: "apikeyofartifactory",
		Sources: []string{"main.go", "fixtures/*"},
		Url: "http://company.com",
		Username: "johndoe",
	}
	plugin = Plugin{
		Config: c,
	}
)
