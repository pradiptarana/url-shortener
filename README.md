# url-shortener
This is a repository for url shortener in Golang

# How to Run
To run the project please do the following:
- Setup MySQL and run the table creation query below:
`CREATE TABLE url
(
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  original_url varchar(300),
  short_url varchar(10),
  created_at int,
  UNIQUE(short_url)
);`

- Create `.env` file. Please refer to `.env.sample` for the env variable that used in this project.
- Run `go run main.go`.

# How to Run Test
To run the unit test please do the following:
- Run `go test ./...`