package nrdb

import (
	"fmt"
	"net/url"
	"time"
)

func (cl client) Cards(filter *CardFilter) ([]*Card, error) {
	var res Response[Card]

	var query url.Values
	if filter != nil {
		q, err := filter.Encode()
		if err != nil {
			return nil, fmt.Errorf("encoding filter: %w", err)
		}
		query = q
	}

	if err := cl.nrdbReq("cards", &res, query); err != nil {
		return nil, err
	}

	return res.Data, nil
}

type Card struct {
	Document[CardPoolAttributes, CardPoolRelationships]
}

type CardAttributes struct {
	AdvancementRequirement int               `json:"advancement_requirement"`
	AgendaPoints           int               `json:"agenda_points"`
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
	Search string
}

func (filter CardFilter) Encode() (url.Values, error) {
	query := url.Values{}

	if filter.Search != "" {
		query.Set("filter[search]", filter.Search)
	}

	return query, nil
}

func (doc Card) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}
