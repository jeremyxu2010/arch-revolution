# 部署指引

1. 新建数据库

在mysql数据库中新建`user-service`、`blog-service`依赖的数据库`demo2_user_db`、`demo2_blog_db`，参考命令如下：

```sql
create database demo2_user_db default character set utf8;
create database demo2_blog_db default character set utf8;
```

初始化数据库表结构，参考命令如下：

```sql
use demo2_user_db;
source user-service/demo2_user_db.sql;
use demo2_blog_db;
source blog-service/demo2_blog_db.sql;
```

2. 修改应用配置文件

相应地修改应用的配置文件`user-service/src/main/resources/application.properties`、`blog-service/src/main/resources/application.properties`，特别注意数据库连接相应字段的设置。

另外为了不绑定IP地址，在应用里均使用`consul-service`、`mysql-service`指代consul、mysql的IP地址，请修改部署机上的/etc/hosts，将这两个名称指向正确的IP地址，如下：

```
127.0.0.1 consul-service
127.0.0.1 mysql-service
```

3. maven打包应用

安装JDK8及maven后，在本机使用mvn命令对应用进行打包，参考命令如下：

```bash
mvn package
```

上述命令会将生成5个微服务模块的jar包

4. 运行服务注册中心consul

本示例使用consul作为服务的注册中心，在运行应用前须先启动consul，参考命令如下：

```bash
curl -O https://releases.hashicorp.com/consul/1.4.4/consul_1.4.4_linux_amd64.zip
unzip consul_1.4.4_linux_amd64.zip
./consul agent -dev -server -ui -client=0.0.0.0
```

5. 启动应用

依次启动各个微服务模块即可，参考命令如下：

```bash
java -jar user-service/target/user-service-0.0.1-SNAPSHOT.jar > user-service.log 2>&1 &
java -jar blog-service/target/blog-service-0.0.1-SNAPSHOT.jar > blog-service.log 2>&1 &
java -jar aggregation-service/target/aggregation-service-0.0.1-SNAPSHOT.jar > aggregation-service.log 2>&1 &
java -jar apigateway-service/target/apigateway-service-0.0.1-SNAPSHOT.jar > apigateway-service.log 2>&1 &
java -jar frontend-service/target/frontend-service-0.0.1-SNAPSHOT.jar > frontend-service.log 2>&1 &
```

各微服务模块均启动后，用浏览器访问`http://${frontend-service-ip}:8085`即可。
