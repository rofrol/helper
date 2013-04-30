package helper

import (
	"code.google.com/p/go.net/html"
	//"reflect"
	"sort"
	"strings"
)

func FindByClass(n *html.Node, nType string, w string) *html.Node {
	//w n.Data może być html, table albo wieloliniowy string z białymi znakami
	if n.Type == html.ElementNode && n.Data == nType {
		//fmt.Print("type ", reflect.TypeOf(n))
		for _, a := range n.Attr {
			if a.Key == "class" {
				if t := FoundClass(a.Val, w); t {
					return n
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := FindByClass(c, nType, w); found != nil {
			return found
		}
	}
	return nil
}

// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/j4vFdmMZa_4
// Musi być posortowane najpierw, a potem jeszcze porównać czy wartość pod zwróconym indeksem jest ta szukana
// czy może został zwrócony indeks gdzie ta wartość byłaby wstawiona
func Found(a []string, w string) bool {
	return a[sort.SearchStrings(a, w)] == w
}
func FoundClass(a string, w string) bool {
	//TODO: what if there is more than one space between class names
	s := strings.Split(a, " ")
	sort.Strings(s)
	return Found(s, w)
}
func FirstChildByTag(nn *html.Node, data string) *html.Node {
	var searched *html.Node
	for searched = nn.FirstChild; searched != nil; searched = searched.NextSibling {
		if searched.Data == data {
			break
		}
	}
	return searched
}
func NextSiblingByTag(nn *html.Node, data string) *html.Node {
	var searched *html.Node
	for searched = nn.NextSibling; searched != nil; searched = searched.NextSibling {
		if searched.Data == data {
			break
		}
	}
	return searched
}

func ElementsByTag(nn *html.Node, data string) []*html.Node {
	var arr []*html.Node

	for searched := nn.FirstChild; searched != nil; searched = searched.NextSibling {
		if searched.Data == data {
			arr = append(arr, searched)
		}
	}
	return arr
}

func ElementsByTagRec(n *html.Node, nType string) []*html.Node {
	var arr []*html.Node
	//w n.Data może być html, table albo wieloliniowy string z białymi znakami
	if n.Type == html.ElementNode && n.Data == nType {
		//fmt.Print("type ", reflect.TypeOf(n))
		arr = append(arr, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := ElementsByTagRec(c, nType); found != nil {
			arr = concat(arr, found)
		}
	}
	return arr
}

func concat(old1 []*html.Node, old2 []*html.Node) []*html.Node {
	newslice := make([]*html.Node, len(old1)+len(old2))
	copy(newslice, old1)
	copy(newslice[len(old1):], old2)
	return newslice
}

func String2CsvCell(title string) string {
	title = strings.TrimSpace(title)
	title = strings.Replace(title, "Tytuł: ", "", 1)
	return title
}
