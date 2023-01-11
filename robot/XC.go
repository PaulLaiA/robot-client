package robot

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type XCAutoLog struct {
	Name        string            //名称
	Time        int64             //10000000 = 1s
	Revive      string            //死亡次数
	Msg         string            //消息
	Acquisition map[string]string //获得品
	Consumables map[string]string //消耗品
	Card        []string          //翻牌
	Book        []string          //图鉴
}

func Parse(path string) XCAutoLog {
	log.Println("parse xc log file {", paths, "}")
	content, _ := os.ReadFile(path)
	utf8, _ := GbkToUtf8(content)
	str := string(utf8)
	data := ParseContent(str)
	log.Println(data)
	return data
}

func ParseContent(content string) XCAutoLog {
	list := regexp.MustCompile(`[\t\n\f\r]`).Split(content, -1)
	var info, revive, acquisition, consumables, msg, card, book int
	for i, s := range list {
		switch s {
		case "[INFO]":
			info = i
		case "[REVIVE]":
			revive = i
		case "[ITEM1]":
			acquisition = i
		case "[ITEM2]":
			consumables = i
		case "[MSG]":
			msg = i
		case "[CARD]":
			card = i
		case "[BOOK]":
			book = i
		}
	}
	autoLog := XCAutoLog{
		Name:        "",
		Time:        0,
		Revive:      "",
		Msg:         "",
		Acquisition: make(map[string]string),
		Consumables: make(map[string]string),
		Card:        []string{},
		Book:        []string{},
	}

	temp := list[info:revive]

	var tTime int64
	for _, s := range temp {
		if strings.HasPrefix(s, "NAME=") {
			autoLog.Name = strings.Split(s, "=")[1]
		} else if strings.HasPrefix(s, "REVIVE=") {
			autoLog.Revive = strings.Split(s, "=")[1]
		} else if strings.HasPrefix(s, "MSG=") {
			autoLog.Msg = strings.Split(s, "=")[1]
		} else if strings.HasPrefix(s, "BEGIN=") {
			i, _ := strconv.ParseInt(strings.Split(s, "BEGIN=")[1], 10, 0)
			tTime = tTime - i
		} else if strings.HasPrefix(s, "END=") {
			i, _ := strconv.ParseInt(strings.Split(s, "END=")[1], 10, 0)
			tTime = tTime + i
		}
	}
	var a int64 = 10000000
	autoLog.Time = tTime / a / 60

	temp = list[acquisition:consumables]
	for _, s := range temp {
		if ts := strings.Split(s, "="); strings.ContainsAny(s, "=") && len(ts) > 1 {
			autoLog.Acquisition[ts[0]] = ts[1]
		}
	}

	temp = list[consumables:msg]
	for _, s := range temp {
		if ts := strings.Split(s, "="); strings.ContainsAny(s, "=") && len(ts) > 1 {
			autoLog.Consumables[ts[0]] = ts[1]
		}
	}

	temp = list[card:book]
	copy(autoLog.Card, temp)

	temp = list[book:]
	copy(autoLog.Book, temp)

	return autoLog
}

// GbkToUtf8 GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
