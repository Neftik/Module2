# Calculator

## What is calculator

> [!IMPORTANT]
>
> It is an HTTP API, designed to process mathematical expressions using numbers and operators like "+", "-", "*", and "/". It also supports brackets "(" and ")".

### Calculator

| Name                         | Symbol       | Supported | Error when                                | Description                                                                |
| ---------------------------- | ------------ | --------- | ---------------------------------------- | -------------------------------------------------------------------------- |
| [Number](https://en.wikipedia.org/wiki/Rational_number) | `float64`   | ☑         | -                                        | A number (e.g., `1`, `3.14`). Nothing special.                             |
| **Multiply**                 | `*`          | ☑         | -                                        | Has higher priority than addition and subtraction, but lower than brackets.|
| **Division**                 | `/`          | ☑         | Division by zero is not allowed          | Has higher priority than addition and subtraction, but lower than brackets.|
| **Adding**                   | `+`          | ☑         | -                                        | Has lower priority (same as subtraction).                                 |
| **Subtraction**              | `-`          | ☑         | -                                        | Has lower priority (same as addition).                                    |
| **Brackets**                 | `(`, `)`     | ☑         | Bracket not opened/closed correctly       | Brackets have the highest priority, and note: `10(1+1)` is `102`.         |
| **Other characters**         | Any          | ☒         | Can't convert to float                   | Avoid using unsupported characters.                                       |

### HTTP

| Name             | Response status | Method | Path                    | Body                               |
| ---------------- | --------------- | ------ | ----------------------- | ---------------------------------- |
| **OK**           | 200             | POST   | `/api/v1/calculate`     | `{"expression":"2+2"}`            |
| **Wrong Method** | 405             | GET    | `/api/v1/calculate`     | `{"expression":"2+2"}`            |
| **Wrong Path**   | 404             | POST   | `/any/unsupported/path` | `{"expression":"2+2"}`            |
| **Invalid Body** | 400             | POST   | `/api/v1/calculate`     | `invalid body`                    |
| **Error calculation** | 422         | POST   | `/api/v1/calculate`     | `{"expression":"2*(2+2"}"`       |

## How requests are sent?

> [!NOTE]
>
> The request has a specific structure in the [input](internal/http/input/input.go):

type Request struct {
   Expression string `json:"expression"`
}
The expression is the mathematical expression you want to evaluate.

In JSON, it looks like this:

json
{
   "expression": "2+2"
}
[!TIP]

You can find a variety of examples with valid expressions in the calculator tests (TestCalc), and invalid expressions in TestCalcErrors.

How will the server respond?
[!NOTE]

OK
For successful responses, we use a structure in responses:

type ResponseOK struct {
   Result float64 `json:"result"`
}
The result is the value obtained after calculating the expression.

In JSON, the response looks like this:

json
{
   "result": 4
}
[!NOTE]

Error
In case of errors, the structure (also in responses) is different:

type ResponseError struct {
   Error string `json:"error"`
}
In JSON, it looks like this:

json
{
   "error": "Error"
}
Errors are formatted using vanerrors, and can include messages and causes.

[!TIP]

Possible error messages include:

"method not allowed"
"invalid body"
"page not found"
"bracket should be opened"
"bracket should be closed"
"number parsing error"
"unknown operator"
"divide by zero not allowed"
"unknown calculator error"
These errors occur due to invalid requests, not server errors. Errors will also contain messages and causes.

How to run the application
[!IMPORTANT]

Configuration file:

You need to create a configuration file.

Example JSON structure can be found in the current config.

Don't forget to edit the path in the main.

[!TIP]

port: the port for the server to run.
path: the endpoint of the API.
do_log: whether to log every request in the calc service.
[!IMPORTANT]

Running:

In the terminal, run the following command:

shell
go run cmd/main.go
Make sure you have Go version >= 1.23.0 installed.

Examples
[!WARNING]

If you're using Windows, it's recommended to use Git Bash or WSL for cURL requests. cURL might not work correctly in Command Prompt or PowerShell.

200 (OK)
shell
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'
Result:

json
{
   "result": 2
}
400 (Bad Request)
shell
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data 'bebebe'
Result:

json

{
   "error": "invalid body"
}
405 (Method Not Allowed)
shell
curl --request GET 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'
Result:

json
{
   "error": "method not allowed"
}
422 (Unprocessable Entity)
shell
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+"}'
Result:

json
{
   "error": "number parsing error: '' is not a number in expression '1+'"
}
(Or other invalid expressions)

[!TIP]

To see more examples, view tests.

License
MIT

go