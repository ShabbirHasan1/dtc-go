package codegen

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParseMessageDocs(t *testing.T) {
	docs, err := ParseMessageDocs(
		"testdata/docs/account.txt",
		"testdata/docs/account_balance.txt",
		"testdata/docs/auth.txt",
		"testdata/docs/historical_price.txt",
		"testdata/docs/logging.txt",
		"testdata/docs/market.txt",
		"testdata/docs/order.txt",
		"testdata/docs/symbol.txt",
		"testdata/docs/trading.txt",
	)
	if err != nil {
		t.Fatal(err)
	}

	types, err := ParseTypeDocs("testdata/docs/field_types.txt")
	if err != nil {
		t.Fatal(err)
	}
	docs.Merge(types)
	docs.NormalizeNames()
	fmt.Println(docs)

	jsonData, err := json.Marshal(docs)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("testdata/docs.json", jsonData, 0755)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
