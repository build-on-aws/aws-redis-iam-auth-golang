# Authenticate Go application to Amazon MemoryDB (and Amazon ElastiCache) for Redis using AWS IAM

Amazon MemoryDB has supported username/password based authentication [using Access Control Lists](https://docs.aws.amazon.com/memorydb/latest/devguide/components.html#whatis.components.acls) since the very beginning. But it also supports [IAM support for authentication](https://aws.amazon.com/about-aws/whats-new/2023/05/amazon-memorydb-redis-iam-authentication/).

[MemoryDB documentation](https://docs.aws.amazon.com/memorydb/latest/devguide/auth-iam.html#auth-iam-Connecting) has an example for a Java application with the Lettuce client. The process is similar for other languages, but you still need to implement it. This repository contains an example for a Go application with the widely used [go-redis](https://github.com/redis/go-redis) client.

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This library is licensed under the MIT-0 License. See the LICENSE file.

