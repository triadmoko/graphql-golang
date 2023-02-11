# Database Migration

## Configs
```sql
SET GLOBAL sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));
```

## Create Migration Database MYSQL
script create database 
```bash
make migration-create table=name
```
 
script migration up
```bash
make migration-create table=name
```

script migration down
```bash
make migration-down  table=name
```
