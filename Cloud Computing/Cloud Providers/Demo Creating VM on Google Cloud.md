# Demo Creating VM on Google Cloud

> GCP 虚机实操增强笔记：围绕创建、配置、SSH 验证、停止删除与磁盘清理，构建“可运行 + 可控成本 + 可复盘”的基础运维闭环。
>
> Enhanced GCP VM hands-on note: from creation and configuration to SSH validation, stop/delete actions, and disk cleanup, building a practical loop of operability, cost control, and reviewability.

## 1. 本节目标 / Lesson Goal

本节不是单纯创建一台机器，而是学习如何把一台云主机从“开通”推进到“验证、管理、清理”的完整生命周期。

This lesson is not only about launching a VM, but about managing its full lifecycle from provisioning to validation, operations, and cleanup.

你需要建立一个意识：在云平台里，资源创建动作就是成本动作，资源保留就是持续费用。

You need a key mindset: in cloud environments, creation is a cost action, and retention means ongoing spend.

## 2. 在 Compute Engine 创建实例 / Create an Instance in Compute Engine

创建实例时需要选择区域、可用区、机器类型、启动盘大小、镜像系统、网络设置等参数。

When creating an instance, you choose region, zone, machine type, boot disk size, OS image, networking, and optional features.

课堂示例选择 e2-medium（2 vCPU、4 GB RAM），这是开发与教学中常见的入门规格。

The classroom example used e2-medium (2 vCPU, 4 GB RAM), a common starter profile for learning and development.

![1774969511266](image/DemoCreatingVMonGoogleCloud/1774969511266.png)

### 2.1 机型选择补充 / Extra Notes on Machine Type Selection

通用机型适合多数基础业务；计算优化机型适合 CPU 密集任务；内存优化机型适合缓存与大内存数据库；GPU 机型用于训练或图形计算。

General-purpose shapes fit most baseline workloads; compute-optimized types fit CPU-bound tasks; memory-optimized types fit cache and large-memory databases; GPU types support ML and graphics workloads.

选型时建议同时看峰值需求和长期平均负载，避免长期过配。

When selecting machine types, evaluate both peak demand and long-term average load to avoid chronic overprovisioning.

## 3. 实例创建后的状态核查 / Post-creation State Verification

实例创建成功后，先在 VM 列表确认状态为 Running，再核对内网 IP、外网 IP、区域与可用区。

After creation, verify Running state in the VM list, then check internal/external IPs, region, and zone.

课堂里展示了实例总览和关联动作入口，这一步帮助你建立“资源可视化管理”习惯。

The class showed instance overview and related action panels, helping build a habit of visual resource governance.

![1774969526504](image/DemoCreatingVMonGoogleCloud/1774969526504.png)

## 4. 通过 SSH 登录并验证资源 / SSH In and Validate Resources

使用浏览器 SSH 可以快速连接实例，适合实验或临时排障。连接后建议立即执行容量核查命令，验证磁盘和内存是否与配置一致。

Browser SSH gives fast access for labs and ad-hoc troubleshooting. After login, run capacity checks to confirm disk and memory match expected configuration.

课堂示例中通过安装小工具验证了网络和包管理链路可用，这等同于一次最小化环境验收。

The demo installed a small package to verify outbound network and package management, equivalent to a minimal environment acceptance check.

关闭 SSH 窗口并不会停止实例。只要实例仍在运行，计算费用就会继续累计。

Closing the SSH session does not stop the VM. As long as the instance is running, compute charges continue.

## 5. 停止、删除与磁盘保留策略 / Stop, Delete, and Disk Retention Policy

停止实例会暂停计算资源计费，但磁盘等持久资源通常仍计费。删除实例时需确认启动盘删除策略，否则可能遗留磁盘产生持续费用。

Stopping an instance usually halts compute charges, but persistent disks still incur cost. During deletion, verify boot disk deletion policy to avoid orphaned storage spend.

课堂演示了在 Disks 页面检查是否有残留磁盘，这一步是实验环境成本控制的关键动作。

The lesson demonstrated checking the Disks page for leftovers, a critical step for cost hygiene in lab environments.

![1774969549895](image/DemoCreatingVMonGoogleCloud/1774969549895.png)

## 6. 课堂外补充：生产环境最小实践 / Extra: Minimum Production Practices

建议使用最小权限服务账号，避免长期共享高权限凭据；对 SSH 访问建立统一入口与审计机制。

Use least-privilege service accounts and avoid long-lived shared privileged credentials; establish controlled SSH entry paths with auditing.

对关键实例启用快照计划、监控告警与自动化停机策略，减少故障损失和空转成本。

Enable snapshot schedules, monitoring alerts, and automated shutdown policies for critical instances to reduce both outage impact and idle cost.

如果工作负载波动明显，可结合可抢占实例或计划任务调度进一步优化成本。

For variable workloads, combine preemptible/spot-like options and schedule-based automation for better cost efficiency.

## 7. 课堂回顾 / Lesson Recap

本节核心能力可以归纳为三点：

The lesson outcome can be summarized in three capabilities:

1. 能按需求选择并创建合适规格的 VM。
2. 能通过 SSH 与系统命令验证资源状态。
3. 能正确执行停止、删除与磁盘清理，避免隐性成本。

1. Create right-sized VMs for workload needs.
2. Validate instance resources via SSH and system checks.
3. Execute stop/delete/disk cleanup correctly to avoid hidden costs.

把这三点做成固定流程，你的云主机运维会更稳定，也更容易标准化。

Turning these three points into a repeatable process makes VM operations more stable and easier to standardize.
