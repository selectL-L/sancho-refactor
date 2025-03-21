package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileUtils contains utility functions for file operations
type FileUtils struct {
	DataDir string
}

// NewFileUtils creates a new FileUtils instance
func NewFileUtils(dataDir string) *FileUtils {
	return &FileUtils{
		DataDir: dataDir,
	}
}

// GetFilePath returns the full path for a file in the data directory
func (f *FileUtils) GetFilePath(filename string) string {
	return filepath.Join(f.DataDir, filename)
}

// ReadFile reads an entire file into memory
func ReadFile(filename string) ([]byte, error) {
	if !FileExists(filename) {
		return nil, fmt.Errorf("file does not exist: %s", filename)
	}

	return os.ReadFile(filename)
}

// WriteFile writes data to a file, creating it if it doesn't exist
func WriteFile(filename string, data []byte) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

// AppendToFile appends data to a file
func AppendToFile(filename string, data []byte) error {
	// Create file if it doesn't exist
	if !FileExists(filename) {
		return WriteFile(filename, data)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// AppendLineToFile appends a line to a file
func AppendLineToFile(filename, line string) error {
	// Make sure line ends with a newline
	if !strings.HasSuffix(line, "\n") {
		line += "\n"
	}

	return AppendToFile(filename, []byte(line))
}

// ReadLines reads all lines from a file
func ReadLines(filename string) ([]string, error) {
	data, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	// Remove empty lines
	var result []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}

	return result, nil
}

// ReadLinesWithPrefix reads lines that start with a specific prefix
func ReadLinesWithPrefix(filename, prefix string) ([]string, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			result = append(result, line)
		}
	}

	return result, nil
}

// ReadKeyValueFile reads a file with key-value pairs (one per line, separated by space or tab)
func ReadKeyValueFile(filename string) (map[string]string, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key != "" {
			result[key] = value
		}
	}

	return result, nil
}

// WriteKeyValueFile writes a map to a file as key-value pairs
func WriteKeyValueFile(filename string, data map[string]string) error {
	var builder strings.Builder

	for key, value := range data {
		builder.WriteString(key)
		builder.WriteString(" ")
		builder.WriteString(value)
		builder.WriteString("\n")
	}

	return WriteFile(filename, []byte(builder.String()))
}

// ReadJSONFile reads a JSON file into a struct
func ReadJSONFile(filename string, v interface{}) error {
	data, err := ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// WriteJSONFile writes a struct to a file as JSON
func WriteJSONFile(filename string, v interface{}, pretty bool) error {
	var data []byte
	var err error

	if pretty {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return WriteFile(filename, data)
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

// DirExists checks if a directory exists
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// EnsureDirExists creates a directory if it doesn't exist
func EnsureDirExists(dirname string) error {
	if !DirExists(dirname) {
		return os.MkdirAll(dirname, 0755)
	}

	return nil
}

// ListFiles lists all files in a directory
func ListFiles(dirname string) ([]string, error) {
	if !DirExists(dirname) {
		return nil, fmt.Errorf("directory does not exist: %s", dirname)
	}

	files, err := os.ReadDir(dirname)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var result []string
	for _, file := range files {
		if !file.IsDir() {
			result = append(result, file.Name())
		}
	}

	return result, nil
}

// BackupFile creates a backup of a file
func BackupFile(filename string) error {
	if !FileExists(filename) {
		return fmt.Errorf("file does not exist: %s", filename)
	}

	backupName := filename + ".bak"

	// Read original file
	data, err := ReadFile(filename)
	if err != nil {
		return err
	}

	// Write backup
	return WriteFile(backupName, data)
}

// ReplaceInFile replaces a string in a file
func ReplaceInFile(filename, old, new string) error {
	data, err := ReadFile(filename)
	if err != nil {
		return err
	}

	newData := strings.ReplaceAll(string(data), old, new)

	return WriteFile(filename, []byte(newData))
}

// CopyFile copies a file
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// DeleteLine deletes a line from a file
func DeleteLine(filename string, lineNum int) error {
	lines, err := ReadLines(filename)
	if err != nil {
		return err
	}

	if lineNum < 0 || lineNum >= len(lines) {
		return fmt.Errorf("line number out of range: %d", lineNum)
	}

	// Remove the line
	lines = append(lines[:lineNum], lines[lineNum+1:]...)

	// Write the file
	return WriteFile(filename, []byte(strings.Join(lines, "\n")))
}

// DeleteLinesWith deletes lines that contain a specific string
func DeleteLinesWith(filename, substring string) error {
	lines, err := ReadLines(filename)
	if err != nil {
		return err
	}

	var newLines []string
	for _, line := range lines {
		if !strings.Contains(line, substring) {
			newLines = append(newLines, line)
		}
	}

	// Write the file
	return WriteFile(filename, []byte(strings.Join(newLines, "\n")))
}

// ReadCSV reads a CSV file
func ReadCSV(filename string) ([][]string, error) {
	data, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var result [][]string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		for i, field := range fields {
			fields[i] = strings.TrimSpace(field)
		}

		result = append(result, fields)
	}

	return result, nil
}

// WriteCSV writes a CSV file
func WriteCSV(filename string, data [][]string) error {
	var builder strings.Builder

	for _, row := range data {
		for i, field := range row {
			if i > 0 {
				builder.WriteString(",")
			}
			builder.WriteString(field)
		}
		builder.WriteString("\n")
	}

	return WriteFile(filename, []byte(builder.String()))
}

// ReadConfig reads a configuration file
func (f *FileUtils) ReadConfig(filename string) (map[string]string, error) {
	path := f.GetFilePath(filename)
	return ReadKeyValueFile(path)
}

// WriteConfig writes a configuration file
func (f *FileUtils) WriteConfig(filename string, config map[string]string) error {
	path := f.GetFilePath(filename)
	return WriteKeyValueFile(path, config)
}

// ReadReminders reads reminder data from a file
func (f *FileUtils) ReadReminders(filename string) ([]string, error) {
	path := f.GetFilePath(filename)
	return ReadLines(path)
}

// WriteReminders writes reminder data to a file
func (f *FileUtils) WriteReminders(filename string, reminders []string) error {
	path := f.GetFilePath(filename)
	return WriteFile(path, []byte(strings.Join(reminders, "\n")))
}

// AppendReminder adds a reminder to the file
func (f *FileUtils) AppendReminder(filename, reminder string) error {
	path := f.GetFilePath(filename)
	return AppendLineToFile(path, reminder)
}

// DeleteReminder removes a reminder from the file
func (f *FileUtils) DeleteReminder(filename, reminderID string) error {
	path := f.GetFilePath(filename)
	return DeleteLinesWith(path, reminderID)
}

// ReadUserTimezones reads user timezone data
func (f *FileUtils) ReadUserTimezones(filename string) (map[string]string, error) {
	path := f.GetFilePath(filename)
	return ReadKeyValueFile(path)
}

// WriteUserTimezones writes user timezone data
func (f *FileUtils) WriteUserTimezones(filename string, timezones map[string]string) error {
	path := f.GetFilePath(filename)
	return WriteKeyValueFile(path, timezones)
}

// GetTokenFromFile reads an auth token from a file
func (f *FileUtils) GetTokenFromFile(filename string) (string, error) {
	path := f.GetFilePath(filename)

	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open token file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", errors.New("token file is empty")
	}

	return scanner.Text(), nil
}
