package coding

import (
	"fmt"
	"sort"
	"testing"
)

// Build maps or indexes to avoid if-else or switch-case statements.

func TestTableDriven (t *testing.T) {
	// direct indexes
	fmt.Println(GetMonth(2, 2021))

	// multi-dimensional indexes
	fmt.Println(GetGoodsDescription(1))

	// sparse indexes, use binary search
	fmt.Println(GetGradeLevel(70))
}

var monthMap = []byte{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func GetMonth(month, year int) string {
	days := monthMap[month-1]
	if month == 2 && year&3 == 0 {
		days++
	}
	return fmt.Sprintf("%d.%d days: %d", year, month, days)
}

// 有 100 种商品，每种商品都有一个 ID 号，但很多商品的描述都差不多，所以只有 30 条不同的描述
// 这样建立索引可以节约空间，当然这并不是实际开发中一个很好的例子
var descriptions = [30]string{"good stuff", "bad stuff", "neither good nor bad stuff"}
var goodsDescMap = [100]int{3, 2, 1, 4, 3, 3, 1}       // 其余省略。。。

func GetGoodsDescription(goodsID int) string {
	if goodsID < 0 || goodsID > 100 {
		return ""
	}
	return descriptions[goodsDescMap[goodsID-1]]
}


// 成绩定级
var gradeLevel = [5]byte{'A', 'B', 'C', 'D', 'F'}
var gradeLevelStandard = [4]int{90, 80, 70, 60}

func GetGradeLevel(grade int) string {
	return string(gradeLevel[sort.Search(len(gradeLevelStandard), func(i int) bool {
		return grade >= gradeLevelStandard[i]
	})])
}
