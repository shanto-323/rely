package dto

type AcademicContext struct {
	Department *string `query:"department" validate:"omitempty"`
	Shift      *string `query:"shift" validate:"omitempty,oneof=1 2"`
	Semester   *string `query:"semester" validate:"omitempty,numeric,min=1,max=8"`
	Section    *string `query:"section" validate:"omitempty,oneof=A B"`
}

func (a *AcademicContext) GetFilter() map[string]string {
	filter := make(map[string]string)

	if a.Department != nil {
		filter["department"] = *a.Department
	}

	if a.Shift != nil {
		filter["shift"] = *a.Shift
	}

	if a.Semester != nil {
		filter["semester"] = *a.Semester
	}

	if a.Section != nil {
		filter["section"] = *a.Section
	}

	return filter
}

type Pagination struct {
	Page  *int `query:"page" validate:"omitempty,min=1"`
	Limit *int `query:"limit" validate:"omitempty,min=1,max=100"`
}

type SortOrder struct {
	Sort  *string `query:"sort" validate:"omitempty,oneof=created_at updated_at name"`
	Order *string `query:"order" validate:"omitempty,oneof=asc desc"`
}
