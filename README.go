# AC3
package main

import "fmt"


type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionarContato(contato Contato, listaContatos *[5]Contato) {
	for i, c := range *listaContatos {
		if c.Nome == "" && c.Email == "" {
			(*listaContatos)[i] = contato
			return
		}
	}
}

func excluirContato(listaContatos *[5]Contato) {
	for i := 4; i >= 0; i-- {
		if (*listaContatos)[i].Nome != "" || (*listaContatos)[i].Email != "" {
			(*listaContatos)[i] = Contato{} // Limpar o contato
			return
		}
	}
}

func main() {
    var contatos [5]Contato

    for {
        var opcao int
        fmt.Println("Escolha uma opção:")
        fmt.Println("1- Adicionar contato")
        fmt.Println("2- Excluir contato")
        fmt.Println("3- Sair")

        fmt.Print("Opção: ")
        fmt.Scanln(&opcao)

        switch opcao {
        case 1:
            var nome, email string
            fmt.Print("Nome do contato: ")
            fmt.Scanln(&nome)
            fmt.Print("Email do contato: ")
            fmt.Scanln(&email)

            novoContato := Contato{Nome: nome, Email: email}
            adicionarContato(novoContato, &contatos)

        case 2:
            excluirContato(&contatos)

        case 3:
            fmt.Println("Saindo")
            return

        default:
            fmt.Println("Opção inválida. Escolha uma opção válida.")
        }
    }
}
