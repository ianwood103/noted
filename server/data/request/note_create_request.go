package request

type NoteCreateRequest struct {
	Text string `valide:"required min=1,max=100" json:"name"`
}