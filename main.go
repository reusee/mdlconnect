package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	pt = fmt.Printf
)

func main() {
	metaFilePath := os.Args[1]
	contentBs, err := ioutil.ReadFile(metaFilePath)
	ce(err, "read meta file")
	content := string(contentBs)

	type Product struct {
		Sku      int64
		ThumbUrl string
	}

	products := []Product{}

	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		spu, err := strconv.ParseInt(parts[0], 10, 64)
		ce(err, "parse spu")
		thumbUrl := strings.TrimSpace(parts[1])
		if !strings.HasPrefix(thumbUrl, "http://") {
			panic(me(nil, "invalid thumbnail url: %s", parts[1]))
		}
		products = append(products, Product{
			Sku:      spu,
			ThumbUrl: thumbUrl,
		})
	}

	headerUrl := "http://img10.360buyimg.com/imgzone/jfs/t1873/285/2917542460/6442/7006162e/56f75e4cNe88b458b.jpg"

	for i, product := range products {
		pt("=> %d %d\n", i+1, product.Sku)
		pt(`<img src="%s" class="connect-head" /><br />`, headerUrl)
		indexes := Ints([]int{})
		for n := 1; n <= 12; n++ {
			indexes = append(indexes, (i+n)%len(products))
		}
		indexes.Shuffle()
		for num, index := range indexes {
			product = products[index]
			pt(`<a target="_blank" href="http://item.jd.com/%d.html"><img src="%s" /></a>`,
				product.Sku,
				product.ThumbUrl,
			)
			if (num+1)%4 == 0 {
				pt("<br />")
			}
		}
		pt(`<br /><br class="connect-tail" />`)
		pt("\n\n")
	}
}
