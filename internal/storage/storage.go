package storage

import (
	"fmt"
	"io"
	"machine"
	"macro-keyboard/internal/buttons"
	"os"
	"strings"

	"tinygo.org/x/tinyfs/littlefs"
)

type Flash struct {
	filesystem *littlefs.LFS
}

func SanitizeString(s string) (out string) {
	out = strings.ReplaceAll(s, "\t", "")
	out = strings.ReplaceAll(out, "\n", "")
	out = strings.ReplaceAll(out, "  ", "")
	return out
}

func New(btn []buttons.Button, format bool) Flash {
	f := Flash{
		filesystem: littlefs.New(machine.Flash),
	}
	err := f.Start()
	if err != nil || format == true {
		fmt.Printf("error while mounting: %v\nresetting flash to default state...\n", err)
		f.Reset()
		for idx := range btn {
			f.WriteButton(&btn[idx])
		}
	}
	return f
}

func (f *Flash) Start() error {
	f.filesystem.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	err := f.filesystem.Mount()
	return err
}

func (f *Flash) Stop() error {
	err := f.filesystem.Unmount()
	return err
}

func (f *Flash) Reset() {
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

func (f *Flash) Write(filename, data string) {
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

func (f *Flash) FilesList() (filenames []string) {
	path := "/"
	dir, err := f.filesystem.Open(path)
	if err != nil {
		fmt.Printf("could not open directory %s: %v\n", path, err)
		return filenames
	}
	defer dir.Close()
	infos, err := dir.Readdir(0)
	if err != nil {
		fmt.Printf("could not read directory %s: %v\n", path, err)
		return filenames
	}
	for _, info := range infos {
		filenames = append(filenames, info.Name())
	}
	return filenames
}

func (f *Flash) Read(filename string) string {
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

func (f *Flash) ReadAll() (files []string) {
	filenames := f.FilesList()
	for idx := range filenames {
		files = append(files, f.Read(filenames[idx]))
	}
	return files
}

func (f *Flash) WriteButton(button *buttons.Button) {
	filename := button.Name
	fmt.Printf("\nWriting button %s to storage:\n", filename)
	fmt.Printf("%v: %s ", filename, button.String())
	f.Write(filename, button.String())
}

func (f *Flash) ReadButton(btn *buttons.Button) {
	fmt.Printf("\nReading button %s from storage...\n", btn.Name)
	data := f.Read(btn.Name)
	fmt.Println(data)
	buttons.ParseButton(data, btn)
	return
}
