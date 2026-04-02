# 数据库基础（Database Fundamentals）

> 本笔记基于课程录音整理，涵盖数据库概念、关系型数据库设计、SQL CRUD 操作、非关系型数据库（NoSQL）、查询优化与索引、以及完整的数据库设计实战。录音中的语音识别错误已修正，无关对话内容已过滤。

---

## 一、什么是数据库

### 1.1 从散乱数据到结构化信息

课程用一个生活场景开场：你在计划一次旅行——餐厅名字存在备忘录里，火车时刻表在 WhatsApp 消息里，酒店信息在邮件里，Airbnb 预订截图不知道在哪个相册里。旅行当天你疯狂翻找、滑动、搜索……

这就是 **散乱数据** 的问题：信息你都有，但它们没有被组织、不可搜索、彼此不关联。

**数据库不只是存储数据——它把数据连接起来。**

### 1.2 数据（Data）vs 信息（Information）

课程用 Cody 的兽医记录来说明这个区别。

**未结构化的原始数据：**
```
17 July checkup Vet Atkins $35
2307 25 vaccination 45 Vet Clark
30 July Vet Brown microchip USD 100
```

这是"数据"——原始的文字、数字和日期混在一起，没有标题、格式不统一、排列无规律。随着记录增多，想找到"下次疫苗接种是什么时候"变得越来越难。

**结构化后的信息：**

| Date | Reason | Vet Name | Cost |
|------|--------|----------|------|
| 2025-07-17 | Checkup | Atkins | $35 |
| 2025-07-23 | Vaccination | Clark | $45 |
| 2025-07-30 | Microchip | Brown | $100 |

清理步骤：
1. 给每一列加上 **字段名（Field）** / 列标题：Date, Reason, Vet Name, Cost
2. 确保每个字段的 **数据类型（Data Type）** 一致：Date 列存日期，Cost 列存数字/货币，其余存文本

现在 Cody 可以按日期排序、按原因筛选、统计总花费。**这就是数据变成信息的过程：结构化、可搜索、有意义。**

### 1.3 核心术语

| 术语 | 含义 |
|------|------|
| **数据库（Database）** | 存储和组织大量相关信息的结构化容器 |
| **字段（Field）** | 表中的一列——如 Name、Date、Cost |
| **记录（Record）** | 表中的一行——一次完整的兽医就诊记录 |
| **数据类型（Data Type）** | 字段应该存储的数据种类 |

常见数据类型：

| 数据类型 | 说明 | 示例 |
|---------|------|------|
| **INT / INTEGER** | 整数 | 1, 42, 1000 |
| **VARCHAR(n)** | 可变长度字符串，最长 n 个字符 | "Fluffy", "alan@email.com" |
| **TEXT** | 长文本（不限长度或很大） | 地址、文章内容 |
| **DATE** | 日期 | 2025-07-17 |
| **DATETIME** | 日期+时间 | 2025-07-17 14:30:00 |
| **DECIMAL(m,d)** | 精确小数，m 位总长，d 位小数 | 99.99 (DECIMAL(4,2)) |
| **BOOLEAN** | 布尔值 | TRUE / FALSE |

### 1.4 数据库无处不在

你每天使用的应用背后都有数据库：

| 应用 | 数据库在做什么 |
|------|--------------|
| **手机通讯录** | 每个联系人是一条记录，Name、Phone、Email 是字段 |
| **音乐播放列表** | Song Title、Artist、Album——可搜索、可排序 |
| **Instagram** | 超过 10 亿用户的帖子、点赞、关注、私信关系 |
| **Amazon** | 数百万商品的详情、价格、库存、评价、订单——实时更新 |

没有数据库，这一切都是不可能的。

---

## 二、关系型数据库设计

### 2.1 平面文件（Flat File）的问题

Cody 和朋友创建了一个猫咪视频平台叫 MeowTube，用电子表格跟踪上传：

| Title | Username | Email | Link | Upload Date |
|-------|----------|-------|------|-------------|
| Cat Skateboard | Fluffy | fluffy@email.com | /watch/1 | 2025-05-25 |
| Epic Cat Jump | Paws | paws@email.com | /watch/2 | 2025-05-27 |
| Cat vs Curtain | Fluffy | flufy@email.com | /watch/3 | 2025-05-29 |

一开始没问题，但随着数据增长，三个致命问题出现了：

**问题一：数据冗余（Data Redundancy）**

Fluffy 的用户名和邮箱在第 1 行和第 3 行重复出现。同一条信息被存了多次。

**问题二：数据不一致（Data Inconsistency）**

注意第 3 行 Fluffy 的邮箱拼错了（`flufy@email.com`）。如果 Fluffy 修改了邮箱，他需要逐行更新每个出现的地方。漏掉一个就会导致同一个人有两个不同的邮箱——到底哪个是对的？

数据量小时也许还好，但想象一下百万行记录中修改一个邮箱地址……

