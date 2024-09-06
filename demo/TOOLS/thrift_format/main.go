package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var helpHint = `Name:
   thriftformat
Usage:
   thriftformat [options] [argument]
options are:
   -f  filePath    string类型。thrift文件路径。支持相对路径和绝对路径。
   -ow overwrite   bool类型。是否直接覆盖源文件。若选择否，则仅输出格式化的内容，不覆盖源文件。(default:true)
   -h  help        bool类型。Display usage information (this message)。
example:
   thriftformat -f=your_file.thrift -ow=false
`

var (
	filePath  string
	overwrite bool
	showHelp  bool
)

func main() {
	initParam()
	process()
}

// StructLine 一行。（结构体类型）
type StructLine struct {
	number          string // 字段序号
	paramNamePrefix string // 字段序号之后、参数之前
	paramName       string // 字段参数名
	defaultValue    string // 字段参数默认值
	description     string // 字段注解
	annotation      string // 字段注释
}

// EnumLine 一行。（Enum类型）
type EnumLine struct {
	number         string // 字段序号
	paramName      string // 字段参数名
	hasDescription bool   // 是否有注解
	description    string // 字段注解
	annotation     string // 字段注释
}

func initParam() {
	flag.BoolVar(&overwrite, "ow", true, "")
	flag.StringVar(&filePath, "f", "", "")
	flag.BoolVar(&showHelp, "h", false, "")
	flag.Parse()

	fmt.Println(overwrite)
}

func process() {

	filePath = "～/webhook.thrift"

	if showHelp {
		fmt.Println(helpHint)
		return
	}

	if filePath == "" {
		fmt.Println(helpHint)
		return
	}

	path, _ := os.Getwd()
	// 默认是相对路径，如果以"/"开头，那么是绝对路径。
	if !strings.HasPrefix(filePath, "/") {
		filePath = path + "/" + filePath
	}

	newFileLines := []string{}  // 新的文件的每一行
	structLines := []string{}   // 一个struct结构体的每一行
	enumLines := []string{}     // 一个enum结构体的每一行
	inStructProcessing := false // 当前是否正在处理struct结构体
	inEnumrocessing := false    // 当前是否正在处理enum结构体
	maxLength1 := 0
	maxLength2 := 0
	maxLength3 := 0
	maxLength4 := 0
	maxLength5 := 0
	maxLength6 := 0

	fmt.Println("正在解析文件...")
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件失败!", err)
		return
	}
	buf := bufio.NewReader(file)
	for {
		lineBytes, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		line := string(lineBytes)

		// 开始遍历文件
		if strings.HasPrefix(strings.TrimSpace(line), "struct") {
			inStructProcessing = true
		} else if strings.HasPrefix(line, "enum") {
			inEnumrocessing = true
		}

		if inStructProcessing {
			// 处理结构体
			processStruct(line, &structLines, &newFileLines, &maxLength1, &maxLength2, &maxLength3, &maxLength4, &inStructProcessing)
		} else if inEnumrocessing {
			// 处理枚举
			processEnum(line, &enumLines, &newFileLines, &maxLength5, &maxLength6, &inEnumrocessing)
		} else {
			// 其他行，原封不动
			newFileLines = append(newFileLines, line)
		}
	}
	file.Close()

	// 输出
	for i := range newFileLines {
		fmt.Print(newFileLines[i] + "\n")
	}

	// 写回源文件
	if overwrite {
		fmt.Println("解析完成，覆盖写回原文件...")
		file2, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println("打开文件失败", err)
			return
		}
		writer := bufio.NewWriter(file2)
		for i := range newFileLines {
			writer.WriteString(newFileLines[i] + "\n")
		}
		writer.Flush()
		file2.Close()
	}

	fmt.Println("成功")
}

