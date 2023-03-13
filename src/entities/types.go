package entities

import "encoding/json"

type Result struct {
	Name, Description, URL string
}

type SearchResults struct {
	Ready   bool
	Query   string
	Results []Result
}

func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	var array []interface{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}
