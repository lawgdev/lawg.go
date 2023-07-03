package lawg

import (
	"encoding/json"
	"fmt"
	types "github.com/lawgdev/lawg.go/internal"
	utils "github.com/lawgdev/lawg.go/pkg/utils"
	"net/http"
)

type Lawg struct {
	Token   string
	Project string
	UA      string
}

type Feed struct {
	Token     string
	Project   string
	FeedName  string
	UserAgent string
}

func NewClient(token, project, ua string) *Lawg {
	return &Lawg{
		Token:   token,
		Project: project,
		UA:      ua,
	}
}

func (l *Lawg) Feed(feedName string) *Feed {
	return &Feed{
		Token:     l.Token,
		Project:   l.Project,
		FeedName:  feedName,
		UserAgent: l.UA,
	}
}

func (l *Lawg) Insight(options types.CreateInsight) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/insights", l.Project)
	data, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	reqOptions := utils.Options{
		Method: "POST",
		Token:  l.Token,
		Data:   data,
	}

	return utils.Request(url, reqOptions)
}

func (l *Lawg) SetInsight(options types.UpdateInsight) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/insights/%s", l.Project, options.ID)
	data := map[string]interface{}{
		"value": map[string]interface{}{
			"set": options.Set,
		},
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	reqOptions := utils.Options{
		Method: "PATCH",
		Token:  l.Token,
		Data:   jsonData,
	}

	return utils.Request(url, reqOptions)
}

func (l *Lawg) IncInsight(options types.UpdateInsight) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/insights/%s", l.Project, options.ID)
	data := map[string]interface{}{
		"value": map[string]interface{}{
			"increment": options.Increment,
		},
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	reqOptions := utils.Options{
		Method: "PATCH",
		Token:  l.Token,
		Data:   jsonData,
	}

	return utils.Request(url, reqOptions)
}

func (l *Lawg) DeleteInsight(id string) (*http.Response, error) {
	url := fmt.Sprintf("projects/%s/insights/%s", l.Project, id)

	reqOptions := utils.Options{
		Method: "DELETE",
		Token:  l.Token,
	}

	return utils.Request(url, reqOptions)
}
