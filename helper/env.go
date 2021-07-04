package helper

import (
	"banking/errs"
	"fmt"
	"github.com/spf13/viper"
)

func GetEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		errs.NewUnexpectedError(fmt.Sprintf("Error while reading config file %s", err))
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		errs.NewUnexpectedError("Invalid type assertion")
	}
	return value
}
