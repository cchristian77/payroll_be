package request

type UpsertOvertime struct {
	Duration uint `json:"duration" validate:"required,lte=3"`
}
