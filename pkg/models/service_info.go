package models

type ServiceInfo struct {
	PatientInfo     PatientInfo `json:"patient_info"`
	ID              int         `json:"id"`
	IntoNumber      string      `json:"into_number"`
	ServiceCode     string      `json:"service_code"`
	RighToTreatment string      `json:"righ_to_treatment"`
	ServicrType     string      `json:"servicr_type"`
	Activity        string      `json:"activity"`
	Location        string      `json:"location"`
	Sine            string      `json:"sine"`
	End             string      `json:"end_"`
	Period          string      `json:"period"`
}
