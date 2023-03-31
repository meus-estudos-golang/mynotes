# MyNotes
Uma aplicação simples para gerenciamento de notas pessoais.

## Instalação

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
```

Por fim, popule a base de dados com algumas informações prévias, utilizando os comandos SQL abaixo:

```sql
INSERT INTO notes (title, content, created) VALUES ('Supermercado', 'Presunto\nRequeijão', UTC_TIMESTAMP());
INSERT INTO notes (title, content, created) VALUES ('Para o final de semana', 'Lavar o carro\nCortar grama', UTC_TIMESTAMP());
INSERT INTO notes (title, content, created) VALUES ('Filmes que quero assitir', 'John Wick 3', UTC_TIMESTAMP());
```

