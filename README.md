# HNG12 Internship Tasks
- Github API - Get user details from Github API
- Numbers API - Get a random number fact

# Technologies used:
- Golang > 1.23



## Setting up the repo locally:

Clone the repo locally:
```bash
git clone https://github.com/wathika-eng/HNG12 --depth 1 && cd HNG12
```

If you don't have Golang (1.23.5) installed, you can just execute the binary under the `bin` directory.
e.g on Linux:
```bash
./bin/gith
```

If you have Golang installed, you can run the following commands:


Run the following command to start the server:
```bash
go run .
# or
make run
```

Then open your browser and navigate to `http://localhost:8080/` to see the API in action.
```bash
curl -X GET http://139.59.144.196:8000/
```

Sample response:
```bash
{
  "email": "wathika.wanini1@students.jkuat.ac.ke",
  "current_datetime": "2025-01-29T15:34:49Z",
  "repos_url": "https://api.github.com/users/wathika-eng/repos"
}
```

## Folder structure

```bash
.
├── api_test.go
├── bin
│   ├── gith
│   ├── gith.darwin
│   └── gith.exe
├── go.mod
├── go.sum
├── http
│   ├── default-user.http
│   └── user.http
├── http-client
├── main.go
├── Makefile
├── README.md
├── route.go
└── user.go
```

## Hire Golang Developers

If you're looking to hire Golang developers, check out [HNG Tech - Hire Golang Developers](https://hng.tech/hire/golang-developers).

## 📝 License


## 📢 Contact
