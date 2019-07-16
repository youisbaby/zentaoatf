package utils

import (
	"github.com/easysoft/zentaoatf/src/misc"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

func RemoveBlankLine(str string) string {
	myExp := regexp.MustCompile(`\n{2,}`) // 连续换行
	ret := myExp.ReplaceAllString(str, "\n")

	ret = strings.Trim(ret, "\n")
	ret = strings.TrimSpace(ret)

	return ret
}

func ScriptToLogName(dir string, file string) string {
	logDir := dir + string(os.PathSeparator) + "logs"
	MkDir(logDir)

	nameSuffix := path.Ext(file)
	nameWithSuffix := path.Base(file)
	name := strings.TrimSuffix(nameWithSuffix, nameSuffix)

	logFile := logDir + string(os.PathSeparator) + name + ".log"

	return logFile
}

func ScriptToExpectName(file string) string {
	fileSuffix := path.Ext(file)
	expectName := strings.TrimSuffix(file, fileSuffix) + ".ex"

	return expectName
}

func BoolToPass(b bool) string {
	if b {
		return misc.PASS.String()
	} else {
		return misc.FAIL.String()
	}
}

func GetOs() string {
	osName := runtime.GOOS

	if osName == "darwin" {
		return "mac"
	} else {
		return osName
	}
}
func IsMac() bool {
	return GetOs() == "mac"
}
