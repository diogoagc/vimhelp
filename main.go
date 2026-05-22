package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//go:embed bindings.json
var bindingsFile []byte

type Binding struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Topic       string `json:"topic"`
	Example     string `json:"example"`
}

func main() {
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

	if query == "--help" || query == "-h" {
		printHelp()
		return
	}

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
	var bindings []Binding

	err := json.Unmarshal(bindingsFile, &bindings)
	if err != nil {
		return nil, err
	}

	return bindings, nil
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
	fmt.Println("  vimhelp --topics     List available topics")
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