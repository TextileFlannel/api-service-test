# Q&A API Service
REST API сервис для управления вопросами и ответами. Построен на Go с использованием чистой архитектуры, PostgreSQL и Docker.

## Структура проекта

```
.
├── cmd/server/          # Точка входа приложения
├── internal/
│   ├── config/          # Конфигурация приложения
│   ├── handlers/        # HTTP обработчики
│   ├── models/          # Модели данных
│   ├── repository/      # Работа с БД
│   └── service/         # Бизнес-логика
├── migrations/          # SQL миграции
│   ├── migrations.go    # Запуск миграций из кода
│   └── *.sql            # Файлы миграций
├── .env                 
├── docker-compose.yml   
├── Dockerfile           
├── Makefile             
└── go.mod               
```

## Запуск приложения

```bash
make docker-up
```

## Запуск тестов

```bash
make tests
```

## Endpoints

- `POST /questions` - Создать вопрос

Request:
```json
{
  "text": "question"
}
```

Response (201 Created):
```json
{
  "id": 1,
  "text": "question",
  "created_at": "2025-01-15T10:30:00Z"
}
```

- `GET /questions` - Список всех вопросов

Response (200 OK):
```json
[
  {
    "id": 1,
    "text": "question",
    "created_at": "2025-01-15T10:30:00Z"
  }
]
```
- `GET /questions/{id}` - Получить вопрос и все ответы на него

Response (200 OK):
```json
{
  "id": 1,
  "text": "question",
  "created_at": "2025-01-15T10:30:00Z",
  "answers": [
    {
      "id": 1,
      "question_id": 1,
      "user_id": "123e4567-e89b-12d3-a456-426614174000",
      "text": "answer",
      "created_at": "2025-01-15T10:35:00Z"
    }
  ]
}
```

- `DELETE /questions` - Удалить вопрос (вместе с ответами)

Response (204 No Content)

- `POST /questions/{id}/answers` - Добавить ответ к вопросу

Request:
```json
{
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "text": "answer"                           
}
```

Response (201 Created)
```json
{
  "id": 1,                                            
  "question_id": 1,                                   
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "text": "answer",
  "created_at": "2024-01-15T10:35:00Z"                
}
```

- `GET /answers` - Получить овтет

Response (200 OK):
```json
{
  "id": 1,
  "question_id": 1,
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "text": "answer",
  "created_at": "2024-01-15T10:35:00Z"
}
```

- `DELETE /answers` - Удалить ответ

Response (204 No Content)

## Бизнес-логика
- Нельзя создать ответ к несуществующему вопросу
- Один пользователь может оставить несколько ответов на один вопрос
- При удалении вопроса каскадно удаляются все его ответы (на уровне БД)

## Стек технологий
- Go
- net/http + Chi
- GORM
- PostgreSQL
- Goose
- Docker & Docker Compose
- sqlmock