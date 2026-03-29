# GPU Architecture

> Explains GPU architecture differences from CPUs and how cores, memory, clock speed, and specialized units affect AI training and inference performance and efficiency

Welcome back. In the previous section you rolled out laptops across your company. Now your R\&D team is building an internal chatbot (think ChatGPT trained on company data). Training that model is painfully slow on CPUs alone—it's like trying to read a library one page at a time. To accelerate training we need to run thousands of calculations in parallel, which is exactly what GPUs are designed to do.

In this lesson you'll learn:

* What architectural differences make GPUs better for certain workloads than CPUs.
* How GPU performance depends on core structure, clock speed, and memory.
* Which GPU features matter for AI training and inference.

  ![1774750264601](image/GPUArchitecture/1774750264601.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/purple-motherboard-diagram-kodekloud-presenter-cat.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=ac7e291ef5f06b381351031524396005" alt="A stylized, purple-themed diagram of a computer motherboard highlighting components like the GPU, CPU, RAM and ports. A presenter wearing a "KodeKloud" shirt stands at the right and a small cartoon cat character appears at the lower left." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/purple-motherboard-diagram-kodekloud-presenter-cat.jpg" />
</Frame>

High-level comparison: CPUs versus GPUs

* CPU: A few powerful cores optimized for complex, branching, single-threaded or moderately parallel tasks.
* GPU: Thousands of simpler cores optimized for performing the same operation over large data sets in parallel.

The GPU approach—massive parallelism—matches workloads such as rendering pixels, applying shaders across vertices, and dense matrix math in machine learning. For deep learning, GPUs can run thousands of arithmetic operations at once, reducing training time from days to hours compared to CPU-only training.

![1774750271056](image/GPUArchitecture/1774750271056.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/compare-gpus-cpus-slide-cat-presenter.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=033aa0ce555222990a56b24a6b49e15d" alt="A dark purple presentation slide titled "Compare GPUs and CPUs" with bullet headings like Architecture, Processing Model, Workloads, and a second point about how GPU performance is influenced by various factors. A cartoon cat is on the left and a presenter wearing a KodeKloud T‑shirt stands on the right." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/compare-gpus-cpus-slide-cat-presenter.jpg" />
</Frame>

Why thousands of cores?

GPUs organize their many cores into clusters—often called streaming multiprocessors (NVIDIA) or compute units (AMD). These clusters let teams of cores work in lockstep, increasing throughput and reducing coordination overhead. Imagine a factory: teams specialize in packaging, assembly, and quality checks; the coordinated teams achieve far higher throughput than the same number of lone workers.

![1774750278991](image/GPUArchitecture/1774750278991.png)

![1774750285751](image/GPUArchitecture/1774750285751.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/gpu-1000-cores-cpu-8-kodekloud.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=ad2525d3be9c4c35b06d2192dca4b6a0" alt="A stylized infographic comparing a GPU (shown with many small gradient squares and labeled "1,000+ Cores") to a CPU (eight larger gradient blocks labeled "8 Cores"). A presenter wearing a KodeKloud t-shirt stands at the right against a dark, tech-themed background." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/gpu-1000-cores-cpu-8-kodekloud.jpg" />
</Frame>

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/kodekloud-presenter-ai-training-infographic-factory.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=388b213e8b1442b34d9fdec6ca3eaa54" alt="A presenter in a black "KodeKloud" T-shirt stands at the right speaking with hands clasped. To the left is a purple-themed infographic showing a factory scene and a GPU cluster grid labeled "AI Training" with panels for Team A (Packaging), Team B (Assembling), and Team C (Quality Checking)." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/kodekloud-presenter-ai-training-infographic-factory.jpg" />
</Frame>

Specialized AI hardware

Many modern GPUs include dedicated accelerators—tensor cores, matrix cores, or vendor-specific AI units—designed for dense matrix operations and mixed-precision math (FP16, BF16). These units dramatically increase throughput for neural-network operations. A training job that takes many hours on a general-purpose GPU can finish in a fraction of that time on a GPU with specialized AI cores.

![1774750297139](image/GPUArchitecture/1774750297139.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/kodekloud-presenter-3x3-cube-specialized-cores.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=4700fd7e317883cd8097aa31bf338f11" alt="A dark tech-themed slide showing a glowing purple-pink 3x3 cube icon with buttons labeled "Tensor Core" and "Matrix Core" and a "Specialized Cores" banner. A presenter wearing a black KodeKloud T‑shirt appears at the right, speaking and gesturing." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/kodekloud-presenter-3x3-cube-specialized-cores.jpg" />
</Frame>

Memory and data movement: VRAM matters

GPU cores only run as fast as they can get data. Video RAM (VRAM) is high-bandwidth memory attached directly to the GPU and stores textures, frame buffers, or large matrices for ML. Two key properties:

* Capacity: how large a model or batch you can store on-device.
* Bandwidth: how fast data moves between VRAM and compute cores.

<Callout icon="lightbulb" color="#1CB2FE">
  VRAM capacity and bandwidth are crucial for AI workloads: capacity limits the size of models and batch sizes you can process, while bandwidth reduces stalls by keeping compute units fed. Running out of VRAM forces data transfers or smaller batches, which slow training.
</Callout>

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/gpu-vram-bandwidth-traffic-lanes-presenter.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=66cb223fac0f9e7a83a1bd3d7d470c85" alt="A presentation slide illustrating GPU and Video RAM with a traffic-lane analogy for low vs. high bandwidth on the left and a grid of GPU blocks on the right. A presenter stands at the lower-right corner explaining the diagram." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/gpu-vram-bandwidth-traffic-lanes-presenter.jpg" />
</Frame>

![1774750325339](image/GPUArchitecture/1774750325339.png)



Clock speed vs parallelism

GPUs typically run at lower per-core clock speeds than CPUs (for example, `2.5 GHz` vs `5 GHz`), but they compensate by having orders of magnitude more cores. Think of a race car (`high single-thread speed`) versus a freight train (`lots of throughput`). GPU designs prioritize total operations per second over single-thread latency.

![1774750337887](image/GPUArchitecture/1774750337887.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/gpu-cpu-clock-speed-kodekloud-presenter.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=17266d5e1144f5f58555fdb3dd6c8aef" alt="A dark presentation slide showing stylized, colorful GPU and CPU icons with callouts reading "2.5GHz 16,000 GPU Cores" and "5GHz 16 Cores" under a "Clock Speed" banner. A man wearing a KodeKloud T‑shirt stands at the right as if presenting." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/gpu-cpu-clock-speed-kodekloud-presenter.jpg" />
</Frame>

Power and efficiency

High-performance GPUs can draw hundreds of watts. Data centers running continuous AI workloads must account for power delivery and cooling. Modern GPU architectures emphasize performance-per-watt with dynamic boost clocks and power scaling to improve cost efficiency.

<Callout icon="warning" color="#FF6B6B">
  High-performance GPUs require adequate cooling and power. When planning for continuous AI training, include power consumption and data-center cooling in hardware selection and total cost calculations.
</Callout>

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/kodekloud-gpu-comparison-performance-power.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=f51d2175c2bf0f4470912edb287f64be" alt="A presenter wearing a KodeKloud t-shirt stands beside a slide. The slide compares "High-Performance GPU" and "Modern GPU" with colorful cube illustrations and notes about power, performance-per-watt, and dynamic power scaling." width="1920" height="1080" data-path="images/Computer-Architecture/Graphics-Processing-Unit-GPU/GPU-Architecture/kodekloud-gpu-comparison-performance-power.jpg" />
</Frame>

![1774750365982](image/GPUArchitecture/1774750365982.png)



Quick comparison table

| Feature           |                                                    CPU | GPU                                                         |
| ----------------- | -----------------------------------------------------: | ----------------------------------------------------------- |
| Cores             | Few, powerful cores optimized for complex control flow | Thousands of simpler cores optimized for parallel workloads |
| Processing model  |                  Serial / branching tasks, low latency | SIMD / SIMT parallelism, high throughput                    |
| Specialized units |                    Vector units, general-purpose cores | Tensor/matrix cores for AI acceleration                     |
| Memory            |                           System RAM (lower bandwidth) | High-bandwidth VRAM (higher throughput)                     |
| Clock speed       |                 Higher per-core clocks (`~3–5 GHz`) | Lower per-core clocks (`~1–3 GHz`) but many cores        |
| Energy            |  Optimized for response latency and single-thread perf | Performance-per-watt focus; can draw high total power       |

![1774750376889](image/GPUArchitecture/1774750376889.png)


Choosing the right GPU for AI

When selecting GPUs for training or inference, weigh:

* VRAM capacity: large models and big batch sizes require more memory.
* Memory bandwidth: reduces stalls and speeds up compute.
* Specialized cores: tensor/matrix cores accelerate common NN ops.
* Power and cooling: ensure infrastructure can support continuous loads.
* Price vs performance: the most expensive GPU isn’t always the best fit—consider workload characteristics (training, inference, or graphics) and cost efficiency.

References and further reading

* [NVIDIA CUDA Documentation](https://developer.nvidia.com/cuda-zone)
* [NVIDIA Tensor Cores overview](https://developer.nvidia.com/tensorcores)
* [AMD GPU Architecture](https://www.amd.com/en/technologies/graphics-architecture)
* [Deep Learning &amp; Mixed Precision Training](https://arxiv.org/abs/1710.03740)

Now that you understand GPU architecture and how it differs from CPUs, you’re better equipped to pick hardware for AI training, inference, or graphics workloads.

<CardGroup>
  <Card title="Watch Video" icon="video" cta="Learn more" href="https://learn.kodekloud.com/user/courses/computer-architecture/module/e3c31d19-97a9-464e-b94f-5ff231dc9677/lesson/4e354eef-e308-46e6-aca0-5ff0325d8743" />
</CardGroup>

Built with [Mintlify](https://mintlify.com).
