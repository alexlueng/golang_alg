package facebook

// Feature 6 Group the similar gibberish posts together so a decoding pattern can be observed.
func generateKey(message string) string {
	key := ""
	for i := 1; i < len(message); i++ {
		diff := int(message[i]) - int(message[i-1])
		if diff < 0 {
			key += string(diff + 26)
		} else {
			key += string(diff)
		}
		key += ","
	}
	return key
}

func combineMessage(messages []string) [][]string {
	messageGroup := make(map[string][]string)
	for _, message := range messages {
		key := generateKey(message)
		var list []string
		if value, ok := messageGroup[key]; ok {
			list = value
		}
		list = append(list, message)
		messageGroup[key] = list
	}
	var groups [][]string
	for _, group := range messageGroup {
		groups = append(groups, group)
	}
	return groups
}
