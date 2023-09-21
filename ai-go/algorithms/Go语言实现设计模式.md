导语| 设计模式是针对软件设计中常见问题的工具箱，其中的工具就是各种经过实践验证的解决方案。即使你从未遇到过这些问题，了解模式仍然非常有用，因为它能指导你如何使用面向对象的设计原则来解决各种问题，提高开发效率，降低开发成本；本文囊括了GO语言实现的经典设计模式示例，每个示例都精心设计，力求符合模式结构，可作为日常编码参考，同时一些常用的设计模式融入了开发实践经验总结，帮助大家在平时工作中灵活运用。



# ![图片](F:\Images\责任链模式.png)                			                    **责任链模式**

## **(一）概念**

责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送。收到请求后，每个处理者均可对请求进行处理，或将其传递给链上的下个处理者。

该模式允许多个对象来对请求进行处理，而无需让发送者类与具体接收者类相耦合。链可在运行时由遵循标准处理者接口的任意处理者动态生成。

一般意义上的责任链模式是说，请求在链上流转时任何一个满足条件的节点处理完请求后就会停止流转并返回，不过还可以根据不同的业务情况做一些改进：

- 请求可以流经处理链的所有节点，不同节点会对请求做不同职责的处理；
- 可以通过上下文参数保存请求对象及上游节点的处理结果，供下游节点依赖，并进一步处理；
- 处理链可支持节点的异步处理，通过实现特定接口判断，是否需要异步处理；
- 责任链对于请求处理节点可以设置停止标志位，不是异常，是一种满足业务流转的中断；
- 责任链的拼接方式存在两种，一种是节点遍历，一个节点一个节点顺序执行；另一种是节点嵌套，内层节点嵌入在外层节点执行逻辑中，类似递归，或者“回”行结构；
- 责任链的节点嵌套拼接方式多被称为拦截器链或者过滤器链，更易于实现业务流程的切面，比如监控业务执行时长，日志输出，权限校验等；



## **（二）示例**

本示例模拟实现机场登机过程，第一步办理登机牌，第二步如果有行李，就办理托运，第三步核实身份，第四步安全检查，第五步完成登机；其中行李托运是可选的，其他步骤必选，必选步骤有任何不满足就终止登机；旅客对象作为请求参数上下文，每个步骤会根据旅客对象状态判断是否处理或流转下一个节点；



### **（三）登机过程**

```go
package chainofresponsibility

import "fmt"

// BoardingProcessor 登机过程中，各节点统一处理接口
type BoardingProcessor interface {
  SetNextProcessor(processor BoardingProcessor)
  ProcessFor(passenger *Passenger)
}

// Passenger 旅客
type Passenger struct {
  name                  string // 姓名
  hasBoardingPass       bool   // 是否办理登机牌
  hasLuggage            bool   // 是否有行李需要托运
  isPassIdentityCheck   bool   // 是否通过身份校验
  isPassSecurityCheck   bool   // 是否通过安检
  isCompleteForBoarding bool   // 是否完成登机
}

// baseBoardingProcessor 登机流程处理器基类
type baseBoardingProcessor struct {
  // nextProcessor 下一个登机处理流程
  nextProcessor BoardingProcessor
}

// SetNextProcessor 基类中统一实现设置下一个处理器方法
func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) {
  b.nextProcessor = processor
}

// ProcessFor 基类中统一实现下一个处理器流转
func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) {
  if b.nextProcessor != nil {
    b.nextProcessor.ProcessFor(passenger)
  }
}

// boardingPassProcessor 办理登机牌处理器
type boardingPassProcessor struct {
  baseBoardingProcessor // 引用基类
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) {
  if !passenger.hasBoardingPass {
    fmt.Printf("为旅客%s办理登机牌;\n", passenger.name)
    passenger.hasBoardingPass = true
  }
  // 成功办理登机牌后，进入下一个流程处理
  b.baseBoardingProcessor.ProcessFor(passenger)
}

// luggageCheckInProcessor 托运行李处理器
type luggageCheckInProcessor struct {
  baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) {
  if !passenger.hasBoardingPass {
    fmt.Printf("旅客%s未办理登机牌，不能托运行李;\n", passenger.name)
    return
  }
  if passenger.hasLuggage {
    fmt.Printf("为旅客%s办理行李托运;\n", passenger.name)
  }
  l.baseBoardingProcessor.ProcessFor(passenger)
}

// identityCheckProcessor 校验身份处理器
type identityCheckProcessor struct {
  baseBoardingProcessor
}

func (i *identityCheckProcessor) ProcessFor(passenger *Passenger) {
  if !passenger.hasBoardingPass {
    fmt.Printf("旅客%s未办理登机牌，不能办理身份校验;\n", passenger.name)
    return
  }
  if !passenger.isPassIdentityCheck {
    fmt.Printf("为旅客%s核实身份信息;\n", passenger.name)
    passenger.isPassIdentityCheck = true
  }
  i.baseBoardingProcessor.ProcessFor(passenger)
}

// securityCheckProcessor 安检处理器
type securityCheckProcessor struct {
  baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessFor(passenger *Passenger) {
  if !passenger.hasBoardingPass {
    fmt.Printf("旅客%s未办理登机牌，不能进行安检;\n", passenger.name)
    return
  }
  if !passenger.isPassSecurityCheck {
    fmt.Printf("为旅客%s进行安检;\n", passenger.name)
    passenger.isPassSecurityCheck = true
  }
  s.baseBoardingProcessor.ProcessFor(passenger)
}

// completeBoardingProcessor 完成登机处理器
type completeBoardingProcessor struct {
  baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessFor(passenger *Passenger) {
  if !passenger.hasBoardingPass ||
    !passenger.isPassIdentityCheck ||
    !passenger.isPassSecurityCheck {
    fmt.Printf("旅客%s登机检查过程未完成，不能登机;\n", passenger.name)
    return
  }
  passenger.isCompleteForBoarding = true
  fmt.Printf("旅客%s成功登机;\n", passenger.name)
}
```

### **（四）测试程序**

```go
package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
  boardingProcessor := BuildBoardingProcessorChain()
  passenger := &Passenger{
    name:                  "李四",
    hasBoardingPass:       false,
    hasLuggage:            true,
    isPassIdentityCheck:   false,
    isPassSecurityCheck:   false,
    isCompleteForBoarding: false,
  }
  boardingProcessor.ProcessFor(passenger)
}

// BuildBoardingProcessorChain 构建登机流程处理链
func BuildBoardingProcessorChain() BoardingProcessor {
  completeBoardingNode := &completeBoardingProcessor{}

  securityCheckNode := &securityCheckProcessor{}
  securityCheckNode.SetNextProcessor(completeBoardingNode)

  identityCheckNode := &identityCheckProcessor{}
  identityCheckNode.SetNextProcessor(securityCheckNode)

  luggageCheckInNode := &luggageCheckInProcessor{}
  luggageCheckInNode.SetNextProcessor(identityCheckNode)

  boardingPassNode := &boardingPassProcessor{}
  boardingPassNode.SetNextProcessor(luggageCheckInNode)
  return boardingPassNode
}
```

### **（五）运行结果** 

```go
=== RUN   TestChainOfResponsibility
为旅客李四办理登机牌;
为旅客李四办理行李托运;
为旅客李四核实身份信息;
为旅客李四进行安检;
旅客李四成功登机;
--- PASS: TestChainOfResponsibility (0.00s)
PASS
```



# ![图片](F:\Images\命令行模式.png)                			                    **命令模式**

## **（一）概念**

命令模式是一种行为设计模式，它可将请求转换为一个包含与请求相关的所有信息的独立对象。该转换让你能根据不同的请求将方法参数化、延迟请求执行或将其放入队列中，且能实现可撤销操作。



方法参数化是指将每个请求参数传入具体命令的工厂方法（go语言没有构造函数）创建命令，同时具体命令会默认设置好接受对象，这样做的好处是不管请求参数个数及类型，还是接受对象有几个，都会被封装到具体命令对象的成员字段上，并通过统一的Execute接口方法进行调用，屏蔽各个请求的差异，便于命令扩展，多命令组装，回滚等；



## **（二）示例**

控制电饭煲做饭是一个典型的命令模式的场景，电饭煲的控制面板会提供设置煮粥、蒸饭模式，及开始和停止按钮，电饭煲控制系统会根据模式的不同设置相应的火力，压强及时间等参数；煮粥，蒸饭就相当于不同的命令，开始按钮就相当命令触发器，设置好做饭模式，点击开始按钮电饭煲就开始运行，同时还支持停止命令；



### **（三）电饭煲接收器**

```go
package command

import "fmt"

// ElectricCooker 电饭煲
type ElectricCooker struct {
  fire     string // 火力
  pressure string // 压力
}

// SetFire 设置火力
func (e *ElectricCooker) SetFire(fire string) {
  e.fire = fire
}

// SetPressure 设置压力
func (e *ElectricCooker) SetPressure(pressure string) {
  e.pressure = pressure
}

// Run 持续运行指定时间
func (e *ElectricCooker) Run(duration string) string {
  return fmt.Sprintf("电饭煲设置火力为%s,压力为%s,持续运行%s;", e.fire, e.pressure, duration)
}

// Shutdown 停止
func (e *ElectricCooker) Shutdown() string {
  return "电饭煲停止运行。"
}

```

### **（四）电饭煲命令**

```go

package command

// CookCommand 做饭指令接口
type CookCommand interface {
  Execute() string // 指令执行方法
}

// steamRiceCommand 蒸饭指令
type steamRiceCommand struct {
  electricCooker *ElectricCooker // 电饭煲
}

func NewSteamRiceCommand(electricCooker *ElectricCooker) *steamRiceCommand {
  return &steamRiceCommand{
    electricCooker: electricCooker,
  }
}

func (s *steamRiceCommand) Execute() string {
  s.electricCooker.SetFire("中")
  s.electricCooker.SetPressure("正常")
  return "蒸饭:" + s.electricCooker.Run("30分钟")
}

// cookCongeeCommand 煮粥指令
type cookCongeeCommand struct {
  electricCooker *ElectricCooker
}

func NewCookCongeeCommand(electricCooker *ElectricCooker) *cookCongeeCommand {
  return &cookCongeeCommand{
    electricCooker: electricCooker,
  }
}

func (c *cookCongeeCommand) Execute() string {
  c.electricCooker.SetFire("大")
  c.electricCooker.SetPressure("强")
  return "煮粥:" + c.electricCooker.Run("45分钟")
}

// shutdownCommand 停止指令
type shutdownCommand struct {
  electricCooker *ElectricCooker
}

func NewShutdownCommand(electricCooker *ElectricCooker) *shutdownCommand {
  return &shutdownCommand{
    electricCooker: electricCooker,
  }
}

func (s *shutdownCommand) Execute() string {
  return s.electricCooker.Shutdown()
}

// ElectricCookerInvoker 电饭煲指令触发器
type ElectricCookerInvoker struct {
  cookCommand CookCommand
}

// SetCookCommand 设置指令
func (e *ElectricCookerInvoker) SetCookCommand(cookCommand CookCommand) {
  e.cookCommand = cookCommand
}

// ExecuteCookCommand 执行指令
func (e *ElectricCookerInvoker) ExecuteCookCommand() string {
  return e.cookCommand.Execute()
}
```

### **（五）测试程序**

```go
package command

import (
  "fmt"
  "testing"
)

func TestCommand(t *testing.T) {
  // 创建电饭煲，命令接受者
  electricCooker := new(ElectricCooker)
  // 创建电饭煲指令触发器
  electricCookerInvoker := new(ElectricCookerInvoker)

  // 蒸饭
  steamRiceCommand := NewSteamRiceCommand(electricCooker)
  electricCookerInvoker.SetCookCommand(steamRiceCommand)
  fmt.Println(electricCookerInvoker.ExecuteCookCommand())

  // 煮粥
  cookCongeeCommand := NewCookCongeeCommand(electricCooker)
  electricCookerInvoker.SetCookCommand(cookCongeeCommand)
  fmt.Println(electricCookerInvoker.ExecuteCookCommand())

  // 停止
  shutdownCommand := NewShutdownCommand(electricCooker)
  electricCookerInvoker.SetCookCommand(shutdownCommand)
  fmt.Println(electricCookerInvoker.ExecuteCookCommand())
}
```

### **（六）运行结果** 

```go
=== RUN   TestCommand
蒸饭:电饭煲设置火力为中,压力为正常,持续运行30分钟;
煮粥:电饭煲设置火力为大,压力为强,持续运行45分钟;
电饭煲停止运行。
--- PASS: TestCommand (0.00s)
PASS
```



# ![图片](F:\Images\迭代器模式.png)                			                    **迭代器模式** 

## **（一）概念**

迭代器模式是一种行为设计模式，让你能在不暴露集合底层表现形式 （列表、 栈和树等）的情况下遍历集合中所有的元素。

在迭代器的帮助下， 客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素。

这里需要注意的是有两个典型的迭代器接口需要分清楚；一个是集合类实现的可以创建迭代器的工厂方法接口一般命名为Iterable，包含的方法类似CreateIterator；另一个是迭代器本身的接口，命名为Iterator，有Next及hasMore两个主要方法；



## **（二）示例**

一个班级类中包括一个老师和若干个学生，我们要对班级所有成员进行遍历，班级中老师存储在单独的结构字段中，学生存储在另外一个slice字段中，通过迭代器，我们实现统一遍历处理；



### **（三）班级成员**

```go

package iterator

import "fmt"

// Member 成员接口
type Member interface {
  Desc() string // 输出成员描述信息
}

// Teacher 老师
type Teacher struct {
  name    string // 名称
  subject string // 所教课程
}

// NewTeacher 根据姓名、课程创建老师对象
func NewTeacher(name, subject string) *Teacher {
  return &Teacher{
    name:    name,
    subject: subject,
  }
}

func (t *Teacher) Desc() string {
  return fmt.Sprintf("%s班主任老师负责教%s", t.name, t.subject)
}

// Student 学生
type Student struct {
  name     string // 姓名
  sumScore int    // 考试总分数
}

// NewStudent 创建学生对象
func NewStudent(name string, sumScore int) *Student {
  return &Student{
    name:     name,
    sumScore: sumScore,
  }
}

func (t *Student) Desc() string {
  return fmt.Sprintf("%s同学考试总分为%d", t.name, t.sumScore)
}

```

