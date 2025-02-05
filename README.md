![Coverage](https://img.shields.io/badge/coverage-80%25-green)
![Go](https://img.shields.io/badge/Go-1.23-blue)
![License](https://img.shields.io/badge/license-MIT-green)

# Numbers API
Get a random number fact from the Numbers API.
Categorise the number

# Technologies used:
- Golang > 1.23

## Setting up the repo locally:

Clone the repo locally:
```bash
git clone https://github.com/wathika-eng/HNG12 --depth 1 && cd HNG12 && git -b numbers-api
```
You must have Golang installed on your machine to run this project. You can download Golang from the [official website](https://golang.org/dl/).


Run the following command to start the server: # will install the dependencies automatically
```bash
go run .
# or
make run
```

Then open your browser and navigate to `http://localhost:8000/api/classify-number` to see the API in action.
```bash
curl -X GET http://localhost:8000/api/classify-number?number=5
```

Sample response:
```bash
{
  "number": 5,
  "is_prime": true,
  "is_perfect": false,
  "properties": [
    "armstrong",
    "odd"
  ],
  "digit_sum": 5,
  "fun_fact": "5 is the number of babies born in a quintuplet."
}
```

## Folder structure

```bash
├── api_test.go    # Unit tests for API endpoints
├── go.mod         # Go module file, manages dependencies
├── go.sum         # Checksum file for module dependencies
├── main.go        # Entry point of the Go application
├── Makefile       # Automation script running commands
├── math.go        # Contains mathematical operations
├── parseNum.go    # Handles parsing and processing of numerical input
├── README.md      # Documentation for the project
└── types.go       # Defines API response structure

```

## Hire Golang Developers

If you're looking to hire Golang developers, check out [HNG Tech - Hire Golang Developers](https://hng.tech/hire/golang-developers).

## 📝 License
MIT License

## 📢 Contact
wathika02@gmail.com