**问题三：安全风险**

所有数据在一个文件里，任何人都能看到所有人的私人联系信息。

这种把所有数据塞在一张大表里的存储方式叫 **平面文件（Flat File）**。它简单、搭建快、适合小数据集，但一旦数据增长，就无法保证一致性、安全性和可维护性。

### 2.2 关系型数据库——拆表 + 连接

解决方案：把一张大表拆成多张小表，每张表只存一类信息，然后用"键"把它们连接起来。

**用户表（Users）：**

| user_id | username | email |
|---------|----------|-------|
| 1 | Fluffy | fluffy@email.com |
| 2 | Paws | paws@email.com |

**视频表（Videos）：**

| video_id | user_id | title | link | upload_date |
|----------|---------|-------|------|-------------|
| 1 | 1 | Cat Skateboard | /watch/1 | 2025-05-25 |
| 2 | 2 | Epic Cat Jump | /watch/2 | 2025-05-27 |
| 3 | 1 | Cat vs Curtain | /watch/3 | 2025-05-29 |

现在：
- Fluffy 的邮箱**只存在一个地方**——修改只需改一次
- 视频和用户通过 `user_id` 连接——不需要重复个人信息
- 可以设置不同的访问权限——个人信息和视频数据分开管理

### 2.3 主键（Primary Key）与外键（Foreign Key）

#### 主键（Primary Key, PK）

主键是表中 **唯一标识每一行** 的字段。

规则：
- 值必须 **唯一**——同一张表中不能有两行主键值相同
- 不能为空（NOT NULL）——每行必须有一个主键值
- 通常是一个简单的 ID 数字（如 `user_id`、`video_id`），并设置 **AUTO_INCREMENT** 让数据库自动递增分配

在上面的例子中：`user_id` 是 Users 表的主键，`video_id` 是 Videos 表的主键。

#### 外键（Foreign Key, FK）

外键是一张表中引用另一张表主键的字段，用于 **建立表与表之间的关系**。

在 Videos 表中，`user_id` 就是外键——它引用 Users 表的 `user_id` 主键。

外键的规则：
- 外键的值必须 **匹配被引用表的主键值**，或者为空（取决于设置）
- 外键确保 **数据完整性（Data Integrity）**——不能给一个不存在的用户添加视频
- 如果没有外键约束，数据库不会阻止你插入无效的 `user_id`

```
Users 表                          Videos 表
┌─────────┐                      ┌──────────┐
│ user_id │◄──── 主键             │ video_id │ ◄── 主键
│ username│     (PK)             │ user_id  │ ◄── 外键 (FK)
│ email   │                      │ title    │     引用 Users.user_id
└─────────┘                      │ link     │
     │                           │ upload_at│
     │         一对多关系          └──────────┘
     │  一个用户可以有多个视频
     └──────────────────────────────┘
```

### 2.4 实体关系图（ERD）

当表的数量增加时，文字描述变得难以理解。**ERD（Entity Relationship Diagram，实体关系图）** 用可视化的方式展示表之间的关系。

#### 基本组成

| 元素 | 含义 |
|------|------|
| **实体（Entity）** | 一个方框 = 一张表 |
| **属性（Attribute）** | 方框内列出的字段 |
| **标识符（Identifier）** | 主键，通常在方框最上方，标注 PK |
| **关系线** | 连接两个实体的线，表示它们之间有关联 |

#### 鸦脚符号（Crow's Foot Notation）

关系线两端的符号表示 **基数（Cardinality）**——即"一个 A 对应多少个 B"：

```
──||──     恰好一个（Exactly One）
──|O──     零个或一个（Zero or One）
──|<──     一个或多个（One or Many）
──O<──     零个或多个（Zero or Many）
──><──     多个（Many）
```

#### MeowTube 的 ERD

```
┌─────────────────┐              ┌─────────────────┐
│     Users       │              │     Videos      │
├─────────────────┤              ├─────────────────┤
│ PK user_id  INT │──────||──O<──│ PK video_id INT │
│    username VCH │              │ FK user_id  INT │
│    email    VCH │              │    title    VCH │
└─────────────────┘              │    link     VCH │
                                 │    upload_  DATE│
 每个用户 → 零个或多个视频         └─────────────────┘
 每个视频 → 恰好一个用户           这是 一对多 (1:N) 关系
```

#### 三种关系类型

| 类型 | 含义 | 示例 |
|------|------|------|
| **一对一 (1:1)** | 一个 A 对应恰好一个 B | 一个用户对应一个用户详情（很少见） |
| **一对多 (1:N)** | 一个 A 对应多个 B | 一个用户上传多个视频 |
| **多对多 (M:N)** | 多个 A 对应多个 B | 多个学生选修多门课程（需要中间表） |

---

## 三、CRUD 操作与 SQL 基础

### 3.1 什么是 CRUD

