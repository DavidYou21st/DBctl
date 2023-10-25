# fctl

> fctl 辅助工具，提供创建项目、快速生成功能模块的功能

## 下载并使用
```bash
go build -o .    #本地go build 
```

目前只支持本地编译，不支持在线安装


## 生成业务模块

### 通过数据库生成gorm gen dal代码

```bash
USAGE:
   fctl gorm [arguments...]

OPTIONS:
   -table  value,  表名称，支持多个表 user,order 中间用逗号隔开,如果为空,则生成全部表 ,也可以在config.yaml  Dal Table 字段配置
   -config value,  配置 默认 config/config.yaml
   -outpath value,   生成model 目录 ,也可以在config.yaml Dal OutPath 字段配置
```

#### 示例：通过数据库生成dal代码

```bash
./fctl gorm  -config config/config.yaml -outpath ./dal/model  -table sys_menus

# 如果在config.yaml配置相关参数则可以执行如下命令即可
./fctl gorm 
```


### 生成dto
```
 ./fctl module dto -a ./schema/user.api -d ./schema/ -f user -n console
```

### 生成biz
```
 ./fctl module biz -d ./schema -f user -t user -n console
```

### 生成service
```
./fctl module service -d ./schema -f sysTenant -t sys_tanant  -n console -a ./schema/tenant.api  
```
### 生成data
```
./fctl module data -d ./schema -f sysTenant -t sys_tanant  -n console
```

### 生成router

```
./fctl module router -d ./schema -f user -n console
```

### 生成inject 文件

```
./fctl inject all -d ./schema -f SysTenant -n console
```

### proto文件转pb文件
```
protoc --proto_path=. \
--proto_path=./third_party \
--go_out=paths=source_relative:. \
--go-http_out=paths=source_relative:. \
--validate_out=paths=source_relative,lang=go:. \
probuf/admin/test.proto

protoc -I ./ \
-I /Users/php/Desktop/langrensha/fctl/third_party \
-I /Users/php/Desktop/langrensha/fctl/third_party/google/api \
--go-grpc_out=paths=source_relative:. \
--go_out=paths=source_relative:. \
--grpc-gateway_out=logtostderr=true:./genproto
probuf/admin/test.proto

protoc工具需要安装，请自行查阅资料安装
```

支持
#### 编译本工具
如果你要打包给其他人使用你最新版本的，可以使用如下脚本进行打包该工具。
该工具依赖 upx ,安装upx [点击参考](https://www.wolai.com/bullteam/jeg4DM9RDKDRiGR5hiPtsy)
```bash
# 授权
sudo chmod 777 ./script/pack.sh
# 执行
./script/pack.sh
```