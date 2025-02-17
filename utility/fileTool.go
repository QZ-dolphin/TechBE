package utility

// fileTool.go
import (
	"fmt"
	"os"
	"path/filepath"
)

func ClearDir(dirPath string) error {
	// 打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		return fmt.Errorf("failed to open directory: %w", err)
	}
	defer dir.Close()

	// 读取目录内容
	files, err := dir.Readdirnames(-1)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// 遍历目录中的每个文件，并删除它
	for _, file := range files {
		filePath := filepath.Join(dirPath, file)
		info, err := os.Lstat(filePath)
		if err != nil {
			return fmt.Errorf("failed to get file info: %w", err)
		}

		if info.IsDir() {
			// 如果是目录，可以跳过或选择递归删除（如果需要的话）
			continue
		}

		// 删除文件
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to remove file: %w", err)
		}
	}

	return nil
}
