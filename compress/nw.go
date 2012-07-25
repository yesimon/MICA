package main

/*
This is adapted from Dan Kortschak's Needleman-Wunsch algorithm from the biogo
package: code.google.com/p/biogo.

It's mostly copied from its original form, but it is optimized specifically
for cablastp to limit allocations and to absolve the need for the biogo/seq.Seq
type.
*/

import (
	"github.com/BurntSushi/cablastp/blosum"

	"code.google.com/p/biogo/align/nw"
	"code.google.com/p/biogo/util"
)

const (
	memSeqSize       = 10000
	dynamicTableSize = memSeqSize * memSeqSize
)

var (
	nwLookUpP util.CTL
	aligner   = &nw.Aligner{
		Matrix:  blosum.Matrix62,
		LookUp:  nwLookUpP,
		GapChar: '-',
	}
)

func init() {
	m := make(map[int]int)
	for i, v := range blosum.Alphabet62 {
		m[int(v)] = i
	}
	nwLookUpP = *util.NewCTL(m)

	aligner.LookUp = nwLookUpP
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type nwMemory struct {
	table    []int
	ref, org []byte
}

func newNwMemory() nwMemory {
	return nwMemory{
		table: make([]int, memSeqSize*memSeqSize),
		ref:   make([]byte, memSeqSize),
		org:   make([]byte, memSeqSize),
	}
}

func appendOne(slice []byte, b byte) []byte {
	if len(slice) == cap(slice) {
		return append(slice, b)
	}
	slice = slice[0 : len(slice)+1]
	slice[len(slice)-1] = b
	return slice
}

func nwAlign(rseq, oseq []byte, mem nwMemory) [2][]byte {
	gap := len(aligner.Matrix) - 1
	r, c := len(rseq)+1, len(oseq)+1
	off := 0

	var table []int
	if r*c > dynamicTableSize {
		table = make([]int, r*c)
	} else {
		table = mem.table[:r*c]
		for i := range table {
			table[i] = 0
		}
	}

	constrained := true
	constraint := r / 4
	if r <= 11 || c <= 11 {
		constrained = false
	}

	var sdiag, sup, sleft, rVal, oVal int
	var j int
	valToCode := aligner.LookUp.ValueToCode
	gapChar := aligner.GapChar
	matrix := aligner.Matrix

	var i2, i3 int
	for i := 1; i < r; i++ {
		i2 = (i - 1) * c
		i3 = i * c
		for j = 1; j < c; j++ {
			if constrained && ((i-j) > constraint || (j-i) > constraint) {
				continue
			}
			rVal, oVal = valToCode[rseq[i-1]], valToCode[oseq[j-1]]
			if rVal < 0 || oVal < 0 {
				continue
			} else {
				off = i2 + (j - 1)
				sdiag = table[off] + matrix[rVal][oVal]
				sup = table[off+1] + matrix[rVal][gap]
				sleft = table[off+c] + matrix[gap][oVal]
				switch {
				case sdiag > sup && sdiag > sleft:
					table[i3+j] = sdiag
				case sup > sleft:
					table[i3+j] = sup
				default:
					table[i3+j] = sleft
				}
			}
		}
	}

	refAln, orgAln := mem.ref[:0], mem.org[:0]

	i, j := r-1, c-1
	for i > 0 && j > 0 {
		rVal, oVal = valToCode[rseq[i-1]], valToCode[oseq[j-1]]
		if rVal < 0 || oVal < 0 {
			continue
		} else {
			sdiag = table[(i-1)*c+(j-1)] + matrix[rVal][oVal]
			sup = table[(i-1)*c+j] + matrix[gap][oVal]
			sleft = table[i*c+(j-1)] + matrix[rVal][gap]
			switch {
			case sdiag > sup && sdiag > sleft:
				i--
				j--
				refAln = appendOne(refAln, rseq[i])
				orgAln = appendOne(orgAln, oseq[j])
			case sup > sleft:
				i--
				refAln = appendOne(refAln, rseq[i])
				orgAln = appendOne(orgAln, gapChar)
			default:
				j--
				refAln = appendOne(refAln, gapChar)
				orgAln = appendOne(orgAln, oseq[j])
			}
		}
	}

	for ; i > 0; i-- {
		refAln = appendOne(refAln, rseq[i-1])
		orgAln = appendOne(orgAln, gapChar)
	}
	for ; j > 0; j-- {
		refAln = appendOne(refAln, gapChar)
		orgAln = appendOne(orgAln, oseq[j-1])
	}

	if len(refAln) == len(orgAln) {
		for i, j := 0, len(refAln)-1; i < j; i, j = i+1, j-1 {
			refAln[i], refAln[j] = refAln[j], refAln[i]
			orgAln[i], orgAln[j] = orgAln[j], orgAln[i]
		}
	} else {
		for i, j := 0, len(refAln)-1; i < j; i, j = i+1, j-1 {
			refAln[i], refAln[j] = refAln[j], refAln[i]
		}
		for i, j := 0, len(orgAln)-1; i < j; i, j = i+1, j-1 {
			orgAln[i], orgAln[j] = orgAln[j], orgAln[i]
		}
	}

	return [2][]byte{refAln, orgAln}
}