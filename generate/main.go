package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"text/template"
	"time"

	"bitbucket.org/metricaid/rnd/timezone_info"
)

const (
	workDir     = "tzdb"
	dbFilename  = "timezonedb.csv.zip"
	dbURL       = "https://timezonedb.com/files/" + dbFilename
	countryFile = "country.csv"
	zoneFile    = "zone.csv"
	outputFile  = "../tz_data.go"
)

type countryColumn int

// Country Columns
const (
	countryCode countryColumn = iota
	countryName
)

type zoneColumn int

// Zone Columns
const (
	ID zoneColumn = iota
	code
	name
)

type byCountryName []tz.Country

func (a byCountryName) Len() int           { return len(a) }
func (a byCountryName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCountryName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type byZoneName []tz.Zone

func (a byZoneName) Len() int           { return len(a) }
func (a byZoneName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byZoneName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {

	tmpl, err := template.New("gen").Parse(output)
	if err != nil {
		log.Fatal("ERROR parsing template:", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("ERROR determining current working DIR:", err)
	}

	cmd := exec.Command("mkdir", "-p", workDir)
	err = cmd.Run()
	if err != nil {
		log.Fatal("ERROR creating working DIR:", err)
	}

	err = os.Chdir(workDir)
	if err != nil {
		log.Fatal("ERROR switching to working DIR:", err)
	}

	cmd = exec.Command("curl", "-O", dbURL)

	err = cmd.Run()
	if err != nil {
		log.Fatal("ERROR downloading file:", err)
	}

	cmd = exec.Command("unzip", dbFilename)

	err = cmd.Run()
	if err != nil {
		log.Fatal("ERROR Unzipping file:", err)
	}

	countries, err := process()
	if err != nil {
		log.Fatal("ERROR processing files:", err)
	}

	err = os.Chdir(cwd)
	if err != nil {
		log.Fatal("ERROR switching to original working DIR:", err)
	}

	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("ERROR writing/creating tz data file:", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, countries)
	if err != nil {
		log.Fatal("ERROR executing template:", err)
	}

	f.Close()

	// after file written run gofmt on file
	cmd = exec.Command("gofmt", "-s", "-w", outputFile)
	if err = cmd.Run(); err != nil {
		log.Fatal("ERROR running gofmt:", err)
	}

	err = os.RemoveAll(workDir)
	if err != nil {
		log.Fatal("ERROR removing working DIR:", err)
	}

}

func process() ([]tz.Country, error) {

	var err error

	cmap := make(map[string]int)
	countries := make([]tz.Country, 0, 10)

	// process countries
	cf, err := os.Open(countryFile)
	if err != nil {
		return nil, err
	}
	defer cf.Close()

	r := csv.NewReader(cf)

	for {

		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		c := tz.Country{
			Code: row[countryCode],
			Name: row[countryName],
		}
		cmap[c.Code] = len(countries)

		countries = append(countries, c)
	}

	// process zones

	zf, err := os.Open(zoneFile)
	if err != nil {
		return nil, err
	}
	defer zf.Close()

	r = csv.NewReader(zf)

	for {

		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		z := tz.Zone{
			CountryCode: row[code],
			Name:        row[name],
		}

		// test zone is working in Go
		_, err = time.LoadLocation(z.Name)
		if err != nil {
			fmt.Println("*********************ERROR:", err)
			continue
		}

		idx, ok := cmap[z.CountryCode]
		if !ok {
			continue
		}

		countries[idx].Zones = append(countries[idx].Zones, z)
	}

	// sort alphabetically
	sort.Sort(byCountryName(countries))

	for _, c := range countries {
		sort.Sort(byZoneName(c.Zones))
	}

	return countries, nil
}

var output = `package tz

import "sync"

// GENERATED FILE DO NOT MODIFY DIRECTLY

var (
	once      sync.Once
	mapped    map[string]Country
	countries = []Country{
			{{ range $c := .}}{
				Code: "{{ $c.Code }}",
				Name: "{{ $c.Name }}",
				Zones: []Zone{
					{{ range $z := $c.Zones }}{
						CountryCode: "{{ $z.CountryCode }}",
						Name: "{{ $z.Name }}",
					},
					{{ end }}
				},
			},
			{{ end }}
	}
)

func init() {
	// load + index countries into map
	// for below functions.

	once.Do(func() {
		mapped = make(map[string]Country)

		for i := 0; i < len(countries); i++ {
			mapped[countries[i].Code] = countries[i]
		}
	})
}

// GetCountries returns an array of all countries.
// Most common use: for loading into a country dropdown
// in HTML.
func GetCountries() []Country {
	return countries
}

// GetCountry returns a single Country that matches the country
// code passed and whether it was found
func GetCountry(code string) (c Country, found bool) {
	c, found = mapped[code]
	return
}
`

// func main() {

// 	time.Local = time.UTC

// 	loc, err := time.LoadLocation("America/Toronto")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	utc := time.Now()

// 	fmt.Println("   NOW UTC:", utc)

// 	local := utc.In(loc)
// 	fmt.Println("LOCAL TIME:", local)

// 	edt, err := time.Parse("2006-01-02", "2016-04-01")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	est, err := time.Parse("2006-01-02", "2016-12-01")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	fmt.Println("EDT UTC:", edt)
// 	fmt.Println("EST UTC:", est)

// 	fmt.Println("EDT LOCAL:", edt.In(loc))
// 	fmt.Println("EST LOCAL:", est.In(loc))
// }
