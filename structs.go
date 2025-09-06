package main

// Root is the top-level structure of Regionen_(MSCI).json
type Root struct {
	Name        string       `json:"name"`
	Color       string       `json:"color"`
	Categories  []Category   `json:"categories"`
	Instruments []Instrument `json:"instruments"`
}

// Category represents a node in the category tree
type Category struct {
	Name     string     `json:"name"`
	Key      string     `json:"key,omitempty"`
	Color    string     `json:"color"`
	Children []Category `json:"children,omitempty"`
}

// Instrument describes one instrument with identifiers and regional weights
type Instrument struct {
	Identifiers InstrumentIdentifiers `json:"identifiers"`
	Categories  []InstrumentCategory  `json:"categories"`
}

// InstrumentCategory links an instrument to a regional path and weight
type InstrumentCategory struct {
	Key    *string  `json:"key,omitempty"`
	Path   []string `json:"path"`
	Weight float64  `json:"weight"`
}

// InstrumentIdentifiers are common instrument IDs
type InstrumentIdentifiers struct {
	Name   string `json:"name"`
	ISIN   string `json:"isin,omitempty"`
	WKN    string `json:"wkn,omitempty"`
	Ticker string `json:"ticker,omitempty"`
}
