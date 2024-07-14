## Connect to postgres

```bash
psql -U postgres
```

## List databases

```psql
\l
```

## Show all databases

```sql
SELECT datname FROM pg_database;
```

## Create database

```sql
CREATE DATABASE test_db;
```

## Change database

```sql
\c test_db;
```

## Drop database

```sql
DROP DATABASE test_db;
```
