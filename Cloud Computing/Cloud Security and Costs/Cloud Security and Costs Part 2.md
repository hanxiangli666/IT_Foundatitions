# Cloud Security and Costs Part 2

> 云成本治理增强笔记：从 CapEx 到 OpEx 的思维迁移、服务与部署模型对成本的影响、FinOps 方法、以及可执行降本动作。
>
> Enhanced cloud cost governance note: mindset shift from CapEx to OpEx, cost implications of service and deployment models, FinOps practice, and executable optimization actions.

## 1. 本节学习目标 / Learning Objectives

本节的目标不是只会看账单，而是建立“成本可预测、责任可归属、动作可自动化”的治理体系。

The objective is not just reading cloud bills, but building governance where cost is predictable, ownership is clear, and optimization actions are automatable.

云的弹性和便捷会放大创新速度，也会放大浪费速度。没有规则时，资源增长通常比业务价值增长更快。

Cloud elasticity and convenience accelerate innovation, but they can also accelerate waste. Without guardrails, resource growth often outpaces business value growth.

![1774968910294](image/CloudSecurityandCostsPart2/1774968910294.png)

## 2. CapEx 到 OpEx 的迁移 / Shift from CapEx to OpEx

传统机房时代，企业常在前期一次性投入硬件和机房资源，属于资本开支。云时代更多是按用量付费，属于运营开支，这带来更低门槛和更高波动。

In traditional datacenter models, organizations made large upfront investments in hardware and facilities, which are capital expenditures. Cloud moves many costs to usage-based operational spending, lowering entry barriers while increasing variability.

这意味着成本管理从“采购决策”转向“持续运营决策”：每一次扩容、每一个未清理资源、每一条错误规格选择，都会体现在月账单上。

This means cost management shifts from one-time procurement decisions to continuous operational decisions: every scale-out, every uncleared resource, and every wrong sizing choice appears on monthly bills.

![1774968929687](image/CloudSecurityandCostsPart2/1774968929687.png)

## 3. 服务模型如何影响成本结构 / How Service Models Shape Cost Structure

IaaS 给你更高控制权，也给你更多犯错空间。你可以精细优化，但如果长期过配，浪费会持续累积。

IaaS gives you greater control and also more room for mistakes. Fine-grained optimization is possible, but persistent overprovisioning accumulates waste.

PaaS 与 SaaS 把更多运维责任交给平台，通常能提升交付效率，但单价不一定最低。是否划算取决于团队的人力成本和稳定性要求。

PaaS and SaaS shift more operations to the provider, often increasing delivery speed, but not always minimizing unit cost. True efficiency depends on team labor cost and reliability requirements.

![1774969000097](image/CloudSecurityandCostsPart2/1774969000097.png)

## 4. 部署模型与隐藏成本 / Deployment Models and Hidden Costs

公有云适合波动业务与快速试错；私有云在合规和可控性上有优势但维护成本高；混合云兼顾迁移与特殊场景，但也引入跨平台协作和治理复杂度。

Public cloud fits variable workloads and rapid experimentation; private cloud can support strict compliance and control at higher operational cost; hybrid cloud helps migration and special-case workloads but introduces cross-platform governance complexity.

很多团队低估了“复杂度成本”：多套网络、监控、身份体系并存时，排障与协同成本会显著上升。

Many teams underestimate complexity cost: when multiple network, monitoring, and identity systems coexist, troubleshooting and coordination overhead can rise sharply.

![1774969052783](image/CloudSecurityandCostsPart2/1774969052783.png)

## 5. FinOps 不是工具，而是协作机制 / FinOps Is a Practice, Not a Tool

FinOps 的核心是让技术、财务、业务共享同一套成本视图和决策节奏。目标不是单纯“降本”，而是在速度、质量、成本之间找到可持续平衡。

FinOps aligns engineering, finance, and business on shared cost visibility and decision cadence. The goal is not cost reduction alone, but sustainable balance among speed, quality, and spend.

有效的 FinOps 需要三个基础动作：资源可归属（标签与账单分摊）、异常可感知（预算与告警）、优化可执行（流程和自动化）。

Effective FinOps needs three foundations: attributable resources (tagging and allocation), detectable anomalies (budgets and alerts), and executable optimization (process and automation).

![1774969064006](image/CloudSecurityandCostsPart2/1774969064006.png)

## 6. 常见成本陷阱与对应策略 / Common Cost Traps and Countermeasures

过度配置是最常见的浪费来源。很多团队担心未来增长而一次性选最大规格，但实际利用率长期偏低。

Overprovisioning is one of the most common waste patterns. Teams often choose maximum sizes for future safety while long-term utilization remains low.

非生产环境常年运行也是典型浪费。开发、测试、演示环境如果 24x7 开机，月度累计费用会非常可观。

