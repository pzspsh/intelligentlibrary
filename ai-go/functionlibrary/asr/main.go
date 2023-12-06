/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 14:48:03
*/
package main

import (
	"os"

	"github.com/liuxp0827/govpr"
	"github.com/liuxp0827/govpr/log"
	"github.com/liuxp0827/govpr/waveIO"
)

type engine struct {
	vprEngine *govpr.VPREngine
}

func NewEngine(sampleRate, delSilRange int, ubmFile, userModelFile string) *engine {
	vprengine, err := govpr.NewVPREngine(sampleRate, delSilRange, true, ubmFile, userModelFile)
	if err != nil {
		return nil
	}
	return &engine{
		vprEngine: vprengine,
	}
}

func (e *engine) DestroyEngine() {
	e.vprEngine = nil
}

func (e *engine) TrainSpeech(buffers [][]byte) error {

	var err error
	count := len(buffers)
	for i := 0; i < count; i++ {
		err = e.vprEngine.AddTrainBuffer(buffers[i])
		if err != nil {
			log.Error(err)
			return err
		}
	}

	defer e.vprEngine.ClearTrainBuffer()
	defer e.vprEngine.ClearAllBuffer()

	err = e.vprEngine.TrainModel()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (e *engine) RecSpeech(buffer []byte) error {

	err := e.vprEngine.AddVerifyBuffer(buffer)
	defer e.vprEngine.ClearVerifyBuffer()
	if err != nil {
		log.Error(err)
		return err
	}

	err = e.vprEngine.VerifyModel()
	if err != nil {
		log.Error(err)
		return err
	}

	Score := e.vprEngine.GetScore()
	log.Infof("vpr score: %f", Score)
	return nil
}

func main() {
	log.SetLevel(log.LevelDebug)

	vprEngine := NewEngine(16000, 50, "../ubm/ubm", "model/test.dat")
	trainlist := []string{
		"wav/train/01_32468975.wav",
		"wav/train/02_58769423.wav",
		"wav/train/03_59682734.wav",
		"wav/train/04_64958273.wav",
		"wav/train/05_65432978.wav",
	}

	trainBuffer := make([][]byte, 0)

	for _, file := range trainlist {
		buf, err := loadWaveData(file)
		if err != nil {
			log.Error(err)
			return
		}
		trainBuffer = append(trainBuffer, buf)
	}

	verifyBuffer, err := waveIO.WaveLoad("wav/verify/34986527.wav")
	if err != nil {
		log.Error(err)
		return
	}

	vprEngine.TrainSpeech(trainBuffer)
	vprEngine.RecSpeech(verifyBuffer)
}

func loadWaveData(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// remove .wav header info 44 bits
	data = data[44:]
	return data, nil
}
