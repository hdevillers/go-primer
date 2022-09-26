package main

import (
	"flag"
	"os"

	"github.com/hdevillers/go-blast"
	"github.com/hdevillers/go-seq/seq"
	"github.com/hdevillers/go-seq/utils"
)

func main() {
	primer1 := flag.String("p1", "", "Input primer 1.")
	primer2 := flag.String("p2", "", "Input primer 2.")
	sequence := flag.String("s", "", "Input sequence file (fasta).")
	flag.Parse()

	if *primer1 == "" {
		panic("You must porvide primer 1 (-p1 argument).")
	}

	if *primer2 == "" {
		panic("You must provide primer 2 (-p2 argument).")
	}

	if *sequence == "" {
		panic("You must provide an input fasta file containing chromosomes/contigs to investigate (-s argument).")
	}

	if _, err := os.Stat(*sequence); err != nil {
		panic("Failed to find the provided fasta file.")
	}

	// Create the seq.Seq object for the two primers
	p1Seq := seq.Seq{
		Id:       "primer1",
		Sequence: []byte(*primer1),
	}
	p2Seq := seq.Seq{
		Id:       "primer2",
		Sequence: []byte(*primer2),
	}

	// Load Chromosomes/contigs
	var chr map[string]seq.Seq
	nchr := utils.LoadSeqInMap(*sequence, "fasta", &chr)
	if nchr == 0 {
		panic("No data found in provided fasta file.")
	}

	// Init. blast search
	blt := blast.NewBlast()
	blt.Db = *sequence
	blt.Par.SetTool("blastn")
	blt.Par.SetEvalue(0.0001)

	// Treat primer 1
	blt.AddQuery(p1Seq)
	blt.Search()
	p1Res := blt.Rst.Iterations
	if len(p1Res) == 0 {
		panic("No it found for primer 1.")
	}
}
