package file

import "os"

// Save writes data to dst.
func Save(dst string, data []byte) error {
	err := os.WriteFile(dst, data, 0666)
	if err != nil {
		return err
	}

	return nil
}
