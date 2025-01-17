# Python 项目建立虚拟环境
```bash
# 切换到项目目录下执行以下操作：
$ python3 -m venv .env  # .env虚拟环境的名称，可以随意更改
$ cd .env/ # 切换到.env
# $ source .env/bin/activate
$ source .env/bin/activate
$ pip freeze

3、关闭虚拟环境
$ deactivate

https://mirrors.aliyun.com/pypi/simple
https://pypi.douban.com/simple

```

```bash
# 安装虚拟环境（如果尚未安装）
pip install virtualenv
# 创建一个新的虚拟环境
virtualenv tf_env
 
# 激活虚拟环境
# 在 Windows 上
tf_env\Scripts\activate
 
# 在 Unix 或 MacOS 上
source tf_env/bin/activate
 
# 在虚拟环境中安装 TensorFlow
pip install tensorflow
```