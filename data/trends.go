package data

import (
	"log"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetAdeData(freq string, indicators []string, markets []string, destinations []string) (result SeriesResults, err error) {
	//language=MySQL
	query :=
		`select i.handle, m.handle, d.handle, dp.date, dp.value
		 from data_points dp
		 left join indicators i on i.id = dp.indicator_id
		 left join markets m on m.id = dp.market_id
		 left join destinations d on d.id = dp.destination_id
		 where dp.module = 'ADE'
		 and dp.frequency = ? `
	var bindVals []interface{}
	bindVals = append(bindVals, freq)
	query += "and i.handle in (" + makeQlist(len(indicators)) + ") "
	bindVals = append(bindVals, indicators)
	if len(markets) > 0 {
		query += "and m.handle in (" + makeQlist(len(markets)) + ") "
		bindVals = append(bindVals, markets)
	}
	if len(destinations) > 0 {
		query += "and d.handle in (" + makeQlist(len(destinations)) + ") "
		bindVals = append(bindVals, destinations)
	}
	query += "order by 1,2,3,4" + "" // extra "" only to make GoLand shut up about an error :(

	results, err := Db.Query(query, bindVals...)
	if err != nil {
		log.Printf("Database error: %s", err.Error())
		return
	}
	currentSlug := ""
	var series Series
	for results.Next() {
		scanObs := ScanObsDim3{}
		err = results.Scan(&scanObs.Dim1, &scanObs.Dim2, &scanObs.Dim3, &scanObs.Date, &scanObs.Value)
		if err != nil {
			return
		}
		dims := []string{nil, nil, nil}
		if scanObs.Dim1.Valid {
			dims[0] = scanObs.Dim1.String
		}
		if scanObs.Dim2.Valid {
			dims[1] = scanObs.Dim2.String
		}
		if scanObs.Dim3.Valid {
			dims[2] = scanObs.Dim3.String
		}
		slug := makeSeriesSlug(dims)
		if slug != currentSlug && currentSlug != "" {
			result.SeriesList = append(result.SeriesList, pSeries)
			currentSlug = slug
			continue
		}
		series = Series{}
		series.Columns = dims
		series.Dates = append(pSeries.Dates, scanObs.Date)
		series.Values = append(pSeries.Values, scanObs.Value)
	}
	return
}
