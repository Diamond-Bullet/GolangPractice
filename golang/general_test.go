package golang

import (
	"fmt"
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
