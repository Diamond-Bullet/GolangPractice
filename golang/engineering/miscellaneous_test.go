package engineering

import (
	"GolangPractice/utils/logger"
	"encoding/base64"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/buger/jsonparser"
	"github.com/bwmarrin/snowflake"
	"github.com/gocarina/gocsv"
	"github.com/google/go-cmp/cmp"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Introduction document: https://www.cnblogs.com/zhichaoma/p/12640064.html

func TestUnicodeDecode(t *testing.T) {
	uContent := "\u62a5\u544a\u6587\u4ef6\n\u5185\u5bb9\u5f02\u5e38"

	text, err := strconv.Unquote(strings.Replace(strconv.Quote(uContent), `\\u`, `\u`, -1))
	fmt.Println(text, err)
}

func TestBase64(t *testing.T) {
	data := "种豆得豆"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
}

func TestRegExp(t *testing.T) {
	// match strings starts with things like `(1234, '`
	r, err := regexp.Compile(`^\([0-9]*[1-9][0-9]*, '`)
	if err != nil {
		color.Redln(err)
		return
	}

	color.Blueln("r.MatchString(\"(1234, 'Good Good'\"):", r.MatchString("(1234, 'Good Good'"))
	color.Blueln("r.FindStringIndex(\"(1234, 'Good Good'\")", r.FindStringIndex("(1234, 'Good Good'"))
}

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
	logger.Info(string(val), " err: ", err)
}

// https://github.com/tealeg/xlsx #seems like the official's.
// https://github.com/qax-os/excelize is the most starred repo by now
func TestExelSheet(t *testing.T) {
	excelFileName := "/foo.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		logger.Error(err)
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

func TestCSV(t *testing.T) {
	// Go team also provides a built-in package `encoding/csv` for csv file.
	type Person struct {
		Name string `csv:"name"`
		Age  int    `csv:"age"`
		City string `csv:"city"`
	}

	// Sample data to write to CSV
	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "Los Angeles"},
		{"Charlie", 35, "Chicago"},
	}

	// Create a CSV file
	file, err := os.Create("people.csv")
	if err != nil {
		log.Fatal("Unable to create file:", err)
	}
	defer file.Close()

	// Marshal the data to the CSV file
	if err := gocsv.MarshalFile(&people, file); err != nil {
		log.Fatal("Error marshaling to CSV:", err)
	}

	log.Println("CSV file written successfully.")
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
		logger.Error(err)
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

// TODO https://github.com/jlaffaye/ftp

func TestToml(t *testing.T) {
	// Initialize a variable to hold the parsed configuration
	var data struct{
		Age int
		Name string
	}

	// Read and parse the TOML file
	if _, err := toml.DecodeFile("config.toml", &data); err != nil {
		logger.Error("Error parsing TOML file:", err)
		return
	}
}
