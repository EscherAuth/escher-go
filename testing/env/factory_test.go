package env_test

import (
	"os"
	"testing"

	. "github.com/EscherAuth/escher/testing/env"
)

func TestSetEnvForTheTestEnvIsAlreadySet(t *testing.T) {
	os.Setenv("TEST_SET_ENV_FOR_THE_TEST", "example")
	defer os.Unsetenv("TEST_SET_ENV_FOR_THE_TEST")

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "example" {
		t.FailNow()
	}

	teardown := SetEnvForTheTest(t, "TEST_SET_ENV_FOR_THE_TEST", "BAZ")

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "BAZ" {
		t.FailNow()
	}

	teardown()

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "example" {
		t.FailNow()
	}

}
func TestSetEnvForTheTestEnvWasUnsetBeforehand(t *testing.T) {

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "" {
		t.FailNow()
	}

	teardown := SetEnvForTheTest(t, "TEST_SET_ENV_FOR_THE_TEST", "BAZ")

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "BAZ" {
		t.FailNow()
	}

	teardown()

	_, envKeyIsSet := os.LookupEnv("TEST_SET_ENV_FOR_THE_TEST")

	if envKeyIsSet {
		t.Fatal("env key for TEST_SET_ENV_FOR_THE_TEST should not be set")
	}

}

func TestUnsetEnvForTheTestEnvIsAlreadySet(t *testing.T) {
	os.Setenv("TEST_SET_ENV_FOR_THE_TEST", "example")
	defer os.Unsetenv("TEST_SET_ENV_FOR_THE_TEST")

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "example" {
		t.FailNow()
	}

	teardown := UnsetEnvForTheTest(t, "TEST_SET_ENV_FOR_THE_TEST")

	currentValue, isSet := os.LookupEnv("TEST_SET_ENV_FOR_THE_TEST")

	if isSet || currentValue != "" {
		t.Fatal("example value should be unset, but it is defined")
	}

	teardown()

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "example" {
		t.FailNow()
	}

}
func TestUnsetEnvForTheTestEnvWasUnsetBeforehand(t *testing.T) {

	if os.Getenv("TEST_SET_ENV_FOR_THE_TEST") != "" {
		t.FailNow()
	}

	teardown := UnsetEnvForTheTest(t, "TEST_SET_ENV_FOR_THE_TEST")

	currentValue, isSet := os.LookupEnv("TEST_SET_ENV_FOR_THE_TEST")

	if isSet || currentValue != "" {
		t.Fatal("example value should be unset, but it is defined")
	}

	teardown()

	_, envKeyIsSet := os.LookupEnv("TEST_SET_ENV_FOR_THE_TEST")

	if envKeyIsSet {
		t.Fatal("env key for TEST_SET_ENV_FOR_THE_TEST should not be set")
	}

}
