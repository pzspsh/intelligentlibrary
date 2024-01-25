# windows 性能提高修改

```
1、 win + X
2、选择 终端管理员
3、输入命令：powercfg -duplicatescheme e9a42b02-d5df-448d-aa00-03f14749eb61
4、回车

5、去控制面板--> 硬件和声音--> 电源选项 -->点击平衡后面的更改计划进行选择
6、
```

# Windows 激活命令

```bash
1、运行管理员终端
2、输入以下任意命令
    irm https://massgrave.dev/get | iex
    irm https://massgrave.dev/get.ps1 | iex
    irm massgrave.dev/get | iex
    irm massgrave.dev/get.ps1 | iex
3、等待激活
4、根据情况输入对应激活数字1、2、3...
这两个命令都可以

```

##### Win11 内存使用率 90%以上如何解决？

```
方法一：
    1、首先，按键盘上的【 Win + X 】组合键，或右键点击任务栏上的【Windows开始徽标】，在打开的右键菜单项中，选择【Windows 终端 （管理员）】
    2、管理员： Windows PowerShell窗口，输入并按回车执行【powercfg /h off】命令。
    3、如果需要开启快速启动，只需要把【off】改为【on】即可。
```
