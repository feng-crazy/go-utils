# 项目名称
简要说明
```
[背景]:
    在调研过程中，经常需要对一些网站进行定向抓取。由于go包含各种强大的库，使用go做定向抓取比较简单。请使用go开发一个迷你定向抓取器mini_spider，实现对种子链接的抓取，并把URL长相符合特定pattern的网页保存到磁盘上。

[程序运行]：
   ./mini_spider -c ../conf -l ../log

[配置文件spider.conf]：
    [spider]
    # 种子文件路径
    urlListFile = ../data/url.data
    # 抓取结果存储目录 
    outputDirectory = ../output
    # 最大抓取深度(种子为0级)
    maxDepth = 1
    # 抓取间隔. 单位: 秒 
    crawlInterval =  1
    # 抓取超时. 单位: 秒 
    crawlTimeout = 1
    # 需要存储的目标网页URL pattern(正则表达式)
    targetUrl = .*.(htm|html)$
    # 抓取routine数 
    threadCount = 8
   
[种子文件为json格式，示例如下]：
   [
     "http://www.baidu.com",
     "http://www.sina.com.cn",
     ...
   ]  
```
## 设计思路
 设计一个队列和协程池来做任务,每个协程从队列中取任务执行，在执行任务中发现子页面就创建任务添加，所有任务执行完成后退出
### 程序初始化

* 读取命令行参数、初始化日志、读取配置文件、url文件
* 创建并初始化协程和队列，添加任务，调用其Start方法开始调度
* 调用Wait方法等待任务结束

### 调度主要逻辑

* 用for循环来创建配置的routine数量
* 每一个routine中循环从任务队列中取出任务，然后执行任务
* 任务中发现新的任务，就添加到任务队列中
* 用一个任务Channel来判断任务是否结束
* 如果队列中的任务为空，且没有协程在执行任务中，就退出协程

### 任务执行主要逻辑

* 判断当前任务的深度，大于等于MaxDepth则返回
* 判断当前任务的URL是否已经爬取过，若是则直接返回，否则开启go routine异步执行任务
* 获取当前任务的域名（站点），检查是否满足爬取间隔的要求
* 根据URL爬取网页，失败则记录日志并返回
* 判断其Content-Type是不是文本，不是则记录日志并返回
* 将爬取到的网页转换成UTF-8格式
* 判断该URL是否满足目标正则表达式，若满足则将其保存至磁盘
* 解析爬取到的网页，将其子URL加入任务队列

### 控制抓取间隔

* 通过sync.Map和time.Timer实现
* sync.Map的Key为hostname，Value为timer
* 每次执行抓取任务前通过任务的URL解析出hostname，通过hostname拿到该站点的timer，等待timer的剩余时间后，重置timer执行抓取任务

### 并发控制

* 通过sync.WaitGroup来做并发控制
* 创建一个channel来判断任务是否执行完

### 退出

* 使用sync.WaitGroup的wait方法等待所有线程结束退出


## 快速开始
如何构建、安装、运行
* sh build_run.sh
## 测试
如何执行自动化测试
* go test 
## 如何贡献
贡献patch流程、质量要求