**CRUD** 是数据库四种基本操作的首字母缩写，几乎所有应用和服务都基于这四种操作：

| 字母 | 操作 | SQL 命令 | 含义 |
|------|------|---------|------|
| **C** | Create（创建） | `INSERT` | 添加新数据 |
| **R** | Read（读取） | `SELECT` | 查询/检索数据 |
| **U** | Update（更新） | `UPDATE` | 修改已有数据 |
| **D** | Delete（删除） | `DELETE` | 移除数据 |

### 3.2 SQL 是什么

**SQL（Structured Query Language，结构化查询语言）** 是与数据库沟通的标准语言。无论你是添加数据、查找数据、修改数据还是删除数据，都用 SQL 来告诉数据库"做什么"。

课程使用的是 **MySQL**——世界上最流行的关系型数据库管理系统（DBMS）之一，被 Google、NASA 等公司使用。开源、轻量，适合从小型项目到大规模系统。

其他常见的关系型 DBMS（补充）：

| DBMS | 特点 |
|------|------|
| **MySQL** | 开源，最流行，Web 应用首选 |
| **PostgreSQL** | 开源，功能更丰富，支持更复杂的查询和数据类型 |
| **SQLite** | 轻量级，嵌入式，单文件，适合移动应用和小工具 |
| **Microsoft SQL Server** | 微软生态，企业级 |
| **Oracle Database** | 企业级，银行和大型系统常用 |

它们都"说" SQL（虽然各自有些小方言），核心语法是相同的。

### 3.3 创建数据库和表

#### 创建数据库

```sql
-- 创建数据库
CREATE DATABASE meowtube;

-- 查看所有数据库
SHOW DATABASES;

-- 切换到该数据库
USE meowtube;

-- 确认当前数据库
SELECT DATABASE();
```

#### 创建表——先建 Users（被引用的表先建）

```sql
CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL
);
```

逐行解析：
| 语句 | 含义 |
|------|------|
| `user_id INT` | 整数类型的字段 |
| `PRIMARY KEY` | 标记为主键——唯一且不可为空 |
| `AUTO_INCREMENT` | 数据库自动分配递增值（1, 2, 3...），无需手动指定 |
| `VARCHAR(100)` | 最长 100 个字符的可变长度字符串 |
| `VARCHAR(255)` | 255 是常见的邮箱长度上限——源自旧版 MySQL 能高效索引的最大长度，且 255 刚好可以用一个字节表示长度 |
| `NOT NULL` | 该字段不能为空——必须填写 |

#### 创建 Videos 表（包含外键）

```sql
CREATE TABLE videos (
    video_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(200) NOT NULL,
    link VARCHAR(500) NOT NULL,
    upload_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON DELETE CASCADE
);
```

关键点：
- `FOREIGN KEY (user_id) REFERENCES users(user_id)` —— 告诉 MySQL：videos 表的 `user_id` 必须对应 users 表中一个已存在的 `user_id`
- `ON DELETE CASCADE` —— 如果删除一个用户，该用户的所有视频也自动删除（级联删除）

**ON DELETE 的选项（补充）：**

| 选项 | 行为 |
|------|------|
| `CASCADE` | 删除父记录时，子记录自动删除 |
| `SET NULL` | 删除父记录时，子记录的外键设为 NULL |
| `RESTRICT` / `NO ACTION` | 如果有子记录存在，阻止删除父记录（默认行为） |

**外键如何实现一对多关系？**

`user_id` 在 Users 表中是主键（每个值只能出现一次），但在 Videos 表中是普通字段（同一个值可以出现多次）。所以一个用户可以有多个视频，但每个视频只属于一个用户。

如果要实现一对一，就需要在 Videos 表的 `user_id` 上加 `UNIQUE` 约束。如果要实现多对多，则需要一张额外的连接表（Junction Table）。

#### 查看表结构

```sql
DESCRIBE users;
DESCRIBE videos;
```

### 3.4 Create——插入数据

```sql
-- 插入用户（不需要指定 user_id，AUTO_INCREMENT 自动处理）
INSERT INTO users (username, email) VALUES
    ('Fluffy', 'fluffy@email.com'),
    ('Paws', 'paws@email.com');

-- 插入视频
INSERT INTO videos (user_id, title, link, upload_date) VALUES
    (1, 'Cat Skateboard', 'meowtube/watch/1', '2025-05-25'),
    (2, 'Epic Cat Jump', 'meowtube/watch/2', '2025-05-27'),
    (1, 'Cat vs Curtain', 'meowtube/watch/3', '2025-05-29');
```

注意事项：
- `VALUES` 中值的顺序必须与前面列出的字段顺序一致
- 不需要包含 `user_id`（主键 + AUTO_INCREMENT），MySQL 自动分配
- 虽然技术上可以手动指定主键值，但让 MySQL 处理更安全，避免重复

### 3.5 Read——查询数据

