## VS Code免密码远程登录Linux

#### 1、打开安装VS Code软件

​		省略

#### 2、打开安装好的VS Code软件，在扩展商店搜索remote如下图所示：

![image-20240821165206108](F:\Images\image-20240821165206108.png)

安装上图所示的扩展Remote-SSH或Remote Development

#### 3、在Windows上进入终端执行ssh-keygen -t rsa命令如下图所示：

![image-20240821163943388](F:\Images\image-20240821163943388.png)

执行完生成id_rsa和id_rsa.pub两个文件



#### 4、**在Linux终端上**，默认在root目录下有文件夹.ssh，如果没有可以动手创建

```bash
mkdir ~/.ssh    # 创建.ssh文件夹
touch ~/.ssh/authorized_key # 创建存储公钥的文件
ls ~/.ssh/
```

![image-20240821164446564](F:\Images\image-20240821164446564.png)

![image-20240821170208416](F:\Images\image-20240821170208416.png)

#### 5、再回到windows上找到第三步id_rsa.pub文件并打开文件，然后复制文件内容

![image-20240821170626894](F:\Images\image-20240821170626894.png)



#### 6、回到Linux打开第四步生的authorized_keys文件，并把复制的id_rsa.pub内容编辑到authorized_keys文件上，如下图所示：

![image-20240821171312078](F:\Images\image-20240821171312078.png)

![image-20240821171245616](F:\Images\image-20240821171245616.png)

#### 7、最后重启VS Code重新进入Linux就可以免密码登录了