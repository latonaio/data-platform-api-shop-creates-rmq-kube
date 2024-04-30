package requests

type Header struct {
	Shop							int		`json:"Shop"`
	ShopType						string	`json:"ShopType"`
	ShopOwner						int		`json:"ShopOwner"`
	ShopOwnerBusinessPartnerRole	string	`json:"ShopOwnerBusinessPartnerRole"`
	Brand							*int	`json:"Brand"`
	PersonResponsible				string	`json:"PersonResponsible"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityStartTime				string	`json:"ValidityStartTime"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	ValidityEndTime					string	`json:"ValidityEndTime"`
	DailyOperationStartTime			string	`json:"DailyOperationStartTime"`
	DailyOperationEndTime			string	`json:"DailyOperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	OperationRemarks				*string	`json:"OperationRemarks"`
	PhoneNumber						*string	`json:"PhoneNumber"`
	Site							int		`json:"Site"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	PointConsumptionType      		string  `json:"PointConsumptionType"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsReleased						*bool	`json:"IsReleased"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
