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
	os.Stdout.WriteString("The " + getEnvString() + " environment is starting up...\n")

	return env
}

func getEnvString() string {
	if env == "" {
		return Setup("")
	}
	return env
}

func String() string {
	return getEnvString()
}

func Development() bool {
	return getEnvString() == DEVELOPMENT
}

func Staging() bool {
	return getEnvString() == STAGING
}

func Demo() bool {
	return getEnvString() == DEMO
}

func Production() bool {
	return getEnvString() == PRODUCTION
}

func Test() bool {
	return getEnvString() == TEST
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
