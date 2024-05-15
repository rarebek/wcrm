package models

type CreateCategory struct {
	OwnerId string `json:"owner_id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
}

type Category struct {
	Id        string `json:"id"`
	OwnerId   string `json:"owner_id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
}

type UpdateCategory struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type CategoryList struct {
	Categories []*Category `protobuf:"bytes,1,rep,name=Categories,proto3" json:"categories,omitempty"`
	Count      int64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}
