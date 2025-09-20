package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	/*
    (?m): This is a multi-line flag. While not strictly 		necessary for this specific problem, it's good practice 	for regex patterns that might be used on multi-line 		strings.

	<: Matches the literal character <.

    [~*=-]+: This is a character set that matches one or 		more (+) of the characters within the brackets.

	>: Matches the literal character >.
	*/
	re := regexp.MustCompile(`(?m)<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
    re := regexp.MustCompile(`(?i)".*password.*"`)
    for _,password := range lines {
        if re.MatchString(password) {
            count ++
        }
    } 

    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

	return re.ReplaceAllString(text, "")
}

// TagWithUserName processes log lines to tag those containing a username.
func TagWithUserName(lines []string) []string {
	// re compiles a regular expression to find a username that follows the word "User".
// - `User`  : Matches the literal characters "User".
// - `\s+`   : Matches one or more whitespace characters (like spaces or tabs).
// - `(\S+)` : This is a capturing group that matches one or more non-whitespace characters (the actual username).
	re := regexp.MustCompile(`User\s+(\S+)`)

	processedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		// FindStringSubmatch looks for a match in the line.
		// If found, it returns a slice where index 0 is the full match
		// and index 1 is the content of the first capture group `(\S+)`.
		matches := re.FindStringSubmatch(line)

		if len(matches) > 1 {
			// A match was found, and the username was captured.
			// `matches[1]` contains the username.
			username := matches[1]
			taggedLine := fmt.Sprintf("[USR] %s %s", username, line)
			processedLines = append(processedLines, taggedLine)
		} else {
			// No match was found, so add the original line unchanged.
			processedLines = append(processedLines, line)
		}
	}

	return processedLines
}
