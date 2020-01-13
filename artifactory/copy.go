package artifactory

import (
	"fmt"

	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/sirupsen/logrus"
)

type (
	// CopyArgs are arguments for searching files
	CopyArgs struct {
		Flat      bool
		Source    string
		Target    string
		Recursive bool
	}
)

// Copy a file inside artifactory
func (a Artifactory) Copy(args CopyArgs) error {
	params := services.NewMoveCopyParams()
	params.ArtifactoryCommonParams = &utils.ArtifactoryCommonParams{
		Pattern:   args.Source,
		Recursive: args.Recursive,
		Target:    args.Target,
	}
	params.Flat = args.Flat

	logrus.WithFields(logrus.Fields{
		"action": "copy",
		"args":   args,
	}).Debug()

	logrus.Infof("Copy %s to %s", args.Source, args.Target)
	_, _, err := a.client.Copy(params)

	if err != nil {
		return err
	}

	return nil
}

// Validate the copy arguments
func (c CopyArgs) Validate() error {
	if len(c.Source) == 0 {
		return fmt.Errorf("No source provided")
	}

	if len(c.Target) == 0 {
		return fmt.Errorf("No target provided")
	}

	return nil
}
