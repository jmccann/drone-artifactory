package artifactory

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

type (
	// PropsArgs are arguments for managing properties on files
	PropsArgs struct {
		Path      string
		Recursive bool
		Props     []Prop
	}

	// Prop is a property on a file
	Prop struct {
		Name   string
		Value  string
		Values []string
	}
)

// SetProps on items in Artifactory
func (a Artifactory) SetProps(args PropsArgs) error {
	searchArgs := SearchArgs{
		Path:      args.Path,
		Recursive: args.Recursive,
	}

	resultItems, err := a.search(searchArgs)

	if err != nil {
		return err
	}

	propsParams := services.NewPropsParams()
	propsParams.Items = resultItems
	propsParams.Props = args.PropsString()

	logrus.WithFields(logrus.Fields{
		"items": resultItems,
		"props": args.PropsString(),
	}).Debug()

	logrus.Infof("Updating props for %s", args.Path)
	_, err = a.client.SetProps(propsParams)

	if err != nil {
		return err
	}

	return nil
}

// Validate the delete arguments
func (p PropsArgs) Validate() error {
	if len(p.Path) == 0 {
		return fmt.Errorf("No path provided")
	}

	if len(p.Props) == 0 {
		return fmt.Errorf("No props provided")
	}

	return nil
}

// PropsString a string representation of Props
func (p PropsArgs) PropsString() string {
	var s []string

	for _, prop := range p.Props {
		s = append(s, prop.String())
	}

	return strings.Join(s, ";")
}

func (p Prop) String() string {
	// TODO: Merge if both a single and multi value provided

	if p.Value == "" {
		return fmt.Sprintf("%s=%s", p.Name, strings.Join(p.Values, ","))
	}

	return fmt.Sprintf("%s=%s", p.Name, p.Value)
}
