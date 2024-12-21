Calculator
Overview
The Calculator is a web-based HTTP API specifically designed to evaluate mathematical expressions. It supports basic operations like addition, subtraction, multiplication, and division, as well as the use of parentheses for precedence control.

Supported Operations
Name	Symbol	Supported	Error When	Description
Number	float64	☑	-	Represents a numerical value in calculations.
<ins>Multiply</ins>	*	☑	-	Has higher precedence over addition and subtraction but lower than parentheses.
<ins>Division</ins>	/	☑	Division by 0 is not allowed	Has higher precedence over addition and subtraction but lower than parentheses.
<ins>Adding</ins>	+	☑	-	Operates with lower precedence, equivalent to subtraction.
<ins>Subtraction</ins>	-	☑	-	Operates with lower precedence, equivalent to addition.
<ins>Brackets</ins>	(, )	☑	Bracket not closed / not opened	Family of highest precedence; for structures like
10(1+1)
which equals
102
.
<ins>Other</ins>	any	☒	Cannot convert to float	Avoid using unsupported symbols or characters.
HTTP Methods
Name	Response Status	Method	Path	Body
OK	200	POST	/api/v1/calculate	
{"expression":"2+2"}
Wrong Method	405	GET	/api/v1/calculate	
{"expression":"2+2"}
Wrong Path	404	POST	/any/unsupported/path	
{"expression":"2+2"}
Invalid Body	400	POST	/api/v1/calculate	
invalid body
Error Calculation	422	POST	/api/v1/calculate	
{"expression":"2*(2+2)"}
Request Structure
To send requests, a specific structure for the input is used. This is defined in input:

type Request struct {
    Expression string `json:"expression"`
}
Example JSON Body
The expression should be formatted as follows in the JSON body:

{
    "expression": "2+2"
}
You can find more examples of valid expressions in the calculator tests (TestCalc), while examples of invalid expressions are available in TestCalcErrors.

Server Response
Successful Response (HTTP 200)
Responses adhere to a structure defined in responses:

type ResponseOK struct {
    Result float64 `json:"result"`
}
In JSON, a successful response will look like this:

{
    "result": 4
}
Error Response
Error responses utilize a different structure, still in responses:

type ResponseError struct {
    Error string `json:"error"`
}
In JSON, this will appear as:

{
    "error": "Error"
}
Various error messages are formatted following the conventions of vanerrors for simple error handling.

Common Error Messages
"method not allowed"
"invalid body"
"page not found"
"bracket should be opened"
"bracket should be closed"
"number parsing error"
"unknown operator"
"divide by zero not allowed"
"unknown calculator error"
All these errors are generated due to invalid requests rather than server issues, and each will have accompanying messages and causes.

Running the Application
Configuration
Before running the application, create a configuration file with the required structure found in current config.

Important Configuration Parameters
port: The port number on which the server will run.
path: The API endpoint.
do_log: Specifies whether the application should log every request.
Starting the Application
To run the application, use the following command in your console:

go run cmd/main.go
Ensure you are using Go version 1.23.0 or above.

Example Usage
Using cURL
Warning: Use Git Bash or WSL for cURL requests, as cURL does not work properly with Command Prompt or PowerShell.

Successful Request (HTTP 200)
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'
Result:

{
    "result": 2
}
Invalid Body (HTTP 400)
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data 'bebebe'
Result:

{
    "error": "invalid body"
}
Method Not Allowed (HTTP 405)
curl --request GET 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'
Result:

{
    "error": "method not allowed"
}
Unprocessable Entity (HTTP 422)
curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+"}'
Result:

{
    "error": "number parsing error: '' is not a number in expression '1+'"
}
(Or may return other errors for invalid expressions.)

Additional Examples
For more examples, please refer to the tests.