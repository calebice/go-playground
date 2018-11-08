package main

//n number of steps, s string determining the hill

// Steps to solve:
// Start at the beginning, keep track of "height"
// track every time go below sea level back up to sea level
// Iterate across string of steps calculating change in height
// UUDDDU +2 -2 = 0 -1 = -1 +1 = 0
func countingValleys(n int32, s string) int32 {
	var height int
	var valleys int32
	for i := 0; i < int(n); i++ {
		prevHeight := height
		height = changeHeight(height, string(s[i]))
		if prevHeight < 0 && height == 0 {
			valleys++
		}
	}

	return valleys
}

func changeHeight(height int, s string) int {
	if s == "U" {
		return height + 1
	}
	return height - 1
}
