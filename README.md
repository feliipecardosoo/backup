# ETL – Backup Diário de Usuários

## Visão Geral
Sistema ETL em Go que realiza backup diário de usuários de um MongoDB inicial (`input`) para um MongoDB de backup (`output`), filtrando apenas usuários modificados no dia.

Fluxo principal:
1. Buscar usuários do dia no banco inicial.
2. Deletar registros existentes no backup (pelo nome).
3. Inserir os usuários novos no backup.
4. Registrar tempo de execução.

## Tecnologias
- Go (Golang)
- MongoDB
- Estrutura de pastas:
    - `/config` – Configurações e conexão Mongo
    - `/exec/input` – Leitura do banco inicial
    - `/exec/output` – Escrita no backup
    - `/model` – Modelos de dados (`Membro`)
    - `main.go` – Ponto de entrada do ETL

## Pacotes

### input
- `GetUsersHoje(ctx)` → Retorna usuários do dia.

### output
- `DeleteExistingUsers(ctx, users)` → Remove usuários existentes pelo nome.
- `InsertUsersToday(ctx, users)` → Insere usuários filtrados no backup.
