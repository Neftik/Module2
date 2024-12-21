# Calculator

## What is Calculator?

The Calculator API is an HTTP service designed to process mathematical expressions using numbers and operators such as "+", "-", "*", and "/". It also supports brackets "(" and ")".

### Calculator Operations

| Name                | Symbol       | Supported | Error when                                 | Description                                                                |
| ------------------- | ------------ | --------- | ----------------------------------------- | -------------------------------------------------------------------------- |
| **Number**           | `float64`    | ☑         | -                                         | A number (e.g., `1`, `3.14`). Nothing special.                             |
| **Multiply**         | `*`          | ☑         | -                                         | Has higher priority than addition and subtraction, but lower than brackets.|
| **Division**         | `/`          | ☑         | Division by zero is not allowed           | Has higher priority than addition and subtraction, but lower than brackets.|
| **Addition**         | `+`          | ☑         | -                                         | Has lower priority (same as subtraction).                                 |
| **Subtraction**      | `-`          | ☑         | -                                         | Has lower priority (same as addition).                                    |
| **Brackets**         | `(`, `)`     | ☑         | Bracket not opened/closed correctly       | Brackets have the highest priority. Note: `10(1+1)` is `102`.             |
| **Other Characters** | Any          | ☒         | Can't convert to float                    | Avoid using unsupported characters.                                       |

### HTTP Requests and Responses

#### Request

The API expects a JSON request with a mathematical expression.

```json
{
   "expression": "2+2"
}

Responses
OK: Successful calculation.

Status Code: 200
Response Body:
Responses
OK: Successful calculation.

Status Code: 200
Response Body:
json

{
   "result": 4
}
Error: Error in processing the expression.

Status Code: 422 (Unprocessable Entity)
Response Body:
json
{
   "error": "number parsing error"
}
Error Messages
Possible error messages include:

method not allowed
invalid body
page not found
bracket should be opened
bracket should be closed
number parsing error
unknown operator
divide by zero not allowed
unknown calculator error
How Requests Are Sent
Use the following format for sending requests:

json
{
   "expression": "2+2"
}
Example of Valid Request (200 OK)
bash
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'
Response:

json
{
   "result": 2
}
Example of Invalid Body (400 Bad Request)
bash
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data 'bebebe'
Response:

json
{
   "error": "invalid body"
}
Example of Method Not Allowed (405)
bash
curl --request GET 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'
Response:

json
{
   "error": "method not allowed"
}
Example of Unprocessable Entity (422)
bash
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+"}'
Response:

json
{
   "error": "number parsing error: '' is not a number in expression '1+'"
}