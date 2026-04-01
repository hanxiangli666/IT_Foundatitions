# Cloud Providers - Bilingual Extended Notes

> 适用范围 / Scope:
> 这份笔记整合并扩展了 Cloud Providers 模块中的课堂内容，覆盖核心云服务、GCP 虚机演示、AWS S3 演示，并补充跨云对照、设计决策、成本与安全实践。
>
> This note consolidates and extends the Cloud Providers module, covering core cloud services, the GCP VM demo, the AWS S3 demo, and additional guidance on cross-cloud mapping, design decisions, and cost-security practices.

---

## 0. 学习目标与课程定位 / Learning Objectives and Course Positioning

本节课程的核心目标有两个：第一，识别 AWS、Azure、Google Cloud 在核心服务上的对应关系；第二，能够亲手创建、使用并清理简单云资源，形成从概念到操作的闭环。

This lesson has two core goals: first, recognize the service mapping across AWS, Azure, and Google Cloud; second, create, use, and clean up simple cloud resources to build a complete concept-to-practice loop.

云学习最常见的障碍不是“看不懂某个服务名”，而是“把服务名当成概念本身”。例如，很多初学者记住了 EC2、S3、RDS，却没有建立“计算、存储、数据库、网络”四大抽象层。课程实际上在训练你用抽象层来理解具体产品。

A common cloud-learning obstacle is not misunderstanding service names, but mistaking service names for the concepts themselves. Many beginners memorize EC2, S3, and RDS without building the four-layer abstraction: compute, storage, databases, and networking. This lesson trains you to think in abstractions, not brand terms.

![课程目标截图](image/CoreCloudServices/1774969308099.png)

---

## 1. 云市场背景与三大厂商定位 / Market Context and Big-3 Positioning

全球公有云长期由三大厂商主导：AWS、Azure、Google Cloud。它们都提供同一组基础能力，但在生态、默认集成、企业合作路径和定价策略上各有侧重。

The public cloud market has long been dominated by AWS, Azure, and Google Cloud. They provide the same foundational capabilities, but differ in ecosystem strength, default integrations, enterprise adoption paths, and pricing strategies.

从时间线上看，AWS 在 2006 年较早商业化基础设施服务，Azure 在 2010 年依托微软企业生态发力，Google Cloud 在 2011 年后凭借数据与平台工程优势逐步扩大影响力。

From a timeline perspective, AWS commercialized infrastructure services early in 2006, Azure accelerated from 2010 through Microsoft enterprise integration, and Google Cloud expanded from 2011 onward with strengths in data and platform engineering.

课堂上的重要结论不是“谁最强”，而是“为什么企业常常是多云”。现实中，团队会因历史系统、数据合规、区域可用性、议价能力和人才储备而采用多云或混合云策略。

The key takeaway is not "who is strongest," but why organizations often adopt multi-cloud. In practice, teams choose multi-cloud or hybrid strategies due to legacy systems, data compliance, regional availability, negotiation leverage, and available talent.

---

## 2. 四大核心服务域总览 / Four Core Service Domains Overview

四大核心服务域是本模块的主线：计算、存储、数据库、网络。只要你能稳定地从这四类角度描述任何云方案，你就已经具备云架构的入门表达能力。

The module is organized around four core domains: compute, storage, databases, and networking. If you can describe any cloud design through these four categories, you already have foundational cloud-architecture literacy.

| 领域 Domain      | AWS                       | Azure                                    | Google Cloud                              |
| ---------------- | ------------------------- | ---------------------------------------- | ----------------------------------------- |
| 计算 Compute     | EC2, ECS/EKS, Lambda      | Virtual Machines, AKS, Functions         | Compute Engine, GKE, Cloud Functions      |
| 存储 Storage     | S3, EBS, EFS              | Blob Storage, Managed Disks, Azure Files | Cloud Storage, Persistent Disk, Filestore |
| 数据库 Databases | RDS, DynamoDB             | Azure SQL, Cosmos DB                     | Cloud SQL, Firestore                      |
| 网络 Networking  | VPC, Security Group, NACL | VNet, NSG                                | VPC, Firewall Rules                       |

