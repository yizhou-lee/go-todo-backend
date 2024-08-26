# go-todo-backend

`todo-backend` 是一个使用 gin + gorm + mysql 创建的 todo list 项目的后端 RESTful API 服务器。

该项目主要目的是探索出一个清晰易懂的 Go 项目结构。

## 项目结构

```
todo-backend
│  .gitignore
│  go.mod
│  go.sum
│  README.md
├─cmd
│  └─app: 程序的入口文件
├─configs: 包含所有的配置文件
├─docs: 包含openapi规范的文档
│  ├─docs.go: swaggo/swag 自动生成的文档
│  ├─swagger.json: json格式的openapi规范
│  └─swagger.yaml: yaml格式的openapi规范
├─internal
│  ├─controllers: 包含所有的控制器，控制器处理具体的请求和响应
│  ├─middlewares: 包含所有的中间件，中间件用于处理请求和响应，例如身份验证、日志记录等
│  ├─models: 定义所有的数据模型，数据模型用于操作数据库
│  └─routes: 定义所有的路由信息，每个路由指向一个控制器
├─pkg: 外部应用程序可以使用的库代码
└─utils: 存放通用的、可重用的代码
```

## 安装

1. 克隆此仓库到本地

2. 在 `configs` 目录下创建一个 `config.yaml` 文件，并添加配置信息。(确保本地已安装 MySQL 并创建了数据库)
   ```
   # MySQL Configuration
   mysql:
     host: 127.0.0.1
     port: 3306
     username: <your_name>
     password: <your_password>
     database: todo_db
   ```

## 使用

运行以下命令启动服务器：

```bash
go run main.go
```

## References

1. [gin-framework](https://masteringbackend.com/posts/gin-framework#getting-started-with-gin)

2. [mux+mongo](https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gorillamux-version-57fh)

3. [optimizing gin in go](https://www.squash.io/optimizing-gin-in-golang-project-structuring-error-handling-and-testing/)
