package country

import (
	"net/http"
)

var (
	chinaPool   cityPool
	americaPool cityPool
	englandPool cityPool
)

type cityFunc func(w http.ResponseWriter, r *http.Request)

type cityPool map[string]cityFunc

func (p cityPool) GetRandomCity() cityFunc {
	for _, city := range p {
		return city
	}
	return nil
}

func init() {
	chinaPool = cityPool{
		"beijing":   Beijing,
		"shanghai":  Shanghai,
		"guangzhou": Guangzhou,
		"handan":    Handan,
	}

	americaPool = cityPool{
		"new-york":    NewYork,
		"los-angeles": LosAngeles,
		"chicago":     Chicago,
		"edmond":      Edmond,
	}

	englandPool = cityPool{
		"london":     London,
		"manchester": Manchester,
		"liverpool":  Liverpool,
		"edinburgh":  Edinburgh,
	}
}
