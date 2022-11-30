package image

import (
	"bytes"
	"sync"
)

// ImageStore is an interface to store laptop images
type ImageStore interface {
	Save(string, string, bytes.Buffer) (string, error)
}

// ImageInfo contains information of the laptop image
type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

// DiskImageStore stores image on disk, and its info on memory
type diskImageStore struct {
	mutex       sync.Mutex
	imageFolder string
	images      map[string]*ImageInfo
}

func NewDiskImageStore(imageFolder string) ImageStore {
	return &diskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}
