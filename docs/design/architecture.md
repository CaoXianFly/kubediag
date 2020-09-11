# Kubernetes 故障诊断恢复平台架构设计

Kubernetes 是一个生产级的容器编排引擎，但是 Kubernetes 仍然存在系统复杂、故障诊断成本高等问题。Kubernetes 故障诊断恢复平台是基于 Kubernetes 云原生基础设施能力打造的框架，旨在解决云原生体系中故障诊断、运维恢复的自动化问题。主要包括以下几个维度：

* 由 Kubernetes 以及 Docker 的 Bug 引起的故障。
* 内核 Bug 导致的故障。
* 基础设施抖动产生的问题。
* 用户在容器化以及使用 Kubernetes 过程中遇到的问题。
* 用户在容器化后遇到的业务相关问题。

## 目标

Kubernetes 故障诊断恢复平台的设计目标包括：

* 通用性：平台依赖通用技术实现，平台组件可以在绝大部分的 Linux 系统下运行并且能够对 Linux 下运行遇到的故障进行诊断和运维。
* 可扩展性：平台组件之间的交互为松耦合接口设计并且整个框架是可插拔式的。
* 可维护性：框架逻辑简洁明了，维护成本与功能数量为线性关系，不同故障的分析和恢复逻辑具有独立性。

## 架构

故障诊断恢复平台 Agent 组件可以监听 APIServer 获取 Abnormal 自定义资源，Abnormal 自定义资源是对故障状态机的抽象。故障诊断恢复平台 Agent 组件可以通过 Event 以及自定义组件产生相对应的 Abnormal 自定义资源并送入故障诊断恢复的流水线。故障诊断恢复平台 Agent 组件使用 DaemonSet 部署在集群中：

```
                                                            ----------------------
                                               Watch        |                    |
                                      --------------------->|     API Server     |
                                      |                     |                    |
                                      |                     |--------------------|
                                      |                     |                    |
-----------------                 ---------                 |        Etcd        |
|               |     Abnormal    |       |     Monitor     |                    |
| Custom Source |---------------->| Agent |---------------->|--------------------|
|               |                 |       |                 |                    |
-----------------                 ---------                 | Controller Manager |
                                      |                     |                    |
                                      |                     |--------------------|
                                      |                     |                    |
                                      |                     |      Scheduler     |
                                      |                     |                    |
                                      |                     ----------------------
                                     \|/
                            ---------------------
                            |                   |
                            |      Kernel       |
                            |      Docker       |
                            |      Kubelet      |
                            |      Cgroup       |
                            |      ......       |
                            |                   |
                            ---------------------
```

故障诊断恢复平台 Agent 组件由下列部分组成：

* 故障管理器（SourceManager）
* 故障分析链（DiagnoserChain）
* 信息管理器（InformationManager）
* 故障恢复链（RecovererChain）

```
-----------------             ----------------------             ------------------             ------------------
|               |  Abnormal   |                    |  Abnormal   |                |  Abnormal   |                |
| SourceManager |------------>| InformationManager |------------>| DiagnoserChain |------------>| RecovererChain |
|               |             |                    |             |                |             |                |
-----------------             ----------------------             ------------------             ------------------
                                         |                                |                              |
                                         |                                |                              |
                                         |                                |                              |
                                        \|/                              \|/                            \|/
                            --------------------------             ---------------                ---------------
                            |                        |             |             |                |             |
                            | InformationCollector 1 |             | Diagnoser 1 |                | Recoverer 1 |
                            |                        |             |             |                |             |
                            --------------------------             ---------------                ---------------
                                         |                                |                              |
                                         |                                |                              |
                                         |                                |                              |
                                        \|/                              \|/                            \|/
                            --------------------------             ---------------                ---------------
                            |                        |             |             |                |             |
                            | InformationCollector 2 |             | Diagnoser 2 |                | Recoverer 2 |
                            |                        |             |             |                |             |
                            --------------------------             ---------------                ---------------
                                         |                                |                              |
                                         |                                |                              |
                                         |                                |                              |
                                        \|/                              \|/                            \|/
                                      .......                          .......                        .......
```

### 功能

故障诊断恢复平台 Agent 组件功能如下：

* 获取 Event 等作为故障源。
* 监听故障诊断 Abnormal 自定义资源并进行处理和状态同步。
* 对本节点故障进行诊断和恢复。
* 通过 Abnormal 自定义资源以及 InformationCollector 自定义资源进行监控扩展和增强。

故障诊断恢复平台中 Abnormal 的状态迁移流程如下：

* 故障诊断恢复平台 Agent 或用户自行创建 Abnormal 自定义资源。
* 将 Abnormal 发送至信息管理器，标记 Abnormal 的状态为 InformationCollecting 并采集故障诊断恢复的信息。
  * 如果信息能够被成功采集则记录 InformationCollected 状况并继续。
  * 如果信息无法被成功采集则将 Abnormal 的状态标记为 Failed 并终止故障诊断恢复流程。
* 将 Abnormal 发送至故障分析链，标记 Abnormal 的状态为 Diagnosing 并对故障进行分析。
  * 如果故障能够被成功识别则记录 Identified 状况并继续。
  * 如果故障无法被成功识别则将 Abnormal 的状态标记为 Failed 并终止故障诊断恢复流程。
* 将 Abnormal 发送至故障恢复链，标记 Abnormal 的状态为 Recovering 并对故障进行恢复。
  * 如果故障能够被成功恢复则记录 Recovered 状况并标记 Abnormal 的状态为 Succeeded。
  * 如果故障无法被成功恢复则标记 Abnormal 的状态为 Failed 并终止故障诊断恢复流程。

故障诊断恢复平台中 Abnormal 的状态迁移图如下：

