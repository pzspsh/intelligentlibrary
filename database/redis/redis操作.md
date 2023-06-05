# Redis操作
```sql
将字符串值 value 关联到 key,如果 key 已经持有其他值， SET 就覆写旧值， 无视类型。当 SET 命令对一个带有生存时间（TTL）的键进行设置之后， 该键原有的 TTL 将被清除。
SET key value [EX seconds] [PX milliseconds] [NX|XX]
```
### Redis SET命令
```sql
对不存在的键进行设置：
redis> SET key "value"
OK
redis> GET key
"value"

对已存在的键进行设置：
redis> SET key "new-value"
OK
redis> GET key
"new-value"

使用 EX 选项：
redis> SET key-with-expire-time "hello" EX 10086
OK
redis> GET key-with-expire-time
"hello"
redis> TTL key-with-expire-time
(integer) 10069

使用 PX 选项：
redis> SET key-with-pexpire-time "moto" PX 123321
OK
redis> GET key-with-pexpire-time
"moto"
redis> PTTL key-with-pexpire-time
(integer) 111939

使用 NX 选项：
redis> SET not-exists-key "value" NX
OK      # 键不存在，设置成功
redis> GET not-exists-key
"value"
redis> SET not-exists-key "new-value" NX
(nil)   # 键已经存在，设置失败
redis> GEt not-exists-key
"value" # 维持原值不变

使用 XX 选项：
redis> EXISTS exists-key
(integer) 0
redis> SET exists-key "value" XX
(nil)   # 因为键不存在，设置失败
redis> SET exists-key "value"
OK      # 先给键设置一个值
redis> SET exists-key "new-value" XX
OK      # 设置新值成功
redis> GET exists-key
"new-value"
```
### Redis SETNX命令
命令在设置成功时返回 1 ， 设置失败时返回 0 。
```sql
SETNX key value
只在键 key 不存在的情况下， 将键 key 的值设置为 value 。
若键 key 已经存在， 则 SETNX 命令不做任何动作。

实例操作：
redis> EXISTS job                # job 不存在
(integer) 0
redis> SETNX job "programmer"    # job 设置成功
(integer) 1
redis> SETNX job "code-farmer"   # 尝试覆盖 job ，失败
(integer) 0
redis> GET job                   # 没有被覆盖
"programmer"
```
### Redis SETEX命令
```sql
SETEX key seconds value
将键 key 的值设置为 value ， 并将键 key 的生存时间设置为 seconds 秒钟。
如果键 key 已经存在， 那么 SETEX 命令将覆盖已有的值。

SETEX 命令的效果和以下两个命令的效果类似：
SET key value
EXPIRE key seconds  # 设置生存时间

实例操作：
在键 key 不存在的情况下执行 SETEX ：
redis> SETEX cache_user_id 60 10086
OK
redis> GET cache_user_id  # 值
"10086"
redis> TTL cache_user_id  # 剩余生存时间
(integer) 49

键 key 已经存在， 使用 SETEX 覆盖旧值：

redis> SET cd "timeless"
OK
redis> SETEX cd 3000 "goodbye my love"
OK
redis> GET cd
"goodbye my love"
redis> TTL cd
(integer) 2997
```

### Redis PSETEX命令
命令在设置成功时返回 OK 。
```sql
PSETEX key milliseconds value
这个命令和 SETEX 命令相似， 但它以毫秒为单位设置 key 的生存时间， 而不是像 SETEX 命令那样以秒为单位进行设置。

实例操作：
redis> PSETEX mykey 1000 "Hello"
OK
redis> PTTL mykey
(integer) 999
redis> GET mykey
"Hello"
```

### Redis GET命令
```sql
GET key
返回与键 key 相关联的字符串值。
如果键 key 不存在， 那么返回特殊值 nil ； 否则， 返回键 key 的值。
如果键 key 的值并非字符串类型， 那么返回一个错误， 因为 GET 命令只能用于字符串值。

实例操作：
对不存在的键 key 或是字符串类型的键 key 执行 GET 命令：
redis> GET db
(nil)
redis> SET db redis
OK
redis> GET db
"redis"

对不是字符串类型的键 key 执行 GET 命令：
redis> DEL db
(integer) 1
redis> LPUSH db redis mongodb mysql
(integer) 3
redis> GET db
(error) ERR Operation against a key holding the wrong kind of value
```

### Redis GETSET命令
```sql
GETSET key value
将键 key 的值设为 value ， 并返回键 key 在被设置之前的旧值。

返回给定键 key 的旧值。
如果键 key 没有旧值， 也即是说， 键 key 在被设置之前并不存在， 那么命令返回 nil 。
当键 key 存在但不是字符串类型时， 命令返回一个错误。

实例操作：
redis> GETSET db mongodb    # 没有旧值，返回 nil
(nil)
redis> GET db
"mongodb"
redis> GETSET db redis      # 返回旧值 mongodb
"mongodb"
redis> GET db
"redis"
```

