# 硬件 / Hardware

> 中文：这是一份中英文对照的硬件笔记，重点解释输入设备、输出设备、外设、供电、散热、主板以及内部组件之间是如何协同工作的。它不只是讲“有哪些零件”，而是讲“数据怎么流动，硬件怎么连接，为什么系统会稳定或不稳定”。
>
> English: This is a bilingual hardware note focused on how input devices, output devices, peripherals, power delivery, cooling, the motherboard, and internal components work together. It is not only about “what parts exist,” but about how data flows, how hardware is connected, and why the system is stable or unstable.

## 1. 硬件是什么 / What Hardware Is

中文：硬件是计算机系统中能够触摸到、看得见的物理部分。它包括键盘、鼠标、显示器、摄像头、麦克风、主板、CPU、GPU、内存、存储、风扇、电源等。软件告诉硬件做什么，而硬件负责真正执行动作。

English: Hardware is the physical part of the computer system that you can see and touch. It includes the keyboard, mouse, monitor, camera, microphone, motherboard, CPU, GPU, memory, storage, fans, and power supply. Software tells hardware what to do, and hardware performs the actual actions.

中文：理解硬件最好的方式不是把它们当成一堆零件，而是把它们看成一个协作系统。输入进来，内部处理，结果输出，整个过程是连续的。

English: The best way to understand hardware is not to treat it as a pile of parts, but as a cooperating system. Input comes in, the system processes it internally, and the result is output. The whole process is continuous.

![Hardware overview](image/Hardware/1774749524477.png)

---

## 2. 外部硬件 / External Hardware

中文：外部硬件是你和计算机直接交互的部分，通常分为输入设备和输出设备。输入设备把人的动作和环境信息送进系统，输出设备把计算结果展示给人。

English: External hardware is the part you interact with directly. It is usually divided into input devices and output devices. Input devices bring human actions and environmental information into the system, while output devices show the results back to the user.

![External hardware](image/Hardware/1774749568674.png)

### 输入设备 / Input Devices

中文：键盘和鼠标是最典型的输入设备。键盘把按键变成扫描码，鼠标把移动和点击变成信号，触摸板把手指位置变成定位信息。它们是人与机器之间最直接的桥梁。

English: Keyboards and mice are the most common input devices. A keyboard turns key presses into scan codes, a mouse turns movement and clicks into signals, and a touchpad turns finger position into pointing information. They are the most direct bridge between humans and the machine.

中文：不同鼠标类型适合不同场景。无线鼠标更灵活，有线鼠标更稳定，光学鼠标适合日常使用，滚球鼠标则属于早期技术。输入设备的差异会影响精度、速度和使用舒适度。

English: Different mouse types fit different scenarios. Wireless mice are more flexible, wired mice are more stable, optical mice are good for everyday use, and roller-ball mice are an older design. Differences in input devices affect precision, speed, and comfort.

![Input devices](image/Hardware/1774749585369.png)

中文：摄像头和麦克风属于另一类输入设备。它们把光和声音转换成数字数据，供计算机处理。也就是说，它们是把模拟世界“翻译”成数字世界的工具。

English: Cameras and microphones are another kind of input device. They convert light and sound into digital data for the computer to process. In other words, they translate the analog world into the digital world.

![Camera and microphone](image/Hardware/1774749597802.png)

### 输出设备 / Output Devices

中文：显示器和扬声器是最常见的输出设备。显示器把数字信号转换成图像，扬声器把电信号转换成声音。输出设备的作用是把系统计算的结果呈现给人。

English: Monitors and speakers are the most common output devices. A monitor converts digital signals into images, and a speaker converts electrical signals into sound. Their job is to present the result of computation to the user.

中文：如果没有输入和输出，计算机虽然还能“运行”，但人就几乎无法使用它。硬件不是为了孤立工作，而是为了完成交互。

English: Without input and output, a computer may still “run,” but it becomes almost impossible for a person to use it. Hardware is not meant to work in isolation; it is meant to enable interaction.

---

## 3. 端口和连接器 / Ports and Connectors

中文：外部设备并不是直接焊在主板上的，它们通过接口连接。USB、HDMI、Ethernet 和音频接口就是最常见的例子。不同接口承担不同用途：USB 连接周边设备，HDMI 连接显示设备，Ethernet 提供网络，音频接口负责声音输入和输出。

