package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

type (
	// Config for plugin
	Config struct {
		Password string
		APIKey   string
		URL      string
		Username string
	}

	// Plugin struct
	Plugin struct {
		Config     Config
		UploadArgs UploadArgs
	}

	// UploadArgs are arguments for uploading files
	UploadArgs struct {
		DryRun      bool
		Explode     bool
		Flat        bool
		IncludeDirs bool
		Path        string
		Recursive   bool
		Regexp      bool
		Sources     []string
	}
)

const jfrogExe = "/bin/jfrog"

// Exec run the plugin
func (p Plugin) Exec() error {
	err := p.Config.validate()

	if err != nil {
		return err
	}

	err = p.UploadArgs.validate()

	if err != nil {
		return err
	}

	err = executeCommand(commandVersion(), false) // jfrog --version

	if err != nil {
		return err
	}

	logrus.Info("Creating CLI config")
	err = executeCommand(commandConfig(p.Config), true) // jfrog rt config
	if err != nil {
		return err
	}

	// jfrog rt upload
	for _, source := range p.UploadArgs.Sources {
		err = executeCommand(commandUpload(source, p.UploadArgs), false)

		if err != nil {
			return err
		}
	}

	return nil
}

// helper function to create the jfrog version command.
func commandVersion() *exec.Cmd {
	return exec.Command(jfrogExe, "--version")
}

// helper function to create the jfrog rt config command.
func commandConfig(c Config) *exec.Cmd {
	if len(c.APIKey) > 0 {
		return exec.Command(
			jfrogExe,
			"rt",
			"config",
			"--interactive=false",
			"--url", c.URL,
			"--user", c.Username,
			"--apikey", c.APIKey, "--enc-password=false",
		)
	}
	return exec.Command(
		jfrogExe,
		"rt",
		"config",
		"--interactive=false",
		"--url", c.URL,
		"--user", c.Username,
		"--password", c.Password, "--enc-password=false",
	)
}

// helper function to create the jfrog rt upload command.
func commandUpload(source string, args UploadArgs) *exec.Cmd {
	return exec.Command(
		jfrogExe,
		"rt",
		"upload",
		fmt.Sprintf("--dry-run=%t", args.DryRun),
		fmt.Sprintf("--explode=%t", args.Explode),
		fmt.Sprintf("--flat=%t", args.Flat),
		fmt.Sprintf("--include-dirs=%t", args.IncludeDirs),
		fmt.Sprintf("--recursive=%t", args.Recursive),
		fmt.Sprintf("--regexp=%t", args.Regexp),
		source,
		args.Path,
	)
}

func executeCommand(cmd *exec.Cmd, sensitive bool) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if !sensitive {
		trace(cmd)
	}

	err := cmd.Run()

	return err
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

func (c Config) validate() error {
	if len(c.Username) == 0 {
		return fmt.Errorf("No username provided")
	}

	if len(c.Password) == 0 && len(c.APIKey) == 0 {
		return fmt.Errorf("No ApiKey or Password provided")
	}

	if len(c.URL) == 0 {
		return fmt.Errorf("No url provided")
	}

	return nil
}

func (u UploadArgs) validate() error {
	if len(u.Sources) == 0 {
		return fmt.Errorf("No sources provided")
	}

	if len(u.Path) == 0 {
		return fmt.Errorf("No path provided")
	}

	return nil
}
