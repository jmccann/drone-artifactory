package artifactory

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestUpload(t *testing.T) {
	g := Goblin(t)

	g.Describe("Validate UploadArgs", func() {
		args := u

		g.BeforeEach(func() {
			args = u
		})

		g.It("should fail on sources", func() {
			args.Sources = []string{}
			err := args.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No sources provided"))
		})
		g.It("should fail on path", func() {
			args.Path = ""
			err := args.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No path provided"))
		})
	})
}

var (
	u = UploadArgs{
		Path:    "thekey/with/path",
		Sources: []string{"main.go", "fixtures/*"},
	}
)
