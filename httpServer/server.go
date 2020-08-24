package httpServer

import (
	"encoding/json"
	"github.com/Anveena/RoomOfRequirement/ezLog"
	"github.com/lezhigb/KanBanMSG/reqModel"
	"github.com/lezhigb/KanBanMSG/wekanConf"
	"io/ioutil"
	"net/http"
	"regexp"
)

type preHandleModel struct {
	handlerFun func(paras []string) error
	reg        *regexp.Regexp
}

func StartHTTPServer() error {
	routeDic := map[string]preHandleModel{
		"act-createCard": {onCreateCard, regexp.MustCompile(`([^ ]*) *created card "([^ ]*)" to list "([^ ]*)" at swimlane "([^ ]*)" at board "([^ ]*)"\n(http://.+)`)},
	}
	http.HandleFunc(wekanConf.Config().WeKanSetting.CallbackPath, func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("X-Wekan-Token") != wekanConf.Config().WeKanSetting.Token {
			http.Error(writer, "token is not valid!", http.StatusForbidden)
			return
		}
		body, err := ioutil.ReadAll(request.Body)
		if err != nil || len(body) < 1 {
			http.Error(writer, "what a http body!", http.StatusBadRequest)
			return
		}
		var m reqModel.WeKanCallbackModel
		err = json.Unmarshal(body, &m)
		if err != nil || len(body) < 1 {
			http.Error(writer, "what a http body!", http.StatusBadRequest)
			return
		}
		handlerModel, valid := routeDic[m.Description]
		if !valid {
			ezLog.SendMessageToDing("unhandled method:", m.Description)
		} else {
			if err := handlerModel.handlerFun(handlerModel.reg.FindStringSubmatch(m.Text)); err != nil {
				ezLog.SendMessageToDing("unhandled method:", m.Description)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}
		writer.WriteHeader(http.StatusOK)
	})
	return http.ListenAndServe(":"+wekanConf.Config().WeKanSetting.ListenPort, nil)
}
