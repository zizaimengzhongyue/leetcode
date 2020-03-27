package main

import "fmt"
import "strings"

func simplifyPath(path string) string {
	ans := [][]byte{}
	bPath := []byte(path)
	tmp := []byte{}
	for _, v := range bPath {
		if v == '/' {
			if len(tmp) == 0 {
				tmp = append(tmp, v)
			} else if tmp[len(tmp)-1] == '/' {
				continue
			} else if len(tmp) == 2 && tmp[1] == '.' {
				tmp = []byte{'/'}
				continue
			} else {
				ans = append(ans, append([]byte{}, tmp...))
				tmp = []byte{'/'}
			}
		} else {
			tmp = append(tmp, v)
		}
	}
	if (len(tmp) > 1 && tmp[1] != '/') || len(ans) == 0 {
		ans = append(ans, tmp)
	}
	res := []string{}
	for _, v := range ans {
		if (len(v) == 2 || len(v) == 3) && v[len(v)-1] == '.' {
			if len(res) > 0 && len(v) == 3 {
				res = res[:len(res)-1]
			} else if len(res) == 0 {
				res = []string{"/"}
			}
		} else {
			if len(res) == 1 && len(res[0]) == 1 {
				res = []string{string(v)}
			} else {
				res = append(res, string(v))
			}
		}
	}
	strAns := strings.Join(res, "")
	if len(strAns) == 0 {
		return "/"
	}
	return strAns
}

func main() {
	str := "/home/"
	fmt.Println(simplifyPath(str))
	str = "/../"
	fmt.Println(simplifyPath(str))
	str = "/home//foo/"
	fmt.Println(simplifyPath(str))
	str = "/a/./b/../../c/"
	fmt.Println(simplifyPath(str))
	str = "/a/../../b/../c//.//"
	fmt.Println(simplifyPath(str))
	str = "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(str))
	str = "/"
	fmt.Println(simplifyPath(str))
	str = "/.."
	fmt.Println(simplifyPath(str))
	str = "/."
	fmt.Println(simplifyPath(str))
	str = "/..."
	fmt.Println(simplifyPath(str))
	str = "/home/../../.."
	fmt.Println(simplifyPath(str))
	str = "/home/of/foo/../../bar/../../is/./here/."
	fmt.Println(simplifyPath(str))
}
