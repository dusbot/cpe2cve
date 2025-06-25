package core

import (
	"fmt"
)

const (
	Disjoint Relation = iota
	Subset
	Equal
	Superset
)

func HasWildcard(s string) bool {
	for n, r := range s {
		if r != '*' && r != '?' {
			continue
		}
		quoted := false
		for i := n - 1; i >= 0; i-- {
			if s[i] != '\\' {
				break
			}
			quoted = !quoted
		}
		if !quoted {
			return true
		}
	}
	return false
}

type Relation int

func (r Relation) String() string {
	switch r {
	case Disjoint:
		return "DISJOINT"
	case Subset:
		return "SUBSET"
	case Equal:
		return "EQUAL"
	case Superset:
		return "SUPERSET"
	default:
		return fmt.Sprintf("Undefined value %d", r)
	}
}

type Comparison struct {
	Part      Relation
	Vendor    Relation
	Product   Relation
	Version   Relation
	Update    Relation
	Edition   Relation
	Language  Relation
	SWEdition Relation
	TargetSW  Relation
	TargetHW  Relation
	Other     Relation
}

func (c Comparison) IsDisjoint() bool {
	switch {
	case c.Part == Disjoint:
		return true
	case c.Vendor == Disjoint:
		return true
	case c.Product == Disjoint:
		return true
	case c.Version == Disjoint:
		return true
	case c.Update == Disjoint:
		return true
	case c.Edition == Disjoint:
		return true
	case c.Language == Disjoint:
		return true
	case c.SWEdition == Disjoint:
		return true
	case c.TargetSW == Disjoint:
		return true
	case c.TargetHW == Disjoint:
		return true
	case c.Other == Disjoint:
		return true
	default:
		return false
	}
}

func (c Comparison) IsEqual() bool {
	switch {
	case c.Part != Equal:
		return false
	case c.Vendor != Equal:
		return false
	case c.Product != Equal:
		return false
	case c.Version != Equal:
		return false
	case c.Update != Equal:
		return false
	case c.Edition != Equal:
		return false
	case c.Language != Equal:
		return false
	case c.SWEdition != Equal:
		return false
	case c.TargetSW != Equal:
		return false
	case c.TargetHW != Equal:
		return false
	case c.Other != Equal:
		return false
	default:
		return true
	}
}

func (c Comparison) IsSubset() bool {
	switch {
	case c.Part != Equal && c.Part != Subset:
		return false
	case c.Vendor != Equal && c.Vendor != Subset:
		return false
	case c.Product != Equal && c.Product != Subset:
		return false
	case c.Version != Equal && c.Version != Subset:
		return false
	case c.Update != Equal && c.Update != Subset:
		return false
	case c.Edition != Equal && c.Edition != Subset:
		return false
	case c.Language != Equal && c.Language != Subset:
		return false
	case c.SWEdition != Equal && c.SWEdition != Subset:
		return false
	case c.TargetSW != Equal && c.TargetSW != Subset:
		return false
	case c.TargetHW != Equal && c.TargetHW != Subset:
		return false
	case c.Other != Equal && c.Other != Subset:
		return false
	default:
		return true
	}
}

func (c Comparison) IsSuperset() bool {
	switch {
	case c.Part != Equal && c.Part != Superset:
		return false
	case c.Vendor != Equal && c.Vendor != Superset:
		return false
	case c.Product != Equal && c.Product != Superset:
		return false
	case c.Version != Equal && c.Version != Superset:
		return false
	case c.Update != Equal && c.Update != Superset:
		return false
	case c.Edition != Equal && c.Edition != Superset:
		return false
	case c.Language != Equal && c.Language != Superset:
		return false
	case c.SWEdition != Equal && c.SWEdition != Superset:
		return false
	case c.TargetSW != Equal && c.TargetSW != Superset:
		return false
	case c.TargetHW != Equal && c.TargetHW != Superset:
		return false
	case c.Other != Equal && c.Other != Superset:
		return false
	default:
		return true
	}
}

func (c Comparison) Relation() Relation {
	if c.IsSubset() {
		return Subset
	}
	if c.IsEqual() {
		return Equal
	}
	if c.IsSuperset() {
		return Superset
	}
	return Disjoint
}

