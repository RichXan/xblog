package designpattern

// MediaPlayer 是目标接口
type MediaPlayer interface {
    Play() 
}

// AdvancedMediaPlayer 是需要适配的高级播放器接口
type AdvancedMediaPlayer interface {
    PlayVlc()
    PlayMp4()
}

// VlcPlayer 实现高级播放器接口
type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc() {
    println("Playing vlc file")
}

func (v *VlcPlayer) PlayMp4() {
    // 不支持MP4
}

// Mp4Player 实现高级播放器接口
type Mp4Player struct{}

func (m *Mp4Player) PlayVlc() {
    // 不支持VLC
}

func (m *Mp4Player) PlayMp4() {
    println("Playing mp4 file")
}

// MediaAdapter 是适配器
type MediaAdapter struct {
    advancedMediaPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
    var advancedMediaPlayer AdvancedMediaPlayer
    if audioType == "vlc" {
        advancedMediaPlayer = &VlcPlayer{}
    } else if audioType == "mp4" {
        advancedMediaPlayer = &Mp4Player{}
    }
    return &MediaAdapter{advancedMediaPlayer}
}

func (m *MediaAdapter) Play() {
    switch m.advancedMediaPlayer.(type) {
    case *VlcPlayer:
        m.advancedMediaPlayer.PlayVlc()
    case *Mp4Player:
        m.advancedMediaPlayer.PlayMp4()
    }
}

// AudioPlayer 是实现了 MediaPlayer 接口的播放器
type AudioPlayer struct {
    mediaAdapter *MediaAdapter
}

func (a *AudioPlayer) Play() {
    if a.mediaAdapter != nil {
        a.mediaAdapter.Play()
    } else {
        println("Playing mp3 file")
    }
}

func NewAudioPlayer(audioType string) *AudioPlayer {
    var adapter *MediaAdapter
    if audioType == "vlc" || audioType == "mp4" {
        adapter = NewMediaAdapter(audioType)
    }
    return &AudioPlayer{adapter}
}