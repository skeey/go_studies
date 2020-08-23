package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	size = 5
	symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

type Repository interface {
	IdExists(id string) bool
	SearchById(id string) *Url
	SearchByUrl(url string) *Url
	Store(url Url) error
	RegisterClick(id string)
	SearchClicks(id string) int
}

type Url struct {
	Id string
	Created_at time.Time
	Destination string
}

type Stats struct {
	Url *Url
	Clicks int
}

var repo Repository

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SetRepository(r Repository) {
	repo = r
}

func RegisterClick(id string) {
	repo.RegisterClick(id)
}

func SearchOrCreateNewUrl(destination string) (u *Url, new bool, err error) {
	if u = repo.SearchByUrl(destination); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(destination); err != nil {
		return nil, false, err
	}

	url := Url{generateId(), time.Now(), destination}
	repo.Store(url)
	return &url, true, nil
}

func Search(id string) *Url {
	return repo.SearchById(id)
}

func (u *Url) Stats() *Stats {
	clicks := repo.SearchClicks(u.Id)
	return &Stats{u, clicks}
}

func generateId() string {
	newId := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = symbols[rand.Intn(len(symbols))]
		}
		return string(id)
	}
	
	for {
		if id := newId(); !repo.IdExists(id) {
			return id
		}
	}
}
