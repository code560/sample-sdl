package main

import (
	"log"
	"time"

	"github.com/code560/sample-sdl/player"
	"github.com/veandco/go-sdl2/mix"
)

func main() {
	test10()
}

func LoadChunks(files []string) []*mix.Chunk {

	chunks := make([]*mix.Chunk, 0)
	for _, file := range files {
		chunk, err := mix.LoadWAV(file)
		if err != nil {
			log.Fatal(err)
		}
		chunks = append(chunks, chunk)

	}
	return chunks
}

func test() {
	err := mix.Init(0)
	if err != nil {
		log.Fatal(err)
	}
	defer mix.Quit()

	chanel := 2 // 2 is stereo
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, chanel, 1024)
	if err != nil {
		log.Fatal(err)
	}
	defer mix.CloseAudio()

	chunks := LoadChunks(
		[]string{
			"asset/audio/bgm_wave.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/se_jump.wav",
			"asset/audio/atari.wav",
		})

	mix.AllocateChannels(len(chunks))

	for i, chunk := range chunks {
		defer chunk.Free()
		chunk.Play(i, 2)
	}

	time.Sleep(10 * time.Second)
}

func test2() {
	err := mix.Init(0)
	if err != nil {
		log.Fatal(err)
	}
	defer mix.Quit()

	chanel := 2 // 2 is stereo
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, chanel, 1024)
	if err != nil {
		log.Fatal(err)
	}
	defer mix.CloseAudio()

	mix.ChannelFinished(func(ch int) {
		chunk := mix.GetChunk(ch)
		if chunk != nil {
			chunk.Free()
			log.Print("free chunk")
		}
	})

	mix.AllocateChannels(100)
	// test 1ch = n chunks
	// chunks := LoadChunks(
	// 	[]string{
	// 		"asset/audio/bgm_wave.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/se_jump.wav",
	// 		"asset/audio/atari.wav",
	// 	})
	// for _, chunk := range chunks {
	// 	ch_, _ := chunk.Play(1, 0)
	// 	if ch_ != 1 {
	// 		log.Printf("not same ch: %d", ch_)
	// 	}
	// }

	// test callback
	chunk, _ := mix.LoadWAV("asset/audio/atari.wav")
	defer chunk.Free()
	chunk.Play(-1, 0)

	time.Sleep(time.Second * 10)
}

func test3() {
	err := mix.Init(0)
	if err != nil {
		log.Fatal(err)
	}
	defer mix.Quit()

	chanel := 2 // 2 is stereo
	for i := 0; i < 100; i++ {
		err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, chanel, 1024)
		if err != nil {
			log.Fatal(err)
		}
	}
	// defer mix.CloseAudio()

	mix.AllocateChannels(200)
	chunk, _ := mix.LoadWAV("asset/audio/atari.wav")
	defer chunk.Free()
	chunk.Play(-1, 0)

	time.Sleep(10 * time.Second)
}

func test9() {
	p1 := player.GetPlayer("hoge")
	go p1.Play("asset/audio/rusuden_04-2.wav")
	time.Sleep(time.Millisecond * 10)

	p1.Stop()
	go p1.Play("asset/audio/okaerigoshujin_01.wav")

	time.Sleep(10 * time.Second)
}

func test10() {
	p1 := player.GetPlayer("hoge")
	p2 := player.GetPlayer("foo")

	go p1.Play("asset/audio/se_rain2.wav")
	go p1.Play("asset/audio/se_jump.wav")
	// go p1.Play("asset/audio/se_jump.wav")
	// go p1.Play("asset/audio/se_jump.wav")
	// go p1.Play("asset/audio/se_jump.wav")
	// go p1.Play("asset/audio/se_jump.wav")
	// go p1.Play("asset/audio/se_jump.wav")
	go p1.Play("asset/audio/se_jump.wav")
	go p1.Play("asset/audio/se_jump.wav")
	go p1.Play("asset/audio/rusuden_04-2.wav")

	time.Sleep(time.Millisecond * 500)

	go p2.Play("asset/audio/info-girl1-start1.wav")
	go p2.Play("asset/audio/se_swing.wav")
	go p2.Play("asset/audio/se_swing.wav")
	go p2.Play("asset/audio/se_swing.wav")
	go p2.Play("asset/audio/baibai_01.wav")

	p1.Stop()

	p1.Play("asset/audio/okaerigoshujin_01.wav")
}

func test20() {
	p1 := player.GetPlayer("hoge")
	p2 := player.GetPlayer("foo")

	go p1.Play("asset/audio/rusuden_04-2.wav")
	go p2.Play("asset/audio/se_rain.wav")

	time.Sleep(time.Second * 1)

	p1.Volume(0)
	p2.Volume(30)

	time.Sleep(time.Second * 1)

	p1.Volume(100)

	time.Sleep(time.Second * 8)
}
