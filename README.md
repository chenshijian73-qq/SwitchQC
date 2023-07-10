## Introduce
以文件分类的形式，将所有的别名和变量集中在一个统一的地方进行管理，更好地管理配置系统中的别名和变量，方便人员快速找到需要的配置项并进行修改，以及配置自己的快捷简洁命令
![appicon.png](build%2Fappicon.png)
- 支持系统：Mac、Linux
- 框架：wails2
- 前端： Vue3 + Arco
- 后端： Golang

## How to build

### install waild
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### develop the program
```
wails dev
```

### build app
```
wails build -clean
```

### run app
```
# mac: After copying, double-click the app to run
cp -r build/bin/QC.app ~/Applications/
```

## How to Use
When you launch the program, the app screen is shown like this:
![img.png](images/img.png)
Click any file in the navigation bar to view and edit it:
![img_1.png](images/img_1.png)
Edit it and Open new terminal to test
![img_7.png](images/img_7.png)
![img_8.png](images/img_8.png)
Delete file:
![img_3.png](images/img_3.png)
Add file:
![img_4.png](images/img_4.png)
Open RecycleBin:
![img_5.png](images/img_5.png)
Retore or Remove file in RecycleBin
![img_6.png](images/img_6.png)

