# simple_github_user_api
A simple API which returns a users GitHub username and email.

## Setting up the repo locally:

Clone the repo locally:
```bash
git clone https://github.com/wathika-eng/simple_github_user_api.git --depth 1 && cd simple_github_user_api
```
```bash
cp .env.example .env
```
If you don't have Golang (1.23.5) installed, you can just execute the binary under the `bin` directory.
e.g on Linux:
```bash
./bin/gith
```

If you have Golang installed, you can run the following commands:
```bash
go mod download
```

Then run the following command to start the server:
```bash
go run .
# or
make run
```

Then open your browser and navigate to `http://localhost:8080/` to see the API in action.
