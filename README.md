# Тестовое задание для golang dev в VK

## 1. Работа c API
*Примечание: Авторизация в swagger через JWT token состоит из 2-х ключевых слов (example: Bearer <your_jwt_token_here>)*
### 1.1 Выбор запуска
#### 1.1.1 Запуск через Docker compose

- ```make build```: start docker containers
- ```make swag_ui```: swagger API

#### 1.1.2 Локальный запуск

- ```make postgresinit```: start postgres in docker container
- ```make server```: server
- ```make swag_ui```: swagger API

### 2. Тесты

- ```make test```: start linter and mock tests