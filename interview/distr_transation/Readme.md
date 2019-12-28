#Seata 原理实践


<h3>简述</h3>
1.开启全局事务    
2.本地事务注册至事务管理器  
3.代理Datasource执行本地事务，commit/rollback前wait   
4.事务管理器根据所有本地事务执行情况计算出所有分支事务commit/rollback     
5.事务管理器通知分支事务commit/rollback    