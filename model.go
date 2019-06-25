package pidx

type LoadingDeliveryReceipt struct {
	Header Header `xml:"http://www.pidx.org/schema/ds/v5.01 Header"`
	Body   Body   `xml:"http://www.pidx.org/schema/ds/v5.01 Body"`
}

type Header struct {
	DocumentSequence   string
	DocumentIdentifier string
	DateTime           string
	LastDateTime       string

	From Partner
	To   Partner
}

type Body struct {
	Postings []Posting `xml:"Posting"`
}

type Posting struct {
	ActionCode       string
	PostingType      string
	PostingReference PostingReference
	Activities       []Activity `xml:"Activity"`
	Location         Location
	Transports       []Transport `xml:"Transport"`
}

type Transport struct {
	Company         Partner
	ModeOfTransport ModeOfTransport
	LoadingParts    []LoadingPart `xml:"LoadingParts>LoadingPart"`
}

type LoadingPart struct {
	LoadPartIdentification LoadPartIdentification
	Compartments           []Compartment `xml:"Compartments>Compartment"`
}

type Compartment struct {
	CompartmentNumber      int
	DetailPostingReference PostingReference
	Reference              Reference
	DetailPartners         []Partner          `xml:"DetailPartners>Party"`
	Products               []ProductIdentifer `xml:"Product>ProductIdentifers"`
	Tax                    Tax
	Measurements           []Measurement `xml:"MeasurementInformation>Measurements>Measurement"`
}

type Reference struct {
	ReferenceQualifier  string
	ReferenceIdentifier string
}

type Measurement struct {
	MeasurementType      string
	MeasurementQualifier string
	UnitQualifier        string
	Quantity             float64
	Unit                 string
}

type Tax struct {
	StockStatus   string
	TaxMovemtCode int
}

type ProductIdentifer struct {
	ProductIdentifierType string
	ProductIdentifer      string
}

type LoadPartIdentification struct {
	TpuType    string
	TpuLicense string
	MotItem    int
}
type ModeOfTransport struct {
	MotType string
}

type PostingReference struct {
	DocumentIdentifier string
	IssuedDate         string
}

type Activity struct {
	ActivityType string
	ActivityDate string
}

type Location struct {
	PIDXTerminalIdentifier string
}

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
