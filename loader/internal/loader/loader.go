package loader

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	ErrInvalidURL       = errors.New("invalid URL, try again with the correct URL")
	ErrConnectionFailed = errors.New("connection failed, check your internet connection")
	ErrDownloadFailed   = errors.New("file download failed, check the availability of the file")
	ErrFileNotFound     = errors.New("file not found on the server")
)

func DownloadFile(fileURL, fileName string) error {
	if fileURL == "" {
		return ErrInvalidURL
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		return ErrInvalidURL
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := client.Do(req)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) {
			return ErrConnectionFailed
		}
		return fmt.Errorf("error in requesting file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return ErrFileNotFound
		}
		return ErrDownloadFailed
	}

	file, err := os.Create(fileName)
	if err != nil {
		return ErrDownloadFailed
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return io.ErrUnexpectedEOF
	}

	return nil
}
