# _**Curso de Go - Construindo a API**_

### :pushpin: **Criação e instalação de dependências**

```powershell
> go mod init api
> go get github.com/gorilla/mux
> go get github.com/joho/godotenv
> go get github.com/go-sql-driver/mysql
> go mod tidy
```

### :pen: **Observações:**

> O comando **[ [go mod tidy][2] ]** é utilizado em projetos Go para limpar e organizar o módulo atual, removendo dependências não utilizadas e baixando quaisquer pacotes necessários que estejam faltando. Ele ajuda a manter o arquivo go.mod e o go.sum atualizados e otimizados.

> Podemos usar o [Scalar][1] para criar modelo visal de acesso aos endpoints

### :pushpin: **Criação do banco de dados**

```powershell
> docker run --name mysql-container `
    -e MYSQL_ROOT_PASSWORD=golang `
    -e MYSQL_USER=golang `
    -e MYSQL_PASSWORD=golang `
    -e MYSQL_DATABASE=devbook `
    -d -p 3306:3306 `
    -v meu-volume-mysql:/var/lib/mysql `
    mysql:latest

> docker exec -it mysql-container mysql -ugolang -psenhagogolang devbook
```

### :pushpin: **Criação de tabelas do banco de dados**

```sql
CREATE DATABASE IF NOT EXISTS devbook;

DROP TABLE IF EXISTS usuarios

CREATE TABLE usuarios (
    ID int UNSIGNED NOT NULL AUTO_INCREMENT,
    Nome VARCHAR(50) NOT NULL,
    Nick VARCHAR(50) NOT NULL unique,
    Email VARCHAR(50) NOT NULL unique,
    Senha VARCHAR(50) NOT NULL,
    CriadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (ID)
);
```

[1]: https://scalar.com "Interface Open Source do OpenApi"
[2]: https://go.dev/ref/mod "Go Modules Reference"
[3]: https://github.com/gorilla/mux "Gorilla/mux"
[4]: https://github.com/joho/godotenv "GoDotEnv"
[5]: https://github.com/go-sql-driver/mysql "Go-MySQL-Driver"
[6]: https://www.robsonalves.dev.br "Web site Robson Candido dos Santos Alves"
