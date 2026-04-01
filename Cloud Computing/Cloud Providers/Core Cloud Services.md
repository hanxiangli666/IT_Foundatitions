# Core Cloud Services

> 核心云服务总览：以“计算、存储、数据库、网络”四大域为主线，建立 AWS、Azure、Google Cloud 的跨平台对照认知，并补充从课堂演示走向实际设计所需的方法。
>
> Core cloud services overview: organize knowledge around Compute, Storage, Databases, and Networking, map capabilities across AWS, Azure, and Google Cloud, and extend classroom demos into practical design thinking.

## 1. 学习目标 / Learning Objectives

本节的目标不是只记住服务名，而是建立“能力层级”的理解方式：当你看到任何云产品时，先判断它属于计算、存储、数据库还是网络，再比较其性能、成本、安全和运维特性。

The goal is not merely memorizing service names, but building capability-first thinking: for any cloud product, identify whether it belongs to Compute, Storage, Databases, or Networking, then compare performance, cost, security, and operations.

通过这个方法，你可以快速把 AWS、Azure、GCP 的同类服务对应起来，并在后续课程中更快理解新服务。

With this method, you can quickly map equivalent services across AWS, Azure, and GCP and absorb new services faster in later modules.

![1774969308099](image/CoreCloudServices/1774969308099.png)

## 2. 三大云厂商与市场背景 / Big-3 Cloud Providers and Market Context

当前主流公有云平台主要是 AWS、Azure 与 Google Cloud。它们都提供类似的核心能力，但在生态、企业集成、数据平台与默认工具链上各有优势。

The major public cloud providers are AWS, Azure, and Google Cloud. They offer similar core capabilities but differ in ecosystem strength, enterprise integration, data platform maturity, and default tooling.

从历史演进看，AWS 率先规模化基础设施即服务；Azure 依托微软企业生态快速扩张；GCP 在数据分析、平台工程与机器学习领域形成特色。

Historically, AWS scaled infrastructure services early; Azure expanded rapidly through Microsoft enterprise ecosystems; GCP built strong differentiation in analytics, platform engineering, and machine learning.

课程中的关键结论是：企业并不一定只选一家云。多云策略常由合规、成本、区域覆盖、团队技能和历史系统共同决定。

A key takeaway is that many organizations do not stay single-cloud. Multi-cloud choices are often driven by compliance, cost, regional availability, team skills, and legacy constraints.

## 3. 四大核心服务域总览 / Four Core Domains at a Glance

四大域是后续所有云学习的坐标系。你可以把它们理解为“云上系统的四个基本层”。

These four domains form the coordinate system for all further cloud learning. Think of them as the four base layers of cloud systems.

| Domain | AWS | Azure | Google Cloud |
| --- | --- | --- | --- |
| Compute | EC2, ECS/EKS, Lambda | Virtual Machines, AKS, Functions | Compute Engine, GKE, Cloud Functions |
| Storage | S3, EBS, EFS | Blob Storage, Managed Disks, Azure Files | Cloud Storage, Persistent Disk, Filestore |
| Databases | RDS, DynamoDB | Azure SQL, Cosmos DB | Cloud SQL, Firestore |
| Networking | VPC, Security Groups, NACLs | VNet, NSG | VPC, Firewall Rules |

你应该优先记“功能对应关系”而不是“品牌名”。这会显著降低跨平台学习成本。

You should prioritize functional mapping over product branding. This significantly lowers cross-platform learning cost.

## 4. 计算服务 / Compute Services

计算层回答的是“代码运行在哪里”。常见模式是虚拟机、容器、无服务器函数。

Compute answers one question: where does your code run? Typical patterns are virtual machines, containers, and serverless functions.

虚拟机适合需要较高系统控制能力的场景；容器适合多服务协同与持续交付；无服务器适合事件驱动、间歇流量和短任务执行。

VMs fit scenarios requiring stronger system-level control; containers fit multi-service delivery and CI/CD; serverless fits event-driven, bursty, short-lived workloads.

课堂强调了一个非常实用的思路：计算选型不仅看技术偏好，还要看团队运维能力和可观测性成熟度。

The class emphasized a practical principle: compute selection is not just a technical preference, but also a decision about team operational maturity and observability readiness.

