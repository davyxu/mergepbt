# mergepbt


## 解决protobuf-net不支持解析文本格式pb数据文件

将pbt合并并转化二进制

# 参数说明

* pb 输入带有客户端配置文件pb类型信息的二进制描述文件, 需要配合github.com/davyxu/pbmeta/protoc-gen-meta插件导出

* msgtype 客户端配置文件pb类型的完整pb类型描述

* pbtdir 客户端配置文件pb类型中字段对应pbt文件的搜索路径

* out 输出文件, 格式为pbt

# 客户端配置文件类型格式

```protobuf

message ClientConfig
{
	optional ItemFile Item = 3;	
	optional AIFile AI = 9;
}


```
* 一个pbt文件对应一个字段

* 字段的名字添加.pbt为输入pbt文件, 例如: Item字段对应Item.pbt



# 安装方法

	go get github.com/davyxu/mergepbt
	
	go install github.com/davyxu/mergepbt

# 使用方法

* 使用github.com/davyxu/tabtoy导出各种pbt

* 根据一个客户端配置文件pb类型描述, 将字段对应的pbt文件合并成1个pbt文件

..\tools\mergepbt.exe --pb=.\obj\game.pb --out=merge.pbt --msgtype=gamedef.ClientConfig --pbtdir=.\obj

* 利用protoc的编码功能转文本为二进制

type merge.pbt | protoc.exe --encode=gamedef.ClientConfig --proto_path ..\proto ..\proto\sys.proto > config.pbb

# 链接

* protobuf-net运行库

	https://github.com/mgravell/protobuf-net
	
* 电子表格强力导出器

	https://github.com/davyxu/tabtoy


# 备注

感觉不错请star, 谢谢!

博客: http://www.cppblog.com/sunicdavy

知乎: http://www.zhihu.com/people/xu-bo-62-87

邮箱: sunicdavy@qq.com
