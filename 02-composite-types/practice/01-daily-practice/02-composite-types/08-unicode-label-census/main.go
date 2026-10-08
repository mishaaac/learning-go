package main

import "fmt"

type LabelRecord struct {
	Text      string
	ByteCount int
	RuneCount int
}

func main() {
	labels := []string{"Go", "Gopher", "世界", "Go", "🌞"}
	records := make([]LabelRecord, 0, len(labels))
	uniqueLabels := make(map[string]struct{}, len(labels))

	totalBytes := 0
	totalRunes := 0

	for _, label := range labels {
		byteCount := len([]byte(label))
		runeCount := len([]rune(label))

		records = append(records, LabelRecord{
			Text:      label,
			ByteCount: byteCount,
			RuneCount: runeCount,
		})

		if _, exists := uniqueLabels[label]; !exists {
			uniqueLabels[label] = struct{}{}
		}

		totalBytes += byteCount
		totalRunes += runeCount
	}

	fmt.Println("=== Unicode Label Census ===")
	for index, record := range records {
		fmt.Printf("Record %d: text=%q bytes=%d runes=%d\n", index+1, record.Text, record.ByteCount, record.RuneCount)
	}

	fmt.Printf("Totals: records=%d unique=%d bytes=%d runes=%d\n", len(records), len(uniqueLabels), totalBytes, totalRunes)
}