```
                                                                                                                                             ----------
                                                                                                                                             |        |
                                          -------------------------------------------------------------------------------------------------->| Failed |
                                         /|\                                    /|\                              /|\                         |        |
                                          |                                      |                                |                          ----------
                                   Failed |                               Failed |                         Failed |
                                          |                                      |                                |
-----------                   -------------------------                   --------------                   --------------                   -------------
|         |                   |                       |                   |            |                   |            |                   |           |
| Created |------------------>| InformationCollecting |------------------>| Diagnosing |------------------>| Recovering |------------------>| Succeeded |
|         |                   |                       |       /|\         |            |       /|\         |            |       /|\         |           |
-----------                   -------------------------        |          --------------        |          --------------        |          -------------
                                          |                    |                 |              |                 |              |
                             Successfully |                    |    Successfully |              |    Successfully |              |
                                          |                    |                 |              |                 |              |
                                         \|/                   |                \|/             |                \|/             |
                               ------------------------        |          --------------        |           -------------        |
                               |                      |        |          |            |        |           |           |        |
                               | InformationCollected |---------          | Identified |---------           | Recovered |---------
                               |                      |                   |            |                    |           |
                               ------------------------                   --------------                    -------------
```

### Abnormal 自定义资源

Abnormal 是故障诊断恢复平台中故障管理器、故障分析链、故障恢复链之间通信的接口。故障事件的详情记录在 Spec 中，故障管理器、故障分析链和故障恢复链对 Abnormal 进行处理并通过变更 Status 字段进行通信。详细信息参考 [Abnormal API 设计](./abnormal.md)。

### InformationCollector 自定义资源

InformationCollector 自定义资源用于注册信息采集器，信息采集器的元数据记录在 Spec 中，包括发现方式和监听地址。InformationCollector 的当前状态记录在 Status 字段。详细信息参考 [InformationCollector API 设计](./information-collector.md)。

### Diagnoser 自定义资源

Diagnoser 自定义资源用于注册故障分析器，故障分析器的元数据记录在 Spec 中，包括发现方式和监听地址。Diagnoser 的当前状态记录在 Status 字段。详细信息参考 [Diagnoser API 设计](./diagnoser.md)。

### Recoverer 自定义资源

Recoverer 自定义资源用于注册故障恢复器，故障恢复器的元数据记录在 Spec 中，包括发现方式和监听地址。Recoverer 的当前状态记录在 Status 字段。详细信息参考 [Recoverer API 设计](./recoverer.md)。

### 故障管理器

故障管理器是获取故障事件的接口，大致可分为以下几类：

* Event：Kubernetes 中的事件支持更细致的故障上报机制。
* Custom：用于自定义故障，用户可以自定义进行扩展。

故障管理器在消费 Event 后会生成 Abnormal 故障事件并发往故障分析链。用户也可以实现故障事件源并直接通过自定义资源来创建故障事件。

### 信息管理器

当故障分析或恢复流程较复杂时，需要从其他接口获取更多信息用于故障的诊断和确认。此时故障分析器或故障恢复器可以调用信息采集器获取更多信息。如果信息无法被任何信息采集器获取则报错并中止。信息采集器一般是一个 HTTP 服务器。常见的信息采集器功能包括：

* Golang 剖析文件采集。
* Java 虚拟机诊断。
* eBPF 信息采集。
* PSI 信息采集。

### 故障分析链

故障分析链是一个调用链框架，本身并不包含故障分析的逻辑，用户需要实现故障分析的具体逻辑并注册到故障分析链中。故障分析链从故障管理器接收故障事件并将故障事件逐一传入被注册的故障分析器中，当故障能够被某个故障分析器识别则中止调用并交由该逻辑进行处理。如果故障无法被任何故障分析器识别则报错并中止。故障分析器一般是一个 HTTP 服务器。常见的故障分析器功能包括：

* Docker 问题分析。
* Kubernetes 问题分析。
* 内核日志分析。
* Pod 磁盘空间使用量分析。

故障分析器在无法识别 Abnormal 故障事件时返回错误，故障分析器在成功识别 Abnormal 故障事件后变更 Status 字段。故障分析链在某个故障分析器成功识别 Abnormal 故障事件后将 Abnormal 故障事件发往故障恢复链。故障分析器在执行诊断时可以通过调用信息采集器获取更多信息。

### 故障恢复链

故障恢复链是一个调用链框架，本身并不包含故障恢复的逻辑，用户需要实现故障恢复的具体逻辑并注册到故障恢复链中。故障恢复链从故障分析链接收故障事件并将故障事件逐一传入被注册的故障恢复器中，当故障能够被某个故障恢复器识别则中止调用并交由该逻辑进行处理。如果故障无法被任何故障恢复器恢复则报错并中止。故障恢复器一般是一个 HTTP 服务器。常见的故障恢复器功能包括：

* 故障进程清理。
* 残余日志文件清理。

## 典型用例

社区 [Issue 3529](https://github.com/containerd/containerd/pull/3529) 中记录的 Bug 会导致 Docker 18.06 及以下版本因为 Shim 的死锁而无法正常终止容器。当该故障出现时可以通过以下步骤实现该故障的诊断恢复：

* 实现故障事件源：获取当前 Terminating 状态的 Pod，当该类 Pod 存在时创建表示该故障的 Abnormal，标记 `.spec.nodeName` 字段为该 Pod 所处节点。
* 实现故障分析器：首先通过进程树查找到该容器对应的 Shim 进程。然后向该 Shim 进程发送 SIGUSR1 信号获取栈信息，在栈信息中查找 `reaper.go` 相关函数以确定问题原因。
* 实现故障恢复器：杀死该容器对应的 Shim 进程进行恢复。