这个映射表建议反复使用：先看业务需求属于哪一类，再在该云平台里找对应产品，而不是先背产品名后猜用途。

Use this mapping table repeatedly: classify the business need first, then find the matching service on that platform. Do not memorize product names first and guess usage later.

---

## 3. 计算服务 Compute：从 VM 到容器再到 Serverless / Compute: VM, Containers, and Serverless

计算服务的本质是让代码“有地方可运行”。差别在于你管理多少层：VM 管理更多基础层，容器强调可移植和编排，Serverless 强调按请求触发和按量计费。

The essence of compute is giving your code a place to run. The difference is operational responsibility: VMs expose more infrastructure control, containers emphasize portability and orchestration, and serverless emphasizes event-driven execution and pay-per-use.

VM 方案适合需要明确控制 OS、内核参数、软件栈的场景；容器方案适合多服务协同、发布频繁、环境一致性要求高的场景；Serverless 适合突发流量和短任务，避免空闲机器成本。

VMs fit scenarios requiring control over OS, kernel settings, and software stack; containers fit multi-service systems with frequent releases and strong environment consistency needs; serverless fits bursty traffic and short-lived tasks while avoiding idle host costs.

课堂中展示了跨云的计算产品对应关系，并强调了“名称不同，能力同类”。你应当把注意力放在扩缩容策略、可观测性、失败恢复、计费曲线四个维度上。

The lesson mapped compute services across clouds and emphasized "different names, similar capability classes." Focus your decisions on scaling strategy, observability, failure recovery, and billing profile.

![计算服务映射截图](image/CoreCloudServices/1774969325875.png)

### 3.1 课堂外补充：计算选型速查 / Extra: Compute Selection Quick Guide

如果你需要稳定长连接、后台常驻进程和完整运维控制，优先 VM；如果你有大量微服务并强调发布效率，优先容器；如果你的业务是事件触发、请求粒度小、峰谷明显，优先 Serverless。

If you need stable long-running processes and full operational control, prefer VMs; if you run many microservices with high release cadence, prefer containers; if your workload is event-driven with sharp traffic variance, prefer serverless.

不要忽略冷启动、配额限制、并发上限、镜像大小、区域可用性这些二阶因素，它们往往在 PoC 阶段不显著，但会在生产环境放大。

Do not ignore second-order factors such as cold starts, quotas, concurrency limits, image size, and regional availability. They may be invisible in PoCs but become critical in production.

---

## 4. 存储服务 Storage：对象、块、文件三分法 / Storage: Object, Block, and File

课堂对存储的关键认知是“先分类型，再谈产品”。对象存储用于文件和非结构化数据，块存储用于挂载到虚机作为磁盘，文件存储用于多实例共享目录。

The key storage concept is "classify storage type before selecting products." Object storage handles files and unstructured data, block storage is attached as disks to VMs, and file storage supports shared directory semantics across instances.

对象存储（S3 / Blob / Cloud Storage）通常是云上最先接触的服务，适合备份、日志、媒体文件、静态站点资源。其优势是高扩展、按量计费、生命周期管理和多存储层级。

Object storage (S3 / Blob / Cloud Storage) is usually the first cloud storage service teams adopt. It fits backups, logs, media assets, and static website artifacts, with strong elasticity, consumption billing, lifecycle policies, and storage-class options.

块存储（EBS / Managed Disks / Persistent Disk）通常与 VM 生命周期相关，IO 性能和快照策略很关键。文件存储（EFS / Azure Files / Filestore）则更接近传统 NAS 使用体验。

Block storage (EBS / Managed Disks / Persistent Disk) is tightly coupled with VM lifecycle and emphasizes IOPS plus snapshot strategy. File storage (EFS / Azure Files / Filestore) resembles managed NAS behavior.

![存储对象服务截图](image/CoreCloudServices/1774969337959.png)

