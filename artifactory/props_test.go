package artifactory

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestProps(t *testing.T) {
	g := Goblin(t)

	g.Describe("PropsArgs", func() {
		propsArgs := basePropsArgs

		g.BeforeEach(func() {
			propsArgs = basePropsArgs
		})

		g.Describe("PropsString()", func() {
			g.It("should generate a string for multiple props", func() {
				expectedString := "some-name=some-value,another-value;some-other-name=some-other-value,another-other-value"
				g.Assert(basePropsArgs.PropsString()).Equal(expectedString)
			})
		})

		g.Describe("Validate()", func() {
			g.It("should fail on path", func() {
				propsArgs.Path = ""
				err := propsArgs.Validate()
				g.Assert(err != nil).IsTrue()
				g.Assert(err).Equal(fmt.Errorf("No path provided"))
			})

			g.It("should fail on props", func() {
				propsArgs.Props = nil
				err := propsArgs.Validate()
				g.Assert(err != nil).IsTrue()
				g.Assert(err).Equal(fmt.Errorf("No props provided"))
			})
		})
	})

	g.Describe("Prop", func() {
		g.It("should generate string for single value", func() {
			expectedString := "some-name=some-value"
			prop := Prop{
				Name:  "some-name",
				Value: "some-value",
			}
			g.Assert(prop.String()).Equal(expectedString)
		})

		g.It("should generate string for multiple values", func() {
			expectedString := "some-name=some-value,another-value"
			prop := Prop{
				Name:   "some-name",
				Values: []string{"some-value", "another-value"},
			}
			g.Assert(prop.String()).Equal(expectedString)
		})
	})
}

var (
	basePropsArgs = PropsArgs{
		Path: "thekey/with/path",
		Props: []Prop{
			Prop{
				Name:   "some-name",
				Values: []string{"some-value", "another-value"},
			},
			Prop{
				Name:   "some-other-name",
				Values: []string{"some-other-value", "another-other-value"},
			},
		},
	}
)
