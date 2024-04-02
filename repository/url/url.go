package url

import (
	"database/sql"
	"fmt"

	"github.com/patrickmn/go-cache"
	"github.com/pradiptarana/url-shortener/model"
)

// URLRepository is responsible for storing and retrieving URL information
type URLRepository struct {
	db   *sql.DB
	cche *cache.Cache
}

// NewURLRepository creates a new URLRepository instance
func NewURLRepository(db *sql.DB, cche *cache.Cache) *URLRepository {
	return &URLRepository{db, cche}
}

func (tr *URLRepository) CreateURL(data *model.URL) error {
	stmt, err := tr.db.Prepare("insert into `url` (`id`, `original_url`, `short_url`, `created_at`) values (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(data.Id, data.OriginalURL, data.ShortURL, data.CreatedAt)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (tr *URLRepository) GetOriginalURL(shortURL string) (string, error) {
	stmt, err := tr.db.Prepare("select original_url from url where short_url = ?")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	var result string
	err = stmt.QueryRow(shortURL).Scan(&result)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return result, nil
}

func (tr *URLRepository) GetMaxID() (int64, error) {
	stmt, err := tr.db.Prepare("select max(id) from url")
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	var result sql.NullInt64
	err = stmt.QueryRow().Scan(&result)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	if !result.Valid {
		return 0, nil
	}

	return result.Int64, nil
}

func (tr *URLRepository) GetOriginalURLCache(shortURL string) string {
	if x, found := tr.cche.Get(shortURL); found {
		return x.(string)
	}
	return ""
}

func (tr *URLRepository) SetURLCache(url *model.URL) {
	tr.cche.Set(url.ShortURL, url.OriginalURL, cache.DefaultExpiration)
}