![块存储概念截图](image/CoreCloudServices/1774969348356.png)

### 4.1 课堂外补充：存储成本模型 / Extra: Storage Cost Model

存储成本不仅是“每 GB 单价”，还包括请求次数（PUT/GET）、跨区传输、检索费用、最短存储时长约束以及对象数量带来的元数据与监控开销。

Storage cost is not just cost per GB. It also includes request charges (PUT/GET), cross-region transfer, retrieval fees, minimum storage-duration constraints, and metadata/monitoring overhead at high object counts.

你应当为每类数据定义“访问频率曲线”和“保留时间曲线”，再映射到具体存储类和生命周期规则，这比手工迁移更稳健。

Define an access-frequency curve and a retention curve for each data type, then map them to storage classes and lifecycle rules. This is more reliable than manual migration.

---

## 5. 数据库服务 Databases：关系型与 NoSQL 的云上对照 / Databases: Relational and NoSQL in Cloud

课堂展示了 RDS / Azure SQL / Cloud SQL 和 DynamoDB / Cosmos DB / Firestore 的对应关系。核心思路是先判定数据模型与事务需求，再决定引擎与托管方式。

The lesson mapped RDS / Azure SQL / Cloud SQL and DynamoDB / Cosmos DB / Firestore. The main decision flow is to determine data model and transactional requirements first, then choose engine and managed offering.

关系型数据库更适合强一致事务和复杂查询；NoSQL 更适合高并发、弹性扩展、半结构化数据和低运维负担场景。

Relational databases fit strongly consistent transactions and complex queries; NoSQL fits high throughput, elastic scaling, semi-structured data, and low operational overhead.

托管数据库节省了补丁、备份、故障切换、版本升级等大量运维工作，但你仍需要负责表结构设计、索引策略、查询优化和数据生命周期。

Managed databases reduce heavy tasks such as patching, backups, failover, and version upgrades, but you still own schema design, indexing strategy, query performance, and data lifecycle.

![数据库服务映射截图](image/CoreCloudServices/1774969362067.png)

---

## 6. 网络服务 Networking：连接、隔离、访问控制 / Networking: Connectivity, Isolation, and Access Control

网络层是所有云资源协同的基础。你在云上做的每一次“可以访问”与“不能访问”，本质都是网络策略与身份策略共同作用的结果。

Networking is the foundation of all cloud resource interaction. Every "reachable" or "blocked" outcome in cloud environments is the combined result of network policy and identity policy.

课堂中对比了 AWS VPC + Security Groups、Azure VNet + NSG、GCP VPC + Firewall Rules。虽然术语不同，但都在解决同一件事：如何把流量限制在最小必要范围。

The lesson compared AWS VPC + Security Groups, Azure VNet + NSG, and GCP VPC + Firewall Rules. Terminology differs, but they solve the same problem: limiting traffic to least-necessary scope.

建议你在学习阶段建立一个固定思维顺序：子网划分 -> 路由 -> 入站/出站规则 -> 主机级策略 -> 审计日志。这个顺序可以显著降低排障难度。

Adopt a stable troubleshooting order while learning: subnet layout -> routing -> ingress/egress rules -> host-level policy -> audit logs. This sequence reduces diagnostic confusion.

![网络服务映射截图](image/CoreCloudServices/1774969374707.png)

---

## 7. 实操一：在 GCP 创建并管理 VM / Hands-on 1: Create and Manage a VM on GCP

本段对应课堂演示中“Launching VM in GCP”的完整流程，包含创建、验证、SSH 连接、停止与删除。

This section maps to the classroom demo "Launching VM in GCP," including creation, validation, SSH access, stop, and delete operations.

### 7.1 创建实例前的配置思维 / Pre-creation Configuration Thinking

在 Compute Engine 创建实例时，你需要一次性决定多个关键参数：区域与可用区、机器类型、启动磁盘、网络接口、防火墙标签，以及是否分配外网 IP。

