# 适配器模式（Adapter）
适配器模式将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

适配器模式（Adapter Pattern）充当两个不兼容接口之间的桥梁，属于结构型设计模式。它通过一个中间件（适配器）将一个类的接口转换成客户期望的另一个接口，使原本不能一起工作的类能够协同工作。

这种模式涉及到一个单一的类，该类负责加入独立的或不兼容的接口功能。举个真实的例子，读卡器是作为内存卡和笔记本之间的适配器。您将内存卡插入读卡器，再将读卡器插入笔记本，这样就可以通过笔记本来读取内存卡。

假设有一个音频播放器，它只能播放 MP3 文件。现在，我们需要播放 VLC 和 MP4 文件，可以通过创建一个适配器来实现：

目标接口：定义一个可以播放多种格式文件的音频播放器接口。
适配者类：现有的音频播放器，只能播放 MP3 文件。
适配器类：创建一个新的类，实现目标接口，并在内部使用适配者类来播放 MP3 文件，同时添加对 VLC 和 MP4 文件的支持。

## 示例
1. 不同数据格式的转换
2. 不同接口的兼容
3. 老系统接口适配
4. 第三方库的集成

## 代码实现

```golang
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
```

## 使用示例

```golang
func main() {
    // 播放 MP3
    audioPlayer := NewAudioPlayer("")
    audioPlayer.Play()

    // 播放 VLC
    vlcPlayer := NewAudioPlayer("vlc")
    vlcPlayer.Play()

    // 播放 MP4
    mp4Player := NewAudioPlayer("mp4")
    mp4Player.Play()
}
```

## 类图
```mermaid
classDiagram
    %% 接口定义
    class MediaPlayer {
        <<interface>>
        +Play()
    }

    class AdvancedMediaPlayer {
        <<interface>>
        +PlayVlc()
        +PlayMp4()
    }

    %% 具体播放器实现
    class VlcPlayer {
        +PlayVlc()
        +PlayMp4()
    }

    class Mp4Player {
        +PlayVlc()
        +PlayMp4()
    }

    %% 适配器
    class MediaAdapter {
        -advancedMediaPlayer AdvancedMediaPlayer
        +Play()
    }

    %% 音频播放器
    class AudioPlayer {
        -mediaAdapter MediaAdapter
        +Play()
    }

    %% 关系
    MediaPlayer <|.. AudioPlayer
    MediaPlayer <|.. MediaAdapter
    AdvancedMediaPlayer <|.. VlcPlayer
    AdvancedMediaPlayer <|.. Mp4Player
    AudioPlayer o-- MediaAdapter
    MediaAdapter o-- AdvancedMediaPlayer
```

## 说明
1. MediaPlayer 是目标接口，定义了客户端使用的统一接口
2. AdvancedMediaPlayer 是需要适配的高级播放器接口
3. VlcPlayer 和 Mp4Player 是具体的高级播放器实现
4. MediaAdapter 是适配器，将高级播放器接口转换为目标接口
5. AudioPlayer 是具体的播放器，通过适配器支持多种格式
6. 这种实现允许系统灵活地支持新的媒体格式