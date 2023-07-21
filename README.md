# golang-backend
clean architecture
### Run this project
- Clone this project

```bash
cd workspace
git clone https://github.com/CGW1996/golang-backend.git
cd golang-backend
cp .env.example .env
# run with docker
docker-compose up -d --build
```
- Access API using http://localhost:8080

- Check test coverage
```bash
cd go-backend
go test -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
open cover.html
```