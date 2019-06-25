package pidx

// LoadingDeliveryReceipt is the pidx loading delivery receipt
type LoadingDeliveryReceipt struct {
	Header Header `xml:"http://www.pidx.org/schema/ds/v5.01 Header"`
	Body   Body   `xml:"http://www.pidx.org/schema/ds/v5.01 Body"`
}

// Header contains informations about the pidx document
type Header struct {
	DocumentSequence   string
	DocumentIdentifier string
	DateTime           string
	LastDateTime       string

	From Partner
	To   Partner
}

// Body contains information about the loading(s)
type Body struct {
	Postings []Posting `xml:"Posting"`
}

// Posting contains one or more loadings and informations about
// the mode of transport for the loadings
type Posting struct {
	ActionCode       string
	PostingType      string
	PostingReference PostingReference
	Activities       []Activity `xml:"Activity"`
	Location         Location
	Transports       []Transport `xml:"Transport"`
}

// Transport contains the type of transport used for a posting
type Transport struct {
	Company         Partner
	ModeOfTransport ModeOfTransport
	LoadingParts    []LoadingPart `xml:"LoadingParts>LoadingPart"`
}

// LoadingPart contains information about the loading parts of one posting
type LoadingPart struct {
	LoadPartIdentification LoadPartIdentification
	Compartments           []Compartment `xml:"Compartments>Compartment"`
}

// Compartment is a single compartment on a transport vehicle
type Compartment struct {
	CompartmentNumber      int
	DetailPostingReference PostingReference
	Reference              Reference
	DetailPartners         []Partner          `xml:"DetailPartners>Party"`
	Products               []ProductIdentifer `xml:"Product>ProductIdentifers"`
	Tax                    Tax
	Measurements           []Measurement `xml:"MeasurementInformation>Measurements>Measurement"`
}

// Reference ...
type Reference struct {
	ReferenceQualifier  string
	ReferenceIdentifier string
}

// Measurement is a single measurement of a compartment
type Measurement struct {
	MeasurementType      string
	MeasurementQualifier string
	UnitQualifier        string
	Quantity             float64
	Unit                 string
}

// Tax ...
type Tax struct {
	StockStatus   string
	TaxMovemtCode int
}

// ProductIdentifer identifies the product in a compartment
type ProductIdentifer struct {
	ProductIdentifierType string
	ProductIdentifer      string
}

// LoadPartIdentification identifies the loading vehicle
type LoadPartIdentification struct {
	TpuType    string
	TpuLicense string
	MotItem    int
}

// ModeOfTransport defines the transportation mode
type ModeOfTransport struct {
	MotType string
}

// PostingReference ...
type PostingReference struct {
	DocumentIdentifier string
	IssuedDate         string
}

// Activity ...
type Activity struct {
	ActivityType string
	ActivityDate string
}

// Location ...
type Location struct {
	PIDXTerminalIdentifier string
}

// Partner contains informations about the partner
type Partner struct {
	PartnerName          string
	PartnerIdentifier    int
	AddressLine          string
	CityName             string
	PostalCode           string
	PostalCountry        string
	ContactEmail         string
	PartnerRoleIndicator string
}
