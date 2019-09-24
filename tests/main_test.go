package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"flag"
	"os"
	"errors"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"

	"github.com/pawmart/northerntech-simpletwitter/config"
	"github.com/pawmart/northerntech-simpletwitter/internal/storage/mongo"

	o "github.com/pawmart/northerntech-simpletwitter/internal/restapi/operations"
	"github.com/pawmart/northerntech-simpletwitter/internal/models"
	"github.com/pawmart/northerntech-simpletwitter/cmd/server"
	"github.com/DATA-DOG/godog/colors"
)

type featureState struct {
	api  *o.NortherntechSimpletwitterAPI
	resp *httptest.ResponseRecorder
	handler http.Handler
}

func TestMain(m *testing.M) {
	flag.Parse()

	opt := godog.Options{}
	opt.Output = colors.Colored(os.Stdout)
	opt.Paths = []string{"../api/features"}
	opt.Format = "pretty"


	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func (a *featureState) resetState(interface{}) {
	a.resp = httptest.NewRecorder()
}

func (a *featureState) prepareSuite() {

	cfg := config.NewConfig()
	storageCnf := cfg.GetDbConfig()
	s := mongo.NewStorage(storageCnf)
	s.Drop()

	a.handler = server.GetAPIServer().GetHandler()

	recordJSON := `{
		"id" : "b8dfdf10-33fa-4301-b859-e19853641651",
		"type": "Tweets",
		"tag": "abc",
		"created_on": 1558772378,
		"modified_on": 1558772378,
		"attributes": {
			"message": "hello my friend",
			"user_id": "b8dfdf10-33fa-4301-b859-e19853641651"
		}
        }`

	recordJSON2 := `{
                "id" : "a8dfdf10-33fa-4301-b859-e19853641655",
		"type": "Tweets",
		"tag": "abc",
		"created_on": 1558772378,
		"modified_on": 1558772378,
		"attributes": {
			"message": "hello my friend v2",
			"user_id": "b8dfdf10-33fa-4301-b859-e19853641651"
		}
        }`

	p := new(models.Tweet)
	p2 := new(models.Tweet)
	json.Unmarshal([]byte(recordJSON), p)
	json.Unmarshal([]byte(recordJSON2), p2)

	if err := s.InsertTweet(p); err != nil {
		errors.New("populate db failed")
	}
	if err := s.InsertTweet(p2); err != nil {
		errors.New("populate db failed")
	}
}

func (a *featureState) callEndpoint(method string, endpoint string, req *http.Request) {

	rr := httptest.NewRecorder()
	a.handler.ServeHTTP(rr, req)
	a.resp = rr
}

func (a *featureState) iSendRequestTo(method, endpoint string) (err error) {

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	a.callEndpoint(method, endpoint, req)
	return
}

func (a *featureState) iSendARequestToWith(method, endpoint string, body *gherkin.DocString) (err error) {

	req, err := http.NewRequest(method, endpoint, strings.NewReader(body.Content))
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	a.callEndpoint(method, endpoint, req)
	return
}

func (a *featureState) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *featureState) theJSONResponseShouldContainAnError() (err error) {
	var error models.APIError
	if err = json.Unmarshal([]byte(a.resp.Body.String()), &error); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err != nil {
		return fmt.Errorf("could not unmarshal error response err %v and response %v", error, a.resp)
	}
	if len(error.ErrorMessage) > 0 && len(error.ErrorCode) > 0 {
		return
	} else {
		return errors.New("no error message")
	}
	return
}

func (a *featureState) theJSONResponseShouldContainTweetData() (err error) {
	var pd *models.TweetCreationResponse
	if err = json.Unmarshal([]byte(a.resp.Body.String()), &pd); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err != nil {
		return fmt.Errorf("could not unmarshal tweet data response %v error %v and body %v", pd, err.Error(), a.resp.Body.String())
	}

	if pd.Data != nil && len(pd.Data.ID) > 0 {
		return
	}
	return
}

func (a *featureState) theResponseShouldHaveHealthStatusUp() (err error) {
	var pd *models.Health
	if err = json.Unmarshal([]byte(a.resp.Body.String()), &pd); err != nil {
		return fmt.Errorf(err.Error())
	}

	if err != nil {
		return fmt.Errorf("could not unmarshal response %v error %v and body %v", pd, err.Error(), a.resp.Body.String())
	}
	if pd.Status != "up" {
		return fmt.Errorf("health status is not up %v error %v and body %v", pd, err.Error(), a.resp.Body.String())
	}
	return
}

func (a *featureState) theJSONResponseShouldContainTweetDataWithMessageAttributeOf(arg1 string) (err error) {
	var pd *models.TweetDetailsResponse
	if err = json.Unmarshal([]byte(a.resp.Body.String()), &pd); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err != nil {
		return fmt.Errorf("could not unmarshal response %v error %v and body %v", pd, err.Error(), a.resp.Body.String())
	}
	if *pd.Data.Attributes.Message != arg1 {
		return fmt.Errorf("wrong reference %v error %v and body %v", pd, err.Error(), a.resp.Body.String())
	}
	return
}

func (a *featureState) theJSONResponseShouldContainTweetDataCollection() (err error) {
	var pd *models.TweetDetailsListResponse
	if err = json.Unmarshal([]byte(a.resp.Body.String()), &pd); err != nil {
		return fmt.Errorf(err.Error())
	}

	if pd.Data == nil || len(pd.Data) < 1 {
		return errors.New("no collection returned")
	}
	if pd.Data[0].ID == "" {
		return errors.New("no collection record returned")
	}

	return
}

func (a *featureState) theJSONResponseShouldContainNoTweetDataCollection() (err error) {
	var pd *models.TweetDetailsListResponse
	if err = json.Unmarshal([]byte(a.resp.Body.String()), &pd); err != nil {
		fmt.Errorf(err.Error())
	}

	for _, v := range pd.Data {
		return fmt.Errorf("tweet data exists and it should not %v error %v and body %v", pd, a.resp.Body.String(), v)
	}

	return
}

func FeatureContext(s *godog.Suite) {
	api := &featureState{}

	s.BeforeSuite(api.prepareSuite)
	s.BeforeScenario(api.resetState)

	s.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
	s.Step(`^I send a "([^"]*)" request to "([^"]*)" with:$`, api.iSendARequestToWith)
	s.Step(`^the response should have health status up$`, api.theResponseShouldHaveHealthStatusUp)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the JSON response should contain an error$`, api.theJSONResponseShouldContainAnError)
	s.Step(`^the JSON response should contain tweet data$`, api.theJSONResponseShouldContainTweetData)
	s.Step(`^the JSON response should contain no tweet data collection$`, api.theJSONResponseShouldContainNoTweetDataCollection)
	s.Step(`^the JSON response should contain tweet data collection$`, api.theJSONResponseShouldContainTweetDataCollection)
	s.Step(`^the JSON response should contain tweet data with message attribute of "([^"]*)"$`, api.theJSONResponseShouldContainTweetDataWithMessageAttributeOf)
}
