package seedloader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrInitializingSeeder = errors.New("failed to initialize seeder, ensure service name matches mod package name")
	ErrOpeningSeed        = errors.New("failed to open seed file. probably file not found or permission err")
	ErrReadingSeed        = errors.New("failed to read seed file. ensure it's not password protected or corrupt")
)

type S struct {
	rootDir string
}

func NewSeedLoader(serviceName, seedsPath string) (S, error) {
	p, err := build.Default.Import(fmt.Sprintf("%s/%s", serviceName, seedsPath), "", build.FindOnly)
	if err != nil {
		fmt.Print(p.Dir, " Dir\n")
		return S{}, ErrInitializingSeeder
	}

	fmt.Print(p.Dir)
	return S{
		rootDir: p.Dir,
	}, nil
}

func (s S) GetSeed(file string) ([]byte, error) {
	seedFile := filepath.Join(s.rootDir, file)
	f, err := os.OpenFile(seedFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, ErrOpeningSeed
	}
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, ErrReadingSeed
	}
	return buff, nil
}

func (s S) ParseSeed(file string, output interface{}) error {
	seed, err := s.GetSeed(file)
	if err != nil {
		return err
	}

	return UnPack(seed, output)
}

func UnPack(in interface{}, target interface{}) error {
	var e1 error
	var b []byte
	switch in := in.(type) {
	case []byte:
		b = in
	// Do something.
	default:
		// Do the rest.
		b, e1 = json.Marshal(in)
		if e1 != nil {
			return e1
		}
	}

	buf := bytes.NewBuffer(b)
	enc := json.NewDecoder(buf)
	enc.UseNumber()
	if err := enc.Decode(&target); err != nil {
		return err
	}
	return nil
}
