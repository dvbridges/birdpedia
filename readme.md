# A simple Golang webapp with Postgres database for local development

Golang code taken from [Soham Kamani](https://medium.com/gojekengineering/build-a-web-application-in-go-golang-991ca05a215f).

See [Dockerfile](./Dockerfile) for Birdpedia build instructions.

## Run the Docker-Compose file
The [docker-compose](./docker-compose.yaml) file will run:
- the birdpedia webapp container
- a Postgres database container
- a PGAdmin container, for creating managing your Postgres databases

```bash
docker-compose pull
docker-compose up
```

The next steps involve setting up the Postgres database via PGAdmin.

## Create Server on PGADMIN
- Log into PGAdmin at localhost:5050 (see pgadmin environment creds in [docker-compose](./docker-compose.yaml))
- On command line, get the postgres docker container IP for the Server connection for pgAdmin: 
```bash
docker inspect postgresql -f “{{json .NetworkSettings.Networks }}”
```
- Right click "Servers" and Create Server
- In General tab, call it anything sensible e.g., "my-postgres"
- In Connections tab, enter postgres container IP in "Host name/address", and in "Username" and "Password", use the postgres container creds (see postgres environment creds in [docker-compose](./docker-compose.yaml), i.e., "user" and "password")
- Click "Save" to create a new Server

## Create Database on new Server on pgAdmin
- Right click "my-postgres" server and click Create > New Database
- Give database the correct name e.g., bird_encyclopedia
- Click Save

## Add the Table to database
- In PGAdmin, go to Servers > Databases > <your-database> > Schemas > Public > Tables
- Right click Tables, and Create Table
- Give Table the correct name of the Table used in your webapp DB calls i.e., "birds"
- Add the necessary columns according to the following SQL, and Save:
```sql
CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  bird VARCHAR(256),
  description VARCHAR(1024)
);
```
- Right click Tables, and select "Query Tool"
- Run a test query:
```sql
INSERT INTO birds(species, description)
VALUES('House Sparrow', 'The house sparrow is a bird of the sparrow family Passeridae, found in most parts of the world');
SELECT species, description FROM birds;
```

## Finally...

With all this in place, you are ready to test out your webapp. Go to `localhost:8080/assets/`


# Running containers separately

If you want to run the containers separetely, without using docker-compose

## Run Birdpedia
```bash
docker run --name birdpedia -p 8080:8080 dvbridges/birdpedia:latest
```

## Run the postgres container
Note, the username and password are required in the webapp for db entry
```bash
docker run --name postgresql -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 -v /data:/var/lib/postgresql/data -d postgres
```
## Start up the pdAdmin container for managing your Postgres Server
Note, the username and password are for the Web UI login
```bash
docker run -p 5050:80 -e "PGADMIN_DEFAULT_EMAIL=pgadmin4@pgadmin.org" -e "PGADMIN_DEFAULT_PASSWORD=test1234" -d dpage/pgadmin4:4.29
```