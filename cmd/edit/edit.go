package main

import "flag"
import "fmt"
import "log"
import "os"
import "phylogenetics/gene"
import "phylogenetics/strings"

func main() {
	debug := flag.Bool("debug", false, "Show debug output")
	flag.Parse()

	strings.DebugEditDistance = *debug

	fasta, err := gene.ReadAllFasta(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	s := fasta[0].Data
	t := fasta[1].Data
	d := strings.EditDistance(s, t)
	fmt.Println(d)
}
