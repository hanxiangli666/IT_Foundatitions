# Demo Creating S3 Bucket on AWS

> Guide to creating and managing an AWS S3 bucket, covering security settings, storage classes, versioning, lifecycle rules, uploading objects, and best practices for cost and access control

In this lesson you’ll create an Amazon S3 bucket and review the key settings used to store files in the cloud — for example, daily database backups or archived logs. Amazon S3 (Simple Storage Service) is AWS’s primary object storage service and is widely used for scalable, durable, and cost-effective storage of objects (files).

A bucket is the top-level container for objects in S3. It behaves like a cloud folder but includes built-in features for security, redundancy, lifecycle management, and storage-class controls. For most mixed-use workloads, a general-purpose bucket with sensible defaults is appropriate.

Note: bucket names must be globally unique across all AWS accounts. If your chosen name is already taken, you will be prompted to pick a different one.

![1774969402294](image/DemoCreatingS3BucketonAWS/1774969402294.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/XqMRTrEG2GqQdgEv/images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/s3-create-kk-demo-bucket-kodekloud.jpg?fit=max&auto=format&n=XqMRTrEG2GqQdgEv&q=85&s=60167f660d81260dadea598acbab62cb" alt="An AWS S3 "Create bucket" console screenshot showing a highlighted bucket name ("kk-demo-bucket") and object ownership/public access settings. A presenter wearing a KodeKloud t-shirt stands to the right, pointing at the screen." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/s3-create-kk-demo-bucket-kodekloud.jpg" />
</Frame>

## Creating a bucket — key settings to review

When creating a bucket, pay attention to these controls and their recommended use:

* Object ownership: Set to “bucket owner enforced” to disable ACLs and ensure objects uploaded by other accounts are owned by your bucket account. This centralizes access control to IAM and bucket policies.
* Block Public Access: Enabled by default. Keeps objects from being publicly accessible unless you explicitly allow it. Recommended for secure-by-default posture.
* Versioning: When enabled, S3 preserves all versions of every object in a bucket. Great for backups and recovering from accidental deletes, but increases storage usage and cost.
* Tags: Key/value pairs for organizing resources. Useful for billing, automation, and filtering.

Table — Quick guidance for bucket settings

| Setting             | Purpose                                  | Recommendation                                                                      |
| ------------------- | ---------------------------------------- | ----------------------------------------------------------------------------------- |
| Object ownership    | Control who owns uploaded objects        | Set to `bucket owner enforced` for centralized control                            |
| Block Public Access | Prevent unintended public access         | Keep enabled unless you have a specific, reviewed use case                          |
| Versioning          | Retain object history and recoverability | Enable if you need point-in-time recovery or protection against accidental deletion |
| Tags                | Organize and track costs                 | Add billing or environment tags (e.g.,`env:prod`, `team:analytics`)             |

## Uploading objects and storage classes

After creating the bucket you can upload objects. Each object has metadata such as content type, last-modified timestamp, size, and a storage class that determines cost and retrieval behavior. You can select a storage class during upload or use lifecycle rules to migrate objects automatically.

Before uploading, verify the Permissions tab to confirm ownership and public access settings match your security requirements. In the Properties tab you can view or set storage class options during upload.

The S3 console includes a pricing table and descriptions for each storage class so you can compare performance, retrieval latency, and cost by region.

![1774969414895](image/DemoCreatingS3BucketonAWS/1774969414895.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/XqMRTrEG2GqQdgEv/images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/aws-s3-pricing-kodekloud-presenter.jpg?fit=max&auto=format&n=XqMRTrEG2GqQdgEv&q=85&s=f64b1261421466c77c11d10e730893c2" alt="A screenshot of an AWS S3 pricing page displayed on the left side of the image. On the right, a person wearing a black KodeKloud t-shirt stands and gestures as if explaining the content." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/aws-s3-pricing-kodekloud-presenter.jpg" />
</Frame>

If you leave objects in the Standard storage class, you get general-purpose, frequently accessed storage. Because the bucket has public access blocked in our example, trying to open an object URL as an anonymous user returns an access denied error. Use the console’s Open button, pre-signed URLs, or proper IAM permissions to view or download objects.

### Storage class summary

| Storage class              | Typical use case                                 | Retrieval / cost characteristics                 |
| -------------------------- | ------------------------------------------------ | ------------------------------------------------ |
| Standard                   | Frequently accessed data                         | Low latency, higher cost                         |
| Standard-IA / One Zone-IA  | Infrequently accessed but quick retrieval needed | Lower storage cost, retrieval fees apply         |
| Intelligent-Tiering        | Auto-optimizes cost for unknown access patterns  | Automated tiering, small monitoring fee          |
| Glacier Instant Retrieval  | Archive with immediate access                    | Low-cost archival, 90-day minimum storage charge |
| Glacier Flexible Retrieval | Archive with configurable retrieval times        | Very low storage cost, retrieval times vary      |
| Glacier Deep Archive       | Long-term archival                               | Lowest storage cost, retrieval may take hours    |

## Lifecycle rules — automate cost optimization

Lifecycle rules let you automate transitions between storage classes or expire objects after a certain time. From the Management tab, click Create lifecycle rule and give it a name. You can then specify filters (prefix, tags, size) and actions (transition current versions, transition noncurrent versions, expire objects).

Example: transition current object versions to a cheaper storage class after 30 days. If you don’t use versioning, select options that apply only to current versions.

Important lifecycle considerations:

* Transitioning objects may incur transition costs.
* Some storage classes have minimum storage duration charges (for example, Glacier Instant Retrieval typically has a 90-day minimum).
* After a rule is active, the console will show the storage class change for objects once conditions are met.

  ![1774969464479](image/DemoCreatingS3BucketonAWS/1774969464479.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/XqMRTrEG2GqQdgEv/images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/s3-create-lifecycle-rule-kodekloud-presenter.jpg?fit=max&auto=format&n=XqMRTrEG2GqQdgEv&q=85&s=a62bf0c0a8ed0c6555a3bfcde334a8b0" alt="A screenshot of the Amazon S3 console’s "Create lifecycle rule" page is shown on the left. On the right, a presenter wearing a black KodeKloud t-shirt is standing and gesturing." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/s3-create-lifecycle-rule-kodekloud-presenter.jpg" />
</Frame>

After saving a lifecycle rule the console presents a summary and success banner indicating your rule is configured.

![1774969475719](image/DemoCreatingS3BucketonAWS/1774969475719.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/XqMRTrEG2GqQdgEv/images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/kodekloud-presenter-s3-lifecycle-success.jpg?fit=max&auto=format&n=XqMRTrEG2GqQdgEv&q=85&s=7acab49fe13f0be97af60b34be1b9f9d" alt="A presenter wearing a black "KodeKloud" t-shirt stands to the right of a large screenshot of the Amazon S3 console showing a "Lifecycle configuration" page and a green success banner. The screen lists a "30-day" lifecycle rule and options to view, edit, delete, or create rules." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/kodekloud-presenter-s3-lifecycle-success.jpg" />
</Frame>

## Deleting objects and buckets

To delete a file manually, use the object actions menu in the console. Remember: to delete a bucket you must first remove all objects and all object versions if versioning is enabled. AWS will not allow deletion of a non-empty bucket.

<Callout icon="lightbulb" color="#1CB2FE">
  Use lifecycle rules to automate retention and cost management. They can transition objects to cheaper tiers or expire objects when they’re no longer needed.
</Callout>

## Pop quiz — best practice

You upload a file to an S3 bucket. Which step ensures resources are created, used, and cleaned up correctly?

A. Upload the file, make it public, and forget about the bucket. It will auto-delete.
B. Use S3 lifecycle rules to archive or delete objects automatically when no longer needed.
C. Store all your files in the root bucket with the same name, and never delete anything.
D. Delete the bucket before removing the objects to save time.

Pause now. Welcome back — the correct answer is B.

Why:

* B is correct: lifecycle rules automate archival or deletion and help control costs.
* A is dangerous: objects and buckets do not auto-delete; leaving public objects or orphaned data creates cost and security risks.
* C is poor practice: poor naming and never deleting data lead to costly and unmanageable storage.
* D won’t work: AWS requires a bucket to be empty before it can be deleted.

## Recap and next steps

* Core cloud services across providers include Compute, Storage, Databases, and Networking. Names differ (EC2, Blob Storage, Cloud SQL) but functions are comparable.
* In this lesson you created an S3 bucket, uploaded objects, reviewed security and ownership settings, inspected storage classes and pricing, and configured lifecycle rules to automate cost control and retention.
* Use versioning, appropriate access controls (IAM and bucket policies), and lifecycle rules together for secure, scalable, and cost-effective storage.

  ![1774969491293](image/DemoCreatingS3BucketonAWS/1774969491293.png)

<Frame>
    <img src="https://mintcdn.com/kodekloud-c4ac6d9a/XqMRTrEG2GqQdgEv/images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/kodekloud-presenter-demos-gcp-s3-security.jpg?fit=max&auto=format&n=XqMRTrEG2GqQdgEv&q=85&s=72ddb6fcaf0df70672b9253a7acdee25" alt="A presenter stands on the right wearing a black KodeKloud t-shirt and gesturing while speaking. On the left is a purple slide titled "Demos" listing: launching a VM in GCP; uploading files to AWS S3; and security settings, storage classes, and lifecycle rules." width="1920" height="1080" data-path="images/Cloud-Computing-Fundamentals/Cloud-Providers/Demo-Creating-S3-Bucket-on-AWS/kodekloud-presenter-demos-gcp-s3-security.jpg" />
</Frame>

Further reading and references:

* [AWS S3 documentation](https://docs.aws.amazon.com/s3/index.html)
* AWS S3 pricing and storage class details (linked from the S3 console)
* [AWS Cloud Practitioner (CLF-C02) course](https://learn.kodekloud.com/user/courses/aws-cloud-practitioner-clf-c02)

<CardGroup>
  <Card title="Watch Video" icon="video" cta="Learn more" href="https://learn.kodekloud.com/user/courses/cloud-computing-fundamentals/module/db2b85ff-442f-481d-a308-21c5eb63344b/lesson/294b3b7a-9321-413e-a2dd-15ae869668b9" />

  `<Card title="Practice Lab" icon="flask-conical" cta="Learn more" href="https://learn.kodekloud.com/user/courses/cloud-computing-fundamentals/module/db2b85ff-442f-481d-a308-21c5eb63344b/lesson/53ff9d21-22af-4ddb-9c7f-fb6504ef0ee1" />`
`</CardGroup>`

Built with [Mintlify](https://mintlify.com).
