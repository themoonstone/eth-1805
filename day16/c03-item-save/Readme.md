## 1. 提取初始页面信息
## 2. 使用正则提取城市链接和名称
## 3. 单任务版爬虫抽象
    1. 封装fetcher
    2. 封装解析器
    3. 实现启动函数engine
    4. 编写cityList测试函数
    5. 实现城市解析器
    6. 实现用户解析器
## 4. 并发(concurrent)
    1. 封装(解析器+Fetcher)worker

## 5. Item数据存储
    1. docker + ElasticSearch
    2. 将数据存入elasticSearch
## docker简述
    1. 为什么要用ES(ElasticSearch)而不是mysql
        不想建表，ES不用建表
    2. 为什么要用docker而不是直接在机器上安装ES
        直接安装需要安装JAVA...，还有可能要配置环境，麻烦！！！
        通过docker只需要一条命令:docker pull elasticsearch:version
        就可以完成安装配置，然后直接使用
        docker:可以把我们的程序、其它的软件等等放在上面运行的一个容器