```sql
-- 查看所有列
SELECT * FROM users;
SELECT * FROM videos;

-- 只查看特定列
SELECT title FROM videos;
SELECT title, upload_date FROM videos;
```

`*` 是通配符，意思是"所有列"。

### 3.6 Update——修改数据

```sql
-- 把 video_id=1 的标题改为 'Fluffy on Skateboard'
UPDATE videos
SET title = 'Fluffy on Skateboard'
WHERE video_id = 1;
```

**⚠️ 关键提醒：`WHERE` 子句极其重要！** 如果忘了写 `WHERE`，`UPDATE` 会修改表中的 **每一行**：

```sql
-- 危险！这会把所有视频的标题都改成同一个
UPDATE videos SET title = 'Fluffy on Skateboard';
```

验证修改：
```sql
SELECT * FROM videos;
```

### 3.7 Delete——删除数据

```sql
-- 删除 video_id=3 的视频
DELETE FROM videos WHERE video_id = 3;

-- 验证
SELECT * FROM videos;
```

**同样的警告：一定要带 `WHERE`！** `DELETE FROM videos;` 不带条件会清空整个表。

**测试级联删除：**

```sql
-- 删除 user_id=1（Fluffy）
DELETE FROM users WHERE user_id = 1;

-- 查看视频表——Fluffy 的所有视频应该也被自动删除了
SELECT * FROM videos;
```

因为设置了 `ON DELETE CASCADE`，删除 Fluffy 用户后，他的所有视频也随之消失。

---

## 四、非关系型数据库（NoSQL）

### 4.1 关系型数据库的扩展性瓶颈

当产品属性多样且不统一时，关系型表会遇到麻烦。课程举了一个在线商店的例子：

最初一张 Products 表工作得很好。但很快你需要给不同产品添加不同属性——衣服需要尺码和颜色，电器需要电压，建材需要材质。你面临两个烂选择：

**选择 A：加列**——大量列是空的（不是所有产品都有"电压"），而且每加一类新属性就需要修改表结构（schema）

**选择 B：EAV 模式（Entity-Attribute-Value）**——每个属性单独一行：

| product_id | attribute | value |
|-----------|-----------|-------|
| 42 | color | red |
| 42 | size | M |
| 42 | price | 14.99 |

看起来灵活，但想象 1000 万产品、每个 100 个属性 = **10 亿行**。搜索"中号红色且价格低于 150"需要数据库分别查找每个条件，再交叉比对结果——极慢。

这就是 **非关系型数据库（NoSQL）** 可能更适合的场景。

### 4.2 NoSQL 的含义

NoSQL 全称 **Not Only SQL**——不仅仅是 SQL。它不是"反对 SQL"，而是"不只限于关系型表格模型"。

### 4.3 六种主要的 NoSQL 数据库类型

#### 文档数据库（Document Store）

最常见的 NoSQL 类型。数据以 **文档（Document）** 的形式存储，通常是 JSON 格式。相关数据放在一起：

```json
{
    "video_id": 1,
    "title": "Cat Skateboard",
    "link": "meowtube/watch/1",
    "upload_date": "2025-05-25",
    "uploader": "Fluffy",
    "comments": [
        { "user": "Paws", "text": "Amazing!", "date": "2025-05-26" },
        { "user": "Whiskers", "text": "LOL", "date": "2025-05-26" }
    ],
    "dislikes": 2
}
```

核心优势：
- 想加新字段（如 `dislikes`）？直接加进文档，不需要修改表结构
- 一次查询就能加载一个视频的所有信息（包括评论），不需要 JOIN 多张表
- 代表产品：**MongoDB**、CouchDB

#### 键值数据库（Key-Value Store）

像一个超级强大的字典——每条数据有一个唯一的键和对应的值：

```
"user:1"   → { "username": "Fluffy", "email": "fluffy@email.com" }
"video:1"  → { "title": "Cat Skateboard", "link": "/watch/1" }
```

- 查找速度极快——告诉我键，我立刻给你值
- 但不能像文档数据库那样自动把相关数据聚合在一起
- 代表产品：**Redis**、Amazon DynamoDB

#### 宽列数据库（Wide Column Store）

乍一看像普通表格，但有一个关键区别：**每一行可以有不同的列**。

```
Video 1: { title, link, daily_views, likes, shares }
Video 2: { title, link }   ← 只有基础信息
Video 3: { title, link, daily_views, regional_stats }
```

- 适合数据结构差异大的海量数据集
- 设计为可以跨多台机器分布式存储
- 代表产品：**Apache Cassandra**、HBase

#### 图数据库（Graph Database）

数据以 **节点（Node）** 和 **边（Edge）** 的形式存储：

```
(Fluffy) ──上传──→ (Cat Skateboard)
                         ↑
(Paws) ──────点赞───────┘
(Paws) ──上传──→ (Epic Cat Jump)
(Whiskers) ──关注──→ (Fluffy)
```

