package artifactory

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
)

type (
	// UploadArgs are arguments for an Upload
	UploadArgs struct {
		DryRun      bool `json:"dryrun"`
		Explode     bool
		Flat        bool
		IncludeDirs bool `json:"include_dirs"`
		Path        string
		Recursive   bool
		Regexp      bool
		Sources     []string
	}
)

// Upload files to artifactory
func (a Artifactory) Upload(args UploadArgs) error {
	err := args.Validate()

	if err != nil {
		return err
	}

	for _, source := range args.Sources {
		logrus.Infof("Uploading %s to %s", source, args.Path)

		params := services.NewUploadParams()

		params.ArtifactoryCommonParams = &utils.ArtifactoryCommonParams{
			IncludeDirs: args.IncludeDirs,
			Pattern:     source,
			Recursive:   args.Recursive,
			Regexp:      args.Regexp,
			Target:      args.Path,
		}

		params.ExplodeArchive = args.Explode
		params.Flat = args.Flat
		params.Retries = 3

		_, n, m, err := a.client.UploadFiles(params)
		logrus.WithFields(logrus.Fields{
			"Success": n,
			"Failed":  m,
		}).Info("Upload Complete")

		if err != nil {
			logrus.Errorf("there was an error: %v", err)
			return err
		}
	}

	return nil
}

// Validate the upload arguments
func (u UploadArgs) Validate() error {
	if len(u.Sources) == 0 {
		return fmt.Errorf("No sources provided")
	}

	if len(u.Path) == 0 {
		return fmt.Errorf("No path provided")
	}

	return nil
}