Always-on non-production environments are another major cost leak. Dev, test, and demo environments running 24x7 can accumulate substantial monthly spend.

缺少标签会让账单不可解释。不能解释就无法归责，无法归责就难以持续优化。

Missing tagging makes bills hard to explain. If cost cannot be explained, ownership cannot be assigned; without ownership, optimization rarely sustains.

## 7. 可执行的降本动作 / Executable Cost Optimization Actions

### 7.1 Right-Sizing 规格匹配

先小后大，按监控数据迭代调整。重点观察 CPU、内存、磁盘与网络指标的长期趋势，而不是某个峰值瞬间。

Start small and scale by metrics. Focus on long-term trends in CPU, memory, disk, and network rather than isolated spikes.

### 7.2 自动关停非生产环境 / Automated Non-production Shutdowns

为非生产资源设定工作时段策略，非工作时段自动关停；同时保留一键启动脚本，减少开发体验损失。

Apply work-hour schedules for non-production resources and shut them down automatically off-hours, while keeping one-click startup workflows to preserve developer productivity.

### 7.3 长期承诺折扣 / Commitment Discounts

对于稳定可预测负载，使用 Reserved Instances、Savings Plans、Committed Use Discounts 可显著降低单价；但覆盖比例应基于真实历史利用率。

For steady predictable workloads, Reserved Instances, Savings Plans, or Committed Use Discounts can reduce unit cost substantially; coverage ratios should be based on real historical utilization.

![1774969084566](image/CloudSecurityandCostsPart2/1774969084566.png)

### 7.4 标签治理 / Tag Governance

建议至少强制以下标签：team、environment、project、cost-center、owner。并在资源创建流程中做缺失校验，避免后补标签。

At minimum, enforce tags such as team, environment, project, cost-center, and owner. Validate these at creation time to avoid unreliable retroactive tagging.

## 8. 课堂判断题复盘 / Quiz Review

题目中正确项是“承诺计划可为可预测长期使用提供折扣”。其余选项要么混淆 CapEx/OpEx，要么误解 Right-Sizing 和 FinOps 概念。

The correct statement is that commitment plans offer discounts for predictable long-term usage. Other options confuse CapEx versus OpEx or misinterpret right-sizing and FinOps.

判断这类题时，建议抓三点：概念定义是否准确、责任边界是否清晰、是否把方法误当成产品。

For similar questions, validate three things: conceptual definition accuracy, clear ownership boundaries, and whether a practice is mistakenly treated as a product.

![1774969099282](image/CloudSecurityandCostsPart2/1774969099282.png)

## 9. 课堂外补充：把降本做成常态机制 / Extra: Making Optimization a Routine Mechanism

建议建立每周成本例会：看预算偏差、看异常增长、看本周可执行动作与责任人。

Run a weekly cost review cadence: compare budget variance, investigate anomalies, and assign concrete optimization actions with clear owners.

建议建立三层告警：预算阈值告警、服务级异常告警、资源级长时间空闲告警。三层组合能兼顾宏观和细节。

Implement three alert layers: budget threshold alerts, service-level anomaly alerts, and resource-level idle-duration alerts. This combination balances macro visibility and micro actionability.

建议把成本指标并入工程目标，例如把单位请求成本、单位租户成本、资源利用率纳入团队看板，避免“只追性能不看费用”。

Integrate cost metrics into engineering scorecards, such as cost per request, cost per tenant, and utilization efficiency, so teams avoid optimizing performance while ignoring spend.

## 10. 本节总结 / Final Summary

本节核心是三句话：云成本是持续运营问题；FinOps 是协作与治理机制；降本要靠可执行动作和自动化闭环。

The lesson can be summarized in three statements: cloud cost is a continuous operations problem, FinOps is a collaborative governance mechanism, and optimization depends on executable actions plus automation loops.

当你把 right-sizing、自动关停、承诺折扣、标签治理与预算告警组合起来，成本将从“月底惊吓”变成“可预测曲线”。

When right-sizing, scheduled shutdowns, commitment discounts, tag governance, and budget alerts are combined, cloud cost shifts from end-of-month surprises to predictable curves.

## 11. 复习清单 / Review Checklist

- 你能否解释 CapEx 与 OpEx 在决策节奏上的区别？
- 你能否说明 IaaS 与 PaaS/SaaS 在成本控制自由度上的差异？
- 你能否给出一个团队级 FinOps 周期动作示例？
- 你能否列出至少三条可立即执行的降本动作？

- Can you explain the decision-cadence difference between CapEx and OpEx?
- Can you describe how cost-control flexibility differs between IaaS and PaaS/SaaS?
- Can you provide one example of a team-level FinOps weekly cycle?
- Can you list at least three immediately executable optimization actions?
