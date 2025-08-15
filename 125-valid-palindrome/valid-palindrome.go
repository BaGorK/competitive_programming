func isPalindrome(s string) bool {
    fr := 0
    bc := len(s) - 1;

    for fr < bc {
        for fr < bc && !isAlphaNum(rune(s[fr])) {
            fr++
        }

        for fr < bc && !isAlphaNum(rune(s[bc])) {
            bc--
        }

        if unicode.ToLower(rune(s[fr])) != unicode.ToLower(rune(s[bc])) {
            return false
        }
        
        fr++
        bc--
    }
    return true
}

func isAlphaNum(c rune) bool {
    return (c >= 'a' && c <= 'z') ||
         (c >= 'A' && c <= 'Z') ||
         (c >= '0' && c <= '9')
}