# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
""" 
BaseException
 +-- SystemExit
 +-- KeyboardInterrupt
 +-- GeneratorExit
 +-- Exception
      +-- StopIteration
      +-- StopAsyncIteration
      +-- ArithmeticError
      |    +-- FloatingPointError
      |    +-- OverflowError
      |    +-- ZeroDivisionError
      +-- AssertionError
      +-- AttributeError
      +-- BufferError
      +-- EOFError
      +-- ImportError
      |    +-- ModuleNotFoundError
      +-- LookupError
      |    +-- IndexError
      |    +-- KeyError
      +-- MemoryError
      +-- NameError
      |    +-- UnboundLocalError
      +-- OSError
      |    +-- BlockingIOError
      |    +-- ChildProcessError
      |    +-- ConnectionError
      |    |    +-- BrokenPipeError
      |    |    +-- ConnectionAbortedError
      |    |    +-- ConnectionRefusedError
      |    |    +-- ConnectionResetError
      |    +-- FileExistsError
      |    +-- FileNotFoundError
      |    +-- InterruptedError
      |    +-- IsADirectoryError
      |    +-- NotADirectoryError
      |    +-- PermissionError
      |    +-- ProcessLookupError
      |    +-- TimeoutError
      +-- ReferenceError
      +-- RuntimeError
      |    +-- NotImplementedError
      |    +-- RecursionError
      +-- SyntaxError
      |    +-- IndentationError
      |         +-- TabError
      +-- SystemError
      +-- TypeError
      +-- ValueError
      |    +-- UnicodeError
      |         +-- UnicodeDecodeError
      |         +-- UnicodeEncodeError
      |         +-- UnicodeTranslateError
      +-- Warning
           +-- DeprecationWarning
           +-- PendingDeprecationWarning
           +-- RuntimeWarning
           +-- SyntaxWarning
           +-- UserWarning
           +-- FutureWarning
           +-- ImportWarning
           +-- UnicodeWarning
           +-- BytesWarning
           +-- EncodingWarning
           +-- ResourceWarning
"""


"""
什么是异常：
1、异常指出了我们的程序有错误
2、有些异常也会在一些合法的情况下发生，比如用户名密码错误，银行卡号不存在
3、内置异常的名字都是以Error结尾：ZeroDivisionError,IndexError，SyntaxError
4、所有异常类都是继承于Exception，（扩展BaseException）
5、当一个异常发生的时候，会立即停止程序的执行，除非正确的处理这个异常
6、异常是一个对象，并且可以继承（通过继承Exception类来实现自己的异常）
"""

# print "hello"
# 抛出异常信息：SyntaxError: Missing parentheses in call to 'print'. Did you mean print("hello")?
# 产生的SyntaxError异常，这就是异常，语法错误的异常

# x = 5 / 0
# print(x)
# 抛出ZeroDivisionError异常，信息：ZeroDivisionError: division by zero，说是不能被除

# lst = [1,2,3]
# print(lst[3])
# 抛出异常：IndexError: list index out of range 说是索引错误


# 通过内置的TypeError和ValueError类来构造异常对象，下面的例子扩展了内置的list，并重写了该内之类的append方法
class MyList(list):  # 继承内置的list类
    def append(self, integer):
        if not isinstance(integer, int):  # 如果不是整数则抛出异常
            raise TypeError("Not an integer")
        if integer % 2:
            raise ValueError("Can not be divisible")  # 不能整除时则抛出异常
        super().append(integer)  # 调用父类的append方法


# mylist = MyList()
# #mylist.append(12.45) #引发TypeError异常
# #mylist.append(87) #引发ValueError异常
# mylist.append(64) #无异常


# 发生异常时，程序是怎样的？
def test_return():
    print("hello")  # 这条是会被执行的
    raise Exception("My God, something went wrong")  # 这里引发异常后，后面的代码永远不会执行，包括return语句
    print("How are you?")
    return "I'm very good"


# test_return()


# 通过另外一个函数来调用test_return函数，看看效果
def call_test_return():
    print("start call...")
    test_return()  # 在这里调用test_return函数
    print("an exception was raised....")
    print("so...")


# call_test_return()
"""
在call_test_return函数中调用test_return函数,在test_return函数中有异常的发生
但是对于call_test_return函数是没有异常语句的。可为什么连call_test_return函数都停止执行了呢？
原因是：
    异常抛出会停止在call_test_return函数调用栈内所有代码的执行
"""

# 即然有异常，就要处理它，如何处理？
# try:
#     test_return()
# except: #在这里捕捉到了异常，因此输出下面的print语句
#     print("test_return Function An exception occurs") #提示有异常发生
# print("end...")
"""
捕捉了异常，并且在发生异常时应该做什么，因此，程序没有被终止掉
对于在test_return函数，在抛出异常的语句之后的代码是没有被执行的
1、try语句可以包含任何可能会发生异常的代码
2、except语句将捕获任何类型的异常，而不是捕获有针对性的异常，它是捕获所有。
3、那么如何捕获指定的异常类型？看下面代码
"""


