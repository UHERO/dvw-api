package data

import (
	"log"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetAdeData(freq string, indicators []string, markets []string, destinations []string) (series Series, err error) {
	//language=MySQL
	var query =
		`select i.handle, m.handle, d.handle, dp.date, dp.value
		 from data_points dp
			join indicators i on i.id = dp.indicator_id
		    join markets m on m.id = dp.market_id
		    join destinations d on d.id = dp.destination_id
		 where dp.module = 'ADE'
		 and dp.frequency = ? `
	var bindVals []interface{}
	bindVals = append(bindVals, freq)
	bindVals = append(bindVals, indicators)
	results, err := Db.Query(query, bindVals...)
	if err != nil {
		log.Printf("Database error: %s", err.Error())
		return
	}
	for results.Next() {

	}
	return
}
