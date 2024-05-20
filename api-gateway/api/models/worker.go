package models

type Worker struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	LoginKey  string `json:"login_key"`
	Password  string `json:"password"`
	OwnerId   string `json:"owner_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type WorkerList struct {
	Workers []Worker `json:"workers"`
	Count   int64    `json:"count"`
}

type CreateWorker struct {
	FullName string `json:"full_name"`
	LoginKey string `json:"login_key"`
	Password string `json:"password"`
	OwnerId  string `json:"owner_id"`
}

type UpdateWorker struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	LoginKey string `json:"login_key"`
	Password string `json:"password"`
	OwnerId  string `json:"owner_id"`
}

// type CheckResponse struct {
// 	Check bool `json:"chack"`
// }
