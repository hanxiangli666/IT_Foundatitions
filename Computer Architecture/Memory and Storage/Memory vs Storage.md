# 内存与存储的差别 / Memory vs Storage

> 中文：这是一份中英文对照的学习笔记，重点解释内存和存储为什么会影响系统性能、为什么“CPU 很强但电脑还是卡”，以及为什么虚拟内存、缓存、寄存器和存储技术会共同决定真实体验。
>
> English: This is a bilingual study note focused on why memory and storage affect system performance, why a computer can still feel slow even with a powerful CPU, and why virtual memory, cache, registers, and storage technology together determine the real user experience.

![Memory vs Storage overview](image/MemoryvsStorage/1774750518094.png)

## 1. 为什么要区分内存和存储 / Why Memory and Storage Must Be Distinguished

中文：很多性能问题看起来像“CPU 不够快”，实际上却是内存或存储拖慢了系统。CPU 计算速度非常高，但它不能凭空处理数据，它需要把数据从更快的临时工作区拿到手里。如果数据不在 RAM 里，而是在更慢的存储里，CPU 就会等待。这个等待通常比计算本身更耗时，所以系统会显得卡顿、冻结或者响应迟缓。

English: Many performance problems look like “the CPU is too slow,” but in reality they are caused by memory or storage bottlenecks. A CPU can compute very quickly, but it cannot process data out of thin air. It needs to bring data into a fast working area first. If the data is not in RAM and instead lives in slower storage, the CPU ends up waiting. That waiting is often more expensive than the computation itself, so the whole system feels laggy, frozen, or unresponsive.

中文：理解这一点之后，你会发现“容量大”和“速度快”不是一回事。存储容量大，只代表你能保存更多文件；内存速度快，才代表程序运行时能更顺畅地读写当前数据。很多人会误以为 512 GB 或 1 TB 的硬盘可以替代更多 RAM，但这两者在系统里的角色根本不同。

English: Once you understand this, it becomes clear that “large capacity” is not the same as “fast speed.” A larger storage drive only means you can keep more files; faster memory means the running program can read and write active data smoothly. Many people assume a 512 GB or 1 TB drive can replace more RAM, but those two components play fundamentally different roles in the system.

---

## 2. 典型故障场景 / A Typical Slow-System Scenario

中文：一个常见场景是：你给设计团队配了性能不错的 CPU 和 GPU，但他们在大型项目上还是会频繁卡顿。第一步不应该是立刻怀疑处理器，而是检查 CPU、GPU、内存和磁盘的活动情况。如果 CPU 和 GPU 使用率都不高，而磁盘活动明显升高，通常说明系统正在做页面交换，也就是把部分内存内容挪到磁盘上。

English: A common scenario is this: you equip a design team with capable CPUs and GPUs, but they still experience frequent slowdowns on large projects. The first step is not to immediately blame the processor. Instead, check CPU, GPU, memory, and disk activity. If CPU and GPU utilization are not high, but disk activity rises sharply, it usually means the system is swapping pages, moving part of memory contents onto disk.

![Performance monitor example](image/MemoryvsStorage/1774750535053.png)

中文：页面交换能帮助系统避免立刻崩溃，但代价非常大。RAM 的访问延迟远低于 SSD，更远低于 HDD；一旦系统依赖磁盘来充当临时工作区，整台机器就会因为等待数据而变慢。这种卡顿不是单个应用的问题，而是系统级的响应变差。

English: Paging helps the system avoid immediate failure, but it comes at a significant cost. RAM latency is much lower than SSD latency, and far lower than HDD latency. Once the system has to rely on disk as temporary working space, the whole machine slows down because it keeps waiting for data. That kind of lag is not just an application issue; it becomes a system-wide responsiveness problem.

> **提示 / Tip**
> 中文：当系统开始频繁使用虚拟内存时，最直接的解决思路通常是增加 RAM 或缩小工作集，而不是单纯换更大的硬盘。
> English: When a system starts relying heavily on virtual memory, the most direct fix is usually to add RAM or reduce the working set, not simply buy a larger drive.

![Laptop activity monitor](image/MemoryvsStorage/1774750535053.png)

中文：排查时要看两个指标：物理内存是否足够，当前工作集是否超过了 RAM 容量。工作集越接近或超过物理 RAM，系统就越可能把页面换到存储里。此时容量再大的硬盘也不等于足够快的内存。

