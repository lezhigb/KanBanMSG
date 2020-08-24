package dingCaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Anveena/RoomOfRequirement/ezLog"
	"github.com/lezhigb/KanBanMSG/wekanConf"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type actionCard struct {
	Title          string `json:"Title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []struct {
		Title     string `json:"title"`
		ActionURL string `json:"actionURL"`
	} `json:"btns"`
}
type dingtalkCallbackModel struct {
	MsgType    string     `json:"msgtype"`
	ActionCard actionCard `json:"actionCard"`
	At         struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"is_at_all"`
	} `json:"at"`
}

func SendDingMsgWithAt(title string, desc string, link string, at []string) {

}
func SendDingMsg(title string, desc string, link string) {
	var m = dingtalkCallbackModel{
		MsgType: "actionCard",
		ActionCard: actionCard{Title: title,
			Text:           fmt.Sprintf("### %s\n> ##### %s\n> 操作时间:%s", title, desc, time.Now().Format("2006-01-02 15:04:05")),
			BtnOrientation: "0",
			Btns: []struct {
				Title     string `json:"title"`
				ActionURL string `json:"actionURL"`
			}{
				{Title: "立刻查看", ActionURL: link},
			}},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"is_at_all"`
		}{
			nil, false,
		},
	}
	tmpData, err := json.Marshal(m)

	if err != nil {
		ezLog.F(err.Error())
		return
	}
	rsp, err := http.Post(wekanConf.Config().WeKanSetting.OptDingPath, "application/json;charset=utf-8", bytes.NewBuffer(tmpData))
	if err != nil {
		ezLog.F(err.Error())
		return
	}
	rspInfo, _ := ioutil.ReadAll(rsp.Body)
	rspStr := string(rspInfo)
	if !strings.Contains(rspStr, `"errcode":0,`) {
		ezLog.F(rspStr)
		return
	}
	_ = rsp.Body.Close()
}
