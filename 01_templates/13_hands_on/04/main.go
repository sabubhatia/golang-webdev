package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

var hm map[string]int = map[string]int{}
var df string

type data struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

var dt []data
var tpl *template.Template

func init() {
	if len(os.Args) < 3 {
		log.Fatalf("Expected at least 3 Args. Got: %d\n", len(os.Args))
	}

	df = os.Args[2]
	if len(df) < 1 {
		log.Fatal("Data file name cannto be empty string\n")
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func main() {

	f, err := os.Open(df)
	if err != nil {
		log.Panicf("Unable to open file %s: %s\n", df, err.Error())
	}
	defer f.Close()
	rdr := csv.NewReader(f)
	cin := reader(rdr)
	appendData(digest(cin), digest(cin))

	log.Println("Num headers: ", len(hm))
	s := struct {
		Headers map[string]int
		Data    []data
	}{
		Headers: hm,
		Data:    dt,
	}

	f = fileutil.OutF("./tmp/", "Table")
	defer fileutil.CloseF(f)
	err = tpl.ExecuteTemplate(f, "Table", s)
	if err != nil {
		log.Panic("Unable to execute template: ", err)
	}
}

func reader(rdr *csv.Reader) <-chan []string {
	cout := make(chan []string)

	// emit file rows/records to the channel
	go func() {
		defer close(cout)
		fr := true
		for {
			r, err := rdr.Read()
			if r == nil && err.Error() == io.EOF.Error() {
				break
			}
			if err != nil {
				log.Panicf("Error when reading file. %s", err)
			}
			if fr {
				log.Println("Processing header..")
				err = header(r)
				if err != nil {
					log.Panicf("Failed during header processing. %v", err)
				}
				// First row is the header so process it and don't emit.
				fr = false
				continue
			}
			log.Println("Processing record..")
			cout <- r
		}
	}()

	return cout
}

func digest(cin <-chan []string) <-chan data {
	cout := make(chan data)

	go func() {
		defer close(cout)
		for r := range cin {
			dr, err := record(r)
			if err != nil {
				log.Panicf("Failed to process %v, %v", r, err)
			}
			cout <- *dr
		}
	}()
	return cout
}

func appendData(cinx ...<-chan data) {

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, cin := range cinx {
		wg.Add(1)
		go func(c <-chan data) {
			defer wg.Done()
			for v := range c {
				log.Println("Appending: ", v)
				mu.Lock()
				dt = append(dt, v)
				mu.Unlock()
			}
		}(cin)
	}
	wg.Wait()
}

func header(sx []string) error {
	if len(sx) < 1 {
		return fmt.Errorf("Header must have at least one column. [%s] is invalid header", sx)
	}
	cnt := 0
	for _, s := range sx {
		if len(s) < 0 {
			return fmt.Errorf("Header cannot be empty string")
		}
		hm[s] = cnt
		cnt++
	}
	return nil
}

func record(sx []string) (*data, error) {

	var d data

	d.Date = sx[0]
	outx := []*float64{
		&d.Open,
		&d.High,
		&d.Low,
		&d.Close,
		&d.Volume,
		&d.AdjClose,
	}
	for i := 0; i < 6; i++ {
		var err error
		(*outx[i]), err = strconv.ParseFloat(sx[i+1], 64)
		if err != nil {
			return nil, fmt.Errorf("Conversion to float64 failed. Value[%s] is not valid for lcoation[%d]", sx[i+1], i)
		}
	}
	return &d, nil
}
