# Тестовое задание для golang dev в VK

### 1. Запуск
*Примечание: Авторизация в swagger через JWT token (example: Bearer <your_jwt_token_here>)*
#### 1.1 Запуск через Docker compose
- Склонировать репозиторий
- ```make build```: start docker containers

#### 1.2 Локальный запуск
- Склонировать репозиторий
- ```make postgresinit```: Start postgres in docker container
- ```make server```: Server
- ```make swag_ui```: Swagger API