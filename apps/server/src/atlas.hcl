data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/migrator",
  ]
}
env "local" {
  url = "postgres://postgres:123456@localhost:5432/norastro?search_path=public&sslmode=disable"
  src = data.external_schema.gorm.url
  dev = "postgres://postgres:123456@localhost:5432/postgres?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
env "local-seed" {
  url = "postgres://postgres:123456@localhost:5432/norastro?search_path=public&sslmode=disable"
  src = data.external_schema.gorm.url
  dev = "docker://postgres"
  migration {
    dir = "file://migrations/seeders"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}