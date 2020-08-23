package url

type memoryRepository struct {
	urls   map[string]*Url
	clicks map[string]int
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{
		make(map[string]*Url),
		make(map[string]int),
	}
}

func (r *memoryRepository) IdExists(id string) bool {
	_, exist := r.urls[id]
	return exist
}

func (r *memoryRepository) SearchById(id string) *Url {
	return r.urls[id]
}

func (r *memoryRepository) SearchByUrl(url string) *Url {
	for _, u := range r.urls {
		if u.Destination == url {
			return u
		}
	}

	return nil
}

func (r *memoryRepository) Store(url Url) error {
	r.urls[url.Id] = &url
	return nil
}

func (r *memoryRepository) RegisterClick(id string) {
	r.clicks[id] += 1
}

func (r *memoryRepository) SearchClicks(id string) int {
	return r.clicks[id]
}