// 遍历message结构体，找出最长的行，以此作为标杆，其他的行都要补齐。
func processStruct(line string, structLines *[]string, newFileLines *[]string, maxLength1 *int, maxLength2 *int, maxLength3 *int, maxLength4 *int, inStructProcessing *bool) {

	// 遍历message结构体，找出最长的行，以此作为标杆，其他的行都要补齐。
	if strings.Contains(line, ":") {
		lineInfo := parseStructOneLine(line)

		if len(lineInfo.number) > *maxLength1 {
			*maxLength1 = len(lineInfo.number)
		}

		if len(lineInfo.paramNamePrefix) > *maxLength2 {
			*maxLength2 = len(lineInfo.paramNamePrefix)
		}

		totalLength := len(lineInfo.paramName)
		if lineInfo.defaultValue != "" {
			totalLength += len(lineInfo.defaultValue) + 3
		}
		if totalLength > *maxLength3 {
			*maxLength3 = totalLength
		}

		if len(lineInfo.description) > *maxLength4 {
			*maxLength4 = len(lineInfo.description)
		}

	}
	*structLines = append(*structLines, line)

	// 结构体结束
	if strings.HasPrefix(line, "}") {
		for _, line := range *structLines {

			// 结构体中的不包含 “:” 的行，直接原封不动不格式化
			if !strings.Contains(line, ":") {
				*newFileLines = append(*newFileLines, line)
				continue
			}

			// 解析本行
			lineInfo := parseStructOneLine(line)

			// 计算当前行与本结构体最长行的长度差
			delta1 := *maxLength1 - len(lineInfo.number)
			delta2 := *maxLength2 - len(lineInfo.paramNamePrefix)
			delta3 := *maxLength3 - len(lineInfo.paramName) - len(lineInfo.defaultValue)
			delta4 := *maxLength4 - len(lineInfo.description)

			// 格式化本行
			formatedLine := formatStructOneLine(lineInfo, delta1, delta2, delta3, delta4)

			*newFileLines = append(*newFileLines, formatedLine)
		}

		// 已经处理完一个结构体。恢复0值
		*inStructProcessing = false
		*structLines = []string{}
		*maxLength1 = 0
		*maxLength2 = 0
		*maxLength3 = 0
		*maxLength4 = 0
	}
}

func processEnum(line string, enumLines *[]string, newFileLines *[]string, maxLength5 *int, maxLength6 *int, inEnumProcessing *bool) {
	// 遍历enum结构体，找出最长的行，以此作为标杆，其他的行都要补齐。
	if strings.Contains(line, "=") {
		lineInfo := parseEnumOneLine(line)
		if len(lineInfo.paramName) > *maxLength5 {
			*maxLength5 = len(lineInfo.paramName)
		}
		if len(lineInfo.number) > *maxLength6 {
			*maxLength6 = len(lineInfo.number)
		}
	}
	*enumLines = append(*enumLines, line)

	// 结构体结束
	if strings.HasPrefix(line, "}") {
		for _, line := range *enumLines {
			// 结构体中的不包含 “=” 的行，不做改动
			if !strings.Contains(line, "=") {
				*newFileLines = append(*newFileLines, line)
				continue
			}

			// 解析本行
			lineInfo := parseEnumOneLine(line)

			// 计算当前行与本结构体最长行的长度差
			delta3 := *maxLength5 - len(lineInfo.paramName)
			delta4 := *maxLength6 - len(lineInfo.number)

			// 格式化本行
			formatedLine := formatEnumOneLine(lineInfo, delta3, delta4)

			*newFileLines = append(*newFileLines, formatedLine)
		}
		// 已经处理完一个结构体。恢复0值
		*inEnumProcessing = false
		*enumLines = []string{}
		*maxLength5 = 0
		*maxLength6 = 0
	}
}

// 解析结构体一行
func parseStructOneLine(line string) *StructLine {
	lineInfo := &StructLine{}

	splitOnColon := strings.Split(line, ":")
	lineInfo.number = strings.TrimLeft(splitOnColon[0], " ")

	temp := strings.Join(splitOnColon[1:], ":") // 为了避免本行有其他的冒号，也被按照冒号分割了， 所以把其他的冒号再join到一起

	splitForAnnotation := strings.Split(temp, "//") // 判断注释
	if len(splitForAnnotation) >= 2 {
		annotation := strings.Join(splitForAnnotation[1:], "//")
		lineInfo.annotation = annotation
	} else {
		lineInfo.annotation = ""
	}

	splitForDescription := strings.Split(splitForAnnotation[0], "(") // 判断注解
	if len(splitForDescription) >= 2 {
		tempSplit := strings.Split(splitForDescription[1], ")")
		lineInfo.description = tempSplit[0]
	} else {
		lineInfo.description = ""
	}

	splitForDefaultVal := strings.Split(splitForDescription[0], "=")
	if len(splitForDefaultVal) >= 2 { // 判断默认值
		equalStr := strings.Join(splitForDefaultVal[1:], "=")
		lineInfo.defaultValue = strings.TrimSpace(equalStr)
	} else {
		lineInfo.defaultValue = ""
	}

	middleStrNoSpace := deleteExtraSpace(strings.Trim(splitForDefaultVal[0], " "))
	middleStrNoSpace = strings.TrimRight(middleStrNoSpace, " ")

	splitForParamName := strings.Split(middleStrNoSpace, " ")
	paramNamePrefix := ""
	for i := 0; i < len(splitForParamName)-1; i++ {
		paramNamePrefix += splitForParamName[i] + " "
	}
	lineInfo.paramNamePrefix = strings.TrimRight(paramNamePrefix, " ")
	lineInfo.paramName = splitForParamName[len(splitForParamName)-1]

	return lineInfo
}

