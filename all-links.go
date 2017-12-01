package htmlPage

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

import (
	hash "github.com/rogeecn/gohash"
)

type FetchRule struct {
	Include   string
	Exclude   string
	StartWith string
	EndWith   string
}

//获取指定结构内的所有URL链接
//这里的URL并不会进行URL自动补全，还需要自己补全
func AllLinks(pageHtml, attr string, rule FetchRule) ([]string, error) {
	if len(pageHtml) == 0 {
		return nil, fmt.Errorf("selector is empty")
	}
	body, err := goquery.NewDocumentFromReader(strings.NewReader(pageHtml))
	if err != nil {
		return nil, err
	}

	var result []string
	body.Find("a").Each(func(idx int, item *goquery.Selection) {
		if itemResult, exist := item.Attr(attr); exist {
			if urlMatchRule(itemResult, rule) {
				result = append(result, itemResult)
			}
		}
	})

	var retResult []string
	tmpMapStorage := make(map[string]string)
	for _, item := range result {
		itemHash := hash.Hash(hash.MD5, item)
		if _, ok := tmpMapStorage[itemHash]; !ok {
			tmpMapStorage[itemHash] = "1"
			retResult = append(retResult, item)
		}
	}
	tmpMapStorage = nil

	return retResult, nil
}

func urlMatchRule(url string, rule FetchRule) bool {
	//包含
	if len(rule.Include) > 0 {
		if !strings.Contains(url, rule.Include) {
			return false
		}
	}

	if len(rule.Exclude) > 0 {
		if strings.Contains(url, rule.Exclude) {
			return false
		}
	}

	if len(rule.StartWith) > 0 {
		if strings.Index(url, rule.StartWith) != 0 {
			return false
		}
	}

	if len(rule.EndWith) > 0 {
		if strings.LastIndex(url, rule.EndWith) != (len(url) - len(rule.EndWith)) {
			return false
		}
	}

	return true
}
