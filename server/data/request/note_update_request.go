package request

type NoteUpdateRequest struct {
	Id int
	Text string `valide:"required min=1,max=100" json:"name"`
}