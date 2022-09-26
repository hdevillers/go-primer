package primer

import (
	"fmt"
	"math"

	"github.com/hdevillers/go-blast"
	"github.com/hdevillers/go-seq/seq"
)

type Match struct {
	Qseq    seq.Seq
	Qfrom   int
	Qto     int
	Sseq    seq.Seq
	Sfrom   int
	Sto     int
	Sstrand int

	Extend bool
}

func NewMatch(qs seq.Seq, e bool) *Match {
	var m Match
	m.Qseq = qs
	m.Extend = e
	m.Sstrand = 1

	return &m
}

func (m *Match) ParseHsp(ss seq.Seq, h blast.Hsp) {
	// Check the strand
	if h.HitFrame < 0 {
		m.Sstrand = -1
	}

	// Check query coverage
	if h.QueryFrom == 1 {
		// Match start at the first base of the query
		m.Qfrom = 1
		if m.Sstrand == 1 {
			m.Sfrom = h.HitFrom
		} else {
			m.Sto = h.HitTo
		}
	} else {
		// Match does not start at the first base of the query
		if m.Extend {
			// Extending match is required
			m.Qfrom = 1
			diff := h.QueryFrom - 1

			// Check if the match can be extended
			if m.Sstrand == 1 {
				// Left extend
				m.Sfrom = int(math.Max(1.0, float64(h.HitFrom-diff)))
			} else {
				// Right extend
				m.Sto = int(math.Min(float64(ss.Length()), float64(h.HitTo+diff)))
			}
		} else {
			// Do not extend
			m.Qfrom = h.QueryFrom
			if m.Sstrand == 1 {
				m.Sfrom = h.HitFrom
			} else {
				m.Sto = h.HitTo
			}
		}
	}

	if h.QueryTo == m.Qseq.Length() {
		// Match end at the last base of the query
		m.Qto = h.QueryTo
		if m.Sstrand == 1 {
			m.Sto = h.HitTo
		} else {
			m.Sfrom = h.HitFrom
		}
	} else {
		if m.Extend {
			m.Qto = m.Qseq.Length()
			diff := m.Qto - h.QueryTo

			// Check if match can be extended
			if m.Sstrand == 1 {
				// Right extend
				m.Sto = int(math.Min(float64(ss.Length()), float64(h.HitTo+diff)))
			} else {
				// Left extend
				m.Sfrom = int(math.Max(1.0, float64(h.HitFrom-diff)))
			}
		} else {
			// Do not extend
			m.Qto = h.QueryTo
			if m.Sstrand == 1 {
				m.Sto = h.HitTo
			} else {
				m.Sfrom = h.HitFrom
			}
		}
	}

	// Create the subject sequence
	m.Sseq = seq.Seq{
		Id:       fmt.Sprintf("%s_%d-%d", ss.Id, m.Sfrom, m.Sto),
		Sequence: ss.Sequence[(m.Sfrom - 1):m.Sto],
	}

}
