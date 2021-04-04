package main

// import (
// 	"errors"
// )

// // type nome []string
// type Pessoa struct {
// 	Nome  string
// 	Idade int
// }

// func (p *Pessoa) setNome(nome string) {
// 	p.Nome = nome
// 	println(p.Nome)
// }

// func (p Pessoa) andou() (string, error) {
// 	if p.Nome != "Leonardo" {
// 		return "", errors.New("Nome deve ser o Leonardo")
// 	}
// 	return p.Nome + " andou!", nil
// }

// func main() {
// var nome nome
// nome = append(nome, "Leonardo")
// nome = append(nome, "Bruna")
// nome = append(nome, "Davi")

// for _, j := range nome {
// 	fmt.Println(j)
// }
// Leonardo := Pessoa{
// 	Nome:  "Leonardo",
// 	Idade: 36,
// }
// // fmt.Println(Leonardo.Nome)
// res, err := Leonardo.andou()

// if err != nil {
// 	fmt.Println(err.Error())
// }
// fmt.Println(res)

//- Bi-References
// nome := "Leonardo"
// var nome2 *string
// nome2 = &nome //-- Endereço de memoria
// fmt.Println(nome2)
// fmt.Println(*nome2)
// pessoa := Pessoa{
// 	Nome:  "Leonardo",
// 	Idade: 20,
// }
// pessoa.setNome("Davi")
// fmt.Println(pessoa.Nome)

// }
func main() {
	product1 := model.NewProduct()
	product1.Name = "Carrinho"

	product2 := model.NewProduct()
	product2.Name = "Boneca"
}
