package lawg

import (
	"encoding/json"
	"fmt"
	"net/http"

	types "github.com/lawgdev/lawg.go/internal"
	utils "github.com/lawgdev/lawg.go/pkg/utils"
)

type Feed struct {
	Token    string
	Project  string
	UA       string
	FeedName string
}

func (f *Feed) FetchEvents() (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/feeds/%s/events", f.Project, f.FeedName)

	reqOptions := utils.Options{
		Method: "GET",
		Token:  f.Token,
	}

	return utils.Request(url, reqOptions)
}

func (f *Feed) Event(options types.CreateEvent) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/feeds/%s/events", f.Project, f.FeedName)

	data, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	reqOptions := utils.Options{
		Method: "POST",
		Token:  f.Token,
		Data:   json.RawMessage(data),
	}

	return utils.Request(url, reqOptions)
}

func (f *Feed) FetchEvent(id string) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/feeds/%s/events/%s", f.Project, f.FeedName, id)

	reqOptions := utils.Options{
		Method: "GET",
		Token:  f.Token,
	}

	return utils.Request(url, reqOptions)
}

func (f *Feed) EditEvent(options types.UpdateEvent) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/feeds/%s/events/%s", f.Project, f.FeedName, *options.ID)

	data, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	reqOptions := utils.Options{
		Method: "PATCH",
		Token:  f.Token,
		Data:   json.RawMessage(data),
	}

	return utils.Request(url, reqOptions)
}

func (f *Feed) DeleteEvent(id string) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/feeds/%s/events/%s", f.Project, f.FeedName, id)

	reqOptions := utils.Options{
		Method: "DELETE",
		Token:  f.Token,
	}

	return utils.Request(url, reqOptions)
}
