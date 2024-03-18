package nrdb_test

import (
	"testing"

	"github.com/mangofeet/nrdb-go"
)

func TestCardCycles(t *testing.T) {

	cl := nrdb.NewClient()

	res, err := cl.CardCycles()
	if err != nil {
		t.Fatal(err)
	}

	for _, doc := range res {
		t.Log(doc)
	}

}

func TestCardPools(t *testing.T) {

	cl := nrdb.NewClient()

	res, err := cl.CardPools()
	if err != nil {
		t.Fatal(err)
	}

	for _, doc := range res {
		t.Log(doc)
	}

}

func TestCardSetTypes(t *testing.T) {

	cl := nrdb.NewClient()

	res, err := cl.CardSetTypes()
	if err != nil {
		t.Fatal(err)
	}

	for _, doc := range res {
		t.Log(doc)
	}

}

func TestCardSets(t *testing.T) {

	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.CardSets(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	t.Run("filter card set type", func(t *testing.T) {
		res, err := cl.CardSets(&nrdb.CardSetFilter{
			CardSetTypeID: "core",
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})

	t.Run("filter card cycle", func(t *testing.T) {
		res, err := cl.CardSets(&nrdb.CardSetFilter{
			CardCycleID: "liberation",
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})

	t.Run("filter cycle and type", func(t *testing.T) {
		res, err := cl.CardSets(&nrdb.CardSetFilter{
			CardCycleID:   "ashes",
			CardSetTypeID: "booster_pack",
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})

}

func TestCardSubtypes(t *testing.T) {

	cl := nrdb.NewClient()

	res, err := cl.CardSubtypes()
	if err != nil {
		t.Fatal(err)
	}

	for _, doc := range res {
		t.Log(doc)
	}

}

func TestCardTypes(t *testing.T) {

	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.CardTypes(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	t.Run("filter side", func(t *testing.T) {
		res, err := cl.CardTypes(&nrdb.CardTypeFilter{
			SideID: "runner",
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})

}

func TestCards(t *testing.T) {

	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.Cards(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	search := "e:system_gateway"

	t.Run("filter by set", func(t *testing.T) {
		res, err := cl.Cards(&nrdb.CardFilter{
			Search: &search,
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})

}

func TestAllCards(t *testing.T) {
	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.AllCards(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	search := "card_cycle:borealis"

	t.Run("filter by cycle", func(t *testing.T) {
		res, err := cl.AllCards(&nrdb.CardFilter{
			Search: &search,
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})
}

func TestCard(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Card("sure_gamble")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

}

func TestFactions(t *testing.T) {
	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.Factions(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	side := "runner"

	t.Run("filter by side", func(t *testing.T) {

		res, err := cl.Factions(&nrdb.FactionFilter{
			SideID: &side,
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})
}

func TestFaction(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Faction("criminal")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

}

func TestFormats(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Formats()
	if err != nil {
		t.Fatal(err)
	}

	for _, doc := range res {
		t.Log(doc)
	}
}

func TestFormat(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Format("standard")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

}

func TestIllustrators(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Illustrators()
	if err != nil {
		t.Fatal(err)
	}

	for _, doc := range res {
		t.Log(doc)
	}
}

func TestIllustrator(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Illustrator("zoe_cohen")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

}

func TestPrintings(t *testing.T) {

	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.Printings(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	search := "e:system_gateway"

	t.Run("filter by set", func(t *testing.T) {
		res, err := cl.Printings(&nrdb.PrintingFilter{
			CardFilter: nrdb.CardFilter{Search: &search},
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})

}

func TestAllPrintings(t *testing.T) {
	cl := nrdb.NewClient()

	t.Run("no filter", func(t *testing.T) {
		res, err := cl.AllPrintings(nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}

	})

	search := "card_cycle:liberation"

	t.Run("filter by cycle", func(t *testing.T) {
		res, err := cl.AllPrintings(&nrdb.PrintingFilter{
			CardFilter: nrdb.CardFilter{Search: &search},
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, doc := range res {
			t.Log(doc)
		}
	})
}

func TestPrinting(t *testing.T) {
	cl := nrdb.NewClient()

	res, err := cl.Printing("30030")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

}
