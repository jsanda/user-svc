The purpose of this project is to facilitate backup/restore testing with [medusa-operator](https://github.com/jsanda/medusa-operator). `user-svc` is a gRPC service that interacts with Cassandra on the backend. It creates a keyspace that defines a single table:

```
CREATE TABLE IF NOT EXISTS medusa_test.users (
    email text PRIMARY KEY, 
    name text
)
```
The gRPC API allows clients to insert and read the `users` table.