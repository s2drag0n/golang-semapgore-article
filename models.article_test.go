// models.article_test.go

package main

import "testing"

// 测试获取所有文章的函数
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// 检查返回的文章列表长度 = 全局变量articleList
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// 检查每个成员
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
				t.Fail()
				break
		}
	}
}