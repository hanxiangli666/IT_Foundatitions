# 存储类型 / Types of Storage

> 中文：这是一份中英文对照的存储类型笔记，重点解释磁存储、光存储、固态存储和云存储之间的差别，以及它们分别适合什么场景。
>
> English: This is a bilingual note on storage types, focused on the differences between magnetic storage, optical storage, solid-state storage, and cloud storage, and the scenarios where each one makes sense.

## 1. Memory vs Storage / 内存 vs 存储

中文：内存通常指 CPU 直接使用的快速临时存储，包括寄存器、缓存和 RAM。它们速度快，但多半是易失性的。存储则指长期保存文件和程序的介质，目标是容量和持久性，而不是极限速度。

English: Memory usually refers to the fast temporary storage used directly by the CPU, including registers, cache, and RAM. These layers are fast but mostly volatile. Storage refers to media used to keep files and programs long term, and its goal is capacity and persistence rather than extreme speed.

中文：ROM 和某些固件存储是例外。它们不属于 RAM，但也是非易失性的，所以它们既不是典型内存，也不是典型存储文件的地方，而是系统启动和基础配置的重要组成部分。

English: ROM and some firmware storage are exceptions. They are not RAM, but they are non-volatile, so they are neither typical memory nor typical file storage. Instead, they are important parts of system boot and foundational configuration.

---

## 2. 存储家族 / Storage Families

中文：存储通常可以分为四大类：磁存储、光存储、固态存储和云存储。它们使用的物理机制不同，擅长的场景也不同。

English: Storage can usually be divided into four major families: magnetic storage, optical storage, solid-state storage, and cloud storage. They use different physical mechanisms and excel in different scenarios.

### 2.1 磁存储 / Magnetic Storage

中文：磁存储包括 HDD、磁带和其他依靠磁性材料记录数据的介质。它的优势是便宜、容量大、适合批量和归档，但随机访问延迟较高。

English: Magnetic storage includes HDDs, tape, and other media that record data using magnetic materials. Its advantages are low cost, large capacity, and suitability for bulk and archival use, but random-access latency is relatively high.

![Magnetic storage](image/TypesofStorage/1774750610279.png)

中文：HDD 是最常见的磁存储形式。它靠盘片旋转和磁头移动读写数据，因此速度和机械状态紧密相关。对于大量静态数据，它依然很有价值；但对于系统盘或高频访问工作负载，它已经逐渐被 SSD 替代。

English: The HDD is the most common form of magnetic storage. It reads and writes data by spinning platters and moving a head, so its speed depends heavily on mechanical motion. It remains valuable for large amounts of static data, but for system drives or high-frequency workloads, SSDs have increasingly replaced it.

> **提示 / Tip**
> 中文：磁存储的核心优势通常是“便宜和大”，不是“快”。
> English: The core strength of magnetic storage is usually “cheap and large,” not “fast.”

### 2.2 光存储 / Optical Storage

中文：光盘类介质通过激光读取表面结构来存储和读取数据。它们适合分发、离线保存和一些兼容性场景，但随机访问速度较慢，而且容量已经落后于现代主流介质。

English: Optical media use lasers to read surface structures for storage and retrieval. They are useful for distribution, offline preservation, and some compatibility scenarios, but random-access speed is slow and capacity has fallen behind modern mainstream media.

### 2.3 固态存储 / Solid-State Storage

中文：固态存储包括 SSD、U 盘、SD 卡等。它们基于闪存，没有机械部件，因此抗震、安静、响应快，非常适合系统盘和应用程序。

English: Solid-state storage includes SSDs, USB flash drives, and SD cards. They are based on flash memory and have no moving parts, so they are shock-resistant, quiet, and responsive, making them ideal for system drives and applications.

中文：固态存储的缺点是每 GB 成本通常高于磁存储，而且写入寿命有限。不过对于大多数日常用途，这个限制已经足够可控。

English: The downside of solid-state storage is that cost per gigabyte is usually higher than magnetic storage, and write endurance is finite. For most everyday uses, however, that limitation is quite manageable.

> **警告 / Warning**
> 中文：写入密集型场景要关注闪存寿命，必要时选择更高等级的 SSD。
> English: For write-intensive workloads, pay attention to flash endurance and choose a higher-grade SSD when needed.

### 2.4 云存储 / Cloud Storage

中文：云存储是通过网络访问的远程存储服务，比如 Google Drive、OneDrive 和 AWS S3。它不是一种新的物理介质，而是把底层磁盘或闪存资源包装成可远程访问的服务。

English: Cloud storage is a remote storage service accessed over the network, such as Google Drive, OneDrive, and AWS S3. It is not a new physical medium; it is a service layer that wraps underlying disks or flash storage for remote access.

![Cloud storage](image/TypesofStorage/1774750617035.png)

中文：云存储的优势是弹性、共享和备份方便；缺点是速度依赖网络，延迟受带宽和连通性影响。它很适合协作和容灾，但不适合要求极低延迟的本地 I/O。

