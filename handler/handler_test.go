package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func TestSearchHandler(t *testing.T) {
	os.Setenv("URL", "http://www.omdbapi.com")
	os.Setenv("API_KEY", "faf7e5bb")

	r := newRequest("GET", "/search?s=Batman&page=2")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: god %v want %v", status, http.StatusOK)
	}

	expected := `{"Search":[{"Title":"Batman: The Killing Joke","Year":"2016","imdbID":"tt4853102","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: Under the Red Hood","Year":"2010","imdbID":"tt1569923","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: The Dark Knight Returns, Part 1","Year":"2012","imdbID":"tt2313197","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"},{"Title":"Batman: Mask of the Phantasm","Year":"1993","imdbID":"tt0106364","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYTRiMWM3MGItNjAxZC00M2E3LThhODgtM2QwOGNmZGU4OWZhXkEyXkFqcGdeQXVyNjExODE1MDc@._V1_SX300.jpg"},{"Title":"Batman: The Dark Knight Returns, Part 2","Year":"2013","imdbID":"tt2166834","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg"},{"Title":"Batman","Year":"1966","imdbID":"tt0060153","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMmM1OGIzM2UtNThhZS00ZGNlLWI4NzEtZjlhOTNhNmYxZGQ0XkEyXkFqcGdeQXVyNTkxMzEwMzU@._V1_SX300.jpg"},{"Title":"Batman: Year One","Year":"2011","imdbID":"tt1672723","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNTJjMmVkZjctNjNjMS00ZmI2LTlmYWEtOWNiYmQxYjY0YWVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: Assault on Arkham","Year":"2014","imdbID":"tt3139086","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZDU1ZGRiY2YtYmZjMi00ZDQwLWJjMWMtNzUwNDMwYjQ4ZTVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman Beyond","Year":"1999–2001","imdbID":"tt0147746","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BZWJhNjA4YTAtODBlMS00NjIzLThhZWUtOGYxMGM3OTRmNDZmXkEyXkFqcGdeQXVyNjk1Njg5NTA@._V1_SX300.jpg"},{"Title":"Batman: Gotham Knight","Year":"2008","imdbID":"tt1117563","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BM2I0YTFjOTUtMWYzNC00ZTgyLTk2NWEtMmE3N2VlYjEwN2JlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"}],"totalResults":"523","Response":"True"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}

func TestDetailHandler(t *testing.T) {
	os.Setenv("URL", "http://www.omdbapi.com")
	os.Setenv("API_KEY", "faf7e5bb")

	vars := map[string]string{
		"id": "tt4853102",
	}

	r := newRequest("GET", "/detail/tt4853102")
	r = mux.SetURLVars(r, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DetailHandler)
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: god %v want %v", status, http.StatusOK)
	}
	expected := `{"Title":"Batman: The Killing Joke","Year":"2016","Rated":"R","Released":"25 Jul 2016","Runtime":"76 min","Genre":"Animation, Action, Crime","Director":"Sam Liu","Writer":"Brian Azzarello, Brian Bolland, Bob Kane","Actors":"Kevin Conroy, Mark Hamill, Tara Strong","Plot":"As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.","Language":"English","Country":"United States","Awards":"1 win & 2 nominations","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.4/10"},{"Source":"Rotten Tomatoes","Value":"36%"}],"Metascore":"N/A","imdbRating":"6.4","imdbVotes":"58,167","imdbID":"tt4853102","Type":"movie","DVD":"02 Aug 2016","BoxOffice":"$3,775,000","Production":"N/A","Website":"N/A","Response":"True"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}
