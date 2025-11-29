package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"gonomics/genome"
)

const PROVIDED_DIR = "data"

var providedFiles = []string{
	"Ferroplasma_acidarmanus.txt",
	"Halobacterium_jilantaiense.txt",
	"Halorubrum_chaoviator.txt",
	"Halorubrum_californiense.txt",
	"Halorientalis_regularis.txt",
	"Halorientalis_persicus.txt",
	"Ferroglobus_placidus.txt",
	"Desulfurococcus_mucosus.txt",
}

func createNewLibrary(library *genome.GenomeMatcher) {
	fmt.Print("Enter minimum search length (3-100): ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	len, err := strconv.Atoi(line)
	if err != nil || len < 3 || len > 100 {
		fmt.Println("Invalid prefix size.")
		return
	}

	*library = genome.NewGenomeMatcher(len)
}

func addOneGenomeManually(library *genome.GenomeMatcher) {
	fmt.Print("Enter name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		fmt.Println("Name must not be empty.")
		return
	}

	fmt.Print("Enter DNA sequence: ")
	sequence, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sequence = strings.TrimSpace(sequence)
	if len(sequence) == 0 {
		fmt.Println("Sequence must not be empty.")
		return
	}
	if strings.Trim(sequence, "ACGTNacgtn") != "" {
		fmt.Println("Invalid character in DNA sequence.")
		return
	}
	sequence = strings.ToUpper(sequence)
	library.AddGenome(genome.NewGenome(name, sequence))
}

func loadFile(filename string, genomes *[]genome.Genome) bool {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file:", filename)
		return false
	}

	genomeSource := bufio.NewReader(f)
	if err = genome.Load(genomeSource, genomes); err != nil {
		fmt.Println("Improperly formatted file:", filename)
		return false
	}
	return true
}

