# Swagger-UI

```
http://127.0.0.1:8888/swagger
```

#### 环境变量

- SWAGGER_PORT

  ```text
  web ui 端口。默认：8888
  ```

- SWAGGER_DOC

  ```
  文档 json 文件夹。默认：./api
  ```

#### 默认加载

```text
修改 swagger-ui/swagger-initializer.js 内 url
```

#### 运行

- ##### Docker

  ```shell
  git clone https://github.com/charlesbases/swagger-ui.git
  
  cd swagger-ui && make
  
  # 需要挂载容器内 '/swagger/api' 文件夹
  ```