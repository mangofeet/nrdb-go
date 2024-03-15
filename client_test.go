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
