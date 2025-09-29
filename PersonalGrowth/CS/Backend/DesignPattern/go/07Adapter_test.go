package designpattern

import "testing"

func TestAdapter(t *testing.T) {
	audioPlayer := NewAudioPlayer("mp3")
	audioPlayer.Play()
	audioPlayer = NewAudioPlayer("mp4")
	audioPlayer.Play()
	audioPlayer = NewAudioPlayer("vlc")
	audioPlayer.Play()
}
