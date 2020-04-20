### 这是什么?  
>   1.这是一个基于go语言gin框架的web项目骨架，其目的主要在于将web项目主线业务逻辑，底层封装完善，开发者只需要关注真正的业务即可。  
>   2.本项目骨架封装了完整的用户注册、登录（获取token）、中间件token认证、用户的增、删、改、查等一个项目基本而又繁琐的基础功能，使用本项目骨架，您只需要关注您真正的业务即可，无需关注业务之外的东西。   

####    开箱即用
>   1.使用 `goland(>=2019.3版本)` 打开本项目，找到`database/db_demo.sql`导入数据库，打开`main.go`,点击`run`运行本项目，自动下载依赖， 片刻后即可启动.  
>   2.建议go语言使用>=1.14版本,才能更好的支持 `go module` 包管理.  
>![avatar](http://139.196.101.31:2080/GinSkeleton.jpg)  

####    压力测试  
>   2核4g云服务器，并发（Qps）可以达到1w+，所有请求100%成功！  
![avatar](http://139.196.101.31:2080/concurrent.png)  

####    本项目测试用例路由  
>GET    /                         
>POST   /Admin/users/register     
>POST   /Admin/users/login        
>GET    /Admin/users/index        
>POST   /Admin/users/create       
>POST   /Admin/users/edit         
>POST   /Admin/users/delete       


#### 版本 
V 1.0.05   2020-04-20 
>   1.增加`json`统一返回逻辑。  
>   2.用户模块核心逻辑全部完成（注册、登录、`token`授权、`token`认证、`CURD`等操作）。   
>   3.全局常量增加`CURD`常用的列表。  
>   4.增加`Service`层逻辑，并提供相关的示例代码。     
>   5.继续精简代码，使本项目骨架逻辑主线更加清晰，快速上手。       
>   6.更新本项目所必须的数据库`db_demo.sql`文件。       
>   7.精简代码，基本的业务操作只保留`tb_users`表的完整操作示例代码。         
>   8.增加文件上传公共模块，供任何有需要上传文件的业务模块调用。            
>   9.日志存储路径调整为全局变量统一定义。        

V 1.0.04   2020-04-19 
>   1.路由——>中间件——>表单验证器——>控制器 上下文数据一致性开发完成。    
>   2.验证器结构调整，将业务部分和系统核心部分分离，开发者只需更多关注业务即可。  
>   3.增加项目骨架所需的demo数据库。      

V 1.0.03   2020-04-17   
>   1.增加`linux`环境`chan signal`监听信号值，使程序在退出时，更加优雅，资源的释放更加完善。    

V 1.0.02   2020-04-16 
>   1.容器、事件注册器调整命名规范，增加模糊处理函数。        

V 1.0.01   2020-04-15 
>   1.增加容器，将一些比较繁琐的功能模块率先注册在容器，方便后续调用。  
>   2.表单参数验证器首先注册在容器，避免在路由模块不停地引入表单验证器造成该文件过于庞大。   
>   3.函数类事件精简代码，删除原有的一个键对应多个事件的逻辑，目前设置为键值一一对应关系。   
>   4.Mysql、Redis数据库连接的释放统一注册在函数事件管理器，由程序退出时统一释放。   
>   5.容器存储变量修改为sync.map，避免了并发情况下发生bug。     

V 1.0.00   2020-04-14 
>   1.基于gin框架的web项目骨架.  
>   2.开发单体应用基本的功能模块全部已经封装完毕.  
