package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hdevillers/go-blast"
	"github.com/hdevillers/go-seq/seq"
)

func main() {
	primer := flag.String("p", "", "Input primer 1.")
	sequence := flag.String("s", "", "Input sequence file (fasta).")
	flag.Parse()

	if *primer == "" {
		panic("You must porvide primer 1 (-p1 argument).")
	}
	if *sequence == "" {
		panic("You must provide an input fasta file containing chromosomes/contigs to investigate (-s argument).")
	}

	if _, err := os.Stat(*sequence); err != nil {
		panic("Failed to find the provided fasta file.")
	}

	// Create the seq.Seq object for the two primers
	pSeq := seq.Seq{
		Id:       "primer",
		Sequence: []byte(*primer),
	}

	// Init. blast search
	blt := blast.NewBlast()
	blt.Db = *sequence
	blt.Par.SetTool("blastn")
	blt.Par.SetTask("blastn")
	blt.Par.SetEvalue(0.001)

	// Run blast on primer
	blt.AddQuery(pSeq)
	blt.Search()
	pRes := blt.Rst.Iterations[0]
	nbHits := len(pRes.Hits)
	if nbHits == 0 {
		fmt.Println("No match found for the provided primer.")
	} else {

	}
}
