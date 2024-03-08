package golang

import (
	"encoding/base64"
	"fmt"
	"github.com/gookit/color"
	"regexp"
	"strconv"
	"strings"
	"testing"
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