func Compare(src, tgt *Attributes) (Comparison, error) {
	var result Comparison
	var err error
	if result.Part, err = CompareAttr(src.Part, tgt.Part); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Part, tgt.Part, err)
	}
	if result.Vendor, err = CompareAttr(src.Vendor, tgt.Vendor); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Vendor, tgt.Vendor, err)
	}
	if result.Product, err = CompareAttr(src.Product, tgt.Product); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Product, tgt.Product, err)
	}
	if result.Version, err = CompareAttr(src.Version, tgt.Version); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Version, tgt.Version, err)
	}
	if result.Update, err = CompareAttr(src.Update, tgt.Update); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Update, tgt.Update, err)
	}
	if result.Edition, err = CompareAttr(src.Edition, tgt.Edition); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Edition, tgt.Edition, err)
	}
	if result.Language, err = CompareAttr(src.Language, tgt.Language); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Language, tgt.Language, err)
	}
	if result.SWEdition, err = CompareAttr(src.SWEdition, tgt.SWEdition); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.SWEdition, tgt.SWEdition, err)
	}
	if result.TargetSW, err = CompareAttr(src.TargetSW, tgt.TargetSW); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.TargetSW, tgt.TargetSW, err)
	}
	if result.TargetHW, err = CompareAttr(src.TargetHW, tgt.TargetHW); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.TargetHW, tgt.TargetHW, err)
	}
	if result.Other, err = CompareAttr(src.Other, tgt.Other); err != nil {
		return result, fmt.Errorf("failed to compare wfns %q to %q: %v", src.Other, tgt.Other, err)
	}
	return result, nil
}

func Match(src, tgt *Attributes) bool {
	if src == nil || tgt == nil {
		return false
	}
	return matchAttr(src.Part, tgt.Part) && matchAttr(src.Vendor, tgt.Vendor) &&
		matchAttr(src.Product, tgt.Product) && matchAttr(src.Version, tgt.Version) &&
		matchAttr(src.Update, tgt.Update) && matchAttr(src.Edition, tgt.Edition) &&
		matchAttr(src.Language, tgt.Language) && matchAttr(src.SWEdition, tgt.SWEdition) &&
		matchAttr(src.TargetHW, tgt.TargetHW) && matchAttr(src.TargetSW, tgt.TargetSW) &&
		matchAttr(src.Other, tgt.Other)
}

func CompareAttr(src, tgt string) (Relation, error) {
	if src != NA && src != Any && HasWildcard(tgt) {
		return Disjoint, fmt.Errorf("target attribute value cannot contain wildcard")
	}
	if src == tgt {
		return Equal, nil
	}
	if src == Any {
		return Superset, nil
	}
	if tgt == Any {
		return Subset, nil
	}
	if src == NA || tgt == NA {
		return Disjoint, nil
	}
	return matchStr(src, tgt), nil
}

func matchAttr(src, tgt string) bool {
	switch {
	case src == Any || tgt == Any || src == tgt:
		return true
	case src == NA || tgt == NA || HasWildcard(tgt):
		return false
	default:
		return matchStr(src, tgt) != Disjoint
	}
}

func matchStr(s, t string) Relation {
	escaped := false
	matchesAs := Equal
	idx := 0
	for ; idx < len(t); idx++ {
		if idx >= len(s) {
			return Disjoint
		}
		if !escaped && s[idx] == '*' {
			if idx == len(s)-1 {
				return Superset
			}
			for i := idx; i < len(t); i++ {
				if matchStr(s[idx+1:], t[i:]) != Disjoint {
					return Superset
				}
			}
			return Disjoint
		}

		if (escaped || s[idx] != '?') && s[idx] != t[idx] {
			return Disjoint
		} else if !escaped && s[idx] == '?' {
			matchesAs = Superset
		}
		if s[idx] == '\\' {
			escaped = !escaped
		} else {
			escaped = false
		}
	}
	for ; idx < len(s); idx++ {
		if s[idx] != '*' {
			return Disjoint
		}
	}
	if len(s) > len(t) {
		return Superset
	}
	return matchesAs
}
