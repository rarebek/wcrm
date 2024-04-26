package models

type Owner struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Tax         string `json:"tax"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type OwnerList struct {
	Owners []Owner `json:"owners"`
}

type CreateOwner struct {
	FullName    string `json:"full_name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Tax         string `json:"tax"`
}

type UpdateOwner struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Tax         string `json:"tax"`
}

// type CheckResponse struct {
// 	Check bool `json:"chack"`
// }
