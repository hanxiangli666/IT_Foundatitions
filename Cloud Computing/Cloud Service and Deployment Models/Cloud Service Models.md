# Cloud Service Models

> 云服务模型增强笔记：围绕 IaaS、PaaS、SaaS 的责任边界、适用场景与选型方法，建立可迁移的中英双语理解框架。
>
> Enhanced cloud service model notes: build transferable understanding of IaaS, PaaS, and SaaS through responsibility boundaries, fit-for-purpose scenarios, and selection logic in bilingual format.

## 1. 本节学习目标 / Learning Objectives

本节的核心不是背定义，而是看懂“谁管理哪一层”。当你能清晰回答“硬件、系统、平台、应用、数据分别由谁负责”时，服务模型就不再抽象。

The core objective is not memorizing definitions, but understanding who manages each layer. Once you can answer who owns hardware, OS, platform, application, and data, service models stop being abstract.

课堂基于 NIST 云计算定义，强调按需访问、资源池化、快速弹性和可度量使用。你应把这些特征与模型选择联系起来，而不是孤立记忆。

The lesson is grounded in the NIST view of cloud computing, highlighting on-demand access, resource pooling, rapid elasticity, and measured usage. You should connect these characteristics to model choice instead of memorizing them in isolation.

![1774763383164](image/CloudServiceModels/1774763383164.png)

## 2. 三层责任视角 / Three-layer Responsibility View

为了便于学习，可把技术栈简化为三层：基础设施层（算力、网络、存储）、平台层（操作系统、运行时、中间件）、应用层（业务代码与数据）。

For learning clarity, the stack can be simplified into three layers: infrastructure (compute, network, storage), platform (OS, runtime, middleware), and application (business code and data).

云服务模型的本质，就是这三层责任在“你”和“云厂商”之间如何分配。

The essence of cloud service models is how responsibility for these layers is split between you and the cloud provider.

## 3. IaaS：高控制，高运维责任 / IaaS: High Control, Higher Ops Ownership

IaaS 中，云厂商提供服务器、存储、网络和虚拟化能力。你负责操作系统、运行环境、应用部署和数据治理。

In IaaS, the provider supplies servers, storage, networking, and virtualization. You own the operating system, runtime environment, application deployment, and data governance.

IaaS 的优势是灵活和可控，适合需要自定义系统行为或特殊网络拓扑的业务。但代价是你要承担补丁、加固、容量规划、故障处理等运维工作。

IaaS offers flexibility and control, fitting workloads that require custom system behavior or specialized network topology. The tradeoff is greater operational burden: patching, hardening, capacity planning, and incident handling.

![1774763390239](image/CloudServiceModels/1774763390239.png)

## 4. PaaS：减少平台运维，聚焦业务代码 / PaaS: Less Platform Operations, More Product Focus

PaaS 在 IaaS 基础上进一步托管了 OS、运行时和平台能力。开发团队主要关注代码与数据逻辑，不需要频繁管理底层系统。

PaaS extends IaaS by managing OS, runtime, and platform operations. Engineering teams focus mostly on code and data logic rather than routine system maintenance.

它特别适合“功能迭代速度优先”的团队，例如快速上线新模块、验证新功能、缩短发布周期。

It is especially suitable for teams prioritizing delivery speed, such as rapid feature rollout, experimentation, and shorter release cycles.

## 5. SaaS：直接使用成品应用 / SaaS: Consume Ready-made Applications

SaaS 由厂商托管完整应用与底层栈。你通常只做租户级配置、用户管理和流程集成。

SaaS providers manage the full application and underlying stack. You typically handle tenant-level configuration, user administration, and workflow integration.

这类模型非常适合协作办公、沟通、文档、工单等通用场景，能以最小技术投入获得可用能力。

This model is ideal for common needs like collaboration, communication, documentation, and workflow tooling, delivering value with minimal engineering overhead.

![1774763398204](image/CloudServiceModels/1774763398204.png)

## 6. 生活类比：租房模型理解云模型 / Analogy: Housing Models for Cloud Models

你可以把 IaaS 理解为“毛坯房”：房屋结构已给你，但装修和维护大部分由你承担；PaaS 像“精装公寓”：底层设施已准备好，你带业务入住；SaaS 像“酒店”：拎包即用。

