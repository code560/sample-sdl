package player

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
)

var (
	SDL_CH          = 2
	SDL_SAMPLE_RATE = 44100
	SDL_BUFFER      = 1024
)

type Player interface {
	Close()
	Play(file string)
}

func GetPlayer(tag string) Player {
	return &imple_player{
		tag: tag,
	}
}

// imple_player

type imple_player struct {
	tag string
	sdl *sdl_player
}

func (i *imple_player) Close() {
	i.sdl.Close()
}

func (i *imple_player) Play(file string) {
	i.sdl.Play(file)
}

// sdl_player

func getInstance() *sdl_player {
	return instance_player
}

var instance_player *sdl_player = newSdlPlayer()

type sdl_player struct {
}

func newSdlPlayer() *sdl_player {
	p := &sdl_player{}
	p.init()
	return p
}

func (p *sdl_player) init() {
	err := mix.Init(mix.INIT_MOD)
	if err != nil {
		log.Fatal(err)
	}
	p.open()
}

func (p *sdl_player) open() {
	err := mix.OpenAudio(SDL_SAMPLE_RATE,
		mix.DEFAULT_FORMAT, SDL_CH, SDL_BUFFER)
	if err != nil {
		log.Fatal(err)
	}
	mix.AllocateChannels(20)
}

func (p *sdl_player) Close() {
	mix.CloseAudio()
	mix.Quit()
}

func (p *sdl_player) Play(file string) {
	chunk := p.load(file)
	if chunk == nil {
		return
	}

	_, err := chunk.Play(-1, 0)
	if err != nil {
		log.Print(err)
		return
	}
	defer chunk.Free()
}

func (p *sdl_player) load(file string) *mix.Chunk {
	chunk, err := mix.LoadWAV(file)
	if err != nil {
		log.Print(err)
		return nil
	}
	return chunk
}
