## Introduce
In the form of file classification, all aliases and variables are centralized in the desktop application for management, so that personnel can quickly find and modify the required configuration items, as well as add quick and concise commands to configure themselves.

![appicon.png](build%2Fappicon.png)
- **Supported system**：Mac、Linux
- **Frame**：wails2
- **Frontend**： Vue3 + Arco
- **Backend**： Golang

## How to build

### install waild
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
**Reference：** https://wails.io/docs/gettingstarted/installation

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


## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=chenshijian73-qq/quickcmd&type=Date)](https://star-history.com/#chenshijian73-qq/quickcmd&Date)
