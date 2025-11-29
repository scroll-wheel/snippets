package genome

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Genome struct {
	name           string
	sequence       string
	sequenceLength int
}

func NewGenome(name string, sequence string) Genome {
	return Genome{name: name, sequence: sequence, sequenceLength: len(sequence)}
}

func Load(genomeSource *bufio.Reader, genomes *[]Genome) error {
	// Tip: Write a validator function (in your head) for definitions

	var currentGenome *Genome = nil

	for {
		line, err := genomeSource.ReadString('\n')

		if err != nil && err != io.EOF {
			return err
		}

		line = strings.TrimRight(line, "\r\n")

		if line == "" {
			if err != io.EOF {
				return errors.New("encountered empty line")
			}
		} else if line[0] == '>' {
			if line == ">" {
				return errors.New("encountered empty name")
			}

			if currentGenome != nil {
				if currentGenome.sequenceLength == 0 {
					return errors.New("encountered empty genome sequence")
				}
				*genomes = append(*genomes, *currentGenome)
			}

			currentGenome = &Genome{name: line[1:]}
		} else {
			line = strings.ToUpper(line)

			if strings.Trim(line, "ATCGN") != "" {
				return errors.New("encounted invalid base line")
			}

			if currentGenome == nil {
				return errors.New("encountered base line without corresponding name line")
			}

			currentGenome.sequence += line
			currentGenome.sequenceLength += len(line)
		}

		if err == io.EOF {
			// Turn this into a function that returns error?
			if currentGenome.sequenceLength == 0 {
				return errors.New("encountered empty genome sequence")
			}
			*genomes = append(*genomes, *currentGenome)

			return nil
		}
	}
}

func (g Genome) Length() int {
	return g.sequenceLength
}

func (g Genome) Name() string {
	return g.name
}

func (g Genome) Extract(position int, length int) (string, error) {
	if position+length > g.sequenceLength {
		return "", errors.New("attempted to extract a string that goes past the end of the genome")
	}

	// Assume all characters are A, T, C, G, N
	return g.sequence[position : position+length], nil
}
