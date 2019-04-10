package artifactory

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestDelete(t *testing.T) {
	g := Goblin(t)

	g.Describe("Validate DeleteArgs", func() {
		args := d

		g.BeforeEach(func() {
			args = d
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
	d = DeleteArgs{
		DryRun: true,
		Path:   "thekey/with/path",
	}
)
