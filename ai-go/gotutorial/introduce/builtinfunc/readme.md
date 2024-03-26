### 内置类型

```bash
bool
int(32 or 64), int8, int16, int32, int64
uint(32 or 64), uint8(byte), uint16, uint32, uint64
float32, float64
string
complex64, complex128
array    -- 固定长度的数组
```

### 引用类型

```bash
slice   -- 序列数组(最常用)
map     -- 映射
chan    -- 管道
```

### 内置函数

```bash
append  		-- 用来追加元素到数组、slice中,返回修改后的数组、slice
close   		-- 主要用来关闭channel
delete    		-- 从map中删除key对应的value
panic    		-- 停止常规的goroutine  （panic和recover：用来做错误处理）
recover 		-- 允许程序定义goroutine的panic动作
imag    		-- 返回complex的实部   （complex、real imag：用于创建和操作复数）
real    		-- 返回complex的虚部
make    		-- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
new        		-- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
cap        		-- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
copy    		-- 用于复制和连接slice，返回复制的数目
len        		-- 来求长度，比如string、array、slice、map、channel ，返回长度
print、println 	-- 底层打印函数，在部署环境中建议使用 fmt 包
```
