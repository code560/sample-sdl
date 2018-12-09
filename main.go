package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/mix"
)

func main() {
	test()
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
