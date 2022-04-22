package action

import (
	"fmt"
	commConsts "github.com/easysoft/zentaoatf/internal/comm/consts"
	commDomain "github.com/easysoft/zentaoatf/internal/comm/domain"
	analysisHelper "github.com/easysoft/zentaoatf/internal/comm/helper/analysis"
	configHelper "github.com/easysoft/zentaoatf/internal/comm/helper/config"
	"github.com/easysoft/zentaoatf/internal/comm/helper/zentao"
	i118Utils "github.com/easysoft/zentaoatf/internal/pkg/lib/i118"
	logUtils "github.com/easysoft/zentaoatf/internal/pkg/lib/log"
	stdinUtils "github.com/easysoft/zentaoatf/internal/pkg/lib/stdin"
	stringUtils "github.com/easysoft/zentaoatf/internal/pkg/lib/string"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

var (
	bug       commDomain.ZtfBug
	bugFields commDomain.ZentaoBugFields
)

func CommitBug(files []string, productId int) {
	var resultDir string
	if len(files) > 0 {
		resultDir = files[0]
	} else {
		stdinUtils.InputForDir(&resultDir, "", "result")
	}

	if productId == 0 {
		productIdStr := stdinUtils.GetInput("\\d+", "",
			i118Utils.Sprintf("pls_enter")+" "+i118Utils.Sprintf("product_id"))
		productId, _ = strconv.Atoi(productIdStr)
	}

	report, err := analysisHelper.ReadReportByWorkspaceSeq(commConsts.WorkDir, resultDir)
	if err != nil {
		return
	}

	ids := make([]string, 0)
	lines := make([]string, 0)
	for _, cs := range report.FuncResult {
		if cs.Status != commConsts.PASS {
			lines = append(lines, fmt.Sprintf("%d. %s %s", cs.Id, cs.Title, coloredStatus(cs.Status)))
			ids = append(ids, strconv.Itoa(cs.Id))
		}
	}

	if len(lines) == 0 {
		logUtils.ExecConsole(color.FgCyan, i118Utils.Sprintf("no_failed_case_to_report_bug"))
		return
	}

	for {
		logUtils.ExecConsole(color.FgCyan, "\n"+i118Utils.Sprintf("enter_case_id_for_report_bug"))
		logUtils.ExecConsole(color.FgCyan, strings.Join(lines, "\n"))
		var caseId string
		fmt.Scanln(&caseId)
		if caseId == "exit" {
			color.Unset()
			os.Exit(0)
		} else {
			if stringUtils.FindInArr(caseId, ids) {
				reportBug(resultDir, caseId, productId)
			} else {
				logUtils.ExecConsole(color.FgRed, i118Utils.Sprintf("invalid_input"))
			}
		}
	}
}

func coloredStatus(status commConsts.ResultStatus) string {
	temp := strings.ToLower(status.String())

	switch temp {
	case "pass":
		return color.GreenString(i118Utils.Sprintf(temp))
	case "fail":
		return color.RedString(i118Utils.Sprintf(temp))
	case "skip":
		return color.YellowString(i118Utils.Sprintf(temp))
	}

	return status.String()
}

func reportBug(resultDir string, caseId string, productId int) error {
	config := configHelper.LoadByWorkspacePath(commConsts.WorkDir)
	bugFields, _ = zentaoHelper.GetBugFiledOptions(config, bug.Product)

	bug = zentaoHelper.PrepareBug(commConsts.WorkDir, resultDir, caseId, productId)

	err := zentaoHelper.CommitBug(bug, config)
	return err
}

func getFirstNoEmptyVal(options []commDomain.BugOption) string {
	for _, opt := range options {
		if opt.Name != "" {
			return opt.Code
		}
	}

	return ""
}

func getNameById(id string, options []commDomain.BugOption) string {
	for _, opt := range options {
		if opt.Code == id {
			return opt.Name
		}
	}

	return ""
}
