package main

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jmccann/drone-artifactory/artifactory"
)

type (
	// Action to perform
	Action struct {
		Name         string          `json:"name"`
		RawArguments json.RawMessage `json:"args"`
		Arguments    interface{}
	}

	// Plugin struct
	Plugin struct {
		Actions []Action
		Config  artifactory.Config
	}
)

// Exec run the plugin
func (p Plugin) Exec() error {
	err := p.Config.Validate()

	if err != nil {
		return err
	}

	logrus.Info("Creating Client")
	arti, err := artifactory.New(p.Config)

	if err != nil {
		return err
	}

	logrus.Info("Pinging Artifactory")
	err = arti.Ping()

	if err != nil {
		return err
	}

	logrus.Info("Validating Actions")
	for idx := range p.Actions {
		err = parseArgs(&p.Actions[idx])

		if err != nil {
			return err
		}
	}

	logrus.Info("Running Actions")
	for _, action := range p.Actions {
		err = do(arti, action)

		if err != nil {
			return err
		}
	}

	return nil
}

func do(arti *artifactory.Artifactory, action Action) error {
	if action.Name == "delete" {
		return arti.Delete(action.Arguments.(artifactory.DeleteArgs))
	}

	if action.Name == "upload" {
		return arti.Upload(action.Arguments.(artifactory.UploadArgs))
	}

	return fmt.Errorf("action '%s' not supported", action)
}

func parseArgs(action *Action) error {
	if action.Name == "delete" {
		var deleteArgs artifactory.DeleteArgs
		err := json.Unmarshal(action.RawArguments, &deleteArgs)

		if err != nil {
			return err
		}

		action.Arguments = deleteArgs
		return deleteArgs.Validate()
	}

	if action.Name == "upload" {
		var uploadArgs artifactory.UploadArgs
		err := json.Unmarshal(action.RawArguments, &uploadArgs)

		if err != nil {
			return err
		}

		action.Arguments = uploadArgs
		return uploadArgs.Validate()
	}

	return fmt.Errorf("action '%s' not supported", action.Name)
}
