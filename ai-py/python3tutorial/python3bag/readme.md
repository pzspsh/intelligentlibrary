# Python3 标准库及功能示例

# [Python 标准库](https://docs.python.org/zh-cn/3.10/library/#the-python-standard-library)

[Python 语言参考手册](https://docs.python.org/zh-cn/3.10/reference/index.html#reference-index) 描述了 Python 语言的具体语法和语义，这份库参考则介绍了与 Python 一同发行的标准库。它还描述了通常包含在 Python 发行版中的一些可选组件。

Python 标准库非常庞大，所提供的组件涉及范围十分广泛，正如以下内容目录所显示的。这个库包含了多个内置模块 (以 C 编写)，Python 程序员必须依靠它们来实现系统级功能，例如文件 I/O，此外还有大量以 Python 编写的模块，提供了日常编程中许多问题的标准解决方案。其中有些模块经过专门设计，通过将特定平台功能抽象化为平台中立的 API 来鼓励和加强 Python 程序的可移植性。

Windows 版本的 Python 安装程序通常包含整个标准库，往往还包含许多额外组件。对于类 Unix 操作系统，Python 通常会分成一系列的软件包，因此可能需要使用操作系统所提供的包管理工具来获取部分或全部可选组件。

在这个标准库以外还存在成千上万并且不断增加的其他组件 (从单独的程序、模块、软件包直到完整的应用开发框架)，访问 [Python 包索引](https://pypi.org/) 即可获取这些第三方包。

- 概述
  - [可用性注释](https://docs.python.org/zh-cn/3.10/library/intro.html#notes-on-availability)
- [内置函数](https://docs.python.org/zh-cn/3.10/library/functions.html)

- 内置常量
  - [由 `site` 模块添加的常量](https://docs.python.org/zh-cn/3.10/library/constants.html#constants-added-by-the-site-module)
- 内置类型
  - [逻辑值检测](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#truth-value-testing)
  - [布尔运算 --- `and`, `or`, `not`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#boolean-operations-and-or-not)
  - [比较运算](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#comparisons)
  - [数字类型 --- `int`, `float`, `complex`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#numeric-types-int-float-complex)
  - [迭代器类型](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#iterator-types)
  - [序列类型 --- `list`, `tuple`, `range`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#sequence-types-list-tuple-range)
  - [文本序列类型 --- `str`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#text-sequence-type-str)
  - [二进制序列类型 --- `bytes`, `bytearray`, `memoryview`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#binary-sequence-types-bytes-bytearray-memoryview)
  - [集合类型 --- `set`, `frozenset`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#set-types-set-frozenset)
  - [映射类型 --- `dict`](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#mapping-types-dict)
  - [上下文管理器类型](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#context-manager-types)
  - [类型注解的类型 --- Generic Alias 、 Union](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#type-annotation-types-generic-alias-union)
  - [其他内置类型](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#other-built-in-types)
  - [特殊属性](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#special-attributes)
  - [整数字符串转换长度限制](https://docs.python.org/zh-cn/3.10/library/stdtypes.html#integer-string-conversion-length-limitation)
- 内置异常
  - [异常上下文](https://docs.python.org/zh-cn/3.10/library/exceptions.html#exception-context)
  - [从内置异常继承](https://docs.python.org/zh-cn/3.10/library/exceptions.html#inheriting-from-built-in-exceptions)
  - [基类](https://docs.python.org/zh-cn/3.10/library/exceptions.html#base-classes)
  - [具体异常](https://docs.python.org/zh-cn/3.10/library/exceptions.html#concrete-exceptions)
  - [警告](https://docs.python.org/zh-cn/3.10/library/exceptions.html#warnings)
  - [异常层次结构](https://docs.python.org/zh-cn/3.10/library/exceptions.html#exception-hierarchy)
- 文本处理服务
  - [`string` --- 常见的字符串操作](https://docs.python.org/zh-cn/3.10/library/string.html)
  - [`re` --- 正则表达式操作](https://docs.python.org/zh-cn/3.10/library/re.html)
  - [`difflib` --- 计算差异的辅助工具](https://docs.python.org/zh-cn/3.10/library/difflib.html)
  - [`textwrap` --- 文本自动换行与填充](https://docs.python.org/zh-cn/3.10/library/textwrap.html)
  - [`unicodedata` --- Unicode 数据库](https://docs.python.org/zh-cn/3.10/library/unicodedata.html)
  - [`stringprep` --- 因特网字符串预备](https://docs.python.org/zh-cn/3.10/library/stringprep.html)
  - [`readline` --- GNU readline 接口](https://docs.python.org/zh-cn/3.10/library/readline.html)
  - [`rlcompleter` --- GNU readline 的补全函数](https://docs.python.org/zh-cn/3.10/library/rlcompleter.html)
- 二进制数据服务
  - [`struct` --- 将字节串解读为打包的二进制数据](https://docs.python.org/zh-cn/3.10/library/struct.html)
  - [`codecs` --- 编解码器注册和相关基类](https://docs.python.org/zh-cn/3.10/library/codecs.html)
- 数据类型
  - [`datetime` --- 基本日期和时间类型](https://docs.python.org/zh-cn/3.10/library/datetime.html)
  - [`zoneinfo` --- IANA 时区支持](https://docs.python.org/zh-cn/3.10/library/zoneinfo.html)
  - [`calendar` --- 日历相关函数](https://docs.python.org/zh-cn/3.10/library/calendar.html)
  - [`collections` --- 容器数据类型](https://docs.python.org/zh-cn/3.10/library/collections.html)
  - [`collections.abc` --- 容器的抽象基类](https://docs.python.org/zh-cn/3.10/library/collections.abc.html)
  - [`heapq` --- 堆队列算法](https://docs.python.org/zh-cn/3.10/library/heapq.html)
  - [`bisect` --- 数组二分查找算法](https://docs.python.org/zh-cn/3.10/library/bisect.html)
  - [`array` --- 高效的数字数组](https://docs.python.org/zh-cn/3.10/library/array.html)
  - [`weakref` --- 弱引用](https://docs.python.org/zh-cn/3.10/library/weakref.html)
  - [`types` --- 动态类型创建和内置类型名称](https://docs.python.org/zh-cn/3.10/library/types.html)
  - [`copy` --- 浅层 (shallow) 和深层 (deep) 复制操作](https://docs.python.org/zh-cn/3.10/library/copy.html)
  - [`pprint` --- 数据美化输出](https://docs.python.org/zh-cn/3.10/library/pprint.html)
  - [`reprlib` --- 另一种 `repr()` 实现](https://docs.python.org/zh-cn/3.10/library/reprlib.html)
  - [`enum` --- 对枚举的支持](https://docs.python.org/zh-cn/3.10/library/enum.html)
  - [`graphlib` --- 操作类似图的结构的功能](https://docs.python.org/zh-cn/3.10/library/graphlib.html)
- 数字和数学模块
  - [`numbers` --- 数字的抽象基类](https://docs.python.org/zh-cn/3.10/library/numbers.html)
  - [`math` --- 数学函数](https://docs.python.org/zh-cn/3.10/library/math.html)
  - [`cmath` --- 关于复数的数学函数](https://docs.python.org/zh-cn/3.10/library/cmath.html)
  - [`decimal` --- 十进制定点和浮点运算](https://docs.python.org/zh-cn/3.10/library/decimal.html)
  - [`fractions` --- 分数](https://docs.python.org/zh-cn/3.10/library/fractions.html)
  - [`random` --- 生成伪随机数](https://docs.python.org/zh-cn/3.10/library/random.html)
  - [`statistics` --- 数学统计函数](https://docs.python.org/zh-cn/3.10/library/statistics.html)
- 函数式编程模块
  - [`itertools` --- 为高效循环而创建迭代器的函数](https://docs.python.org/zh-cn/3.10/library/itertools.html)
  - [`functools` --- 高阶函数和可调用对象上的操作](https://docs.python.org/zh-cn/3.10/library/functools.html)
  - [`operator` --- 标准运算符替代函数](https://docs.python.org/zh-cn/3.10/library/operator.html)
- 文件和目录访问
  - [`pathlib` --- 面向对象的文件系统路径](https://docs.python.org/zh-cn/3.10/library/pathlib.html)
  - [`os.path` --- 常用路径操作](https://docs.python.org/zh-cn/3.10/library/os.path.html)
  - [`fileinput` --- 迭代来自多个输入流的行](https://docs.python.org/zh-cn/3.10/library/fileinput.html)
  - [`stat` --- 解析 `stat()` 结果](https://docs.python.org/zh-cn/3.10/library/stat.html)
  - [`filecmp` --- 文件及目录的比较](https://docs.python.org/zh-cn/3.10/library/filecmp.html)
  - [`tempfile` --- 生成临时文件和目录](https://docs.python.org/zh-cn/3.10/library/tempfile.html)
  - [`glob` --- Unix 风格路径名模式扩展](https://docs.python.org/zh-cn/3.10/library/glob.html)
  - [`fnmatch` --- Unix 文件名模式匹配](https://docs.python.org/zh-cn/3.10/library/fnmatch.html)
  - [`linecache` --- 随机读写文本行](https://docs.python.org/zh-cn/3.10/library/linecache.html)
  - [`shutil` --- 高阶文件操作](https://docs.python.org/zh-cn/3.10/library/shutil.html)
- 数据持久化
  - [`pickle` --- Python 对象序列化](https://docs.python.org/zh-cn/3.10/library/pickle.html)
  - [`copyreg` --- 注册配合 `pickle` 模块使用的函数](https://docs.python.org/zh-cn/3.10/library/copyreg.html)
  - [`shelve` --- Python 对象持久化](https://docs.python.org/zh-cn/3.10/library/shelve.html)
  - [`marshal` --- 内部 Python 对象序列化](https://docs.python.org/zh-cn/3.10/library/marshal.html)
  - [`dbm` --- Unix "数据库" 接口](https://docs.python.org/zh-cn/3.10/library/dbm.html)
  - [`sqlite3` --- SQLite 数据库 DB-API 2.0 接口模块](https://docs.python.org/zh-cn/3.10/library/sqlite3.html)
- 数据压缩和存档
  - [`zlib` --- 与 **gzip** 兼容的压缩](https://docs.python.org/zh-cn/3.10/library/zlib.html)
  - [`gzip` --- 对 **gzip** 格式的支持](https://docs.python.org/zh-cn/3.10/library/gzip.html)
  - [`bz2` --- 对 **bzip2** 压缩算法的支持](https://docs.python.org/zh-cn/3.10/library/bz2.html)
  - [`lzma` --- 用 LZMA 算法压缩](https://docs.python.org/zh-cn/3.10/library/lzma.html)
  - [`zipfile` --- 使用 ZIP 存档](https://docs.python.org/zh-cn/3.10/library/zipfile.html)
  - [`tarfile` --- 读写 tar 归档文件](https://docs.python.org/zh-cn/3.10/library/tarfile.html)
- 文件格式
  - [`csv` --- CSV 文件读写](https://docs.python.org/zh-cn/3.10/library/csv.html)
  - [`configparser` --- 配置文件解析器](https://docs.python.org/zh-cn/3.10/library/configparser.html)
  - [`netrc` --- netrc 文件处理](https://docs.python.org/zh-cn/3.10/library/netrc.html)
  - [`plistlib` --- 生成与解析 Apple `.plist` 文件](https://docs.python.org/zh-cn/3.10/library/plistlib.html)
- 加密服务
  - [`hashlib` --- 安全哈希与消息摘要](https://docs.python.org/zh-cn/3.10/library/hashlib.html)
  - [`hmac` --- 基于密钥的消息验证](https://docs.python.org/zh-cn/3.10/library/hmac.html)
  - [`secrets` --- 生成管理密码的安全随机数](https://docs.python.org/zh-cn/3.10/library/secrets.html)
- 通用操作系统服务
  - [`os` --- 多种操作系统接口](https://docs.python.org/zh-cn/3.10/library/os.html)
  - [`io` --- 处理流的核心工具](https://docs.python.org/zh-cn/3.10/library/io.html)
  - [`time` --- 时间的访问和转换](https://docs.python.org/zh-cn/3.10/library/time.html)
  - [`argparse` --- 命令行选项、参数和子命令解析器](https://docs.python.org/zh-cn/3.10/library/argparse.html)
  - [`getopt` --- C 风格的命令行选项解析器](https://docs.python.org/zh-cn/3.10/library/getopt.html)
  - [`logging` --- Python 的日志记录工具](https://docs.python.org/zh-cn/3.10/library/logging.html)
  - [`logging.config` --- 日志记录配置](https://docs.python.org/zh-cn/3.10/library/logging.config.html)
  - [`logging.handlers` --- 日志处理程序](https://docs.python.org/zh-cn/3.10/library/logging.handlers.html)
  - [`getpass` --- 便携式密码输入工具](https://docs.python.org/zh-cn/3.10/library/getpass.html)
  - [`curses` --- 终端字符单元显示的处理](https://docs.python.org/zh-cn/3.10/library/curses.html)
  - [`curses.textpad` --- 用于 curses 程序的文本输入控件](https://docs.python.org/zh-cn/3.10/library/curses.html#module-curses.textpad)
  - [`curses.ascii` --- 用于 ASCII 字符的工具](https://docs.python.org/zh-cn/3.10/library/curses.ascii.html)
  - [`curses.panel` --- curses 的面板栈扩展](https://docs.python.org/zh-cn/3.10/library/curses.panel.html)
  - [`platform` --- 获取底层平台的标识数据](https://docs.python.org/zh-cn/3.10/library/platform.html)
  - [`errno` --- 标准 errno 系统符号](https://docs.python.org/zh-cn/3.10/library/errno.html)
  - [`ctypes` --- Python 的外部函数库](https://docs.python.org/zh-cn/3.10/library/ctypes.html)
- 并发执行
  - [`threading` --- 基于线程的并行](https://docs.python.org/zh-cn/3.10/library/threading.html)
  - [`multiprocessing` --- 基于进程的并行](https://docs.python.org/zh-cn/3.10/library/multiprocessing.html)
  - [`multiprocessing.shared_memory` --- 可跨进程直接访问的共享内存](https://docs.python.org/zh-cn/3.10/library/multiprocessing.shared_memory.html)
  - [`concurrent` 包](https://docs.python.org/zh-cn/3.10/library/concurrent.html)
  - [`concurrent.futures` --- 启动并行任务](https://docs.python.org/zh-cn/3.10/library/concurrent.futures.html)
  - [`subprocess` --- 子进程管理](https://docs.python.org/zh-cn/3.10/library/subprocess.html)
  - [`sched` --- 事件调度器](https://docs.python.org/zh-cn/3.10/library/sched.html)
  - [`queue` --- 一个同步的队列类](https://docs.python.org/zh-cn/3.10/library/queue.html)
  - [`contextvars` --- 上下文变量](https://docs.python.org/zh-cn/3.10/library/contextvars.html)
  - [`_thread` --- 底层多线程 API](https://docs.python.org/zh-cn/3.10/library/_thread.html)
- 网络和进程间通信
  - [`asyncio` --- 异步 I/O](https://docs.python.org/zh-cn/3.10/library/asyncio.html)
  - [`socket` --- 底层网络接口](https://docs.python.org/zh-cn/3.10/library/socket.html)
  - [`ssl` --- 套接字对象的 TLS/SSL 包装器](https://docs.python.org/zh-cn/3.10/library/ssl.html)
  - [`select` --- 等待 I/O 完成](https://docs.python.org/zh-cn/3.10/library/select.html)
  - [`selectors` --- 高级 I/O 复用库](https://docs.python.org/zh-cn/3.10/library/selectors.html)
  - [`signal` --- 设置异步事件处理程序](https://docs.python.org/zh-cn/3.10/library/signal.html)
  - [`mmap` --- 内存映射文件支持](https://docs.python.org/zh-cn/3.10/library/mmap.html)
- 互联网数据处理
  - [`email` --- 电子邮件与 MIME 处理包](https://docs.python.org/zh-cn/3.10/library/email.html)
  - [`json` --- JSON 编码和解码器](https://docs.python.org/zh-cn/3.10/library/json.html)
  - [`mailbox` --- 操作多种格式的邮箱](https://docs.python.org/zh-cn/3.10/library/mailbox.html)
  - [`mimetypes` --- 映射文件名到 MIME 类型](https://docs.python.org/zh-cn/3.10/library/mimetypes.html)
  - [`base64` --- Base16, Base32, Base64, Base85 数据编码](https://docs.python.org/zh-cn/3.10/library/base64.html)
  - [`binhex` --- 对 binhex4 文件进行编码和解码](https://docs.python.org/zh-cn/3.10/library/binhex.html)
  - [`binascii` --- 二进制和 ASCII 码互转](https://docs.python.org/zh-cn/3.10/library/binascii.html)
  - [`quopri` --- 编码与解码经过 MIME 转码的可打印数据](https://docs.python.org/zh-cn/3.10/library/quopri.html)
- 结构化标记处理工具
  - [`html` --- 超文本标记语言支持](https://docs.python.org/zh-cn/3.10/library/html.html)
  - [`html.parser` --- 简单的 HTML 和 XHTML 解析器](https://docs.python.org/zh-cn/3.10/library/html.parser.html)
  - [`html.entities` --- HTML 一般实体的定义](https://docs.python.org/zh-cn/3.10/library/html.entities.html)
  - [XML 处理模块](https://docs.python.org/zh-cn/3.10/library/xml.html)
  - [`xml.etree.ElementTree` --- ElementTree XML API](https://docs.python.org/zh-cn/3.10/library/xml.etree.elementtree.html)
  - [`xml.dom` --- 文档对象模型 API](https://docs.python.org/zh-cn/3.10/library/xml.dom.html)
  - [`xml.dom.minidom` --- 最小化的 DOM 实现](https://docs.python.org/zh-cn/3.10/library/xml.dom.minidom.html)
  - [`xml.dom.pulldom` --- 支持构建部分 DOM 树](https://docs.python.org/zh-cn/3.10/library/xml.dom.pulldom.html)
  - [`xml.sax` --- 支持 SAX2 解析器](https://docs.python.org/zh-cn/3.10/library/xml.sax.html)
  - [`xml.sax.handler` --- SAX 处理句柄的基类](https://docs.python.org/zh-cn/3.10/library/xml.sax.handler.html)
  - [`xml.sax.saxutils` --- SAX 工具集](https://docs.python.org/zh-cn/3.10/library/xml.sax.utils.html)
  - [`xml.sax.xmlreader` --- 用于 XML 解析器的接口](https://docs.python.org/zh-cn/3.10/library/xml.sax.reader.html)
  - [`xml.parsers.expat` --- 使用 Expat 的快速 XML 解析](https://docs.python.org/zh-cn/3.10/library/pyexpat.html)
- 互联网协议和支持
  - [`webbrowser` --- 方便的 Web 浏览器控制工具](https://docs.python.org/zh-cn/3.10/library/webbrowser.html)
  - [`wsgiref` --- WSGI 工具和参考实现](https://docs.python.org/zh-cn/3.10/library/wsgiref.html)
  - [`urllib` --- URL 处理模块](https://docs.python.org/zh-cn/3.10/library/urllib.html)
  - [`urllib.request` --- 用于打开 URL 的可扩展库](https://docs.python.org/zh-cn/3.10/library/urllib.request.html)
  - [`urllib.response` --- urllib 使用的 Response 类](https://docs.python.org/zh-cn/3.10/library/urllib.request.html#module-urllib.response)
  - [`urllib.parse` 用于解析 URL](https://docs.python.org/zh-cn/3.10/library/urllib.parse.html)
  - [`urllib.error` --- urllib.request 引发的异常类](https://docs.python.org/zh-cn/3.10/library/urllib.error.html)
  - [`urllib.robotparser` --- robots.txt 语法分析程序](https://docs.python.org/zh-cn/3.10/library/urllib.robotparser.html)
  - [`http` --- HTTP 模块](https://docs.python.org/zh-cn/3.10/library/http.html)
  - [`http.client` --- HTTP 协议客户端](https://docs.python.org/zh-cn/3.10/library/http.client.html)
  - [`ftplib` --- FTP 协议客户端](https://docs.python.org/zh-cn/3.10/library/ftplib.html)
  - [`poplib` --- POP3 协议客户端](https://docs.python.org/zh-cn/3.10/library/poplib.html)
  - [`imaplib` --- IMAP4 协议客户端](https://docs.python.org/zh-cn/3.10/library/imaplib.html)
  - [`smtplib` --- SMTP 协议客户端](https://docs.python.org/zh-cn/3.10/library/smtplib.html)
  - [`uuid` --- **RFC 4122** 定义的 UUID 对象](https://docs.python.org/zh-cn/3.10/library/uuid.html)
  - [`socketserver` --- 用于网络服务器的框架](https://docs.python.org/zh-cn/3.10/library/socketserver.html)
  - [`http.server` --- HTTP 服务器](https://docs.python.org/zh-cn/3.10/library/http.server.html)
  - [`http.cookies` --- HTTP 状态管理](https://docs.python.org/zh-cn/3.10/library/http.cookies.html)
  - [`http.cookiejar` —— HTTP 客户端的 Cookie 处理](https://docs.python.org/zh-cn/3.10/library/http.cookiejar.html)
  - [`xmlrpc` --- XMLRPC 服务端与客户端模块](https://docs.python.org/zh-cn/3.10/library/xmlrpc.html)
  - [`xmlrpc.client` --- XML-RPC 客户端访问](https://docs.python.org/zh-cn/3.10/library/xmlrpc.client.html)
  - [`xmlrpc.server` --- 基本 XML-RPC 服务器](https://docs.python.org/zh-cn/3.10/library/xmlrpc.server.html)
  - [`ipaddress` --- IPv4/IPv6 操作库](https://docs.python.org/zh-cn/3.10/library/ipaddress.html)
- 多媒体服务
  - [`wave` --- 读写 WAV 格式文件](https://docs.python.org/zh-cn/3.10/library/wave.html)
  - [`colorsys` --- 颜色系统间的转换](https://docs.python.org/zh-cn/3.10/library/colorsys.html)
- 国际化
  - [`gettext` --- 多语种国际化服务](https://docs.python.org/zh-cn/3.10/library/gettext.html)
  - [`locale` --- 国际化服务](https://docs.python.org/zh-cn/3.10/library/locale.html)
- 程序框架
  - [`turtle` --- 海龟绘图](https://docs.python.org/zh-cn/3.10/library/turtle.html)
  - [`cmd` --- 支持面向行的命令解释器](https://docs.python.org/zh-cn/3.10/library/cmd.html)
  - [`shlex` —— 简单的词法分析](https://docs.python.org/zh-cn/3.10/library/shlex.html)
- Tk 图形用户界面(GUI)
  - [`tkinter` —— Tcl/Tk 的 Python 接口](https://docs.python.org/zh-cn/3.10/library/tkinter.html)
  - [`tkinter.colorchooser` --- 颜色选择对话框](https://docs.python.org/zh-cn/3.10/library/tkinter.colorchooser.html)
  - [`tkinter.font` --- Tkinter 字体封装](https://docs.python.org/zh-cn/3.10/library/tkinter.font.html)
  - [Tkinter 对话框](https://docs.python.org/zh-cn/3.10/library/dialog.html)
  - [`tkinter.messagebox` --- Tkinter 消息提示](https://docs.python.org/zh-cn/3.10/library/tkinter.messagebox.html)
  - [`tkinter.scrolledtext` --- 滚动文字控件](https://docs.python.org/zh-cn/3.10/library/tkinter.scrolledtext.html)
  - [`tkinter.dnd` --- 拖放操作支持](https://docs.python.org/zh-cn/3.10/library/tkinter.dnd.html)
  - [`tkinter.ttk` --- Tk 风格的控件](https://docs.python.org/zh-cn/3.10/library/tkinter.ttk.html)
  - [`tkinter.tix` --- TK 扩展包](https://docs.python.org/zh-cn/3.10/library/tkinter.tix.html)
  - [IDLE](https://docs.python.org/zh-cn/3.10/library/idle.html)
- 开发工具
  - [`typing` —— 对类型提示的支持](https://docs.python.org/zh-cn/3.10/library/typing.html)
  - [`pydoc` --- 文档生成器和在线帮助系统](https://docs.python.org/zh-cn/3.10/library/pydoc.html)
  - [Python 开发模式](https://docs.python.org/zh-cn/3.10/library/devmode.html)
  - [Python 开发模式的效果](https://docs.python.org/zh-cn/3.10/library/devmode.html#effects-of-the-python-development-mode)
  - [ResourceWarning 示例](https://docs.python.org/zh-cn/3.10/library/devmode.html#resourcewarning-example)
  - [文件描述符错误示例](https://docs.python.org/zh-cn/3.10/library/devmode.html#bad-file-descriptor-error-example)
  - [`doctest` --- 测试交互性的 Python 示例](https://docs.python.org/zh-cn/3.10/library/doctest.html)
  - [`unittest` --- 单元测试框架](https://docs.python.org/zh-cn/3.10/library/unittest.html)
  - [`unittest.mock` --- 模拟对象库](https://docs.python.org/zh-cn/3.10/library/unittest.mock.html)
  - [`unittest.mock` --- 上手指南](https://docs.python.org/zh-cn/3.10/library/unittest.mock-examples.html)
  - [2to3 --- 自动化的 Python 2 到 3 代码转写](https://docs.python.org/zh-cn/3.10/library/2to3.html)
  - [`test` --- Python 回归测试包](https://docs.python.org/zh-cn/3.10/library/test.html)
  - [`test.support` --- 针对 Python 测试套件的工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support)
  - [`test.support.socket_helper` --- 用于套接字测试的工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.socket_helper)
  - [`test.support.script_helper` --- 用于 Python 执行测试工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.script_helper)
  - [`test.support.bytecode_helper` --- 用于测试正确字节码生成的支持工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.bytecode_helper)
  - [`test.support.threading_helper` --- 用于线程测试的工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.threading_helper)
  - [`test.support.os_helper` --- 用于操作系统测试的工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.os_helper)
  - [`test.support.import_helper` --- 用于导入测试的工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.import_helper)
  - [`test.support.warnings_helper` --- 用于警告测试的工具](https://docs.python.org/zh-cn/3.10/library/test.html#module-test.support.warnings_helper)
- 调试和分析
  - [审计事件表](https://docs.python.org/zh-cn/3.10/library/audit_events.html)
  - [`bdb` --- 调试器框架](https://docs.python.org/zh-cn/3.10/library/bdb.html)
  - [`faulthandler` —— 转储 Python 的跟踪信息](https://docs.python.org/zh-cn/3.10/library/faulthandler.html)
  - [`pdb` --- Python 的调试器](https://docs.python.org/zh-cn/3.10/library/pdb.html)
  - [Python 性能分析器](https://docs.python.org/zh-cn/3.10/library/profile.html)
  - [`timeit` --- 测量小代码片段的执行时间](https://docs.python.org/zh-cn/3.10/library/timeit.html)
  - [`trace` —— 跟踪 Python 语句的执行](https://docs.python.org/zh-cn/3.10/library/trace.html)
  - [`tracemalloc` --- 跟踪内存分配](https://docs.python.org/zh-cn/3.10/library/tracemalloc.html)
- 软件打包和分发
  - [`distutils` --- 构建和安装 Python 模块](https://docs.python.org/zh-cn/3.10/library/distutils.html)
  - [`ensurepip` --- 引导 `pip` 安装器](https://docs.python.org/zh-cn/3.10/library/ensurepip.html)
  - [`venv` --- 创建虚拟环境](https://docs.python.org/zh-cn/3.10/library/venv.html)
  - [`zipapp` —— 管理可执行的 Python zip 打包文件](https://docs.python.org/zh-cn/3.10/library/zipapp.html)
- Python 运行时服务
  - [`sys` --- 系统相关的形参和函数](https://docs.python.org/zh-cn/3.10/library/sys.html)
  - [`sysconfig` —— 提供对 Python 配置信息的访问支持](https://docs.python.org/zh-cn/3.10/library/sysconfig.html)
  - [`builtins` --- 内建对象](https://docs.python.org/zh-cn/3.10/library/builtins.html)
  - [`__main__` --- 最高层级代码环境](https://docs.python.org/zh-cn/3.10/library/__main__.html)
  - [`warnings` —— 警告信息的控制](https://docs.python.org/zh-cn/3.10/library/warnings.html)
  - [`dataclasses` --- 数据类](https://docs.python.org/zh-cn/3.10/library/dataclasses.html)
  - [`contextlib` --- 为 `with`语句上下文提供的工具](https://docs.python.org/zh-cn/3.10/library/contextlib.html)
  - [`abc` --- 抽象基类](https://docs.python.org/zh-cn/3.10/library/abc.html)
  - [`atexit` --- 退出处理器](https://docs.python.org/zh-cn/3.10/library/atexit.html)
  - [`traceback` —— 打印或读取堆栈的跟踪信息](https://docs.python.org/zh-cn/3.10/library/traceback.html)
  - [`__future__` --- Future 语句定义](https://docs.python.org/zh-cn/3.10/library/__future__.html)
  - [`gc` --- 垃圾回收器接口](https://docs.python.org/zh-cn/3.10/library/gc.html)
  - [`inspect` --- 检查对象](https://docs.python.org/zh-cn/3.10/library/inspect.html)
  - [`site` —— 指定域的配置钩子](https://docs.python.org/zh-cn/3.10/library/site.html)
- 自定义 Python 解释器
  - [`code` --- 解释器基类](https://docs.python.org/zh-cn/3.10/library/code.html)
  - [`codeop` --- 编译 Python 代码](https://docs.python.org/zh-cn/3.10/library/codeop.html)
- 导入模块
  - [`zipimport` --- 从 Zip 存档中导入模块](https://docs.python.org/zh-cn/3.10/library/zipimport.html)
  - [`pkgutil` --- 包扩展工具](https://docs.python.org/zh-cn/3.10/library/pkgutil.html)
  - [`modulefinder` --- 查找脚本使用的模块](https://docs.python.org/zh-cn/3.10/library/modulefinder.html)
  - [`runpy` ——查找并执行 Python 模块](https://docs.python.org/zh-cn/3.10/library/runpy.html)
  - [`importlib` --- `import` 的实现](https://docs.python.org/zh-cn/3.10/library/importlib.html)
  - [使用 `importlib.metadata`](https://docs.python.org/zh-cn/3.10/library/importlib.metadata.html)
- Python 语言服务
  - [`ast` --- 抽象语法树](https://docs.python.org/zh-cn/3.10/library/ast.html)
  - [`symtable` ——访问编译器的符号表](https://docs.python.org/zh-cn/3.10/library/symtable.html)
  - [`token` --- 与 Python 解析树一起使用的常量](https://docs.python.org/zh-cn/3.10/library/token.html)
  - [`keyword` --- 检验 Python 关键字](https://docs.python.org/zh-cn/3.10/library/keyword.html)
  - [`tokenize` --- 对 Python 代码使用的标记解析器](https://docs.python.org/zh-cn/3.10/library/tokenize.html)
  - [`tabnanny` --- 模糊缩进检测](https://docs.python.org/zh-cn/3.10/library/tabnanny.html)
  - [`pyclbr` --- Python 模块浏览器支持](https://docs.python.org/zh-cn/3.10/library/pyclbr.html)
  - [`py_compile` --- 编译 Python 源文件](https://docs.python.org/zh-cn/3.10/library/py_compile.html)
  - [`compileall` --- 字节编译 Python 库](https://docs.python.org/zh-cn/3.10/library/compileall.html)
  - [`dis` --- Python 字节码反汇编器](https://docs.python.org/zh-cn/3.10/library/dis.html)
  - [`pickletools` --- pickle 开发者工具集](https://docs.python.org/zh-cn/3.10/library/pickletools.html)
- Windows 系统相关模块
  - [`msvcrt` --- 来自 MS VC++ 运行时的有用例程](https://docs.python.org/zh-cn/3.10/library/msvcrt.html)
  - [`winreg` --- 访问 Windows 注册表](https://docs.python.org/zh-cn/3.10/library/winreg.html)
  - [`winsound` —— Windows 系统的音频播放接口](https://docs.python.org/zh-cn/3.10/library/winsound.html)
- Unix 专有服务
  - [`posix` --- 最常见的 POSIX 系统调用](https://docs.python.org/zh-cn/3.10/library/posix.html)
  - [`pwd` --- 用户密码数据库](https://docs.python.org/zh-cn/3.10/library/pwd.html)
  - [`grp` --- 组数据库](https://docs.python.org/zh-cn/3.10/library/grp.html)
  - [`termios` --- POSIX 风格的 tty 控制](https://docs.python.org/zh-cn/3.10/library/termios.html)
  - [`tty` --- 终端控制功能](https://docs.python.org/zh-cn/3.10/library/tty.html)
  - [`pty` --- 伪终端工具](https://docs.python.org/zh-cn/3.10/library/pty.html)
  - [`fcntl` —— 系统调用 `fcntl` 和 `ioctl`](https://docs.python.org/zh-cn/3.10/library/fcntl.html)
  - [`resource` --- 资源使用信息](https://docs.python.org/zh-cn/3.10/library/resource.html)
  - [Unix syslog 库例程](https://docs.python.org/zh-cn/3.10/library/syslog.html)
- 被取代的模块
  - [`aifc` --- 读写 AIFF 和 AIFC 文件](https://docs.python.org/zh-cn/3.10/library/aifc.html)
  - [`asynchat` --- 异步套接字指令/响应处理程序](https://docs.python.org/zh-cn/3.10/library/asynchat.html)
  - [`asyncore` --- 异步套接字处理器](https://docs.python.org/zh-cn/3.10/library/asyncore.html)
  - [`audioop` --- 处理原始音频数据](https://docs.python.org/zh-cn/3.10/library/audioop.html)
  - [`cgi` --- 通用网关接口支持](https://docs.python.org/zh-cn/3.10/library/cgi.html)
  - [`cgitb` --- 用于 CGI 脚本的回溯管理器](https://docs.python.org/zh-cn/3.10/library/cgitb.html)
  - [`chunk` --- 读取 IFF 分块数据](https://docs.python.org/zh-cn/3.10/library/chunk.html)
  - [`crypt` —— 验证 Unix 口令的函数](https://docs.python.org/zh-cn/3.10/library/crypt.html)
  - [`imghdr` --- 推测图像类型](https://docs.python.org/zh-cn/3.10/library/imghdr.html)
  - [`imp` —— 由代码内部访问 import 。](https://docs.python.org/zh-cn/3.10/library/imp.html)
  - [`mailcap` --- Mailcap 文件处理](https://docs.python.org/zh-cn/3.10/library/mailcap.html)
  - [`msilib` --- 读写 Microsoft Installer 文件](https://docs.python.org/zh-cn/3.10/library/msilib.html)
  - [`nis` --- Sun 的 NIS (黄页) 接口](https://docs.python.org/zh-cn/3.10/library/nis.html)
  - [`nntplib` --- NNTP 协议客户端](https://docs.python.org/zh-cn/3.10/library/nntplib.html)
  - [`optparse` --- 命令行选项的解析器](https://docs.python.org/zh-cn/3.10/library/optparse.html)
  - [`ossaudiodev` --- 访问兼容 OSS 的音频设备](https://docs.python.org/zh-cn/3.10/library/ossaudiodev.html)
  - [`pipes` --- 终端管道接口](https://docs.python.org/zh-cn/3.10/library/pipes.html)
  - [`smtpd` --- SMTP 服务器](https://docs.python.org/zh-cn/3.10/library/smtpd.html)
  - [`sndhdr` --- 推测声音文件的类型](https://docs.python.org/zh-cn/3.10/library/sndhdr.html)
  - [`spwd` —— shadow 密码库](https://docs.python.org/zh-cn/3.10/library/spwd.html)
  - [`sunau` --- 读写 Sun AU 文件](https://docs.python.org/zh-cn/3.10/library/sunau.html)
  - [`telnetlib` -- Telnet 客户端](https://docs.python.org/zh-cn/3.10/library/telnetlib.html)
  - [`uu` --- 对 uuencode 文件进行编码与解码](https://docs.python.org/zh-cn/3.10/library/uu.html)
  - [`xdrlib` --- 编码与解码 XDR 数据](https://docs.python.org/zh-cn/3.10/library/xdrlib.html)
- [安全考量](https://docs.python.org/zh-cn/3.10/library/security_warnings.html)

  下列模块具有专门的安全事项:

  - [`base64`](https://docs.python.org/zh-cn/3.10/library/base64.html#module-base64): [base64 安全事项](https://docs.python.org/zh-cn/3.10/library/base64.html#base64-security)，参见 [**RFC 4648**](https://datatracker.ietf.org/doc/html/rfc4648.html)
  - [`cgi`](https://docs.python.org/zh-cn/3.10/library/cgi.html#module-cgi): [CGI 安全事项](https://docs.python.org/zh-cn/3.10/library/cgi.html#cgi-security)
  - [`hashlib`](https://docs.python.org/zh-cn/3.10/library/hashlib.html#module-hashlib): [所有构造器都接受一个 "usedforsecurity" 仅限关键字参数以停用已知的不安全和已封禁的算法](https://docs.python.org/zh-cn/3.10/library/hashlib.html#hashlib-usedforsecurity)
  - [`http.server`](https://docs.python.org/zh-cn/3.10/library/http.server.html#module-http.server) 不适合生产用途，只实现了基本的安全检查。 请参阅 [安全性考量](https://docs.python.org/zh-cn/3.10/library/http.server.html#http-server-security)。
  - [`logging`](https://docs.python.org/zh-cn/3.10/library/logging.html#module-logging): [日志记录配置使用了 eval()](https://docs.python.org/zh-cn/3.10/library/logging.config.html#logging-eval-security)
  - [`multiprocessing`](https://docs.python.org/zh-cn/3.10/library/multiprocessing.html#module-multiprocessing): [Connection.recv() 使用了 pickle](https://docs.python.org/zh-cn/3.10/library/multiprocessing.html#multiprocessing-recv-pickle-security)
  - [`pickle`](https://docs.python.org/zh-cn/3.10/library/pickle.html#module-pickle): [在 pickle 中限制全局变量](https://docs.python.org/zh-cn/3.10/library/pickle.html#pickle-restrict)
  - [`random`](https://docs.python.org/zh-cn/3.10/library/random.html#module-random) 不应当被用于安全目的，而应改用 [`secrets`](https://docs.python.org/zh-cn/3.10/library/secrets.html#module-secrets)
  - [`shelve`](https://docs.python.org/zh-cn/3.10/library/shelve.html#module-shelve): [shelve 是基于 pickle 的因此不适用于处理不受信任的源](https://docs.python.org/zh-cn/3.10/library/shelve.html#shelve-security)
  - [`ssl`](https://docs.python.org/zh-cn/3.10/library/ssl.html#module-ssl): [SSL/TLS 安全事项](https://docs.python.org/zh-cn/3.10/library/ssl.html#ssl-security)
  - [`subprocess`](https://docs.python.org/zh-cn/3.10/library/subprocess.html#module-subprocess): [子进程安全事项](https://docs.python.org/zh-cn/3.10/library/subprocess.html#subprocess-security)
  - [`tempfile`](https://docs.python.org/zh-cn/3.10/library/tempfile.html#module-tempfile): [mktemp 由于存在竞争条件缺陷已被弃用](https://docs.python.org/zh-cn/3.10/library/tempfile.html#tempfile-mktemp-deprecated)
  - [`xml`](https://docs.python.org/zh-cn/3.10/library/xml.html#module-xml): [XML 安全缺陷](https://docs.python.org/zh-cn/3.10/library/xml.html#xml-vulnerabilities)
  - [`zipfile`](https://docs.python.org/zh-cn/3.10/library/zipfile.html#module-zipfile): [恶意处理的 .zip 文件可能导致硬盘空间耗尽](https://docs.python.org/zh-cn/3.10/library/zipfile.html#zipfile-resources-limitations)

[Python3 中文官网](https://docs.python.org/zh-cn/3.6/library/)
[Python3 官方教程](https://docs.python.org/3/tutorial/stdlib2.html)