### **（四）班级成员迭代器**

```go

package iterator

// Iterator 迭代器接口
type Iterator interface {
  Next() Member  // 迭代下一个成员
  HasMore() bool // 是否还有
}

// memberIterator 班级成员迭代器实现
type memberIterator struct {
  class *Class // 需迭代的班级
  index int    // 迭代索引
}

func (m *memberIterator) Next() Member {
  // 迭代索引为-1时，返回老师成员，否则遍历学生slice
  if m.index == -1 {
    m.index++
    return m.class.teacher
  }
  student := m.class.students[m.index]
  m.index++
  return student
}

func (m *memberIterator) HasMore() bool {
  return m.index < len(m.class.students)
}

// Iterable 可迭代集合接口，实现此接口返回迭代器
type Iterable interface {
  CreateIterator() Iterator
}

// Class 班级，包括老师和同学
type Class struct {
  name     string
  teacher  *Teacher
  students []*Student
}

// NewClass 根据班主任老师名称，授课创建班级
func NewClass(name, teacherName, teacherSubject string) *Class {
  return &Class{
    name:    name,
    teacher: NewTeacher(teacherName, teacherSubject),
  }
}

// CreateIterator 创建班级迭代器
func (c *Class) CreateIterator() Iterator {
  return &memberIterator{
    class: c,
    index: -1, // 迭代索引初始化为-1，从老师开始迭代
  }
}

func (c *Class) Name() string {
  return c.name
}

// AddStudent 班级添加同学
func (c *Class) AddStudent(students ...*Student) {
  c.students = append(c.students, students...)
}

```

### **（五）测试程序** 

```go
package iterator

import (
  "fmt"
  "testing"
)

func TestIterator(t *testing.T) {
  class := NewClass("三年级一班", "王明", "数学课")
  class.AddStudent(NewStudent("张三", 389),
    NewStudent("李四", 378),
    NewStudent("王五", 347))

  fmt.Printf("%s成员如下:\n", class.Name())
  classIterator := class.CreateIterator()
  for classIterator.HasMore() {
    member := classIterator.Next()
    fmt.Println(member.Desc())
  }
}
```

### **（六）运行结果**

```go
=== RUN   TestIterator
三年级一班成员如下:
王明班主任老师负责教数学课
张三同学考试总分为389
李四同学考试总分为378
王五同学考试总分为347
--- PASS: TestIterator (0.00s)
PASS
```



# ![图片](F:\Images\中介者模式.png)                			                    **中介者模式**

## **（一）概念**

中介者模式是一种行为设计模式，能让你减少对象之间混乱无序的依赖关系。该模式会限制对象之间的直接交互，迫使它们通过一个中介者对象进行合作，将网状依赖变为星状依赖。

中介者能使得程序更易于修改和扩展，而且能更方便地对独立的组件进行复用，因为它们不再依赖于很多其他的类。

中介者模式与观察者模式之间的区别是，中介者模式解决的是同类或者不同类的多个对象之间多对多的依赖关系，观察者模式解决的是多个对象与一个对象之间的多对一的依赖关系。



## **（二）示例**

机场塔台调度系统是一个体现中介者模式的典型示例，假设是一个小机场，每次只能同时允许一架飞机起降，每架靠近机场的飞机需要先与塔台沟通是否可以降落，如果没有空闲的跑道，需要在天空盘旋等待，如果有飞机离港，等待的飞机会收到塔台的通知，按先后顺序降落；这种方式，免去多架飞机同时到达机场需要相互沟通降落顺序的复杂性，减少多个飞机间的依赖关系，简化业务逻辑，从而降低系统出问题的风险。



### **（三）飞机对象**

```go
package mediator

import "fmt"

// Aircraft 飞机接口
type Aircraft interface {
  ApproachAirport() // 抵达机场空域
  DepartAirport()   // 飞离机场
}

// airliner 客机
type airliner struct {
  name            string          // 客机型号
  airportMediator AirportMediator // 机场调度
}

// NewAirliner 根据指定型号及机场调度创建客机
func NewAirliner(name string, mediator AirportMediator) *airliner {
  return &airliner{
    name:            name,
    airportMediator: mediator,
  }
}

func (a *airliner) ApproachAirport() {
  if !a.airportMediator.CanLandAirport(a) { // 请求塔台是否可以降落
    fmt.Printf("机场繁忙，客机%s继续等待降落;\n", a.name)
    return
  }
  fmt.Printf("客机%s成功滑翔降落机场;\n", a.name)
}

func (a *airliner) DepartAirport() {
  fmt.Printf("客机%s成功滑翔起飞，离开机场;\n", a.name)
  a.airportMediator.NotifyWaitingAircraft() // 通知等待的其他飞机
}

// helicopter 直升机
type helicopter struct {
  name            string
  airportMediator AirportMediator
}

// NewHelicopter 根据指定型号及机场调度创建直升机
func NewHelicopter(name string, mediator AirportMediator) *helicopter {
  return &helicopter{
    name:            name,
    airportMediator: mediator,
  }
}

func (h *helicopter) ApproachAirport() {
  if !h.airportMediator.CanLandAirport(h) { // 请求塔台是否可以降落
    fmt.Printf("机场繁忙，直升机%s继续等待降落;\n", h.name)
    return
  }
  fmt.Printf("直升机%s成功垂直降落机场;\n", h.name)
}

func (h *helicopter) DepartAirport() {
  fmt.Printf("直升机%s成功垂直起飞，离开机场;\n", h.name)
  h.airportMediator.NotifyWaitingAircraft() // 通知其他等待降落的飞机
}

```

### **（四）机场塔台**

```go
package mediator

// AirportMediator 机场调度中介者
type AirportMediator interface {
  CanLandAirport(aircraft Aircraft) bool // 确认是否可以降落
  NotifyWaitingAircraft()                // 通知等待降落的其他飞机
}

// ApproachTower 机场塔台
type ApproachTower struct {
  hasFreeAirstrip bool
  waitingQueue    []Aircraft // 等待降落的飞机队列
}

func (a *ApproachTower) CanLandAirport(aircraft Aircraft) bool {
  if a.hasFreeAirstrip {
    a.hasFreeAirstrip = false
    return true
  }
  // 没有空余的跑道，加入等待队列
  a.waitingQueue = append(a.waitingQueue, aircraft)
  return false
}

func (a *ApproachTower) NotifyWaitingAircraft() {
  if !a.hasFreeAirstrip {
    a.hasFreeAirstrip = true
  }
  if len(a.waitingQueue) > 0 {
    // 如果存在等待降落的飞机，通知第一个降落
    first := a.waitingQueue[0]
    a.waitingQueue = a.waitingQueue[1:]
    first.ApproachAirport()
  }
}

```

### **（五）测试程序**

```go
package mediator

import "testing"

func TestMediator(t *testing.T) {
  // 创建机场调度塔台
  airportMediator := &ApproachTower{hasFreeAirstrip: true}
  // 创建C919客机
  c919Airliner := NewAirliner("C919", airportMediator)
  // 创建米-26重型运输直升机
  m26Helicopter := NewHelicopter("米-26", airportMediator)

  c919Airliner.ApproachAirport()  // c919进港降落
  m26Helicopter.ApproachAirport() // 米-26进港等待

  c919Airliner.DepartAirport()  // c919飞离，等待的米-26进港降落
  m26Helicopter.DepartAirport() // 最后米-26飞离
}
```

### **（六）运行结果**

```go
=== RUN   TestMediator
客机C919成功滑翔降落机场;
机场繁忙，直升机米-26继续等待降落;
客机C919成功滑翔起飞，离开机场;
直升机米-26成功垂直降落机场;
直升机米-26成功垂直起飞，离开机场;
--- PASS: TestMediator (0.00s)
PASS
```



# ![图片](F:\Images\备忘录模式.png)                			                    **备忘录模式**

## **（一）概念**

备忘录模式是一种行为设计模式， 允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。

备忘录不会影响它所处理的对象的内部结构， 也不会影响快照中保存的数据。

一般情况由原发对象保存生成的备忘录对象的状态不能被除原发对象之外的对象访问，所以通过内部类定义具体的备忘录对象是比较安全的，但是go语言不支持内部类定义的方式，因此go语言实现备忘录对象时，首先将备忘录保存的状态设为非导出字段，避免外部对象访问，其次将原发对象的引用保存到备忘录对象中，当通过备忘录对象恢复时，直接操作备忘录的恢复方法，将备份数据状态设置到原发对象中，完成恢复。



## **（二）示例**

大家平时玩的角色扮演闯关游戏的存档机制就可以通过备忘录模式实现，每到一个关键关卡，玩家经常会先保存游戏存档，用于闯关失败后重置，存档会把角色状态及场景状态保存到备忘录中，同时将需要恢复游戏的引用存入备忘录，用于关卡重置；



### **（三）闯关游戏**

```go
package memento

import "fmt"

// Originator 备忘录模式原发器接口
type Originator interface {
  Save(tag string) Memento // 当前状态保存备忘录
}

// RolesPlayGame 支持存档的RPG游戏
type RolesPlayGame struct {
  name          string   // 游戏名称
  rolesState    []string // 游戏角色状态
  scenarioState string   // 游戏场景状态
}

// NewRolesPlayGame 根据游戏名称和角色名，创建RPG游戏
func NewRolesPlayGame(name string, roleName string) *RolesPlayGame {
  return &RolesPlayGame{
    name:          name,
    rolesState:    []string{roleName, "血量100"}, // 默认满血
    scenarioState: "开始通过第一关",                   // 默认第一关开始
  }
}

// Save 保存RPG游戏角色状态及场景状态到指定标签归档
func (r *RolesPlayGame) Save(tag string) Memento {
  return newRPGArchive(tag, r.rolesState, r.scenarioState, r)
}

func (r *RolesPlayGame) SetRolesState(rolesState []string) {
  r.rolesState = rolesState
}

func (r *RolesPlayGame) SetScenarioState(scenarioState string) {
  r.scenarioState = scenarioState
}

// String 输出RPG游戏简要信息
func (r *RolesPlayGame) String() string {
  return fmt.Sprintf("在%s游戏中，玩家使用%s,%s,%s;", r.name, r.rolesState[0], r.rolesState[1], r.scenarioState)
}
```

### **（四）游戏存档**

```go
package memento

import "fmt"

// Memento 备忘录接口
type Memento interface {
  Tag() string // 备忘录标签
  Restore()    // 根据备忘录存储数据状态恢复原对象
}

// rpgArchive rpg游戏存档，
type rpgArchive struct {
  tag           string         // 存档标签
  rolesState    []string       // 存档的角色状态
  scenarioState string         // 存档游戏场景状态
  rpg           *RolesPlayGame // rpg游戏引用
}

// newRPGArchive 根据标签，角色状态，场景状态，rpg游戏引用，创建游戏归档备忘录
func newRPGArchive(tag string, rolesState []string, scenarioState string, rpg *RolesPlayGame) *rpgArchive {
  return &rpgArchive{
    tag:           tag,
    rolesState:    rolesState,
    scenarioState: scenarioState,
    rpg:           rpg,
  }
}

func (r *rpgArchive) Tag() string {
  return r.tag
}

// Restore 根据归档数据恢复游戏状态
func (r *rpgArchive) Restore() {
  r.rpg.SetRolesState(r.rolesState)
  r.rpg.SetScenarioState(r.scenarioState)
}

// RPGArchiveManager RPG游戏归档管理器
type RPGArchiveManager struct {
  archives map[string]Memento // 存储归档标签对应归档
}

func NewRPGArchiveManager() *RPGArchiveManager {
  return &RPGArchiveManager{
    archives: make(map[string]Memento),
  }
}

// Reload 根据标签重新加载归档数据
func (r *RPGArchiveManager) Reload(tag string) {
  if archive, ok := r.archives[tag]; ok {
    fmt.Printf("重新加载%s;\n", tag)
    archive.Restore()
  }
}

// Put 保存归档数据
func (r *RPGArchiveManager) Put(memento Memento) {
  r.archives[memento.Tag()] = memento
}
```

### **（五）测试程序**

```go
package memento

import (
  "fmt"
  "testing"
)

func TestMemento(t *testing.T) {
  // 创建RPG游戏存档管理器
  rpgManager := NewRPGArchiveManager()
  // 创建RPG游戏
  rpg := NewRolesPlayGame("暗黑破坏神2", "野蛮人战士")
  fmt.Println(rpg)                  // 输出游戏当前状态
  rpgManager.Put(rpg.Save("第一关存档")) // 游戏存档

  // 第一关闯关失败
  rpg.SetRolesState([]string{"野蛮人战士", "死亡"})
  rpg.SetScenarioState("第一关闯关失败")
  fmt.Println(rpg)

  // 恢复存档，重新闯关
  rpgManager.Reload("第一关存档")
  fmt.Println(rpg)
}
```

### **（六）运行结果**

```go
=== RUN   TestMemento
在暗黑破坏神2游戏中，玩家使用野蛮人战士,血量100,开始通过第一关;
在暗黑破坏神2游戏中，玩家使用野蛮人战士,死亡,第一关闯关失败;
重新加载第一关存档;
在暗黑破坏神2游戏中，玩家使用野蛮人战士,血量100,开始通过第一关;
--- PASS: TestMemento (0.00s)
PASS
```



# ![图片](F:\Images\观察者模式.png)                			                    **观察者模式**

## **（一）概念**

观察者模式是一种行为设计模式，允许你定义一种订阅机制，可在对象事件发生时通知多个 “观察” 该对象的其他对象。

观察者模式提供了一种作用于任何实现了订阅者接口的对象的机制，可对其事件进行订阅和取消订阅。

观察者模式是最常用的模式之一，是事件总线，分布式消息中间件等各种事件机制的原始理论基础，常用于解耦多对一的对象依赖关系；

增强的实现功能包括：

- 当被观察者通过异步实现通知多个观察者时就相当于单进程实例的消息总线；

- 同时还可以根据业务需要，将被观察者所有数据状态变更进行分类为不同的主题，观察者通过不同主题进行订阅；

- 同一个主题又可分为增加，删除，修改事件行为；

