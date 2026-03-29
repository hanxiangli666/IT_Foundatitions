# Course Introduction

> Introduces computer architecture fundamentals, explaining CPU, GPU, memory, storage, motherboard, and peripherals and how they cooperate to produce computing results

If you remember personal computers from a few decades ago, you probably recall how magical they felt. Those early machines were impressive for their time, but they were extremely limited compared with modern computers. Today’s laptops and desktops perform tasks that once required entire rooms of equipment.

What changed? Advances across core components — the CPU, GPU, memory, storage, motherboard, and peripherals — transformed those early systems into the powerful devices we use every day. In this lesson, we’ll examine each of these building blocks and how they cooperate to convert simple inputs into complex results.

![1774749410256](image/Introduction/1774749410256.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Introduction/Course-Introduction/man-kodekloud-tshirt-purple-pc-parts.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=60e4cde07721ec3a11dbbd0772218427" alt="A man wearing a KodeKloud t-shirt stands centered against a purple gradient backdrop. Surrounding him are stylized pink illustrations of computer parts like RAM sticks, a GPU, a motherboard, a CPU and cooling fans." width="1920" height="1080" data-path="images/Computer-Architecture/Introduction/Course-Introduction/man-kodekloud-tshirt-purple-pc-parts.jpg" />
</Frame>

Hello, and welcome to the Computer Architecture course. I’m Alan — I’ll guide you through the components and concepts that make modern computing possible.

What we’ll cover in this lesson

* Peripherals: how users interact with a computer (input and output)
* CPU: the general-purpose processor that executes instructions
* GPU: the parallel processor optimized for graphics and data-parallel workloads
* Memory & Storage: fast temporary workspace (RAM) vs persistent storage (SSD/HDD)
* Motherboard: the backbone that connects everything and provides power, buses, and interfaces

We’ll begin with the devices you use to interact with a computer.

Peripherals — input and output devices
Peripherals are the external devices you use every day: keyboard, mouse, monitor, camera, microphones, and more. They bring data into the system (input) and display or transmit results back to you (output). Without peripherals, a computer would be difficult to operate — how would you type, click, or see output?

![1774749438764](image/Introduction/1774749438764.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-presenter-peripheral-devices-monitor.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=1dd61f614c34652caf3b74c4a6500a1f" alt="A presenter wearing a KodeKloud T‑shirt stands against a purple gradient slide titled "01 Peripheral Devices." To his right is a graphic of a computer monitor and keyboard showing the text "Hello, How are you doing?"" width="1920" height="1080" data-path="images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-presenter-peripheral-devices-monitor.jpg" />
</Frame>

Central Processing Unit (CPU) — the system’s general-purpose processor
The CPU executes instructions, schedules tasks, and performs general computation. While older CPUs handled many calculations at once, modern CPUs improve throughput and responsiveness through multiple cores, higher clock speeds, and better power efficiency. These improvements enable multitasking, fast context switching, and support for increasingly complex software.

![1774749446774](image/Introduction/1774749446774.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Introduction/Course-Introduction/cpu-slide-kodekloud-presenter-old-laptop.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=5c868b85cd4396ab22693fb9603794c9" alt="A slide titled "02 Central Processing Unit" with a presenter wearing a KodeKloud shirt shown on the left. On the right are purple illustrations of an "Old Laptop" (and a faint "Modern Ultrabook" graphic) with gear icons." width="1920" height="1080" data-path="images/Computer-Architecture/Introduction/Course-Introduction/cpu-slide-kodekloud-presenter-old-laptop.jpg" />
</Frame>

Graphics Processing Unit (GPU) — parallel, specialized computation
GPUs are built for massively parallel workloads. They excel at many similar, independent calculations — rendering frames, accelerating video encoding, and training or running many machine learning models. For highly parallel tasks, modern GPUs often outperform multiple older machines combined.

![1774749454036](image/Introduction/1774749454036.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-presenter-gpu-motherboard-slide.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=a6e60fdef4a477554cf48d33c9deac19" alt="A presenter wearing a "KodeKloud" T‑shirt stands in front of a large slide. The slide shows stylized GPU and motherboard graphics with the headings "03 Graphics Processing Unit" and "Parallel Processing."" width="1920" height="1080" data-path="images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-presenter-gpu-motherboard-slide.jpg" />
</Frame>

Memory and Storage — workspace versus long-term data
RAM (random access memory) is the fast, temporary workspace where running applications and active data reside. Storage (SSDs or HDDs) holds persistent data such as the operating system, applications, documents, photos, and videos. Because SSDs provide much lower latency and higher throughput than traditional HDDs, they significantly improve boot times, application launch times, and overall system responsiveness.

![1774749462598](image/Introduction/1774749462598.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-memory-storage-ram-ssd-hdd.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=400222be236d2b7d12a18f25c0fbc5d3" alt="A presenter wearing a "KodeKloud" t‑shirt stands beside a slide titled "04 Memory & Storage" with stylized RAM, SSD, and HDD icons on a dark purple background. Music and video icons are shown on the right." width="1920" height="1080" data-path="images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-memory-storage-ram-ssd-hdd.jpg" />
</Frame>

Motherboard — the system backbone
The motherboard ties everything together. It routes power and data between the CPU, GPU, memory, storage, and peripherals. It also contains controllers, buses, slots, and interfaces that determine expandability and performance characteristics. From vintage laptops to modern ultrabooks, the motherboard’s design defines the platform’s capabilities and limitations.

![1774749468606](image/Introduction/1774749468606.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-presenter-motherboard-binary-slide.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=75202ea9a524c3ec04cc90732222e7ca" alt="A presenter wearing a KodeKloud t-shirt stands in front of a purple, binary-themed slide titled "Motherboard." The background shows stylized icons of an old laptop, a modern ultrabook and a motherboard diagram with tech-related symbols." width="1920" height="1080" data-path="images/Computer-Architecture/Introduction/Course-Introduction/kodekloud-presenter-motherboard-binary-slide.jpg" />
</Frame>

Quick reference table

| Component         | Primary role                                            | Why it matters                                                                 |
| ----------------- | ------------------------------------------------------- | ------------------------------------------------------------------------------ |
| CPU               | Execute instructions, run OS and applications           | Determines single-thread performance, multitasking, and instruction throughput |
| GPU               | Parallel numeric processing (graphics, ML, encoding)    | Accelerates workloads with many similar operations                             |
| RAM               | Fast temporary workspace for active data                | Affects application responsiveness and multitasking capacity                   |
| Storage (SSD/HDD) | Persistent data storage                                 | Affects boot times, load times, and data throughput                            |
| Motherboard       | Interconnects components; provides power and interfaces | Defines expandability, bus speeds, and compatibility                           |
| Peripherals       | Input and output devices                                | Enable user interaction and data capture/display                               |

Community and learning tips

<Callout icon="lightbulb" color="#1CB2FE">
  Join a learning community to accelerate your progress. Ask questions, share projects, and discuss trade-offs — connecting with other learners (and instructors) is one of the fastest ways to deepen your understanding of computer architecture.
</Callout>

Safety and handling reminder

<Callout icon="warning" color="#FF6B6B">
  If you plan to open or handle internal components, take ESD (electrostatic discharge) precautions and power down the device. Improper handling can damage delicate hardware.
</Callout>

Whether you’re studying a museum piece or a modern ultrabook, the same core components are at work. Understanding how they interact clarifies the design trade-offs that lead to different form factors, performance profiles, and price points.

Ready to dig deeper? Let’s get started and explore how each component works internally and how they communicate to produce the computing experience you rely on every day.

Further reading and references

* Microsoft Mechanics: [https://www.microsoft.com/en-us/mechanics](https://www.microsoft.com/en-us/mechanics)

<CardGroup>
  <Card title="Watch Video" icon="video" cta="Learn more" href="https://learn.kodekloud.com/user/courses/computer-architecture/module/960b647b-9129-4f10-ac03-7f598a5ca09f/lesson/10a5b425-c045-4ed2-b1cb-8ddcf45c2530" />
</CardGroup>

Built with [Mintlify](https://mintlify.com).
