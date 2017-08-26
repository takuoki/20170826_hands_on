package fruit

import "log"

type Product struct {
	Name string
}

type Data struct {
	Name     string
	Products []*Product
}

func init() {
	log.Println("import fruit")
}

func GetList() Data {
	p1 := Product{Name: "りんご"}
	p2 := Product{Name: "なし"}
	p3 := Product{Name: "ばなな"}
	p4 := Product{Name: "ぶどう"}
	return Data{Name: "くだもの", Products: []*Product{&p1, &p2, &p3, &p4}}
}
