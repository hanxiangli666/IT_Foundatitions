> ## Documentation Index
>
> Fetch the complete documentation index at: https://notes.kodekloud.com/llms.txt
> Use this file to discover all available pages before exploring further.

# What Happens in the Cloud

> Explains how compute, storage, databases, and networking work together to process, store, index, and deliver media for scalable cloud video applications.

Yahoo’s new Remix feature just went viral. Millions of users upload clips, edit videos in the browser, and watch content around the world simultaneously — with no buffering and no late-night calls to IT. How does that scale?

The secret isn’t a single super-server. It’s a whole factory of cloud services working together: scaling compute, storing durable media, organizing metadata in databases, and delivering content over optimized networks. This article follows the typical steps a Remix video takes — from upload to global playback — and explains the core cloud building blocks that make it possible.

![1774763577337](image/WhatHappensintheCloud/1774763577337.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-upload-process-store-stream.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=a09437c1ff1b4564fa495daf0cd3f0b1" alt="A presenter wearing a KodeKloud t-shirt stands on the right against a dark background. To the left are four colorful buttons labeled "Upload," "Process," "Store," and "Stream."" width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-upload-process-store-stream.jpg" />
</Frame>

Cloud dashboards can look like a jungle of services,

![1774763581829](image/WhatHappensintheCloud/1774763581829.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-aws-console-access-denied.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=d210c4a0dcb156b02780402d4d11faf2" alt="A person stands on the right wearing a black KodeKloud t-shirt, appearing to present. On the left is a large framed screen showing an AWS Console dashboard with "Access denied" messages." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-aws-console-access-denied.jpg" />
</Frame>

but most systems are built from four repeatable building blocks: Compute, Storage, Databases, and Networking. These foundations power modern platforms and keep front-end experiences smooth and responsive.

After reading this article, you’ll be able to describe how compute, storage, databases, and networking work together to run, store, manage, and deliver applications at scale.

## 1. Compute — the brain that runs the work

When Remix launched, millions began editing, cropping, and rendering videos in the browser. That processing doesn’t all run on each user’s device — cloud compute services perform the heavy lifting.

Cloud compute executes tasks, processes data, and makes fast decisions. In practice, this means launching virtual machines (VMs) or containers to run code.

![1774763588402](image/WhatHappensintheCloud/1774763588402.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-purple-cloud-infographic.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=4cb12c9736488577873877459102e4a6" alt="A presenter wearing a KodeKloud T‑shirt stands on the right beside a stylized purple cloud infographic. The graphic shows gears, a VM cube and a container box with labels reading "Carry Task", "Crunch Numbers", and "Fast Decisions."" width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-purple-cloud-infographic.jpg" />
</Frame>

* Virtual machines (VMs) simulate a full computer with their own OS. Use VMs for strong isolation, custom OS-level configuration, or long-running tasks like large video transcodes or scheduled batch jobs.
* Containers package an app and its dependencies while sharing the host kernel. They start quickly and scale horizontally, making them ideal for short-lived micro-tasks: trimming clips, applying filters, or handling uploads.

MiaoTube (the Remix backend example) typically uses containers for interactive editing and fast autoscaling. For heavier workloads (long transcoding runs or large batch jobs), VMs can be used in the background so the UI stays responsive even under peak load.

## 2. Storage — the cloud’s filing cabinet

Once a user finishes editing, the results must be saved safely and durably. Cloud storage provides that persistence.

![1774763596278](image/WhatHappensintheCloud/1774763596278.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-cloud-storage-icon.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=796af12fa2e5bc83f3b494a8d16161f0" alt="A presenter wearing a KodeKloud t-shirt stands against a black background. To his left is a purple stylized drawer icon with a gradient file and the label "Cloud Storage."" width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-cloud-storage-icon.jpg" />
</Frame>

Although we say “the cloud,” storage is still physical disks and SSDs in data centers. The cloud abstracts that hardware behind APIs, letting you focus on files and objects instead of rack-level maintenance.

Common cloud storage types:

| Storage Type   | Characteristics                                                 | Best for                                                  |
| -------------- | --------------------------------------------------------------- | --------------------------------------------------------- |
| Block storage  | Fixed-size blocks, low latency, presented as volumes            | VM disks, active databases                                |
| File storage   | Hierarchical filesystem with folders and paths                  | Shared workloads, legacy apps expecting POSIX filesystems |
| Object storage | Flat namespace, objects with metadata and keys, highly scalable | Large media files, thumbnails, static assets, archives    |

Example file path (file storage):

```text
/projects/remix/video1.mov
```

In object storage, the media is accessed by a key and metadata, not a nested path. Object stores are the common choice for media-heavy apps because they scale cheaply and integrate well with CDNs and metadata-based searches.

<Callout icon="lightbulb" color="#1CB2FE">
  Object storage is optimized for durability and scale. Use block storage for fast random I/O and file storage when an application expects a filesystem interface.
</Callout>

## 3. Databases — the cloud’s working memory

Every time a user likes a video, saves a draft, or edits a remix, those actions must be recorded and available across devices. Databases provide the structured and unstructured data stores that keep the app consistent and searchable.

![1774763604060](image/WhatHappensintheCloud/1774763604060.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/inventory-list-db-structured-unstructured.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=bd4c6536ba60a46a45f1ad944edec2c3" alt="A purple "Inventory List" graphic with a stacked database icon and labeled boxes for "Structured" (e.g., Users, Payments) and "Unstructured" (e.g., Tags, Comments, View History). A presenter stands on the right, gesturing toward the visual." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/inventory-list-db-structured-unstructured.jpg" />
</Frame>

Relational (SQL) databases are great for structured data, strong consistency, and complex queries — for example, user accounts and payments. NoSQL stores are designed for flexible schemas, massive scale, and fast writes — useful for feeds, metadata, session data, and edit histories.

![1774763609728](image/WhatHappensintheCloud/1774763609728.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/relational-vs-nosql-kodekloud-presenter.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=664910d481d0cd9f87f92bd5ab4af2d4" alt="A slide comparing "Relational" databases (example Users and Videos tables) on the left with various "NoSQL" store types listed on the right. A presenter stands to the right wearing a black KodeKloud t-shirt." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/relational-vs-nosql-kodekloud-presenter.jpg" />
</Frame>

In modern cloud architectures, databases are commonly provided as managed services. That means the provider handles provisioning, scaling, replication, monitoring, and backups so developers can focus on schema design and queries.

![1774763614843](image/WhatHappensintheCloud/1774763614843.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-database-comparison-slide.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=9f0bc3fb78dc88b3e5eb2a94c4f3d259" alt="A presenter wearing a KodeKloud T-shirt stands to the right of a slide. The slide compares "Relational Database" (user details, access control) and "Non-Relational Database" (metadata, personalization, edit history) with accompanying icons." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-database-comparison-slide.jpg" />
</Frame>

MiaoTube uses both: a relational DB for user accounts and access control, and a flexible NoSQL store for Remix metadata, personalization, and edit history. Managed databases reduce operational burden and make it easier to scale safely.

## 4. Networking — routing, balancing, and delivering content

After processing and storage, media must be delivered worldwide. Cloud networking is a system of smart highways: routing requests, balancing load, and delivering cached content from the closest edge.

Think about these four networking responsibilities:

* Availability: Multiple regions and availability zones provide geographic distribution and fault tolerance.
* Routing: Traffic is directed to healthy, nearby, and lightly loaded servers for optimal latency.
* Load balancing: Distributes incoming requests across server instances and reroutes traffic when instances fail.
* Delivery: CDNs cache popular assets at edge locations to serve content quickly to users worldwide.

  ![1774763623637](image/WhatHappensintheCloud/1774763623637.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/multi-cloud-load-balancing-diagram.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=416355c8bcf9125b49de43633eb612bf" alt="A presenter stands beside an illustrated diagram of load balancers routing traffic to three stacks of servers, with user avatars above each stack. Cloud provider logos for AWS, Azure and Google Cloud appear in a purple cloud at the top." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/multi-cloud-load-balancing-diagram.jpg" />
</Frame>

CDNs (Content Delivery Networks) are particularly important for media-heavy platforms: they cache videos, thumbnails, and static assets closer to users to reduce startup latency and improve streaming quality.

![1774763628369](image/WhatHappensintheCloud/1774763628369.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/cdn-purple-globe-cloud-logos-kodekloud.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=455aa333dec2c6de94e281ea8dcc2d13" alt="A purple stylized globe with glowing location pins and the text "Content Delivery Networks" sits on the left, with AWS, Azure, and Google Cloud logos above. On the right, a person wearing a KodeKloud t‑shirt gestures against a black background." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/cdn-purple-globe-cloud-logos-kodekloud.jpg" />
</Frame>

Bringing the pieces together:

* Compute runs Remix at scale (containers for quick tasks, VMs for heavy jobs).
* Storage keeps files safe and accessible (block, file, and object options).
* Databases manage metadata, user preferences, and transactional records.
* Networking connects users and accelerates delivery with load balancers and CDNs.

  ![1774763633013](image/WhatHappensintheCloud/1774763633013.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-cloud-computing-graphics.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=589e9d0e41804926dfa0fd934b4bb290" alt="A person wearing a KodeKloud t-shirt stands on the right, gesturing as if explaining something. On the left are purple cloud-computing graphics including a server drawer, stacked disks, a globe-style network, and UI icons." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-cloud-computing-graphics.jpg" />
</Frame>

## Quick quiz

Which of the following statements is true?

A. Compute services provide virtual machines and containers to store application data.
B. Storage services run your application logic and automatically scale based on demand.
C. Database services help organize and query structured data, like video titles and tags.
D. Networking services are only used to connect internal cloud components, not end users.

![1774763638956](image/WhatHappensintheCloud/1774763638956.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-cloud-quiz-which-is-true.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=b9577762569156ce6dd02cb2a67ff82d" alt="A quiz slide titled "Which of the following is TRUE?" lists four multiple-choice statements about cloud services (compute, storage, database, networking). A presenter wearing a KodeKloud t-shirt stands to the right against a black background." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-cloud-quiz-which-is-true.jpg" />
</Frame>

Pause now.

Welcome back. The correct answer is C.

* Why C is true: Databases are designed to store and query structured data such as video titles, tags, and user records.
* Why the others are false:
  * A is false: Compute runs code (VMs and containers); storage is used for persistent data.
  * B is false: Storage persists data; compute handles running logic and scaling.
  * D is false: Networking handles both internal and external traffic, including user-facing delivery.

    ![1774763646707](image/WhatHappensintheCloud/1774763646707.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/k9suVL7cFPUhhb5j/images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-four-building-blocks.jpg?fit=max&auto=format&n=k9suVL7cFPUhhb5j&q=85&s=11d8574ba51d240e6cb9f46bafe4b25a" alt="A presenter stands on the right wearing a black "KodeKloud" T-shirt, speaking against a black background. To the left is a purple slide titled "Four Building Blocks" listing: Compute, Storage, Databases, and Networking." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/What-Happens-in-the-Cloud/What-Happens-in-the-Cloud/kodekloud-presenter-four-building-blocks.jpg" />
</Frame>

Storage systems persist files, backups, and media (block for low-latency disk, file for compatibility, object for scale). Databases manage structured and unstructured data with managed services that scale on demand. Networking connects everything, routes traffic across regions, and uses CDNs to accelerate global delivery.

<Callout icon="warning" color="#FF6B6B">
  Cloud architecture also needs security controls and cost management. We’ll cover cloud security and cost control in a dedicated follow-up article.
</Callout>

## Links and references

* [Kubernetes Basics](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/) — orchestration for containers.
* [AWS Compute Services](https://aws.amazon.com/compute/) — VM and container offerings.
* [Cloud Storage Concepts](https://cloud.google.com/storage/docs) — object, file, and block storage patterns.
* [Content Delivery Networks (CDNs)](https://en.wikipedia.org/wiki/Content_delivery_network) — how CDNs improve delivery.

You’ve now seen the end-to-end flow: upload → compute → store → index in a database → deliver via networking. These four building blocks — Compute, Storage, Databases, and Networking — form the backbone of modern cloud-native applications.

<CardGroup>
  <Card title="Watch Video" icon="video" cta="Learn more" href="https://learn.kodekloud.com/user/courses/cloud-computing-fundamentals/module/f032d7b9-edae-4154-a304-efac660806f1/lesson/f992fb27-fe20-48cf-9643-f106e4af5272" />
</CardGroup>
