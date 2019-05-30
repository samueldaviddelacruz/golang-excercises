package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/samueldaviddelacruz/golang-exercises/secret_api_cli/cipher"
)

func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}
func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&v.keyValues)
}
func (v *Vault) load() error {
	file, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer file.Close()
	r, err := cipher.DecryptReader(v.encodingKey, file)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)

}
func (v *Vault) save() error {
	file, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {

		return err
	}
	defer file.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, file)
	if err != nil {

		return err
	}
	return v.writeKeyValues(w)

	return nil
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	if err != nil {
		return err
	}
	return nil
}