- 节点 = 事物（用户、视频、评论）
- 边 = 关系（上传、点赞、关注）
- 在关系型数据库中用 JOIN 查询"所有点赞了 Fluffy 上传的视频的用户"会很慢；图数据库天生擅长沿着这些关系高效遍历
- 代表产品：**Neo4j**、Amazon Neptune
- 应用场景：社交网络、推荐系统、欺诈检测

#### 时序数据库（Time Series Database）

专为 **持续变化的、带时间戳的数据** 设计。不是更新一条记录，而是每次变化都存一个新条目：

```
2025-07-01  Video:1  views: 500
2025-07-02  Video:1  views: 1450
2025-07-03  Video:1  views: 3200
```

- 容易发现趋势、做时间段对比、保留完整历史
- 代表产品：**InfluxDB**、TimescaleDB
- 应用场景：服务器监控、IoT 传感器数据、股票行情

#### 向量数据库（Vector Database）

数据存储为 **向量（一组数字列表）**，通常由 AI 模型生成，用于捕捉"含义"或"相似度"。

搜索"cat skateboard"不仅能找到完全匹配的结果，还能找到"muffin kitten longboard"这样词不同但语义相似的内容。

- 不是匹配精确 ID，而是找 **最近邻**——数学意义上最相似的项
- 代表产品：**Pinecone**、Milvus、Weaviate
- 应用场景：AI 搜索、推荐系统、语义匹配

### 4.4 关系型 vs 非关系型——全面对比

课程用了"图书馆 vs 文件柜"的类比：

| 维度 | 关系型（图书馆） | 非关系型（文件柜） |
|------|----------------|------------------|
| **结构** | 每本书严格按同一套分类系统编目（固定 Schema）| 每个抽屉可以有不同的组织方式，随时加新信息（灵活/无 Schema） |
| **数据完整性** | 图书管理员检查每一个细节，ISBN 不能重复（强一致性） | 可以先快速存入，之后再整理（最终一致性 / Eventual Consistency） |
| **性能** | 擅长深度研究——跨多个书架交叉引用（复杂 JOIN 查询） | 如果知道要打开哪个抽屉，直接拿出来就行（简单查询极快） |
| **扩展性** | 图书馆满了就建更大的图书馆（垂直扩展 / Scale Up）| 加更多文件柜并排放（水平扩展 / Scale Out） |

**关键理解：不是"哪个更好"，而是"当前需求适合哪个"。很多真实系统同时使用两种类型。**

- 需要严格一致性和复杂关联查询（如银行系统）→ 关系型
- 需要灵活结构和快速读写的海量数据（如社交媒体动态、商品标签）→ NoSQL
- 需要 AI 语义搜索 → 向量数据库
- 需要高频时间序列数据 → 时序数据库

### 4.5 MongoDB CRUD 演示

课程演示了如何在 MongoDB（最流行的文档数据库）中完成相同的 CRUD 操作，与 SQL 做对照。

#### 创建数据库和集合（Collection = SQL 中的表）

```javascript
// 切换到数据库（不存在则自动创建）
use meowtube

// 创建集合（类似 SQL 的 CREATE TABLE，但不需要定义字段和类型）
db.createCollection("users")
db.createCollection("videos")

// 查看集合
show collections
```

与 SQL 的关键区别：**不需要预先定义字段和数据类型**（No Schema）。

#### Create——插入数据

```javascript
db.users.insertMany([
    { username: "Fluffy", email: "fluffy@email.com" },
    { username: "Paws", email: "paws@email.com" }
])

db.videos.insertMany([
    { user_id: 1, title: "Cat Skateboard", link: "meowtube/watch/1", upload_date: "2025-05-25" },
    { user_id: 2, title: "Epic Cat Jump", link: "meowtube/watch/2", upload_date: "2025-05-27" },
    { user_id: 1, title: "Cat vs Curtain", link: "meowtube/watch/3", upload_date: "2025-05-29" }
])
```

MongoDB 自动为每个文档分配一个唯一的 `_id`。

#### Read——查询数据

```javascript
// 查看所有用户
db.users.find()

// 查看所有视频
db.videos.find()

// 只显示标题（1 = 显示，0 = 隐藏）
db.videos.find({}, { title: 1 })
```

#### Update——修改数据

```javascript
db.videos.updateOne(
    { title: "Cat Skateboard" },          // 查找条件
    { $set: { title: "Fluffy on Skateboard" } }  // 要修改的内容
)
```

#### Delete——删除数据

```javascript
db.videos.deleteOne({ title: "Cat vs Curtain" })
```

#### SQL vs MongoDB 对照表

