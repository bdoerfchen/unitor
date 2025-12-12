package unitor

import (
	"cmp"
	"errors"
	"slices"
)

type unitPrefix struct {
	Prefix     string
	BaseFactor float64
	unit       *Unit
}

type prefixManager struct {
	all        map[string]*unitPrefix
	mainSorted []*unitPrefix
}

func NewPrefixes() *prefixManager {
	return &prefixManager{
		all: make(map[string]*unitPrefix),
	}
}

func (m *prefixManager) Add(prefix *unitPrefix, aliases ...string) error {
	if prefix == nil || prefix.BaseFactor == 0 {
		return errors.New("illegal input prefix")
	}

	// Check if prefix or alias is already registered
	newPrefixes := append(aliases, prefix.Prefix)
	for _, p := range newPrefixes {
		if _, used := m.all[p]; used {
			return errors.New("prefix (or one of its aliases) is already registered in the unit")
		}

		m.all[p] = prefix
	}

	// Add prefix into sorted slice
	m.mainSorted = append(m.mainSorted, prefix)
	slices.SortFunc(m.mainSorted, func(a, b *unitPrefix) int { return cmp.Compare(a.BaseFactor, b.BaseFactor) })

	return nil
}

// Adds or overwrites alias for base prefix
func (m *prefixManager) AddAlias(base string, alias string) error {
	prefix := m.For(base)
	if prefix == nil {
		return ErrPrefixNotFound
	}

	m.all[alias] = prefix

	return nil
}

func (m *prefixManager) For(prefix string) *unitPrefix {
	return m.all[prefix]
}

func (m *prefixManager) Closest(baseValue float64) *unitPrefix {
	// Find next highest index
	index, _ := slices.BinarySearchFunc(m.mainSorted, baseValue, func(prefix *unitPrefix, target float64) int {
		return cmp.Compare(prefix.BaseFactor, baseValue)
	})

	// If not smaller than smallest prefix, use one prefix before
	if index != 0 {
		index--
	}

	return m.mainSorted[index]
}

func (p *unitPrefix) String() string {
	return p.Prefix + p.unit.Symbol()
}
