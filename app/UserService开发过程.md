## 1. 设计用户前端
- 设计signin.tmpl
- 注册路由 sign-in
- #### 设计login接口(接收表单)
  - 1.编写idl--` auth_page.proto ` 
  - 2.引入第三方中间件`session` --注册中间件
  - 3.编写相应业务逻辑 `service/login.go`
- #### 设计register接口(接收表单)
  - 1.更新idl--` auth_page.proto ` --添加相应服务
  - 2.编写相应业务逻辑
  