English: When troubleshooting, look at two things: whether physical memory is sufficient and whether the current working set exceeds RAM capacity. The closer the working set gets to or exceeds physical RAM, the more likely the system is to move pages into storage. At that point, even a huge drive does not equal fast enough memory.

---

## 3. 内存和存储的几个维度 / Three Ways to Compare Memory and Storage

中文：我们可以从访问方式、物理构成和系统架构三个角度来理解内存和存储的区别。这样不仅能知道它们“是什么”，还能知道它们为什么在性能上差这么多。

English: We can compare memory and storage from three angles: access method, physical composition, and system architecture. That lets us understand not just what they are, but also why they differ so much in performance.

### 3.1 访问方式 / Access Method

中文：RAM 是随机访问存储器，访问任意地址的时间大致相同，所以 CPU 能快速、直接地读取和写入它需要的数据。相比之下，某些顺序访问介质必须按顺序读取，随机定位成本很高，因此更适合特定场景而不是通用计算。

English: RAM is random access memory, which means the time to access any address is roughly the same. That allows the CPU to read and write the data it needs quickly and directly. By contrast, some sequential-access media must be read in order, so random positioning is expensive and they are better for specific use cases rather than general-purpose computing.

![Random vs sequential access](image/MemoryvsStorage/1774750552955.png)

中文：HDD 也属于随机访问设备，但由于它有机械磁头和旋转盘片，随机访问的延迟明显高于 SSD，更远高于 DRAM。SSD 没有机械运动，所以更快、更安静，也更抗震，但它仍然比 RAM 慢很多。

English: HDDs are also random-access devices, but because they use mechanical heads and spinning platters, their random-access latency is much higher than SSDs and far higher than DRAM. SSDs have no mechanical movement, so they are faster, quieter, and more shock-resistant, but they are still much slower than RAM.

### 3.2 物理构成 / Physical Composition

中文：DRAM 依靠电容和晶体管存储状态，速度很快，但断电后内容会消失，所以它是易失性的。SSD 使用 NAND 闪存，通过电荷陷阱存储数据，断电后依然能保留内容，所以它是非易失性的。HDD 则是通过磁性盘片和磁头来读写数据，容量大、成本低，但有机械磨损。

English: DRAM stores state using capacitors and transistors, so it is fast, but its contents vanish when power is removed, which makes it volatile. SSDs use NAND flash, storing data by trapping charge, so they retain contents without power and are non-volatile. HDDs store data on magnetic platters using a read/write head, which gives them high capacity and low cost, but also mechanical wear.

![Semiconductor comparison](image/MemoryvsStorage/1774750560846.png)

中文：你可以把这看成不同的工程目标。RAM 追求速度，SSD 追求速度与持久性的平衡，HDD 追求低成本和高容量。没有一种方案在所有维度上都最优，系统设计的关键就在于取舍。

English: Think of this as different engineering goals. RAM prioritizes speed, SSDs balance speed and persistence, and HDDs prioritize low cost and high capacity. No single solution is best in every dimension; the real art of system design is making the right trade-offs.

### 3.3 系统架构 / Architecture

中文：内存和存储不仅在内部构造上不同，它们和 CPU 的连接方式也不同。内存通常通过更直接、更低延迟的通道连接到内存控制器，而存储则通过 SATA、NVMe 或其他接口连接，虽然带宽很高，但延迟仍然远大于 RAM。

English: Memory and storage differ not only in internal construction, but also in how they connect to the CPU. Memory is usually attached through more direct, lower-latency channels to the memory controller, while storage connects through SATA, NVMe, or similar interfaces. These storage interfaces can offer high bandwidth, but latency is still much higher than RAM.

![System architecture comparison](image/MemoryvsStorage/1774750571728.png)

中文：这也是为什么即使 SSD 已经很快，系统仍然会把 RAM 视为更高层级的工作区。RAM 直接服务于运行中的程序，而存储主要负责持久保存。

English: That is why even though SSDs are already fast, the system still treats RAM as a higher-level working area. RAM directly serves running programs, while storage primarily provides durable persistence.

---

## 4. CPU 内存层级 / The CPU Memory Hierarchy

