package lc44

const NORMAL = "NORMAL"
const STAR = "STAR"

func isMatch(s string, p string) bool {
	state := NORMAL

	sIdx, pIdx, starIdx, latestMatch := 0, 0, -1, 0
	for sIdx < len(s) {
		if pIdx < len(p) && (s[sIdx] == p[pIdx] || p[pIdx] == '?') {
			state = NORMAL
			sIdx++
			pIdx++
		} else if pIdx < len(p) && p[pIdx] == '*' {
			state = STAR
			starIdx = pIdx
			latestMatch = sIdx
			pIdx++
		} else if state == STAR {
			latestMatch++
			sIdx++
		} else if starIdx != -1 && pIdx > (starIdx+1) {
			state = STAR
			pIdx = starIdx + 1
			latestMatch++
			sIdx = latestMatch
		} else {
			return false
		}
	}
	if pIdx < len(p) {
		for ; pIdx < len(p); pIdx++ {
			if p[pIdx] != '*' {
				return false
			}
		}
	} else if pIdx >= len(p) && sIdx < len(s) && state != STAR {
		return false
	}
	return true
}
