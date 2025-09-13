package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"sodnix/apps/server/src/common/converter"
	"sodnix/apps/server/src/modules/accounts"
	"sodnix/apps/server/src/modules/categories"
	"sodnix/apps/server/src/modules/transactions"
	"sodnix/apps/server/src/modules/types"
	"sodnix/apps/server/src/modules/users"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(converter.ConcatToAny(categories.Models(), types.Models(), transactions.Models(), accounts.Models(), users.Models())...)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
