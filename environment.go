package environment

import (
	"os"
)

var env string

const DEFAULT_VAR = "SERVER_ENV"
const PRODUCTION = "production"
const DEVELOPMENT = "development"
const DEMO = "demo"
const STAGING = "staging"
const TEST = "test"

func Setup(environmentVariable string) string {
	envVar := environmentVariable
	if envVar == "" {
		envVar = DEFAULT_VAR
	}

	env = os.Getenv(envVar)
	if env == "" {
		env = DEVELOPMENT
	}

	if Unknown() {
		os.Stderr.WriteString("Could not determine runtime environment, continuing in development\n")
		env = DEVELOPMENT
	}
	os.Stdout.WriteString("The " + String() + " environment is starting up...\n")

	return env
}

func String() string {
	if env == "" {
		return Setup("")
	}
	return env
}

func Development() bool {
	return String() == DEVELOPMENT
}

func Staging() bool {
	return String() == STAGING
}

func Demo() bool {
	return String() == DEMO
}

func Production() bool {
	return String() == PRODUCTION
}

func Test() bool {
	return String() == TEST
}

func Known() bool {
	return Production() || Demo() || Staging() || Test() || Development()
}

func Local() bool {
	return Development() || Test()
}

func Remote() bool {
	return !Local()
}

func Unknown() bool {
	return !Known()
}
