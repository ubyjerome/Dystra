package routine

import (
	"fmt"
	"ubyjerome/dystra/logger"
	"ubyjerome/dystra/utilities"
	// "github.com/robfig/cron"
)

func ScanForIncludedFolders(){
	File, err := utilities.FetchFile("./.dystra.config.json")
	// scanSchedule := cron.New()
	logger.Info("Folder Scan Routine Started")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(File)
}