package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/fileutil"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-io/go-utils/v2/pathutil"
	"github.com/bitrise-steplib/steps-authenticate-host-with-netrc/netrcutil"
)

const (
	gitHost = "bitbucket.org"
	apiHost = "api.bitbucket.org"
)

type ConfigsModel struct {
	Username       string          `env:"username,required"`
	AccessToken    stepconf.Secret `env:"access_token,required"`
	AtlassianEmail stepconf.Secret `env:"atlassian_email"`
}

func main() {
	logger := log.NewLogger()
	fileManager := fileutil.NewFileManager()

	failf := func(message string, args ...interface{}) {
		logger.Errorf(message, args...)
		os.Exit(1)
	}

	var configs ConfigsModel

	parser := stepconf.NewInputParser(env.NewRepository())
	if err := parser.Parse(&configs); err != nil {
		failf("Issue with input: %s", err)
	}

	stepconf.Print(configs)
	fmt.Println()

	netRC := netrcutil.New()

	logger.Infof("Other configs:")
	logger.Printf("- OutputPath: %s", netRC.OutputPth)

	fmt.Println()

	logger.Infof("Adding host config...")

	// Git over HTTPS on bitbucket.org authenticates with the Bitbucket username and the secret.
	netRC.AddItemModel(netrcutil.NetRCItemModel{Machine: gitHost, Login: configs.Username, Password: string(configs.AccessToken)})
	logger.Printf("- Added: %s", gitHost)

	// REST calls to api.bitbucket.org require the Atlassian account email when using a Bitbucket API
	// token, so it gets a separate entry with a different login. App passwords used the Bitbucket
	// username for both hosts, so the email is optional and the entry is skipped when it is empty.
	if configs.AtlassianEmail != "" {
		netRC.AddItemModel(netrcutil.NetRCItemModel{Machine: apiHost, Login: string(configs.AtlassianEmail), Password: string(configs.AccessToken)})
		logger.Printf("- Added: %s", apiHost)
	} else {
		logger.Printf("- Skipped: %s (no atlassian_email provided)", apiHost)
	}

	fmt.Println()

	logger.Infof("Writing .netrc file...")

	isExists, err := pathutil.NewPathChecker().IsPathExists(netRC.OutputPth)
	if err != nil {
		failf("Failed to check path (%s), error: %s", netRC.OutputPth, err)
	}

	if !isExists {
		logger.Printf("No .netrc file found at (%s), creating new...", netRC.OutputPth)

		if err := netRC.CreateFile(); err != nil {
			failf("Failed to write .netrc file, error: %s", err)
		}
	} else {
		logger.Warnf("File already exists at (%s)", netRC.OutputPth)

		backupPth := fmt.Sprintf("%s%s", strings.Replace(netRC.OutputPth, ".netrc", ".bk.netrc", -1), time.Now().Format("2006_01_02_15_04_05"))

		originalContent, err := os.ReadFile(netRC.OutputPth)
		if err != nil {
			failf("Failed to read file (%s), error: %s", netRC.OutputPth, err)
		}
		if err := fileManager.WriteBytes(backupPth, originalContent); err != nil {
			failf("Failed to write file (%s), error: %s", backupPth, err)
		}
		logger.Printf("Backup created at: %s", backupPth)

		logger.Printf("Appending config to the existing .netrc file...")

		if err := netRC.Append(); err != nil {
			failf("Failed to write .netrc file, error: %s", err)
		}
	}
	logger.Donef("Success")
}
