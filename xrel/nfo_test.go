package xrel

import (
	"testing"
)

func TestGetNFOByID(t *testing.T) {
	configureOAuth2ForTests(t)
	// https://www.xrel.to/movie-nfo/60557/Keinohrhasen-German-AC3-DVDRip-XviD-CRUCiAL.html
	bytes, err := GetNFOByID("f638d1cfec8d", false)
	if err != nil {
		t.Error("Unexpected error: " + err.Error())
	} else {
		if len(bytes) == 0 {
			t.Error("Expected any image data.")
		}
	}

	// https://www.xrel.to/p2p/11238-Killer-Elite-2011-German-AC3D-DL-720p-HDTV-x264-TwixX/nfo.html
	bytes, err = GetNFOByID("6dbb52db2be6", true)
	if err != nil {
		t.Error("Unexpected error: " + err.Error())
	} else {
		if len(bytes) == 0 {
			t.Error("Expected any image data.")
		}
	}
}
