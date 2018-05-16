package main

import (
	"github.com/yanzay/tbot"
	"log"
	"os"
)

const (
	REGRAS = `Regras:
  - Se comunicar com respeito e consideração aos outros usuários.
	- Compartilhar conteúdo protegido por leis de direitos autorais.
  - Não enviar nudes.
  - Nao enviar porno.
  - Não me invocar de 1 em 1 minuto.
  - Não se passar por outra pessoa (fake).
  - Evitar mensagens que não sejam de interesse comum do grupo.
  - Falar do grupo para os coleguinhas.

	Ao agir contra essas regras, a pessoa será advertida em privado pela administração.
	Em caso de recorrência, será expulsa do grupo.
`
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Println(token)
	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/regras", REGRAS)
	bot.ListenAndServe()

}
