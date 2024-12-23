## **滑动时间窗口算法**

- 算法思想

滑动时间窗口算法，是从对普通时间窗口计数的优化。使用普通时间窗口时，我们会为每个user_id/ip维护一个KV: uidOrIp: timestamp_requestCount。假设限制1秒1000个请求，那么第100ms有一个请求，这个KV变成 uidOrIp: timestamp_1，递200ms有1个请求，我们先比较距离记录的timestamp有没有超过1s，如果没有只更新count，此时KV变成 uidOrIp: timestamp_2。当第1100ms来一个请求时，更新记录中的timestamp并重置计数，KV变成 uidOrIp: newtimestamp_1 普通时间窗口有一个问题，假设有500个请求集中在前1s的后100ms，500个请求集中在后1s的前100ms，其实在这200ms没就已经请求超限了，但是由于时间窗每经过1s就会重置计数，就无法识别到此时的请求超限。

对于滑动时间窗口，我们可以把1ms的时间窗口划分成10个time slot, 每个time slot统计某个100ms的请求数量。每经过100ms，有一个新的time slot加入窗口，早于当前时间100ms的time slot出窗口。窗口内最多维护10个time slot，储存空间的消耗同样是比较低的。



- 适用场景

与令牌桶一样，有应对突发流量的能力



- go语言实现

主要就是实现sliding window算法。可以参考Bilibili开源的kratos框架里circuit breaker用循环列表保存time slot对象的实现，他们这个实现的好处是不用频繁的创建和销毁time slot对象。下面给出一个简单的基本实现：