- 每个主题可以实现一个线程池，多个主题通过不同的线程池进行处理隔离，线程池可以设置并发线程大小、缓冲区大小及调度策略，比如先进先出，优先级等策略；

- 观察者处理事件时有可能出现异常，所以也可以注册异常处理函数，异常处理也可以通过异常类型进行分类；
- 根据业务需求也可以实现通知异常重试，延迟通知等功能；



**（二）示例**

信用卡业务消息提醒可通过观察者模式实现，业务消息包括日常消费，出账单，账单逾期，消息提醒包括短信、邮件及电话，根据不同业务的场景会采用不同的消息提醒方式或者多种消息提醒方式，这里信用卡相当于被观察者，观察者相当于不同的通知方式；日常消费通过短信通知，出账单通过邮件通知，账单逾期三种方式都会进行通知；



### **（三）通知方式**

```go
package observer

import "fmt"

// Subscriber 订阅者接口
type Subscriber interface {
  Name() string          //订阅者名称
  Update(message string) //订阅更新方法
}

// shortMessage 信用卡消息短信订阅者
type shortMessage struct{}

func (s *shortMessage) Name() string {
  return "手机短息"
}

func (s *shortMessage) Update(message string) {
  fmt.Printf("通过【%s】发送消息:%s\n", s.Name(), message)
}

// email 信用卡消息邮箱订阅者
type email struct{}

func (e *email) Name() string {
  return "电子邮件"
}

func (e *email) Update(message string) {
  fmt.Printf("通过【%s】发送消息:%s\n", e.Name(), message)
}

// telephone 信用卡消息电话订阅者
type telephone struct{}

func (t *telephone) Name() string {
  return "电话"
}

func (t *telephone) Update(message string) {
  fmt.Printf("通过【%s】告知:%s\n", t.Name(), message)
}
```

### **（四）信用卡业务**

```go
package observer

import "fmt"

// Subscriber 订阅者接口
type Subscriber interface {
  Name() string          //订阅者名称
  Update(message string) //订阅更新方法
}

// shortMessage 信用卡消息短信订阅者
type shortMessage struct{}

func (s *shortMessage) Name() string {
  return "手机短息"
}

func (s *shortMessage) Update(message string) {
  fmt.Printf("通过【%s】发送消息:%s\n", s.Name(), message)
}

// email 信用卡消息邮箱订阅者
type email struct{}

func (e *email) Name() string {
  return "电子邮件"
}

func (e *email) Update(message string) {
  fmt.Printf("通过【%s】发送消息:%s\n", e.Name(), message)
}

// telephone 信用卡消息电话订阅者
type telephone struct{}

func (t *telephone) Name() string {
  return "电话"
}

func (t *telephone) Update(message string) {
  fmt.Printf("通过【%s】告知:%s\n", t.Name(), message)
}
```

### **（五）测试程序**

```go

package observer

import "testing"

func TestObserver(t *testing.T) {
  // 创建张三的信用卡
  creditCard := NewCreditCard("张三")
  // 短信通知订阅信用卡消费及逾期消息
  creditCard.Subscribe(new(shortMessage), ConsumeType, ExpireType)
  // 电子邮件通知订阅信用卡账单及逾期消息
  creditCard.Subscribe(new(email), BillType, ExpireType)
  // 电话通知订阅信用卡逾期消息，同时逾期消息通过三种方式通知
  creditCard.Subscribe(new(telephone), ExpireType)

  creditCard.Consume(500.00) // 信用卡消费
  creditCard.Consume(800.00) // 信用卡消费
  creditCard.SendBill()      // 信用卡发送账单
  creditCard.Expire()        // 信用卡逾期

  // 信用卡逾期消息取消电子邮件及短信通知订阅
  creditCard.Unsubscribe(new(email), ExpireType)
  creditCard.Unsubscribe(new(shortMessage), ExpireType)
  creditCard.Consume(300.00) // 信用卡消费
  creditCard.Expire()        // 信用卡逾期
}
```

### **（六）运行结果**

```go
=== RUN   TestObserver
通过【手机短息】发送消息:尊敬的持卡人张三,您当前消费500.00元;
通过【手机短息】发送消息:尊敬的持卡人张三,您当前消费800.00元;
通过【电子邮件】发送消息:尊敬的持卡人张三,您本月账单已出，消费总额1300.00元;
通过【手机短息】发送消息:尊敬的持卡人张三,您本月账单已逾期，请及时还款，总额1300.00元;
通过【电子邮件】发送消息:尊敬的持卡人张三,您本月账单已逾期，请及时还款，总额1300.00元;
通过【电话】告知:尊敬的持卡人张三,您本月账单已逾期，请及时还款，总额1300.00元;
通过【手机短息】发送消息:尊敬的持卡人张三,您当前消费300.00元;
通过【电话】告知:尊敬的持卡人张三,您本月账单已逾期，请及时还款，总额1600.00元;
--- PASS: TestObserver (0.00s)
PASS
```



# ![图片](F:\Images\状态模式.png)                			                    **状态模式**

# **（一）概念**

状态模式是一种行为设计模式，让你能在一个对象的内部状态变化时改变其行为，使其看上去就像改变了自身所属的类一样。

该模式将与状态相关的行为抽取到独立的状态类中，让原对象将工作委派给这些类的实例，而不是自行进行处理。

状态迁移有四个元素组成，起始状态、触发迁移的事件，终止状态以及要执行的动作，每个具体的状态包含触发状态迁移的执行方法，迁移方法的实现是执行持有状态对象的动作方法，同时设置状态为下一个流转状态；持有状态的业务对象包含有触发状态迁移方法，这些迁移方法将请求委托给当前具体状态对象的迁移方法。



## **（二）示例**

IPhone手机充电就是一个手机电池状态的流转，一开始手机处于有电状态，插入充电插头后，继续充电到满电状态，并进入断电保护，拔出充电插头后使用手机，由满电逐渐变为没电，最终关机；

状态迁移表：

| 起始状态 | 触发事件   | 终止状态 | 执行动作 |
| -------- | ---------- | -------- | -------- |
| 有电     | 插入充电线 | 满电     | 充电     |
| 有电     | 拔出充电线 | 没电     | 耗电     |
| 满电     | 插入充电线 | 满电     | 停止充电 |
| 满电     | 拔出充电线 | 有电     | 耗电     |
| 没电     | 插入充电线 | 有电     | 充电     |
| 没电     | 拔出充电线 | 没电     | 关机     |



### **（三）电池状态**

```go

package state

import "fmt"

// BatteryState 电池状态接口，支持手机充电线插拔事件
type BatteryState interface {
  ConnectPlug(iPhone *IPhone) string
  DisconnectPlug(iPhone *IPhone) string
}

// fullBatteryState 满电状态
type fullBatteryState struct{}

func (s *fullBatteryState) String() string {
  return "满电状态"
}

func (s *fullBatteryState) ConnectPlug(iPhone *IPhone) string {
  return iPhone.pauseCharge()
}

func (s *fullBatteryState) DisconnectPlug(iPhone *IPhone) string {
  iPhone.SetBatteryState(PartBatteryState)
  return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, PartBatteryState)
}

// emptyBatteryState 空电状态
type emptyBatteryState struct{}

func (s *emptyBatteryState) String() string {
  return "没电状态"
}

func (s *emptyBatteryState) ConnectPlug(iPhone *IPhone) string {
  iPhone.SetBatteryState(PartBatteryState)
  return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, PartBatteryState)
}

func (s *emptyBatteryState) DisconnectPlug(iPhone *IPhone) string {
  return iPhone.shutdown()
}

// partBatteryState 部分电状态
type partBatteryState struct{}

func (s *partBatteryState) String() string {
  return "有电状态"
}

func (s *partBatteryState) ConnectPlug(iPhone *IPhone) string {
  iPhone.SetBatteryState(FullBatteryState)
  return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, FullBatteryState)
}

func (s *partBatteryState) DisconnectPlug(iPhone *IPhone) string {
  iPhone.SetBatteryState(EmptyBatteryState)
  return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, EmptyBatteryState)
}
```

### **（四）IPhone手机**

```go
package state

import "fmt"

// 电池状态单例，全局统一使用三个状态的单例，不需要重复创建
var (
  FullBatteryState  = new(fullBatteryState)  // 满电
  EmptyBatteryState = new(emptyBatteryState) // 空电
  PartBatteryState  = new(partBatteryState)  // 部分电
)

// IPhone 已手机充电为例，实现状态模式
type IPhone struct {
  model        string       // 手机型号
  batteryState BatteryState // 电池状态
}

// NewIPhone 创建指定型号手机
func NewIPhone(model string) *IPhone {
  return &IPhone{
    model:        model,
    batteryState: PartBatteryState,
  }
}

// BatteryState 输出电池当前状态
func (i *IPhone) BatteryState() string {
  return fmt.Sprintf("iPhone %s 当前为%s", i.model, i.batteryState)
}

// ConnectPlug 连接充电线
func (i *IPhone) ConnectPlug() string {
  return fmt.Sprintf("iPhone %s 连接电源线,%s", i.model, i.batteryState.ConnectPlug(i))
}

// DisconnectPlug 断开充电线
func (i *IPhone) DisconnectPlug() string {
  return fmt.Sprintf("iPhone %s 断开电源线,%s", i.model, i.batteryState.DisconnectPlug(i))
}

// SetBatteryState 设置电池状态
func (i *IPhone) SetBatteryState(state BatteryState) {
  i.batteryState = state
}

func (i *IPhone) charge() string {
  return "正在充电"
}

func (i *IPhone) pauseCharge() string {
  return "电已满,暂停充电"
}

func (i *IPhone) shutdown() string {
  return "手机关闭"
}

func (i *IPhone) consume() string {
  return "使用中,消耗电量"
}
```

### **（五）测试程序**

```go
package state

import (
  "fmt"
  "testing"
)

func TestState(t *testing.T) {
  iPhone13Pro := NewIPhone("13 pro") // 刚创建的手机有部分电

  fmt.Println(iPhone13Pro.BatteryState()) // 打印部分电状态
  fmt.Println(iPhone13Pro.ConnectPlug())  // 插上电源插头，继续充满电
  fmt.Println(iPhone13Pro.ConnectPlug())  // 满电后再充电，会触发满电保护

  fmt.Println(iPhone13Pro.DisconnectPlug()) // 拔掉电源，使用手机消耗电量，变为有部分电
  fmt.Println(iPhone13Pro.DisconnectPlug()) // 一直使用手机，直到没电
  fmt.Println(iPhone13Pro.DisconnectPlug()) // 没电后会关机

  fmt.Println(iPhone13Pro.ConnectPlug()) // 再次插上电源一会，变为有电状态
}
```

### **（六）运行结果**

```go
=== RUN   TestState
iPhone 13 pro 当前为有电状态
iPhone 13 pro 连接电源线,正在充电,有电状态转为满电状态
iPhone 13 pro 连接电源线,电已满,暂停充电
iPhone 13 pro 断开电源线,使用中,消耗电量,满电状态转为有电状态
iPhone 13 pro 断开电源线,使用中,消耗电量,有电状态转为没电状态
iPhone 13 pro 断开电源线,手机关闭
iPhone 13 pro 连接电源线,正在充电,没电状态转为有电状态
--- PASS: TestState (0.00s)
PASS
```



# ![图片](F:\Images\策略模式.png)                			                    **策略模式**

## **（一）概念**

策略模式是一种行为设计模式，它能让你定义一系列算法，并将每种算法分别放入独立的类中，以使算法的对象能够相互替换。

原始对象被称为上下文，它包含指向策略对象的引用并将执行行为的任务分派给策略对象。为了改变上下文完成其工作的方式，其他对象可以使用另一个对象来替换当前链接的策略对象。

策略模式是最常用的设计模式，也是比较简单的设计模式，是以多态替换条件表达式重构方法的具体实现，是面向接口编程原则的最直接体现；



## **（二）示例**

北京是一个四季分明的城市，每个季节天气情况都有明显特点；我们定义一个显示天气情况的季节接口，具体的四季实现，都会保存一个城市和天气情况的映射表，城市对象会包含季节接口，随着四季的变化，天气情况也随之变化；



### **（三）四季天气**

```go
package strategy

import "fmt"

// Season 季节的策略接口，不同季节表现得天气不同
type Season interface {
  ShowWeather(city string) string // 显示指定城市的天气情况
}

type spring struct {
  weathers map[string]string // 存储不同城市春天气候
}

func NewSpring() *spring {
  return &spring{
    weathers: map[string]string{"北京": "干燥多风", "昆明": "清凉舒适"},
  }
}

func (s *spring) ShowWeather(city string) string {
  return fmt.Sprintf("%s的春天，%s;", city, s.weathers[city])
}

type summer struct {
  weathers map[string]string // 存储不同城市夏天气候
}

func NewSummer() *summer {
  return &summer{
    weathers: map[string]string{"北京": "高温多雨", "昆明": "清凉舒适"},
  }
}

func (s *summer) ShowWeather(city string) string {
  return fmt.Sprintf("%s的夏天，%s;", city, s.weathers[city])
}

type autumn struct {
  weathers map[string]string // 存储不同城市秋天气候
}

func NewAutumn() *autumn {
  return &autumn{
    weathers: map[string]string{"北京": "凉爽舒适", "昆明": "清凉舒适"},
  }
}

func (a *autumn) ShowWeather(city string) string {
  return fmt.Sprintf("%s的秋天，%s;", city, a.weathers[city])
}

type winter struct {
  weathers map[string]string // 存储不同城市冬天气候
}

func NewWinter() *winter {
  return &winter{
    weathers: map[string]string{"北京": "干燥寒冷", "昆明": "清凉舒适"},
  }
}

func (w *winter) ShowWeather(city string) string {
  return fmt.Sprintf("%s的冬天，%s;", city, w.weathers[city])
}
```

### **（四）城市气候**

```go

package strategy

import (
  "fmt"
)

// City 城市
type City struct {
  name    string
  feature string
  season  Season
}

// NewCity 根据名称及季候特征创建城市
func NewCity(name, feature string) *City {
  return &City{
    name:    name,
    feature: feature,
  }
}

// SetSeason 设置不同季节，类似天气在不同季节的不同策略
func (c *City) SetSeason(season Season) {
  c.season = season
}

// String 显示城市的气候信息
func (c *City) String() string {
  return fmt.Sprintf("%s%s，%s", c.name, c.feature, c.season.ShowWeather(c.name))
}
```

