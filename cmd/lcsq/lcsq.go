package main

import "fmt"
import "log"
import "os"
import "phylogenetics/gene"
import "phylogenetics/strings"

func main() {
	fasta, err := gene.ReadAllFasta(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	s := fasta[0].Data
	t := fasta[1].Data
	lcs := strings.LongestCommonSubsequence(s, t)
	fmt.Println(lcs)
}
