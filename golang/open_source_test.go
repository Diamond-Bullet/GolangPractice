package golang

import (
	"GolangPractice/utils/logger"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/buger/jsonparser"
	"github.com/bwmarrin/snowflake"
	"github.com/google/go-cmp/cmp"
	"github.com/gookit/color"
	"github.com/panjf2000/ants/v2"
	pkgerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// https://github.com/buger/jsonparser #Parse json data dynamically，support Get、Set。
func TestJsonParser(t *testing.T) {
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	val, err := jsonparser.Set(data, []byte("https://github.com"), "person", "avatars", "[0]", "backup_url")
	logger.Infoln(string(val), " err: ", err)
}

// https://github.com/tealeg/xlsx #seems like the official's.
// https://github.com/qax-os/excelize is the most starred repo by now
func TestHandleExelSheet(t *testing.T) {
	excelFileName := "/foo.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		logger.Errorln(err)
		return
	}

	// traverse the sheet.
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s ", text)
			}
		}
		println()
	}

	// read individual cell.
	xlFile.Sheet["规则表名"].Cell(0, 2).String()
	xlFile.Sheets[0].Row(0).Cells[2].String()
	// write specified cell.
	xlFile.Sheets[0].Rows[0].Cells[2].SetBool(true)
	// add sell.
	newCell := xlFile.Sheets[0].Row(0).AddCell()
	newCell.SetDate(time.Now())
}

func TestColorfulPrint(t *testing.T) {
	// https://github.com/gookit/color
	// basic colors below
	// for more information, see ReadMe.MD
	color.Blueln("I am great Blue")
	color.Magentaln("color Magenta")
	color.New(color.FgRed, color.BgCyan).Println("FgRed BgCyan")
	// partial rendering
	fmt.Println(color.FgRed.Render("red"), "line")

	// 256 colors
	color.C256(132).Println("what the fucking color132")
	color.S256(110, 120).Println("fg110 bg120")

	// RGB colors
	color.RGB(100, 200, 30).Println("color.RGB r100 g200 b30")
	color.HEX("#1976D2").Println("color.HEX #1976D2")
	color.RGBStyleFromString("170,187,204", "70,87,4").Println("RGBStyleFromString")

	// work on linux or macOS.
	// https://stackoverflow.com/questions/5947742/how-to-change-the-output-color-of-echo-in-linux
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	fmt.Println(colorRed, "test")
	fmt.Println(colorGreen, "test")
	fmt.Println(colorYellow, "test")
	fmt.Println(colorBlue, "test")
	fmt.Println(colorPurple, "test")
	fmt.Println(colorWhite, "test")
	fmt.Println(colorCyan, "test", colorReset)
	fmt.Println()
}

// snowflake
// see uuid (universally unique identifier) at https://github.com/search?q=uuid+language%3AGo+&type=repositories&s=stars&o=desc
func TestSnowFlake(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Errorln(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	logger.Infof("Int64  ID: %d", id)
	logger.Infof("String ID: %s", id)
	logger.Infof("Base2  ID: %s", id.Base2())
	logger.Infof("Base64 ID: %s", id.Base64())

	// Print out the ID's timestamp
	logger.Infof("ID Time  : %d", id.Time())

	// Print out the ID's node number
	logger.Infof("ID Node  : %d", id.Node())

	// Print out the ID's sequence number
	logger.Infof("ID Step  : %d", id.Step())

	// Generate and print, all in one.
	logger.Infof("ID       : %d", node.Generate().Int64())
}

// https://github.com/pkg/errors error with stack trace
// alternative: https://github.com/go-errors/errors
// learn about new error handling draft `Go2 errors` by Go team.
func TestStackError(t *testing.T) {
	err := pkgerrors.Errorf("err: %s", "i want to bring out an error")
	logger.Infoln(err)
	logger.Infof("%+v", err)

	err1 := pkgerrors.Wrap(err, "err1")
	logger.Infof("%+v", err1)

	err2 := StackError1()
	logger.Infof("%+v", err2)
}

func StackError1() error {
	return StackError2()
}

func StackError2() error {
	return pkgerrors.New("error here")
}

// provided by Go team. simply wrap an error with new prefix.
func TestWrapError(t *testing.T) {
	err := errors.New("error here")
	err1 := fmt.Errorf("layer1: %w", err)
	err2 := fmt.Errorf("layer2: %w", err1)
	logger.Errorln(err2)
}

// https://golang.org/x/sync/errgroup
// Slightly different from `go func()...`, it handles errors.
// ErrGroup does NOT offer the functionality of recovering from panic.
func TestErrGroup(t *testing.T) {
	g := new(errgroup.Group)

	g.Go(func() error {
		return nil
	})
	g.Go(func() error {
		return errors.New("very good")
	})

	err := g.Wait()
	if err != nil {
		logger.Errorln(err)
	}
}

// https://github.com/google/go-cmp/cmp
// Compare 2 objects according to customized rules which are implemented in `Equal()` method of the type.
func TestCmp(t *testing.T) {
	cmp.Equal(1, 2)
}

// https://github.com/sirupsen/logrus Logging component.
// pretty nice, but I don't like the format.
func TestLogRus(t *testing.T) {
	logrus.Errorln("error")

	localLogger := logrus.New()
	localLogger.SetNoLock()
	localLogger.SetReportCaller(true)
	localLogger.SetFormatter(&logrus.TextFormatter{})

	localLogger.Errorln("error")
}

// https://go.uber.org/zap Logging Component.
func TestUberZap(t *testing.T) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}

// TODO　https://github.com/jlaffaye/ftp

// https://github.com/panjf2000/ants/v2 Goroutine Pool.
func TestGoroutinePool(t *testing.T) {
	const TaskNum = 1e3

	pool, err := ants.NewPool(TaskNum / 10)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer pool.Release()

	waitGroup := new(sync.WaitGroup)

	for i := 0; i < TaskNum; i++ {
		_ = pool.Submit(func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			fmt.Println("Good task")
		})
	}

	waitGroup.Wait()
}
