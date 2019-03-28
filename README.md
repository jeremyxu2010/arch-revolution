# 应用架构演进

本项目是一个展示应用架构演进的示例工程，包含以下几个子工程：

## 1. 单体应用架构示例

`demo1`是一个传统的单体应用，其基于Spring Boot构建，整体采用成熟的MVC分层架构，传统绝大多数Java应用都是这种架构。采用JSP作为视图层，采用SpringMVC架构实现控制层，业务数据保存在外部mysql数据库中，所有各层模块等均被揉合在一个项目中。

部署指引请参考[demo1的README](demo1/README.md)。

## 2. 微服务架构示例

`demo2`是一个进行了微服务改造后的应用，其基于Spring Cloud构建。其被拆分为`user-service`、`blog-service`、`aggregation-service`、`apigateway-service`、`frontend-service`5个微服务子项目，各微服务的功能概要如下：

1. `user-service`、`blog-service`分别实现对各自业务模型的restful接口； 
2. `aggregation-service`调用`user-service`、`blog-service`，实现了跨业务模型的restful接口；
3. `apigateway-service`通过配置，将`user-service`、`blog-service`、`aggregation-service`的restful接口以统一的形式对暴露API接口; 
4. `frontend-service`采用[Vue.js](https://cn.vuejs.org/index.html)重写了前端，其调用`apigateway-service`暴露的API接口，实现业务逻辑

部署指引请参考[demo2的README](demo2/README.md)。

## 3. Istio服务网格纳管传统微服务应用

`demo3`在`demo2`基础上进行了简单的调整，以适应将其部署到虚拟机，并由Istio服务网格纲管的场景。

具体的调整事项及部署指引请参考[demo3的README](demo3/README.md)。

## 4. 微服务应用架构优化

`demo4`在`demo3`的基础之上，为优化应用性能，加入了缓存服务、消息队列服务。

详细的优化原因及优化方案请参考[demo4的README](demo4/README.md)。

## 5. 微服务应用容器化

`demo5`在`demo4`的基础上，将应用各模块打包成容器镜像，并以标准的[Chart](https://helm.sh/docs/developing_charts/)包形式描述整个应用部署信息，便于部署到kubernetes里去，以实现真正的应用云原生化。
