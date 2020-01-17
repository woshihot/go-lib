package os

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func IsWindows() bool {
	return "windows" == runtime.GOOS
}
func IsLinux() bool {
	return "linux" == runtime.GOOS
}

const (
	RELATIVE_PREFIX = "./" //以当前目录相对路径
)

func DeviceChooseString(bean ...string) string {

	var length = len(bean)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return bean[0]
	}
	if length == 2 {
		//第一位window 第二位linux
		return deviceChooseString(bean[0], bean[1])
	}
	//default return linux
	return bean[1]

}

func deviceChooseString(w string, l string) string {
	if IsWindows() {
		return w
	} else if IsLinux() {
		return l
	}
	return ""
}
func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

func GetAbsolutely(path string) string {

	if IsLinux() {
		return getLinuxAbsolutely(path)
	} else if IsWindows() {
		return getWinAbsolutely(path)
	}

	return path
}
func getLinuxAbsolutely(path string) string {
	if strings.HasPrefix(path, "/") {
		//当前就是绝对路径直接返回
		return path
	}
	path = strings.TrimPrefix(path, RELATIVE_PREFIX) //将./统一去掉，拼接当前路径
	currentPath := GetCurPath()
	return currentPath + "/" + path
}

func getWinAbsolutely(path string) string {

	if strings.HasPrefix(path, RELATIVE_PREFIX) {
		//以相对路径开头
		path = strings.TrimPrefix(path, RELATIVE_PREFIX) //将./统一去掉，拼接当前路径
		currentPath := GetCurPath()

		for {
			if strings.HasPrefix(path, "../") {
				path = strings.TrimPrefix(path, "../")
				currentPath = strings.TrimPrefix(currentPath, "\\")
				index := strings.LastIndex(currentPath, "\\")
				currentPath = currentPath[:index]
			} else {
				break
			}
		}

		return currentPath + "\\" + path
	} else {
		return path
	}

}

func Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func MkdirAll(name string, perm os.FileMode) error {
	return os.MkdirAll(name, perm)
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func Remove(path string) error {
	return os.Remove(path)
}
