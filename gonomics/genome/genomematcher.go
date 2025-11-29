package genome

import (
	"errors"

	"gonomics/trie"
)

type position struct {
	genome Genome
	index  int
}

type GenomeMatcher struct {
	minSearchLength int
	genomes         []Genome
	sequences       trie.Trie[position]
}

type DNAMatch struct {
	GenomeName string
	Position   int
	Length     int
}

type GenomeMatch struct {
	GenomeName   string
	PercentMatch float64
}

func NewGenomeMatcher(minSearchLength int) GenomeMatcher {
	return GenomeMatcher{minSearchLength: minSearchLength}
}

func (gm *GenomeMatcher) AddGenome(genome Genome) {
	gm.genomes = append(gm.genomes, genome)

	for i := 0; i+gm.minSearchLength <= genome.Length(); i++ {
		sequence, err := genome.Extract(i, gm.minSearchLength)

		if err != nil {
			panic(err)
		}

		gm.sequences.Insert(sequence, position{genome: genome, index: i})
	}
}

func (gm GenomeMatcher) MinimumSearchLength() int {
	return gm.minSearchLength
}

// TODO: Make this function return []DNAMatch?
func (gm GenomeMatcher) FindGenomesWithThisDNA(fragment string, minimumLength int, exactMatchOnly bool, matches *[]DNAMatch) error {
	// Assumption: gm.minSearchLength <= minimumLength <= len(fragment)

	if len(fragment) < minimumLength {
		return errors.New("fragment's length is less than minimumLength")
	}

	if minimumLength < gm.minSearchLength {
		return errors.New("minimumLength is less than GenomeMatcher's minSearchLength")
	}

	m := make(map[Genome]DNAMatch)

	for _, p := range gm.sequences.Find(fragment[:gm.MinimumSearchLength()], exactMatchOnly) {
		sequence, _ := p.genome.Extract(p.index, len(fragment))

		exact := exactMatchOnly
		matchLength := 1
		for matchLength < len(fragment) {
			if sequence[matchLength] != fragment[matchLength] {
				if !exact {
					exact = true
				} else {
					break
				}
			}
			matchLength++
		}

		if matchLength >= minimumLength {
			if v, prs := m[p.genome]; !prs || v.Length < matchLength || (v.Length == matchLength && v.Position > p.index) {
				match := DNAMatch{GenomeName: p.genome.Name(), Position: p.index, Length: matchLength}
				m[p.genome] = match
			}
		}
	}

	for _, v := range m {
		*matches = append(*matches, v)
	}

	return nil
}

// TODO: Make this function return []GenomeMatch?
func (gm GenomeMatcher) FindRelatedGenomes(query Genome, fragmentMatchLength int, exactMatchOnly bool, matchPercentThreshold float64, results *[]GenomeMatch) error {
	if fragmentMatchLength < gm.MinimumSearchLength() {
		return errors.New("fragmentMatchLength is less than GenomeMatcher's minSearchLength")
	}

	s := query.Length() / fragmentMatchLength
	dnaMatches := []DNAMatch{}
	for i := 0; i < s; i++ {
		sequence, err := query.Extract(i*fragmentMatchLength, fragmentMatchLength)

		if err != nil {
			return err
		}

		gm.FindGenomesWithThisDNA(sequence, fragmentMatchLength, exactMatchOnly, &dnaMatches)
	}

	count := make(map[string]int)
	for _, match := range dnaMatches {
		count[match.GenomeName]++
	}

	for k, v := range count {
		percent := float64(v) / float64(s)
		if percent > matchPercentThreshold {
			match := GenomeMatch{GenomeName: k, PercentMatch: percent}
			*results = append(*results, match)
		}
	}

	return nil
}

// Stack - Automatic deallocation
// Heap - Garbage collection

// Slice internals
//	Pointer to array
// 		O(1) - Stack
// 		O(N) - Heap
//	Length
//	Capacity

// Strings ~ slices
