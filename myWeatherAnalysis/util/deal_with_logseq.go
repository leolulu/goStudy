package util

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var LogseqBaseDirPath = `C:\坚果云\Logseq数据文件夹\new-personal\journals`

func FindLogseqJournalsPath(journalsDir string) (journalFilesPath []string, err error) {
	entries, err := os.ReadDir(journalsDir)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}
	slices.Reverse(entries)
	for _, entry := range entries[:3] {
		if !entry.IsDir() {
			abs, err := filepath.Abs(filepath.Join(journalsDir, entry.Name()))
			if err != nil {
				log.Println("Error finding absolute path:", err)
				continue
			}
			journalFilesPath = append(journalFilesPath, abs)
		}
	}
	return journalFilesPath, nil
}

func ReadJournalContent(journalFilePath string) (content string, err error) {
	byteContent, err := os.ReadFile(journalFilePath)
	if err != nil {
		return
	}
	return string(byteContent), nil
}

func GenJournalContentOfLastThreeDays(journalsDir string) (resultContentString string, err error) {
	journalsPath, err := FindLogseqJournalsPath(journalsDir)
	if err != nil {
		return "", err
	}
	resultContent := make([]string, len(journalsPath))
	for _, path := range journalsPath {
		content, err := ReadJournalContent(path)
		if err != nil {
			return "", err
		}
		resultContent = append(resultContent, filepath.Base(path))
		resultContent = append(resultContent, content)
		resultContent = append(resultContent, "\n\n")
	}
	return strings.TrimSpace(strings.Join(resultContent, "\n")), nil
}
