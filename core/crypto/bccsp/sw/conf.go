package sw

import (
	"errors"
	"path/filepath"

	"os"

	"github.com/spf13/viper"
)

type config struct {
	keystorePath string

	configurationPathProperty string
}

func (conf *config) init() error {
	conf.configurationPathProperty = "security.bccsp.default.keyStorePath"

	// Check mandatory fields
	var rootPath string
	if err := conf.checkProperty(conf.configurationPathProperty); err != nil {
		logger.Warning("'security.bccsp.default.keyStorePath' not set. Using the default directory [%s] for temporary files", os.TempDir())
		rootPath = os.TempDir()
	} else {
		rootPath = viper.GetString(conf.configurationPathProperty)
	}
	logger.Infof("Root Path [%s]", rootPath)
	// Set configuration path
	rootPath = filepath.Join(rootPath, "crypto")

	// Set ks path
	conf.keystorePath = filepath.Join(rootPath, "ks")

	return nil
}

func (conf *config) checkProperty(property string) error {
	res := viper.GetString(property)
	if res == "" {
		return errors.New("Property not specified in configuration file. Please check that property is set: " + property)
	}
	return nil
}

func (conf *config) getKeyStorePath() string {
	return conf.keystorePath
}

func (conf *config) getPathForAlias(alias, suffix string) string {
	return filepath.Join(conf.getKeyStorePath(), alias+"_"+suffix)
}
