// file create by daihao, time is 2018/4/13 17:11
package utils

import "fmt"

func GetUrl(path string) string {
	return fmt.Sprintf("%s%s", HtmlPath, path)
}
