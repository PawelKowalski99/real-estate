Real estate analyzer - server


⚠️ Now, the system need connect to a database, by default you can use Postgres or CockroachDB
- Just create a database
- Upload the schema on your DB
```
/db/schema.sql
```
- Change the enverioment variables located in .env.development
```bash
DB_ENGINE = "postgres or mysql"
DB_HOST = "host"
DB_PORT = 1234
DB_DATABASE = "db name"
DB_USERNAME = "user name"
DB_PASSWORD = "your secret password"

# For cockroach
DB_OPTIONS = "--cluster=cockroach-cluser-id"

# For postgres
DB_OPTIONS = "sslmode=disable timezone=UTC connect_timeout=5"

# For mysql
DB_OPTIONS = ""
```
<br />

## ✅ Run

If you want run in local, simply

```bash
go get
go run main.go
```

If you want run with Docker, simply
```bash
make docker-up 
```

or
```
docker compose up --build

If you want run with Heroku local, simply
```bash
# For Linux/MacOS
make heroku-run

# For Windows
heroku-run-win
```

⚠️ If you use Windows, you need change **Procfile** in the root of project for
```bash
# For windows use 
web: bin\real-estate.exe
# web: bin/hex-arch-go
```
And now you can run Heroku local for Windows
<br /><br />

## ✅ Test
For unit tests, simply
```bash
make unit-test
```

⚠️For integration tests, first need configure the **.env.test** vars adding the database test connection, after, simply
```bash
make integration-test
```
Or both of them
```bash
make test
```
<br />

## 🌳 Understanding the folder structure
```bash
.
├── /.github/workflows       # Github Actions!
├── /cmd                     # Start the application with server and database
├── /core                    # The CORE of hexagonal architecture: infrastructure, application and domain
│   ├── /application         # Handlers and the entry point of data
│   ├── /entities            # The entities what conform the domain
│   └── /infrastructure      # Gateways for the domain logic and Storage/Repository for the implementation of database
├── /db-data                      # Simply the schema of DB for you first run
├── /env                     # .env loader
├── /internal                # Elemental logic common for all the system
│   ├── /database            # Connection with database implemented
│   └── /helpers             # Reusable functions around the app, like a UUID generation
│       └── tests            # Unit tests for helpers 
└── /server                  # The server listener and endpoints of API REST
```
