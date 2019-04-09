// +build cli

package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	"github.com/jmccann/drone-artifactory/fixtures"
)

func TestPluginCLI(t *testing.T) {
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
			uploaded := len(p.UploadArgs.Sources)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
			g.Assert(uploaded == 2).IsTrue(fmt.Sprintf("Should have uploaded 2 files instead of %d files", uploaded))
		})

		g.It("Should upload a file", func() {
			// Set to a single file
			p.UploadArgs.Sources = []string{"main.go"}
			err := p.Exec()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
		})

		g.It("Should upload glob of files", func() {
			// Set to a single file
			p.UploadArgs.Sources = []string{"*.go"}

			err := p.Exec()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to upload stuff: %s", err))
		})
	})
}