// 解析Enum一行
func parseEnumOneLine(line string) *EnumLine {
	enumLine := &EnumLine{}

	arr1 := strings.Split(line, "=")
	enumLine.paramName = strings.TrimSpace(arr1[0])

	arr2 := strings.Split(arr1[1], "//")
	// 有注释
	if len(arr2) == 2 {
		enumLine.annotation = strings.TrimLeft(arr2[1], " ")
	}
	enumLine.number = strings.TrimSpace(arr2[0])
	return enumLine
}

// 格式化一行。（结构体）
func formatStructOneLine(lineInfo *StructLine, delta1, delta2, delta3, delta4 int) string {
	var buffer bytes.Buffer
	buffer.WriteString("    ")
	buffer.WriteString(lineInfo.number)
	buffer.WriteString(": ")
	for i := 0; i < delta1; i++ {
		buffer.WriteString(" ")
	}

	buffer.WriteString(lineInfo.paramNamePrefix)

	for i := 0; i < delta2; i++ {
		buffer.WriteString(" ")
	}

	buffer.WriteString(" ")
	buffer.WriteString(lineInfo.paramName)

	if lineInfo.defaultValue != "" {
		buffer.WriteString(" = ")
		buffer.WriteString(lineInfo.defaultValue)
	}
	if lineInfo.description != "" || lineInfo.annotation != "" {
		if lineInfo.defaultValue != "" {
			delta3 -= 3
		}
		for i := 0; i < delta3; i++ {
			buffer.WriteString(" ")
		}
	}

	if lineInfo.description != "" {
		buffer.WriteString(" (")
		buffer.WriteString(lineInfo.description)
		buffer.WriteString(")")
	}
	if lineInfo.annotation != "" {
		for i := 0; i < delta4; i++ {
			buffer.WriteString(" ")
		}
		if lineInfo.description == "" {
			for i := 0; i < 3; i++ {
				buffer.WriteString(" ")
			}
		}
		buffer.WriteString(" //")
		buffer.WriteString(lineInfo.annotation)
	}

	return buffer.String()
}

func formatEnumOneLine(lineInfo *EnumLine, delta3 int, delta4 int) string {
	var buffer bytes.Buffer
	buffer.WriteString("    ")

	buffer.WriteString(lineInfo.paramName)

	for i := 0; i < delta3; i++ {
		buffer.WriteString(" ")
	}

	buffer.WriteString(" = ")
	buffer.WriteString(lineInfo.number)

	for i := 0; i < delta4; i++ {
		buffer.WriteString(" ")
	}

	if lineInfo.annotation != "" {
		buffer.WriteString("    // ")
	}
	buffer.WriteString(lineInfo.annotation)
	return buffer.String()
}

//删除字符串中的多余空格，有多个空格时，仅保留一个空格
func deleteExtraSpace(s string) string {
	s1 := strings.Replace(s, " ", " ", -1)      // 替换tab为空格
	regstr := "\\s{2,}"                         // 两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)            //编译正则表达式
	s2 := make([]byte, len(s1))                 //定义字符数组切片
	copy(s2, s1)                                //将字符串复制到切片
	spcIndex := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spcIndex) > 0 {
		s2 = append(s2[:spcIndex[0]+1], s2[spcIndex[1]:]...) //删除多余空格
		spcIndex = reg.FindStringIndex(string(s2))           //继续在字符串中搜索
	}
	return string(s2)
}
