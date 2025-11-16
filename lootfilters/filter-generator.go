package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// Windows API constants and functions
const (
	FOREGROUND_BLUE      = 0x0001
	FOREGROUND_GREEN     = 0x0002
	FOREGROUND_RED       = 0x0004
	FOREGROUND_INTENSITY = 0x0008
	BACKGROUND_BLUE      = 0x0010
	BACKGROUND_GREEN     = 0x0020
	BACKGROUND_RED       = 0x0040
	BACKGROUND_INTENSITY = 0x0080
)

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
)

// Color constants for Windows
const (
	ColorDefault = FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_BLUE
	
	ColorRed     = FOREGROUND_RED | FOREGROUND_INTENSITY
	ColorGreen   = FOREGROUND_GREEN | FOREGROUND_INTENSITY
	ColorYellow  = FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_INTENSITY
	ColorBlue    = FOREGROUND_BLUE | FOREGROUND_INTENSITY
	ColorPurple  = FOREGROUND_RED | FOREGROUND_BLUE | FOREGROUND_INTENSITY
	ColorCyan    = FOREGROUND_GREEN | FOREGROUND_BLUE | FOREGROUND_INTENSITY
	ColorWhite   = FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_BLUE | FOREGROUND_INTENSITY
	
	// Backgrounds
	BgRed    = BACKGROUND_RED | BACKGROUND_INTENSITY
	BgGreen  = BACKGROUND_GREEN
	BgYellow = BACKGROUND_RED | BACKGROUND_GREEN | BACKGROUND_INTENSITY
)

// Color functions
func setColor(color uint16) {
	handle, _ := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	procSetConsoleTextAttribute.Call(uintptr(handle), uintptr(color))
}

func resetColor() {
	setColor(ColorDefault)
}

func colorPrint(color uint16, text string) {
	setColor(color)
	fmt.Print(text)
	resetColor()
}

func colorPrintf(color uint16, format string, args ...interface{}) {
	setColor(color)
	fmt.Printf(format, args...)
	resetColor()
}

func colorPrintln(color uint16, text string) {
	setColor(color)
	fmt.Println(text)
	resetColor()
}

// Helper functions for specific colors
func printRed(text string) {
	colorPrint(ColorRed, text)
}

func printGreen(text string) {
	colorPrint(ColorGreen, text)
}

func printYellow(text string) {
	colorPrint(ColorYellow, text)
}

func printCyan(text string) {
	colorPrint(ColorCyan, text)
}

func printPurple(text string) {
	colorPrint(ColorPurple, text)
}

func printBgRed(text string) {
	colorPrint(BgRed | FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_BLUE | FOREGROUND_INTENSITY, text)
}

func printfRed(format string, args ...interface{}) {
	colorPrintf(ColorRed, format, args...)
}

func printfGreen(format string, args ...interface{}) {
	colorPrintf(ColorGreen, format, args...)
}

func printfYellow(format string, args ...interface{}) {
	colorPrintf(ColorYellow, format, args...)
}

func printfCyan(format string, args ...interface{}) {
	colorPrintf(ColorCyan, format, args...)
}

func printfPurple(format string, args ...interface{}) {
	colorPrintf(ColorPurple, format, args...)
}

func printfBgRed(format string, args ...interface{}) {
	colorPrintf(BgRed | FOREGROUND_RED | FOREGROUND_GREEN | FOREGROUND_BLUE | FOREGROUND_INTENSITY, format, args...)
}

// TagMap defines which tags to exclude for each filter type
var TagMap = map[string][]string{
	"summoner-nossf": {"filter-rarity", "filter-ssf", "filter-ssf-mage"},
	"summoner-ssf":   {"filter-rarity", "filter-nossf", "filter-ssf-mage", "filter-mage"},
	"mage-nossf":     {"filter-rarity", "filter-ssf", "filter-ssf-summoner"},
	"mage-ssf":       {"filter-rarity", "filter-nossf", "filter-ssf-summoner", "filter-summoner"},
	"rarity":         {"filter-ssf", "filter-summoner", "filter-mage", "filter-norarity"},
}

// Global variables
var debugMode bool = false
var localMode bool = true

func main() {
	fmt.Println("Filter generator started")
	fmt.Println("Enter 'exit' to quit, 'debug' to toggle debug mode, 'local' to toggle local mode")
	
	// Main application loop
	for {
		// Show selection menu
		showMenu()
		
		// Read user choice
		choice := readChoice()
		
		// Process commands
		switch choice {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "debug":
			debugMode = !debugMode
			if debugMode {
				printCyan("Debug mode ENABLED\n")
			} else {
				printCyan("Debug mode DISABLED\n")
			}
			continue
		case "local":
			localMode = !localMode
			if localMode {
				printCyan("Local mode ENABLED\n")
			} else {
				printCyan("Local mode DISABLED\n")
			}
			continue
		case "invalid":
			printRed("Invalid command. Exiting...\n")
			return
		case "all":
			// Generate all files
			generateAllFiles()
		default:
			// Process filter selection
			success, filteredBlocks := processChoice(choice)
			printGenerationResult(choice, success, filteredBlocks)
		}
	}
}

