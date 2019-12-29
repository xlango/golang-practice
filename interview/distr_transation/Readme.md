#分布式事务原理实践
<h2>一、Seata 原理实践<h2>

<h3>简述</h3>
1.开启全局事务    
2.本地事务注册至事务管理器  
3.代理Datasource执行本地事务，commit/rollback前wait   
4.事务管理器根据所有本地事务执行情况计算出所有分支事务commit/rollback     
5.事务管理器通知分支事务commit/rollback   

<h3>推荐教程</h3>

    https://www.bilibili.com/video/av50531999

<h2>一、基于mq 原理实践<h2> 
1.rocketmq


<h3>推荐教程</h3>

    https://www.bilibili.com/video/av64434096?t=16