![1774969325875](image/CoreCloudServices/1774969325875.png)

## 5. 存储服务 / Storage Services

存储应先按类型划分，再选产品。对象存储适合文件与归档；块存储适合挂载给虚机作为系统盘或数据盘；文件存储适合多实例共享目录。

For storage, classify by storage type before selecting product. Object storage is for files and archives; block storage is attached to VMs as system/data disks; file storage is for shared filesystem semantics.

对象存储是课堂实操中的重点，因为它与成本控制关系最直接：存储类别、生命周期策略、访问频率和数据保留期限共同决定长期费用。

Object storage is a classroom focus because it most directly links to cost control: storage class, lifecycle rules, access frequency, and retention horizon jointly shape long-term spending.

![1774969337959](image/CoreCloudServices/1774969337959.png)

![1774969348356](image/CoreCloudServices/1774969348356.png)

### 5.1 存储成本补充 / Storage Cost Add-on

很多新手只看每 GB 价格，但总成本还包含请求费用、检索费用、最短存储时长、跨区域流量等。

Many beginners only look at per-GB pricing, but total cost also includes request charges, retrieval fees, minimum storage duration, and cross-region transfer.

建议先画出数据访问曲线与保留周期，再映射到对应的存储层级。

A practical method is to model access patterns and retention windows first, then map data to proper storage tiers.

## 6. 数据库服务 / Database Services

云数据库的价值是把大量重复运维工作托管化，比如备份、打补丁、故障转移与高可用。

Managed cloud databases reduce repetitive operations such as backups, patching, failover, and high-availability management.

关系型数据库适合强事务与复杂查询；NoSQL 适合高吞吐、弹性扩展和灵活模式。

Relational databases fit strong transactions and complex querying; NoSQL fits high throughput, elasticity, and schema flexibility.

选型时建议优先确认一致性需求、读写比例、查询复杂度与延迟目标。

During selection, prioritize consistency requirements, read/write ratio, query complexity, and latency targets.

![1774969362067](image/CoreCloudServices/1774969362067.png)

## 7. 网络服务 / Networking Services

网络层用于连接与隔离资源，是安全与性能的共同基础。VPC 或 VNet 定义边界，安全组或防火墙规则控制流量，路由决定路径。

Networking connects and isolates resources, forming the shared base for both security and performance. VPC/VNet defines boundaries, security rules control traffic, and routing determines paths.

课堂强调：云网络并不是全新知识，而是把传统网络原理在更大规模上自动化。

The class emphasized that cloud networking is not entirely new knowledge; it is traditional networking principles applied at scale with automation.

![1774969374707](image/CoreCloudServices/1774969374707.png)

## 8. 从课堂到实战的补充框架 / Practical Extension Beyond the Lesson

建议你在每个资源上都问四个问题：

Use these four questions for every resource you create:

1. 业务目标是什么，如何度量成功？
2. 谁可以访问，如何最小权限？
3. 成本由哪些因素驱动，是否自动化治理？
4. 出现故障如何恢复，恢复是否演练过？

1. What is the business objective, and how is success measured?
2. Who can access it, and how is least privilege enforced?
3. What drives cost, and is cost governance automated?
4. How does recovery work, and has recovery been practiced?

当你能用这四问复盘 VM、对象存储、数据库和网络配置时，你就不再只是“会点控制台”，而是具备了初级云架构思维。

When you can review VMs, object storage, databases, and networking using these questions, you move beyond console clicking and into foundational cloud-architecture thinking.

## 9. 课堂回顾与行动建议 / Recap and Actionable Next Steps

本课核心是同一能力在不同平台的映射。真正可迁移的是抽象层能力，而不是厂商命名。

The lesson core is capability mapping across platforms. What truly transfers is abstraction-level thinking, not vendor-specific naming.

建议课后完成一个最小实操闭环：创建资源、验证可用、记录成本、清理资源、复盘改进点。

For post-class practice, complete a minimal loop: create resources, validate behavior, record costs, clean up, and run a short retrospective.

完成这个闭环后，你对云平台的理解会从“知识点”升级为“可执行能力”。

After completing this loop, your understanding shifts from isolated knowledge points to executable capability.
