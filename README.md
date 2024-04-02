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

# Full Documentation
Please open `https://mesquite-plier-172.notion.site/URL-Shortener-8fe4e5bb3a0e40d7b383e67a18136dd0` to see full documentation of the system.