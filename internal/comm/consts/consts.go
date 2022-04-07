package commConsts

const (
	App        = "ztf"
	AppServer  = "server"
	AppAgent   = "agent"
	AppCommand = "cmd"

	ConfigVersion = "3.0"
	ConfigDir     = "conf"
	ConfigFile    = "ztf.conf"
	LogDirName    = "log"

	ExtNameSuite = "cs"
	LogText      = "log.txt"
	ResultText   = "result.txt"
	ResultJson   = "result.json"
	ResultZip    = "result.zip"

	ExpectResultPass = "pass"

	PathInfo = "PATH_INFO"
	Get      = "GET"
)

var (
	UnitBuildToolMap = map[string]BuildTool{
		"mvn": Maven,
	}

	AutoTestTypeSelenium = "selenium"
	AutoTestTypeAppium   = "appium"
	AutoTestTypes        = []string{AutoTestTypeSelenium, AutoTestTypeAppium}

	UnitTestTypeJunit   = "junit"
	UnitTestTypeTestNG  = "testng"
	UnitTestTypeRobot   = "robot"
	UnitTestTypeCypress = "cypress"
	UnitTestPhpUnit     = "phpunit"
	UnitTestTypePyTest  = "pytest"
	UnitTestTypeJest    = "jest"
	UnitTestTypeCppUnit = "cppunit"
	UnitTestTypeGTest   = "gtest"
	UnitTestTypeQTest   = "qtest"
	UnitTestTypes       = []string{
		UnitTestTypeJunit, UnitTestTypeTestNG, UnitTestTypeRobot, UnitTestTypeCypress,
		UnitTestPhpUnit, UnitTestTypePyTest, UnitTestTypeJest, UnitTestTypeCppUnit, UnitTestTypeGTest, UnitTestTypeQTest,
	}

	UnitTestToolMvn   = "mvn"
	UnitTestToolRobot = "robot"
)
