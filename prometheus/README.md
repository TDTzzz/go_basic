#promethes+node_exporter+grafana监控体系钻研



> 参考资料：

> > [promethes-book](https://yunlzheng.gitbook.io/prometheus-book/parti-prometheus-ji-chu/quickstart/why-monitor)

> > [写的很不错的博客](https://www.hwholiday.com/2019/docker_install_prometheus/)



### prometheus搭建

> docker run -d -p 9090:9090 --name prometheus -v /Users/tt/docker/prometheus:/data -v /Users/tt/docker/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml  prom/prometheus



普普通通的用docker搭建，端口9090，挂载 data文件夹 和 配置文件。



##### Prometheus 组成及架构

Prometheus 生态圈中包含了多个组件，其中许多组件是可选的：

- **Prometheus Server**: 用于收集和存储时间序列数据。
- **Client Library**: 客户端库，为需要监控的服务生成相应的 metrics 并暴露给 Prometheus server。当 Prometheus server 来 pull 时，直接返回实时状态的 metrics。
- **Push Gateway**: 主要用于短期的 jobs。由于这类 jobs 存在时间较短，可能在 Prometheus 来 pull 之前就消失了。为此，这次 jobs 可以直接向 Prometheus server 端推送它们的 metrics。这种方式主要用于服务层面的 metrics，对于机器层面的 metrices，需要使用 node exporter。
- **Exporters**: 用于暴露已有的第三方服务的 metrics 给 Prometheus。
- **Alertmanager**: 从 Prometheus server 端接收到 alerts 后，会进行去除重复数据，分组，并路由到对收的接受方式，发出报警。常见的接收方式有：电子邮件，pagerduty，OpsGenie, webhook 等。
- 一些其他的工具。



### 配置node_exporter

这里官方是建议在本机跑，不建议在docker里搭。毕竟这个是监控主机硬件指标运行情况的。

直接从 [prometheus.io](https://prometheus.io/download/) 里下载对应版本的对应二进制包，然后解压执行，启动9100端口即可。

[项目地址](https://github.com/prometheus/node_exporter)



**注意的坑**

docker里的prometheus是无法通过127.0.0.1 connect到其他容器或者宿主机的localhost的。所以如果配置node_exporter的target为127.0.0.1:9100，会出现浏览器可以访问，但是prometheus的targets里会down。



![image-20200926154130279](/Users/tt/Library/Application Support/typora-user-images/image-20200926154130279.png)

**解决方法**

/etc/hosts文件里加 ```127.0.0.1       host.docker.internal```

配置文件里也用

```yaml
  - job_name: 'node2'
    static_configs:
      - targets: ['host.docker.internal:9100']

```



**配置文件实例**

```yaml
global:
  scrape_interval: 15s
  external_labels:
    monitor: 'codelab-monitor'
scrape_configs:
  - job_name: 'go-pro'
    metrics_path: '/metrics'
    scrape_interval: 5s
    static_configs:
      - targets: ['127.0.0.1:9090']
  - job_name: 'node2'
    static_configs:
      - targets: ['host.docker.internal:9100']

```





### Grafana搭建并连prometheus

1.启动grafana的docker

```
docker run -d -p 3000:3000 grafana/grafana
```

2.访问localhost:3000,进去后登录。账号/密码 默认都是admin

3.找到data source，配置到prometheus的url。如 127.0.0.1:9090

4.创建dashboard，选择数据源为刚才创的prometheus

5.在metrics那里填写和promQL，语句和promtheus自导的graph一样。

![image-20200926154755256](/Users/tt/Library/Application Support/typora-user-images/image-20200926154755256.png)







### todo

1. 研究promQL及prometheus的数据模型

2. 研究其他的prometheus配套组件，如Alertmanager告警处理

3. 结合微服务。用golang的Client library，为需要监控的服务生成相应的metrics并暴露给promethes。

4. .....

5. .......

