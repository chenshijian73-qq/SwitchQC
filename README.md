# README

## About

This is the official Wails Vanilla template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Introduce
基于 Wails2 实现类似 switchhost 的功能。
1、左边栏显示文件列表，文件来源设置的本地文件目录，文件后缀为 .quickcmd;
2、文件列表每个文件增加一个开关，可以开启或关闭
3、右边栏为文本框，默认显示左边栏文件列表第一个文件的内容，可以手动点击任意文件，文本框显示对应文件内容。文本框支持增删改。
4、增加按钮，可以点击创建文件、删除文件

前端实现： Vue3
后端实现： Golang