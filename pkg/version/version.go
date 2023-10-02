package version

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

var PROJECT_VERSION string

func init() {
	file, err := os.Open("./.version")
	if err != nil {
		log.WithError(err).Error("Error to open version file")
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.WithError(err).Error("Error to close version file")
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		PROJECT_VERSION = scanner.Text()

		break
	}

	if PROJECT_VERSION == "" {
		log.WithField("project_version", PROJECT_VERSION).Error("Project version is empty")
	}
}
