package storage

import (
	"fmt"
	"io"
	"machine"
	"macro-keyboard/internal/buttons"
	"os"

	"tinygo.org/x/tinyfs/littlefs"
)

type Flash struct {
	filesystem *littlefs.LFS
}

func New(btn []buttons.Button, format bool) Flash {
	f := Flash{
		filesystem: littlefs.New(machine.Flash),
	}
	err := f.start()
	if err != nil || format == true {
		fmt.Printf("error while mounting: %v\nresetting flash to default state...\n", err)
		f.reset()
		for idx := range btn {
			f.WriteButton(&btn[idx])
		}
	}
	return f
}

func (f *Flash) Stop() error {
	err := f.filesystem.Unmount()
	return err
}

func (f *Flash) WriteButton(button *buttons.Button) {
	filename := button.Name
	fmt.Printf("\nWriting button %s to storage:\n", filename)
	fmt.Printf("%v: %s ", filename, button.String())
	f.write(filename, button.String())
}

func (f *Flash) ReadButton(btn *buttons.Button) {
	fmt.Printf("\nReading button %s from storage...\n", btn.Name)
	data := f.read(btn.Name)
	buttons.ParseButton(data, btn)
	return
}

/* internal functions */

func (f *Flash) start() error {
	f.filesystem.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	err := f.filesystem.Mount()
	return err
}

func (f *Flash) reset() {
	err := f.filesystem.Format()
	if err != nil {
		fmt.Printf("error while formatting: %v\n", err)
		return
	}
	err = f.filesystem.Mount()
	if err != nil {
		fmt.Printf("error while mounting: %v\nrebooting...\n\n", err)
		machine.CPUReset()
	}
}

func (f *Flash) write(filename, data string) {
	file, err := f.filesystem.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		fmt.Printf("error opening %s: %s\n", filename, err.Error())
		return
	}
	defer file.Close()
	bytes, err := file.Write([]byte(data))
	if err != nil {
		fmt.Printf("error writing %s: %s\n", filename, err.Error())
		return
	}
	fmt.Printf("written %v bytes.\n", bytes)
}

func (f *Flash) read(filename string) string {
	info, err := f.filesystem.Stat(filename)
	if err != nil {
		fmt.Printf("error getting info from %s: %s\n", filename, err.Error())
		return ""
	}
	file, err := f.filesystem.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s: %s\n", filename, err.Error())
		return ""
	}
	defer file.Close()
	buf := make([]byte, info.Size()+1)
	for {
		_, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading %s: %s\n", filename, err.Error())
		}
	}
	return string(buf[:len(buf)-1])
}
