# Cloud Security and Costs Part 1

> 云安全基础增强笔记：围绕共享责任模型、身份与网络控制、CIA 三元组与常见攻击面，构建可执行的安全认知框架。
>
> Enhanced cloud security fundamentals: build an actionable security framework around shared responsibility, identity and network controls, the CIA triad, and common attack paths.

## 1. 本节学习目标 / Learning Objectives

这一节的重点不是记住几个安全名词，而是建立一套判断方法：当你看到一个云系统时，能快速识别“谁负责什么、哪里最容易出错、如何最小代价补上防护”。

The goal of this lesson is not to memorize security terms, but to form a decision method: when you see a cloud system, quickly identify who owns what, where failure is most likely, and how to close gaps with the least operational friction.

课堂把“安全”和“成本”放在一起讲，是因为它们都属于横切能力。一个错误权限可能引发数据泄露；一台长期开机的测试机可能造成持续浪费。两者都源于治理不到位。

Security and cost are taught together because both are cross-cutting capabilities. One misconfigured permission can cause data exposure; one always-on test environment can create silent cost waste. Both are governance failures.

![1774763761677](image/CloudSecurityandCostsPart1/1774763761677.png)

## 2. 共享责任模型 / Shared Responsibility Model

云厂商负责“云之下”的安全：物理机房、底层网络设备、硬件生命周期、虚拟化基础设施等。用户负责“云之上”的安全：账户、权限、数据、应用和配置。

Cloud providers secure what is under the cloud: physical datacenters, foundational networking hardware, hardware lifecycle, and virtualization substrate. Customers secure what is in the cloud: identities, permissions, data, applications, and configuration.

一个常见误区是“用了云服务就自动安全”。事实是：厂商给你的是安全能力与控制面，不是替你做完策略。是否开放公网、是否最小权限、是否审计留痕，都由你决定。

A common misunderstanding is that cloud adoption automatically means secure systems. In reality, providers give you security capabilities and control planes, not completed policies. Public exposure, least privilege, and audit traceability are your decisions.

## 3. 两条控制主线：网络与身份 / Two Control Planes: Network and Identity

在实践中，安全控制通常分成两条主线：网络访问控制与身份权限控制。前者回答“谁能连进来”，后者回答“连进来后能做什么”。

In practice, security controls usually run along two planes: network access control and identity permission control. The first answers who can connect, and the second answers what they can do after connecting.

Security Group、Firewall Rule、NACL 这类能力用于限制入口、端口和来源范围；IAM、Role、Policy 这类能力用于限制操作权限和资源范围。

Controls such as security groups, firewall rules, and NACLs restrict entry points, ports, and source ranges; IAM roles and policies restrict actions and resource scope.

如果把系统比作办公室大楼：网络控制像门禁与前台，身份控制像门内权限卡与文件柜钥匙。两者缺一不可。

If you compare systems to an office building, network controls are doors and reception checks, while identity controls are internal access cards and cabinet keys. You need both.

![1774763772457](image/CloudSecurityandCostsPart1/1774763772457.png)

![1774763778009](image/CloudSecurityandCostsPart1/1774763778009.png)

![1774763783026](image/CloudSecurityandCostsPart1/1774763783026.png)

## 4. CIA 三元组：把安全问题结构化 / CIA Triad: Structuring Security Problems

CIA 三元组是云安全分析最实用的结构之一：Confidentiality 关注“不能被不该看的人看到”，Integrity 关注“不能被不该改的人改掉”，Availability 关注“该用的时候必须可用”。

The CIA triad is one of the most practical structures for cloud security analysis: Confidentiality means unauthorized parties must not read data; Integrity means unauthorized parties must not alter data; Availability means systems must remain usable when needed.

把事故映射到 CIA，可以帮助你更快找到治理方向。比如数据外泄属于机密性问题，配置被篡改属于完整性问题，服务不可达属于可用性问题。

Mapping incidents to CIA helps you choose mitigation direction faster. Data leakage is a confidentiality issue, tampered configuration is an integrity issue, and service outages are availability issues.

![1774763787519](image/CloudSecurityandCostsPart1/1774763787519.png)

## 5. 机密性 Confidentiality：防止数据被看见 / Preventing Unauthorized Reading

机密性最常见的破口是错误公开、凭据泄露和社工攻击。对象存储公开读、代码仓库里硬编码密钥、钓鱼拿到账号，都是高频事故入口。

The most common confidentiality breaks are unintended exposure, credential leakage, and social engineering. Public object storage, hard-coded secrets in repositories, and phishing-based account takeover are frequent entry points.

治理策略应包含三层：默认私有与最小权限、秘密管理与自动轮转、身份加强与用户教育。仅做其中一层很难形成长期防线。

