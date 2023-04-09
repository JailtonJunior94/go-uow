# Curso - FullCyle [Go Expert]
## Módulo [Unit of Work]

[Migrate](https://github.com/golang-migrate/migrate) </br>
[SQLC](https://sqlc.dev/)

### Criando migration
`migrate create -ext=sql -dir=sql/migrations -seq init`

### Executando migration
`migrate --path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up`

### Gerando package com sqlc
`brew install sqlc` </br>
`sqlc generate`