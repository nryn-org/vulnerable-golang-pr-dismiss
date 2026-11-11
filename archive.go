package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/ulikunitz/xz"
	"gopkg.in/yaml.v2"
)

// unpackFixture decompresses an xz archive and returns a request ID,
// the archive checksum, and the parsed manifest.
func unpackFixture(path string) (string, string, map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", nil, err
	}
	defer f.Close()

	r, err := xz.NewReader(f)
	if err != nil {
		return "", "", nil, err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return "", "", nil, err
	}

	sum := md5.Sum(data)
	manifest := map[string]string{}
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return "", "", nil, err
	}
	return uuid.NewV4().String(), hex.EncodeToString(sum[:]), manifest, nil
}
