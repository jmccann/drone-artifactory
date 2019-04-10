package artifactory

import (
	"fmt"

	"github.com/Sirupsen/logrus"

	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
)

type (
	// DeleteArgs are arguments for Delete
	DeleteArgs struct {
		DryRun    bool
		Path      string
		Recursive bool
	}
)

// Delete items from Artifactory
func (a Artifactory) Delete(args DeleteArgs) error {
	params := services.NewDeleteParams()
	params.ArtifactoryCommonParams = &utils.ArtifactoryCommonParams{
		Pattern:   args.Path,
		Recursive: args.Recursive,
	}

	pathsToDelete, err := a.client.GetPathsToDelete(params)
	if err != nil {
		return err
	}
	logrus.Debugf("Paths to delete: %v", pathsToDelete)

	deletedFileCnt, err := a.client.DeleteFiles(pathsToDelete)

	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Deleted Files": deletedFileCnt,
	}).Info("Delete Complete")

	return nil
}

// Validate the delete arguments
func (d DeleteArgs) Validate() error {
	if len(d.Path) == 0 {
		return fmt.Errorf("No path provided")
	}

	return nil
}
