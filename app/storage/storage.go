package storage

import (
	"sync"
	"time"

	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
	"github.com/google/uuid"
)

var storage = Storage{Mutex: &sync.Mutex{}, Data: map[string]*Session{}}

type (
	Storage struct {
		Mutex *sync.Mutex
		Data  map[string]*Session
	}
	Session struct {
		ID      string
		Name    string
		Timer   timer.Timer
		Spotify Spotify
	}
	Spotify struct {
		AccessToken   string
		RefreshToken  string
		ExpiresIn     int
		RefreshedAt   time.Time
		RefreshFailed bool
	}
)

func (s Spotify) TokenDue() timer.Due {
	const refreshBeforeExpiration = time.Minute

	return timer.Due{
		Start:    s.RefreshedAt,
		Duration: (time.Duration(s.ExpiresIn) * time.Second) - refreshBeforeExpiration,
	}
}

func Get(id string) (*Session, bool) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	session, ok := storage.Data[id]

	return session, ok
}

func Set(session Session) string {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	for id, u := range storage.Data {
		if u.ID == session.ID {
			storage.Data[id] = &session

			return id
		}
	}

	var id = uuid.New().String()

	storage.Data[id] = &session

	return id
}

func All() []Session {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	var all = make([]Session, 0, len(storage.Data))

	for _, session := range storage.Data {
		all = append(all, *session)
	}

	return all
}

func MutateAll(mutate func(s *Session)) {
	storage.Mutex.Lock()

	for _, session := range storage.Data {
		mutate(session)
	}

	storage.Mutex.Unlock()
}
