# opentracing


> [好文章](https://xie.infoq.cn/article/2a19ef287069f4d51cf0d8070)
> [好教程](https://github.com/yurishkuro/opentracing-tutorial/tree/master/go)

### 概念
* Opentracing 是分布式链路追踪的一种规范标准，它提供一套分布式追踪协议，是一个与平台无关，编程语言无关，接口统一的分布式追踪系统

* OpenTracing 是一个轻量级的标准化层，它位于应用程序/类库和追踪或日志分析程序之间。


### 数据模型

#### Span

**Span**是链路中的集锦本元素，一个span表示一个独立的工作单元，在入侵式链路追踪中可以表示一个接口的调用，一个数据库操作的调用等。
一个span中包含如下内容。
> * 服务名称
> * 服务的开始和结束时间
> * Span Tag
> * Span Log
> * SpanContext,Span的上下文对象
> * References(Span间关系)

##### Tags

Tags 是一个 K/V类型的键值对，用户可以自定义该标签，但是 value只能是 string类型。主要用于链路追踪结果对查询和过滤。需要注意的是， tags不会传递给下一个 span调用，即仅自己可见。
```
span.SetTag("http_code":"400")
span.SetTag("http_method":"GET")
```


#### Logs
Logs 也是一个 K/V类型的键值对，主要是用来存储时间发生时间，value的类型没有做限制。
```
span.LogFields(
  			log.String("database","mysql"),
  			log.Int("used_time":5),
  			log.Int("start_ts":1596335100),
)
```


