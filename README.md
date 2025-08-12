# yumiShell

##### 本项目是一个基于ssh的一个模拟终端程序，基于ssh进行远程连接。支持文件的拖拽上传和密码以及私钥登录远程主机的方式。界面主要用electron+vue进行开发，go语言主要处理ssh连接部分。

#### 1.使用

打开可执行程序后，点击Add new group建立一个新的组用于管理会话，然后在你的组上新建ssh会话，双击你新建的会话即可连接到远程主机
![image](https://github.com/yumilengjiao/yumiShell/blob/master/assets/addSession.gif?raw=true)


当然你也可以通过鼠标拖拽的方式上传你的文件(不支持拖拽文件夹),上传暂不支持上传大文件，文件大小最好不超过1GB
![image](https://github.com/yumilengjiao/yumiShell/blob/master/assets/fileUpload.gif?raw=true)


你也可以在设置界面搭配你自己喜欢的主题
![image](https://github.com/yumilengjiao/yumiShell/blob/master/assets/setting.png?raw=true)

#### 2.调试

如果你想调试源码，可以通过以下流程启动程序

1. 克隆仓库,将源码拉取到本地

```git clone https://github.com/yumilengjiao/yumiShell.git```

2. 启动后端代码，用你喜欢的编辑器或者直接在终端中打开**yumi-shell\src\backend**目录下的所有文件，找到**main**目录下的main.go文件，使用```go run build main.go```编译并启动，或者使用你的编辑器启动main.go文件
3. 启动前端代码，打开**yumi-shell**目录，终端中输入```pnpm i```安装依赖包，安装后输入```pnpm dev```启动