### **（五）测试程序**

```go
package strategy

import (
  "fmt"
  "testing"
)

func TestStrategy(t *testing.T) {
  Beijing := NewCity("北京", "四季分明")

  Beijing.SetSeason(NewSpring())
  fmt.Println(Beijing)

  Beijing.SetSeason(NewSummer())
  fmt.Println(Beijing)

  Beijing.SetSeason(NewAutumn())
  fmt.Println(Beijing)

  Beijing.SetSeason(NewWinter())
  fmt.Println(Beijing)
}
```

### **（六）运行结果**

```go
=== RUN   TestStrategy
北京四季分明，北京的春天，干燥多风;
北京四季分明，北京的夏天，高温多雨;
北京四季分明，北京的秋天，凉爽舒适;
北京四季分明，北京的冬天，干燥寒冷;
--- PASS: TestStrategy (0.00s)
PASS
```



# ![图片](F:\Images\模板方法模式.png)                			                    **模板方法模式**

## **（一）概念**

模板方法模式是一种行为设计模式，它在超类中定义了一个算法的框架，允许子类在不修改结构的情况下重写算法的特定步骤。

由于GO语言没有继承的语法，模板方法又是依赖继承实现的设计模式，因此GO语言实现模板方法比较困难， GO语言支持隐式内嵌字段“继承”其他结构体的字段与方法，但是这个并不是真正意义上的继承语法，外层结构重写隐式字段中的算法特定步骤后，无法动态绑定到“继承”过来的算法的框架方法调用中，因此不能实现模板方法模式的语义。



## **（二）示例**

本示例给出一种间接实现模板方法的方式，也比较符合模板方法模式的定义：

- 将多个算法特定步骤组合成一个接口；

- 基类隐式内嵌算法步骤接口，同时调用算法步骤接口的各方法，实现算法的模板方法，此时基类内嵌的算法步骤接口并没有真正的处理行为；

- 子类隐式内嵌基类，并覆写算法步骤接口的方法；

- 通过工厂方法创建具体子类，并将自己的引用赋值给基类中算法步骤接口字段；

  

以演员装扮为例，演员的装扮是分为化妆，穿衣，配饰三步骤，三个步骤又根据不同角色的演员有所差别，因此演员基类实现装扮的模板方法，对于化妆，穿衣，配饰的三个步骤，在子类演员中具体实现，子类具体演员分为，男演员、女演员和儿童演员；



### **（三）演员基类**

```go
package templatemethod

import (
  "bytes"
  "fmt"
)

// IActor 演员接口
type IActor interface {
  DressUp() string // 装扮
}

// dressBehavior 装扮的多个行为，这里多个行为是私有的，通过DressUp模版方法调用
type dressBehavior interface {
  makeUp() string // 化妆
  clothe() string // 穿衣
  wear() string   // 配饰
}

// BaseActor 演员基类
type BaseActor struct {
  roleName      string // 扮演角色
  dressBehavior        // 装扮行为
}

// DressUp 统一实现演员接口的DressUp模版方法，装扮过程通过不同装扮行为进行扩展
func (b *BaseActor) DressUp() string {
  buf := bytes.Buffer{}
  buf.WriteString(fmt.Sprintf("扮演%s的", b.roleName))
  buf.WriteString(b.makeUp())
  buf.WriteString(b.clothe())
  buf.WriteString(b.wear())
  return buf.String()
}
```

### **（四）具体演员**

```go
package templatemethod

// womanActor 扩展装扮行为的女演员
type womanActor struct {
  BaseActor
}

// NewWomanActor 指定角色创建女演员
func NewWomanActor(roleName string) *womanActor {
  actor := new(womanActor)    // 创建女演员
  actor.roleName = roleName   // 设置角色
  actor.dressBehavior = actor // 将女演员实现的扩展装扮行为，设置给自己的装扮行为接口
  return actor
}

// 化妆
func (w *womanActor) makeUp() string {
  return "女演员涂着口红，画着眉毛；"
}

// 穿衣
func (w *womanActor) clothe() string {
  return "穿着连衣裙；"
}

// 配饰
func (w *womanActor) wear() string {
  return "带着耳环，手拎着包；"
}

// manActor 扩展装扮行为的男演员
type manActor struct {
  BaseActor
}

func NewManActor(roleName string) *manActor {
  actor := new(manActor)
  actor.roleName = roleName
  actor.dressBehavior = actor // 将男演员实现的扩展装扮行为，设置给自己的装扮行为接口
  return actor
}

func (m *manActor) makeUp() string {
  return "男演员刮净胡子，抹上发胶；"
}

func (m *manActor) clothe() string {
  return "穿着一身西装；"
}

func (m *manActor) wear() string {
  return "带上手表，抽着烟；"
}

// NewChildActor 扩展装扮行为的儿童演员
type childActor struct {
  BaseActor
}

func NewChildActor(roleName string) *childActor {
  actor := new(childActor)
  actor.roleName = roleName
  actor.dressBehavior = actor // 将儿童演员实现的扩展装扮行为，设置给自己的装扮行为接口
  return actor
}

func (c *childActor) makeUp() string {
  return "儿童演员抹上红脸蛋；"
}

func (c *childActor) clothe() string {
  return "穿着一身童装；"
}

func (c *childActor) wear() string {
  return "手里拿着一串糖葫芦；"
}
```

### **（五）测试程序**

```go
package templatemethod

import (
  "fmt"
  "testing"
)

func TestTemplateMethod(t *testing.T) {
  showActors(NewWomanActor("妈妈"), NewManActor("爸爸"), NewChildActor("儿子"))
}

// showActors 显示演员的装扮信息
func showActors(actors ...IActor) {
  for _, actor := range actors {
    fmt.Println(actor.DressUp())
  }
}
```

### **（六）运行结果**

```go
=== RUN   TestTemplateMethod
扮演妈妈的女演员涂着口红，画着眉毛；穿着连衣裙；带着耳环，手拎着包；
扮演爸爸的男演员刮净胡子，抹上发胶；穿着一身西装；带上手表，抽着烟；
扮演儿子的儿童演员抹上红脸蛋；穿着一身童装；手里拿着一串糖葫芦；
--- PASS: TestTemplateMethod (0.00s)
PASS
```



# ![图片](F:\Images\访问者模式.png)                			                    **访问者模式**

## **（一）概念**

访问者模式是一种行为设计模式，它能将算法与其所作用的对象隔离开来。允许你在不修改已有代码的情况下向已有类层次结构中增加新的行为。

访问者接口需要根据被访问者具体类，定义多个相似的访问方法，每个具体类对应一个访问方法；每个被访问者需要实现一个接受访问者对象的方法，方法的实现就是去调用访问者接口对应该类的访问方法；这个接受方法可以传入不同目的访问者接口的具体实现，从而在不修改被访问对象的前提下，增加新的功能；



## **（二）示例**

公司中存在多种类型的员工，包括产品经理、软件工程师、人力资源等，他们的KPI指标不尽相同，产品经理为上线产品数量及满意度，软件工程师为实现的需求数及修改bug数，人力资源为招聘员工的数量；公司要根据员工完成的KPI进行表彰公示，同时根据KPI完成情况定薪酬，这些功能都是员工类职责之外的，不能修改员工本身的类，我们通过访问者模式，实现KPI表彰排名及薪酬发放；



### **（三）员工结构**

```go
package visitor

import "fmt"

// Employee 员工接口
type Employee interface {
  KPI() string                    // 完成kpi信息
  Accept(visitor EmployeeVisitor) // 接受访问者对象
}

// productManager 产品经理
type productManager struct {
  name         string // 名称
  productNum   int    // 上线产品数
  satisfaction int    // 平均满意度
}

func NewProductManager(name string, productNum int, satisfaction int) *productManager {
  return &productManager{
    name:         name,
    productNum:   productNum,
    satisfaction: satisfaction,
  }
}

func (p *productManager) KPI() string {
  return fmt.Sprintf("产品经理%s，上线%d个产品，平均满意度为%d", p.name, p.productNum, p.satisfaction)
}

func (p *productManager) Accept(visitor EmployeeVisitor) {
  visitor.VisitProductManager(p)
}

// softwareEngineer 软件工程师
type softwareEngineer struct {
  name           string // 姓名
  requirementNum int    // 完成需求数
  bugNum         int    // 修复问题数
}

func NewSoftwareEngineer(name string, requirementNum int, bugNum int) *softwareEngineer {
  return &softwareEngineer{
    name:           name,
    requirementNum: requirementNum,
    bugNum:         bugNum,
  }
}

func (s *softwareEngineer) KPI() string {
  return fmt.Sprintf("软件工程师%s，完成%d个需求，修复%d个问题", s.name, s.requirementNum, s.bugNum)
}

func (s *softwareEngineer) Accept(visitor EmployeeVisitor) {
  visitor.VisitSoftwareEngineer(s)
}

// hr 人力资源
type hr struct {
  name       string // 姓名
  recruitNum int    // 招聘人数
}

func NewHR(name string, recruitNum int) *hr {
  return &hr{
    name:       name,
    recruitNum: recruitNum,
  }
}

func (h *hr) KPI() string {
  return fmt.Sprintf("人力资源%s，招聘%d名员工", h.name, h.recruitNum)
}

func (h *hr) Accept(visitor EmployeeVisitor) {
  visitor.VisitHR(h)
}
```

### **（四）员工访问者**

```go
package visitor

import (
  "fmt"
  "sort"
)

// EmployeeVisitor 员工访问者接口
type EmployeeVisitor interface {
  VisitProductManager(pm *productManager)     // 访问产品经理
  VisitSoftwareEngineer(se *softwareEngineer) // 访问软件工程师
  VisitHR(hr *hr)                             // 访问人力资源
}

// kpi kpi对象
type kpi struct {
  name string // 完成kpi姓名
  sum  int    // 完成kpi总数量
}

// kpiTopVisitor 员工kpi排名访问者
type kpiTopVisitor struct {
  top []*kpi
}

func (k *kpiTopVisitor) VisitProductManager(pm *productManager) {
  k.top = append(k.top, &kpi{
    name: pm.name,
    sum:  pm.productNum + pm.satisfaction,
  })
}

func (k *kpiTopVisitor) VisitSoftwareEngineer(se *softwareEngineer) {
  k.top = append(k.top, &kpi{
    name: se.name,
    sum:  se.requirementNum + se.bugNum,
  })
}

func (k *kpiTopVisitor) VisitHR(hr *hr) {
  k.top = append(k.top, &kpi{
    name: hr.name,
    sum:  hr.recruitNum,
  })
}

// Publish 发布KPI排行榜
func (k *kpiTopVisitor) Publish() {
  sort.Slice(k.top, func(i, j int) bool {
    return k.top[i].sum > k.top[j].sum
  })
  for i, curKPI := range k.top {
    fmt.Printf("第%d名%s：完成KPI总数%d\n", i+1, curKPI.name, curKPI.sum)
  }
}

// salaryVisitor 薪酬访问者
type salaryVisitor struct{}

func (s *salaryVisitor) VisitProductManager(pm *productManager) {
  fmt.Printf("产品经理基本薪资：1000元，KPI单位薪资：100元，")
  fmt.Printf("%s，总工资为%d元\n", pm.KPI(), (pm.productNum+pm.satisfaction)*100+1000)
}

func (s *salaryVisitor) VisitSoftwareEngineer(se *softwareEngineer) {
  fmt.Printf("软件工程师基本薪资：1500元，KPI单位薪资：80元，")
  fmt.Printf("%s，总工资为%d元\n", se.KPI(), (se.requirementNum+se.bugNum)*80+1500)
}

func (s *salaryVisitor) VisitHR(hr *hr) {
  fmt.Printf("人力资源基本薪资：800元，KPI单位薪资：120元，")
  fmt.Printf("%s，总工资为%d元\n", hr.KPI(), hr.recruitNum*120+800)
}
```

### **（五）测试程序**

```go
package visitor

import "testing"

func TestVisitor(t *testing.T) {
  allEmployees := AllEmployees() // 获取所有员工
  kpiTop := new(kpiTopVisitor)   // 创建KPI排行访问者
  VisitAllEmployees(kpiTop, allEmployees)
  kpiTop.Publish() // 发布排行榜

  salary := new(salaryVisitor) // 创建薪酬访问者
  VisitAllEmployees(salary, allEmployees)
}

// VisitAllEmployees 遍历所有员工调用访问者
func VisitAllEmployees(visitor EmployeeVisitor, allEmployees []Employee) {
  for _, employee := range allEmployees {
    employee.Accept(visitor)
  }
}

// AllEmployees 获得所有公司员工
func AllEmployees() []Employee {
  var employees []Employee
  employees = append(employees, NewHR("小明", 10))
  employees = append(employees, NewProductManager("小红", 4, 7))
  employees = append(employees, NewSoftwareEngineer("张三", 10, 5))
  employees = append(employees, NewSoftwareEngineer("李四", 3, 6))
  employees = append(employees, NewSoftwareEngineer("王五", 7, 1))
  return employees
}
```

### **（六）运行结果**

```go
=== RUN   TestVisitor
第1名张三：完成KPI总数15
第2名小红：完成KPI总数11
第3名小明：完成KPI总数10
第4名李四：完成KPI总数9
第5名王五：完成KPI总数8
人力资源基本薪资：800元，KPI单位薪资：120元，人力资源小明，招聘10名员工，总工资为2000元
产品经理基本薪资：1000元，KPI单位薪资：100元，产品经理小红，上线4个产品，平均满意度为7，总工资为2100元
软件工程师基本薪资：1500元，KPI单位薪资：80元，软件工程师张三，完成10个需求，修复5个问题，总工资为2700元
软件工程师基本薪资：1500元，KPI单位薪资：80元，软件工程师李四，完成3个需求，修复6个问题，总工资为2220元
软件工程师基本薪资：1500元，KPI单位薪资：80元，软件工程师王五，完成7个需求，修复1个问题，总工资为2140元
--- PASS: TestVisitor (0.00s)
```


