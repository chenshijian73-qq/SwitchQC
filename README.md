<div align="center">
  <img width="200" src="build/appicon.png">
  <h1> SwitchQC</h1>
</div>


## Introduce For SwitchQC
1. Similar to SwitchHost, SwitchQC is used to manage system variables and aliases for mac/linux;
2. Nest user-supplied AI Chat links to chat with AI;

- **Supported system**：Mac、Linux
- **Frame**：wails2
- **Frontend**： Vue3 + Arco.design
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
cp -r build/bin/SwitchQC.app /Applications
```

## How to Use
### 1、SwitchQC Usage
**When you launch the program, the app screen is shown like this**
![img_1.png](images/img_1.png)

**Add file**

![img_2.png](images/img_2.png)

**Delete file**

![img_3.png](images/img_3.png)

**Open RecycleBin**

![img_4.png](images/img_4.png)

### 2、AI Chat Usage
**Hover the cat，click to AI chat page**

![img_5.png](images/img_5.png)

**Select which AI link to visit**

![img_6.png](images/img_6.png)

![img_7.png](images/img_7.png)

**Manage your AI link** 

![img_8.png](images/img_8.png)

### 3、How to get your AI Chat
- Come to https://cloud.dify.ai/apps?oauth_login=success to create your app，then go into app to get link

![img_9.png](images/img_9.png)
![img_10.png](images/img_10.png)
![img_12.png](images/img_12.png)

- develop your AI chat by yourself and apply to get link, you can refer it: https://github.com/Yidadaa/ChatGPT-Next-Web
![img_11.png](images/img_11.png)



## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=chenshijian73-qq/SwitchQC&type=Date)](https://star-history.com/#chenshijian73-qq/SwitchQC&Date)
