package app

func toProto(n *EditorNote) *Note {
	return &Note{
		Body: n.Body,
	}
}
