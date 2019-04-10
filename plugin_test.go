package main

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
	"github.com/jmccann/drone-artifactory/artifactory"
)

func TestPlugin(t *testing.T) {
	g := Goblin(t)

	g.Describe("parseArgs", func() {
		g.It("should validate default input", func() {
			expectedArguments := artifactory.UploadArgs{
				Explode: false,
				Path:    "some/path",
				Sources: []string{"some/source"},
			}
			action := Action{
				Name:         "upload",
				RawArguments: []byte(`{"name":"upload","sources":["some/source"],"path":"some/path"}`),
			}
			err := parseArgs(&action)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to parse arguments: %s", err))
			g.Assert(action.Arguments).Equal(expectedArguments)
		})

		g.It("should error on unsupported action", func() {
			action := Action{
				Name:         "bad-action",
				RawArguments: []byte(`{"name":"bad-action","sources":["some/source"],"path":"some/path"}`),
			}
			err := parseArgs(&action)
			g.Assert(err != nil).IsTrue("should have failed on unsupported action")
			g.Assert(err.Error()).Equal("action 'bad-action' not supported")
		})

		g.It("should parse raw upload arguments", func() {
			expectedArguments := artifactory.UploadArgs{
				Explode: true,
				Path:    "some/path",
				Sources: []string{"some/source"},
			}
			action := Action{
				Name:         "upload",
				RawArguments: []byte(`{"name":"upload","sources":["some/source"],"path":"some/path","explode":true}`),
			}
			err := parseArgs(&action)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to parse arguments: %s", err))
			g.Assert(action.Arguments).Equal(expectedArguments)
		})

		g.It("should parse raw delete arguments", func() {
			expectedArguments := artifactory.DeleteArgs{
				Path: "some/path",
			}
			action := Action{
				Name:         "delete",
				RawArguments: []byte(`{"name":"delete","path":"some/path"}`),
			}
			err := parseArgs(&action)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Failed to parse arguments: %s", err))
			g.Assert(action.Arguments).Equal(expectedArguments)
		})
	})
}
