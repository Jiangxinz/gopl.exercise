package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// byArtist
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// custom
type cmpFunc func(lhs, rhs *Track) bool
type swapFunc func(lhs, rhs *Track)

type custom struct {
	t   []*Track
	cmp cmpFunc
	// swap swapFunc
}

func (c *custom) Len() int           { return len(c.t) }
func (c *custom) Less(i, j int) bool { return c.cmp(c.t[i], c.t[j]) }
func (c *custom) Swap(i, j int)      { c.t[i], c.t[j] = c.t[j], c.t[i] }

var cmpMap = map[string]cmpFunc{
	"Title":  func(lhs, rhs *Track) bool { return lhs.Title < rhs.Title },
	"Artist": func(lhs, rhs *Track) bool { return lhs.Artist < rhs.Artist },
	"Album":  func(lhs, rhs *Track) bool { return lhs.Album < rhs.Album },
	"Year":   func(lhs, rhs *Track) bool { return lhs.Year < rhs.Year },
	"Length": func(lhs, rhs *Track) bool { return lhs.Length < rhs.Length },
}

func click(column string) {
	cmp, ok := cmpMap[column]
	if !ok {
		return
	}
	sort.Stable(&custom{tracks, cmp})
}

func clickAndPrint(column string) {
	fmt.Println("by ", column)
	click(column)
	printTracks(tracks)
}

func main() {
	printTracks(tracks)
	// sort.Stable(byArtist(tracks))
	sort.Sort(byArtist(tracks))

	clickAndPrint("Artist")
	clickAndPrint("Album")
	clickAndPrint("Length")
}
