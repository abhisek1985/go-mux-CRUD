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
- Create merchant (/api/new-merchant)
- Get all merchants (/api/merchant or /api/merchant?PageNum=<page number>&PageSize=<page size>)
- Get merchant (/api/merchant/{id})
- Update merchant (/api/merchant/{id})
- Delete merchant (/api/del-merchant/{id})

### Team APIs:
- Create team (/api/new-team)
- Get all teams (/api/team or /api/team?PageNum=<page number>&PageSize=<page size>)
- Get team (/api/team/{id})
- Update team (/api/team/{id})
- Delete team (/api/del-team/{id})
- Get list of team members per merchant (/api/team/merchant/{merchant_id} or /api/team/merchant/{merchant_id?PageNum=<page number>&PageSize=<page size>}