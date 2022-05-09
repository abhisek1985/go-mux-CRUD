# CRUD operations using:
- GO
- PostgreSQL
- Gorilla MUX router
- Docker
- Custom migrations with golang-migrate
- Basic Authentication

### Requirements:
* Docker and Docker Compose
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### Build and start the services with:
```bash
$ docker-compose up --build
```
### Merchant APIs:
- Create merchant (/api/create/merchant)
- Get all merchants (/api/merchants or /api/merchants?PageNum=<page number>&PageSize=<page size>)
- Get merchant (/api/merchant/{id})
- Update merchant (/api/update/merchant/{id})
- Delete merchant (/api/delete/merchant/{id})

### Team APIs:
- Create team (/api/create/team)
- Get all teams (/api/teams or /api/teams?PageNum=<page number>&PageSize=<page size>)
- Get team (/api/team/{id})
- Update team (/api/update/team/{id})
- Delete team (/api/delete/team/{id})
- Get teams per merchant (/api/teams/merchant/{merchant_id} or /api/teams/merchant/{merchant_id?PageNum=<page number>&PageSize=<page size>}