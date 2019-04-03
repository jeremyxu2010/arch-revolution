# 应用架构演进

本项目是一个展示应用架构演进的示例工程，包含以下几个子工程：

## 1. 单体应用架构示例

`demo1`是一个传统的单体应用，该示例的详细说明见[demo1的README](demo1/README.md)。

## 2. 微服务架构示例

`demo2`是一个进行了微服务改造后的应用，该示例的详细说明见[demo2的README](demo2/README.md)。

## 3. Istio服务网格纳管传统微服务应用

`demo3`在`demo2`基础上进行了简单的调整，以适应将其部署到虚拟机，并由Istio服务网格纲管的场景。

该示例的详细说明见[demo3的README](demo3/README.md)。

## 4. 微服务应用架构优化

`demo4`在`demo3`的基础之上，为优化应用性能，加入了缓存服务、消息队列服务。

该示例的详细说明见[demo4的README](demo4/README.md)。

## 5. 微服务应用云原生化

`demo5`在`demo4`的基础上，将应用各模块打包成容器镜像，并以标准的[Chart](https://helm.sh/docs/developing_charts/)包形式描述整个应用部署信息，便于部署到kubernetes里去，以实现真正的应用云原生化，真正体验到云原生应用的优势。

该示例的详细说明见[demo5的README](demo5/README.md)。