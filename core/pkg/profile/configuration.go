package profile

import (
	"github.com/erickmaria/kangaroo/core/pkg/handler"
	"github.com/erickmaria/kangaroo/core/pkg/utils"
)

var (
	profileFile string = "profile.yaml"
)

func Init(application interface{}, profilesPath string, envPrefix string) error {

	var err error

	// Profiles configs

	var profile Profile
	profilePathFile := profilesPath + "/" + profileFile

	err = utils.CheckFileExist(profilePathFile)
	if handler.Validate(err) {
		return err
	}

	loadProfilesError := utils.YamlUnmarshal(&profile, profilePathFile)

	if handler.Validate(loadProfilesError) {
		return loadProfilesError
	}

	active, selectProfileError := selectProfile(profile)
	if handler.Validate(selectProfileError) {
		return selectProfileError
	}

	// Application configs

	applicationFile := getProfileApplication(profile, active)

	err = utils.CheckFileExist(applicationFile)
	if handler.Validate(err) {
		return err
	}

	err = utils.YamlUnmarshal(application, applicationFile)
	if handler.Validate(err) {
		return err
	}

	err = SyncEnv(application, envPrefix)
	if handler.Validate(err) {
		return err
	}

	return nil
}
