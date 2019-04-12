package artifactory

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestCopy(t *testing.T) {
	g := Goblin(t)

	g.Describe("Validate CopyArgs", func() {
		args := baseCopyArgs

		g.BeforeEach(func() {
			args = baseCopyArgs
		})

		g.It("should fail on source", func() {
			args.Source = ""
			err := args.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No source provided"))
		})

		g.It("should fail on path", func() {
			args.Target = ""
			err := args.Validate()
			g.Assert(err != nil).IsTrue()
			g.Assert(err).Equal(fmt.Errorf("No target provided"))
		})
	})
}

var (
	baseCopyArgs = CopyArgs{
		Source: "some-source.file",
		Target: "thekey/with/path",
	}
)
