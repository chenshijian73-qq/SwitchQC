## Introduce
介绍：以文件的形式，统一管理配置系统的别名，变量
支持系统：Mac、Linux
框架：wails2
前端实现： Vue3 + Arco
后端实现： Golang

# 开发项
- 文件列表☑️
- 新增文件☑️
- 保存内容☑️
- 删除文件☑️
- 文件使用开关☑️
- 文本框更换为代码编辑器☑️
- 回收站功能☑️
- 数据为空的显示
- 删除文件不存在报错处理
- 删除文件 编辑框依旧可用

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
wails build -clean
```