func generateAllFiles() {
	fmt.Println("\nGenerating all files...")
	
	// List of all variants for generation
	variants := []string{"summoner-nossf", "summoner-ssf", "mage-nossf", "mage-ssf", "rarity"}
	
	for _, variant := range variants {
		success, filteredBlocks := processChoice(variant)
		printGenerationResult(variant, success, filteredBlocks)
	}
}

func showMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Select filter type:")
	fmt.Println("0. All files")
	fmt.Println("1. summoner-nossf")
	fmt.Println("2. summoner-ssf") 
	fmt.Println("3. mage-nossf")
	fmt.Println("4. mage-ssf")
	fmt.Println("5. rarity")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  debug - toggle debug mode")
	fmt.Println("  local - toggle local mode")
	fmt.Println("  exit - quit program")
	
	// Show current debug status
	if debugMode {
		printCyan("\nDebug mode: ON\n")
	} else {
		printCyan("\nDebug mode: OFF\n")
	}
	
	// Show current local mode status
	if localMode {
		username := os.Getenv("USERNAME")
		poe2Path := fmt.Sprintf("C:\\Users\\%s\\Documents\\My Games\\Path of Exile 2", username)
		printfYellow("Local mode: ON\n")
		printfYellow("When generating lootfilters, they will also be updated in your Path of Exile 2 client,\ndirectory: %s\n", poe2Path)
	} else {
		printCyan("Local mode: OFF\n")
	}
}

func readChoice() string {
	fmt.Print("\nEnter number (0-5) or command: ")
	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		printfRed("Input read error: %v\n", err)
		return "invalid"
	}
	
	// Remove extra spaces and newlines
	input = strings.TrimSpace(input)
	
	// Check commands
	switch strings.ToLower(input) {
	case "exit", "quit", "q":
		return "exit"
	case "debug", "d":
		return "debug"
	case "local", "l":
		return "local"
	case "0":
		return "all"
	case "1":
		return "summoner-nossf"
	case "2":
		return "summoner-ssf"
	case "3":
		return "mage-nossf"
	case "4":
		return "mage-ssf"
	case "5":
		return "rarity"
	default:
		// If unsupported string entered - exit
		return "invalid"
	}
}

func processChoice(choice string) (bool, [][]string) {
	debugPrintfCyan("\nSelected: %s\n", choice)
	debugPrintfCyan("Excluded tags: %v\n", TagMap[choice])
	
	// 1. Read overall.filter content
	content, err := readFile("overall.filter")
	if err != nil {
		printfBgRed("File read error: %v\n", err)
		return false, nil
	}
	
	// 2. Split into rule blocks
	blocks := splitIntoBlocks(content)
	debugPrintfCyan("Found blocks of rules AND comments: %d\n", len(blocks))
	
	// 3. Filter blocks by tags
	filteredBlocks := filterBlocks(blocks, choice)
	debugPrintfCyan("Blocks of rules AND comments after filtering: %d\n", len(filteredBlocks))
	
	// 4. Count informative blocks after filtering
	informativeBlocks := countInformativeBlocks(filteredBlocks)
	debugPrintfCyan("Blocks of rules after filtering: %d\n", informativeBlocks)
	
	// 5. Check generation success (at least one informative block)
	if informativeBlocks == 0 {
		printRed("No suitable informative blocks found for filter\n")
		return false, filteredBlocks
	}
	
	return true, filteredBlocks
}

// countInformativeBlocks counts only blocks that contain Show/Hide commands
func countInformativeBlocks(blocks [][]string) int {
	count := 0
	for _, block := range blocks {
		if isInformativeBlock(block) {
			count++
		}
	}
	return count
}

// isInformativeBlock checks if block contains Show or Hide command
func isInformativeBlock(block []string) bool {
	for _, line := range block {
		trimmedLine := strings.TrimSpace(line)
		// Skip comment lines
		if strings.HasPrefix(trimmedLine, "#") {
			continue
		}
		// Check if line starts with Show or Hide
		if strings.HasPrefix(trimmedLine, "Show") || strings.HasPrefix(trimmedLine, "Hide") {
			return true
		}
	}
	return false
}

