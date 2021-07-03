//Package tcore config file
package tcore

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// 1 创建 TConfig 对象
// 2 调用Analysis方法，传入要解析的文件，返回解析状态true/false
// 3 调用结构体的GetDataInt32/GetDataFloat32/GetDataString方法，传入文件节点名字，获取到文件信息
// 4 配置文件参考LogicServer.cfg

//TConfig export
type TConfig struct {
	mMap map[string]string
}

//Analysis export
func (pOwn *TConfig) Analysis(aFileName string) bool {
	file, err := os.Open(aFileName)
	if err != nil {
		return false
	}
	defer file.Close()
	pOwn.mMap = make(map[string]string)
	reader := bufio.NewReader(file)
	for {
		bytes, bIsPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil || bIsPrefix == true {
			return false
		}

		line := string(bytes)
		str1 := strings.Replace(line, " ", "", -1)
		str1 = strings.Replace(str1, "\n", "", -1)
		str1 = strings.Replace(str1, "\r", "", -1)
		str2 := strings.Split(str1, "#")
		if str2[0] == "" {
			continue
		}
		str3 := strings.Split(str2[0], "=")
		if len(str3) == 2 && str3[0] != "" && str3[1] != "" {
			pOwn.mMap[str3[0]] = str3[1]
		} else {
			return false
		}
	}
	return true
}

//GetDataInt32 export
func (pOwn *TConfig) GetDataInt32(aKey string) (int32, bool) {
	value, bIsOK := pOwn.mMap[aKey]
	if bIsOK == false {
		return 0, false
	}
	nRes, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return int32(nRes), true
}

//GetDataFloat32 export
func (pOwn *TConfig) GetDataFloat32(aKey string) (float32, bool) {
	value, isOK := pOwn.mMap[aKey]
	if isOK == false {
		return 0, false
	}
	fRes, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, false
	}
	return float32(fRes), true
}

//GetDataString export
func (pOwn *TConfig) GetDataString(aKey string) (string, bool) {
	value, isOK := pOwn.mMap[aKey]
	return value, isOK
}
