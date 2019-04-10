package artifactory

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestArtifactory(t *testing.T) {
	g := Goblin(t)

	g.Describe("New", func() {
		g.It("should create a client", func() {
			_, err := New(c)

			g.Assert(err == nil).IsTrue(fmt.Sprintf("should not error: %s", err))
		})

		g.It("should append / to end of URL", func() {
			arti, err := New(c)

			g.Assert(err == nil).IsTrue(fmt.Sprintf("should not error: %s", err))
			g.Assert(arti.client.GetConfig().GetArtDetails().GetUrl()).Equal("http://company.com/artifactory/")
		})
	})
}