English: Cloud storage’s strengths are elasticity, sharing, and easy backup; its weakness is that speed depends on the network and latency is affected by bandwidth and connectivity. It is excellent for collaboration and disaster recovery, but not ideal for extremely low-latency local I/O.

---

## 3. 数据尺寸 / Data Sizes

中文：计算机以 bit 和 byte 为基本单位。bit 是 0 或 1，8 个 bit 等于 1 byte。容量单位在技术文档里最好区分二进制前缀和十进制前缀，这样才不会混淆。

English: Computers use bits and bytes as their basic units. A bit is 0 or 1, and 8 bits make 1 byte. In technical documentation, it is best to distinguish binary prefixes from decimal prefixes so that capacity reporting stays precise.

```text
01101000
```

| Unit | Value |
| --- | --- |
| 1 KiB | 1024 bytes |
| 1 MiB | 1024 KiB |
| 1 GiB | 1024 MiB |
| 1 TiB | 1024 GiB |

中文：很多厂商会把 1000 和 1024 混着写，所以你在看容量时要留意单位来源。对于日常理解，不必过度纠结；但在比较磁盘、系统报告和实际容量时，这个差异很重要。

English: Many vendors mix 1000 and 1024 in marketing copy, so you should pay attention to the source of the units. For everyday understanding, you do not need to obsess over it, but when comparing drives, system reports, and actual capacity, the difference matters.

---

## 4. 存储技术的对比 / Storage Trade-offs

中文：不同存储技术的差异可以归结为速度、成本、容量、耐用性和用途。没有一种技术在所有维度上都最强，因此要按场景选。

English: The differences among storage technologies come down to speed, cost, capacity, durability, and use case. No single technology is strongest in every dimension, so you choose according to the scenario.

| Storage Type | 中文理解 / Chinese meaning | English meaning | 典型用途 / Typical use |
| --- | --- | --- | --- |
| Optical | 便于分发与离线保存 | Good for distribution and offline retention | 介质发放、归档 |
| Magnetic | 低成本大容量 | Low cost, high capacity | 备份、冷数据、归档 |
| Solid-state | 高速、抗震 | Fast and shock-resistant | 系统盘、应用、移动设备 |
| Cloud | 远程、可扩展 | Remote and scalable | 备份、协作、共享 |

中文：如果你的目标是性能，优先考虑 SSD；如果你的目标是便宜和容量，HDD 仍然有价值；如果你的目标是分享和远程访问，云存储更方便；如果你的目标是长期离线保存，光盘或磁带可能仍然适用。

English: If your goal is performance, prioritize SSDs; if your goal is low cost and capacity, HDDs still make sense; if your goal is sharing and remote access, cloud storage is more convenient; if your goal is long-term offline preservation, optical media or tape may still be appropriate.

---

## 5. 位置与层级 / Where Storage Sits in the Hierarchy

中文：从 CPU 往外看，寄存器和缓存最快，RAM 次之，SSD 或 HDD 更慢，光存储和云存储通常更慢。层级越往外，容量越大、单位成本越低，但延迟也越高。

English: From the CPU outward, registers and cache are fastest, RAM comes next, SSDs and HDDs are slower, and optical or cloud storage is often slower still. The farther out you go in the hierarchy, the larger and cheaper the capacity tends to be, but latency also rises.

中文：这也是为什么系统会先把热数据放在内存里，再根据需要写到存储上。性能优化的目标，就是让最常用的数据停留在最快的层级。

English: That is why the system first keeps hot data in memory and only writes it to storage when needed. The goal of performance optimization is to keep the most frequently used data in the fastest possible layer.

---

## 6. 选择建议 / How to Choose Storage

中文：如果你关心速度，优先选 NVMe SSD；如果你关心大容量和低成本，选 HDD；如果你需要跨设备访问、共享和备份，选云存储；如果你要做离线长期保留，再考虑光学介质或磁带。

English: If speed matters most, choose an NVMe SSD; if large capacity and low cost matter most, choose an HDD; if you need cross-device access, sharing, and backup, choose cloud storage; if you need long-term offline retention, consider optical media or tape.

---

## 7. 小结 / Summary

中文：内存和存储不是竞争关系，而是分工关系。内存负责快、临时、频繁变化的数据，存储负责慢一些但持久的数据。理解这一点后，你就能更准确地判断系统瓶颈，也能更好地选择硬件。

English: Memory and storage are not in competition; they are complementary. Memory handles fast, temporary, frequently changing data, while storage handles slower but persistent data. Once you understand this, you can diagnose bottlenecks more accurately and choose hardware more wisely.

## Further Reading

- [Watch Video](https://learn.kodekloud.com/user/courses/computer-architecture/module/79580b70-d812-41b0-9704-6c333005a949/lesson/6c3b0277-d9f7-4fe1-9ced-b97f17ab4a2a)