Mitigation should include three layers: private-by-default with least privilege, centralized secrets management with rotation, and stronger identity practices with user awareness. A single layer is rarely enough for durable protection.

![1774763794256](image/CloudSecurityandCostsPart1/1774763794256.png)

![1774763800499](image/CloudSecurityandCostsPart1/1774763800499.png)

![1774763824592](image/CloudSecurityandCostsPart1/1774763824592.png)

### 5.1 课堂外补充：秘密管理最小实践 / Extra: Minimum Secrets Management Practices

不要把 API Key 放进源码与镜像层。把密钥放到 Secrets Manager 或同类服务，通过运行时注入；一旦怀疑泄露，立刻轮转并审计调用日志。

Do not place API keys in source code or image layers. Store them in a secrets manager and inject at runtime; if exposure is suspected, rotate immediately and audit usage logs.

对 CI/CD 设置专用短期凭据，避免长期静态凭据在多个系统间复制传播。

Use dedicated short-lived credentials for CI/CD pipelines to prevent long-lived static secrets from propagating across systems.

## 6. 完整性 Integrity：防止数据与配置被篡改 / Preventing Unauthorized Modification

完整性问题不仅来自攻击者，也来自误操作。权限过大、变更无审计、缺少版本回滚，都会让一次小失误放大成生产事故。

Integrity failures come not only from attackers but also from operational mistakes. Over-privileged access, poor change auditability, and missing rollback paths can turn small errors into major production incidents.

建议把权限分层为只读、变更、管理三类角色，并为关键动作启用审计日志与审批机制。

A practical baseline is role segmentation into read-only, change, and admin privileges, with audit logs and approval workflows for critical operations.

![1774763832347](image/CloudSecurityandCostsPart1/1774763832347.png)

## 7. 可用性 Availability：系统持续可用 / Keeping Services Available

可用性风险常见于单点架构、容量规划不足和突发流量。即便数据没泄露、也没被篡改，服务不可用同样会造成业务损失。

Availability risks commonly stem from single points of failure, weak capacity planning, and burst traffic. Even if data is secure and correct, unavailability still causes direct business damage.

可用性治理建议遵循四步：多可用区部署、负载均衡分流、自动扩缩容、健康检查与自动替换。

A practical availability pattern has four steps: multi-zone deployment, load balancing, auto scaling, and health-based replacement.

![1774763842641](image/CloudSecurityandCostsPart1/1774763842641.png)

![1774763848717](image/CloudSecurityandCostsPart1/1774763848717.png)

## 8. 快速判断题复盘 / Quick Quiz Review

题目核心在考察共享责任边界：云厂商不会替你定义业务权限，也不会自动收敛你开放过大的访问策略。正确答案是“用户负责保护自己的数据和代码”。

The quiz checks your understanding of shared responsibility boundaries: cloud providers do not define your business permissions nor automatically tighten your overexposed policies. The correct statement is that users are responsible for securing their own data and code.

这道题背后的方法是：遇到安全争议，先问“谁拥有该配置的决策权”。能改这条规则的人，就是责任主体。

The underlying method is simple: in any security ambiguity, ask who owns the configuration decision. The party that can change the rule owns the responsibility.

![1774763856024](image/CloudSecurityandCostsPart1/1774763856024.png)

## 9. 小结与迁移到下一节 / Summary and Transition to Cost Management

本节建立了云安全的三个核心框架：共享责任模型、网络与身份双控制面、CIA 三元组。它们会在后续所有平台实操中反复出现。

This lesson established three core frameworks for cloud security: shared responsibility, the dual control planes of network and identity, and the CIA triad. These frameworks recur in every platform-specific practice.

下一节进入成本治理时，你会发现同样的方法仍然适用：先划分责任，再建立可观测与可执行规则，最后通过自动化降低人为偏差。

When moving to cost governance in the next section, the same method still applies: define ownership first, build observable and executable rules, and use automation to reduce human variance.

![1774763865792](image/CloudSecurityandCostsPart1/1774763865792.png)

## 10. 复习清单 / Review Checklist

- 你能否解释共享责任模型中“云厂商负责什么、你负责什么”？
- 你能否区分网络控制与 IAM 控制分别解决什么问题？
- 你能否把一个真实事故归类到 CIA 三元组之一并给出缓解方案？
- 你能否给出 Secrets 管理、最小权限、MFA、审计日志四者之间的协作关系？

- Can you explain what providers secure versus what you secure in the shared responsibility model?
- Can you distinguish what network controls solve versus what IAM controls solve?
- Can you map a real incident to one CIA pillar and propose mitigation?
- Can you describe how secrets management, least privilege, MFA, and audit logs work together?
