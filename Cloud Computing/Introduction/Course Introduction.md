# Course Introduction / 课程导览

> Overview of cloud computing fundamentals covering service models, deployment options, scalability, storage, security, cost management, and hands-on labs with major providers.
>
> 本章概览云计算的核心基础，包括服务模型、部署方式、可扩展性、存储、安全、成本管理，以及在主流云平台上的动手实验。

Picture this: you snap a photo on your phone and it’s instantly backed up and accessible anywhere in the world. Or a business launches a website that serves millions of users overnight without buying a single new server. How is this possible? Cloud computing is the technology powering much of our connected world.

想象一下：你在手机上拍了一张照片，它立刻就被备份，并且能在世界任何地方访问。又或者，一家企业一夜之间就能让网站服务数百万用户，而完全不用先购买新的服务器。这是怎么做到的？答案就是云计算，它支撑着我们今天大量的互联服务。

Hi, I’m Alan. In this lesson we’ll unpack the core ideas behind the cloud, explore how it changes IT operations, and show you how to make the most of its flexibility and power.

大家好，我是 Alan。在这一课里，我们会拆解云背后的核心概念，看看它如何改变 IT 运营方式，并说明如何充分利用它的灵活性和能力。

## What you’ll learn / 你将学到什么

* Core differences between cloud and traditional IT: renting compute and storage instead of owning and maintaining hardware. / 云与传统 IT 的核心区别：租用算力和存储，而不是自己购买并维护硬件。
* Cloud service models (IaaS, PaaS, SaaS) and how each one shifts control and responsibility. / 云服务模型（IaaS、PaaS、SaaS）以及它们如何重新分配控制权和责任。
* Deployment approaches (public, private, hybrid, multi-cloud) and where each one fits best. / 部署方式（公有云、私有云、混合云、多云）以及各自最适合的场景。
* How cloud providers run code at scale, store and distribute data globally, and support different database types. / 云厂商如何在大规模环境下运行代码、全球存储与分发数据，并支持不同类型的数据库。
* Security responsibilities, common threats, and practical cost-management techniques. / 安全责任划分、常见威胁，以及实用的成本管理方法。
* Hands-on experience creating, monitoring, and cleaning up cloud resources on major providers. / 在主流云平台上动手创建、监控并清理云资源的实践经验。

  ![1774763222156](image/CourseIntroduction/1774763222156.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/kodekloud-presenter-single-vs-multiple-servers.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=2e9347cfb51068fe517e4af0bac54572" alt="A presenter in a KodeKloud t-shirt stands to the right of a graphic showing cloud icons and server illustrations. The slide compares a single server versus multiple servers with the text "Why own one?" and "VS."" width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/kodekloud-presenter-single-vs-multiple-servers.jpg" />
</Frame>

## Why cloud is different / 云计算为何不同

Cloud changes the operating model. Instead of committing to servers up front, you can request capacity when you need it, release it when you do not, and let the provider handle much of the underlying complexity.

云计算改变了运行模式。你不需要一开始就投入服务器采购，而是可以在需要时申请资源，在不需要时释放资源，并把很多底层复杂性交给云厂商处理。

* Rent vs own: Cloud lets you pay for capacity when you need it rather than buying and maintaining physical servers. / 租用而非自有：云让你按需为容量付费，而不是自己购买和维护物理服务器。
* Elasticity: Scale up or down quickly to match demand. / 弹性：可以快速扩容或缩容，以匹配真实需求。
* Operational trade-offs: You trade some direct control for faster provisioning, managed services, and operational simplicity. / 运营权衡：你会牺牲一部分直接控制权，换来更快的开通、托管服务和更简单的运维。
* Business impact: Faster iteration, reduced time-to-market, and often better cost-efficiency when workloads are variable. / 商业影响：迭代更快、上市时间更短，而且在负载波动较大的情况下通常更具成本效率。

## Cloud service models at a glance / 云服务模型速览

| Service Model / 服务模型                            | What you manage / 你负责什么                        | What the provider manages / 云厂商负责什么                           | Typical use case / 典型场景                                                      |
| --------------------------------------------------- | --------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| IaaS (Infrastructure as a Service) / 基础设施即服务 | OS, runtimes, applications / 操作系统、运行时、应用 | Servers, virtualization, networking / 服务器、虚拟化、网络           | Lift-and-shift VMs, custom platforms / 迁移现有虚拟机、自定义平台                |
| PaaS (Platform as a Service) / 平台即服务           | Applications, data / 应用、数据                     | OS, runtimes, middleware, scaling / 操作系统、运行时、中间件、扩缩容 | Web apps, microservices, developer productivity / Web 应用、微服务、提升开发效率 |
| SaaS (Software as a Service) / 软件即服务           | User data, configuration / 用户数据、配置           | Entire stack and app / 整个技术栈和应用本身                          | Email, CRM, collaboration tools / 邮件、CRM、协作工具                            |

IaaS gives you the most control, PaaS removes much of the platform maintenance burden, and SaaS removes almost all infrastructure responsibility from your team.

IaaS 提供最多控制权，PaaS 会去掉大量平台维护工作，而 SaaS 则把几乎所有基础设施责任都从你的团队中移走。

## Deployment approaches and when to choose them / 部署方式与选择时机

| Deployment Model / 部署模型 | When to use it / 何时使用                                                                     | Example / 示例                                            |
| --------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| Public cloud / 公有云       | Rapid scaling, pay-as-you-go, no datacenter management / 需要快速扩展、按量付费、不想管理机房 | Startups, bursty workloads / 初创公司、突发型负载         |
| Private cloud / 私有云      | Strict compliance, full control / 合规要求严格、需要完全控制                                  | Regulated industries / 受监管行业                         |
| Hybrid cloud / 混合云       | Mix of on-prem and cloud for flexibility / 本地与云结合以获得灵活性                           | Gradual cloud migration / 渐进式上云                      |
| Multi-cloud / 多云          | Avoid vendor lock-in or leverage best-of-breed services / 避免厂商锁定，或组合各家最优服务    | Large enterprises with diverse needs / 需求复杂的大型企业 |

The main idea is simple: service models describe what you rent, while deployment models describe where and how that rented environment is hosted.

核心思路很简单：服务模型描述的是“你租什么”，而部署模型描述的是“这些资源部署在哪里、如何托管”。

Next, we’ll break down the main cloud service models and show how each one shifts the balance between user control and provider convenience, with practical examples.

接下来，我们会进一步拆解主要的云服务模型，并用实际例子说明它们如何在用户控制权与云厂商便利性之间重新平衡。

You’ll also learn the key approaches to deploying cloud environments and how each fits different needs, with quick examples of where and why organizations choose them.

你还会学习云环境的关键部署方式，以及它们分别适合什么需求，并快速了解组织为什么会选择这些模式。

Then we’ll peek behind the scenes. You’ll see how the cloud runs code at scale, stores massive amounts of data, distributes it globally with low latency, and supports a range of database types.

然后我们会看一看幕后。你将了解云是如何大规模运行代码、存储海量数据、以低延迟在全球分发内容，并支持多种数据库类型的。

![1774763229808](image/CourseIntroduction/1774763229808.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/kodekloud-instructor-purple-tech-background.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=3a659ec927f7f5998c8b0513535263a9" alt="A man in a black KodeKloud T-shirt is speaking in front of a purple, tech-themed background. The backdrop shows stylized cloud icons, code file symbols (HTML/CSS) and cartoon avatars labeled with countries like Argentina, Canada, Korea, Thailand, Philippines and Peru." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/kodekloud-instructor-purple-tech-background.jpg" />
</Frame>

## How cloud services operate at scale / 云服务如何在大规模下运行

* Compute: Auto-scaling groups, serverless functions, and container orchestration run ephemeral workloads and long-lived services. / 计算：自动扩缩组、无服务器函数和容器编排既能运行短暂任务，也能支撑长期在线服务。
* Storage: Object storage for large, durable files; block storage for VM disks; file systems for shared access. / 存储：对象存储适合大文件和高耐久数据；块存储适合虚拟机磁盘；文件系统适合共享访问。
* Networking & delivery: CDNs and global load balancers minimize latency for distributed users. / 网络与分发：CDN 和全球负载均衡器能够为分布式用户降低延迟。
* Data & databases: Relational databases, NoSQL stores, data warehouses, and analytics pipelines support different workloads and SLAs. / 数据与数据库：关系型数据库、NoSQL、数据仓库和分析流水线分别支撑不同负载和 SLA 需求。

The important pattern is that cloud systems are designed to separate responsibilities, scale horizontally when needed, and place services closer to users.

关键模式在于：云系统会把职责拆开、在需要时横向扩展，并尽量把服务部署到离用户更近的位置。

## Security and cost control matter / 安全与成本控制同样重要

* Shared responsibility: Cloud providers secure the infrastructure; customers secure their data, identities, and configurations. / 共享责任：云厂商负责基础设施安全，客户负责数据、身份和配置安全。
* Common threats: Misconfigured storage, exposed credentials, insecure network rules. / 常见威胁：存储配置错误、凭证泄露、不安全的网络规则。
* Cost techniques: Right-sizing, reserved instances or savings plans, scheduling non-production shutdowns, and monitoring spend with alerts and budgets. / 成本技巧：合理选型、预留实例或节省计划、安排非生产环境停机时间、使用告警和预算监控支出。

  ![1774763237636](image/CourseIntroduction/1774763237636.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/spotting-threats-cloud-key-theft.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=5924363805cf874315bd61fb2919d8a9" alt="A presenter stands on the right wearing a KodeKloud shirt, speaking against a dark purple background. On the left is an illustration of a cloud with a shadowy figure reaching for a key and a person handing over cash, with the title "Spotting Threats."" width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/spotting-threats-cloud-key-theft.jpg" />
</Frame>

<Callout icon="warning" color="#FF6B6B">
  Cloud resources can incur real costs if left running. Always practice creating and cleaning up resources during labs, and use provider cost controls such as budgets, alerts, and shutdown schedules to avoid surprises.
  <br />
  云资源如果一直开着，就会产生真实费用。做实验时一定要练习创建资源后及时清理资源，并使用预算、告警和关机计划等成本控制手段，避免账单意外增加。
</Callout>

## Major cloud providers and hands-on practice / 主流云厂商与动手实践

You’ll meet the major cloud providers — [AWS](https://learn.kodekloud.com/user/courses/aws-for-beginners-with-hands-on-labs), [Azure](https://learn.kodekloud.com/user/courses/az900-microsoft-azure-fundamentals), and [Google Cloud](https://learn.kodekloud.com/user/courses/gcp-cloud-digital-leader-certification). This course includes guided labs where you’ll create, inspect, and clean up cloud resources while learning to use provider dashboards for monitoring security posture and costs.

你会接触到主流云厂商——[AWS](https://learn.kodekloud.com/user/courses/aws-for-beginners-with-hands-on-labs)、[Azure](https://learn.kodekloud.com/user/courses/az900-microsoft-azure-fundamentals) 和 [Google Cloud](https://learn.kodekloud.com/user/courses/gcp-cloud-digital-leader-certification)。本课程还提供引导式实验，你会在实验中创建、检查并清理云资源，同时学习如何使用云厂商控制台来监控安全状况和成本。

![1774763245051](image/CourseIntroduction/1774763245051.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/cloud-logos-s3-security-kodekloud.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=7680dc20d2226f0beac752c7838d5f01" alt="The image shows cloud provider logos (AWS, Azure, Google Cloud) above a laptop mockup displaying an AWS S3 settings page and a "Security" badge. A man wearing a KodeKloud t-shirt stands to the right as if presenting." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Introduction/Course-Introduction/cloud-logos-s3-security-kodekloud.jpg" />
</Frame>

<Callout icon="lightbulb" color="#1CB2FE">
  To get the most from the labs: create a free-tier account if available, follow the cleanup instructions after each lab, and ask questions in the KodeKloud community. Hands-on practice is the fastest path to confidence.
  <br />
  想把实验做得更有效：如果可用就创建免费层账号；每个实验结束后都按照说明清理资源；有问题就去 KodeKloud 社区提问。动手实践是建立信心最快的方式。
</Callout>

## Who this lesson is for / 这节课适合谁

Whether you’re aiming to break into tech, modernize an existing business, or simply satisfy your curiosity, this lesson gives you the foundational knowledge and practical skills to start working with cloud technology confidently.

无论你是想进入技术行业、为现有业务做现代化改造，还是单纯出于兴趣，这一节课都会给你打下基础知识和实践技能，让你更有信心地开始使用云技术。

Bring your questions, try the labs, and join the global learning community at KodeKloud.

带着你的问题来学习，去做实验，并加入 KodeKloud 的全球学习社区。

<CardGroup>
  <Card title="Watch Video" icon="video" cta="Learn more" href="https://learn.kodekloud.com/user/courses/cloud-computing-fundamentals/module/011c3039-0ef4-42f9-8ac1-0633bb3bf667/lesson/32a30d02-13db-4f86-ba8d-6142328a3a1a" />
</CardGroup>

Built with [Mintlify](https://mintlify.com).
