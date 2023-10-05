# OnlineStore

## Архитектура

![screenshot1](https://github.com/kerucko/OnlineStore/blob/main/images/database.png)

Создание БД:
```
sudo -u postgres psql
create database onlinestore;
\q
sudo -u postgres psql -d onlinestore -a -f datebase.sql
```