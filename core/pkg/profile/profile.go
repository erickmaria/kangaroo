package profile

import (
	"errors"
	"flag"
	"os"
	"strings"
)

type Profile struct {
	Configs struct {
		Path   string   `yaml:"path"`
		File   string   `yaml:"file"`
		Target []string `yaml:"target"`
	} `yaml:"configs"`
}

var (
	profileEnv           string = "KANGAROO_PROFILE_ACTIVE"
	profileDefaultKey    string = "{default}"
	profileConfigFileKey string = "#{target}#"
)

func selectProfile(profile Profile) (string, error) {

	var active string
	var getDefault bool = false

	active = flagProfile()
	if active == "" {
		active = os.Getenv(profileEnv)
	}

	if active == "" {
		getDefault = true
	}

	for _, target := range profile.Configs.Target {

		targetSplit := strings.Split(target, " ")
		if getDefault {
			if len(targetSplit) >= 2 {
				if targetSplit[1] == profileDefaultKey {
					active = targetSplit[0]

					return active, nil
				}
			}
			continue
		}
		if targetSplit[0] == active {
			return active, nil
		}
	}

	if getDefault {
		return active, errors.New("target default not was informed, put {default} on profile target to load as default")
	}

	return active, errors.New("profile target '" + active + "' not found")

}

func flagProfile() string {

	var profileActive string

	flag.StringVar(&profileActive, "profile", "", "load application profile")
	flag.Parse()

	return profileActive
}
