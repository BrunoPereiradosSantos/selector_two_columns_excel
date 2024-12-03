package handlers

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetDownloadsDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Não foi possível encontrar o diretório do usuário")
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(homeDir, "Downloads"), nil
	case "darwin":
		return filepath.Join(homeDir, "Downloads"), nil
	case "linux":
		cmd := exec.Command("xdg-user-dir", "DOWNLOAD")
		output, err := cmd.Output()
		if err != nil {
			return filepath.Join(homeDir, "Downloads"), nil
		}
		return strings.TrimSpace(string(output)), nil
	default:
		return "", errors.New("Sistema operacional não suportado")
	}
}

func SaveFileToDownloads(filename string, data []byte) (string, error) {
	downloadsDir, err := GetDownloadsDir()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(downloadsDir); os.IsNotExist(err) {
		err := os.MkdirAll(downloadsDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("erro ao criar o diretório Downloads: %v", err)
		}
	}

	outputPath := filepath.Join(downloadsDir, filename)

	err = os.WriteFile(outputPath, data, 0o644)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar o arquivo no diretório Downloads: %v", err)
	}

	return outputPath, nil
}
