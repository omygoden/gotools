package paths

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

//需要过滤的目录
//返回项目目录和项目名称
func ProjectPathAndName(filter ...string) (string, string) {
	var filterStr = "/bin$|/test$|cron$|job$|queue$"
	if len(filter) > 0 {
		filterStr = fmt.Sprintf("%s|%s$", filterStr, strings.Join(filter, "$|"))
	}
	pwd, _ := os.Getwd()

	reg := regexp.MustCompile(filterStr)

	pwd = reg.ReplaceAllString(pwd, "")
	p := strings.Split(pwd, "/")

	return pwd, p[len(p)-1]
}
