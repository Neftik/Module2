# Калькулятор

**Calc Go** — это сервис для выполнения математических вычислений на основе переданного выражения. Сервис предоставляет API для обработки запросов на вычисления, а также включает в себя набор тестов для проверки корректности работы.

Проект написан на языке Go, организован в модульной структуре и содержит примеры использования.

---

## Установка и запуск

Для запуска проекта выполните следующие шаги:

1. Склонируйте репозиторий:

```bash
git clone https://github.com/Andreyka-coder9192/calc_go.git
cd calc_go
```

2. Убедитесь, что Go установлен и находится в `$PATH` (проверить версию можно командой `go version`).

3. Запустите API-сервер:

```bash
go run ./cmd/main.go
```

Сервер запустится на порту `8080` по умолчанию. Если необходимо изменить порт, установите переменную окружения `PORT` перед запуском:

### В Linux и macOS

```bash
export PORT=9090
go run ./cmd/main.go
```

### В Windows (PowerShell)

```powershell
$env:PORT=9090
go run ./cmd/main.go
```

---

## Использование API

### Эндпоинт

```
POST /api/v1/calculate
```

### Заголовки

- `Content-Type: application/json`

### Тело запроса

Пример:

```json
{
  "expression": "2+2*2"
}
```

### Ответы

1. **Успешный запрос**

   **Статус-код:** `200 OK`  
   **Пример ответа:**

   ```json
   {
     "result": "6"
   }
   ```

2. **Ошибка обработки выражения**

   **Статус-код:** `422 Unprocessable Entity`  
   **Пример ответа:**

   ```json
   {
     "error": "Error calculation"
   }
   ```

3. **Неподдерживаемый метод**

   **Статус-код:** `405 Method Not Allowed`  
   **Пример ответа:**

   ```json
   {
     "error": "Wrong Method"
   }
   ```

4. **Некорректное тело запроса**

   **Статус-код:** `400 Bad Request`  
   **Пример ответа:**

   ```json
   {
     "error": "Invalid Body"
   }
   ```

---

## Примеры использования

1. **Успешный запрос**:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

Ответ:

```json
{
  "result": "6"
}
```

2. **Ошибка: некорректное выражение**:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2*(2+2{)"
}'
```

Ответ:

```json
{
  "error": "Error calculation"
}
```

3. **Ошибка: неверный метод**:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json'
```

Ответ:

```json
{
  "error": "Wrong Method"
}
