package wekanConf

import (
	"encoding/base64"
	"errors"
	"github.com/Anveena/RoomOfRequirement/ezCrypto"
	"github.com/Anveena/RoomOfRequirement/ezLog"
	"sync"
)

type conf struct {
	LogConf      ezLog.EZLoggerModel `json:"log_conf"`
	WeKanSetting struct {
		Token              string
		TokenEncoded       string   `json:"token_encoded"`
		CallbackPath       string   `json:"callback_path"`
		ListenPort         string   `json:"listen_port"`
		OptHistoryFilePath string   `json:"opt_history_file_path"`
		OptDingPath        string   `json:"opt_ding_path"`
		BoardsOfInterest   []string `json:"boards_of_interest"`
		BOIDic             map[string]bool
	} `json:"we_kan_setting"`
}

var instance *conf
var once sync.Once

func Config() *conf {
	once.Do(func() {
		instance = &conf{}
	})
	return instance
}
func (c *conf) Check() error {
	if c.WeKanSetting.TokenEncoded == "" {
		return errors.New("empty WeKan token")
	}
	token, e := GetPasswordFromBase64Str(c.WeKanSetting.TokenEncoded)
	if e != nil {
		return e
	}
	c.WeKanSetting.Token = token
	c.WeKanSetting.BOIDic = make(map[string]bool, len(c.WeKanSetting.BoardsOfInterest))
	for _, s := range c.WeKanSetting.BoardsOfInterest {
		c.WeKanSetting.BOIDic[s] = true
	}
	return nil
}

func MakePasswordBase64Str(origPwd string) (string, error) {
	origPwdData := []byte(origPwd)
	encData, err := ezCrypto.EZEncrypt(&origPwdData, "this code may be not working", 9499)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(*encData), nil
}
func GetPasswordFromBase64Str(base64Str string) (string, error) {
	encData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	origData, err := ezCrypto.EZDecrypt(&encData, "this code may be not working", 9499)
	if err != nil {
		return "", err
	}
	return string(*origData), nil
}