English: External devices are not soldered directly to the motherboard. They connect through interfaces. USB, HDMI, Ethernet, and audio jacks are common examples. Different interfaces serve different purposes: USB connects peripherals, HDMI connects displays, Ethernet provides networking, and audio jacks handle sound input and output.

![Ports and connectors](image/Hardware/1774749630117.png)

中文：这些接口之所以重要，是因为它们决定了扩展能力和兼容性。一个系统是否容易接显示器、键盘、网络、存储和其他外设，往往就看这些接口是否齐全。

English: These interfaces matter because they determine expandability and compatibility. Whether a system can easily connect to displays, keyboards, networks, storage, and other peripherals often depends on how complete the interface set is.

> 中文：接口标准会随着时代变化。旧接口会被新接口替代，但很多系统仍然要兼容旧设备，所以你会看到 VGA、HDMI、USB 这些不同年代的标准同时存在。
>
> English: Interface standards change over time. Older interfaces get replaced by newer ones, but many systems still need to support older devices, so you often see VGA, HDMI, and USB standards from different eras living side by side.

---

## 4. 从按键到字符 / From Keypress to Character

中文：一个很经典的例子是按下键盘上的字母 k。键盘内部的矩阵检测到按键闭合，微控制器生成扫描码，扫描码通过 USB 或其他接口传给操作系统，操作系统再把它映射成字符 k 或 K。

English: A classic example is pressing the letter k on a keyboard. The keyboard matrix detects the key closure, the microcontroller generates a scan code, the scan code is sent through USB or another interface to the operating system, and the OS maps it to the character k or K.

中文：这个流程说明外设并不只是“接上就能用”。它们会先把物理动作转换成电信号，再由操作系统和应用程序解释成有意义的内容。

English: This flow shows that peripherals are not just “plug and play” in a simple sense. They first convert physical action into electrical signal, and then the operating system and applications interpret that signal into meaningful content.

---

## 5. 电脑的内部硬件 / Internal Hardware

中文：打开机箱之后，你会看到主板、CPU、内存、存储、GPU、供电模块和散热系统。它们构成了计算机真正的“内部世界”。主板像底座和交通网络，CPU 像大脑，内存像工作台，存储像长期仓库，GPU 像专门处理图像和并行任务的工厂。

English: When you open the case, you see the motherboard, CPU, memory, storage, GPU, power circuitry, and cooling system. These form the computer’s real internal world. The motherboard is like the base and transport network, the CPU is like the brain, memory is the workbench, storage is the long-term warehouse, and the GPU is a factory specialized for graphics and parallel work.

![Motherboard and internal layout](image/Hardware/1774749643625.png)

中文：内部硬件之间的数据流动都要经过主板。它不仅负责“连起来”，还负责“让它们正确地说话”。如果连接错了，系统就会无法启动、识别不到设备，或者性能下降。

English: Data flow between internal components always passes through the motherboard. Its job is not only to “connect things,” but also to make sure they “speak correctly.” If the connection is wrong, the system may fail to boot, fail to detect devices, or lose performance.

---

## 6. 内部核心部件 / Core Internal Components

### CPU

中文：CPU 是系统的计算核心，负责执行指令、控制流程和协调各部分工作。现代 CPU 通常有多个核心和多个功能单元，因此可以同时处理更多任务。

English: The CPU is the system’s compute core, responsible for executing instructions, controlling flow, and coordinating the rest of the system. Modern CPUs usually have multiple cores and multiple functional units, so they can handle more work at the same time.

### RAM

中文：RAM 是短期记忆。CPU 在处理任务时会把当前需要的数据和程序放在 RAM 里，因为 RAM 快。RAM 是易失性的，断电后内容会消失。

English: RAM is short-term memory. The CPU keeps the data and programs it is currently using in RAM because RAM is fast. RAM is volatile, so its contents disappear when power is lost.

### Storage

中文：SSD 和 HDD 属于长期存储。HDD 便宜、容量大，适合批量保存；SSD 快得多，能明显提升开机和加载速度。它们都负责保存程序和文件，但速度差别非常大。

