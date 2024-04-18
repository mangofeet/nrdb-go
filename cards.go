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

func (doc Card) AdvancementRequirement() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.AdvancementRequirement == nil {
		return 0
	}
	return *doc.Attributes.AdvancementRequirement
}

func (doc Card) AgendaPoints() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.AgendaPoints == nil {
		return 0
	}
	return *doc.Attributes.AgendaPoints
}

func (doc Card) Attribution() string {
	if doc.Attributes == nil {
		return ""
	}
	if doc.Attributes.Attribution == nil {
		return ""
	}
	return *doc.Attributes.Attribution
}

func (doc Card) BaseLink() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.BaseLink == nil {
		return 0
	}
	return *doc.Attributes.BaseLink
}

func (doc Card) CardAbilities() *CardAbilities {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.CardAbilities
}

func (doc Card) CardCycleIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.CardCycleIDs
}

func (doc Card) CardPoolIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.CardPoolIDs
}

func (doc Card) CardSetIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.CardSetIDs
}

func (doc Card) CardSubtypeIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.CardSubtypeIDs
}

func (doc Card) CardTypeID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.CardTypeID
}

func (doc Card) Cost() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.Cost == nil {
		return 0
	}
	return *doc.Attributes.Cost
}

func (doc Card) DateRelease() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.DateRelease
}

func (doc Card) DeckLimit() int {
	if doc.Attributes == nil {
		return 0
	}
	return doc.Attributes.DeckLimit
}

func (doc Card) DesignedBy() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.DesignedBy
}

func (doc Card) DisplaySubtypes() string {
	if doc.Attributes == nil {
		return ""
	}
	if doc.Attributes.DisplaySubtypes == nil {
		return ""
	}
	return *doc.Attributes.DisplaySubtypes
}

func (doc Card) FactionID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.FactionID
}

func (doc Card) FormatIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.FormatIDs
}

func (doc Card) InRestriction() bool {
	if doc.Attributes == nil {
		return false
	}
	return doc.Attributes.InRestriction
}

func (doc Card) InfluenceCost() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.InfluenceCost == nil {
		return 0
	}
	return *doc.Attributes.InfluenceCost
}

func (doc Card) InfluenceLimit() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.InfluenceLimit == nil {
		return 0
	}
	return *doc.Attributes.InfluenceLimit
}

func (doc Card) IsUnique() bool {
	if doc.Attributes == nil {
		return false
	}
	return doc.Attributes.IsUnique
}

func (doc Card) LatestPrintingID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.LatestPrintingID
}

func (doc Card) MemoryCost() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.MemoryCost == nil {
		return 0
	}
	return *doc.Attributes.MemoryCost
}

func (doc Card) MinimumDeckSize() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.MinimumDeckSize == nil {
		return 0
	}
	return *doc.Attributes.MinimumDeckSize
}

func (doc Card) NumPrintings() int {
	if doc.Attributes == nil {
		return 0
	}
	return doc.Attributes.NumPrintings
}

func (doc Card) PrintingIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.PrintingIDs
}

func (doc Card) PrintingsReleasedBy() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.PrintingsReleasedBy
}

func (doc Card) Pronouns() string {
	if doc.Attributes == nil {
		return ""
	}
	if doc.Attributes.Pronouns == nil {
		return ""
	}
	return *doc.Attributes.Pronouns
}

func (doc Card) RestrictionIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.RestrictionIDs
}

func (doc Card) Restrictions() *CardRestrictions {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.Restrictions
}

func (doc Card) SideID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.SideID
}

func (doc Card) SnapshotIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.SnapshotIDs
}

func (doc Card) Strength() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.Strength == nil {
		return 0
	}
	return *doc.Attributes.Strength
}

func (doc Card) StrippedText() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.StrippedText
}

func (doc Card) StrippedTitle() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.StrippedTitle
}

func (doc Card) Title() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.Title
}

func (doc Card) TrashCost() int {
	if doc.Attributes == nil {
		return 0
	}
	if doc.Attributes.TrashCost == nil {
		return 0
	}
	return *doc.Attributes.TrashCost
}

func (doc Card) UpdatedAt() time.Time {
	if doc.Attributes == nil {
		return time.Time{}
	}
	return doc.Attributes.UpdatedAt
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
	Cost                   *int              `json:"cost"`
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
