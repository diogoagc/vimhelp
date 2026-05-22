package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed bindings.json
var defaultBindingsFile []byte

type Binding struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Topic       string `json:"topic"`
	Example     string `json:"example"`
}

func main() {
	if len(os.Args) > 1 {
		query := strings.ToLower(os.Args[1])

		if query == "--help" || query == "-h" {
			printHelp()
			return
		}

		if query == "init" {
			initUserBindings()
			return
		}

		if query == "config" || query == "--config" {
			printConfigPath()
			return
		}
	}

	bindings, err := loadBindings()
	if err != nil {
		fmt.Println("Error loading bindings:", err)
		return
	}

	if len(os.Args) == 1 {
		printAll(bindings)
		return
	}

	query := strings.ToLower(os.Args[1])

	if query == "topics" || query == "--topics" {
		printTopics(bindings)
		return
	}

	if printTopic(bindings, query) {
		return
	}

	if printKey(bindings, query) {
		return
	}

	fmt.Println("No binding or topic found:", os.Args[1])
}

func loadBindings() ([]Binding, error) {
	content := defaultBindingsFile

	userBindingsPath, err := getUserBindingsPath()
	if err == nil {
		userContent, readErr := os.ReadFile(userBindingsPath)
		if readErr == nil {
			content = userContent
		}
	}

	var bindings []Binding

	err = json.Unmarshal(content, &bindings)
	if err != nil {
		return nil, err
	}

	return bindings, nil
}

func getUserBindingsPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "vimhelp", "bindings.json"), nil
}

func initUserBindings() {
	userBindingsPath, err := getUserBindingsPath()
	if err != nil {
		fmt.Println("Error finding config directory:", err)
		return
	}

	if _, err := os.Stat(userBindingsPath); err == nil {
		fmt.Println("User bindings already exist:")
		fmt.Println(userBindingsPath)
		return
	}

	err = os.MkdirAll(filepath.Dir(userBindingsPath), 0755)
	if err != nil {
		fmt.Println("Error creating config directory:", err)
		return
	}

	err = os.WriteFile(userBindingsPath, defaultBindingsFile, 0644)
	if err != nil {
		fmt.Println("Error creating user bindings file:", err)
		return
	}

	fmt.Println("User bindings created:")
	fmt.Println(userBindingsPath)
	fmt.Println()
	fmt.Println("You can now edit this file to customize your Vim cheat sheet.")
}

func printConfigPath() {
	userBindingsPath, err := getUserBindingsPath()
	if err != nil {
		fmt.Println("Error finding config directory:", err)
		return
	}

	fmt.Println(userBindingsPath)
}

func printAll(bindings []Binding) {
	topics := getTopics(bindings)

	for index, topic := range topics {
		if index > 0 {
			fmt.Println()
		}

		printTitle(topic)

		for _, binding := range bindings {
			if binding.Topic == topic {
				printBinding(binding)
			}
		}
	}
}

func printTopic(bindings []Binding, topic string) bool {
	found := false

	for _, binding := range bindings {
		if strings.ToLower(binding.Topic) == topic {
			if !found {
				printTitle(topic)
			}

			printBinding(binding)
			found = true
		}
	}

	return found
}

func printKey(bindings []Binding, key string) bool {
	for _, binding := range bindings {
		if strings.ToLower(binding.Key) == key {
			printTitle(binding.Topic)
			printBinding(binding)
			return true
		}
	}

	return false
}

func printHelp() {
	fmt.Println("vimhelp - simple Vim bindings cheat sheet")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  vimhelp              Show all bindings grouped by topic")
	fmt.Println("  vimhelp <topic>      Show bindings for a topic")
	fmt.Println("  vimhelp <key>        Show help for a specific key")
	fmt.Println("  vimhelp topics       List available topics")
	fmt.Println("  vimhelp init         Create editable user bindings file")
	fmt.Println("  vimhelp config       Show user bindings file path")
	fmt.Println("  vimhelp --help       Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  vimhelp movement")
	fmt.Println("  vimhelp editing")
	fmt.Println("  vimhelp dd")
	fmt.Println("  vimhelp 'ci\"'")
}

func printTopics(bindings []Binding) {
	printTitle("topics")

	for _, topic := range getTopics(bindings) {
		fmt.Println("- " + topic)
	}
}

func printTitle(topic string) {
	title := formatTitle(topic)

	fmt.Println("[ " + title + " ]")
}

func formatTitle(value string) string {
	words := strings.Fields(strings.ReplaceAll(value, "-", " "))

	for index, word := range words {
		words[index] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}

func printBinding(binding Binding) {
	fmt.Printf("%-8s %s", binding.Key, binding.Description)

	if binding.Example != "" {
		fmt.Printf(" (%s)", binding.Example)
	}

	fmt.Println()
}

func getTopics(bindings []Binding) []string {
	var topics []string

	for _, binding := range bindings {
		if !contains(topics, binding.Topic) {
			topics = append(topics, binding.Topic)
		}
	}

	return topics
}

func contains(items []string, value string) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}

	return false
}