English: SSDs and HDDs are long-term storage. HDDs are cheaper and offer larger capacity, which makes them good for bulk storage; SSDs are much faster and noticeably improve boot and loading times. Both store programs and files, but their speed difference is very large.

![Storage and motherboard](image/Hardware/1774749652880.png)

### GPU

中文：GPU 负责图形和大规模并行任务。它的作用不是取代 CPU，而是在适合并行处理的场景里提供更高吞吐量。

English: The GPU handles graphics and large-scale parallel tasks. Its purpose is not to replace the CPU, but to provide higher throughput in workloads that are suitable for parallel processing.

---

## 7. 供电和散热 / Power and Cooling

中文：任何计算机都离不开稳定电力。笔记本通常有电池和外部电源适配器，台式机则通过电源供应器提供稳定电流。主板上的电源管理模块负责把输入电压转换成各部件所需的工作电压。

English: No computer works without stable power. Laptops usually have both a battery and an external power adapter, while desktops rely on a power supply unit. Power-management circuitry on the motherboard converts the incoming voltage into the working voltages needed by different components.

中文：散热同样关键。CPU 和 GPU 工作时会发热，如果热量散不出去，系统就会降频，甚至不稳定。风扇、散热片、热管和液冷系统的作用，就是把热量从芯片带走。

English: Cooling is just as important. CPU and GPU generate heat while running, and if that heat is not removed, the system will lower its clock speed or become unstable. Fans, heat sinks, heat pipes, and liquid cooling systems all exist to move heat away from the chip.

> 中文：供电稳定和散热良好不仅影响性能，还直接影响硬件寿命和数据安全。
>
> English: Stable power and good cooling affect not only performance, but also hardware lifespan and data safety.

![Cooling system](image/Hardware/1774749568674.png)

---

## 8. 一个完整的数据路径 / A Complete Data Path

中文：如果你把前面所有内容串起来，就会得到一个完整路径：外设采集输入，主板接收信号，CPU 处理指令，RAM 暂存数据，GPU 渲染图像，显示器输出结果。这个过程看起来很快，但它其实经历了许多层转换。

English: If you connect everything above, you get a complete path: peripherals capture input, the motherboard receives the signal, the CPU processes instructions, RAM temporarily stores data, the GPU renders images, and the monitor outputs the result. It looks instant, but it actually passes through many layers of transformation.

中文：这就是硬件学习最重要的地方：不要只记名字，要看它们之间的关系。每个部件都不是孤立存在的，它们都在同一条数据路径上。

English: This is the most important part of learning hardware: do not just memorize names, but understand the relationships between them. No component exists in isolation; they all participate in the same data path.

---

## 9. 为什么这些知识重要 / Why This Matters

中文：理解硬件有助于你判断系统为什么慢、为什么不稳定、为什么某个接口不能用、为什么升级后效果不明显。很多看似“复杂”的问题，其实都可以追溯到硬件连接、供电、散热或瓶颈部件。

English: Understanding hardware helps you judge why a system is slow, why it is unstable, why an interface does not work, or why an upgrade does not help as much as expected. Many problems that seem “complex” can actually be traced back to hardware connections, power, cooling, or bottleneck components.

中文：当你能把输入、处理、存储、输出和供电散热这几条线串起来时，你就已经具备了分析计算机系统的基础能力。

English: Once you can connect input, processing, storage, output, power, and cooling into one mental model, you already have the basic ability to analyze a computer system.

---

## 10. 小结 / Summary

中文：硬件不是零件的集合，而是一个协作系统。输入设备让人和系统沟通，主板把部件连起来，CPU 和 GPU 负责计算，RAM 负责临时工作区，存储负责长期保存，电源和散热负责稳定运行。

English: Hardware is not just a collection of parts; it is a cooperative system. Input devices let humans communicate with the system, the motherboard connects the parts, the CPU and GPU do the computing, RAM provides temporary workspace, storage provides long-term retention, and power plus cooling keep the system stable.

## Further Reading

- [Watch Video](https://learn.kodekloud.com/user/courses/computer-architecture/module/1fdcb0d1-7a06-4ea5-8a79-5598e809a097/lesson/cb5d281e-8b98-440a-ae22-2a01b7336c98)
