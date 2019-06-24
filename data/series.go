package data

import "log"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetAdeAirseatData(module, freq string, indicators, markets, destinations []string) (result SeriesResults, err error) {
	//language=MySQL
	query :=
		`select i.handle, m.handle, d.handle, dp.date, dp.value
		 from data_points dp
		 left join indicators i on i.id = dp.indicator_id
		 left join markets m on m.id = dp.market_id
		 left join destinations d on d.id = dp.destination_id
		 where dp.module = ? and dp.frequency = ? `

	var bindVals []interface{}
	bindVals = append(bindVals, module, freq)
	query += "and i.handle in (" + makeQlist(len(indicators)) + ")\n"
	for _, ind := range indicators {
		bindVals = append(bindVals, ind)
	}
	if len(markets) > 0 {
		query += "and m.handle in (" + makeQlist(len(markets)) + ")\n"
		for _, mkt := range markets {
			bindVals = append(bindVals, mkt)
		}
	}
	if len(destinations) > 0 {
		query += "and d.handle in (" + makeQlist(len(destinations)) + ")\n"
		for _, dest := range destinations {
			bindVals = append(bindVals, dest)
		}
	}
	query += "order by 1,2,3,4" + "\n" // extra "\n" only to make GoLand shut up about an error :(

	dbResults, err := Db.Query(query, bindVals...)
	if err != nil {
		log.Printf("Database error: %s", err.Error())
		return
	}
	currentSlug := ""
	var series Series
	for dbResults.Next() {
		scanObs := ScanObservation{}
		err = dbResults.Scan(&scanObs.Dim1, &scanObs.Dim2, &scanObs.Dim3, &scanObs.Date.Time, &scanObs.Value)
		if err != nil {
			return
		}
		dims := []string{"", "", ""}
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
		if slug != currentSlug {
			if currentSlug != "" {
				result.SeriesList = append(result.SeriesList, series)
			}
			series = Series{}
			series.Columns = dims
			currentSlug = slug
		}
		series.Dates  = append(series.Dates,  scanObs.Date)
		series.Values = append(series.Values, scanObs.Value)

		series.ObsStart.updateIfEarlier(scanObs.Date)
		result.ObsStart.updateIfEarlier(scanObs.Date)
		series.ObsEnd.updateIfLater(scanObs.Date)
		result.ObsEnd.updateIfLater(scanObs.Date)
	}
	result.SeriesList = append(result.SeriesList, series) // the last series being read when the loop ended
	result.Module = module
	result.Frequency = freq
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetHotelData(module, freq string, indicators, categories []string) (result SeriesResults, err error) {
	//language=MySQL
	query :=
		`select i.handle, c.handle, dp.date, dp.value
		 from data_points dp
		 left join indicators i on i.id = dp.indicator_id
		 left join categories c on c.id = dp.category_id
		 where dp.module = ? and dp.frequency = ? `

	var bindVals []interface{}
	bindVals = append(bindVals, module, freq)
	query += "and i.handle in (" + makeQlist(len(indicators)) + ")\n"
	for _, ind := range indicators {
		bindVals = append(bindVals, ind)
	}
	if len(categories) > 0 {
		query += "and c.handle in (" + makeQlist(len(categories)) + ")\n"
		for _, cat := range categories {
			bindVals = append(bindVals, cat)
		}
	}
	query += "order by 1,2,3" + "\n" // extra "\n" only to make GoLand shut up about an error :(

	dbResults, err := Db.Query(query, bindVals...)
	if err != nil {
		log.Printf("Database error: %s", err.Error())
		return
	}
	currentSlug := ""
	var series Series
	for dbResults.Next() {
		scanObs := ScanObservation{}
		err = dbResults.Scan(&scanObs.Dim1, &scanObs.Dim2, &scanObs.Date.Time, &scanObs.Value)
		if err != nil {
			return
		}
		dims := []string{"", ""}
		if scanObs.Dim1.Valid {
			dims[0] = scanObs.Dim1.String
		}
		if scanObs.Dim2.Valid {
			dims[1] = scanObs.Dim2.String
		}
		slug := makeSeriesSlug(dims)
		if slug != currentSlug {
			if currentSlug != "" {
				result.SeriesList = append(result.SeriesList, series)
			}
			series = Series{}
			series.Columns = dims
			currentSlug = slug
		}
		series.Dates  = append(series.Dates,  scanObs.Date)
		series.Values = append(series.Values, scanObs.Value)

		series.ObsStart.updateIfEarlier(scanObs.Date)
		result.ObsStart.updateIfEarlier(scanObs.Date)
		series.ObsEnd.updateIfLater(scanObs.Date)
		result.ObsEnd.updateIfLater(scanObs.Date)
	}
	result.SeriesList = append(result.SeriesList, series) // the last series being read when the loop ended
	result.Module = module
	result.Frequency = freq
	return
}
