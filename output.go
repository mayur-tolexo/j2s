package jsonToStruct

//User model
type User struct {
	Data     Data     `json:"data"`
	Debug    Debug    `json:"debug"`
	Message  string   `json:"message"`
	MetaData MetaData `json:"metaData"`
	Status   string   `json:"status"`
}

//MetaData model
type MetaData struct {
	UserType string `json:"user_type"`
}

//Debug model
type Debug struct {
	Dcall     Dcall     `json:"dcall"`
	Dmessage  string    `json:"dmessage"`
	Dresponse Dresponse `json:"dresponse"`
}

//Dresponse model
type Dresponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

//Dcall model
type Dcall struct {
	APIBody    string  `json:"apiBody"`
	APITimeout float64 `json:"api_timeout"`
	Headers    Headers `json:"headers"`
	Hostname   string  `json:"hostname"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	Port       string  `json:"port"`
}

//Headers model
type Headers struct {
	CONSUMER            string  `json:"CONSUMER"`
	ContentLength       float64 `json:"Content-Length"`
	ContentType         string  `json:"Content-Type"`
	DEBUGTACHYON        bool    `json:"DEBUG_TACHYON"`
	LOGINCOMPANYID      float64 `json:"LOGIN_COMPANY_ID"`
	LOGINCOMPANYTYPE    string  `json:"LOGIN_COMPANY_TYPE"`
	LOGINUID            float64 `json:"LOGIN_UID"`
	SELECTEDCUSTOMERCID float64 `json:"SELECTED_CUSTOMER_CID"`
	SELECTEDSELLERCID   float64 `json:"SELECTED_SELLER_CID"`
}

//Data model
type Data struct {
	AllowBackorder         bool          `json:"allow_backorder"`
	BrandName              string        `json:"brand_name"`
	Config                 []Config      `json:"config"`
	DisplayPackableQty     bool          `json:"display_packable_qty"`
	DisplayStock           bool          `json:"display_stock"`
	Media                  Media         `json:"media"`
	Moq                    string        `json:"moq"`
	Name                   string        `json:"name"`
	ShowDealerPriceWithTax bool          `json:"show_dealer_price_with_tax"`
	Sku                    []Sku         `json:"sku"`
	Spec                   []interface{} `json:"spec"`
	Tsin                   string        `json:"tsin"`
}

//Sku model
type Sku struct {
	ID                 float64       `json:"_id"`
	ConfigAttr         ConfigAttr    `json:"config_attr"`
	DealerPrice        float64       `json:"dealer_price"`
	DealerPriceInclTax float64       `json:"dealer_price_incl_tax"`
	Discount           float64       `json:"discount"`
	Do                 float64       `json:"do"`
	InStock            bool          `json:"in_stock"`
	Lbt                float64       `json:"lbt"`
	Moq                float64       `json:"moq"`
	MrpPerUnit         float64       `json:"mrp_per_unit"`
	Msrp               float64       `json:"msrp"`
	MysqlArticleID     float64       `json:"mysql_article_id"`
	Name               string        `json:"name"`
	NumberOfItems      string        `json:"number_of_items"`
	PackPrice          float64       `json:"pack_price"`
	PendingQty         float64       `json:"pending_qty"`
	PkgUnit            string        `json:"pkg_unit"`
	Sku                string        `json:"sku"`
	SpPerUnit          float64       `json:"sp_per_unit"`
	Stock              float64       `json:"stock"`
	TierRule           []interface{} `json:"tier_rule"`
}

//ConfigAttr model
type ConfigAttr struct {
	Color string `json:"color"`
	Size  string `json:"size"`
	Sole  string `json:"sole"`
}

//Media model
type Media struct {
	Images []string      `json:"images"`
	VIDeo  []interface{} `json:"video"`
}

//Config model
type Config struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}
