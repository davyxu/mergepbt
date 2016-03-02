package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/davyxu/pbmeta"
)

// 输入协议二进制描述文件,通过protoc配合github.com/davyxu/pbmeta/protoc-gen-meta插件导出
var paramPbFile = flag.String("pb", "PB", "input protobuf binary descript file, export by protoc-gen-meta plugins")
var paramMsgType = flag.String("msgtype", "", "final output msg type")
var paramPbtDir = flag.String("pbtdir", "", "pbt file directory")
var paramOut = flag.String("out", "", "output file name")

func main() {

	flag.Parse()

	// 输入协议二进制描述文件,通过protoc配合github.com/davyxu/pbmeta/protoc-gen-meta插件导出
	// 创建描述文件池
	pool, err := pbmeta.CreatePoolByFile(*paramPbFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgD := pool.MessageByFullName(*paramMsgType)
	if msgD == nil {
		fmt.Printf("mrege2msg not found: %s\n", *paramMsgType)
		return
	}

	var outBuff bytes.Buffer

	for i := 0; i < msgD.FieldCount(); i++ {

		fd := msgD.Field(i)

		outBuff.WriteString(fmt.Sprintf("%s {\n", fd.Name()))

		pbtFileName := path.Join(*paramPbtDir, fd.Name()) + ".pbt"

		// 读取子配置文件
		content, err := ioutil.ReadFile(pbtFileName)
		if err != nil {
			fmt.Printf("read pbt file failed, %v\n", err)
			return
		}

		outBuff.WriteString(string(content))

		outBuff.WriteString(fmt.Sprintf("\n}#%s\n\n", fd.Name()))

	}

	ioutil.WriteFile(*paramOut, []byte(outBuff.String()), 0644)

}
