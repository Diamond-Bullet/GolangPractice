package engineering

import (
	"GolangPractice/utils/logger"
	"fmt"
	"syscall"
	"testing"
	"time"
)

// see more layouts in src/time/format.go
const (
	TimeFormatDate = "2006-01-02"
	TimeFormatDate2 = "20060102"

	TimeFormatTime = "2006-01-02 15:04:05"
	TimeFormatTime2 = "2006/01/02 15:04:05"

	TimeFormatMS = "2006-01-02 15:04:05.000"

	TimeFormatISO = "2006-01-02T15:04:05Z"
)

func TestTimeFormat(t *testing.T) {
	// Format current time.
	nowStamp := time.Now()

	timeStr := nowStamp.Format(TimeFormatMS)
	logger.Info("Current Time: %s", timeStr)
}

func TestTimeParse(t *testing.T) {
	// Assign location.
	loc, _ := time.LoadLocation("Local") // UTC， or other valid location name

	// another way to get local time zone is time.Local.
	// UTC is accessible if using time.UTC.
	timeStr, _ := time.ParseInLocation(TimeFormatTime2, "2021/11/02 15:04:05", loc)
	logger.Info("Local time.Time: %s", timeStr)

	NewYorkLoc, _ := time.LoadLocation("America/New_York")

	// you can either define a format by yourself or use predefined patterns in `time` package.
	timeStr, _ = time.ParseInLocation(time.RFC3339, "2021/11/02 15:04:05", NewYorkLoc)
	logger.Info("New York time.Time: %s", timeStr)
}

func TestTimeCalc(t *testing.T) {
	logger.Info(time.Now().AddDate(1,1,1).String())
}

func TestGetFileCreationTime(t *testing.T) {
	var st syscall.Stat_t
	fileFullPath := "/root/demo.txt"
	err := syscall.Stat(fileFullPath, &st)
	if err != nil {
		fmt.Println(err)
		return
	}
	createTime := time.Unix(st.Ctim.Sec, 0)
	fmt.Println("file is created at ", createTime)
}
