package audio

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func PlayWavAudio(body []byte) {
	fmt.Println("Playing audio...")

	// read wav stream
	r := bytes.NewReader(body)
	rc := ioutil.NopCloser(r)

	// decode wav stream
	streamer, format, err := wav.Decode(rc)
	if err != nil {
		panic(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// play wav stream
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
