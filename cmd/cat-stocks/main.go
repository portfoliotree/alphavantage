package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := getData("TIME_SERIES_WEEKLY", "MSFT")
	if err != nil {
		log.Fatal(err)
	}

	data = data[len(data)-100:]

	var (
		xMin, xMax = data[0].Timestamp.Unix(), data[len(data)-1].Timestamp.Unix()

		yMin, yMax float64 = math.MaxFloat64, 0
	)

	for _, d := range data {
		y := d.Mid()
		if y > yMax {
			yMax = y
		}
		if y < yMin {
			yMin = y
		}
	}

	fmt.Println(xMin, xMax)
	fmt.Println(yMin, yMax)
	fmt.Println()

	var corrdinates []string
	for _, d := range data {
		cx := float64(d.Timestamp.Unix()-xMin) / float64(xMax-xMin) * 100
		cy := (1 - (d.Mid()-yMin)/(yMax-yMin)) * 100
		corrdinates = append(corrdinates, fmt.Sprintf("%f,%f", float64(cx), cy))
	}

	f, _ := os.Create("index.html")
	fmt.Fprintf(f, tmpl, strings.Join(corrdinates, " "))
}

const (
	tmpl = `<!DOCTYPE html>
<html>
<body>
	<svg class="chart" viewBox="0 0 100 100">
		<polyline
	     fill="none""
	     stroke="#0074d9"
	     stroke-width=".05"
	     points=%q/>
	</svg>
</body>
</html>`
)

func getData(fn, symbol string) (Data, error) {
	var data Data

	// query := make(url.Values)
	// query.Add("apikey", os.Getenv("KEY"))
	// query.Add("function", fn)
	// query.Add("symbol", symbol)
	//
	// href := "https://www.alphavantage.co/query?" + query.Encode()
	// fmt.Println(href)
	// res, err := http.Get(href)
	// if err != nil {
	// 	return data, err
	// }
	// if res.StatusCode >= 300 {
	// 	log.Fatalf("not success code %d", res.StatusCode)
	// }

	f, _ := os.Open("msft.json")

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return data, err
	}

	var body map[string]interface{}
	if err := json.Unmarshal(buf, &body); err != nil {
		return data, err
	}

	var set Data

	delete(body, "Meta Data")

	for date, numbers := range body["Weekly Time Series"].(map[string]interface{}) {
		var datum Datum

		datum.Timestamp, _ = time.Parse("2006-01-02", date) // " 15:04:05"

		num := make(map[string]interface{})
		for k, v := range numbers.(map[string]interface{}) {
			num[strings.Join(strings.Fields(k)[1:], "_")] = v
		}

		datum.Open = getFloat(num, "open")
		datum.High = getFloat(num, "high")
		datum.Low = getFloat(num, "low")
		datum.Close = getFloat(num, "close")
		datum.Volume = getFloat(num, "volume")
		datum.AdjustedClose = getFloat(num, "adjusted_close")
		datum.DividendAmount = getFloat(num, "dividend_amount")

		set = append(set, datum)
	}

	sort.Sort(set)

	return set, err
}

func getFloat(mp map[string]interface{}, key string) float64 {
	v, ok := mp[key]
	if !ok {
		return math.NaN()
	}
	str, ok := v.(string)
	if !ok {
		return math.NaN()
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return math.NaN()
	}
	return f
}

type Datum struct {
	Timestamp time.Time
	Open, High, Low, Close, Volume,
	AdjustedClose, DividendAmount float64
}

func (datum Datum) Mid() float64 {
	return datum.High - datum.Low
}

type Data []Datum

func (data Data) Len() int { return len([]Datum(data)) }

func (data Data) Less(i, j int) bool { return data[i].Timestamp.Before(data[j].Timestamp) }

func (data Data) Swap(i, j int) { data[i], data[j] = data[j], data[i] }

func (data Data) XY(index int) (x, y float64) {
	return float64(data[index].Timestamp.Unix()), data[index].Low
}