When creating a Compute Engine instance, you choose several critical parameters: region/zone, machine type, boot disk, network interface, firewall tags, and whether to assign an external IP.

课堂示例选择了 e2-medium（2 vCPU、4GB RAM），并展示了控制台中按机型估算月成本的视图。这一点非常重要，因为云资源的创建动作本质就是成本决策动作。

The classroom example selected e2-medium (2 vCPU, 4 GB RAM) and showed the monthly estimate in console. This is important because creating cloud resources is fundamentally a cost decision.

![GCP 创建实例与价格估算](image/DemoCreatingVMonGoogleCloud/1774969511266.png)

### 7.2 创建后检查：实例与磁盘状态 / Post-creation Checks: Instance and Disk State

实例创建成功后，应立即检查三类信息：实例是否 Running、内外网 IP 是否符合预期、启动盘是否挂载且容量正确。

After instance creation, immediately verify three categories: running status, expected internal/external IP assignment, and boot-disk attachment with correct capacity.

课堂中强调了一个经常被忽视的点：即便实例停止，磁盘通常仍保留并持续计费，除非你在删除策略中明确设定随实例删除。

The lesson emphasized a commonly missed point: even if the instance is stopped, the disk usually remains and may keep incurring charges unless deletion policy is explicitly configured.

![GCP VM 列表与状态](image/DemoCreatingVMonGoogleCloud/1774969526504.png)

### 7.3 SSH 登录与资源验证 / SSH Access and Resource Verification

通过浏览器内置 SSH 可以快速进入实例，适合教学与临时运维。进入后建议先用 df -h 和 free -m 验证磁盘和内存是否与配置一致。

Browser-based SSH is convenient for training and ad-hoc operations. Once inside, run df -h and free -m to verify disk and memory match requested configuration.

课堂演示安装了一个轻量软件（ninvaders）来证明实例可正常联网和安装包，这在实务中相当于最小化功能验收。

The demo installed a lightweight package (ninvaders) to verify network and package-manager functionality, which serves as a minimal functional acceptance test in real operations.

### 7.4 停止、删除、清理 / Stop, Delete, and Cleanup

关闭 SSH 窗口不等于停止实例。实例仍然运行并计费，必须在控制台执行 Stop 或 Delete。

Closing the SSH terminal does not stop the VM. The instance keeps running and billing until you explicitly Stop or Delete it.

Delete 之后请进入 Disks 页面核对是否仍有遗留磁盘，避免实验环境长期累积“看不见但持续收费”的资源。

After deleting the VM, always check the Disks page for leftovers, preventing hidden persistent-cost accumulation in lab environments.

![GCP 磁盘清理页面](image/DemoCreatingVMonGoogleCloud/1774969549895.png)

### 7.5 课堂外补充：GCP VM 最佳实践 / Extra: GCP VM Best Practices

生产环境建议使用最小权限服务账号、禁用不必要外网暴露、通过 IAP 或堡垒机统一入口，并配置快照策略和告警。

For production, use least-privilege service accounts, avoid unnecessary public exposure, centralize access via IAP or bastion patterns, and configure snapshot plus alerting policies.

对于间歇负载，优先考虑自动停机计划或可抢占实例策略，以降低闲置成本。

For intermittent workloads, consider scheduled shutdown or preemptible/spot strategies to reduce idle costs.

---

## 8. 实操二：在 AWS 创建并管理 S3 Bucket / Hands-on 2: Create and Manage an S3 Bucket on AWS

本段对应课堂演示中“Uploading files to AWS S3 + Security settings + Storage classes + Lifecycle rules”的完整流程。

This section maps to the classroom sequence: uploading files to S3 plus security settings, storage classes, and lifecycle rules.

### 8.1 创建 Bucket 与命名规则 / Bucket Creation and Naming Rules

S3 Bucket 名称必须全局唯一，课堂演示中特别展示了重名提示。建议采用组织前缀 + 环境 + 业务 + 随机尾缀的命名规范，降低冲突概率。

