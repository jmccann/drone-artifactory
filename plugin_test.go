package main

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestPlugin(t *testing.T) {
	g := Goblin(t)
	p := plugin

	g.Describe("Validate Config", func() {
		config := p.Config

		g.BeforeEach(func() {
			config = p.Config
		})

		g.It("- should validate input", func() {
			err := config.validate()
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to validate input: %s", err))
		})
		g.It("- should fail on password/api key", func() {
			config.Password = ""
			config.APIKey = ""
			err := config.validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No ApiKey or Password provided"))
		})
		g.It("- should fail on url", func() {
			config.URL = ""
			err := config.validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No url provided"))
		})
		g.It("- should fail on username", func() {
			config.Username = ""
			err := config.validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No username provided"))
		})
	})

	g.Describe("Validate UploadArgs", func() {
		args := p.UploadArgs

		g.BeforeEach(func() {
			args = p.UploadArgs
		})

		g.It("- should fail on sources", func() {
			args.Sources = []string{}
			err := args.validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No sources provided"))
		})
		g.It("- should fail on path", func() {
			args.Path = ""
			err := args.validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No path provided"))
		})
	})
}

var (
	c = Config{
		Password: "supersecret",
		APIKey:   "apikeyofartifactory",
		URL:      "http://company.com",
		Username: "johndoe",
	}
	u = UploadArgs{
		DryRun:  true,
		Path:    "thekey/with/path",
		Sources: []string{"main.go", "fixtures/*"},
	}
	plugin = Plugin{
		Config:     c,
		UploadArgs: u,
	}
)
