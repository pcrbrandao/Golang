package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"Golang/ReceitasGo/mensagem"
	"unicode"
	"Golang/ReceitasGo/lib"
)

type Usuario struct {

	// Contém ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model

	Nome string `gorm:"size:255"`
	Email Email
	Senha []byte
}

/*
 * Password rules:
 * at least 7 letters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character
 */
func (u *Usuario)SetSenha(senha string) error {

	if len(u.verificaSenha(senha)) > 0 {
		return fmt.Errorf("%s %s", mensagem.NAOEVALIDO)
	}

	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("%s %s %s", mensagem.ERRO, senha, mensagem.NAOEVALIDO)
	}

	u.Senha = hashedSenha
	return nil
}

// Passa por todos os caracteres da string e remove de condicoes
// o tipo correspondente, caso haja.
// Se condicoes retornar vazio, é porque todas as condicoes foram cumpridas
// As condicoes que retornaram são as que não foram satisfeitas.
func (u *Usuario) verificaSenha(s string) (condicoes []lib.Char) {

	condicoes = []lib.Char{
		lib.NUMERO, lib.MAIUSCULA,
		lib.ESPECIAL, lib.MENORQUESETE}

	for _, s := range s {
		switch {
		case unicode.IsNumber(s):
			fmt.Printf("remove número...%q\n", s)
			condicoes = lib.CharSliceRemoveFrom(condicoes, lib.NUMERO)
		case unicode.IsUpper(s):
			fmt.Printf("remove maiúscula... %q\n", s)
			condicoes = lib.CharSliceRemoveFrom(condicoes, lib.MAIUSCULA)
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			fmt.Printf("remove especial... %q\n", s)
			condicoes = lib.CharSliceRemoveFrom(condicoes, lib.ESPECIAL)
		case unicode.IsLetter(s) || s == ' ':
			fmt.Printf("letras minúscula... %q\n", s)
		default:
			fmt.Printf("Não sei que caracter é esse... %q", s)
			return condicoes
		}
	}
	if l := len(s); l == 7 {
		fmt.Printf("comprimento correto... %d\n", len(s))
		condicoes = lib.CharSliceRemoveFrom(condicoes, lib.MENORQUESETE)
	}
	return condicoes
}

func (a *Usuario)GetID() uint {
	return a.ID
}