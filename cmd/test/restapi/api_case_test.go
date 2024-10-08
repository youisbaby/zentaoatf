package main

import (
	"fmt"
	"testing"

	commonTestHelper "github.com/easysoft/zentaoatf/cmd/test/helper/common"
	constTestHelper "github.com/easysoft/zentaoatf/cmd/test/helper/conf"
	httpHelper "github.com/easysoft/zentaoatf/cmd/test/helper/http"
	"github.com/easysoft/zentaoatf/cmd/test/restapi/config"
	commConsts "github.com/easysoft/zentaoatf/internal/pkg/consts"
	commDomain "github.com/easysoft/zentaoatf/internal/pkg/domain"
	zentaoHelper "github.com/easysoft/zentaoatf/internal/pkg/helper/zentao"
	stringUtils "github.com/easysoft/zentaoatf/pkg/lib/string"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/tidwall/gjson"
)

func TestCaseApi(t *testing.T) {
	suite.RunSuite(t, new(CaseApiSuite))
}

type CaseApiSuite struct {
	suite.Suite
}

func (s *CaseApiSuite) BeforeEach(t provider.T) {
	commonTestHelper.ReplaceLabel(t, "CaseApi")
}

func (s *CaseApiSuite) TestCaseListApi(t provider.T) {
	t.ID("7612")
	token := httpHelper.Login()

	params := map[string]interface{}{
		"limit": 10,
	}
	url := zentaoHelper.GenApiUrl(fmt.Sprintf("/products/%d/testcases", config.ProductId), params, constTestHelper.ZentaoSiteUrl)

	bodyBytes, _ := httpHelper.Get(url, token)

	firstCaseId := gjson.Get(string(bodyBytes), "testcases.0.id").Int()

	t.Require().Greater(firstCaseId, int64(0), "list testcases failed, url: "+url)
}

func (s *CaseApiSuite) TestCaseListByModuleApi(t provider.T) {
	t.ID("7635")
	token := httpHelper.Login()

	moduleId := getModuleMinId()

	url := zentaoHelper.GenApiUrl(fmt.Sprintf("/products/%d/testcases?module=%d", config.ProductId, moduleId),
		nil, constTestHelper.ZentaoSiteUrl)

	bodyBytes, _ := httpHelper.Get(url, token)

	firstCaseId := gjson.Get(string(bodyBytes), "testcases.0.id").Int()

	t.Require().Greater(firstCaseId, int64(0), "list testcases failed, url: "+url)
}

func (s *CaseApiSuite) TestCaseListBySuiteApi(t provider.T) {
	t.ID("7614")
	token := httpHelper.Login()

	url := zentaoHelper.GenApiUrl(fmt.Sprintf("/testsuites/%d", config.SuiteId), nil, constTestHelper.ZentaoSiteUrl)

	bodyBytes, _ := httpHelper.Get(url, token)

	firstCaseId := gjson.Get(string(bodyBytes), "testcases.0.id").Int()

	t.Require().Greater(firstCaseId, int64(0), "list testcases failed, url: "+url)
}

func (s *CaseApiSuite) TestCaseListByTaskApi(t provider.T) {
	t.ID("7615")
	token := httpHelper.Login()

	url := zentaoHelper.GenApiUrl(fmt.Sprintf("/testtasks/%d", config.TaskId), nil, constTestHelper.ZentaoSiteUrl)

	bodyBytes, _ := httpHelper.Get(url, token)

	firstCaseId := gjson.Get(string(bodyBytes), "testcases.0.id").Int()

	t.Require().Greater(firstCaseId, int64(0), "list testcases failed, url: "+url)
}

func (s *CaseApiSuite) TestCaseDetailApi(t provider.T) {
	t.ID("7613")
	token := httpHelper.Login()

	url := zentaoHelper.GenApiUrl(fmt.Sprintf("testcases/%d", config.CaseId), nil, constTestHelper.ZentaoSiteUrl)

	bodyBytes, _ := httpHelper.Get(url, token)

	title := gjson.Get(string(bodyBytes), "title").String()

	t.Require().Greater(len(title), 0, "get testcases failed, url: "+url)
}

func (s *CaseApiSuite) TestCaseCheckinApi(t provider.T) {
	t.ID("7616")
	token := httpHelper.Login()

	url := zentaoHelper.GenApiUrl(fmt.Sprintf("testcases/%d", config.CaseId), nil, constTestHelper.ZentaoSiteUrl)

	steps := []commDomain.ZentaoCaseStep{
		{Type: commConsts.Step, Desc: "Step 1", Expect: "Expect 1"},
		{Type: commConsts.Step, Desc: "Step 2", Expect: "Expect 1"},
		{Type: commConsts.Step, Desc: "Step 3", Expect: "Expect 1"},
	}

	title := "用例新名字" + stringUtils.NewUuid()
	requestObj := map[string]interface{}{
		"type":  "feature",
		"title": title,
		"steps": steps,

		"path":   "path_of_case",
		"script": "script_of_case",
		"lang":   "php",
	}

	bodyBytes, _ := httpHelper.Put(url, token, requestObj)

	actualTitle := gjson.Get(string(bodyBytes), "title").String()
	t.Require().Equal(actualTitle, title, "checkin testcases failed, url: "+url)

	newCase := getCase(config.CaseId)
	titleFromRemote := newCase["title"]
	t.Require().Equal(titleFromRemote, title, "get testcases failed, url: "+url)
}

func getCase(id int) (cs map[string]interface{}) {
	token := httpHelper.Login()

	url := zentaoHelper.GenApiUrl(fmt.Sprintf("testcases/%d", id), nil, constTestHelper.ZentaoSiteUrl)

	bodyBytes, _ := httpHelper.Get(url, token)

	cs = map[string]interface{}{}

	cs["id"] = gjson.Get(string(bodyBytes), "id").Int()
	cs["title"] = gjson.Get(string(bodyBytes), "title").String()

	return
}
