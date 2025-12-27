package v1

import (
	"fmt"

	"github.com/charmbracelet/huh"
	fstreev1 "github.com/rcdmrl/go-sandbox/fstree/v1"
	fstreev2 "github.com/rcdmrl/go-sandbox/fstree/v2"
	todoappv1 "github.com/rcdmrl/go-sandbox/todoapp/v1"
)

const (
	ProjFSTree   = "fstree"
	ProjTodoApp  = "todo"
	ProjSayonara = "sayonara"

	V1 = "v1"
	V2 = "v2"
)

type MainForm struct {
	projectName    string
	projectVersion string
	// deps
	tree1    *fstreev1.ParallelDir
	tree2    *fstreev2.ParallelDir
	todoapp1 *todoappv1.TodoApp
}

func NewMainForm(tree1 *fstreev1.ParallelDir, tree2 *fstreev2.ParallelDir, todoapp1 *todoappv1.TodoApp) *MainForm {
	return &MainForm{"", "", tree1, tree2, todoapp1}
}

// Run executes the multi-step flow: pick project, then (if needed) pick a version.
func (f *MainForm) Run() error {
	if err := f.runProjectSelect(); err != nil {
		return err
	}

	// In case projects have more screens, like one to choose versions
	switch f.projectName {
	case ProjFSTree:
		return f.runProjectVersionSelect(ProjFSTree, V1, V2)
	case ProjTodoApp:
		return f.runProjectVersionSelect(ProjTodoApp, V1)
	default:
		return nil
	}
}

// Dispatch runs the selected project/version after the user has gone through the TUI options
func (f *MainForm) Dispatch() error {
	switch f.projectName {
	case ProjFSTree:
		switch f.projectVersion {
		case V1:
			f.tree1.Run()
		case V2:
			f.tree2.Run()
		default:
			return fmt.Errorf("unknown fs tree version %q", f.projectVersion)
		}
	case ProjTodoApp:
		switch f.projectVersion {
		case V1:
			f.todoapp1.Run()
		default:
			return fmt.Errorf("unknown todo app version %q", f.projectVersion)
		}
	case ProjSayonara:
		fmt.Println("You called quits. Cya!")
	default:
		return fmt.Errorf("the Dispatch method don't know about the %q project.", f.projectName)
	}
	return nil
}

// runProjectSelect shows the top-level project chooser
func (f *MainForm) runProjectSelect() error {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which project?").
				Options(
					huh.NewOption("FS Tree", ProjFSTree),
					huh.NewOption("TODO App", ProjTodoApp),
					huh.NewOption("Sayonara", ProjSayonara),
				).
				Value(&f.projectName),
		),
	).Run()
}

// runProjectVersionSelect shows the available versions for a given project.
func (f *MainForm) runProjectVersionSelect(proj string, versions ...string) error {
	// Transforming the string versions into an array of huh.Option
	versionOps := make([]huh.Option[string], 0, len(versions))
	for _, opt := range versions {
		versionOps = append(versionOps, huh.NewOption(opt, opt))
	}
	// The form
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which " + proj + " project version?").
				Options(versionOps...).
				Value(&f.projectVersion),
		),
	).Run()
}