中文：从最快到最慢，典型内存层级大致是寄存器、缓存、主内存、本地存储、远程或云端存储。越靠近 CPU 的层级越快，但容量越小、成本越高；越远离 CPU 的层级越慢，但容量更大、价格更低。

English: From fastest to slowest, the typical memory hierarchy is registers, cache, main memory, local storage, and remote or cloud storage. The closer a level is to the CPU, the faster it is, but the smaller and more expensive it becomes; the farther away it is, the slower it gets, but the capacity grows and the cost drops.

![CPU cache and registers](image/MemoryvsStorage/1774750580751.png)

中文：寄存器位于 CPU 核心内部，是最快但最小的存储层；缓存位于核心附近，用来保存高频访问数据；主内存则是程序运行时的主要工作区。缓存和寄存器能减少 CPU 等待数据的时间，这对整体性能非常关键。

English: Registers sit inside the CPU core and are the fastest but smallest storage layer; cache is close to the core and holds frequently accessed data; main memory is the primary workspace for running programs. Cache and registers reduce the time the CPU spends waiting for data, which is crucial for overall performance.

中文：除了 RAM 和缓存，ROM 也很重要。ROM 或基于闪存的 BIOS/UEFI 保存开机所需的固件代码，它们不用于频繁读写，而是用于在上电时初始化硬件并引导系统启动。

English: In addition to RAM and cache, ROM matters too. ROM, or flash-based BIOS/UEFI, stores the firmware code needed during boot. It is not meant for frequent reads and writes; it is meant to initialize hardware and start the system when power is applied.

![ROM and BIOS](image/MemoryvsStorage/1774750586872.png)

---

## 5. 核心结论 / Key Conclusions

中文：更快的内存通常离 CPU 更近，使用更高速度的半导体技术，因此每 GB 成本更高；更慢的存储通常离 CPU 更远，容量更大，成本更低，但延迟也更高。理解这一点后，你就不会再把“容量大”误当成“速度快”。

English: Faster memory is usually physically closer to the CPU and uses higher-speed semiconductor technology, so it costs more per gigabyte; slower storage is farther away, offers larger capacity, and costs less, but has higher latency. Once you understand that, you stop confusing “more capacity” with “more speed.”

> **警告 / Warning**
> 中文：单纯升级存储容量不会替代 RAM 的速度优势。如果应用已经超过了可用内存，应该优先增加 RAM 或减少工作集，而不是只买更大的 SSD。
> English: Simply upgrading storage capacity does not replace the speed advantage of RAM. If applications already exceed available memory, you should prioritize more RAM or a smaller working set instead of only buying a larger SSD.

中文：排查卡顿系统时，优先检查 RAM 使用率、分页情况和当前工作集是否合理。如果系统已经开始频繁换页，即使 CPU 看起来很空闲，用户体验也可能很差。

English: When troubleshooting a sluggish system, prioritize RAM usage, paging activity, and whether the current working set is reasonable. If the system is already paging frequently, the user experience may be poor even if the CPU looks idle.

---

## 6. 快速参考 / Quick Reference

| 层级 / Tier | 中文理解 / Chinese meaning | English meaning | 典型特征 / Key traits |
| --- | --- | --- | --- |
| Registers | CPU 内部最快工作区 | Fastest internal CPU workspace | 极快、极小、易失 |
| Cache | 靠近核心的高速缓冲区 | High-speed buffer near the core | 很快、小、易失 |
| Main memory / DRAM | 程序运行中的主要工作区 | Primary working memory for programs | 快、容量中等、易失 |
| Local storage | SSD / HDD / NVMe | Persistent local storage | 更大、更慢、非易失 |
| Remote / cloud | 网络文件系统、对象存储 | Networked or cloud storage | 最慢、最灵活、非易失 |

---

## Further Reading

中文：如果你想继续深入，可以去看这门课里关于 CPU、内存与存储、主板和硬件的后续章节。它们会把这里提到的层级关系讲得更细。

English: If you want to go deeper, continue with the CPU, memory and storage, motherboard, and hardware chapters in this course. They will explain the hierarchy described here in much more detail.

- [Watch Video](https://learn.kodekloud.com/user/courses/computer-architecture/module/79580b70-d812-41b0-9704-6c333005a949/lesson/b7dc2762-8197-4c99-90d6-6e96cf61a75c)
