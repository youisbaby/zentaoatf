package main

/**

cid=0
pid=0

1.提交bug到禅道 >> Success

*/
import (
	"fmt"
	"regexp"
	"runtime"
	"time"

	expect "github.com/google/goexpect"
)

var (
	continueRE = regexp.MustCompile("Which case do you want to report bug for|请输入您想提交缺陷的用例ID")
	successRE  = regexp.MustCompile("Success to report bug for case \\d+|成功为用例\\d+提交缺陷")
	newline    = "\n"
)

func testCi() {
	cmd := "ztf cb demo/001 -p 1"
	if runtime.GOOS == "windows" {
		cmd = "ztf cb demo\\001 -p 1"
	}
	child, _, err := expect.Spawn(cmd, -1)
	if err != nil {
		fmt.Println(err)
	}
	defer child.Close()
	if _, _, err := child.Expect(continueRE, 2*time.Second); err != nil {
		fmt.Printf("%s: %s%s", continueRE, err, newline)
		return
	}

	if err = child.Send("1" + newline); err != nil {
		fmt.Println(err)
		return
	}

	if _, _, err = child.Expect(successRE, 10*time.Second); err != nil {
		fmt.Printf("%s: %s%s", successRE, err, newline)
		return
	}

	fmt.Println("Success")
}

func main() {
	if runtime.GOOS == "windows" {
		newline = "\r\n"
	}
	testCi()
}
