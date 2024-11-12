# devbook
Curso de Go - Construindo a API


```bash
go mod init api
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
```

```powershell
docker run --name mysql-container `
    -e MYSQL_ROOT_PASSWORD=golang `
    -e MYSQL_USER=golang `
    -e MYSQL_PASSWORD=golang `
    -e MYSQL_DATABASE=devbook `
    -d -p 3306:3306 `
    -v meu-volume-mysql:/var/lib/mysql `
    mysql:latest

docker exec -it mysql-container mysql -ugolang -psenhagogolang devbook
```