# ![图片](F:\Images\解释器模式.png)                			                    **解释器模式**

## **（一）概念**

解释器模式用于描述如何使用面向对象语言构成一个简单的语言解释器。在某些情况下，为了更好地描述某一些特定类型的问题，我们可以创建一种新的语言，这种语言拥有自己的表达式和结构，即文法规则，这些问题的实例将对应为该语言中的句子。此时，可以使用解释器模式来设计这种新的语言。对解释器模式的学习能够加深我们对面向对象思想的理解，并且掌握编程语言中文法规则的解释过程。



## **（二）示例**

定义一个解析特征值的语句解释器，提供是否包含特征值的终结表达式，并提供或表达式与且表达式，同时，生成南极洲特征判断表达式，及美国人特征判断表达式，最后测试程序根据对象特征值描述，通过表达式判断是否为真。



### **（三）特征值解释器**

```go

package interpreter

import "strings"

// Expression 表达式接口，包含一个解释方法
type Expression interface {
  Interpret(context string) bool
}

// terminalExpression 终结符表达式，判断表达式中是否包含匹配数据
type terminalExpression struct {
  matchData string
}

func NewTerminalExpression(matchData string) *terminalExpression {
  return &terminalExpression{matchData: matchData}
}

// Interpret 判断是否包含匹配字符
func (t *terminalExpression) Interpret(context string) bool {
  if strings.Contains(context, t.matchData) {
    return true
  }
  return false
}

// orExpression 或表达式
type orExpression struct {
  left, right Expression
}

func NewOrExpression(left, right Expression) *orExpression {
  return &orExpression{
    left:  left,
    right: right,
  }
}

func (o *orExpression) Interpret(context string) bool {
  return o.left.Interpret(context) || o.right.Interpret(context)
}

// andExpression 与表达式
type andExpression struct {
  left, right Expression
}

func NewAndExpression(left, right Expression) *andExpression {
  return &andExpression{
    left:  left,
    right: right,
  }
}

func (o *andExpression) Interpret(context string) bool {
  return o.left.Interpret(context) && o.right.Interpret(context)
}
```

### **（四）测试程序**

```go
package interpreter

import (
  "fmt"
  "testing"
)

func TestInterpreter(t *testing.T) {
  isAntarcticaExpression := generateCheckAntarcticaExpression()
  // 大洲描述1
  continentDescription1 := "此大洲生活着大量企鹅，全年低温，并且伴随着有暴风雪"
  fmt.Printf("%s，是否是南极洲？%t\n", continentDescription1, isAntarcticaExpression.Interpret(continentDescription1))
  // 大洲描述2
  continentDescription2 := "此大洲生活着狮子，全年高温多雨"
  fmt.Printf("%s，是否是南极洲？%t\n", continentDescription2, isAntarcticaExpression.Interpret(continentDescription2))

  isAmericanExpression := generateCheckAmericanExpression()
  peopleDescription1 := "此人生活在北美洲的黑人，说着英语，持有美国绿卡"
  fmt.Printf("%s，是否是美国人？%t\n", peopleDescription1, isAmericanExpression.Interpret(peopleDescription1))

  peopleDescription2 := "此人生活在欧洲，说着英语，是欧洲议会议员"
  fmt.Printf("%s，是否是南极洲？%t\n", peopleDescription2, isAmericanExpression.Interpret(peopleDescription2))

}

// generateCheckAntarcticaExpression 生成校验是否是南极洲表达式
func generateCheckAntarcticaExpression() Expression {
  // 判断南极洲的动物，或关系
  animalExpression := NewOrExpression(NewTerminalExpression("企鹅"),
    NewTerminalExpression("蓝鲸"))
  // 判断南极洲的天气，与关系
  weatherExpression := NewAndExpression(NewTerminalExpression("低温"),
    NewTerminalExpression("暴风雪"))
  // 最终返回动物与天气的与关系
  return NewAndExpression(animalExpression, weatherExpression)
}

// generateCheckAmericanExpression 生成检查美国人表达式
func generateCheckAmericanExpression() Expression {
  // 人种判断，或关系
  raceExpression := NewOrExpression(NewTerminalExpression("白人"),
    NewTerminalExpression("黑人"))
  // 生活方式，与关系
  lifeStyleExpression := NewAndExpression(NewTerminalExpression("英语"),
    NewTerminalExpression("北美洲"))
  // 身份，与关系
  identityExpression := NewAndExpression(lifeStyleExpression, NewTerminalExpression("美国绿卡"))
  return NewAndExpression(raceExpression, identityExpression)
}
```

### **（五）运行结果**

```go
=== RUN   TestInterpreter
此大洲生活着大量企鹅，全年低温，并且伴随着有暴风雪，是否是南极洲？true
此大洲生活着狮子，全年高温多雨，是否是南极洲？false
此人生活在北美洲的黑人，说着英语，持有美国绿卡，是否是美国人？true
此人生活在欧洲，说着英语，是欧洲议会议员，是否是美国人？false
--- PASS: TestInterpreter (0.00s)
PASS
```

# ![图片](F:\Images\适配器模式.png)                			                    **适配器模式**

### **（一）概念**

适配器模式是一种结构型设计模式，它能使接口不兼容的对象能够相互合作。

适配器可担任两个对象间的封装器，它会接收对于一个对象的调用， 并将其转换为另一个对象可识别的格式和接口。



### **（二）示例**

通过充电宝给不同充电接口的手机充电是一个非常符合适配器模式特征的生活示例；一般充电宝提供USB电源输出接口，手机充电输入接口则分为两类一是苹果手机的lightning接口，另一类是安卓手机的typeC接口，这两类接口都需要通过适配电源线连接充电宝的USB接口，这里USB接口就相当于充电宝的通用接口，lightning或typeC接口要想充电需要通过充电线适配。



### **（三）手机充电插头**

```go
package adapter

import "fmt"

// HuaweiPlug 华为手机充电插槽接口
type HuaweiPlug interface {
  ConnectTypeC() string
}

// HuaweiPhone 华为系列手机
type HuaweiPhone struct {
  model string
}

// NewHuaweiPhone 华为手机创建方法
func NewHuaweiPhone(model string) *HuaweiPhone {
  return &HuaweiPhone{
    model: model,
  }
}

// ConnectTypeC 华为手机TypeC充电插槽
func (h *HuaweiPhone) ConnectTypeC() string {
  return fmt.Sprintf("%v connect typeC plug", h.model)
}

// ApplePlug 苹果手机充电插槽
type ApplePlug interface {
  ConnectLightning() string
}

// IPhone 苹果系列手机
type IPhone struct {
  model string
}

// NewIPhone 苹果手机创建方法
func NewIPhone(model string) *IPhone {
  return &IPhone{
    model: model,
  }
}

// ConnectLightning 苹果手机Lightning充电插槽
func (i *IPhone) ConnectLightning() string {
  return fmt.Sprintf("%v connect lightning plug", i.model)
}
```

### **（四）充电宝适配器**

```go
package adapter

import "fmt"


// CommonPlug 通用的USB电源插槽
type CommonPlug interface {
  ConnectUSB() string
}

// HuaweiPhonePlugAdapter 华为TypeC充电插槽适配通用USB充电插槽
type HuaweiPhonePlugAdapter struct {
  huaweiPhone HuaweiPlug
}

// NewHuaweiPhonePlugAdapter 创建华为手机适配USB充电插槽适配器
func NewHuaweiPhonePlugAdapter(huaweiPhone HuaweiPlug) *HuaweiPhonePlugAdapter {
  return &HuaweiPhonePlugAdapter{
    huaweiPhone: huaweiPhone,
  }
}

// ConnectUSB 链接USB
func (h *HuaweiPhonePlugAdapter) ConnectUSB() string {
  return fmt.Sprintf("%v adapt to usb ", h.huaweiPhone.ConnectTypeC())
}

// ApplePhonePlugAdapter 苹果Lightning充电插槽适配通用USB充电插槽
type ApplePhonePlugAdapter struct {
  iPhone ApplePlug
}

// NewApplePhonePlugAdapter 创建苹果手机适配USB充电插槽适配器
func NewApplePhonePlugAdapter(iPhone ApplePlug) *ApplePhonePlugAdapter {
  return &ApplePhonePlugAdapter{
    iPhone: iPhone,
  }
}

// ConnectUSB 链接USB
func (a *ApplePhonePlugAdapter) ConnectUSB() string {
  return fmt.Sprintf("%v adapt to usb ", a.iPhone.ConnectLightning())
}

// PowerBank 充电宝
type PowerBank struct {
  brand string
}

// Charge 支持通用USB接口充电
func (p *PowerBank) Charge(plug CommonPlug) string {
  return fmt.Sprintf("%v power bank connect usb plug, start charge for %v", p.brand, plug.ConnectUSB())
}
```

### **（五）测试程序**

```go
package adapter

import (
  "fmt"
  "testing"
)

func TestAdapter (t *testing.T) {
  huaweiMate40Pro := NewHuaweiPhone("华为 mate40 pro")
  iphone13MaxPro := NewIPhone("苹果 iphone13 pro max")

  powerBank := &PowerBank{"飞利浦"}
  fmt.Println(powerBank.Charge(NewHuaweiPhonePlugAdapter(huaweiMate40Pro)))
  fmt.Println(powerBank.Charge(NewApplePhonePlugAdapter(iphone13MaxPro)))
}
```

### **（六）运行结果**

```go
=== RUN   TestAdapter
飞利浦 power bank connect usb plug, start charge for 华为 mate40 pro connect typeC plug adapt to usb 
飞利浦 power bank connect usb plug, start charge for 苹果 iphone13 pro max connect lightning plug adapt to usb 
--- PASS: TestAdapter (0.00s)
PASS
```



# ![图片](F:\Images\桥接模式.png)                			                    **桥接模式**

## **（一）概念**

桥接是一种结构型设计模式，可将业务逻辑或一个大类拆分为不同的层次结构， 从而能独立地进行开发。

层次结构中的第一层（通常称为抽象部分）将包含对第二层 （实现部分） 对象的引用。抽象部分将能将一些（有时是绝大部分）对自己的调用委派给实现部分的对象。所有的实现部分都有一个通用接口，因此它们能在抽象部分内部相互替换。

简单的说，一个事物存在多个维度的变化点，每一个维度都抽象出一个接口，事物引用这些接口实现整体行为逻辑，而每一个接口都可以存在多个变化的实现。

更简单的一句话：依赖接口编程。



## **（二）示例**

对于一段经历的描述，经历就可能有多种实现，比如旅游经历，探险经历这相当于第一层次的类结构，同时描述旅游经历或探险经历又包含多个维度，比如如何到达目的地，在目的地开展了什么活动等，到达目的地有很多种方式，比如飞机、火车、汽车等；开展的活动又根据地点不同而不同，海边可以冲浪，山地可以攀岩，荒漠可以徒步穿越等；这两个维度的变化点对于描述经历来说相当于第二层次类实现，通过接口被第一层次引用。

这里对于经历描述存在三个维度的变化，

1.经历本身的两个实现：旅游经历与探险经历。

2.交通方式的两个实现：飞机和汽车。

3.开展活动的三个实现：冲浪、攀岩与徒步穿越。

如果用一个类层次去实现就需要2*2*3=12个不同的实现类，如果用桥接模式仅需要2+2+3=7个不同的类，而且两种方式的加速度也不一样，比如增加一个交通方式火车，非桥接模式需要增加2*3*3-12=6个实现类，桥接模式2+3+3-7=1个实现类；桥接模式大大增加了类之间组合的灵活性。



### **（三）交通工具**

```go
package bridge

// Traffic 交通工具
type Traffic interface {
   Transport() string
}

// airplane 飞机
type airplane struct{}

// Transport 坐飞机
func (a *airplane) Transport() string {
   return "by airplane"
}

// car 汽车
type car struct{}

// Transport 坐汽车
func (t *car) Transport() string {
   return "by car"
}
```

### **（四）目的地**

```go
package bridge

import "fmt"

// Location 地点
type Location interface {
  Name() string // 地点名称
  PlaySports() string // 参与运动
}

// namedLocation 被命名的地点，统一引用此类型，声明名字字段及获取方法
type namedLocation struct {
  name string
}

// Name 获取地点名称
func (n namedLocation) Name() string {
  return n.name
}

// seaside 海边
type seaside struct {
  namedLocation
}

// NewSeaside 创建指定名字的海边，比如三亚湾
func NewSeaside(name string) *seaside {
  return &seaside{
    namedLocation: namedLocation{
      name: name,
    },
  }
}

// PlaySports 海边可以冲浪
func (s *seaside) PlaySports() string {
  return fmt.Sprintf("surfing")
}

// mountain 山
type mountain struct {
  namedLocation
}

// NewMountain 创建指定名字的山，比如泰山
func NewMountain(name string) *mountain {
  return &mountain{
    namedLocation: namedLocation{
      name: name,
    },
  }
}

// PlaySports 可以爬山
func (m *mountain) PlaySports() string {
  return fmt.Sprintf("climbing")
}

// desert 荒漠
type desert struct {
  namedLocation
}

// NewDesert 创建指定名字的荒漠，比如罗布泊
func NewDesert(name string) *desert {
  return &desert{
    namedLocation: namedLocation{
      name: name,
    },
  }
}

// PlaySports 荒漠可以徒步穿越
func (d *desert) PlaySports() string {
  return fmt.Sprintf("trekking")
}
```

### **（五）经历描述**

```go

package bridge

import "fmt"

// Experience 经历
type Experience interface {
  Describe() string // 描述经历
}

// travelExperience 旅游经历
type travelExperience struct {
  subject  string
  traffic  Traffic
  location Location
}

// NewTravelExperience 创建旅游经历，包括主题、交通方式、地点
func NewTravelExperience(subject string, traffic Traffic, location Location) *travelExperience {
  return &travelExperience{
    subject:  subject,
    traffic:  traffic,
    location: location,
  }
}

// Describe 描述旅游经历
func (t *travelExperience) Describe() string {
  return fmt.Sprintf("%s is to %s %s and %s", t.subject, t.location.Name(), t.traffic.Transport(), t.location.PlaySports())
}

// adventureExperience 探险经历
type adventureExperience struct {
  survivalTraining string
  travelExperience
}

// NewAdventureExperience 创建探险经历，包括探险需要的培训，其他的与路由参数类似
func NewAdventureExperience(training string, subject string, traffic Traffic, location Location) *adventureExperience {
  return &adventureExperience{
    survivalTraining: training,
    travelExperience: *NewTravelExperience(subject, traffic, location),
  }
}

// Describe 描述探险经历
func (a *adventureExperience) Describe() string {
  return fmt.Sprintf("after %s, %s", a.survivalTraining, a.travelExperience.Describe())
}
```

