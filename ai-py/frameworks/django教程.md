# Django 教程

### 1、下载安装 django

```python
pip install django -i https://pypi.douban.com/simple

例如：
# ai-py\framelearn\projects> pip install django -i https://pypi.douban.com/simple
```

### 2、创建 django 项目

放 django 文件的的文件路径上输入 cmd 进入终端输入下面的命令

```python
django-admin startproject 项目名称

例如：
# ai-py\framelearn\projects> django-admin startproject PanView
```

### 3、app 的创建

进入创建好的项目里面输入一下的命令

```python
python manage.py startapp app名称

例如：
# ai-py\framelearn\projects\PanView> python .\manage.py startapp panApp
```

### 4、运行启动 django 项目命令

```python
方式一：
python manage.py runserver

方式二：设置ip端口
python manage.py runserver 0.0.0.0:8080
```
