package jsonToStruct

//User model
type User struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

//Data model
type Data struct {
	Cart   Cart   `json:"cart"`
	Config Config `json:"config"`
	Error  Error  `json:"error"`
}

//Error model
type Error struct {
	Article ErrorArticle `json:"article"`
	Msg     string       `json:"msg"`
}

//ErrorArticle model
type ErrorArticle struct {
}

//Config model
type Config struct {
	ItemWiseDiscount bool `json:"item_wise_discount"`
}

//Cart model
type Cart struct {
	Documents   Documents `json:"documents"`
	Edd         string    `json:"edd"`
	Info        Info      `json:"info"`
	Instruction string    `json:"instruction"`
	Overdue     Overdue   `json:"overdue"`
	Po          string    `json:"po"`
	PoDate      string    `json:"po_date"`
	Retailer    Retailer  `json:"retailer"`
	Shipper     float64   `json:"shipper"`
	Tnc         Tnc       `json:"tnc"`
}

//Tnc model
type Tnc struct {
	Disclaimer string `json:"disclaimer"`
	Doc        string `json:"doc"`
}

//Retailer model
type Retailer struct {
	Name              string  `json:"name"`
	RetailerCompanyID float64 `json:"retailer_company_id"`
}

//Overdue model
type Overdue struct {
	Message string `json:"message"`
}

//Info model
type Info struct {
	BillAddressID           float64   `json:"bill_address_id"`
	Charge                  Charge    `json:"charge"`
	Discount                Discount  `json:"discount"`
	IsCustomDiscountApplied bool      `json:"is_custom_discount_applied"`
	Price                   float64   `json:"price"`
	Product                 []Product `json:"product"`
	ShipAddressID           float64   `json:"ship_address_id"`
	SubTotal                float64   `json:"sub_total"`
	Tax                     float64   `json:"tax"`
	Total                   float64   `json:"total"`
	TotalQty                float64   `json:"total_qty"`
	Voucher                 string    `json:"voucher"`
}

//Product model
type Product struct {
	Article []Article `json:"article"`
	Brand   string    `json:"brand"`
	BrandID float64   `json:"brand_id"`
	ID      float64   `json:"id"`
	Name    string    `json:"name"`
	Tsin    string    `json:"tsin"`
}

//Article model
type Article struct {
	Cp                    float64        `json:"cp"`
	CpPerUnit             float64        `json:"cp_per_unit"`
	CustomDiscount        CustomDiscount `json:"custom_discount"`
	CustomDiscountApplied string         `json:"custom_discount_applied"`
	Hsn                   string         `json:"hsn"`
	ID                    float64        `json:"id"`
	Image                 string         `json:"image"`
	Lbt                   float64        `json:"lbt"`
	LpPerUnit             float64        `json:"lp_per_unit"`
	Moq                   float64        `json:"moq"`
	MrpPerUnit            float64        `json:"mrp_per_unit"`
	Mxoq                  float64        `json:"mxoq"`
	Name                  string         `json:"name"`
	PkgQty                float64        `json:"pkg_qty"`
	PkgUnit               string         `json:"pkg_unit"`
	Price                 float64        `json:"price"`
	Qty                   float64        `json:"qty"`
	Sku                   string         `json:"sku"`
	SubTotal              float64        `json:"sub_total"`
	Tax                   float64        `json:"tax"`
	TaxableAmount         float64        `json:"taxable_amount"`
	TaxablePercent        float64        `json:"taxable_percent"`
	Total                 float64        `json:"total"`
}

//CustomDiscount model
type CustomDiscount struct {
	DiscountAmount float64 `json:"discount_amount"`
	DiscountName   string  `json:"discount_name"`
}

//Discount model
type Discount struct {
	Rule  []Rule  `json:"rule"`
	Total float64 `json:"total"`
}

//Rule model
type Rule struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

//Charge model
type Charge struct {
	Rule  []interface{} `json:"rule"`
	Total float64       `json:"total"`
}

//Documents model
type Documents struct {
}