### **（六）测试程序**

```go
package bridge

import (
   "fmt"
   "testing"
)

func TestBridge(t *testing.T) {
   // 坐飞机去三亚度蜜月
   honeymoonTravel := NewTravelExperience("honeymoon", new(airplane), NewSeaside("SanyaYalongBay"))
   fmt.Println(honeymoonTravel.Describe())
   // 坐车去泰山毕业旅游
   graduationTrip := NewTravelExperience("graduationTrip", new(car), NewMountain("Tarzan"))
   fmt.Println(graduationTrip.Describe())

   // 野外生存培训后，坐车去罗布泊，徒步穿越
   desertAdventure := NewAdventureExperience("wilderness survival training", "adventure", new(car), NewDesert("Lop Nor"))
   fmt.Println(desertAdventure.Describe())
}
```

### **（七）运行结果**

```go
=== RUN   TestBridge
honeymoon is to SanyaYalongBay by airplane and surfing
graduationTrip is to Tarzan by car and climbing
after wilderness survival training, adventure is to Lop Nor by car and trekking
--- PASS: TestBridge (0.00s)
PASS
```



# ![图片](F:\Images\组合模式.png)                			                    **组合模式**

## **（一）概念**

组合是一种结构型设计模式，你可以使用它将对象组合成树状结构，并且能像使用独立对象一样使用它们。

对于绝大多数需要生成树状结构的问题来说，组合都是非常受欢迎的解决方案。组合最主要的功能是在整个树状结构上递归调用方法并对结果进行汇总。



## **（二）示例**

一般来说一个地区统计人口或经济总量，总是通过行政区划一层层上报汇总得出结果，区镇是最低一级行政区划，需要落实统计人口及经济总量的工作，再上一级行政区划需要将所辖区镇的数据汇总统计，以此类推每一级行政区划都需要统计人口与经济总量，就像一个倒过来的树状结构，各级行政区划统一的组件接口是统计人口与经济总量，区镇相当于最底层的叶子节点，中间级别行政区划相当于组合节点；下面代码以苏州市为例；



### **（三）组件接口**

```go
package composite

// Region 行政区，作为组合模式component接口
type Region interface {
   Name() string    // 名称
   Population() int //人口
   GDP() float64    // gdp
}
```

### **（四）区镇实现**

```go
package composite

// town 区镇，组合模式中相当于叶子节点
type town struct {
  name       string
  population int
  gdp        float64
}

// NewTown 创建区镇，根据名称、人口、GDP
func NewTown(name string, population int, gdp float64) *town {
  return &town{
    name:       name,
    population: population,
    gdp:        gdp,
  }
}

func (c *town) Name() string {
  return c.name
}

func (c *town) Population() int {
  return c.population
}

func (c *town) GDP() float64 {
  return c.gdp
}
```

### **（五）县市地市实现**

```go

package composite

// cities 市，包括县市或者地市，组合模式中相当于composite
type cities struct {
  name    string
  regions map[string]Region
}

// NewCities 创建一个市
func NewCities(name string) *cities {
  return &cities{
    name:    name,
    regions: make(map[string]Region),
  }
}

func (c *cities) Name() string {
  return c.name
}

func (c *cities) Population() int {
  sum := 0
  for _, r := range c.regions {
    sum += r.Population()
  }
  return sum
}

func (c *cities) GDP() float64 {
  sum := 0.0
  for _, r := range c.regions {
    sum += r.GDP()
  }
  return sum
}

// Add 添加多个行政区
func (c *cities) Add(regions ...Region) {
  for _, r := range regions {
    c.regions[r.Name()] = r
  }
}

// Remove 递归删除行政区
func (c *cities) Remove(name string) {
  for n, r := range c.regions {
    if n == name {
      delete(c.regions, name)
      return
    }
    if city, ok := r.(*cities); ok {
      city.Remove(name)
    }
  }
}

func (c *cities) Regions() map[string]Region {
  return c.regions
}
```

### **（六）测试程序**

```go

package composite

import (
   "fmt"
   "testing"
)

func TestComposite(t *testing.T) {
   gusu := NewTown("姑苏区", 100, 2000.00)
   fmt.Println(ShowRegionInfo(gusu))
   wuzhong := NewTown("吴中区", 150, 2600.00)
   fmt.Println(ShowRegionInfo(wuzhong))
   huqiu := NewTown("虎丘区", 80, 1800.00)
   fmt.Println(ShowRegionInfo(huqiu))

   kunshan := NewCities("昆山市")
   kunshan.Add(NewTown("玉山镇", 60, 1200.00),
      NewTown("周庄镇", 68, 1900.00),
      NewTown("花桥镇", 78, 2200.00))
   fmt.Println(ShowRegionInfo(kunshan))

   changshu := NewCities("常熟市")
   changshu.Add(NewTown("沙家浜镇", 55, 1100.00),
      NewTown("古里镇", 59, 1300.00),
      NewTown("辛庄镇", 68, 2100.00))
   fmt.Println(ShowRegionInfo(changshu))

   suzhou := NewCities("苏州市")
   suzhou.Add(gusu, wuzhong, huqiu, kunshan, changshu)
   fmt.Println(ShowRegionInfo(suzhou))

}

func ShowRegionInfo(region Region) string {
   return fmt.Sprintf("%s, 人口:%d万, GDP:%.2f亿", region.Name(), region.Population(), region.GDP())
}
```

### **（七）运行结果**

```go
=== RUN   TestComposite
姑苏区, 人口:100万, GDP:2000.00亿
吴中区, 人口:150万, GDP:2600.00亿
虎丘区, 人口:80万, GDP:1800.00亿
昆山市, 人口:206万, GDP:5300.00亿
常熟市, 人口:182万, GDP:4500.00亿
苏州市, 人口:718万, GDP:16200.00亿
--- PASS: TestComposite (0.00s)
PASS
```



# ![图片](F:\Images\装饰模式.png)                			                    **装饰模式**

## **（一）概念**

装饰是一种结构设计模式，允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。

由于目标对象和装饰器遵循同一接口，因此你可用装饰来对对象进行无限次的封装。结果对象将获得所有封装器叠加而来的行为。



## **（二）示例**

地铁进站的过程一般情况下只需要买票，检票进站，但是如果你是带行李，就需要进行安全检查，如果是疫情时期，就需要进行疫情防护检查，比如戴口罩、测量体温等，这里买票进站相当于通用进站流程，安检及防疫检查就相当于加强的修饰行为。



### **（三）修饰器实现**

```go
package decorator

import "fmt"

// Station 车站，修饰器模式统一接口
type Station interface {
  Enter() string // 进站
}

// subwayStation 地铁站
type subwayStation struct {
  name string
}

// NewSubwayStation 创建指定站名地铁站
func NewSubwayStation(name string) *subwayStation {
  return &subwayStation{
    name: name,
  }
}

// Enter 进地铁站
func (s *subwayStation) Enter() string {
  return fmt.Sprintf("买票进入%s地铁站。", s.name)
}

// securityCheckDecorator 进站安检修饰器
type securityCheckDecorator struct {
  station Station
}

func NewSecurityCheckDecorator(station Station) *securityCheckDecorator {
  return &securityCheckDecorator{
    station: station,
  }
}

func (s *securityCheckDecorator) Enter() string {
  return "行李通过安检；" + s.station.Enter()
}

// epidemicProtectionDecorator 进站疫情防护修饰器
type epidemicProtectionDecorator struct {
  station Station
}

func NewEpidemicProtectionDecorator(station Station) *epidemicProtectionDecorator {
  return &epidemicProtectionDecorator{
    station: station,
  }
}

func (e *epidemicProtectionDecorator) Enter() string {
  return "测量体温，佩戴口罩；" + e.station.Enter()
}
```

### **（四）测试代码**

```go
package decorator

import (
   "fmt"
   "testing"
)

func TestDecorator(t *testing.T) {
   xierqiStation := NewSubwayStation("西二旗")
   fmt.Println(EnhanceEnterStationProcess(xierqiStation, false, false).Enter())
   fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, false).Enter())
   fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, true).Enter())
}

// EnhanceEnterStationProcess 根据是否有行李，是否处于疫情，增加进站流程
func EnhanceEnterStationProcess(station Station, hasLuggage bool, hasEpidemic bool) Station {
   if hasLuggage {
      station = NewSecurityCheckDecorator(station)
   }
   if hasEpidemic {
      station = NewEpidemicProtectionDecorator(station)
   }
   return station
}
```

### **（五）运行结果**

```go
=== RUN   TestDecorator
买票进入西二旗地铁站。
行李通过安检；买票进入西二旗地铁站。
测量体温，佩戴口罩；行李通过安检；买票进入西二旗地铁站。
--- PASS: TestDecorator (0.00s)
PASS
```



# ![图片](F:\Images\外观模式.png)                			                    **外观模式**

## **（一）概念**

外观是一种结构型设计模式，能为复杂系统、程序库或框架提供一个简单 （但有限） 的接口。

尽管外观模式降低了程序的整体复杂度，但它同时也有助于将不需要的依赖移动到同一个位置。



## **（二）示例**

用户在淘宝电商系统买商品时，只需要选好商品在结算页点击提交即可完成下单；在客户端系统仅需要一个创建订单的方法，但是整个订单的生成需要很多步骤，比如查询用户配送地址，查询商品价格，使用优惠券，扣减商品库存，支付相应价钱等。



### **（三）淘宝电商系统**

```go

package facade

import "fmt"

// TaobaoFacade 淘宝网站门面，在淘宝网站下单涉及到多个系统配合调用，包括用户系统，商品系统，优惠券系统，库存系统，支付系统，最终生成订单
type TaobaoFacade struct {
  userService    *UserService
  productService *ProductService
  couponService  *CouponService
  stockService   *StockService
  paymentService *PaymentService
}

// NewTaobaoFacade 创建淘宝网站
func NewTaobaoFacade() *TaobaoFacade {
  return &TaobaoFacade{
    userService: &UserService{},
    productService: &ProductService{
      products: map[string]float64{"笔记本电脑": 6666.66},
    },
    couponService:  &CouponService{},
    stockService:   &StockService{},
    paymentService: &PaymentService{},
  }
}

// CreateOrder 根据用户名，商品名，商品数量生成购买订单
func (t *TaobaoFacade) CreateOrder(userName string, productName string, count int) string {
  // 使用优惠券
  couponInfo := t.couponService.useCoupon()
  // 扣减库存
  stockInfo := t.stockService.decreaseFor(productName, count)
  // 计算商品总价
  sumPrice := t.productService.getProductPrice(productName) * float64(count)
  // 支付价格
  payInfo := t.paymentService.pay(sumPrice)
  return fmt.Sprintf("用户%s,购买了%d件%s商品,%s,%s,%s,送货到%s", userName, count, productName, couponInfo,
    stockInfo, payInfo, t.userService.getUserAddress(userName))
}

// UserService 用户系统
type UserService struct{}

func (u *UserService) getUserAddress(userName string) string {
  return fmt.Sprintf("%s地址是:北京市海淀区中关村大街，1号院2号楼3单元402", userName)
}

// ProductService 商品系统
type ProductService struct {
  products map[string]float64
}

func (p *ProductService) getProductPrice(productName string) float64 {
  return p.products[productName]
}

// CouponService 优惠券系统
type CouponService struct{}

func (c *CouponService) useCoupon() string {
  return "使用满100减20优惠券"
}

// StockService 库存系统
type StockService struct{}

func (s *StockService) decreaseFor(productName string, count int) string {
  return fmt.Sprintf("扣减%d件%s商品库存", count, productName)
}

// PaymentService 支付系统
type PaymentService struct{}

func (p *PaymentService) pay(amount float64) string {
  return fmt.Sprintf("支付金额%.2f", amount)
}
```

### **（四）测试程序**

```go
package facade

import (
   "fmt"
   "testing"
)

func TestFacade(t *testing.T) {
   // 通过门面模式，隐藏下单过程中，后端多个系统的复杂交互
   taobao := NewTaobaoFacade()
   fmt.Println(taobao.CreateOrder("张三", "笔记本电脑", 1))
}
```

### **（五）运行结果**

```go
=== RUN   TestFacade
用户张三,购买了1件笔记本电脑商品,使用满100减20优惠券,扣减1件笔记本电脑商品库存,支付金额6666.66,送货到张三地址是:北京市海淀区中关村大街，1号院2号楼3单元402
--- PASS: TestFacade (0.00s)
PASS
```



# ![图片](F:\Images\享元模式.png)                			                    **享元模式**

## **（一）概念**

享元是一种结构型设计模式，它允许你在消耗少量内存的情况下支持大量对象。

模式通过共享多个对象的部分状态来实现上述功能。换句话来说，享元会将不同对象的相同数据进行缓存以节省内存。



## **（二）示例**

北京出租车调度系统，需要每隔一分钟记录一下全市出租车的位置信息，假设为了提高系统响应速度，近一天的数据需要存储在内存中，每个位置信息包括出租车辆信息及位置信息，位置信息在系统中就是一个(x,y)坐标，车辆信息包括车的号牌，颜色，品牌和所属公司，在调度系统存储的出租车行驶轨迹中，位置是实时在变化的，但车辆信息就可以通过享元模式共用一个对象引用，来减少内存消耗。



### **（三）出租车享元对象**

