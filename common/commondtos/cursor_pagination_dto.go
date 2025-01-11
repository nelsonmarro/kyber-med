package commondtos

// Cursor representa la información del cursor decodificada
type Cursor struct {
	ID         string `json:"id"`
	CreatedAt  int    `json:"created_at"`  // Unix seconds
	PointsNext bool   `json:"points_next"` // true => “siguiente”, false => “previo”
}

// CursorPaginatedResult es la respuesta con datos + paginación
type PaginationInfo struct {
	NextCursor string `json:"next_cursor,omitempty"`
	PrevCursor string `json:"prev_cursor,omitempty"`
}
