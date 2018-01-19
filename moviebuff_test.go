package moviego

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMoviebuff_GetMovie(t *testing.T) {
	m := new(Moviebuff)
	m.Init(os.Getenv("API_TOKEN"), log.New(os.Stdout, "", log.Llongfile))
	_, err := m.GetMovie("dd083a9d-823b-4ffc-b057-d8885840fcf7")
	ShouldBeEqual(t, nil, err)
}

func TestMovie_GetEarliestReleaseYear(t *testing.T) {
	m := new(Moviebuff)
	m.Init(os.Getenv("API_TOKEN"), log.New(os.Stdout, "", log.Llongfile))
	mv, err := m.GetMovie("dd083a9d-823b-4ffc-b057-d8885840fcf7")
	ShouldBeEqual(t, 2014, mv.GetEarliestReleaseYear())
	ShouldBeEqual(t, nil, err)
}

func TestMoviebuff_GetPerson(t *testing.T) {
	m := new(Moviebuff)
	m.Init(os.Getenv("API_TOKEN"), log.New(os.Stdout, "", log.Llongfile))
	_, err := m.GetPerson("9494a132-4447-42e9-ae49-a66ac071a36a")
	ShouldBeEqual(t, nil, err)
}

func TestMoviebuff_GetEntity(t *testing.T) {
	m := new(Moviebuff)
	m.Init(os.Getenv("API_TOKEN"), log.New(os.Stdout, "", log.Llongfile))
	_, err := m.GetEntity("1e14fd01-3a4f-4b0b-a164-fe50c0fab13d")
	ShouldBeEqual(t, nil, err)
}

func TestMoviebuff_GetResources(t *testing.T) {
	m := new(Moviebuff)
	m.Init(os.Getenv("API_TOKEN"), log.New(os.Stdout, "", log.Llongfile))
	_, err := m.GetResources("people", 20, 1)
	ShouldBeEqual(t, nil, err)
}

func ShouldBeEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("Expected %v. Got %v. Location: %s:%d", expected, actual, fn, line)
	}
}
