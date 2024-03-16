# Тестовое задание для golang dev в VK

### 1. Запуск
#### 1.1 Локальный запуск
- Склонировать репозиторий
- ```make postgresinit```: Start postgres in docker container
- ```make server```: Server
- ```make swag_ui```: Swagger API

#### 1.2 Docker compose
- Склонировать репозиторий
- ```make build```: start docker containers

### 1. Примеры запросов для каждого endpoint

- POST localhost:8000/auth/sign-up
```
{
    "username": "useruser",
    "password": "0000",
    "role": "admin"
}
```

- POST localhost:8000/auth/sign-in
```
{
    "username": "useruser",
    "password": "0000"
}
```