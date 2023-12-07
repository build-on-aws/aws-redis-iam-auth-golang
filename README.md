# Using IAM authentication for Redis on AWS

*How to securely connect your Go applications to Amazon MemoryDB (or ElastiCache) for Redis using IAM*

Amazon MemoryDB has supported username/password based authentication [using Access Control Lists](https://docs.aws.amazon.com/memorydb/latest/devguide/components.html#whatis.components.acls) since the very beginning. But it also supports [IAM for authentication](https://aws.amazon.com/about-aws/whats-new/2023/05/amazon-memorydb-redis-iam-authentication/).

![](https://community.aws/_next/image?url=https%3A%2F%2Fassets.community.aws%2Fa%2F2ZCVX81lcmA658o2P05GmRPjRCU.jpeg%3FimgSize%3D918x370&w=1920&q=75)

[MemoryDB documentation](https://docs.aws.amazon.com/memorydb/latest/devguide/auth-iam.html#auth-iam-Connecting) has an example for a Java application with the Lettuce client. The process is similar for other languages, but you still need to implement it. This repository contains an example for a Go application with the widely used [go-redis](https://github.com/redis/go-redis) client.

For a deep dive, [check out this blog post](https://community.aws/content/2ZCKrwaaaTglCCWISSaKv1d7bI3/using-iam-authentication-for-redis-on-aws).

## License

This library is licensed under the MIT-0 License. See the LICENSE file.

