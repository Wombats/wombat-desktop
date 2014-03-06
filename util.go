package main

import ()


func IsMember (s string, arr []string) (bool) {
	for i := 0; i < len(arr); i++ {
		if s == arr[i] {
			return true
		}
	}
	return false
}
