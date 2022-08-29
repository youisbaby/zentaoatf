package main

import (
	"flag"
	"github.com/easysoft/zentaoatf/internal/command/action"
	commandConfig "github.com/easysoft/zentaoatf/internal/command/config"
	"github.com/easysoft/zentaoatf/internal/pkg/consts"
	unitHelper "github.com/easysoft/zentaoatf/internal/pkg/helper/unit"
	fileUtils "github.com/easysoft/zentaoatf/pkg/lib/file"
	i118Utils "github.com/easysoft/zentaoatf/pkg/lib/i118"
	logUtils "github.com/easysoft/zentaoatf/pkg/lib/log"
	stringUtils "github.com/easysoft/zentaoatf/pkg/lib/string"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	AppVersion string
	BuildTime  string
	GoVersion  string
	GitHash    string

	language        string
	independentFile bool
	keywords        string

	productId    string
	moduleId     string
	taskIdOrName string
	suiteId      string
	taskName     string

	noNeedConfirm bool

	flagSet *flag.FlagSet
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("ztf", flag.ContinueOnError)

	flagSet.StringVar(&productId, "p", "", "")
	flagSet.StringVar(&productId, "product", "", "")

	flagSet.StringVar(&moduleId, "m", "", "")
	flagSet.StringVar(&moduleId, "module", "", "")

	flagSet.StringVar(&suiteId, "s", "", "")
	flagSet.StringVar(&suiteId, "suite", "", "")

	flagSet.StringVar(&taskIdOrName, "t", "", "")
	flagSet.StringVar(&taskIdOrName, "task", "", "")

	flagSet.StringVar(&language, "l", "", "")
	flagSet.StringVar(&language, "language", "", "")

	flagSet.BoolVar(&independentFile, "i", false, "")
	flagSet.BoolVar(&independentFile, "independent", false, "")

	flagSet.StringVar(&commConsts.Interpreter, "I", "", "")
	flagSet.StringVar(&commConsts.Interpreter, "interpreter", "", "")

	flagSet.StringVar(&keywords, "k", "", "")
	flagSet.StringVar(&keywords, "keywords", "", "")

	flagSet.BoolVar(&commConsts.AutoCommitResult, "cr", false, "")
	flagSet.BoolVar(&commConsts.AutoCommitBug, "cb", false, "")
	flagSet.BoolVar(&noNeedConfirm, "y", false, "")
	flagSet.BoolVar(&commConsts.Verbose, "verbose", false, "")

	flagSet.StringVar(&commConsts.AllureReportDir, "allureReportDir", "", "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "run", ".")
	}

	switch os.Args[1] {
	case "set", "-set":
		flagSet.Parse(os.Args[2:])
		action.Set()

	case "expect":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		action.GenExpectFiles(files)

	case "extract":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		action.Extract(files)

	case "checkout", "co":
		checkout()

	case "ci":
		ci()

	case "cr":
		cr()

	case "cb":
		cb()

	case "list", "ls", "-l":
		list()

	case "view", "-v":
		view()

	case "clean", "-clean", "-c":
		action.Clean()

	case "version", "--version":
		logUtils.PrintVersion(AppVersion, BuildTime, GoVersion, GitHash)

	case "help", "-h", "-help", "--help":
		action.PrintUsage()

	case "run", "-r":
		run(os.Args)

	default: // run
		if len(os.Args) > 1 {
			args := []string{os.Args[0], "run"}
			args = append(args, os.Args[1:]...)

			run(args)
		} else {
			action.PrintUsage()
		}
	}
}

func checkout() {
	if err := flagSet.Parse(os.Args[2:]); err == nil {
		action.Checkout(productId, moduleId, suiteId, taskIdOrName, independentFile, language)
	}
}

func ci() {
	files := fileUtils.GetFilesFromParams(os.Args[2:])
	if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
		action.CheckIn(files)
	}
}

func cr() {
	files := fileUtils.GetFilesFromParams(os.Args[2:])
	if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
		action.CommitZTFTestResult(files, stringUtils.ParseInt(productId),
			taskIdOrName, noNeedConfirm)
	}
}

func cb() {
	files := fileUtils.GetFilesFromParams(os.Args[2:])
	if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
		action.CommitBug(files, stringUtils.ParseInt(productId), noNeedConfirm)
	}
}

func list() {
	files := fileUtils.GetFilesFromParams(os.Args[2:])
	if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
		if len(files) == 0 {
			files = append(files, ".")
		}

		action.List(files, keywords)
	}
}

func view() {
	files := fileUtils.GetFilesFromParams(os.Args[2:])
	if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
		action.View(files, keywords)
	}
}

func run(args []string) {
	if len(args) >= 3 && stringUtils.FindInArr(args[2], commConsts.UnitTestTypes) { // unit test
		runUnitTest(args)
	} else { // ztf test
		runFuncTest(args)

		if commConsts.AutoCommitResult && productId != "" {
			action.CommitZTFTestResult([]string{commConsts.ExecLogDir},
				stringUtils.ParseInt(productId), taskIdOrName, true)
		}

		if commConsts.AutoCommitBug && productId != "" {
			action.CommitBug([]string{commConsts.ExecLogDir}, stringUtils.ParseInt(productId), true)
		}
	}
}

func runFuncTest(args []string) {
	files := fileUtils.GetFilesFromParams(args[2:])

	err := flagSet.Parse(args[len(files)+2:])
	if err != nil {
		action.PrintUsage()
		return
	}
	if len(files) > 0 && files[0] != "" && !fileUtils.FileExist(files[0]) {
		action.PrintUsage()
		return
	}

	commConsts.ProductId = productId

	if len(files) == 0 {
		files = append(files, ".")
	}

	if commConsts.Interpreter != "" {
		msgStr := i118Utils.Sprintf("run_with_specific_interpreter", commConsts.Interpreter)
		logUtils.ExecConsolef(color.FgCyan, msgStr)
	}
	action.RunZTFTest(files, moduleId, suiteId, taskIdOrName)
}

func runUnitTest(args []string) {
	// junit -p 1 mvn clean package test
	commConsts.UnitTestType = args[2]
	end := 8
	if end > len(args)-1 {
		end = len(args) - 1
	}
	flagSet.Parse(args[3:])

	start := 3
	if productId != "" {
		start = start + 2
		commConsts.ProductId = productId
	}
	if taskIdOrName != "" {
		start = start + 2
	}
	if commConsts.AllureReportDir != "" {
		start = start + 2
	}
	if commConsts.Verbose {
		start = start + 1
	}

	unitHelper.GetUnitTools(args, start)

	cmd := strings.Join(args[start:], " ")

	action.RunUnitTest(cmd, taskIdOrName)
}

func init() {
	cleanup()
	commandConfig.Init()
}

func cleanup() {
	color.Unset()
}
