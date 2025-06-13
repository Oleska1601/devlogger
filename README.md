# DevLogger - Структурированный логгер для логов разработки

Cтруктурированный логгер на основе `slog` с кастомными уровнями логирования и удобным форматом вывода.
**Вывод логов** осуществляется в консоль в формате JSON

## Формат

1. Временная метка
2. В каком файле и строчке
3. Уровень лога (debug, info, warning, error, fatal)
4. Сообщение (структура):
```json
{
  "action": "строка с описанием действия",
  "status": "try|success|failure (только для level=info)",
  "details": "дополнительные данные в произвольном формате",
  "error": "текст ошибки (для level=error|fatal)"
}
```

## Установка

```bash
go get github.com/Oleska1601/devlogger
```

### Описание уровней логирования
### Уровни логирования

| Уровень | Значение |
|---------|----------|
| debug   | -4       |
| info    | 0        |
| warn    | 4        |
| error   | 8        |
| fatal   | 12       |

### Правила формирования сообщения по уровням

| Уровень | Обязательные поля                          | Описание                                                                      |
|---------|--------------------------------------------|-------------------------------------------------------------------------------|
| debug   | action, details                            | Произвольная отладочная информация                                            |
| info    | action, status, details                    | Действие + статус выполнения ( try / success / failure ) + релевантные данные |
| warn    | action, details (expected, actual, reason) | Нестандартное, но обработанное поведение системы                              |
| error   | action, error, details                     | Критическая ошибка с контекстом возникновения                                 |
| fatal   | action, error, details                     | Фатальная ошибка, приводящая к остановке приложения                           |

## Использование

### Базовая инициализация

```go
logger := devlogger.New("info") // debug, info, warn, error, fatal
```
Если указан неизвестный уровень, по умолчанию используется значение «info».
### Примеры логирования

```go
// Debug
logger.Debug("validate user input", map[string]map[string]string{
		"input": {
			"username": "test_user",
			"email":    "test@example.com",
		},
	})
```
```go
// Info
logger.Info("validate user input", logger.InfoStatusSuccess, map[string]map[string]string{
		"user": {
			"username": "test_user",
			"role":     "CA",
			"context":  "test_ctx",
		},
	})
```
```go
// Warning
logger.Warning("refresh token", "refresh token", "re-login required", "refresh token expired")
```
```go
// Error
logger.Error("create user", err, map[string]map[string]string{
		"user": {
			"username": "test_user",
			"email":    "test@example.com",
		},
	})
```
```go
// Fatal
logger.Fatal("load configuration", err, map[string]string{"path": "/etc/app/config.yaml"})
```

## Формат вывода

Пример вывода логов:

Debug:
```json
{
  "timestamp": "2023-11-15T14:22:10.123Z",
  "file": "user_service.go:42",
  "level": "debug",
  "message": {
    "action": "validate user input",
    "details": {
      "input": {"username": "test_user", "email": "test@example.com"}
    }
  }
}
```

Info:
```json
{
  "timestamp": "2023-11-15T14:23:15.456Z",
  "file": "user_service.go:78",
  "level": "info",
  "message": {
    "action": "create user",
    "status": "success",
    "details": {
      "user": {"username": "test_user", "role": "CA", "context": "test_ctx"}
    }
  }
}
```

Warning:
```json
{
  "timestamp": "2023-11-15T14:24:20.789Z",
  "file": "auth_service.go:112",
  "level": "warning",
  "message": {
    "action": "refresh token",
    "details": {
      "expected": "refresh token",
      "actual": "re-login required",
      "reason": "refresh token expired"
    }
  }
}
```

Error:
```json
{
  "timestamp": "2023-11-15T14:25:30.012Z",
  "file": "user_service.go:85",
  "level": "error",
  "message": {
    "action": "create user",
    "error": "username already exists",
    "details": {
      "user": {"username": "test_user", "email": "test@example.com"}
    }
  }
}
```

Fatal:
```json
{
  "timestamp": "2023-11-15T14:26:40.345Z",
  "file": "main.go:25",
  "level": "fatal",
  "message": {
    "action": "load configuration",
    "error": "config file not found",
    "details": {
      "path": "/etc/app/config.yaml"
    }
  }
}
```