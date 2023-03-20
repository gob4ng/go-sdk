package file

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)
func NewConfig(fileName string) *error {

	splits := strings.Split(filepath.Base(fileName), ".")
	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(fileName))

	if err := viper.ReadInConfig();err != nil {
		return &err
	}

	return nil
}

func MustGetString(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func MustGetInt(key string) int {
	checkKey(key)
	return viper.GetInt(key)
}

func MustGetBool(key string) bool {
	checkKey(key)
	return viper.GetBool(key)
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		os.Exit(1)
	}
}
