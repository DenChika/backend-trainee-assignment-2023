# Сервис динамического сегментирования пользователей

## Проблема

В Авито часто проводятся различные эксперименты — тесты новых продуктов, тесты интерфейса, скидочные и многие другие.
На архитектурном комитете приняли решение централизовать работу с проводимыми экспериментами и вынести этот функционал в отдельный сервис.

## Задача

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

## Запуск приложения

```
make run
```
Если приложение запускается впервые, то необходимо применить миграции к базе данных:
```
make migrate-up
```

## Описание методов

### Регистрация пользователя

Пример запроса:
```
curl -X POST -H "Content-Type: application/json" 
    -d '{"username": "denchika", "password": "abobus123"}' 
    http://localhost:8080/auth/sign-up
```

Пример ответа:
```json
{
  "status": 200
}
```

### Аутентификация пользователя

Пример запроса:
```
curl -X GET -H "Content-Type: application/json" 
    -d '{"username": "denchika", "password": "abobus123"}' 
    http://localhost:8080/auth/sign-in
```

Пример ответа:
```json
{
  "Token": "Your access-token"
}
```

### Добавление сегмента

Пример запроса:
```
curl -X POST -H "Content-Type: application/json"
    -H "Authorization: Bearer `Your token`
    -d '{"slug": "ABOBA"}'
    http://localhost:8080/segment/
```

Пример ответа:
```json
{
  "Id": 1
}
```

### Удаление сегмента

Пример запроса:
```
curl -X DELETE -H "Content-Type: application/json"
    -H "Authorization: Bearer `Your token`
    -d '{"slug": "ABOBA"}'
    http://localhost:8080/segment/
```

Пример ответа:
```json
{
  "status": 200
}
```

### Добавление и удаление сегментов для пользователя

Пример запроса:
```
curl -X POST -H "Content-Type: application/json"
    -H "Authorization: Bearer `Your token`
    -d '{"slugs-to-add": ["ABOBA"], "slugs-to-remove": []}'
    http://localhost:8080/users-segments/
```

Пример ответа:
```json
{
  "slugs-that-have-been-added": [
    "ABOBA"
  ],
  "slugs-that-have-been-removed": null
}
```

### Получение сегментов пользователя

Пример запроса:
```
curl -X GET -H "Content-Type: application/json"
    -H "Authorization: Bearer `Your token`
    http://localhost:8080/users-segments/
```

Пример ответа:
```json
{
  "slugs": ["ABOBA"]
}
```

## Swagger

Swagger файл описан в директории [docs](docs)

При работающем приложении можно посмотреть [здесь](http://localhost:8080/swagger/index.html)

