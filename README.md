## Introduce
框架：wails2
前端实现： Vue3
后端实现： Golang

基于 Wails2 实现类似 switchhost 的功能。
1、左边栏显示文件列表，文件来源默认本地文件目录，如 /.quickcmd，文件后缀为 .quickcmd;
2、文件列表每个文件增加一个开关，可以开启或关闭
3、右边栏为文本框，默认显示左边栏文件列表第一个文件的内容，可以手动点击任意文件，文本框显示对应文件内容。文本框支持增删改。
4、增加按钮，可以点击创建文件、删除文件

# 后端功能接口
- gorm 数据库存储☑️
- 数据初始化☑️
- 获取文件列表及文件内容☑️
- 保存文件修改☑️
- 新增文件☑️
- 删除文件☑️


## 安装 waild
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 运行项目测试
```
wails dev
```

## 编译项目
```
wails build
```