package main

import "fmt"

type AcceptedEvent struct {
	ID        string
	Category  string
	Title     string
	ByteCount int
	RuneCount int
}

func main() {
	allowedCategories := [3]string{"build", "test", "deploy"}
	allowedSet := make(map[string]struct{}, len(allowedCategories))
	for _, category := range allowedCategories {
		allowedSet[category] = struct{}{}
	}

	incomingEvents := []struct {
		ID       string
		Category string
		Title    string
	}{
		{ID: "A1", Category: "build", Title: "Compile"},
		{ID: "A2", Category: "test", Title: "世界"},
		{ID: "A1", Category: "deploy", Title: "duplicate"},
		{ID: "A3", Category: "unknown", Title: "skip"},
		{ID: "A4", Category: "deploy", Title: "Ship 🌞"},
	}

	accepted := make([]AcceptedEvent, 0, len(incomingEvents))
	seenIDs := make(map[string]struct{}, len(incomingEvents))
	categoryCounts := make(map[string]int, len(allowedCategories))

	duplicateRejections := 0
	categoryRejections := 0
	totalBytes := 0
	totalRunes := 0

	for _, event := range incomingEvents {
		if _, allowed := allowedSet[event.Category]; !allowed {
			categoryRejections++
			continue
		}

		if _, seen := seenIDs[event.ID]; seen {
			duplicateRejections++
			continue
		}

		byteCount := len([]byte(event.Title))
		runeCount := len([]rune(event.Title))
		accepted = append(accepted, AcceptedEvent{
			ID:        event.ID,
			Category:  event.Category,
			Title:     event.Title,
			ByteCount: byteCount,
			RuneCount: runeCount,
		})

		// Un ID se registra solamente después de que todas las validaciones del evento pasan.
		seenIDs[event.ID] = struct{}{}
		categoryCounts[event.Category]++
		totalBytes += byteCount
		totalRunes += runeCount
	}

	fmt.Println("=== Event Intake Audit ===")
	for index, event := range accepted {
		fmt.Printf("Accepted %d: id=%s category=%s title=%q bytes=%d runes=%d\n", index+1, event.ID, event.Category, event.Title, event.ByteCount, event.RuneCount)
	}

	fmt.Printf("Rejections: duplicate-id=%d invalid-category=%d\n", duplicateRejections, categoryRejections)
	for _, category := range allowedCategories {
		fmt.Printf("Category %s: %d\n", category, categoryCounts[category])
	}
	fmt.Printf("Title totals: bytes=%d runes=%d\n", totalBytes, totalRunes)

	snapshot := make([]AcceptedEvent, len(accepted))
	copy(snapshot, accepted)
	clear(accepted)

	fmt.Printf("Cleared working slice: len=%d values=%+v\n", len(accepted), accepted)
	fmt.Printf("Independent snapshot: len=%d values=%+v\n", len(snapshot), snapshot)
}
