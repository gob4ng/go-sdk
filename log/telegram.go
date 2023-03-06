package log

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func (t TelegramBot) HitTelegram(message string) (*string, *error) {

	telegramUrl, err := url.Parse(t.TelegramUrl)
	if err != nil {
		return nil, &err
	}

	rawQuery := telegramUrl.Query()
	rawQuery.Set("chat_id", t.GroupChatID)
	rawQuery.Set("text", message)
	telegramUrl.RawQuery = rawQuery.Encode()

	request, err := http.NewRequest(http.MethodGet, telegramUrl.String(), nil)
	if err != nil {
		return nil, &err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, &err
	}

	byteResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &err
	}

	strResponse := string(byteResult)

	return &strResponse, nil
}
