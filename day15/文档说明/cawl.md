# 爬虫简述

## 背景

- 应用场景：大数据时代，AI时代，需要大量的数据进行分析，这些数据的来源，通常都是爬虫
- 分类
  - 通用爬虫：比如百度、google
  - 聚焦爬虫：比如专门爬取博客，知乎问答等，主要是从互联网获取结构化数据
- 作用：把网页转换为数据
## 设计
- ![项目架构](D:\Project\go\src\eth-1805\day14\01-分布式网络爬虫\img\1 爬虫项目总体架构.jpg)
- 项目设计
  - 网络数据爬取--从html页面中爬取数据
  - 数据清洗(筛选)
  - 数据存储：mysql+ElasticSearch
  - 使用go语言标准模板库实现HTTP数据展示部分
- 爬虫的主题
  - 相亲网站、求职网站
  - 本项目爬去相亲网站(珍爱网)
- 总体算法
  - 如何发现用户
    - 通过城市列表-->城市-->(下一页)-->用户
    - 通过用户-->猜你喜欢
    - 通过已有用户ID+1来猜测用户ID
  - ![](D:\Project\go\src\eth-1805\day14\01-分布式网络爬虫\img\2 爬虫总体算法.jpg)
  - ![](D:\Project\go\src\eth-1805\day14\01-分布式网络爬虫\img\3 爬虫总体算法2.jpg)
- 项目设计
  - 数据存储：mysql+ElasticSearch
  - 使用go语言标准模板库实现HTTP数据展示部分
- 爬虫的主题
  - 相亲网站、求职网站
  - 本项目爬去相亲网站
- 总体算法
  - 如何发现用户
    - 通过城市列表-->城市-->(下一页)-->用户
    - 通过用户-->猜你喜欢
    - 通过已有用户ID+1来猜测用户ID
- ![实现步骤](D:\Project\go\src\eth-1805\day14\01-分布式网络爬虫\img\4 爬虫实现步骤.jpg)

##实现单任务版爬虫

- 目标：获取并打印所有城市第一页用户的详细信息
- 目标地址：https://www.zhenai.com/zhenghun
- ![单任务版爬虫总体算法](D:\Project\go\src\eth-1805\day14\01-分布式网络爬虫\img\5 单任务版爬虫总体算法.jpg)
- ![](D:\Project\go\src\eth-1805\day14\01-分布式网络爬虫\img\6 单任务版爬虫架构.jpg)
- 单任务版爬虫实现思路
  - 首先，有一个engine(引擎)，将整个爬取行为进行驱动
  - seed代表种子，起始的URL
  - 将种子和解析器进行封装，作为 request传递给engine
  - engine将请求加载到任务队列中，方便维护
  - engine从任务队列中获取request中的URL传递给fetcher
  - fetcher从网络中获取数据，然后返回文本text给engine
  - engine获取到文本之后，传递给parser
  - parser返回requests,items。
  - requests被添加任务队列中
  - items打印输出

##实现并发版爬虫

- 背景：单任务爬虫太慢，性能很低
- ![](D:\Project\go\src\1805\day14\文档说明\img\7 并发版爬虫架构.jpg)
- 考虑一下在并发时，我们应该对哪些模块进行并发？对耗时最多，等待时间最长的模块进行并发，在这里面，应该是fetcher，但实际上我们的正则表达式处理模块parser也消耗了比较长的时间，而且fetcher的输出也就是parser的输入，所以我们可以考虑将二者放到一个大的模块中，称为worker，如下图![](D:\Project\go\src\1805\day14\文档说明\img\8 并发版爬虫(2).jpg)
- 多个goroutine共同抢夺一个channel

## 并发版本实现二：实现对worker和Request的控制

![](D:\Users\Ocam\eth-1805\day15\01-分布式网络爬虫\img\10 并发爬虫实现一架构更新.jpg)

![](D:\Users\Ocam\eth-1805\day15\01-分布式网络爬虫\img\11 并发爬虫scheduler实现三 request队列和worker队列.jpg)

1. 改进思路
   1. 将每一个request,存入到一个专门的request队列中，由Scheduler进行维护
   2. 把worker也存入一个队列，同样有scheduler进行维护
   3. 100个worker对应100个channel，进行数据交互时，将100个channel传入100个workerChan
2. ![](D:\Users\Ocam\eth-1805\day15\01-分布式网络爬虫\img\12 并发版爬虫架构(2).jpg)

