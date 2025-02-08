package commonhelpers

import (
	"encoding/base64"
	"encoding/json"

	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	pEntities "github.com/nelsonmarro/kyber-med/internal/pacient/entities"
)

func EncodeCursor(cursor *commondtos.Cursor) string {
	if cursor == nil {
		return ""
	}
	raw, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(raw)
}

func DecodeCursor(cursor string) (commondtos.Cursor, error) {
	var cd commondtos.Cursor
	if cursor == "" {
		return cd, nil
	}
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return cd, err
	}
	err = json.Unmarshal(b, &cd)
	return cd, err
}

func GetPaginationOperator(pointsNext bool, sortOrder string) (string, string) {
	if pointsNext && sortOrder == "asc" {
		return ">", ""
	}
	if pointsNext && sortOrder == "desc" {
		return "<", ""
	}
	if !pointsNext && sortOrder == "asc" {
		return "<", "desc"
	}
	if !pointsNext && sortOrder == "desc" {
		return ">", "asc"
	}

	return "", ""
}

func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func CalculatePagination(isFirstPage bool, hasPagination bool, limit int, data []pEntities.Pacient, pointsNext bool) commondtos.PaginationInfo {
	pagination := commondtos.PaginationInfo{}
	nextCur := &commondtos.Cursor{}
	prevCur := &commondtos.Cursor{}

	if isFirstPage {
		if hasPagination {
			nextCur.ID = data[limit-1].ID
			nextCur.CreatedAt = data[limit-1].CreatedAt
			nextCur.PointsNext = true

			pagination = GeneratePager(nextCur, nil)
		}
	} else {
		if pointsNext {
			// si apunta a un siguiente siempre va a tener un prev, pero no siempre un next
			if hasPagination {
				nextCur.ID = data[limit-1].ID
				nextCur.CreatedAt = data[limit-1].CreatedAt
				nextCur.PointsNext = true
			} else {
				nextCur = nil
			}
			prevCur.ID = data[0].ID
			prevCur.CreatedAt = data[0].CreatedAt
			prevCur.PointsNext = false

			pagination = GeneratePager(nextCur, prevCur)
		} else {
			nextCur.ID = data[limit-1].ID
			nextCur.CreatedAt = data[limit-1].CreatedAt
			nextCur.PointsNext = true
			if hasPagination {
				prevCur.ID = data[0].ID
				prevCur.CreatedAt = data[0].CreatedAt
				prevCur.PointsNext = false
			} else {
				prevCur = nil
			}
			pagination = GeneratePager(nextCur, prevCur)
		}
	}
	return pagination
}

func GeneratePager(next *commondtos.Cursor, prev *commondtos.Cursor) commondtos.PaginationInfo {
	return commondtos.PaginationInfo{
		NextCursor: EncodeCursor(next),
		PrevCursor: EncodeCursor(prev),
	}
}
