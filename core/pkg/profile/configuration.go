package profile

import (
	"github.com/erickmaria/kangaroo/core/pkg/handler"
	"github.com/erickmaria/kangaroo/core/pkg/utils"
)

var (
	profileFile string = "profile.yaml"
)

func Init(properties interface{}, profilesPath string, envPrefix string) error {

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

	// properties configs

	propertiesFile := getProfileProperties(profile, active)

	err = utils.CheckFileExist(propertiesFile)
	if handler.Validate(err) {
		return err
	}

	if properties != nil {
		err = utils.YamlUnmarshal(properties, propertiesFile)
		if handler.Validate(err) {
			return err
		}

		err = StructSyncEnv(properties, envPrefix)
		if handler.Validate(err) {
			return err
		}

		return nil
	}

	var ppt interface{}
	err = utils.YamlUnmarshal(&ppt, propertiesFile)
	if handler.Validate(err) {
		return err
	}

	err = MapSyncEnv(ppt, envPrefix)
	if handler.Validate(err) {
		return err
	}

	return nil
}
