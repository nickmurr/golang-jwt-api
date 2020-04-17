package config

import (
	"github.com/sirupsen/logrus"

	"os"
	"strconv"
)

var (
	PORT = 3001
	Log  = logrus.New()
)

func Load() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Log.Warn(err)
		PORT = 3001
	}

}
