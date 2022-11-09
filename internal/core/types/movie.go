package types

// Book defines an example domain resource
type Movie struct {
	ID string `json:"id"`
	MovieProperties
}

// Movie defines an example domain resource
type MovieProperties struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Year     uint16 `json:"year"`
	Director string `json:"director"`
	Deleted  bool   `json:"deleted"`
	Status   int    `json:"status"`
}

// Discover
func (m *Movie) Discover() *Movie {
	return m
}

// // SerializeModel
// func (m *Movie) SerializeModel(r *MovieRepoResult) (*Movie, error) {
// 	return &Movie{}, nil
// }

// SerializeResponse
func (m *Movie) SerializeResponse(r *MovieRepoResult, single bool) (JSONResponse, error) {
	if single {
		res := &JSONResponseDetail{
			Data: &Resource{
				Type:       "movie",
				ID:         "1234",
				Properties: r.Data[0].Properties,
			},
		}
		return res, nil
	} else {
		res := &JSONResponseList{
			Meta: &APIMetadata{
				Paging: &ListPaging{
					Limit:  r.Metadata.Paging.Limit,
					Offset: r.Metadata.Paging.Offset,
					Total:  r.Metadata.Paging.Total,
				},
			},
			Data: &[]Resource{},
		}
		return res, nil
	}
}
