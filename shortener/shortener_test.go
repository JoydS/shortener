package shortener

import (
	"net/url"
	"testing"
)

// TestSlugURL checks that SlugURL returns the expected slug for a given URL.
func TestSlugURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Google.com",
			input: "https://www.google.com",
			want:  "20f5081d41",
		},
		{
			name:  "URL with path",
			input: "https://www.example.com/path/to/page",
			want:  "3b03f99239",
		},
		{
			name:  "URL with query params",
			input: "https://www.example.com/search?q=golang",
			want:  "d99b8bbd00",
		},
		{
			name:  "URL no secure",
			input: "http://www.example.com",
			want:  "2f7d4c4d8a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.input)
			if err != nil {
				t.Fatalf("Erreur lors de l'analyse de l'URL %q: %v", tt.input, err)
			}

			got, err := SlugURL(u)
			if err != nil {
				t.Errorf("SlugURL(%q) a renvoyé une erreur : %v", tt.input, err)
				return
			}

			if got != tt.want {
				t.Errorf("SlugURL(%q) = %q; attendu %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestSlugURLError checks that SlugURL returns an error when the URL is nil.
func TestSlugURLError(t *testing.T) {
	_, err := SlugURL(nil)
	if err == nil {
		t.Errorf("SlugURL(nil) aurait dû renvoyer une erreur")
	}
}
