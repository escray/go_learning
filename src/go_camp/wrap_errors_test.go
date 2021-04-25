package go_camp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		log.Println("unable to write:", err) // annotated error goes to log file
		return err	// unannotated error returned to caller
	}
	return nil
}

type Config struct {

}

func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		log.Printf("could not marshal config: %v", err )
		// oops, forgot to return
		return err
	}
	if err := WriteAll(w, buf); err != nil {
		log.Printf("could not write config: %v", err)
		return err
	}
	return nil
}

func TestWriteConfig(t *testing.T) {
	conf := Config{}
	f := * new(io.Writer)
	err := WriteConfig(f, &conf)
	fmt.Println(err)
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	return nil, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "could not read config")
}

func TestErrorsCause(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		fmt.Printf("origin error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}
}

func TestErrorsWithMessage(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	return errors.Wrap(err, "write failed")
}
