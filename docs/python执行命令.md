# Python 

### Python执行命令
```bash
python -m pip install --upgrade pip
python -m pip install requests
python -m http.server 8080
python -m http.server port --bind ip
pip install requests -i https://pypi.douban.com/simple --trusted-host pypi.douban.com
pip install --timeout 60 your-package # 增加超时时间
pip config set global.timeout 60 # 配置中设置全局超时
pip install your-package --ignore-requires-python # 后面添加--ignore-requires-python忽略版本
```

### Python下载包源
```bash
https://pypi.org/simple
https://pypi.tuna.tsinghua.edu.cn/simple
https://pypi.mirrors.ustc.edu.cn/simple
https://pypi.douban.com/simple
https://mirrors.aliyun.com/pypi/simple
https://mirrors.cloud.tencent.com/pypi/simple
https://mirrors.huaweicloud.com/repository/pypi/simple
https://mirrors.ustc.edu.cn/pypi/web/simple
```