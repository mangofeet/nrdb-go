package nrdb

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func (cl client) Printing(printingID string) (*Printing, error) {
	var res Response[Printing]

	if err := cl.nrdbReq(fmt.Sprintf("printings/%s", printingID), &res, nil); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (cl client) Printings(filter *PrintingFilter) ([]*Printing, error) {

	res, err := cl.printingReq(filter)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (cl client) AllPrintings(filter *PrintingFilter) ([]*Printing, error) {

	res, err := cl.printingReq(filter)
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
		filter = &PrintingFilter{}
	}

	filter.PageOffset = &pageOffset

	next, err := cl.AllPrintings(filter)
	if err != nil {
		return nil, fmt.Errorf("getting offset %d: %w", pageOffset, err)
	}

	return append(res.Data, next...), nil
}

type Printing struct {
	Document[PrintingAttributes, PrintingRelationships]
}

type PrintingAttributes struct {
	CardID                 string            `json:"card_id"`
	CardCycleID            string            `json:"card_cycle_id"`
	CardCycleName          string            `json:"card_cycle_name"`
	CardSetID              string            `json:"card_set_id"`
	CardSetName            string            `json:"card_set_name"`
	Flavor                 string            `json:"flavor"`
	DisplayIllustrators    string            `json:"string"`
	IllustratorIDs         []string          `json:"illustrator_ids"`
	IllustratorNames       []string          `json:"illustrator_names"`
	Position               int               `json:"position"`
	PositionInSet          int               `json:"position_in_set"`
	Quantity               int               `json:"quantity"`
	DateRelease            string            `json:"date_release"`
	UpdatedAt              time.Time         `json:"updated_at"`
	AdvancementRequirement *int              `json:"advancement_requirement"`
	AgendaPoints           *int              `json:"agenda_points"`
	BaseLink               *int              `json:"base_link"`
	CardTypeID             string            `json:"card_type_id"`
	Cost                   *int              `json:"cont"`
	DeckLimit              int               `json:"deck_limit"`
	DisplaySubtypes        *string           `json:"display_subtypes"`
	CardSubtypeIDs         []string          `json:"card_subtype_ids"`
	CardSubtypeNames       []string          `json:"card_subtype_names"`
	FactionID              string            `json:"faction_id"`
	InfluenceCost          *int              `json:"influence_cost"`
	InfluenceLimit         *int              `json:"influence_limit"`
	IsUnique               bool              `json:"is_unique"`
	MemoryCost             *int              `json:"memory_cost"`
	MinimumDeckSize        *int              `json:"minimum_deck_size"`
	SideID                 string            `json:"side_id"`
	Strength               *int              `json:"strength"`
	StrippedText           string            `json:"stripped_text"`
	StrippedTitle          string            `json:"stripped_title"`
	Text                   string            `json:"text"`
	Title                  string            `json:"title"`
	TrashCost              *int              `json:"trash_cost"`
	PrintingIDs            []string          `json:"printing_ids"`
	NumPrintings           int               `json:"num_printings"`
	IsLatestPrinting       bool              `json:"is_latest_printing"`
	RestrictionIDs         []string          `json:"restriction_ids"`
	InRestriction          bool              `json:"in_restriction"`
	FormatIDs              []string          `json:"format_ids"`
	CardPoolIDs            []string          `json:"card_pool_ids"`
	SnapshotIDs            []string          `json:"snapshot_ids"`
	CardCycleIDs           []string          `json:"card_cycle_ids"`
	CardSetIDs             []string          `json:"card_set_ids"`
	Attribution            *string           `json:"attribution"`
	ReleasedBy             string            `json:"released_by"`
	PrintingsReleasedBy    []string          `json:"printings_released_by"`
	DesignedBy             string            `json:"designed_by"`
	Pronouns               *string           `json:"pronouns"`
	CardAbilities          *CardAbilities    `json:"card_abilities"`
	Images                 *PrintingImages   `json:"images"`
	LatestPrintingID       string            `json:"latest_printing_id"`
	Restrictions           *CardRestrictions `json:"restrictions"`
}

type PrintingImages struct {
	NRDBClassic *PrintingImageLinks `json:"nrdb_classic"`
}

type PrintingImageLinks struct {
	Tiny   string `json:"tiny"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type PrintingRelationships struct {
	Card        *Relationship `json:"card"`
	CardCycle   *Relationship `json:"card_cycle"`
	CardSet     *Relationship `json:"card_set"`
	CardType    *Relationship `json:"card_type"`
	Faction     *Relationship `json:"faction"`
	Illustrator *Relationship `json:"illustrator"`
	Side        *Relationship `json:"side"`
}

type PrintingFilter struct {
	Params
	CardFilter
	DistinctCards *bool
}

func (filter PrintingFilter) Query() (url.Values, error) {
	query := filter.SetPageInfo(url.Values{})

	if filter.Search != nil {
		query.Set("filter[search]", *filter.Search)
	}

	return query, nil
}

func (doc Printing) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Title, doc.ID)
}

func (cl client) printingReq(filter *PrintingFilter) (*Response[[]*Printing], error) {
	var res Response[[]*Printing]

	var query url.Values
	if filter != nil {
		q, err := filter.Query()
		if err != nil {
			return nil, fmt.Errorf("encoding filter: %w", err)
		}
		query = q
	}

	if err := cl.nrdbReq("printings", &res, query); err != nil {
		return nil, err
	}

	return &res, nil
}
