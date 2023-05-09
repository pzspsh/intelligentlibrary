

# 互斥锁

在[编程](https://baike.baidu.com/item/编程/139828?fromModule=lemma_inlink)中，引入了对象互斥锁的概念，来保证共享数据操作的完整性。每个[对象](https://baike.baidu.com/item/对象/2331271?fromModule=lemma_inlink)都对应于一个可称为互斥锁的标记，这个标记用来保证在任一时刻，只能有一个[线程](https://baike.baidu.com/item/线程/103101?fromModule=lemma_inlink)访问该[对象](https://baike.baidu.com/item/对象/2331271?fromModule=lemma_inlink)。

- 中文名

  互斥锁

- 领  域

  汇编语言

相关视频

查看全部

![img](https://bkimg.cdn.bcebos.com/smart/96dda144ad345982b2b77ec8aba126adcbef76092e69-bkimg-process,v_1,rw_16,rh_9,maxl_640,pad_1?x-bce-process=image/format,f_auto)

![img](https://bkssl.bdimg.com/static/wiki-lemma/widget/lemma_content/configModule/second/types/common/videoList/img/play-icon_68f96a9.png)

1.1万播放

07:14

Go语言互斥锁之加锁逻辑源码分析

![img](https://bkimg.cdn.bcebos.com/smart/63d0f703918fa0ec08fa286ebbc14eee3d6d55fb6d03-bkimg-process,v_1,rw_16,rh_9,maxl_640,pad_1?x-bce-process=image/format,f_auto)

![img](https://bkssl.bdimg.com/static/wiki-lemma/widget/lemma_content/configModule/second/types/common/videoList/img/play-icon_68f96a9.png)

8370播放

02:54

go如何使用互斥锁同步协程？

![img](https://bkimg.cdn.bcebos.com/smart/902397dda144ad3459821fdd77f71bf431adcbef2f69-bkimg-process,v_1,rw_16,rh_9,maxl_640,pad_1?x-bce-process=image/format,f_auto)

![img](https://bkssl.bdimg.com/static/wiki-lemma/widget/lemma_content/configModule/second/types/common/videoList/img/play-icon_68f96a9.png)

1.1万播放

06:34

Go语言互斥锁的介绍

![img](https://bkimg.cdn.bcebos.com/smart/908fa0ec08fa513d269722bda03b42fbb2fb43167303-bkimg-process,v_1,rw_16,rh_9,maxl_640,pad_1?x-bce-process=image/format,f_auto)

![img](https://bkssl.bdimg.com/static/wiki-lemma/widget/lemma_content/configModule/second/types/common/videoList/img/play-icon_68f96a9.png)

6765播放

02:09

互斥锁3 #python #编程 #程序员 #python全栈

![img](https://bkimg.cdn.bcebos.com/smart/a1ec08fa513d269759ee413ec8ada5fb43166d227203-bkimg-process,v_1,rw_16,rh_9,maxl_640,pad_1?x-bce-process=image/format,f_auto)

![img](https://bkssl.bdimg.com/static/wiki-lemma/widget/lemma_content/configModule/second/types/common/videoList/img/play-icon_68f96a9.png)

6969播放

02:58

互斥锁4 #python #编程 #程序员 #python全栈开发

![img](https://bkimg.cdn.bcebos.com/smart/622762d0f703918fa0ec70a9cc6b319759ee3d6d6e03-bkimg-process,v_1,rw_16,rh_9,maxl_640,pad_1?x-bce-process=image/format,f_auto)

![img](https://bkssl.bdimg.com/static/wiki-lemma/widget/lemma_content/configModule/second/types/common/videoList/img/play-icon_68f96a9.png)

6447播放

01:56

互斥锁1 #python #编程 #程序员 #python全栈开发

## 目录

1. 1 [示例](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#1)
2. 2 [属性对象](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#2)
3. 3 [销毁对象](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#3)
4. 4 [设置范围](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#4)

1. 5 [获取范围](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#5)
2. 6 [类型属性](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#6)
3. 7 [获取属性](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#7)
4. 8 [设置协议](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#8)

1. 9 [获取协议](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#9)
2. 10 [设置上限](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#10)
3. 11 [互斥上限](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#11)
4. 12 [互斥锁](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#12)

1. 13 [获取设置](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#13)
2. 14 [强健属性](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#14)
3. 15 [互斥锁的相关实现与效率问题](https://baike.baidu.com/item/互斥锁?fromId=1889552&redirected=seachword#15)

## 示例

[编辑](javascript:;)[ 播报](javascript:;)

下面举例：

在Posix Thread中定义有一套专门用于[线程同步](https://baike.baidu.com/item/线程同步?fromModule=lemma_inlink)的mutex函数。

1． 创建和销毁

有两种方法创建互斥锁，静态方式和动态方式。POSIX定义了一个[宏](https://baike.baidu.com/item/宏/2648286?fromModule=lemma_inlink)PTHREAD_MUTEX_INITIALIZER来静态初始化互斥锁，方法如下： pthread_mutex_t mutex=PTHREAD_MUTEX_INITIALIZER; 在LinuxThreads实现中，pthread_mutex_t是一个结构，而PTHREAD_MUTEX_INITIALIZER则是一个结构常量。

动态方式是采用pthread_mutex_init()函数来初始化互斥锁，API定义如下： int pthread_mutex_init(pthread_mutex_t *mutex, const pthread_mutexattr_t *mutexattr) 其中mutexattr用于指定互斥锁属性（见下），如果为NULL则使用缺省属性。

pthread_mutex_destroy ()用于注销一个互斥锁，API定义如下： int pthread_mutex_destroy(pthread_mutex_t *mutex) 销毁一个互斥锁即意味着释放它所占用的资源，且要求锁当前处于开放状态。由于在Linux中，互斥锁并不占用任何资源，因此LinuxThreads中的 pthread_mutex_destroy()除了检查锁状态以外（锁定状态则返回EBUSY）没有其他动作。

2． 互斥锁属性

互斥锁的属性在创建锁的时候指定，在LinuxThreads实现中仅有一个锁类型属性，不同的锁类型在试图对一个已经被锁定的互斥锁加锁时表现不同。当前（glibc2.2.3,linuxthreads0.9）有四个值可供选择：

\* PTHREAD_MUTEX_TIMED_NP，这是[缺省值](https://baike.baidu.com/item/缺省值?fromModule=lemma_inlink)，也就是普通锁。当一个线程加锁以后，其余请求锁的线程将形成一个[等待队列](https://baike.baidu.com/item/等待队列?fromModule=lemma_inlink)，并在解锁后按优先级获得锁。这种锁策略保证了资源分配的公平性。

\* PTHREAD_MUTEX_RECURSIVE_NP，嵌套锁，允许同一个线程对同一个锁成功获得多次，并通过多次unlock解锁。如果是不同线程请求，则在加锁线程解锁时重新竞争。

\* PTHREAD_MUTEX_ERRORCHECK_NP，检错锁，如果同一个线程请求同一个锁，则返回EDEADLK，否则与PTHREAD_MUTEX_TIMED_NP类型动作相同。这样就保证当不允许多次加锁时不会出现最简单情况下的死锁。

\* PTHREAD_MUTEX_ADAPTIVE_NP，适应锁，动作最简单的锁类型，仅等待解锁后重新竞争。

3．[锁](https://baike.baidu.com/item/锁/379548?fromModule=lemma_inlink)操作

锁操作主要包括加锁pthread_mutex_lock()、解锁pthread_mutex_unlock()和测试加锁 pthread_mutex_trylock()三个，不论哪种类型的锁，都不可能被两个不同的线程同时得到，而必须等待解锁。对于普通锁和适应锁类型，解锁者可以是同进程内任何线程；而检错锁则必须由加锁者解锁才有效，否则返回EPERM；对于嵌套锁，文档和实现要求必须由加锁者解锁，但实验结果表明并没有这种限制，这个不同还没有得到解释。在同一进程中的线程，如果加锁后没有解锁，则任何其他线程都无法再获得锁。

int pthread_mutex_lock(pthread_mutex_t *mutex)

int pthread_mutex_unlock(pthread_mutex_t *mutex)

int pthread_mutex_trylock(pthread_mutex_t *mutex)

pthread_mutex_trylock()语义与pthread_mutex_lock()类似，不同的是在锁已经被占据时返回EBUSY而不是挂起等待。

4． 其他

[POSIX](https://baike.baidu.com/item/POSIX?fromModule=lemma_inlink) 线程锁机制的[Linux](https://baike.baidu.com/item/Linux?fromModule=lemma_inlink)实现都不是取消点，因此，延迟取消类型的线程不会因收到取消信号而离开加锁等待。值得注意的是，如果线程在加锁后解锁前被取消，锁将永远保持锁定状态，因此如果在关键区段内有取消点存在，或者设置了异步取消类型，则必须在退出[回调函数](https://baike.baidu.com/item/回调函数?fromModule=lemma_inlink)中解锁。

这个锁机制同时也不是异步信号安全的，也就是说，不应该在[信号处理](https://baike.baidu.com/item/信号处理?fromModule=lemma_inlink)过程中使用互斥锁，否则容易造成死锁。

互斥锁属性使用互斥锁（互斥）可以使线程按[顺序执行](https://baike.baidu.com/item/顺序执行?fromModule=lemma_inlink)。通常，互斥锁通过确保一次只有一个线程执行代码的临界段来同步多个线程。互斥锁还可以保护[单线程](https://baike.baidu.com/item/单线程?fromModule=lemma_inlink)代码。

要更改缺省的互斥锁属性，可以对属性对象进行声明和初始化。通常，互斥锁属性会设置在应用程序开头的某个位置，以便可以快速查找和轻松修改。表 4–1列出了用来处理互斥锁属性的函数。

表 4–1 互斥锁属性例程

| **操作**                   | **相关函数说明**                      |
| -------------------------- | ------------------------------------- |
| 初始化互斥锁属性对象       | pthread_mutexattr_init 语法           |
| 销毁互斥锁属性对象         | pthread_mutexattr_destroy 语法        |
| 设置互斥锁范围             | pthread_mutexattr_setpshared 语法     |
| 获取互斥锁范围             | pthread_mutexattr_getpshared 语法     |
| 设置互斥锁的类型属性       | pthread_mutexattr_settype 语法        |
| 获取互斥锁的类型属性       | pthread_mutexattr_gettype 语法        |
| 设置互斥锁属性的协议       | pthread_mutexattr_setprotocol 语法    |
| 获取互斥锁属性的协议       | pthread_mutexattr_getprotocol 语法    |
| 设置互斥锁属性的优先级上限 | pthread_mutexattr_setprioceiling 语法 |
| 获取互斥锁属性的优先级上限 | pthread_mutexattr_getprioceiling 语法 |
| 设置互斥锁的优先级上限     | pthread_mutex_setprioceiling 语法     |
| 获取互斥锁的优先级上限     | pthread_mutex_getprioceiling 语法     |
| 设置互斥锁的强健属性       | pthread_mutexattr_setrobust_np 语法   |
| 获取互斥锁的强健属性       | pthread_mutexattr_getrobust_np 语法   |

表 4–2中显示了在定义互斥范围时 Solaris 线程和 POSIX 线程之间的差异。

表 4–2 互斥锁范围比较

| **Solaris**          | **POSIX**               | **定义**                         |
| -------------------- | ----------------------- | -------------------------------- |
| USYNC_PROCESS        | PTHREAD_PROCESS_SHARED  | 用于同步该进程和其他进程中的线程 |
| USYNC_PROCESS_ROBUST | 无 POSIX 等效项         | 用于在进程间**可靠地**同步线程   |
| USYNC_THREAD         | PTHREAD_PROCESS_PRIVATE | 用于仅同步该进程中的线程         |

## 属性对象

[编辑](javascript:;)[ 播报](javascript:;)

使用pthread_mutexattr_init(3C)可以将与互斥锁对象相关联的属性初始化为其缺省值。在执行过程中，线程系统会为每个属性对象分配[存储空间](https://baike.baidu.com/item/存储空间?fromModule=lemma_inlink)。

**pthread_mutexattr_init 语法**

int pthread_mutexattr_init(pthread_mutexattr_t *mattr);

\#include <pthread.h>

pthread_mutexattr_t mattr;

int ret;/* initialize an attribute to default value */

ret = pthread_mutexattr_init(&mattr);

调用此函数时，pshared 属性的[缺省值](https://baike.baidu.com/item/缺省值?fromModule=lemma_inlink)为 PTHREAD_PROCESS_PRIVATE。该值表示可以在进程内使用经过初始化的互斥锁。

mattr 的类型为 opaque，其中包含一个由系统分配的属性对象。mattr 范围可能的值为 PTHREAD_PROCESS_PRIVATE 和 PTHREAD_PROCESS_SHARED。PTHREAD_PROCESS_PRIVATE 是[缺省值](https://baike.baidu.com/item/缺省值?fromModule=lemma_inlink)。

对于互斥锁属性对象，必须首先通过调用 pthread_mutexattr_destroy(3C) 将其销毁，才能重新初始化该对象。pthread_mutexattr_init()调用会导致分配类型为 opaque 的对象。如果未销毁该对象，则会导致内存泄漏。

ENOMEM

描述:[内存不足](https://baike.baidu.com/item/内存不足?fromModule=lemma_inlink)，无法初始化互斥锁属性对象。

## 销毁对象

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_destroy(3C)可用来取消分配用于维护**pthread_mutexattr_init()**所创建的属性对象的[存储空间](https://baike.baidu.com/item/存储空间?fromModule=lemma_inlink)。

**pthread_mutexattr_destroy 语法**

int pthread_mutexattr_destroy(pthread_mutexattr_t *mattr)#include <pthread.h>pthread_mutexattr_t mattr;int ret;/* destroy an attribute */ret = pthread_mutexattr_destroy(&mattr);

**pthread_mutexattr_destroy 返回值**

**pthread_mutexattr_destroy()**成功完成之后会返回零。其他任何返回值都表示出现了错误。如果出现以下情况，该函数将失败并返回对应的值。

EINVAL

**描述:**由 mattr 指定的值无效。

## 设置范围

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_setpshared(3C)可用来设置互斥锁[变量](https://baike.baidu.com/item/变量?fromModule=lemma_inlink)的[作用域](https://baike.baidu.com/item/作用域?fromModule=lemma_inlink)。

**pthread_mutexattr_setpshared 语法**

int pthread_mutexattr_setpshared(pthread_mutexattr_t *mattr, int pshared);

\#include <pthread.h>

pthread_mutexattr_t mattr;

int ret;

ret = pthread_mutexattr_init(&mattr);/* * resetting to its default value: private */

ret = pthread_mutexattr_setpshared(&mattr, PTHREAD_PROCESS_PRIVATE);

互斥锁[变量](https://baike.baidu.com/item/变量?fromModule=lemma_inlink)可以是进程专用的（进程内）变量，也可以是系统范围内的（进程间）变量。要在多个进程中的线程之间共享互斥锁，可以在[共享内存](https://baike.baidu.com/item/共享内存?fromModule=lemma_inlink)中创建互斥锁，并将**pshared**属性设置为 PTHREAD_PROCESS_SHARED。 此行为与最初的 Solaris 线程实现中**mutex_init()**中的 USYNC_PROCESS 标志等效。

如果互斥锁的**pshared**属性设置为 PTHREAD_PROCESS_PRIVATE，则仅有那些由同一个进程创建的线程才能够处理该互斥锁。

## 获取范围

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_getpshared(3C)可用来返回由**pthread_mutexattr_setpshared()**定义的互斥锁[变量](https://baike.baidu.com/item/变量?fromModule=lemma_inlink)的范围。

**pthread_mutexattr_getpshared 语法**

int pthread_mutexattr_getpshared(pthread_mutexattr_t *mattr, int *pshared);#include <pthread.h>pthread_mutexattr_t mattr;int pshared, ret;/* get pshared of mutex */ret = pthread_mutexattr_getpshared(&mattr, &pshared); 此函数可为属性对象 mattr 获取 pshared 的当前值。该值为 PTHREAD_PROCESS_SHARED 或 PTHREAD_PROCESS_PRIVATE。

## 类型属性

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_settype(3C)可用来设置互斥锁的**type**属性。

**pthread_mutexattr_settype 语法**

\#include <pthread.h>int pthread_mutexattr_settype(pthread_mutexattr_t *attr , int type);类型属性的[缺省值](https://baike.baidu.com/item/缺省值?fromModule=lemma_inlink)为 PTHREAD_MUTEX_DEFAULT。

**type**参数指定互斥锁的类型。以下列出了有效的互斥锁类型：

PTHREAD_MUTEX_NORMAL

**描述:**此类型的互斥锁不会检测死锁。如果线程在不首先解除互斥锁的情况下尝试重新锁定该互斥锁，则会产生死锁。尝试解除由其他线程锁定的互斥锁会产生不确定的行为。如果尝试解除锁定的互斥锁未锁定，则会产生不确定的行为。

PTHREAD_MUTEX_ERRORCHECK

**描述:**此类型的互斥锁可提供错误检查。如果线程在不首先解除锁定互斥锁的情况下尝试重新锁定该互斥锁，则会返回错误。如果线程尝试解除锁定的互斥锁已经由其他线程锁定，则会返回错误。如果线程尝试解除锁定的互斥锁未锁定，则会返回错误。

PTHREAD_MUTEX_RECURSIVE

**描述:**如果线程在不首先解除锁定互斥锁的情况下尝试重新锁定该互斥锁，则可成功锁定该互斥锁。 与 PTHREAD_MUTEX_NORMAL 类型的互斥锁不同，对此类型互斥锁进行重新锁定时不会产生死锁情况。多次锁定互斥锁需要进行相同次数的解除锁定才可以释放该锁，然后其他线程才能获取该互斥锁。如果线程尝试解除锁定的互斥锁已经由其他线程锁定，则会返回错误。 如果线程尝试解除锁定的互斥锁未锁定，则会返回错误。

PTHREAD_MUTEX_DEFAULT

**描述:**如果尝试以[递归](https://baike.baidu.com/item/递归?fromModule=lemma_inlink)方式锁定此类型的互斥锁，则会产生不确定的行为。对于不是由调用线程锁定的此类型互斥锁，如果尝试对它解除锁定，则会产生不确定的行为。对于尚未锁定的此类型互斥锁，如果尝试对它解除锁定，也会产生不确定的行为。允许在实现中将该互斥锁映射到其他互斥锁类型之一。对于 Solaris 线程，PTHREAD_PROCESS_DEFAULT 会映射到 PTHREAD_PROCESS_NORMAL。

**pthread_mutexattr_settype 返回值**

如果运行成功，pthread_mutexattr_settype 函数会返回零。否则，将返回用于指明错误的错误号。

EINVAL

**描述:**值为**type**无效。

EINVAL

**描述:attr**指定的值无效。

## 获取属性

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_gettype(3C)可用来获取由**pthread_mutexattr_settype()**设置的互斥锁的**type**属性。

**pthread_mutexattr_gettype 语法**

\#include <pthread.h>int pthread_mutexattr_gettype(pthread_mutexattr_t *attr , int *type);类型属性的缺省值为 PTHREAD_MUTEX_DEFAULT。

**type**参数指定互斥锁的类型。有效的互斥锁类型包括：

PTHREAD_MUTEX_NORMAL

PTHREAD_MUTEX_ERRORCHECK

PTHREAD_MUTEX_RECURSIVE

PTHREAD_MUTEX_DEFAULT

有关每种类型的说明，请参见pthread_mutexattr_settype 语法。

**pthread_mutexattr_gettype 返回值**

如果成功完成，**pthread_mutexattr_gettype()**会返回 0。其他任何返回值都表示出现了错误。

## 设置协议

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_setprotocol(3C)可用来设置互斥锁属性对象的协议属性。

**pthread_mutexattr_setprotocol 语法**

\#include <pthread.h>int pthread_mutexattr_setprotocol(pthread_mutexattr_t *attr, int protocol);attr 指示以前调用**pthread_mutexattr_init()**时创建的互斥锁属性对象。

protocol 可定义应用于互斥锁属性对象的协议。

pthread.h 中定义的 protocol 可以是以下值之一：PTHREAD_PRIO_NONE、PTHREAD_PRIO_INHERIT 或 PTHREAD_PRIO_PROTECT。

PTHREAD_PRIO_NONE

线程的优先级和调度不会受到互斥锁拥有权的影响。

PTHREAD_PRIO_INHERIT

此协议值（如 thrd1）会影响线程的优先级和调度。如果更高优先级的线程因 thrd1 所拥有的一个或多个互斥锁而被阻塞，而这些互斥锁是用 PTHREAD_PRIO_INHERIT 初始化的，则 thrd1 将以高于它的优先级或者所有正在等待这些互斥锁（这些互斥锁是 thrd1 指所拥有的互斥锁）的线程的最高优先级运行。

如果 thrd1 因另一个线程 (thrd3) 拥有的互斥锁而被阻塞，则相同的优先级继承效应会以[递归](https://baike.baidu.com/item/递归?fromModule=lemma_inlink)方式传播给 thrd3。

使用 PTHREAD_PRIO_INHERIT 可以避免优先级倒置。低优先级的线程持有较高优先级线程所需的锁时，便会发生优先级倒置。只有在较低优先级的线程释放该锁之后，较高优先级的线程才能继续使用该锁。设置 PTHREAD_PRIO_INHERIT，以便按与预期的优先级相反的优先级处理每个线程。

如果为使用协议属性值 PTHREAD_PRIO_INHERIT 初始化的互斥锁定义了 _POSIX_THREAD_PRIO_INHERIT，则互斥锁的属主失败时会执行以下操作。属主失败时的行为取决于**pthread_mutexattr_setrobust_np()**的 robustness 参数的值。

解除锁定互斥锁。

互斥锁的下一个属主将获取该互斥锁，并返回错误 EOWNERDEAD。

互斥锁的下一个属主会尝试使该互斥锁所保护的状态一致。如果上一个属主失败，则状态可能会不一致。如果属主成功使状态保持一致，则可针对该互斥锁调用**pthread_mutex_init()**并解除锁定该互斥锁。

**注 –**如果针对以前初始化的但尚未销毁的互斥锁调用**pthread_mutex_init()**，则该互斥锁不会重新初始化。

如果属主无法使状态保持一致，**请勿**调用**pthread_mutex_init()**，而是解除锁定该互斥锁。在这种情况下，所有等待的线程都将被唤醒。以后对**pthread_mutex_lock()**的所有调用将无法获取互斥锁，并将返回错误代码 ENOTRECOVERABLE。现在，通过调用**pthread_mutex_destroy()**来取消初始化该互斥锁，即可使其状态保持一致。调用**pthread_mutex_init()**可重新初始化互斥锁。

如果已获取该锁的线程失败并返回 EOWNERDEAD，则下一个属主将获取该锁及错误代码 EOWNERDEAD。

PTHREAD_PRIO_PROTECT

当线程拥有一个或多个使用 PTHREAD_PRIO_PROTECT 初始化的互斥锁时，此协议值会影响其他线程（如 thrd2）的优先级和调度。thrd2 以其较高的优先级或者以 thrd2 拥有的所有互斥锁的最高优先级上限运行。基于被 thrd2 拥有的任一互斥锁阻塞的较高优先级线程对于 thrd2 的调度没有任何影响。

如果某个线程调用**sched_setparam()**来更改初始优先级，则调度程序不会采用新优先级将该线程移到调度队列末尾。

线程拥有使用 PTHREAD_PRIO_INHERIT 或 PTHREAD_PRIO_PROTECT 初始化的互斥锁

线程解除锁定使用 PTHREAD_PRIO_INHERIT 或 PTHREAD_PRIO_PROTECT 初始化的互斥锁

一个[线程](https://baike.baidu.com/item/线程?fromModule=lemma_inlink)可以同时拥有多个混合使用 PTHREAD_PRIO_INHERIT 和 PTHREAD_PRIO_PROTECT 初始化的互斥锁。在这种情况下，该线程将以通过其中任一协议获取的最高优先级执行。

**pthread_mutexattr_setprotocol 返回值**

如果成功完成，**pthread_mutexattr_setprotocol()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下任一情况，**pthread_mutexattr_setprotocol()**将失败并返回对应的值。

ENOSYS

**描述:**选项 _POSIX_THREAD_PRIO_INHERIT 和 _POSIX_THREAD_PRIO_PROTECT 均未定义并且该实现不支持此函数。

ENOTSUP

**描述:**protocol 指定的值不受支持。

如果出现以下任一情况，**pthread_mutexattr_setprotocol()**可能会失败并返回对应的值。

EINVAL

**描述:**attr 或 protocol 指定的值无效。

EPERM

**描述:**调用方无权执行该操作。

## 获取协议

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_getprotocol(3C)可用来获取互斥锁属性对象的协议属性。

**pthread_mutexattr_getprotocol 语法**

\#include <pthread.h>int pthread_mutexattr_getprotocol(const pthread_mutexattr_t *attr, int *protocol);attr 指示以前调用**pthread_mutexattr_init()**时创建的互斥锁属性对象。

protocol 包含以下协议属性之一：PTHREAD_PRIO_NONE、PTHREAD_PRIO_INHERIT 或 PTHREAD_PRIO_PROTECT。

**pthread_mutexattr_getprotocol 返回值**

如果成功完成，**pthread_mutexattr_getprotocol()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下情况，**pthread_mutexattr_getprotocol()**将失败并返回对应的值。

ENOSYS

**描述:**_POSIX_THREAD_PRIO_INHERIT 选项和 _POSIX_THREAD_PRIO_PROTECT 选项均未定义并且该实现不支持此函数。

如果出现以下任一情况，**pthread_mutexattr_getprotocol()**可能会失败并返回对应的值。

EINVAL

**描述:**attr 指定的值无效。

EPERM

**描述:**调用方无权执行该操作。

## 设置上限

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_setprioceiling(3C)可用来设置互斥锁属性对象的优先级上限属性。

**pthread_mutexattr_setprioceiling 语法**

\#include <pthread.h>int pthread_mutexattr_setprioceiling(pthread_mutexatt_t *attr, int prioceiling, int *oldceiling);attr 指示以前调用**pthread_mutexattr_init()**时创建的互斥锁属性对象。

prioceiling 指定已初始化互斥锁的优先级上限。优先级上限定义执行互斥锁保护的临界段时的最低优先级。prioceiling 位于 SCHED_FIFO 所定义的优先级的最大范围内。要避免优先级倒置，请将 prioceiling 设置为高于或等于可能会锁定特定互斥锁的所有线程的最高优先级。

oldceiling 包含以前的优先级上限值。

**pthread_mutexattr_setprioceiling 返回值**

如果成功完成，**pthread_mutexattr_setprioceiling()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下任一情况，**pthread_mutexattr_setprioceiling()**将失败并返回对应的值。

ENOSYS

**描述:**选项 _POSIX_THREAD_PRIO_PROTECT 未定义并且该实现不支持此函数。

如果出现以下任一情况，**pthread_mutexattr_setprioceiling()**可能会失败并返回对应的值。

EINVAL

**描述:**attr 或 prioceiling 指定的值无效。

EPERM

**描述:**调用方无权执行该操作。

## 互斥上限

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_getprioceiling(3C)可用来获取互斥锁属性对象的优先级上限属性。

**pthread_mutexattr_getprioceiling 语法**

\#include <pthread.h>int pthread_mutexattr_getprioceiling(const pthread_mutexatt_t *attr, int *prioceiling);attr 指定以前调用**pthread_mutexattr_init()**时创建的属性对象。

**注 –**

仅当定义了 _POSIX_THREAD_PRIO_PROTECT 符号时，attr 互斥锁属性对象才会包括优先级上限属性。

**pthread_mutexattr_getprioceiling()**返回 prioceiling 中已初始化互斥锁的优先级上限。优先级上限定义执行互斥锁保护的临界段时的最低优先级。prioceiling 位于 SCHED_FIFO 所定义的优先级的最大范围内。要避免优先级倒置，请将 prioceiling 设置为高于或等于可能会锁定特定互斥锁的所有线程的最高优先级。

**pthread_mutexattr_getprioceiling 返回值**

如果成功完成，**pthread_mutexattr_getprioceiling()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下任一情况，**pthread_mutexattr_getprioceiling()**将失败并返回对应的值。

ENOSYS

**描述:**_POSIX_THREAD_PRIO_PROTECT 选项未定义并且该实现不支持此函数。

如果出现以下任一情况，**pthread_mutexattr_getprioceiling()**可能会失败并返回对应的值。

EINVAL

**描述:**attr 指定的值无效。

EPERM

**描述:**调用方无权执行该操作。

## 互斥锁

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_setprioceiling(3C)可用来设置互斥锁的优先级上限。

**pthread_mutex_setprioceiling 语法**

\#include <pthread.h>int pthread_mutex_setprioceiling(pthread_mutex_t *mutex, int prioceiling, int *old_ceiling);**pthread_mutex_setprioceiling()**可更改互斥锁 mutex 的优先级上限 prioceiling。**pthread_mutex_setprioceiling()**可锁定互斥锁（如果未锁定的话），或者一直处于[阻塞状态](https://baike.baidu.com/item/阻塞状态?fromModule=lemma_inlink)，直到**pthread_mutex_setprioceiling()**成功锁定该互斥锁，更改该互斥锁的优先级上限并将该互斥锁释放为止。锁定互斥锁的过程无需遵循优先级保护协议。

如果**pthread_mutex_setprioceiling()**成功，则将在 old_ceiling 中返回以前的优先级上限值。如果**pthread_mutex_setprioceiling()**失败，则互斥锁的优先级上限保持不变。

**pthread_mutex_setprioceiling 返回值**

如果成功完成，**pthread_mutex_setprioceiling()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下情况，**pthread_mutexatt_setprioceiling()**将失败并返回对应的值。

ENOSYS

**描述:**选项_POSIX_THREAD_PRIO_PROTECT 未定义并且该实现不支持此函数。

如果出现以下任一情况，**pthread_mutex_setprioceiling()**可能会失败并返回对应的值。

EINVAL

**描述:**prioceiling 所请求的优先级超出了范围。

EINVAL

**描述:**mutex 指定的值不会引用当前存在的互斥锁。

ENOSYS

**描述:**该实现不支持互斥锁的优先级上限协议。

EPERM

**描述:**调用方无权执行该操作。

## 获取设置

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_getprioceiling(3C)可用来获取互斥锁的优先级上限。

**pthread_mutex_getprioceiling 语法**

\#include <pthread.h>int pthread_mutex_getprioceiling(const pthread_mutex_t *mutex, int *prioceiling);**pthread_mutex_getprioceiling()**会返回 mutex 的优先级上限 prioceiling。

**pthread_mutex_getprioceiling 返回值**

如果成功完成，**pthread_mutex_getprioceiling()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下任一情况，**pthread_mutexatt_getprioceiling()**将失败并返回对应的值。

ENOSYS

**描述:**_POSIX_THREAD_PRIO_PROTECT 选项未定义并且该实现不支持此函数。

如果出现以下任一情况，**pthread_mutex_getprioceiling()**可能会失败并返回对应的值。

EINVAL

**描述:**mutex 指定的值不会引用当前存在的互斥锁。

ENOSYS

**描述:**该实现不支持互斥锁的优先级上限协议。

EPERM

**描述:**调用方无权执行该操作。

## 强健属性

[编辑](javascript:;)[ 播报](javascript:;)

pthread_mutexattr_setrobust_np(3C)可用来设置互斥锁属性对象的强健属性。

**pthread_mutexattr_setrobust_np 语法**

\#include <pthread.h>int pthread_mutexattr_setrobust_np(pthread_mutexattr_t *attr, int *robustness);**注 –**

仅当定义了符号 _POSIX_THREAD_PRIO_INHERIT 时，**pthread_mutexattr_setrobust_np()**才适用。

attr 指示以前通过调用**pthread_mutexattr_init()**创建的互斥锁属性对象。

robustness 定义在互斥锁的属主失败时的行为。pthread.h 中定义的 robustness 的值为 PTHREAD_MUTEX_ROBUST_NP 或 PTHREAD_MUTEX_STALLED_NP。缺省值为 PTHREAD_MUTEX_STALLED_NP。

PTHREAD_MUTEX_ROBUST_NP

如果互斥锁的属主失败，则以后对**pthread_mutex_lock()**的所有调用将以不确定的方式被阻塞。

PTHREAD_MUTEX_STALLED_NP

互斥锁的属主失败时，将会解除锁定该互斥锁。互斥锁的下一个属主将获取该互斥锁，并返回错误 EOWNWERDEAD。

**注 –**应用程序必须检查pthread_mutex_lock()的返回代码，查找返回错误 EOWNWERDEAD 的互斥锁。

互斥锁的新属主应使该互斥锁所保护的状态保持一致。如果上一个属主失败，则互斥锁状态可能会不一致。

如果新属主能够使状态保持一致，请针对该互斥锁调用pthread_mutex_consistent_np()，并解除锁定该互斥锁。

如果新属主无法使状态保持一致，请勿针对该互斥锁调用pthread_mutex_consistent_np()，而是解除锁定该互斥锁。

所有等待的线程都将被唤醒，以后对pthread_mutex_lock()的所有调用都将无法获取该互斥锁。返回代码为 ENOTRECOVERABLE。通过调用pthread_mutex_destroy()取消对互斥锁的初始化，并调用pthread_mutex_int()重新初始化该互斥锁，可使该互斥锁保持一致。

如果已获取该锁的线程失败并返回 EOWNERDEAD，则下一个属主获取该锁时将返回代码 EOWNERDEAD。

**pthread_mutexattr_setrobust_np 返回值**

如果成功完成，**pthread_mutexattr_setrobust_np()**会返回 0。其他任何返回值都表示出现了错误。

如果出现以下任一情况，**pthread_mutexattr_setrobust_np()**将失败并返回对应的值。

ENOSYS

**描述:**选项 _POSIX_THREAD_PRIO__INHERIT 未定义，或者该实现不支持**pthread_mutexattr_setrobust_np()**。

ENOTSUP

描述:robustness 指定的值不受支持。

pthread_mutexattr_setrobust_np()可能会在出现以下情况时失败：

EINVAL

描述:attr 或 robustness 指定的值无效。

## 互斥锁的相关实现与效率问题

[编辑](javascript:;)[ 播报](javascript:;)

互斥锁实际的效率还是可以让人接受的，加锁的时间大概100ns左右，而实际上互斥锁的一种可能的实现是先自旋一段时间，当自旋的时间超过阀值之后再将线程投入睡眠中，因此在并发运算中使用互斥锁（每次占用锁的时间很短）的效果可能不亚于使用自旋锁。