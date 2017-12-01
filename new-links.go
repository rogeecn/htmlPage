package htmlPage

import (
	hash "github.com/rogeecn/gohash"
)

const (
	ORDER_HEAD = iota
	ORDER_TAIL
)

// 链接列表中差异化的最新链接地址
func NewLinks(allLinks []string, compareId string, hashType, order int) []string {
	keyIndex := -1

	for index, link := range allLinks {
		linkHash := hash.Hash(hashType, link)
		if linkHash == compareId {
			keyIndex = index
			break
		}
	}

	// 如果是-1返回所有LINKS
	if keyIndex == -1 {
		return allLinks
	}

	if order == ORDER_HEAD {
		if keyIndex == 0 {
			return nil
		}
		return allLinks[0:keyIndex]
	}

	if order == ORDER_TAIL {
		if keyIndex == len(allLinks)-1 {
			return nil
		}
		return allLinks[keyIndex+1:len(allLinks)]
	}

	return nil
}