# 捕捉指定的异常类型
def func_a(number):
    try:
        return 100 / number
    except ZeroDivisionError:
        print("Can not be 0")


# print(func_a(0)) #抛出ZeroDivisionError类型的异常
# print(func_a("abcdef")) #抛出TypeError类型的异常，但目前为止的代码，没有写捕获TypeError类型的异常，因此这个异常无法被捕捉
# 也就是说，TypeError异常，不包含在要处理的异常类型中
# 那么如何才能同时捕捉多种类型的异常？改进代码，如下


def func_b(number):
    try:
        return 100 / number
    except (ZeroDivisionError, TypeError):
        print("Unknown value...")


# print(func_b(0))
# print(func_b("abcdef"))
# 非常完美，貌似这两种异常类型都捕捉到了
# 但是，这里有一个弊端，就是，我想捕获不同类型的异常并且对它们做出不同的操作，目前的代码是无法实现的，
# 那么，为了实现这个想法，继续改进代码，如下
def func_c(number):
    try:
        return 100 / number
    except ZeroDivisionError:
        print("Unknown value...")  # 捕获到ZeroDivisionError异常，就执行此操作
    except TypeError:
        print("Value Type Error...")  # 捕获到TypeError异常，就执行这个操作


# func_c(0)
# func_c("abc")
# print(func_c(89))
# 因此，非常完美！


# 思考一个问题，如果捕获任何类型的异常再最前面会是什么情况？看如下代码
def func_d(number):
    try:
        return 100 / number
    except Exception:  # 捕获任何异常类型
        print("Exception....")
    except ZeroDivisionError:
        print("Unknown value...")
    except TypeError:
        print("Value Type Error...")


# func_d(0)
# func_d("abc")
# 效果是，虽然明确知道会发生哪种类型的异常并有针对性的捕获，但是捕获任何类型异常在最前面，导致有针对性的捕获根本就没有捕获
# 为什么会这样？因为ZeroDivisionError，TypeError这些内置的异常类型都是从Exception类继承而来的，也就是说，已经捕获了，就没必要再去有针对性的捕获。
# OK，那么如果将他们的顺序反过来（也就是except Exception在最后面），又会是啥情况？看下面的代码（下面的代码去掉了捕获TypeError类型）
def func_e(number):
    try:
        return 100 / number
    except ZeroDivisionError:
        print("Unknown value...")
    except Exception:  # 捕获任何异常类型
        print("all Exception....")


# func_e(0) #这里引发的是ZeroDivisionError类型的异常
# func_e("abc") #这里原本是引发TypeError类型的异常，但去掉后因此无法捕捉，所以由except Exception语句来负责捕获剩下的所有异常
# 通过效果，非常完美，而该在怎样的应用场景去使用它已经不必多说，很显然知道怎么去应用它了

"""
有没有注意到，在上面的例子中，捕获异常之后所做的操作是打印一句话，
但是实际的操作不只是打印一句话，也可以是做别的操作，比如做运算，或者继续循环，或者断开连接等等操作
有一种情况是，我不想仅仅只是做操作，我还想知道它所引发的异常的具体信息。那么如何查看？看下面的代码
"""


def func_f(number):
    try:
        return 100 / number
    except ZeroDivisionError as err:
        print("[Error]:%s [args]:%s" % (err, err.args))


# func_f(0)
# 这里是通过as关键字来捕获到异常作为变量来访问，err.args则是获取传给函数的参数
# 关键字as，在异常中使用，是在python3版本中，而对于python2，则使用的是一个逗号

"""
之前说了，把可能发生异常的代码丢进try中，那么如果被丢进try的代码没有发生异常呢？
如果没有异常，不仅要执行try中的代码，并且同时我还需要执行别的操作，
如果有异常的发生，那么就只捕获异常，并执行对应的动作，无需执行别的额外操作。
那么请看下面改进后的代码
"""


def func_g(number):
    try:
        ret = 100 / number
        print(ret)
    except ZeroDivisionError as err:
        print("[Error]:%s [args]:%s" % (err, err.args))
    else:
        print("calculation done...")


# func_g(5) #传入5到函数中进行计算，没有异常，没有异常并且也要执行else后面的语句，因此达到了目的
# func_g(0) #传入0，则引发异常，那么仅仅只是执行了except中的捕获操作，else后面的语句没有被执行

"""
上面的代码似乎非常完美，我又有一个需求，就是语句无论是否发生异常都将执行我指定的操作
改进如下
"""


def func_h(number):
    try:
        ret = 100 / number
        print(ret)
    except ZeroDivisionError as err:
        print("[Error]:%s [args]:%s" % (err, err.args))
    else:
        print("calculation done...")
    finally:
        print("code end...")


# func_h(0) #传入0，引发异常，并且，也继续执行了finally后面语句，但是else则没有执行，非常完美，达到我的目的
# func_h(8) #传入的是8，没有异常则无需捕获，那么else是在没有异常的情况下才执行，那么他执行了。这非常正确，finally后面的语句也执行了
# 通过看到的效果，finally确实是不管有没有异常的发生，都确实是会执行
