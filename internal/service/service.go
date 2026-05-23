package service

import (
	"fmt"
	"strings"
)

var textToMorse = map[string]string{
	"А":  ".-",
	"Б":  "-...",
	"В":  ".--",
	"Г":  "--.",
	"Д":  "-..",
	"Е":  ".",
	"Ж":  "...-",
	"З":  "--..",
	"И":  "..",
	"Й":  ".---",
	"К":  "-.-",
	"Л":  ".-..",
	"М":  "--",
	"Н":  "-.",
	"О":  "---",
	"П":  ".--.",
	"Р":  ".-.",
	"С":  "...",
	"Т":  "-",
	"У":  "..-",
	"Ф":  "..-.",
	"Х":  "....",
	"Ц":  "-.-.",
	"Ч":  "---.",
	"Ш":  "----",
	"Щ":  "--.-",
	"ЪЬ": "-..-",
	"Ы":  "-.--",
	"Э":  "..-..",
	"Ю":  "..--",
	"Я":  ".-.-",
}
var morseToText = map[string]string{
	".-":    "А",
	"...-":  "Б",
	"--.":   "В",
	".--":   "Г",
	"..-":   "Д",
	".":     "Е",
	"-...":  "Ж",
	"..--":  "З",
	"..":    "И",
	"---.":  "Й",
	"-.-":   "К",
	"..-.":  "Л",
	"--":    "М",
	"-.":    "Н",
	"---":   "О",
	".--.":  "П",
	".-.":   "Р",
	"...":   "С",
	"-":     "Т",
	"-..":   "У",
	".-..":  "Ф",
	"....":  "Х",
	".-.-":  "Ц",
	".---":  "Ч",
	"----":  "Ш",
	"-.--":  "Щ",
	"-..-":  "ЪЬ",
	"--.-":  "Ы",
	"..-..": "Э",
	"--..":  "Ю",
	"-.-.":  "Я",
}

func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return false
		}
	}
	return true
}
func Convert(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("empty input")
	}
	if isMorse(input) {
		return Decode(input)
	}
	return Encode(input)
}

// Text to Morse
func Encode(text string) (string, error) {

	var result []string

	text = strings.ToUpper(text)

	for _, ch := range text {
		v := string(ch)
		code, ok := textToMorse[v]
		if !ok {
			return "", fmt.Errorf("unknown symbol: %s", v)
		}
		result = append(result, code)
	}
	return strings.Join(result, " "), nil
}

// Morse to Text
func Decode(code string) (string, error) {
	devLine := strings.Split(code, " ")
	var result []string

	for _, ch := range devLine {
		v, ok := morseToText[ch]
		if !ok {
			return "", fmt.Errorf("unknown morse code: %s", ch)
		}
		result = append(result, v)
	}
	return strings.Join(result, ""), nil
}
