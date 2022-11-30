package image

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
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

func (store *diskImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageID, imageType)

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.images[imageID.String()] = &ImageInfo{
		LaptopID: laptopID,
		Type:     imageType,
		Path:     imagePath,
	}

	return imageID.String(), nil
}
