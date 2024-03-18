package nrdb

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func (cl client) Card(cardID string) (*Card, error) {
	var res Response[Card]

	if err := cl.nrdbReq(fmt.Sprintf("cards/%s", cardID), &res, nil); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (cl client) Cards(filter *CardFilter) ([]*Card, error) {

	res, err := cl.cardReq(filter)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (cl client) AllCards(filter *CardFilter) ([]*Card, error) {

	res, err := cl.cardReq(filter)
	if err != nil {
		return nil, err
	}

	// if no links, return now
	if res.Links == nil {
		return res.Data, nil
	}

	// if no "next" link, return now
	if res.Links.Next == nil {
		return res.Data, nil
	}

	nextURLStr := *res.Links.Next
	nextURL, err := url.Parse(nextURLStr)
	if err != nil {
		return nil, fmt.Errorf(`invalid "next" link %s: %w`, nextURLStr, err)
	}

	nextQuery := nextURL.Query()
	nextOffset := nextQuery.Get("page[offset]")
	pageOffset, err := strconv.ParseUint(nextOffset, 10, 64)
	if err != nil {
		return nil, fmt.Errorf(`invalid "next" page offset %s: %w`, nextOffset, err)
	}

	if filter == nil {
		filter = &CardFilter{}
	}

	filter.PageOffset = &pageOffset

	next, err := cl.AllCards(filter)
	if err != nil {
		return nil, fmt.Errorf("getting offset %d: %w", pageOffset, err)
	}

	return append(res.Data, next...), nil
}

type Card struct {
	Document[CardAttributes, CardRelationships]
}

type CardAttributes struct {
	AdvancementRequirement *int              `json:"advancement_requirement"`
	AgendaPoints           *int              `json:"agenda_points"`
	Attribution            *string           `json:"attribution"`
	BaseLink               *int              `json:"base_link"`
	CardAbilities          *CardAbilities    `json:"card_abilities"`
	CardCycleIDs           []string          `json:"card_cycle_ids"`
	CardPoolIDs            []string          `json:"card_pool_ids"`
	CardSetIDs             []string          `json:"card_set_ids"`
	CardSubtypeIDs         []string          `json:"card_subtype_ids"`
	CardTypeID             string            `json:"card_type_id"`
	Cost                   *int              `json:"cont"`
	DateRelease            string            `json:"date_release"`
	DeckLimit              int               `json:"deck_limit"`
	DesignedBy             string            `json:"designed_by"`
	DisplaySubtypes        *string           `json:"display_subtypes"`
	FactionID              string            `json:"faction_id"`
	FormatIDs              []string          `json:"format_ids"`
	InRestriction          bool              `json:"in_restriction"`
	InfluenceCost          *int              `json:"influence_cost"`
	InfluenceLimit         *int              `json:"influence_limit"`
	IsUnique               bool              `json:"is_unique"`
	LatestPrintingID       string            `json:"latest_printing_id"`
	MemoryCost             *int              `json:"memory_cost"`
	MinimumDeckSize        *int              `json:"minimum_deck_size"`
	NumPrintings           int               `json:"num_printings"`
	PrintingIDs            []string          `json:"printing_ids"`
	PrintingsReleasedBy    []string          `json:"printings_released_by"`
	Pronouns               *string           `json:"pronouns"`
	RestrictionIDs         []string          `json:"restriction_ids"`
	Restrictions           *CardRestrictions `json:"restrictions"`
	SideID                 string            `json:"side_id"`
	SnapshotIDs            []string          `json:"snapshot_ids"`
	Strength               *int              `json:"strength"`
	StrippedText           string            `json:"stripped_text"`
	StrippedTitle          string            `json:"stripped_title"`
	Title                  string            `json:"title"`
	TrashCost              *int              `json:"trash_cost"`
	UpdatedAt              time.Time         `json:"updated_at"`
}

type CardRestrictions struct {
	Banned               []string       `json:"banned"`
	GlobalPenalty        []string       `json:"global_penalty"`
	Points               map[string]int `json:"points"`
	Restricted           []string       `json:"restricted"`
	UniversalFactionCost map[string]int `json:"universal_faction_cost"`
}

type CardAbilities struct {
	AddtionalCost            bool `json:"additional_cost"`
	Advanceable              bool `json:"advanceable"`
	GainsSubroutines         bool `json:"gains_subroutines"`
	Interrupt                bool `json:"interrupt"`
	LinkProvided             *int `json:"link_provided"`
	MUProvided               *int `json:"mu_provided"`
	NumPrintedSubroutines    *int `json:"num_printed_subroutines"`
	OnEncounterEffect        bool `json:"on_encounter_effect"`
	PerformsTrace            bool `json:"performs_trace"`
	RecurringCreditsProvided *int `json:"recurring_credits_provided"`
	RezEffect                bool `json:"rez_effect"`
	TrashAbility             bool `json:"trash_ability"`
}

type CardRelationships struct {
	CardSubtypes *Relationship `json:"card_subtypes"`
	CardType     *Relationship `json:"card_type"`
	Faction      *Relationship `json:"faction"`
	Printings    *Relationship `json:"printings"`
	Rulings      *Relationship `json:"rulings"`
	Side         *Relationship `json:"side"`
}

type CardFilter struct {
	Params
	Search *string
}

func (filter CardFilter) Query() (url.Values, error) {
	query := filter.SetPageInfo(url.Values{})

	if filter.Search != nil {
		query.Set("filter[search]", *filter.Search)
	}

	return query, nil
}

func (doc Card) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Title, doc.ID)
}

func (cl client) cardReq(filter *CardFilter) (*Response[[]*Card], error) {
	var res Response[[]*Card]

	var query url.Values
	if filter != nil {
		q, err := filter.Query()
		if err != nil {
			return nil, fmt.Errorf("encoding filter: %w", err)
		}
		query = q
	}

	if err := cl.nrdbReq("cards", &res, query); err != nil {
		return nil, err
	}

	return &res, nil
}
