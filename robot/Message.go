package robot

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	RCtemplate "robot-client/template"
	"strings"
)

type data string

func (d *data) Write(p []byte) (n int, err error) {
	var s1 = string(p)
	d2 := data(s1)
	*d = *d + d2
	return 0, nil
}

type PushPlus struct {
	Token       string `json:"token"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Template    string `json:"template"`
	Channel     string `json:"channel"`
	Webhook     string `json:"webhook"`
	CallbackUrl string `json:"callbackUrl"`
	Timestamp   string `json:"timestamp"`
}

func PushPlusPost(message XCAutoLog, token string) {
	var d data = ""
	_template, _ := template.New("CHD").Parse(RCtemplate.PushPlusHTML)
	err := _template.ExecuteTemplate(&d, "CHD", message)
	if err != nil {
		return
	}
	body := PushPlus{Token: token, Title: message.Name, Content: string(d), Template: "html", Channel: "wechat"}
	marshal, _ := json.Marshal(body)
	request, _ := http.NewRequest("POST", "http://www.pushplus.plus/send", strings.NewReader(string(marshal)))
	request.Header.Add("Content-Type", "application/json")
	do, _ := http.DefaultClient.Do(request)
	log.Println(do.Body)
}
