package logging

import (
	"fmt"
	"log"
	"os"
	"path"
)

func InitLogging() (*os.File, *log.Logger, error)  {
	logFile := path.Join(os.TempDir(), fmt.Sprintf("postx.log"))
	f, err := os.OpenFile(
		logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise log file.\n")
	}
	LstdFlags := log.Ldate | log.Lshortfile
    logger := log.New(f,"postx ", LstdFlags)
    logger.Println("initialise logger")
    return f, logger, nil 
}