| 操作 | SQL (MySQL) | MongoDB |
|------|------------|---------|
| 创建库 | `CREATE DATABASE x` | `use x` |
| 创建表/集合 | `CREATE TABLE x (...)` | `db.createCollection("x")` |
| 插入 | `INSERT INTO x VALUES (...)` | `db.x.insertMany([{...}])` |
| 查询全部 | `SELECT * FROM x` | `db.x.find()` |
| 查询特定列 | `SELECT title FROM x` | `db.x.find({}, {title:1})` |
| 更新 | `UPDATE x SET ... WHERE ...` | `db.x.updateOne({...}, {$set:{...}})` |
| 删除 | `DELETE FROM x WHERE ...` | `db.x.deleteOne({...})` |

---

## 五、SQL 查询进阶与索引

### 5.1 SQL 子句的执行顺序

SQL 查询中各子句必须按特定顺序书写，课程的比喻是"烤蛋糕——材料重要，但步骤顺序也必须对"：

```sql
SELECT columns        -- 要显示哪些列
FROM table            -- 从哪张表
JOIN other_table ON   -- 连接其他表（如果需要）
WHERE condition       -- 筛选条件
ORDER BY column       -- 排序
LIMIT n;              -- 限制返回行数
```

**记忆顺序：FROM → JOIN → WHERE → ORDER BY → LIMIT**

逻辑上的理解：
1. **FROM / JOIN**：先确定从哪些表取数据
2. **WHERE**：筛选出符合条件的行
3. **ORDER BY**：对结果排序
4. **LIMIT**：截取前 N 条

### 5.2 WHERE——筛选

```sql
-- 只显示 2025-05-27 上传的视频
SELECT * FROM videos
WHERE upload_date = '2025-05-27';
```

### 5.3 ORDER BY——排序

```sql
-- 按上传日期降序（最新的在前）
SELECT * FROM videos ORDER BY upload_date DESC;

-- 升序（最旧的在前）
SELECT * FROM videos ORDER BY upload_date ASC;
```

### 5.4 LIMIT——限制结果数量

```sql
-- 只显示最新的 2 个视频
SELECT * FROM videos ORDER BY upload_date DESC LIMIT 2;
```

`LIMIT` 永远放在最后——你需要先确定要什么、怎么排，最后才截取。

### 5.5 JOIN——连接多张表

```sql
-- 把视频和用户信息合在一起显示
SELECT *
FROM videos
JOIN users ON videos.user_id = users.user_id
ORDER BY upload_date DESC;
```

- `JOIN` 紧跟在 `FROM` 之后——先确定数据来源
- `ON` 指定连接条件——两张表中哪个字段相匹配
- **点号表示法（Dot Notation）**：当两张表有同名字段时，用 `表名.字段名` 来区分，如 `videos.user_id` 和 `users.user_id`

**JOIN 的常见类型（补充）：**

| 类型 | 含义 |
|------|------|
| **INNER JOIN（默认）** | 只返回两张表中匹配的行 |
| **LEFT JOIN** | 返回左表所有行，右表无匹配的地方填 NULL |
| **RIGHT JOIN** | 返回右表所有行，左表无匹配的地方填 NULL |
| **FULL OUTER JOIN** | 返回两张表的所有行（MySQL 不直接支持） |

### 5.6 索引（Index）——数据库查询性能的秘密武器

#### 问题：全表扫描

没有索引时，数据库查询数据的方式就像从头到尾翻阅一整本书来找一句话——它需要检查 **每一行**。数据量小时无所谓，百万行时就是灾难。

#### 索引是什么

索引就像书后面的"索引页"——一组排好序的书签，告诉数据库"如果你找的是 7 月份的视频，直接跳到第 12 行"。数据库不需要逐行翻找，直接跳到目标位置。

#### 课程的性能演示

课程插入了 **100 万条视频数据**，然后执行查询：

```sql
-- "Fluffy 在 7 月上传了多少视频？"
SELECT COUNT(*)
FROM videos
JOIN users ON videos.user_id = users.user_id
WHERE users.username = 'Fluffy'
AND upload_date BETWEEN '2025-07-01' AND '2025-07-31';
```

**无索引：约 2 秒**

用 `EXPLAIN` 分析查询计划：
- Users 表：`key = NULL`（没有使用任何索引，全表扫描）
- Videos 表：`Extra: Using where`（找到 Fluffy 的视频后，还要逐一检查日期）

```sql
EXPLAIN SELECT COUNT(*) FROM videos
JOIN users ON videos.user_id = users.user_id
WHERE users.username = 'Fluffy'
AND upload_date BETWEEN '2025-07-01' AND '2025-07-31';
```

#### 创建索引

```sql
-- 在 username 上创建索引：让数据库能直接定位到 "Fluffy"
CREATE INDEX idx_username ON users(username);

-- 在 user_id + upload_date 上创建复合索引：
-- 让数据库能直接定位到"Fluffy 的 7 月视频"，一步到位
CREATE INDEX idx_user_date ON videos(user_id, upload_date);
```

**创建索引耗时约 3 秒**（一次性成本）。

