package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-shop-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
			Shop:							data.Shop,
			ShopType:						data.ShopType,
			ShopOwner:						data.ShopOwner,
			ShopOwnerBusinessPartnerRole:	data.ShopOwnerBusinessPartnerRole,
			Brand:							data.Brand,
			PersonResponsible:				data.PersonResponsible,
			ValidityStartDate:				data.ValidityStartDate,
			ValidityStartTime:				data.ValidityStartTime,
			ValidityEndDate:				data.ValidityEndDate,
			ValidityEndTime:				data.ValidityEndTime,
			DailyOperationStartTime:		data.DailyOperationStartTime,
			DailyOperationEndTime:			data.DailyOperationEndTime,
			Description:					data.Description,
			LongText:						data.LongText,
			Introduction:					data.Introduction,
			OperationRemarks:				data.OperationRemarks,
			PhoneNumber:					data.PhoneNumber,
			Site:							data.Site,
			Project:						data.Project,
			WBSElement:						data.WBSElement,
			Tag1:							data.Tag1,
			Tag2:							data.Tag2,
			Tag3:							data.Tag3,
			Tag4:							data.Tag4,
			PointConsumptionType:			data.PointConsumptionType,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		Shop:                    dataHeader.Shop,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		Shop:        	dataHeader.Shop,
		AddressID:   	data.AddressID,
		PostalCode:  	data.PostalCode,
		LocalSubRegion: data.LocalSubRegion,
		LocalRegion: 	data.LocalRegion,
		Country:     	data.Country,
		GlobalRegion: 	data.GlobalRegion,
		TimeZone:	 	data.TimeZone,
		District:    	data.District,
		StreetName:  	data.StreetName,
		CityName:    	data.CityName,
		Building:    	data.Building,
		Floor:       	data.Floor,
		Room:        	data.Room,
		XCoordinate: 	data.XCoordinate,
		YCoordinate: 	data.YCoordinate,
		ZCoordinate: 	data.ZCoordinate,
		Site:			data.Site,
	}
}