```go
package flyweight

import (
  "fmt"
)

// Taxi 出租车，享元对象，保存不变的内在属性信息
type Taxi struct {
  licensePlate string // 车牌
  color        string // 颜色
  brand        string // 汽车品牌
  company      string // 所属公司
}

// LocateFor 获取定位信息
func (t *Taxi) LocateFor(monitorMap string, x, y int) string {
  return fmt.Sprintf("%s,对于车牌号%s,%s,%s品牌,所属%s公司,定位(%d,%d)", monitorMap,
    t.licensePlate, t.color, t.brand, t.company, x, y)
}

// taxiFactoryInstance 出租车工厂单例
var taxiFactoryInstance = &TaxiFactory{
  taxis: make(map[string]*Taxi),
}

// GetTaxiFactory 获取出租车工厂单例
func GetTaxiFactory() *TaxiFactory {
  return taxiFactoryInstance
}

// TaxiFactory 出租车工厂类
type TaxiFactory struct {
  taxis map[string]*Taxi // key为车牌号
}

// getTaxi 获取出租车
func (f *TaxiFactory) getTaxi(licensePlate, color, brand, company string) *Taxi {
  if _, ok := f.taxis[licensePlate]; !ok {
    f.taxis[licensePlate] = &Taxi{
      licensePlate: licensePlate,
      color:        color,
      brand:        brand,
      company:      company,
    }
  }
  return f.taxis[licensePlate]
}
```

### **（四）出租车调度系统**

```go
package flyweight

import "bytes"

// TaxiPosition 出租车位置信息 x,y为外在数据信息，taxi为内在数据信息（享元对象）
type TaxiPosition struct {
  x    int
  y    int
  taxi *Taxi
}

func NewTaxiPosition(taxi *Taxi, x, y int) *TaxiPosition {
  return &TaxiPosition{
    taxi: taxi,
    x:    x,
    y:    y,
  }
}

// LocateFor 定位信息
func (p *TaxiPosition) LocateFor(monitorMap string) string {
  return p.taxi.LocateFor(monitorMap, p.x, p.y)
}

// TaxiDispatcher 出租车调度系统
type TaxiDispatcher struct {
  name   string
  traces map[string][]*TaxiPosition // 存储出租车当天轨迹信息，key为车牌号
}

func NewTaxiDispatcher(name string) *TaxiDispatcher {
  return &TaxiDispatcher{
    name:   name,
    traces: make(map[string][]*TaxiPosition),
  }
}

// AddTrace 添加轨迹
func (t *TaxiDispatcher) AddTrace(licensePlate, color, brand, company string, x, y int) {
  taxi := GetTaxiFactory().getTaxi(licensePlate, color, brand, company)
  t.traces[licensePlate] = append(t.traces[licensePlate], NewTaxiPosition(taxi, x, y))
}

// ShowTraces 显示轨迹
func (t *TaxiDispatcher) ShowTraces(licensePlate string) string {
  bytesBuf := bytes.Buffer{}
  for _, trace := range t.traces[licensePlate] {
    bytesBuf.WriteString(trace.LocateFor(t.name))
    bytesBuf.WriteByte('\n')
  }
  return bytesBuf.String()
}
```

### **（五）测试程序**

```go
package flyweight

import (
   "fmt"
   "testing"
)

func TestFlyweight(t *testing.T) {
   dispatcher := NewTaxiDispatcher("北京市出租车调度系统")
   dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 10, 20)
   dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 20, 30)
   dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 30, 40)
   dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 40, 50)

   dispatcher.AddTrace("京B.567890", "红色", "一汽大众", "首汽", 20, 40)
   dispatcher.AddTrace("京B.567890", "红色", "一汽大众", "首汽", 50, 50)

   fmt.Println(dispatcher.ShowTraces("京B.123456"))
   fmt.Println(dispatcher.ShowTraces("京B.567890"))
}
```

### **（六）运行结果**

```go
=== RUN   TestFlyweight
北京市出租车调度系统,对于车牌号京B.123456,黄色,北京现代品牌,所属北汽公司,定位(10,20)
北京市出租车调度系统,对于车牌号京B.123456,黄色,北京现代品牌,所属北汽公司,定位(20,30)
北京市出租车调度系统,对于车牌号京B.123456,黄色,北京现代品牌,所属北汽公司,定位(30,40)
北京市出租车调度系统,对于车牌号京B.123456,黄色,北京现代品牌,所属北汽公司,定位(40,50)

北京市出租车调度系统,对于车牌号京B.567890,红色,一汽大众品牌,所属首汽公司,定位(20,40)
北京市出租车调度系统,对于车牌号京B.567890,红色,一汽大众品牌,所属首汽公司,定位(50,50)
--- PASS: TestFlyweight (0.00s)
PASS
```



# ![图片](F:\Images\代理模式.png)                			                    **代理模式**

## **（一）概念**

代理是一种结构型设计模式，让你能提供真实服务对象的替代品给客户端使用。代理接收客户端的请求并进行一些处理 （访问控制和缓存等）， 然后再将请求传递给服务对象。

代理对象拥有和服务对象相同的接口，这使得当其被传递给客户端时可与真实对象互换。

修饰与代理是非常相似的设计模式，都是基于组合设计原则，也就是说一个对象应该将部分工作委派给另一个对象。但两者之间不同点我认为是，修饰器模式总是要执行服务对象，对于执行之前或执行之后结果进行加强，服务对象基本是客户端创建好再嵌套外层的修饰对象；而代理模式不一定执行服务对象，有可能通过缓存，延迟加载等没有访问服务对象，同时服务对象什么时候创建也是由代理类决定的。



## **（二）示例**

房屋中介代理帮助房东卖房子，这个过程就是一个代理模式的过程，中介会收集尽量多的卖房信息，并通过各种渠道发布，同时中介会随时带客户看房，并初步商讨价格，如果达成初步购买意向，才会约房东讨论房屋价格，最后签约卖房；房屋中介与房东都实现卖房接口，中介会提前坐一些前期工作，如果都没问题，才会约房东执行真正的签约卖房流程。



### **（三）房屋中介卖房**

```go
package proxy

import (
  "bytes"
  "fmt"
)

// HouseSeller 房屋出售者
type HouseSeller interface {
  SellHouse(address string, buyer string) string
}

// houseProxy 房产中介代理
type houseProxy struct {
  houseSeller HouseSeller
}

func NewHouseProxy(houseSeller HouseSeller) *houseProxy {
  return &houseProxy{
    houseSeller: houseSeller,
  }
}

// SellHouse 中介卖房，看房->初步谈价->最终和房东签协议
func (h *houseProxy) SellHouse(address string, buyer string) string {
  buf := bytes.Buffer{}
  buf.WriteString(h.viewHouse(address, buyer) + "\n")
  buf.WriteString(h.preBargain(address, buyer) + "\n")
  buf.WriteString(h.houseSeller.SellHouse(address, buyer))
  return buf.String()
}

// viewHouse 看房介绍基本情况
func (h *houseProxy) viewHouse(address string, buyer string) string {
  return fmt.Sprintf("带买家%s看位于%s的房屋，并介绍基本情况", buyer, address)
}

// preBargain 初步沟通价格
func (h *houseProxy) preBargain(address string, buyer string) string {
  return fmt.Sprintf("讨价还价后，初步达成购买意向")
}

// houseOwner 房东
type houseOwner struct{}

// SellHouse 房东卖房，商讨价格，签署购房协议
func (h *houseOwner) SellHouse(address string, buyer string) string {
  return fmt.Sprintf("最终商讨价格后，与%s签署购买地址为%s的购房协议。", buyer, address)
}
```

### **（四）测试程序**

```go
package proxy

import (
  "fmt"
  "testing"
)

func TestProxy(t *testing.T) {
  proxy := NewHouseProxy(&houseOwner{})
  fmt.Println(proxy.SellHouse("北京市海淀区中关村大街，2号院1号楼4单元502室", "李四"))
}
```

### **（五）运行结果**

```go
=== RUN   TestProxy
带买家李四看位于北京市海淀区中关村大街，2号院1号楼4单元502室的房屋，并介绍基本情况
讨价还价后，初步达成购买意向
最终商讨价格后，与李四签署购买地址为北京市海淀区中关村大街，2号院1号楼4单元502室的购房协议。
--- PASS: TestProxy (0.00s)
PASS
```



# ![图片](F:\Images\工厂方法模式.png)                			                    **工厂方法模式**

## **（一）概念**

工厂方法模式是一种创建型设计模式，其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。



## **（二）示例**

摊煎饼的小贩需要先摊个煎饼，再卖出去，摊煎饼就可以类比为一个工厂方法，根据顾客的喜好摊出不同口味的煎饼。



### **（三）接口**

```go
package factorymethod

// Pancake 煎饼
type Pancake interface {
  // ShowFlour 煎饼使用的面粉
  ShowFlour() string
  // Value 煎饼价格
  Value() float32
}

// PancakeCook 煎饼厨师
type PancakeCook interface {
  // MakePancake 摊煎饼
  MakePancake() Pancake
}

// PancakeVendor 煎饼小贩
type PancakeVendor struct {
  PancakeCook
}

// NewPancakeVendor ...
func NewPancakeVendor(cook PancakeCook) *PancakeVendor {
  return &PancakeVendor{
    PancakeCook: cook,
  }
}

// SellPancake 卖煎饼，先摊煎饼，再卖
func (vendor *PancakeVendor) SellPancake() (money float32) {
  return vendor.MakePancake().Value()
}
```

### **（四）实现**

各种面的煎饼实现

```go
package factorymethod

// cornPancake 玉米面煎饼
type cornPancake struct{}

// NewCornPancake ...
func NewCornPancake() *cornPancake {
  return &cornPancake{}
}

func (cake *cornPancake) ShowFlour() string {
  return "玉米面"
}

func (cake *cornPancake) Value() float32 {
  return 5.0
}

// milletPancake 小米面煎饼
type milletPancake struct{}

func NewMilletPancake() *milletPancake {
  return &milletPancake{}
}

func (cake *milletPancake) ShowFlour() string {
  return "小米面"
}

func (cake *milletPancake) Value() float32 {
  return 8.0
}
```

制作各种口味煎饼的工厂方法实现

```go
package factorymethod

// cornPancakeCook 制作玉米面煎饼厨师
type cornPancakeCook struct{}

func NewCornPancakeCook() *cornPancakeCook {
  return &cornPancakeCook{}
}

func (cook *cornPancakeCook) MakePancake() Pancake {
  return NewCornPancake()
}

// milletPancakeCook 制作小米面煎饼厨师
type milletPancakeCook struct{}

func NewMilletPancakeCook() *milletPancakeCook {
  return &milletPancakeCook{}
}

func (cook *milletPancakeCook) MakePancake() Pancake {
  return NewMilletPancake()
}
```

### **（五）运用**

```go
package factorymethod

import (
  "fmt"
  "testing"
)

func TestFactoryMethod(t *testing.T) {
  pancakeVendor := NewPancakeVendor(NewCornPancakeCook())
  fmt.Printf("Corn pancake value is %v\n", pancakeVendor.SellPancake())

  pancakeVendor = NewPancakeVendor(NewMilletPancakeCook())
  fmt.Printf("Millet pancake value is %v\n", pancakeVendor.SellPancake())
}
```

### **（六）输出**

```go
=== RUN   TestFactoryMethod
Corn pancake value is 5
Millet pancake value is 8
--- PASS: TestFactoryMethod (0.00s)
PASS
```



# ![图片](F:\Images\抽象工厂模式.png)                			                    **抽象工厂模式**

## **（一）概念**

抽象工厂是一种创建型设计模式，它能创建一系列相关的对象，而无需指定其具体类。

抽象工厂定义了用于创建不同产品的接口，但将实际的创建工作留给了具体工厂类。每个工厂类型都对应一个特定的产品变体。

在创建产品时，客户端代码调用的是工厂对象的构建方法，而不是直接调用构造函数 （ new操作符）。由于一个工厂对应一种产品变体，因此它创建的所有产品都可相互兼容。

客户端代码仅通过其抽象接口与工厂和产品进行交互。该接口允许同一客户端代码与不同产品进行交互。你只需创建一个具体工厂类并将其传递给客户端代码即可。



## **（二）示例**

厨师准备一餐时，会分别做吃的和喝的，根据早、中、晚三餐饮食习惯，会分别制作不同的饮食，厨师就相当于抽象工厂，制作三餐的不同烹饪方式就好比不同抽象工厂的实现。



### **（三）接口**

```go
package abstractfactory
// Cook 厨师接口，抽象工厂
type Cook interface {
  // MakeFood 制作主食
  MakeFood() Food
  // MakeDrink 制作饮品
  MakeDrink() Drink
}

// Food 主食接口
type Food interface {
  // Eaten 被吃
  Eaten() string
}

// Drink 饮品接口
type Drink interface {
  // Drunk 被喝
  Drunk() string
}
```

### **（四）实现**

三餐不同厨师接口的实现

```go
package abstractfactory

// breakfastCook 早餐厨师
type breakfastCook struct{}

func NewBreakfastCook() *breakfastCook {
  return &breakfastCook{}
}

func (b *breakfastCook) MakeFood() Food {
  return &cakeFood{"切片面包"}
}

func (b *breakfastCook) MakeDrink() Drink {
  return &gruelDrink{"小米粥"}
}

// lunchCook 午餐厨师
type lunchCook struct{}

func NewLunchCook() *lunchCook {
  return &lunchCook{}
}

func (l *lunchCook) MakeFood() Food {
  return &dishFood{"烤全羊"}
}

func (l *lunchCook) MakeDrink() Drink {
  return &sodaDrink{"冰镇可口可乐"}
}

// dinnerCook 晚餐厨师
type dinnerCook struct{}

func NewDinnerCook() *dinnerCook {
  return &dinnerCook{}
}

func (d *dinnerCook) MakeFood() Food {
  return &noodleFood{"大盘鸡拌面"}
}

func (d *dinnerCook) MakeDrink() Drink {
  return &soupDrink{"西红柿鸡蛋汤"}
}
```

不同吃的

```go

package abstractfactory

import "fmt"

// cakeFood 蛋糕
type cakeFood struct {
  cakeName string
}

func (c *cakeFood) Eaten() string {
  return fmt.Sprintf("%v被吃", c.cakeName)
}

// dishFood 菜肴
type dishFood struct {
  dishName string
}

func (d *dishFood) Eaten() string {
  return fmt.Sprintf("%v被吃", d.dishName)
}

// noodleFood 面条
type noodleFood struct {
  noodleName string
}

func (n *noodleFood) Eaten() string {
  return fmt.Sprintf("%v被吃", n.noodleName)
}
```

不同喝的

