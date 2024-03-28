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
