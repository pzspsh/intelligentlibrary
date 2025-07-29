# mitmproxy使用教程
### 1. 安装 Python
    ```
    省略
    ```
### 2. 安装 mitmproxy
    ```
    pip install mitmproxy
    ```
### 3. 启动 mitmproxy
    ```
    mitmproxy
    mitmproxy
    mitmdump # 默认
    mitmweb  # mitmproxy 的 Web UI
    ```

### 4. 先设置系统代理默认端口8080
    ```
    系统设置 -> 网络 -> 代理 -> 手动 -> 代理服务器 -> 输入localhost:8080
    ```

### 5、然后就可以下载mitmproxy的证书
    ```
    http://mitm.it # 下载你对应系统的证书, 并安装到系统中
    ```