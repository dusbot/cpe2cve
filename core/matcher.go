package core

type Matcher interface {
	Match(attrs []*Attributes, requireVersion bool) (matches []*Attributes)
	Config() []*Attributes
}

func (a *Attributes) Config() []*Attributes {
	return []*Attributes{a}
}

func (a *Attributes) MatchOnlyVersion(attr *Attributes) bool {
	if a == nil || attr == nil {
		return a == attr
	}
	return matchAttr(a.Version, attr.Version)
}

func (a *Attributes) MatchWithoutVersion(attr *Attributes) bool {
	if a == nil || attr == nil {
		return a == attr
	}
	return matchAttr(a.Product, attr.Product) &&
		matchAttr(a.Vendor, attr.Vendor) && matchAttr(a.Part, attr.Part) &&
		matchAttr(a.Update, attr.Update) && matchAttr(a.Edition, attr.Edition) &&
		matchAttr(a.Language, attr.Language) && matchAttr(a.SWEdition, attr.SWEdition) &&
		matchAttr(a.TargetHW, attr.TargetHW) && matchAttr(a.TargetSW, attr.TargetSW) &&
		matchAttr(a.Other, attr.Other)
}

func MatchAll(ms ...Matcher) Matcher {
	return &multiMatcher{ms, true}
}

func MatchAny(ms ...Matcher) Matcher {
	return &multiMatcher{ms, false}
}

func DontMatch(m Matcher) Matcher {
	return notMatcher{m}
}

type multiMatcher struct {
	matchers []Matcher
	allMatch bool
}

func (mm *multiMatcher) Match(attrs []*Attributes, requireVersion bool) []*Attributes {
	matched := make(map[*Attributes]bool)
	for _, matcher := range mm.matchers {
		matches := matcher.Match(attrs, requireVersion)
		if mm.allMatch && len(matches) == 0 {

			return nil
		}
		for _, m := range matches {
			matched[m] = true
		}
	}

	matches := make([]*Attributes, 0, len(matched))
	for m := range matched {
		matches = append(matches, m)
	}
	return matches
}

func (mm *multiMatcher) Config() []*Attributes {
	var attrs []*Attributes
	for _, matcher := range mm.matchers {
		attrs = append(attrs, matcher.Config()...)
	}
	return attrs
}

type notMatcher struct {
	Matcher
}

func (nm notMatcher) Match(attrs []*Attributes, requireVersion bool) (matches []*Attributes) {
	matched := make(map[*Attributes]bool)
	for _, m := range nm.Matcher.Match(attrs, requireVersion) {
		matched[m] = true
	}

	for _, a := range attrs {
		if !matched[a] {
			matches = append(matches, a)
		}
	}
	return matches
}
