package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strings"
)

// go 网络编程
// https://studygolang.com/articles/9240

type DNSHeader struct {
	ID            uint16
	Flag          uint16
	QuestionCount uint16
	AnswerRRs     uint16 //RRs is Resource Records
	AuthorityRRs  uint16
	AdditionalRRs uint16
}

type DNSQuery struct {
	QuestionType  uint16
	QuestionClass uint16
}

func main() {
	var (
		dnsHeader   DNSHeader
		dnsQuestion DNSQuery
	)

	//填充dns首部
	dnsHeader.ID = 0xFFFF
	dnsHeader.SetFlag(0, 0, 0, 0, 1, 0, 0)
	dnsHeader.QuestionCount = 1
	dnsHeader.AnswerRRs = 0
	dnsHeader.AuthorityRRs = 0
	dnsHeader.AdditionalRRs = 0

	//填充dns查询首部
	dnsQuestion.QuestionType = 1 //IPv4
	dnsQuestion.QuestionClass = 1

	var (
		conn net.Conn
		err  error

		buffer bytes.Buffer
	)

	//DNS是基于UDP协议的
	if conn, err = net.Dial("udp", "127.0.0.53:53"); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	//buffer中是我们要发送的数据，里面的内容是DNS首部+查询内容+DNS查询首部
	_ = binary.Write(&buffer, binary.BigEndian, dnsHeader)
	_ = binary.Write(&buffer, binary.BigEndian, ParseDomainName("www.baidu.com"))
	_ = binary.Write(&buffer, binary.BigEndian, dnsQuestion)
	fmt.Println(string(buffer.Bytes()))

	// 写入请求
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("send success.")

	// 读取响应
	//time.Sleep(2 * time.Second)
	//var resp []byte
	//if _, err := conn.Read(resp); err != nil {
	//	fmt.Println(string(resp))
	//}
	reader := bufio.NewReader(conn)
	for {
		// ReadString 会一直阻塞直到遇到分隔符 '\n'
		// 遇到分隔符后 ReadString 会返回上次遇到分隔符到现在收到的所有数据
		// 若在遇到分隔符之前发生异常, ReadString 会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			// 通常遇到的错误是连接中断或被关闭，用io.EOF表示
			if err == io.EOF {
				fmt.Println(err)
			} else {
				fmt.Println(err)
			}
			return
		}
		fmt.Println(msg)
	}
}

func (header *DNSHeader) SetFlag(QR uint16, OperationCode uint16, AuthoritativeAnswer uint16, Truncation uint16, RecursionDesired uint16, RecursionAvailable uint16, ResponseCode uint16) {
	header.Flag = QR<<15 + OperationCode<<11 + AuthoritativeAnswer<<10 + Truncation<<9 + RecursionDesired<<8 + RecursionAvailable<<7 + ResponseCode
}

func ParseDomainName(domain string) []byte {
	//要将域名解析成相应的格式，例如：
	//"www.google.com"会被解析成"0x03www0x06google0x03com0x00"
	//就是长度+内容，长度+内容……最后以0x00结尾
	var (
		buffer   bytes.Buffer
		segments []string = strings.Split(domain, ".")
	)
	for _, seg := range segments {
		_ = binary.Write(&buffer, binary.BigEndian, byte(len(seg)))
		_ = binary.Write(&buffer, binary.BigEndian, []byte(seg))
	}
	_ = binary.Write(&buffer, binary.BigEndian, byte(0x00))

	return buffer.Bytes()
}