func loadOneDataFile(library *genome.GenomeMatcher) {
	fmt.Print("Enter file name: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	filename = strings.TrimSpace(filename)
	if len(filename) == 0 {
		fmt.Println("No file name entered.")
		return
	}

	genomes := []genome.Genome{}
	if !loadFile(filename, &genomes) {
		return
	}
	for _, g := range genomes {
		library.AddGenome(g)
	}
	fmt.Println("Successfully loaded", len(genomes), "genomes.")
}

func loadProvidedFiles(library *genome.GenomeMatcher) {
	for _, f := range providedFiles {
		genomes := []genome.Genome{}
		if loadFile(PROVIDED_DIR+"/"+f, &genomes) {
			for _, g := range genomes {
				library.AddGenome(g)
			}
			fmt.Println("Loaded", len(genomes), "genomes from", f)
		}
	}
}

func findGenome(library *genome.GenomeMatcher, exactMatch bool) {
	if exactMatch {
		fmt.Print("Enter DNA sequence for which to find exact matches: ")
	} else {
		fmt.Print("Enter DNA sequence for which to find exact matches and SNiPs: ")
	}
	reader := bufio.NewReader(os.Stdin)
	sequence, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sequence = strings.TrimSpace(sequence)
	minLength := library.MinimumSearchLength()
	if len(sequence) < minLength {
		fmt.Println("DNA sequence length must be at least", minLength)
		return
	}

	fmt.Print("Enter minimum sequence match length: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	minMatchLength, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if minMatchLength > len(sequence) {
		fmt.Println("Minimum match length must be at least the sequence length.")
		return
	}
	matches := []genome.DNAMatch{}
	if err := library.FindGenomesWithThisDNA(sequence, minMatchLength, exactMatch, &matches); err != nil {
		fmt.Print("No matches ")
		if !exactMatch {
			fmt.Print("and/or SNiPs ")
		}
		fmt.Println("of", sequence, "were found.")
		return
	}
	fmt.Print(len(matches), " matches ")
	if !exactMatch {
		fmt.Print("and/or SNiPs ")
	}
	fmt.Println("of", sequence, "were found:")
	for _, m := range matches {
		fmt.Println("  length", m.Length, "position", m.Position, "in", m.GenomeName)
	}
}

func getFindRelatedParams(pct *float64, exactMatchOnly *bool) bool {
	fmt.Print("Enter match percentage threshold (0-100): ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	num, err := strconv.Atoi(line)
	if err != nil || num < 3 || num > 100 {
		fmt.Println("Percentage must be in the range 0 to 100.")
		return false
	}
	*pct = float64(num)

	fmt.Print("Require (e)xact match or allow (S)NiPs (e or s): ")
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	switch rn, _ := utf8.DecodeRuneInString(line); unicode.ToLower(rn) {
	case 'e':
		*exactMatchOnly = true
	case 's':
		*exactMatchOnly = false
	default:
		fmt.Println("Response must be e or s.")
		return false
	}
	return true
}

func findRelatedGenomesManual(library *genome.GenomeMatcher) {
	fmt.Print("Enter DNA sequence: ")
	reader := bufio.NewReader(os.Stdin)
	sequence, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sequence = strings.TrimSpace(sequence)
	minLength := library.MinimumSearchLength()
	if len(sequence) < minLength {
		fmt.Println("DNA sequence length must be at least", minLength)
	}
	var pctThreshold float64
	var exactMatchOnly bool
	if !getFindRelatedParams(&pctThreshold, &exactMatchOnly) {
		return
	}

	matches := []genome.GenomeMatch{}
	library.FindRelatedGenomes(genome.NewGenome("x", sequence), 2*minLength, exactMatchOnly, pctThreshold, &matches)
	if len(matches) == 0 {
		fmt.Println("    No related genomes were found")
		return
	}
	fmt.Println("   ", len(matches), "related genomes were found:")
	for _, m := range matches {
		fmt.Println("", m.PercentMatch, "%", m.GenomeName)
	}
}

func findRelatedGenomesFromFile(library *genome.GenomeMatcher) {
	fmt.Print("Enter name of file containing one or more genomes to find matches for: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	filename = strings.TrimSpace(filename)
	if len(filename) == 0 {
		fmt.Println("No file name entered.")
		return
	}
	genomes := []genome.Genome{}
	if !loadFile(filename, &genomes) {
		return
	}
	var pctThreshold float64
	var exactMatchOnly bool
	if !getFindRelatedParams(&pctThreshold, &exactMatchOnly) {
		return
	}

	minLength := library.MinimumSearchLength()
	for _, g := range genomes {
		matches := []genome.GenomeMatch{}
		library.FindRelatedGenomes(g, 2*minLength, exactMatchOnly, pctThreshold, &matches)
		fmt.Println("  For", g.Name())
		if len(matches) == 0 {
			fmt.Println("    No related genomes were found")
			continue
		}
		fmt.Println("   ", len(matches), "related genomes were found:")
		for _, m := range matches {
			fmt.Println("    ", m.PercentMatch, "%", m.GenomeName)
		}
	}
}

func showMenu() {
	fmt.Println("        Commands:")
	fmt.Println("         c - create new genome library      s - find matching SNiPs")
	fmt.Println("         a - add one genome manually        r - find related genomes (manual)")
	fmt.Println("         l - load one data file             f - find related genomes (file)")
	fmt.Println("         d - load all provided data files   ? - show this menu")
	fmt.Println("         e - find matches exactly           q - quit")
}

func main() {
	const defaultMinSearchLength = 10

	fmt.Println("Welcome to the Gee-nomics test harness!")
	fmt.Println("The genome library is initially empty, with a default minSearchLength of", defaultMinSearchLength)
	showMenu()

	library := genome.NewGenomeMatcher(defaultMinSearchLength)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		command, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		switch rn, _ := utf8.DecodeRuneInString(command); unicode.ToLower(rn) {
		case 'q':
			return
		case '?':
			showMenu()
		case 'c':
			createNewLibrary(&library)
		case 'a':
			addOneGenomeManually(&library)
		case 'l':
			loadOneDataFile(&library)
		case 'd':
			loadProvidedFiles(&library)
		case 'e':
			findGenome(&library, true)
		case 's':
			findGenome(&library, false)
		case 'r':
			findRelatedGenomesManual(&library)
		case 'f':
			findRelatedGenomesFromFile(&library)
		default:
			fmt.Println("Invalid command", command)
		}
	}
}
