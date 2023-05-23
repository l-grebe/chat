package chat

import "unicode"

// InArray 判断s是否在数组arr中
func InArray[T comparable](arr []T, s T) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func ContainsNativeLanguage(s string) bool {
	for _, r := range s {
		switch DefaultSetting.NativeLanguage {
		case "zh-CHS":
			if unicode.Is(unicode.Han, r) {
				return true
			}
		case "zh-CHT":
			if unicode.Is(unicode.Han, r) {
				return true
			}
		case "ja":
			if unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r) || (unicode.Is(unicode.Han, r) && unicode.In(r, unicode.Han)) {
				return true
			}
		case "ko":
			if unicode.In(r, unicode.Hangul) {
				return true
			}
			//case "fr":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "es":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "pt":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "it":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "ru":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "vi":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "de":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
			//case "ar":
			//	if unicode.Is(unicode.Han, r) {
			//		return true
			//	}
		}
	}
	return false
}
