package main

import (
	"strconv"
	"strings"
	"testing"

	playwright "github.com/playwright-community/playwright-go"
)

var runBrowser playwright.Browser

func RunScript(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("Click node fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err := locator.Locator("text=1_string_match.php")
	if err != nil {
		t.Errorf("Find 1_string_match.php fail: %v", err)
		t.FailNow()
	}
	err = scriptLocator.Click()
	if err != nil {
		t.Errorf("Click script fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".tabs-nav-toolbar>>[title=\"Run\"]")
	if err != nil {
		t.Errorf("Click run fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行1个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec result fail: %v", err)
		t.FailNow()
	}
	element, err := page.QuerySelector("#log-list>>.msg-span>>:has-text('执行1个用例，耗时')")
	innerText, err := element.InnerText()
	if err != nil {
		t.Errorf("Find result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "1(100.0%) 失败") {
		t.Errorf("Exec 1_string_match.php fail: %v", err)
		t.FailNow()
	}
	resultTitleElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
		t.FailNow()
	}
	resultTitle, err := resultTitleElement.InnerText()
	if err != nil || resultTitle != "1_string_match.php" {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	timeElement, err := page.QuerySelector("#log-list .item .time")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err := timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTimeElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTime, err := resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}

	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunSelectedScripts(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("Click node fail: %v", err)
		t.FailNow()
	}
	err = page.Click(`[title="批量选择"]`)
	if err != nil {
		t.Errorf("The Click select btn fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err := locator.Locator(".tree-node-item:has-text('1_string_match.php')")
	if err != nil {
		t.Errorf("Find 1_string_match.php fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err = scriptLocator.Locator(".tree-node-check")
	if err != nil {
		t.Errorf("Find 1_string_match.php checkbox fail: %v", err)
		t.FailNow()
	}
	err = scriptLocator.Click()
	if err != nil {
		t.Errorf("Click 1_string_match.php checkbox fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err = locator.Locator(".tree-node-item:has-text('2_webpage_extract.php')")
	if err != nil {
		t.Errorf("Find 2_webpage_extract.php fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err = scriptLocator.Locator(".tree-node-check")
	if err != nil {
		t.Errorf("Find 2_webpage_extract.php checkbox fail: %v", err)
		t.FailNow()
	}
	err = scriptLocator.Click()
	if err != nil {
		t.Errorf("Click 2_webpage_extract.php checkbox fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".run-selected")
	if err != nil {
		t.Errorf("Click run fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行2个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec result fail: %v", err)
		t.FailNow()
	}
	element, err := page.QuerySelector("#log-list>>.msg-span>>:has-text('执行2个用例，耗时')")
	innerText, err := element.InnerText()
	if err != nil {
		t.Errorf("Find result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "1(50.0%) 通过，1(50.0%) 失败") {
		t.Errorf("Exec 1_string_match.php,2_webpage_extract.php fail: %v", err)
		t.FailNow()
	}
	resultTitleElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
		t.FailNow()
	}
	resultTitle, err := resultTitleElement.InnerText()
	if err != nil || resultTitle != "单元测试工作目录(2)" {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	timeElement, err := page.QuerySelector("#log-list .item .time")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err := timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTimeElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTime, err := resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunOpenedAndLast(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator("#siteMenuToggle")
	if err != nil {
		t.Errorf("The siteMenuToggle is missing: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("The Click is fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#navbar .list-item")
	if err != nil {
		t.Errorf("Wait for site list nav fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-title>>text=单元测试站点")
	if err != nil {
		t.Errorf("The Click site nav fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil || c == 0 {
		t.Errorf("Click workspace fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".tree-node-item:has-text('1_string_match.php')")
	if err != nil {
		t.Errorf("Click 1_string_match.php fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".tree-node-item:has-text('2_webpage_extract.php')")
	if err != nil {
		t.Errorf("Click 2_webpage_extract.php fail: %v", err)
		t.FailNow()
	}
	err = page.Click("#batchRunMenuToggle")
	if err != nil {
		t.Errorf("Click batchRunMenuToggle fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-content:has-text('执行打开文件')")
	if err != nil {
		t.Errorf("The Click Run opened scripts fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行2个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec opened scripts result fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator("#log-list>>code:has-text('执行2个用例，耗时')")
	if err != nil {
		t.Errorf("Find exec opened scripts result fail: %v", err)
		t.FailNow()
	}
	innerText, err := locator.InnerText()
	if err != nil {
		t.Errorf("Find exec opened scripts result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "1(50.0%) 通过，1(50.0%) 失败") {
		t.Errorf("Exec opened scripts fail: %v", err)
		t.FailNow()
	}

	resultTitleElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
	}
	resultTitle, err := resultTitleElement.InnerText()
	if err != nil || resultTitle != "单元测试工作目录(2)" {
		t.Errorf("Find result in rightPane fail: %v", err)
	}
	timeElement, err := locator.Locator(".time>>span")
	if err != nil {
		t.Errorf("Find log time element in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err := timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
	}
	resultTimeElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
	}
	resultTime, err := resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
	}

	err = page.Click("#batchRunMenuToggle")
	if err != nil {
		t.Errorf("Click batchRunMenuToggle fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-content:has-text('执行上次')")
	if err != nil {
		t.Errorf("The Click Run last time fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行2个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec last time result fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator("#log-list>>code:has-text('执行2个用例，耗时')")
	innerText, err = locator.InnerText()
	if err != nil {
		t.Errorf("Find exec last time result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "1(50.0%) 通过，1(50.0%) 失败") {
		t.Errorf("Exec last time fail: %v", err)
		t.FailNow()
	}
	resultTitleElement, err = page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
		t.FailNow()
	}
	resultTitle, err = resultTitleElement.InnerText()
	if err != nil || resultTitle != "单元测试工作目录(2)" {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	timeElement, err = locator.Locator(".time>>span")
	if err != nil {
		t.Errorf("Find log time element in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err = timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTimeElement, err = page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTime, err = resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunAll(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator("#siteMenuToggle")
	if err != nil {
		t.Errorf("The siteMenuToggle is missing: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("The Click is fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#navbar .list-item")
	if err != nil {
		t.Errorf("Wait for site list nav fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-title>>text=单元测试站点")
	if err != nil {
		t.Errorf("The Click site nav fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}

	err = page.Click("#batchRunMenuToggle")
	if err != nil {
		t.Errorf("Click batchRunMenuToggle fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-content:has-text('执行所有文件')")
	if err != nil {
		t.Errorf("The Click Run all scripts fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行3个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec all scripts result fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator("#log-list>>code:has-text('执行3个用例，耗时')")
	if err != nil {
		t.Errorf("Find exec all scripts result fail: %v", err)
		t.FailNow()
	}
	innerText, err := locator.InnerText()
	if err != nil {
		t.Errorf("Find exec all scripts result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "2(66.0%) 通过，1(33.0%) 失败") {
		t.Errorf("Exec all fail: %v", err)
		t.FailNow()
	}
	resultTitleElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
		t.FailNow()
	}
	resultTitle, err := resultTitleElement.InnerText()
	if err != nil || resultTitle != "单元测试工作目录(3)" {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	timeElement, err := locator.Locator(".time>>span")
	if err != nil {
		t.Errorf("Find log time element in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err := timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTimeElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTime, err := resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunReExecFailCase(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}

	locator, err := page.Locator("#siteMenuToggle")
	if err != nil {
		t.Errorf("The siteMenuToggle is missing: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("The Click is fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#navbar .list-item")
	if err != nil {
		t.Errorf("Wait for site list nav fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-title>>text=单元测试站点")
	if err != nil {
		t.Errorf("The Click site nav fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	err = page.Click("#rightPane .result-list-item .list-item-title>>nth=0")
	if err != nil {
		t.Errorf("Click first result fail: %v", err)
	}
	err = page.Click(".result-action .btn:has-text('重新执行失败用例')")
	if err != nil {
		t.Errorf("Click re-exec failed case btn fail: %v", err)
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行1个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec script result fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator("#log-list>>code:has-text('执行1个用例，耗时')")
	if err != nil {
		t.Errorf("Find exec script log fail: %v", err)
		t.FailNow()
	}
	innerText, err := locator.InnerText()
	if err != nil {
		t.Errorf("Find exec script result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "0(0.0%) 通过，1(100.0%) 失败") {
		t.Errorf("Exec failed case fail: %v", err)
		t.FailNow()
	}
	resultTitleElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
		t.FailNow()
	}
	resultTitle, err := resultTitleElement.InnerText()
	if err != nil || resultTitle != "1_string_match.php" {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	timeElement, err := locator.Locator(".time>>span")
	if err != nil || resultTitle != "1_string_match.php" {
		t.Errorf("Find log time element in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err := timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTimeElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTime, err := resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunReExecAllCase(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}

	locator, err := page.Locator("#siteMenuToggle")
	if err != nil {
		t.Errorf("The siteMenuToggle is missing: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("The Click is fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#navbar .list-item")
	if err != nil {
		t.Errorf("Wait for site list nav fail: %v", err)
		t.FailNow()
	}
	err = page.Click(".list-item-title>>text=单元测试站点")
	if err != nil {
		t.Errorf("The Click site nav fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	err = page.Click("#rightPane .result-list-item .list-item-title>>nth=0")
	if err != nil {
		t.Errorf("Click first result fail: %v", err)
	}
	err = page.Click(".result-action .btn:has-text('重新执行所有用例')")
	if err != nil {
		t.Errorf("Click re-exec failed case btn fail: %v", err)
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行3个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec script result fail: %v", err)
		t.FailNow()
	}
	locator, err = page.Locator("#log-list>>code:has-text('执行3个用例，耗时')")
	if err != nil {
		t.Errorf("Find exec script log fail: %v", err)
		t.FailNow()
	}
	innerText, err := locator.InnerText()
	if err != nil {
		t.Errorf("Find exec script result fail: %v", err)
		t.FailNow()
	}
	if !strings.Contains(innerText, "2(66.0%) 通过，1(33.0%) 失败") {
		t.Errorf("Exec failed case fail: %v", err)
		t.FailNow()
	}
	resultTitleElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-title")
	if err != nil {
		t.Errorf("Find log title in logPane fail: %v", err)
		t.FailNow()
	}
	resultTitle, err := resultTitleElement.InnerText()
	if err != nil || resultTitle != "单元测试工作目录(3)" {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	timeElement, err := locator.Locator(".time>>span")
	if err != nil {
		t.Errorf("Find log time element in logPane fail: %v", err)
		t.FailNow()
	}
	logTime, err := timeElement.InnerText()
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTimeElement, err := page.QuerySelector("#rightPane .result-list-item .list-item-trailing-text")
	if err != nil {
		t.Errorf("Find log time in logPane fail: %v", err)
		t.FailNow()
	}
	resultTime, err := resultTimeElement.InnerText()
	if err != nil || logTime[:5] != resultTime {
		t.Errorf("Find result in rightPane fail: %v", err)
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunFailStatistic(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("Click node fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err := locator.Locator("text=1_string_match.php")
	if err != nil {
		t.Errorf("Find 1_string_match.php fail: %v", err)
		t.FailNow()
	}
	err = scriptLocator.Click()
	if err != nil {
		t.Errorf("Click script fail: %v", err)
		t.FailNow()
	}
	elements, _ := page.QuerySelectorAll(".statistic>>span")
	runTimes, _ := elements[0].InnerText()
	failTimes, _ := elements[2].InnerText()
	runTimesInt, _ := strconv.Atoi(runTimes)
	failTimesInt, _ := strconv.Atoi(failTimes)
	err = page.Click(".tabs-nav-toolbar>>[title=\"Run\"]")
	if err != nil {
		t.Errorf("Click run fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行1个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec result fail: %v", err)
		t.FailNow()
	}
	page.WaitForTimeout(200)
	elements, _ = page.QuerySelectorAll(".statistic>>span")
	runTimes2, _ := elements[0].InnerText()
	failTimes2, _ := elements[2].InnerText()
	runTimes2Int, _ := strconv.Atoi(runTimes2)
	failTimes2Int, _ := strconv.Atoi(failTimes2)
	if runTimes2Int-runTimesInt != 1 || failTimes2Int-failTimesInt != 1 {
		t.Error("statistic error")
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunSuccessStatistic(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("Click node fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err := locator.Locator("text=2_webpage_extract.php")
	if err != nil {
		t.Errorf("Find 2_webpage_extract.php fail: %v", err)
		t.FailNow()
	}
	err = scriptLocator.Click()
	if err != nil {
		t.Errorf("Click script fail: %v", err)
		t.FailNow()
	}
	elements, _ := page.QuerySelectorAll(".statistic>>span")
	runTimes, _ := elements[0].InnerText()
	succTimes, _ := elements[1].InnerText()
	runTimesInt, _ := strconv.Atoi(runTimes)
	succTimesInt, _ := strconv.Atoi(succTimes)
	err = page.Click(".tabs-nav-toolbar>>[title=\"Run\"]")
	if err != nil {
		t.Errorf("Click run fail: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector("#log-list>>.msg-span>>:has-text('执行1个用例，耗时')")
	if err != nil {
		t.Errorf("Wait exec result fail: %v", err)
		t.FailNow()
	}
	page.WaitForTimeout(200)
	elements, _ = page.QuerySelectorAll(".statistic>>span")
	runTimes2, _ := elements[0].InnerText()
	succTimes2, _ := elements[1].InnerText()
	runTimes2Int, _ := strconv.Atoi(runTimes2)
	succTimes2Int, _ := strconv.Atoi(succTimes2)
	if runTimes2Int-runTimesInt != 1 || succTimes2Int-succTimesInt != 1 {
		t.Error("statistic error")
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func RunBugStatistic(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	headless := true
	var slowMo float64 = 100
	runBrowser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless, SlowMo: &slowMo})
	if err != nil {
		t.Errorf("Fail to launch the web runBrowser: %v", err)
		t.FailNow()
	}
	page, err := runBrowser.NewPage()
	if err != nil {
		t.Errorf("Create the new page fail: %v", err)
		t.FailNow()
	}
	if _, err = page.Goto("http://127.0.0.1:8000/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded}); err != nil {
		t.Errorf("The specific URL is missing: %v", err)
		t.FailNow()
	}
	_, err = page.WaitForSelector(".tree-node")
	if err != nil {
		t.Errorf("Wait tree-node fail: %v", err)
		t.FailNow()
	}
	locator, err := page.Locator(".tree-node", playwright.PageLocatorOptions{HasText: "单元测试工作目录"})
	c, err := locator.Count()
	if err != nil || c == 0 {
		t.Errorf("Find workspace fail: %v", err)
		t.FailNow()
	}
	err = locator.Click()
	if err != nil {
		t.Errorf("Click node fail: %v", err)
		t.FailNow()
	}
	scriptLocator, err := locator.Locator("text=1_string_match.php")
	if err != nil {
		t.Errorf("Find 1_string_match.php fail: %v", err)
		t.FailNow()
	}
	err = scriptLocator.Click()
	if err != nil {
		t.Errorf("Click script fail: %v", err)
		t.FailNow()
	}
	page.WaitForTimeout(200)
	elements, _ := page.QuerySelectorAll(".statistic>>span")
	bugTimes, _ := elements[3].InnerText()
	bugTimesInt, _ := strconv.Atoi(bugTimes)
	err = page.Click(".statistic>>span>>nth=3")
	if err != nil {
		t.Errorf("Click bug btn fail: %v", err)
		t.FailNow()
	}
	elements, _ = page.QuerySelectorAll("#bugsModal>>tr")

	bugTimes2Int := len(elements)
	if bugTimes2Int-1 != bugTimesInt {
		t.Error("statistic error")
		t.FailNow()
	}
	if err = runBrowser.Close(); err != nil {
		t.Errorf("The runBrowser cannot be closed: %v", err)
		t.FailNow()
	}
	if err = pw.Stop(); err != nil {
		t.Errorf("The playwright cannot be stopped: %v", err)
		t.FailNow()
	}
}

func TestUiRun(t *testing.T) {
	t.Run("RunScript", RunScript)
	t.Run("RunSelectedScripts", RunSelectedScripts)
	t.Run("RunOpenedAndLast", RunOpenedAndLast)
	t.Run("RunAll", RunAll)
	t.Run("RunReExecAllCase", RunReExecAllCase)
	t.Run("RunReExecFailCase", RunReExecFailCase)
	t.Run("RunSuccessStatistic", RunSuccessStatistic)
	t.Run("RunFailStatistic", RunFailStatistic)
	t.Run("RunBugStatistic", RunBugStatistic)
}