S3 bucket names must be globally unique, and the demo explicitly showed duplicate-name prompts. A naming pattern like org-prefix + environment + workload + random suffix significantly reduces collisions.

创建时优先保持 Object Ownership 为 Bucket owner enforced，并保持 Block Public Access 默认开启，这是面向初学者最稳妥的安全基线。

During creation, keep Object Ownership as Bucket owner enforced and keep Block Public Access enabled by default. This is the safest baseline for beginners.

![S3 创建 Bucket 页面](image/DemoCreatingS3BucketonAWS/1774969402294.png)

### 8.2 上传对象与查看属性 / Uploading Objects and Inspecting Properties

上传对象后，建议检查对象大小、时间戳、加密状态、存储类别、版本信息。这些元数据直接影响后续的取回性能、合规审计和生命周期动作。

After uploading objects, inspect size, timestamps, encryption status, storage class, and version information. These metadata fields directly affect retrieval behavior, compliance audits, and lifecycle transitions.

课堂中也强调了权限验证：当公共访问被阻止时，匿名访问对象 URL 会被拒绝，这是预期行为而非错误。

The class also emphasized permission validation: when public access is blocked, anonymous object URL access is denied by design, not by accident.

### 8.3 存储类别与价格理解 / Understanding Storage Classes and Pricing

S3 提供多种存储类别，不同类别在单位存储成本、检索延迟、检索费用、最短存储时长方面差异明显。课堂演示通过定价页面帮助你建立这组权衡关系。

S3 offers multiple storage classes with major differences in per-GB cost, retrieval latency, retrieval fees, and minimum-duration constraints. The class used the pricing page to build this tradeoff awareness.

建议将“访问频率”作为第一决策变量：高频访问选择 Standard，访问模式不确定选择 Intelligent-Tiering，长期归档选择 Glacier 系列。

Use access frequency as your first decision variable: Standard for frequent access, Intelligent-Tiering for uncertain patterns, and Glacier tiers for long-term archival.

![S3 存储定价页面](image/DemoCreatingS3BucketonAWS/1774969414895.png)

### 8.4 生命周期规则 Lifecycle Rules：自动化降本关键点 / Lifecycle Rules: Core Cost-Automation Mechanism

生命周期规则是本课最重要的实务技能之一。它可以在对象达到某个年龄后自动转存到更低价层，或按保留策略自动过期删除。

Lifecycle rules are one of the most practical skills in this lesson. They can automatically transition aged objects to cheaper classes or expire data according to retention policy.

课堂示例创建了一个 30 天后转 Glacier Instant Retrieval 的规则，并在规则列表中验证生效状态。这是“成本治理可自动化”的典型案例。

The demo created a 30-day transition rule to Glacier Instant Retrieval and validated it in the lifecycle list. This is a classic example of automating cost governance.

![S3 生命周期规则创建](image/DemoCreatingS3BucketonAWS/1774969464479.png)

![S3 生命周期规则生效](image/DemoCreatingS3BucketonAWS/1774969475719.png)

### 8.5 删除对象与删除 Bucket / Delete Objects and Delete Bucket

删除 Bucket 前必须先清空对象；如果启用了版本控制，还必须清理所有版本与 Delete Markers。课堂的关键提醒是：云资源不会自动回收，清理动作必须显式执行。

Before deleting a bucket, you must empty objects; with versioning enabled, all versions and delete markers must also be removed. The key classroom reminder is that cloud resources are not auto-recycled; cleanup must be explicit.

### 8.6 课堂外补充：S3 安全与治理建议 / Extra: S3 Security and Governance Recommendations

建议默认启用服务端加密（SSE-S3 或 SSE-KMS），对敏感数据使用 KMS CMK 并配置最小权限 IAM 策略。

Enable server-side encryption by default (SSE-S3 or SSE-KMS). For sensitive data, use KMS CMKs with least-privilege IAM policies.

如需公开对象，优先使用 CloudFront + OAC 的受控分发方式，而非直接开放 Bucket 公网读取。

