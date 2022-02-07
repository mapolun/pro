package fn

import (
	"os"
	"path/filepath"
)

//判断文件是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

//判断目录是否存在
func DirExists(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//创建文件
func GetFile(filePath string) (*os.File, error) {
	path := filepath.Dir(filePath)

	//写入路径
	isDir := DirExists(path)
	if isDir == false {
		err := os.MkdirAll(path, 0775)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//获取根目录
func GetRootPath() string {
	binary, _ := os.Executable()
	root := filepath.Dir(filepath.Dir(binary))
	return root
}
