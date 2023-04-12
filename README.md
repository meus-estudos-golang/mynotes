# MyNotes
Uma aplicação simples para gerenciamento de notas pessoais.

## Configuração da base de dados

Instale e configure o [Docker](https://docs.docker.com/engine/install/ubuntu/) em seu sistema operacional. Depois, execute o comando abaixo para criar um *container* com o SGBD MySQL instalado:

`docker run --name mynotes-db -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mysql:latest`

Entre no *container* criado anteriormente, utilizando do comando:

`docker exec -it mynotes-db bash`

Entre no *prompt* de comandos do MySQL, executando: 

`mysql -u root -p`. 

Quando solicitado, informe a senha utilizada durante a criação do *container*.

Para criar e configurar a base de dados do projeto, execute os comandos SQL abaixo, em ordem:

```sql
CREATE DATABASE mynotes CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE mynotes;
CREATE TABLE notes (id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, title VARCHAR(100) NOT NULL, content TEXT NOT NULL, created DATETIME NOT NULL);
CREATE INDEX idx_notes_created ON notes(created);
CREATE TABLE sessions (token CHAR(43) PRIMARY KEY, data BLOB NOT NULL, expiry TIMESTAMP(6) NOT NULL);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);
CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,name VARCHAR(255) NOT NULL,email VARCHAR(255) NOT NULL,hashed_password CHAR(60) NOT NULL,created DATETIME NOT NULL);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
CREATE DATABASE test_mynotes CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Agora, você deve popular a base de dados com algumas informações prévias, utilizando os comandos SQL abaixo:

```sql
INSERT INTO notes (title, content, created) VALUES ('Supermercado', 'Presunto\nRequeijão', UTC_TIMESTAMP());
INSERT INTO notes (title, content, created) VALUES ('Para o final de semana', 'Lavar o carro\nCortar grama', UTC_TIMESTAMP());
INSERT INTO notes (title, content, created) VALUES ('Filmes que quero assitir', 'John Wick 3', UTC_TIMESTAMP());
```

Para que aplicação funcione corretamente, crie um certificado auto-assinado, com a ferramenta `generate_cert.go`. Para isso, crie um diretório na raiz do projeto, denominado `tls`:

```
mkdir tls
cd tls
```

Por fim, execute o comando abaixo:

`go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`


## Execução da aplicação

Para executar a aplicação, use o comando:

`go run ./cmd/web`

A aplicação estará exeuctando em [http://localhost:4000/](http://localhost:4000/)

Enjoy!


