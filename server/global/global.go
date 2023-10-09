package global

type Record struct {
	AccessToken  string               `json:"name" form:"name"`
	ProjectId    string               `json:"project_id" form:"name"`
	ExperimentId string               `json:"experiment_id" form:"name"`
	StartAt      string               `json:"start_at" form:"name"`
	Os           map[string]string    `json:"os" form:"name"`
	Trainning    []map[string]float32 `json:"trainning" form:"name"`
	Resource     []map[string]float32 `json:"resource" form:"name"`
}

func (b *Record) AddTrainning() {}
func (b *Record) AddResource()  {}

var R *Record
