# bookstore_oauth-api

Oauth API

```bash
cqlsh -u cassandra -p cassandra
```

```cqlsh
CREATE KEYSPACE oauth WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 1
};
```

```cqlsh
DESCRIBE KEYSPACES;
```

```cqlsh
USE oauth;
```

```cqlsh
DESCRIBE TABLES;
```

```cqlsh
CREATE TABLE access_tokens (
    access_token varchar PRIMARY KEY,
    user_id bigint,
    client_id bigint,
    expires bigint
);
```

```cqlsh
SELECT * FROM access_tokens;
```
