package main

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	return shortestPath(0, 0, c)

}

func shortestPath(index int, clouds int32, c []int32) int32 {

	if len(c)-1-index == 0 {
		return clouds
	}

	if c[index+1] == 0 {
		if index+2 == len(c) || c[index+2] == 1 {
			return shortestPath(index+1, clouds+1, c)
		}

		sPath1 := shortestPath(index+1, clouds+1, c)
		sPath2 := shortestPath(index+2, clouds+1, c)

		if sPath1 < sPath2 {
			return sPath1
		}
		return sPath2
	}

	return shortestPath(index+2, clouds+1, c)

}
