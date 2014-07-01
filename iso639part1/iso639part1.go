package iso639part1

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	indexCode     = 2
	indexEngNames = 0
	indexLocNames = 1
)

var (
	MalformedDataError = errors.New("Data could not be parsed correctly.")
)

type Map map[string]*Entry

type Entry struct {
	Code       string   `json:"code"`
	EngName    string   `json:"en_name"`
	AltEngName []string `json:"alternative_en_names"`
	LocName    string   `json:"local_name"`
	AltLocName []string `json:"alernative_local_names"`
}

func GetMap(f *os.File) (m Map, err error) {
	s := bufio.NewScanner(f)
	m = map[string]*Entry{}
	for s.Scan() {
		// Initial line processing
		l := s.Text()
		lp := strings.Split(l, "	")
		if len(lp) != 3 {
			return m, MalformedDataError
		}

		// Create entry and set code
		e := Entry{}
		e.Code = lp[indexCode]

		// Set English names
		enp := strings.Split(lp[indexEngNames], ",")
		e.EngName = enp[0]
		if len(enp) > 1 {
			e.AltEngName = enp[1:]
		}

		// Set localized names
		lnp := strings.Split(lp[indexLocNames], ",")
		e.LocName = lnp[0]
		if len(lnp) > 1 {
			e.AltLocName = lnp[1:]
		}

		m[e.Code] = &e
	}
	return m, err
}
