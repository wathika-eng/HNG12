![Go](https://img.shields.io/badge/Go-1.23-blue)
![License](https://img.shields.io/badge/license-MIT-green)
# HNG12 Internship Tasks
- Github API - Get user details from Github API [GitHub link](https://github.com/wathika-eng/HNG12/tree/github-api)
- Numbers API - Get a random number fact [GitHub link](https://github.com/wathika-eng/HNG12/tree/numbers-api)

# Technologies used:


- Golang > 1.23 [Golang](https://golang.org/)
- Render [Render](http://render.com)

## Setting up the repo locally:
Clone the repo locally:
```bash
git clone https://github.com/wathika-eng/HNG12 --depth 1 && cd HNG12
```
Switch to the branch you want to work on:
```bash
git -b github-api
# or
git -b numbers-api
```

If you have Golang installed, you can run the following commands:

Run the following command to start the server: (# dependecies will be installed automatically)
```bash
go run .
# or
make run 
```

Then open your browser and navigate to `http://localhost:8000/` to see the API in action or check the terminal for the response.
```bash
# for github api
curl -X GET http://139.59.144.196:8000/
```

Sample response:
```bash
{
  "email": "xxxxx@gmail.com",
  "current_datetime": "2025-01-29T15:34:49Z",
  "repos_url": "https://api.github.com/users/wathika-eng/repos"
}
```

## Folder structure

```bash
.
├── api_test.go # test file
├── go.mod # go module file indicates dependecies and go version
├── go.sum 
├── http # http folder contains http files
│   ├── default-user.http
│   └── user.http
├── main.go # main file
├── Makefile # make file for running the server
├── README.md # readme file
```

## Hire Golang Developers

If you're looking to hire Golang developers, check out [HNG Tech - Hire Golang Developers](https://hng.tech/hire/golang-developers).

## 📝 License
MIT License

## 📢 Contact
wathika02@gmail.com

![Golang Logo](https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg)
