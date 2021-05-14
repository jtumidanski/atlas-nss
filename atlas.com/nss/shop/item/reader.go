package item

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

func readDataDirectory(l logrus.FieldLogger, d string) ([]JSONObject, error) {
	f, err := os.Open(d)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, errors.New("data directory provided is a file not a directory")
	}

	fs, err := ioutil.ReadDir(d)
	if err != nil {
		return nil, err
	}

	var items []JSONObject
	for _, cf := range fs {
		l.Infof("Found %s for parsing.", cf.Name())
		v, err := readDataFile(l, d+"/"+cf.Name())
		if err != nil {
			return nil, err
		}
		items = append(items, v...)
	}
	return items, nil
}

func readDataFile(l logrus.FieldLogger, p string) ([]JSONObject, error) {
	l.Debugf("Reading %s.", p)
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	holder := struct {
		Items []JSONObject `json:"items"`
	}{}

	var v = &holder
	err = fromJSON(v, f)
	if err != nil {
		return nil, err
	}
	return v.Items, err
}

// fromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func fromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