**有索引后执行同一查询：约 0.2 秒（快了约 10 倍）**

`EXPLAIN` 现在显示：
- Users 表：使用 `idx_username`——直接定位到 Fluffy
- Videos 表：使用 `idx_user_date`——直接跳到 Fluffy 的 7 月上传记录

**真实世界的影响：**
- 演示只跑了 1 次查询。实际应用可能每天跑这个查询数千次
- 每次查询从 2 秒降到 0.2 秒，意味着同样的硬件可以处理 **20 倍的流量**
- 索引的创建是一次性成本，但每次查询都受益

#### 索引的代价（补充）

索引不是免费的午餐：

| 代价 | 说明 |
|------|------|
| **存储空间** | 索引本身占用磁盘空间 |
| **写入变慢** | 每次 INSERT、UPDATE、DELETE 都需要同时更新索引 |
| **维护成本** | 索引可能随时间碎片化，需要偶尔重建 |

**所以不是"给每个字段都建索引"——而是给经常被 WHERE、JOIN、ORDER BY 使用的字段建索引。**

### 5.7 良好的查询习惯

| 习惯 | 说明 |
|------|------|
| **只查需要的列** | 用 `SELECT title, date` 而不是 `SELECT *`——减少数据库和网络负担 |
| **限制结果数量** | `LIMIT 10` 比返回一万行更快更实用 |
| **保持查询简单** | 与其写一个超复杂的巨型查询，不如拆成几个简单查询 |
| **善用索引和过滤** | 让数据库通过索引直接跳到目标，而不是逐行扫描 |

---

## 六、完整实战——Feline Foods 外卖应用

### 6.1 需求分析

课程最后用一个完整的外卖应用案例（Feline Foods）串联所有知识。功能类似 Uber Eats / DoorDash，顾客可以选择套餐并下单。

需要追踪的信息：
- **顾客**：谁在订餐
- **套餐**：菜单上有什么
- **订单**：谁点了什么、什么时候点的

### 6.2 设计 ERD

```
┌─────────────────┐         ┌─────────────────┐
│   Customers     │         │     Combos      │
├─────────────────┤         ├─────────────────┤
│PK customer_id   │         │PK combo_id      │
│   name    VCH   │         │   name    VCH   │
│   address TEXT   │         │   price DECIMAL │
│   contact VCH   │         └────────┬────────┘
└────────┬────────┘                  │
         │                           │
         │    ┌───────────────┐      │
         │    │    Orders     │      │
         │    ├───────────────┤      │
         └──O<│PK order_id    │>O────┘
              │FK customer_id │
              │FK combo_id    │
              │   order_time  │
              └───────────────┘

Customers 1 ──── O< Orders >O ──── 1 Combos
每个顾客可以有零或多个订单
每个订单属于恰好一个顾客
每个套餐可以出现在零或多个订单中
每个订单包含恰好一个套餐
```

### 6.3 建表

```sql
CREATE DATABASE feline_foods;
USE feline_foods;

-- 顾客表
CREATE TABLE customers (
    customer_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,
    contact VARCHAR(20) NOT NULL
);
```

**为什么 `contact` 用 `VARCHAR(20)` 而不是 `INT`？** 因为电话号码包含前导零、加号、括号、短横线等格式字符。如果用整数类型，`+86-138-0000-0000` 的格式就全丢了。`VARCHAR(20)` 保持格式灵活，20 个字符足够容纳任何格式的电话号码。

```sql
-- 套餐表
CREATE TABLE combos (
    combo_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(4,2)
);
```

`DECIMAL(4,2)` 表示最多 4 位数字、其中 2 位小数——允许的最大值是 99.99。

```sql
-- 订单表（包含两个外键）
CREATE TABLE orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    combo_id INT,
    order_time DATETIME,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id),
    FOREIGN KEY (combo_id) REFERENCES combos(combo_id)
);
```

### 6.4 插入数据

```sql
INSERT INTO customers (name, address, contact) VALUES
    ('Cody', '42 Feline Farm', '5550100'),
    ('Fluffy', '12 Catlin Close', '5550101'),
    ('Paws', '99 Whisker Way', '5550102');

INSERT INTO combos (name, price) VALUES
    ('Fish and Chips', 8.50),
    ('Sushi Bento', 12.00),
    ('Pasta Combo', 10.00);

INSERT INTO orders (customer_id, combo_id, order_time) VALUES
    (1, 3, '2025-08-10 12:00:00'),
    (2, 2, '2025-08-10 13:00:00'),
    (3, 1, '2025-08-11 18:00:00');
```

### 6.5 多表查询

```sql
-- 查看所有订单的详情（谁点了什么、什么时候）
SELECT orders.order_id,
       customers.name,
       combos.name,
       orders.order_time
FROM orders
JOIN customers ON orders.customer_id = customers.customer_id
JOIN combos ON orders.combo_id = combos.combo_id;
```

