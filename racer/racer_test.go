package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("simple check for slow vs faster", func(t *testing.T) {
		slowServer := makeDelayedServer(10 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, error := Racer(slowURL, fastURL)

		if error != nil {
			t.Errorf("Error is not expected")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("return and error if server doesn't respond <10s", func(t *testing.T) {
		serverA := makeDelayedServer(2 * time.Millisecond)
		serverB := makeDelayedServer(2 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 1*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't got")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
