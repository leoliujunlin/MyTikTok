## 文件目录说明

- main.go 为入口文件
- router.go 为路由，提供了URL到控制器的映射关系
- middlewares/ 为中间件目录
- static/: 静态资源目录，存放CSS、JS、图片等静态文件。
- common 为自己定义的资源包，比如包括数据库初始化文件database.go
--------------
另外主要需要编写的地方在：
- ①controllers/ 为控制器目录（MVC）
- ②services/: 服务目录，封装业务逻辑，供控制器调用
- ③model 为模型目录，提供了数据库中表单映射的结构体类型，也提供了一些基本的增删改查的功能(也称之为Dao层)
- --->这三层是自顶向下的关系,处理逻辑主要在此

> controllers内需要定义一系列func (c *gin.Context)方法
> <br>其中：c *gin.Context里保存了请求的上下文信息，它是所有请求处理器的入口参数
> <br>目的：在每个control这里面需要完成 **读取请求参数-->处理请求(完成数据库操作)--->返回响应**

------------

> 上面的调用关系是：
> <br> main.go--->router--->middlewares---->controllers---->(Services(sessions))---->models--->database
> <br> 当前端发来请求，后端通过router来找到对应的路由，并进行必要的中间件方法调用middleware。简单中间件auth鉴权后，进入API，当API接口来验证(BindAndValid) 参数response格式是否正确。在API处适当的调用service处的方法，为API提供Helper或是数据库操作，其中不免用到model/databse数据库对象。并返回对应处理，顺便通过中间件记录日志到logs
## 一般流程：
1. 连接数据库初始化文件,/common/database.go,主要用作连接数据等参数
2. 定义数据模型操作文件,如操作用户相关/model/user.go 包含模型，以及常用数据 新建/编辑/删除/查询 等相关定义
3. 定义数据查询接口,如用户相关/controller/userController.go
4. 定义对外访问路由,在/router/router.go中定义

## 使用说明
1. 首先使用下载mysql,创建数据库,并修改database.go文件
2. 使用go get下载gorm,gin等资源包
3. 使用go run main.go即可运行