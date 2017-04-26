package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bitrise-io/depman/pathutil"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-authenticate-host-with-netrc/netrcutil"
	"github.com/bitrise-tools/go-steputils/input"
)

// ConfigsModel ...
type ConfigsModel struct {
	BitbucketHosts []string
	Username       string
	AccessToken    string
}

func createConfigsModelFromEnvs() ConfigsModel {
	return ConfigsModel{
		BitbucketHosts: []string{"bitbucket.org", "api.bitbucket.org"},
		Username:       os.Getenv("username"),
		AccessToken:    os.Getenv("access_token"),
	}
}

func (configs *ConfigsModel) validate() error {
	if err := input.ValidateIfNotEmpty(configs.Username); err != nil {
		return errors.New("Username, parameter not specified")
	}
	if err := input.ValidateIfNotEmpty(configs.AccessToken); err != nil {
		return errors.New("AppPassword, parameter not specified")
	}
	return nil
}

func (configs *ConfigsModel) print() {
	log.Infof("Configs:")
	log.Printf("- Username: %s", configs.Username)
	log.Printf("- AppPassword: %s", input.SecureInput(configs.AccessToken))
}

func failf(message string, args ...interface{}) {
	log.Errorf(message, args...)
	os.Exit(1)
}

func main() {
	configs := createConfigsModelFromEnvs()

	configs.print()
	fmt.Println()

	if err := configs.validate(); err != nil {
		failf("Issue with input: %s", err)
	}

	netRC := netrcutil.New()

	log.Infof("Other configs:")
	log.Printf("- OutputPath: %s", netRC.OutputPth)

	fmt.Println()

	log.Infof("Adding host config...")
	for _, host := range configs.BitbucketHosts {
		netRC.AddItemModel(netrcutil.NetRCItemModel{Machine: host, Login: configs.Username, Password: configs.AccessToken})
		log.Printf("- Added: %s", host)
	}

	fmt.Println()

	log.Infof("Writing .netrc file...")

	isExists, err := pathutil.IsPathExists(netRC.OutputPth)
	if err != nil {
		failf("Failed to check path (%s), error: %s", netRC.OutputPth, err)
	}

	if !isExists {
		log.Printf("No .netrc file found at (%s), creating new...", netRC.OutputPth)

		if err := netRC.CreateFile(); err != nil {
			failf("Failed to write .netrc file, error: %s", err)
		}
	} else {
		log.Warnf("File already exists at (%s)", netRC.OutputPth)

		backupPth := fmt.Sprintf("%s%s", strings.Replace(netRC.OutputPth, ".netrc", ".bk.netrc", -1), time.Now().Format("2006_01_02_15_04_05"))

		if originalContent, err := fileutil.ReadBytesFromFile(netRC.OutputPth); err != nil {
			failf("Failed to read file (%s), error: %s", netRC.OutputPth, err)
		} else if err := fileutil.WriteBytesToFile(backupPth, originalContent); err != nil {
			failf("Failed to write file (%s), error: %s", backupPth, err)
		} else {
			log.Printf("Backup created at: %s", backupPth)
		}

		log.Printf("Appending config to the existing .netrc file...")

		if err := netRC.Append(); err != nil {
			failf("Failed to write .netrc file, error: %s", err)
		}
	}
	log.Donef("Success")
}
