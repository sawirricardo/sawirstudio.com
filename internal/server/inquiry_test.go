package server

import "testing"

func TestLooksLikeSpam(t *testing.T) {
	for _, c := range []struct {
		msg   string
		needs []string
		spam  bool
	}{
		{"READ THE NEWS AND EARN $50,000,000 PER MONTH ASSUREDLY https://telegra.ph/Read-the-latest-news", nil, true},
		{"Visit www.example.com for cheap SEO", nil, true},
		{"Check HTTP://EXAMPLE.COM", nil, true},
		{"Our current site is https://acme.example, we want a new one.", []string{"Website"}, false},
		{"Salut, ech wollt Äre Präis wëssen.", nil, false},
		{"We need a booking app.", []string{}, false},
	} {
		if got := (inquiryInput{Message: c.msg, Needs: c.needs}).looksLikeSpam(); got != c.spam {
			t.Errorf("looksLikeSpam(%q, %v) = %v, want %v", c.msg, c.needs, got, c.spam)
		}
	}
}