You can think of IaaS as an unfurnished apartment: structure is provided, but most interior setup and upkeep are yours. PaaS is like a furnished apartment: key facilities are ready and you bring your business logic. SaaS is like a hotel: immediate consumption with minimal setup.

这个类比的关键不是“形象”，而是提醒你：控制权与管理负担始终成反比。

The important point of the analogy is not visual appeal, but the tradeoff: greater control usually means greater management burden.

![1774763406502](image/CloudServiceModels/1774763406502.png)

![1774763412910](image/CloudServiceModels/1774763412910.png)

## 7. MiaoTube 案例：需求到模型的映射 / MiaoTube Case: Mapping Needs to Models

案例中，MiaoTube 有三类需求：运行核心视频处理流水线、快速迭代新功能、使用稳定后台协作工具。单一模型难以同时最优，组合模型更现实。

In the case study, MiaoTube has three needs: run core video pipelines, deliver new features quickly, and use reliable admin collaboration tools. A single model rarely optimizes all three; mixed-model strategy is more realistic.

视频流水线需要资源控制和性能调优，适合 IaaS；新功能迭代强调速度，适合 PaaS；办公协作偏通用能力，适合 SaaS。

Video pipelines need control and performance tuning, fitting IaaS; feature iteration prioritizes speed, fitting PaaS; office collaboration uses standardized capabilities, fitting SaaS.

![1774763419564](image/CloudServiceModels/1774763419564.png)

![1774763426012](image/CloudServiceModels/1774763426012.png)

![1774763431786](image/CloudServiceModels/1774763431786.png)

## 8. 责任分工对照表（核心复习） / Responsibility Matrix (Core Review)

| Layer / Model | On-Prem | IaaS | PaaS | SaaS |
| --- | --- | --- | --- | --- |
| Physical infrastructure | You | Provider | Provider | Provider |
| Virtualization | You | Provider | Provider | Provider |
| OS and runtime | You | You | Provider | Provider |
| Application and data | You | You | You | Mostly provider-managed app |

这张表建议背后的逻辑去理解：并非“谁全托管就一定更好”，而是“是否适合当前团队目标、能力和风险承受度”。

Understand the logic behind this table: it is not that more managed is always better, but whether the model matches team goals, capability, and risk tolerance.

## 9. 快速判断题复盘 / Quick Quiz Review

正确结论是：SaaS 由供应商托管完整后端，你可以直接登录使用。常见误区是把 IaaS 当 SaaS，把 PaaS 当 IaaS。

The correct takeaway is that SaaS provides a fully managed backend and is consumed directly by users. A common misconception is confusing IaaS with SaaS and PaaS with IaaS.

判断题的实用技巧：看“谁负责操作系统”，如果仍由你管理，通常不是 PaaS 或 SaaS。

A practical exam tip: check who manages the OS. If you still manage it, the model is usually not PaaS or SaaS.

## 10. 课堂外补充：选型四问法 / Extra: Four Questions for Model Selection

第一问：团队是否有足够运维能力支撑 OS 与平台层？

Question 1: Does the team have enough operational capability to manage OS and platform layers?

第二问：当前目标是“快速上线”还是“深度控制”？

Question 2: Is the current goal rapid delivery or deep control?

第三问：合规要求对数据驻留与审计有多强？

Question 3: How strict are compliance requirements around data residency and auditability?

第四问：未来 6-12 个月需求是否稳定，是否要为变化预留弹性？

Question 4: Are requirements stable in the next 6-12 months, and how much flexibility is needed?

当这四问有答案时，模型选择通常会清晰很多。

When these four questions are answered, model selection usually becomes far clearer.

## 11. 总结 / Final Summary

云服务模型描述的是“你租了多少、你管了多少”。IaaS 给你最大控制，PaaS 给你开发速度，SaaS 给你即时可用性。实际企业往往组合使用，而不是单选。

Cloud service models describe how much you rent and how much you operate. IaaS offers maximum control, PaaS emphasizes delivery speed, and SaaS offers immediate usability. Real organizations usually combine models rather than choosing only one.

![1774763441077](image/CloudServiceModels/1774763441077.png)
