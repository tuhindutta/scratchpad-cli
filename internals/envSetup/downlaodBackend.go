package envSetup

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadBackend(url, destDir string) error {
	// 2. Fetch the ZIP file from GitHub
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad HTTP status: %s", resp.Status)
	}

	// 3. Create a temporary file to save the streaming ZIP data
	tmpFile, err := os.CreateTemp("", "github-*.zip")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temp zip file when done
	defer tmpFile.Close()

	// Stream the download response body into the temporary file
	if _, err = io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("failed to save zip contents: %w", err)
	}

	// Rewind the temp file pointer to the beginning before reading it as a ZIP
	if _, err = tmpFile.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek temp file: %w", err)
	}

	// Get file info to pass its size to the zip reader
	stat, err := tmpFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat temp file: %w", err)
	}

	// 4. Initialize the ZIP archive reader
	zipReader, err := zip.NewReader(tmpFile, stat.Size())
	if err != nil {
		return fmt.Errorf("failed to open zip reader: %w", err)
	}

	// 5. Iterate through files and folders in the ZIP archive
	for _, file := range zipReader.File {
		// Prevent Zip Slip vulnerability by cleaning the file path
		cleanedPath := filepath.Clean(file.Name)
		if strings.HasPrefix(cleanedPath, "..") || strings.HasPrefix(cleanedPath, "/") {
			return fmt.Errorf("illegal file path detected in zip: %s", file.Name)
		}

		// Construct the final extraction path on the local drive
		targetPath := filepath.Join(destDir, cleanedPath)

		// Create parent directories if this is a file, or create the directory itself
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(targetPath, file.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return fmt.Errorf("failed to create parent directories for %s: %w", targetPath, err)
		}

		// Open the compressed file inside the ZIP archive
		zipFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zipped file %s: %w", file.Name, err)
		}

		// Create the destination file on disk
		destFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			zipFile.Close()
			return fmt.Errorf("failed to create target file %s: %w", targetPath, err)
		}

		// Uncompress and copy the content to disk
		_, err = io.Copy(destFile, zipFile)
		zipFile.Close()
		destFile.Close()
		if err != nil {
			return fmt.Errorf("failed to extract file content for %s: %w", targetPath, err)
		}
	}

	return nil
}