这里做了两个 JOIN——先连接顾客信息，再连接套餐信息，把三张表的数据合并展示。

```sql
-- 最新的订单排在前面
... ORDER BY order_time DESC;

-- 只看最近一个订单
... ORDER BY order_time DESC LIMIT 1;

-- 只看 Cody 的订单
... WHERE customers.name = 'Cody';
```

### 6.6 Update 与 Delete

```sql
-- 修正套餐名称
UPDATE combos SET name = 'Sushi Combo' WHERE combo_id = 2;

-- 取消一个订单
DELETE FROM orders WHERE order_id = 2;
```

课程中老师实际操作时犯了一个错：写了 `WHERE item_id = 2` 而不是 `WHERE combo_id = 2`——MySQL 报错提示"未知列 item_id"。这是一个好的提醒：**认真读错误信息，它通常直接告诉你哪里错了。**

---

## 七、总结与知识地图

```
                         ┌──────────────────┐
                         │   数据库基础      │
                         └────────┬─────────┘
          ┌──────────────────────┼──────────────────────┐
          ▼                      ▼                      ▼
    ┌───────────┐         ┌───────────┐          ┌───────────┐
    │ 基本概念   │         │ 关系型设计 │          │ SQL CRUD  │
    │           │         │           │          │           │
    │ 数据vs信息 │         │ 平面文件问题│         │ CREATE    │
    │ 字段/记录  │         │ 拆表+连接  │         │ INSERT    │
    │ 数据类型   │         │ 主键/外键  │         │ SELECT    │
    │ 无处不在   │         │ ERD / 鸦脚 │         │ UPDATE    │
    └───────────┘         │ 一对多/多对多│        │ DELETE    │
                          └───────────┘          │ JOIN      │
                                                 └───────────┘
          ┌──────────────────────┼──────────────────────┐
          ▼                      ▼                      ▼
    ┌───────────┐         ┌───────────┐          ┌───────────┐
    │ NoSQL     │         │ 查询进阶   │          │ 实战项目   │
    │           │         │           │          │           │
    │ 文档数据库 │         │ WHERE     │          │ ERD 设计  │
    │ 键值数据库 │         │ ORDER BY  │          │ 建表      │
    │ 宽列数据库 │         │ LIMIT     │          │ 插入数据  │
    │ 图数据库   │         │ JOIN      │          │ 多表查询  │
    │ 时序数据库 │         │ 索引(Index)│         │ 修改/删除 │
    │ 向量数据库 │         │ EXPLAIN   │          │ 完整流程  │
    │ SQL对比   │         │ 查询习惯   │          └───────────┘
    └───────────┘         └───────────┘
```

### 核心要点回顾

1. **数据库 = 结构化存储 + 关联查询**。不只是存数据，更重要的是让数据变成可搜索、可分析的信息
2. **关系型数据库通过拆表 + 主键/外键解决冗余和不一致**——一条信息只存一个地方，用键来连接
3. **ERD 是数据库设计的蓝图**——先画 ERD 再写 SQL，就像先画图纸再盖房子
4. **CRUD 是一切数据库操作的基础**——Create / Read / Update / Delete
5. **SQL 子句有严格的顺序**：FROM → JOIN → WHERE → ORDER BY → LIMIT
6. **索引是性能的王牌**——创建一次，每次查询都受益。百万行数据查询速度可提升 10 倍以上
7. **NoSQL 不是替代关系型，而是针对不同场景的工具**——文档、键值、图、时序、向量各有所长
8. **现实系统通常混合使用**——用关系型处理需要强一致性的核心数据，用 NoSQL 处理灵活或高吞吐的数据

---

## 附录：SQL 命令速查表

```sql
-- ===== 数据库管理 =====
CREATE DATABASE db_name;
SHOW DATABASES;
USE db_name;
SELECT DATABASE();
DROP DATABASE db_name;        -- 删除数据库（危险！）

-- ===== 表管理 =====
CREATE TABLE table_name (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    FOREIGN KEY (col) REFERENCES other_table(col)
);
DESCRIBE table_name;           -- 查看表结构
DROP TABLE table_name;         -- 删除表（危险！）

-- ===== CRUD =====
-- Create
INSERT INTO table (col1, col2) VALUES ('val1', 'val2');

-- Read
SELECT * FROM table;
SELECT col1, col2 FROM table WHERE condition;
SELECT * FROM table ORDER BY col DESC LIMIT 10;
SELECT * FROM t1 JOIN t2 ON t1.id = t2.fk_id;

-- Update
UPDATE table SET col = 'new_value' WHERE id = 1;

-- Delete
DELETE FROM table WHERE id = 1;

-- ===== 索引 =====
CREATE INDEX idx_name ON table(column);
CREATE INDEX idx_name ON table(col1, col2);  -- 复合索引
EXPLAIN SELECT ...;                          -- 分析查询计划
```