```go
package abstractfactory

import "fmt"

// gruelDrink 粥
type gruelDrink struct {
  gruelName string
}

func (g *gruelDrink) Drunk() string {
  return fmt.Sprintf("%v被喝", g.gruelName)
}

// sodaDrink 汽水
type sodaDrink struct {
  sodaName string
}

func (s *sodaDrink) Drunk() string {
  return fmt.Sprintf("%v被喝", s.sodaName)
}

// soupDrink 汤
type soupDrink struct {
  soupName string
}

func (s *soupDrink) Drunk() string {
  return fmt.Sprintf("%v被喝", s.soupName)
}
```

### **（五）运用**

```go
package abstractfactory

import (
  "fmt"
  "testing"
)

func TestAbstractFactory(t *testing.T) {
  fmt.Printf("breakfast: %v\n", HaveMeal(NewBreakfastCook()))
  fmt.Printf("lunch: %v\n", HaveMeal(NewLunchCook()))
  fmt.Printf("dinner: %v\n", HaveMeal(NewDinnerCook()))
}

// HaveMeal 吃饭
func HaveMeal(cook Cook) string {
  return fmt.Sprintf("%s %s", cook.MakeFood().Eaten(), cook.MakeDrink().Drunk())
}
```

### **（六）输出**

```go
=== RUN   TestAbstractFactory
breakfast: 切片面包被吃 小米粥被喝
lunch: 烤全羊被吃 冰镇可口可乐被喝
dinner: 大盘鸡拌面被吃 西红柿鸡蛋汤被喝
--- PASS: TestAbstractFactory (0.00s)
PASS
```



# ![图片](F:\Images\生成器模式.png)                			                    **生成器模式**

## **（一）概念**

生成器是一种创建型设计模式，使你能够分步骤创建复杂对象。

与其他创建型模式不同，生成器不要求产品拥有通用接口。这使得用相同的创建过程生成不同的产品成为可能。



## **（二）示例**

还是摊煎饼的例子，摊煎饼分为四个步骤，1放面糊、2放鸡蛋、3放调料、4放薄脆，通过四个创建过程，制作好一个煎饼，这个摊煎饼的过程就好比煎饼生成器接口，不同生成器的实现就相当于摊不同品类的煎饼，比如正常的煎饼，健康的煎饼（可能用的是粗粮面、柴鸡蛋、非油炸薄脆、不放酱等），生成器接口方法也可以通过参数控制煎饼的大小，比如放两勺面糊，放2个鸡蛋等。

生成器的使用者为了避免每次都调用相同的构建步骤，也可以通过包装类固定几种构建过程，生成几类常用的产品，就好像摊煎饼有几类常卖固定成品，比如普通的，加两个鸡蛋的，不要香菜的等等，这几类固定构建过程提前定制好，直接通过简单工厂方法就直接创建，如果用户再需要细粒度的定制构建，再通过生成器创建。



### **（三）接口**

```go
package builder

// Quantity 分量
type Quantity int

const (
  Small  Quantity = 1
  Middle Quantity = 5
  Large  Quantity = 10
)

type PancakeBuilder interface {
  // PutPaste 放面糊
  PutPaste(quantity Quantity)
  // PutEgg 放鸡蛋
  PutEgg(num int)
  // PutWafer 放薄脆
  PutWafer()
  // PutFlavour 放调料 Coriander香菜，Shallot葱 Sauce酱
  PutFlavour(hasCoriander, hasShallot, hasSauce bool)
  // Build 摊煎饼
  Build() *Pancake
}

// Pancake  煎饼
type Pancake struct {
  pasteQuantity Quantity // 面糊分量
  eggNum        int      // 鸡蛋数量
  wafer         string   // 薄脆
  hasCoriander  bool     // 是否放香菜
  hasShallot    bool     // 是否放葱
  hasSauce      bool     // 是否放酱
}

```

### **（四）实现**

正常煎饼创建器

```go
package builder

type normalPancakeBuilder struct {
  pasteQuantity Quantity // 面糊量
  eggNum        int      // 鸡蛋数量
  friedWafer    string   // 油炸薄脆
  hasCoriander  bool     // 是否放香菜
  hasShallot    bool     // 是否放葱
  hasHotSauce   bool     // 是否放辣味酱
}

func NewNormalPancakeBuilder() *normalPancakeBuilder {
  return &normalPancakeBuilder{}
}

func (n *normalPancakeBuilder) PutPaste(quantity Quantity) {
  n.pasteQuantity = quantity
}

func (n *normalPancakeBuilder) PutEgg(num int) {
  n.eggNum = num
}

func (n *normalPancakeBuilder) PutWafer() {
  n.friedWafer = "油炸的薄脆"
}

func (n *normalPancakeBuilder) PutFlavour(hasCoriander, hasShallot, hasSauce bool) {
  n.hasCoriander = hasCoriander
  n.hasShallot = hasShallot
  n.hasHotSauce = hasSauce
}

func (n *normalPancakeBuilder) Build() *Pancake {
  return &Pancake{
    pasteQuantity: n.pasteQuantity,
    eggNum:        n.eggNum,
    wafer:         n.friedWafer,
    hasCoriander:  n.hasCoriander,
    hasShallot:    n.hasShallot,
    hasSauce:      n.hasHotSauce,
  }
}
```

健康煎饼创建器

```go

package builder

type healthyPancakeBuilder struct {
  milletPasteQuantity Quantity // 小米面糊量
  chaiEggNum          int      // 柴鸡蛋数量
  nonFriedWafer       string   // 非油炸薄脆
  hasCoriander        bool     // 是否放香菜
  hasShallot          bool     // 是否放葱
}

func NewHealthyPancakeBuilder() *healthyPancakeBuilder {
  return &healthyPancakeBuilder{}
}

func (n *healthyPancakeBuilder) PutPaste(quantity Quantity) {
  n.milletPasteQuantity = quantity
}

func (n *healthyPancakeBuilder) PutEgg(num int) {
  n.chaiEggNum = num
}

func (n *healthyPancakeBuilder) PutWafer() {
  n.nonFriedWafer = "非油炸的薄脆"
}

func (n *healthyPancakeBuilder) PutFlavour(hasCoriander, hasShallot, _ bool) {
  n.hasCoriander = hasCoriander
  n.hasShallot = hasShallot
}

func (n *healthyPancakeBuilder) Build() *Pancake {
  return &Pancake{
    pasteQuantity: n.milletPasteQuantity,
    eggNum:        n.chaiEggNum,
    wafer:         n.nonFriedWafer,
    hasCoriander:  n.hasCoriander,
    hasShallot:    n.hasShallot,
    hasSauce:      false,
  }
}
```

煎饼生成器的封装类-厨师

```go

package builder

// PancakeCook 摊煎饼师傅
type PancakeCook struct {
  builder PancakeBuilder
}

func NewPancakeCook(builder PancakeBuilder) *PancakeCook {
  return &PancakeCook{
    builder: builder,
  }
}

// SetPancakeBuilder 重新设置煎饼构造器
func (p *PancakeCook) SetPancakeBuilder(builder PancakeBuilder) {
  p.builder = builder
}

// MakePancake 摊一个一般煎饼
func (p *PancakeCook) MakePancake() *Pancake {
  p.builder.PutPaste(Middle)
  p.builder.PutEgg(1)
  p.builder.PutWafer()
  p.builder.PutFlavour(true, true, true)
  return p.builder.Build()
}

// MakeBigPancake 摊一个巨无霸煎饼
func (p *PancakeCook) MakeBigPancake() *Pancake {
  p.builder.PutPaste(Large)
  p.builder.PutEgg(3)
  p.builder.PutWafer()
  p.builder.PutFlavour(true, true, true)
  return p.builder.Build()
}

// MakePancakeForFlavour 摊一个自选调料霸煎饼
func (p *PancakeCook) MakePancakeForFlavour(hasCoriander, hasShallot, hasSauce bool) *Pancake {
  p.builder.PutPaste(Large)
  p.builder.PutEgg(3)
  p.builder.PutWafer()
  p.builder.PutFlavour(hasCoriander, hasShallot, hasSauce)
  return p.builder.Build()
}
```

### **（五）运用**

```go
package builder

import (
  "fmt"
  "testing"
)

func TestBuilder(t *testing.T) {
  pancakeCook := NewPancakeCook(NewNormalPancakeBuilder())
  fmt.Printf("摊一个普通煎饼 %#v\n", pancakeCook.MakePancake())

  pancakeCook.SetPancakeBuilder(NewHealthyPancakeBuilder())
  fmt.Printf("摊一个健康的加量煎饼 %#v\n", pancakeCook.MakeBigPancake())
}
```

### **（六）输出**

```go
=== RUN   TestBuilder
摊一个普通煎饼 &builder.Pancake{pasteQuantity:5, eggNum:1, wafer:"油炸的薄脆", hasCoriander:true, hasShallot:true, hasSauce:true}
摊一个健康的加量煎饼 &builder.Pancake{pasteQuantity:10, eggNum:3, wafer:"非油炸的薄脆", hasCoriander:true, hasShallot:true, hasSauce:false}
--- PASS: TestBuilder (0.00s)
PASS
```



# ![图片](F:\Images\原型模式.png)                			                    **原型模式**

## **（一）概念**

原型是一种创建型设计模式，使你能够复制对象，甚至是复杂对象，而又无需使代码依赖它们所属的类。

所有的原型类都必须有一个通用的接口， 使得即使在对象所属的具体类未知的情况下也能复制对象。原型对象可以生成自身的完整副本， 因为相同类的对象可以相互访问对方的私有成员变量。



## **（二）示例**

纸质文件可以通过复印机轻松拷贝出多份，设置Paper接口，包含读取文件内容和克隆文件两个方法。同时声明两个类报纸（Newspaper）和简历（Resume）实现了Paper接口，通过复印机（Copier）复印出两类文件的副本，并读取文件副本内容。



### **（三）接口实现**

```go

package prototype

import (
  "bytes"
  "fmt"
  "io"
)

// Paper 纸张，包含读取内容的方法，拷贝纸张的方法，作为原型模式接口
type Paper interface {
  io.Reader
  Clone() Paper
}

// Newspaper 报纸 实现原型接口
type Newspaper struct {
  headline string
  content  string
}

func NewNewspaper(headline string, content string) *Newspaper {
  return &Newspaper{
    headline: headline,
    content:  content,
  }
}

func (np *Newspaper) Read(p []byte) (n int, err error) {
  buf := bytes.NewBufferString(fmt.Sprintf("headline:%s,content:%s", np.headline, np.content))
  return buf.Read(p)
}

func (np *Newspaper) Clone() Paper {
  return &Newspaper{
    headline: np.headline + "_copied",
    content:  np.content,
  }
}

// Resume 简历 实现原型接口
type Resume struct {
  name       string
  age        int
  experience string
}

func NewResume(name string, age int, experience string) *Resume {
  return &Resume{
    name:       name,
    age:        age,
    experience: experience,
  }
}

func (r *Resume) Read(p []byte) (n int, err error) {
  buf := bytes.NewBufferString(fmt.Sprintf("name:%s,age:%d,experience:%s", r.name, r.age, r.experience))
  return buf.Read(p)
}

func (r *Resume) Clone() Paper {
  return &Resume{
    name:       r.name + "_copied",
    age:        r.age,
    experience: r.experience,
  }
}
```

### **（四）运用**

```go
package prototype

import (
  "fmt"
  "reflect"
  "testing"
)

func TestPrototype(t *testing.T) {
  copier := NewCopier("云打印机")
  oneNewspaper := NewNewspaper("Go是最好的编程语言", "Go语言十大优势")
  oneResume := NewResume("小明", 29, "5年码农")

  otherNewspaper := copier.copy(oneNewspaper)
  copyNewspaperMsg := make([]byte, 100)
  byteSize, _ := otherNewspaper.Read(copyNewspaperMsg)
  fmt.Println("copyNewspaperMsg:" + string(copyNewspaperMsg[:byteSize]))

  otherResume := copier.copy(oneResume)
  copyResumeMsg := make([]byte, 100)
  byteSize, _ = otherResume.Read(copyResumeMsg)
  fmt.Println("copyResumeMsg:" + string(copyResumeMsg[:byteSize]))
}

// Copier 复印机
type Copier struct {
  name string
}

func NewCopier(n string) *Copier {
  return &Copier{name: n}
}

func (c *Copier) copy(paper Paper) Paper {
  fmt.Printf("copier name:%v is copying:%v ", c.name, reflect.TypeOf(paper).String())
  return paper.Clone()
}
```

### **（五）输出**

```go
=== RUN   TestPrototype
copier name:云打印机 is copying:*prototype.Newspaper copyNewspaperMsg:headline:Go是最好的编程语言_copied,content:Go语言十大优势
copier name:云打印机 is copying:*prototype.Resume copyResumeMsg:name:小明_copied,age:29,experience:5年码农
--- PASS: TestPrototype (0.00s)
PASS
```



# ![图片](F:\Images\单例模式.png)                			                    **单例模式**

### **（一）概念**

单例是一种创建型设计模式，让你能够保证一个类只有一个实例，并提供一个访问该实例的全局节点。

单例拥有与全局变量相同的优缺点。尽管它们非常有用，但却会破坏代码的模块化特性。



### **（二）示例**

通过地球对象实现单例，earth不能导出，通过TheEarth方法访问全局唯一实例，并通过sync.Once实现多协程下一次加载。



### **（三）接口实现**

```go
package singleton

import "sync"

var once sync.Once

// 不可导出对象
type earth struct {
  desc string
}

func (e *earth) String() string {
  return e.desc
}

// theEarth 地球单实例
var theEarth *earth

// TheEarth 获取地球单实例
func TheEarth() *earth {
  if theEarth == nil {
    once.Do(func() {
      theEarth = &earth{
        desc: "美丽的地球，孕育了生命。",
      }
    })
  }
  return theEarth
}
```

### **（四）运用**

```go
package singleton

import (
  "fmt"
  "testing"
)

func TestSingleton(t *testing.T) {
  fmt.Println(TheEarth().String())
}
```

### **（五）输出**

```go
=== RUN   TestSingleton
美丽的地球，孕育了生命。
--- PASS: TestSingleton (0.00s)
PASS
```

