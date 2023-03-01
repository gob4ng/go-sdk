package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gob4ng/go-sdk/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (h Context) HitClient() (*Response, *error) {

	if h.ZapLog == nil {
		newErr := errors.New("please define zap log")
		return nil, &newErr
	}

	// set http new request
	request, errHttpNewRequest := setHttpNewRequest(h)
	if errHttpNewRequest != nil {
		return nil, errHttpNewRequest
	}

	// hit http client
	response, err := h.OptionalContext.HttpClient.Do(request)
	if err != nil {
		return nil, &err
	}

	// read http client message
	byteResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &err
	}

	clientResponse := Response{
		HttpCode:     response.StatusCode,
		ResponseBody: string(byteResult),
	}

	printHttpMessage(h, clientResponse)

	return &clientResponse, nil

}

func setHttpNewRequest(h Context) (*http.Request, *error) {

	// set byte request
	byteRequest, err := setByteRequest(h)
	if err != nil {
		return nil, err
	}

	// set nil if byteRequest nil
	if byteRequest == nil {
		request, errNewRequest := http.NewRequest(h.HttpMethod, h.URL, nil)

		// set query param if any
		if h.OptionalContext.QueryParam != nil {
			queryParam := request.URL.Query()
			for key, value := range *h.OptionalContext.QueryParam {
				queryParam.Add(key, value)
			}
			request.URL.RawQuery = queryParam.Encode()
		}

		if errNewRequest != nil {
			return nil, &errNewRequest
		}
		return request, nil
	}

	// set http new header
	request, errNewRequest := http.NewRequest(h.HttpMethod, h.URL, bytes.NewReader(*byteRequest))

	// set query param if any
	if h.OptionalContext.QueryParam != nil {
		queryParam := request.URL.Query()
		for key, value := range *h.OptionalContext.QueryParam {
			queryParam.Add(key, value)
		}
		request.URL.RawQuery = queryParam.Encode()
	}

	if errNewRequest != nil {
		return nil, &errNewRequest
	}

	// set header
	for key, value := range h.Header {
		// check if content type is multipart form data
		if key == gin.MIMEMultipartPOSTForm {
			writer, err := setMultiPartFormData(h)
			if err != nil {
				return nil, err
			}
			request.Header.Set("Content-Type", writer.FormDataContentType())
		}
		request.Header.Set(key, value)
	}

	// set base auth if any
	if h.OptionalContext.BaseAuth != nil {
		for key, value := range *h.OptionalContext.BaseAuth {
			request.SetBasicAuth(key, value)
		}
	}

	return request, nil

}

func setByteRequest(h Context) (*[]byte, *error) {

	contentType := h.Header["Content-Type"]
	requestBody := h.OptionalContext.RequestBody

	if requestBody != nil {
		if contentType == binding.MIMEJSON {
			byteRequest, err := json.Marshal(requestBody)
			if err != nil {
				return nil, &err
			}
			return &byteRequest, nil
		}

		if contentType == binding.MIMEMultipartPOSTForm {
			// if form data is not define
			if h.OptionalContext.FormData == nil {
				newError := errors.New("please define form data request")
				return nil, &newError
			}
		}

		if contentType == binding.MIMEXML || contentType == binding.MIMEXML2 {
			byteRequest, err := xml.Marshal(requestBody)
			if err != nil {
				return nil, &err
			}
			return &byteRequest, nil
		}

		if contentType == binding.MIMEPOSTForm {
			byteRequest, err := json.Marshal(requestBody)
			if err != nil {
				return nil, &err
			}

			mapJson := map[string]string{}
			if err := json.Unmarshal(byteRequest, &mapJson); err != nil {
				return nil, &err
			}

			data := url.Values{}
			for key, value := range mapJson {
				data.Set(key, value)
			}

			bytePostForm := []byte(data.Encode())

			return &bytePostForm, nil
		}

	}

	return nil, nil
}

func setMultiPartFormData(h Context) (*multipart.Writer, *error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range *h.OptionalContext.FormData {
		for a, b := range value {
			lKey := strings.ToLower(key)

			if lKey == "file" {
				if err := setFileWriter(a, b, *writer); err != nil {
					return nil, err
				}
			} else if lKey == "text" {
				if err := setFieldWriter(a, b, *writer); err != nil {
					return nil, err
				}
			} else {
				newError := errors.New("key must be file or text")
				return nil, &newError
			}

		}
	}

	return writer, nil
}

func setFileWriter(key string, value string, writer multipart.Writer) *error {

	fw, err := writer.CreateFormFile(key, value)
	if err != nil {
		return &err
	}

	file, err := os.Open(value)
	if err != nil {
		return &err
	}

	if _, err := io.Copy(fw, file); err != nil {
		return &err
	}

	return nil
}

func setFieldWriter(key string, value string, writer multipart.Writer) *error {

	fw, err := writer.CreateFormField(key)
	if err != nil {
		return &err
	}

	if _, err = io.Copy(fw, strings.NewReader(value)); err != nil {
		return &err
	}

	return nil
}

func printHttpMessage(h Context, clientResponse Response) {

	if h.Debug {
		if h.OptionalContext.HiddenLogContext == nil {

			trackingContext := log.ZapTrackingContext{
				GinContext:     nil,
				ClientResponse: clientResponse,
				ClientContext:  h,
				LogID:          h.OptionalContext.LogID,
				UnixTimestamp:  h.OptionalContext.UnixTimestamp,
			}

			if clientResponse.HttpCode == http.StatusOK {
				h.ZapLog.ClientDebug(trackingContext)
			} else {
				h.ZapLog.ClientError(trackingContext)
			}

		} else {
			optionalPrintHttpMessage(h, clientResponse)
		}

	}

}

func optionalPrintHttpMessage(h Context, clientResponse Response) {

}
