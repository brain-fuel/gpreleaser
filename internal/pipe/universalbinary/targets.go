package universalbinary

import "goforge.dev/gpreleaser/internal/tmpl"

type unitarget struct{}

func (unitarget) String() string { return "darwin_all" }

func (unitarget) Fields() map[string]string {
	return map[string]string{
		tmpl.KeyOS:   "darwin",
		tmpl.KeyArch: "all",
	}
}