If public delivery is required, prefer CloudFront with OAC for controlled distribution rather than direct bucket-level public access.

通过标签、生命周期、清理脚本和预算告警的组合，可以把对象存储从“易失控成本项”变成“可预测成本项”。

Using tags, lifecycle policies, cleanup scripts, and budget alerts together turns object storage from an unpredictable cost center into a manageable one.

![课堂演示总结页](image/DemoCreatingS3BucketonAWS/1774969491293.png)

---

## 9. 跨云命名翻译表（实战记忆版） / Cross-Cloud Naming Translation (Practical Memory Version)

学习云平台最容易卡在“同物不同名”。下面是课堂内容扩展后的高频翻译表，建议作为术语速查。

A frequent blocker in cloud learning is "same capability, different names." The table below extends classroom content into a practical high-frequency translation guide.

| 能力 Capability                   | AWS      | Azure            | Google Cloud    | 说明 Notes        |
| --------------------------------- | -------- | ---------------- | --------------- | ----------------- |
| 虚拟机 VM                         | EC2      | Virtual Machines | Compute Engine  | IaaS 计算基础     |
| 容器编排 Kubernetes               | EKS      | AKS              | GKE             | 托管 K8s 集群     |
| 无服务器函数 Serverless Functions | Lambda   | Azure Functions  | Cloud Functions | 事件触发执行      |
| 对象存储 Object Storage           | S3       | Blob Storage     | Cloud Storage   | 文件与归档数据    |
| 块存储 Block Storage              | EBS      | Managed Disks    | Persistent Disk | VM 系统盘/数据盘  |
| 文件存储 File Storage             | EFS      | Azure Files      | Filestore       | 共享目录语义      |
| 托管关系数据库 Managed SQL        | RDS      | Azure SQL        | Cloud SQL       | MySQL/Postgres 等 |
| 托管 NoSQL                        | DynamoDB | Cosmos DB        | Firestore       | 键值/文档模型     |
| 虚拟私有网络 Virtual Network      | VPC      | VNet             | VPC             | 网络隔离边界      |

---

## 10. 课堂逻辑补全：从“会点控制台”到“会做设计” / Filling the Logic Gap: From Console Operation to Design Thinking

课堂演示重在建立直觉和手感，但实际工作还需要把单步操作串成可复用设计。建议用以下四问做每个云资源的设计校验。

The demos build intuition, but production work requires turning isolated actions into repeatable design. Use these four questions as a design check for every cloud resource.

第一问：这个资源的业务目标是什么，是否可量化（例如延迟、可用性、吞吐、RPO/RTO）？

Question 1: What business objective does this resource serve, and is it measurable (latency, availability, throughput, RPO/RTO)?

第二问：它的安全边界在哪里，谁可以访问，凭什么访问，审计日志在哪里看？

Question 2: Where is the security boundary, who can access it, on what identity basis, and where are audit logs?

第三问：它的成本驱动因素是什么，是否具备自动化降本机制？

Question 3: What drives cost for this resource, and is there an automated cost-control mechanism?

第四问：它在故障时如何恢复，恢复流程是否被演练过？

Question 4: How does recovery work during failures, and has the recovery path been practiced?

当你能用这四问复盘 GCP VM 与 S3 两个演示时，说明你已经从“会点按钮”进入“会做基础云设计”的阶段。

If you can apply these four questions to both the GCP VM and S3 demos, you have progressed from "clicking console buttons" to foundational cloud design thinking.

---

## 11. 课堂内容回顾清单 / Classroom Coverage Checklist

以下清单用于确认本笔记已覆盖课堂主线并完成扩展补充。

This checklist confirms the note covers the classroom storyline and includes expanded material.

