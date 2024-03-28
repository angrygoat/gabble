package gabble

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

/*
configuration of gabble via viper
*/

type GabbleConfiguration struct {
	ConnectionTimeoutSeconds   int
	ConnectionKeepaliveSeconds int
	ConnectionUseMultipathTcp  bool
}

/*
configuration for test suites
*/
type GabbleTestConfiguration struct {
	Host             string
	Zone             string
	Port             int
	DefaultResource  string
	User1            string
	Password1        string
	User2            string
	Password2        string
	User3            string
	Password3        string
	RodsUser         string
	RodsPassword     string
	LocalScratchPath string
	FederationTest   bool
	RemoteRescTest   bool
}

// call to establish the configuration from properties
func EstablishGabbleConfig() (GabbleConfiguration, error) {
	log.Info().Msg("EstablishGabbleConfig()")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("gabble") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	err := viper.ReadInConfig()

	if err != nil {
		log.Error().Msg("unable to read gabble.yaml")
		return GabbleConfiguration{}, fmt.Errorf("EstablishGabbleConfig: failed reading gabble.config: %w", err)

		//Msg("Can't find the file .env : ", err)
	}

	gabbleConfiguration := GabbleConfiguration{}
	err = viper.Unmarshal(&gabbleConfiguration)

	if err != nil {
		log.Error().Msg("unable to unmarshal gabbleConfiguration")
		return GabbleConfiguration{}, fmt.Errorf("EstablishGabbleConfig: failed unmarshalling gabble.config: %w", err)
	}

	return gabbleConfiguration, nil

}

// call to establish the test configuration from properties
func EstablishGabbleTestConfig() (GabbleTestConfiguration, error) {
	log.Info().Msg("EstablishGabbleTestConfig()")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("testing") // Register config file name (no extension)
	viper.SetConfigType("yaml")    // Look for specific type
	err := viper.ReadInConfig()

	if err != nil {
		log.Error().Msg("unable to read gabble testing.yaml")
		return GabbleTestConfiguration{}, fmt.Errorf("EstablishGabbleConfig: failed reading testing.yaml: %w", err)

		//Msg("Can't find the file .env : ", err)
	}

	gabbleTestConfiguration := GabbleTestConfiguration{}
	err = viper.Unmarshal(&gabbleTestConfiguration)

	if err != nil {
		log.Error().Msg("unable to unmarshal gabbleTestConfiguration")
		return GabbleTestConfiguration{}, fmt.Errorf("EstablishGabbleTestConfig: failed unmarshalling testing.config: %w", err)
	}

	return gabbleTestConfiguration, nil

}