func printGenerationResult(choice string, success bool, filteredBlocks [][]string) {
	outputFile := "fps30_" + choice + ".filter"
	
	// Success counters
	currentSuccess := false
	localSuccess := false
	var currentBlocks int = 0
	var localBlocks int = 0
	
	// Generate to current directory
	currentPath := outputFile
	if success {
		err := writeOutputFile(currentPath, filteredBlocks)
		if err != nil {
			currentSuccess = false
		} else {
			// Verify file was actually created and not empty
			if fileExists(currentPath) {
				content, err := readFile(currentPath)
				if err == nil {
					blocks := splitIntoBlocks(content)
					informativeBlocks := countInformativeBlocks(blocks)
					if informativeBlocks > 0 {
						currentSuccess = true
						currentBlocks = informativeBlocks
					}
				}
			}
		}
	}
	
	// Generate to local Path of Exile 2 directory
	localPath := ""
	if localMode && success {
		username := os.Getenv("USERNAME")
		localPath = fmt.Sprintf("C:\\Users\\%s\\Documents\\My Games\\Path of Exile 2\\%s", username, outputFile)
		
		// Create directory if it doesn't exist
		dir := filepath.Dir(localPath)
		os.MkdirAll(dir, 0755)
		
		err := writeOutputFile(localPath, filteredBlocks)
		if err != nil {
			localSuccess = false
		} else {
			// Verify file was actually created and not empty
			if fileExists(localPath) {
				content, err := readFile(localPath)
				if err == nil {
					blocks := splitIntoBlocks(content)
					informativeBlocks := countInformativeBlocks(blocks)
					if informativeBlocks > 0 {
						localSuccess = true
						localBlocks = informativeBlocks
					}
				}
			}
		}
	}
	
	// Output result
	if !success {
		// Failed to generate any files
		printfRed("[-] %s - file was not generated (no matching blocks found)\n", outputFile)
		fmt.Println()
	} else if localMode && currentSuccess && localSuccess {
		// Both files generated successfully
		printfGreen("[++] %s (rules: %d)\n", outputFile, currentBlocks)
	} else if !localMode && currentSuccess {
		// Only current directory (local mode disabled)
		printfGreen("[+] %s (rules: %d)\n", outputFile, currentBlocks)
	} else {
		// Mixed result - output separately with full paths
		if currentSuccess {
			printfGreen("[+] %s (rules: %d)\n", currentPath, currentBlocks)
		} else {
			printfRed("[-] %s - generation error\n", currentPath)
		}
		
		if localMode {
			if localSuccess {
				printfGreen("[+] %s (rules: %d)\n", localPath, localBlocks)
			} else {
				printfRed("[-] %s - generation error\n", localPath)
			}
		}
		fmt.Println()
	}
}

// debugPrintf outputs message only in debug mode
func debugPrintfCyan(format string, args ...interface{}) {
	if debugMode {
		printfCyan(format, args...)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func splitIntoBlocks(content string) [][]string {
	var blocks [][]string
	var currentBlock []string
	
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		
		// If empty line or line with only spaces - end current block
		if trimmedLine == "" {
			if len(currentBlock) > 0 {
				blocks = append(blocks, currentBlock)
				currentBlock = nil
			}
			continue
		}
		
		// Add line to current block
		currentBlock = append(currentBlock, line)
	}
	
	// Add last block if not empty
	if len(currentBlock) > 0 {
		blocks = append(blocks, currentBlock)
	}
	
	return blocks
}

func filterBlocks(blocks [][]string, choice string) [][]string {
	tagsToExclude := TagMap[choice]
	var filteredBlocks [][]string
	
	for _, block := range blocks {
		if len(block) == 0 {
			continue
		}
		
		firstLine := block[0]
		trimmedFirstLine := strings.TrimSpace(firstLine)
		
		// Check if first line starts with "#"
		if strings.HasPrefix(trimmedFirstLine, "#") {
			// Look for tags to exclude
			shouldExclude := false
			for _, tag := range tagsToExclude {
				if strings.Contains(trimmedFirstLine, tag) {
					shouldExclude = true
					debugPrintfCyan("Excluded block with tag '%s': %s\n", tag, getFirstLinePreview(trimmedFirstLine))
					break
				}
			}
			
			// If found at least one excluding tag - skip block
			if shouldExclude {
				continue
			}
		}
		
		// If not excluded - add block to result
		filteredBlocks = append(filteredBlocks, block)
	}
	
	return filteredBlocks
}

// Helper function for pretty output of excluded blocks
func getFirstLinePreview(line string) string {
	if len(line) > 50 {
		return line[:50] + "..."
	}
	return line
}

func writeOutputFile(filename string, blocks [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	
	for i, block := range blocks {
		// Write all block lines
		for _, line := range block {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
		
		// Add empty line between blocks (except last one)
		if i < len(blocks)-1 {
			_, err := writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	
	return writer.Flush()
}