- 已覆盖：三大云厂商定位与四大服务域。
- Covered: Big-3 positioning and the four core service domains.
- 已覆盖：Compute/Storage/Databases/Networking 的跨云映射。
- Covered: Cross-cloud mapping of compute, storage, databases, and networking.
- 已覆盖：GCP VM 创建、SSH 验证、停止/删除、磁盘残留与计费意识。
- Covered: GCP VM creation, SSH verification, stop/delete, disk leftovers, and billing awareness.
- 已覆盖：S3 Bucket 创建、权限基线、存储类别、生命周期规则、清理流程。
- Covered: S3 bucket creation, permission baseline, storage classes, lifecycle rules, and cleanup flow.
- 已补充：选型策略、成本模型、安全治理、术语翻译、设计四问。
- Added: Selection strategy, cost model, security governance, terminology translation, and four design-check questions.

---

## 12. 高频误区与纠正 / Common Pitfalls and Corrections

误区一：停止实例就不花钱了。

Pitfall 1: "Stopping an instance means no more cost."

纠正：停止通常只停止计算计费，磁盘、快照、静态 IP、日志等可能继续收费。

Correction: Stopping usually halts compute charges only; disks, snapshots, static IPs, and logs may still incur cost.

误区二：对象存储只看每 GB 单价。

Pitfall 2: "Object storage cost is only per-GB price."

纠正：请求费、检索费、最短存储时长、跨区流量都可能成为主要成本来源。

Correction: Request fees, retrieval charges, minimum-duration constraints, and cross-region transfer can dominate total cost.

误区三：有防火墙就等于安全。

Pitfall 3: "If firewall exists, security is done."

纠正：身份权限、密钥管理、加密策略、日志审计同样是安全闭环的一部分。

Correction: Identity access, key management, encryption, and audit logging are equally essential parts of the security loop.

---

## 13. 建议的课后实操路线 / Recommended Post-class Practice Path

建议你用 1-2 小时完成一个最小跨云练习：在 GCP 创建 VM 并验证资源，在 AWS 创建 S3 Bucket 并配置 30 天生命周期规则，然后写出你的清理清单并实际执行。

Use 1-2 hours for a minimum cross-cloud practice: create and verify a VM on GCP, create an S3 bucket on AWS with a 30-day lifecycle rule, then write and execute a cleanup checklist.

如果要进一步提升，可增加一个数据库服务（Cloud SQL 或 RDS）并设计基础网络隔离（私有子网 + 最小开放端口），把课堂四大域完整串起来。

To level up, add one managed database service (Cloud SQL or RDS) and basic network isolation (private subnet plus minimum open ports), connecting all four classroom domains into one design.

---

## 14. 附：原课堂插图索引 / Appendix: Original Classroom Image Index

以下图片均来自原始笔记并已在正文适配位置插入，路径保持不变，便于你后续继续维护。

All images below come from the original notes and were reinserted in context with unchanged paths for easy maintenance.

- image/CoreCloudServices/1774969308099.png
- image/CoreCloudServices/1774969325875.png
- image/CoreCloudServices/1774969337959.png
- image/CoreCloudServices/1774969348356.png
- image/CoreCloudServices/1774969362067.png
- image/CoreCloudServices/1774969374707.png
- image/DemoCreatingS3BucketonAWS/1774969402294.png
- image/DemoCreatingS3BucketonAWS/1774969414895.png
- image/DemoCreatingS3BucketonAWS/1774969464479.png
- image/DemoCreatingS3BucketonAWS/1774969475719.png
- image/DemoCreatingS3BucketonAWS/1774969491293.png
- image/DemoCreatingVMonGoogleCloud/1774969511266.png
- image/DemoCreatingVMonGoogleCloud/1774969526504.png
- image/DemoCreatingVMonGoogleCloud/1774969549895.png

---

## 15. 总结 / Final Summary

如果你只记住一句话，请记住：云平台之间真正可迁移的是抽象能力，而不是产品名字。把计算、存储、数据库、网络作为主坐标，再叠加安全、成本、可观测性和清理习惯，你就具备了持续增长的云基础。

If you remember one sentence, make it this: what transfers across cloud platforms is abstraction capability, not product names. Use compute, storage, databases, and networking as your main coordinates, then layer security, cost control, observability, and cleanup discipline on top.
