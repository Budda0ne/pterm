package pterm_test

// Behavioral tests for TreePrinter: exact connector layout (corner vs. tee
// connectors, vertical continuation lines), indentation math, custom connector
// strings and the LeveledList conversion. The builder/contract plumbing is
// covered in contract_test.go, one representative output in snapshot_test.go.

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func TestTreePrinter_ConnectorLayout(t *testing.T) {
	printer := pterm.DefaultTree.WithRoot(pterm.TreeNode{
		Text: "root",
		Children: []pterm.TreeNode{
			{Text: "a"},
			{Text: "b", Children: []pterm.TreeNode{
				{Text: "b1"},
				{Text: "b2"},
			}},
			{Text: "c", Children: []pterm.TreeNode{
				{Text: "c1", Children: []pterm.TreeNode{
					{Text: "c1a"},
				}},
			}},
		},
	})

	// - non-last leaf:            ├──
	// - non-last with children:   ├─┬ and a │ continuation for its subtree
	// - last child:               └── / └─┬ with plain-space continuation
	expected := "" +
		"root\n" +
		"├──a\n" +
		"├─┬b\n" +
		"│ ├──b1\n" +
		"│ └──b2\n" +
		"└─┬c\n" +
		"  └─┬c1\n" +
		"    └──c1a\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTreePrinter_SingleRootWithoutChildren(t *testing.T) {
	printer := pterm.DefaultTree.WithRoot(pterm.TreeNode{Text: "root"})

	assert.Equal(t, "root\n", srenderPlain(t, printer))
}

func TestTreePrinter_EmptyRootTextOmitsRootLine(t *testing.T) {
	printer := pterm.DefaultTree.WithRoot(pterm.TreeNode{
		Children: []pterm.TreeNode{{Text: "only child"}},
	})

	assert.Equal(t, "└──only child\n", srenderPlain(t, printer))
}

func TestTreePrinter_IndentControlsConnectorAndPrefixWidth(t *testing.T) {
	printer := pterm.DefaultTree.WithIndent(4).WithRoot(pterm.TreeNode{
		Children: []pterm.TreeNode{
			{Text: "p", Children: []pterm.TreeNode{
				{Text: "q"},
			}},
			{Text: "r"},
		},
	})

	// Indent 4: three horizontals before the ┬ of a parent, four before leaf
	// text, and continuation prefixes that are exactly four cells wide.
	expected := "" +
		"├───┬p\n" +
		"│   └────q\n" +
		"└────r\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTreePrinter_CustomConnectorStrings(t *testing.T) {
	printer := pterm.DefaultTree.
		WithTopRightDownStringOngoing("Y").
		WithTopRightCornerString("X").
		WithHorizontalString("_").
		WithRoot(pterm.TreeNode{
			Children: []pterm.TreeNode{
				{Text: "a"},
				{Text: "b"},
			},
		})

	expected := "" +
		"Y__a\n" +
		"X__b\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTreePrinter_TextStyleAppliedToNodeText(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	printer := pterm.DefaultTree.WithTextStyle(style).WithRoot(pterm.TreeNode{
		Text:     "root",
		Children: []pterm.TreeNode{{Text: "leaf"}},
	})

	styled, err := printer.Srender()
	require.NoError(t, err)

	assert.Contains(t, styled, style.Sprint("root"))
	assert.Contains(t, styled, style.Sprint("leaf"))
}

func TestTreePrinter_NewTreeFromLeveledList(t *testing.T) {
	// Note: the (deprecated) converter sets the root's Text to the first
	// item's text and gives leaves empty, non-nil Children slices.
	leaf := func(text string) pterm.TreeNode {
		return pterm.TreeNode{Text: text, Children: []pterm.TreeNode{}}
	}

	t.Run("builds nested nodes from levels", func(t *testing.T) {
		root := pterm.NewTreeFromLeveledList(pterm.LeveledList{
			{Level: 0, Text: "a"},
			{Level: 1, Text: "a1"},
			{Level: 1, Text: "a2"},
			{Level: 0, Text: "b"},
		})

		assert.Equal(t, pterm.TreeNode{
			Text: "a",
			Children: []pterm.TreeNode{
				{Text: "a", Children: []pterm.TreeNode{
					leaf("a1"),
					leaf("a2"),
				}},
				leaf("b"),
			},
		}, root)
	})

	t.Run("clamps level jumps to one below the previous item", func(t *testing.T) {
		root := pterm.NewTreeFromLeveledList(pterm.LeveledList{
			{Level: 0, Text: "a"},
			{Level: 10, Text: "b"}, // invalid jump, must become level 1
		})

		assert.Equal(t, pterm.TreeNode{
			Text: "a",
			Children: []pterm.TreeNode{
				{Text: "a", Children: []pterm.TreeNode{
					leaf("b"),
				}},
			},
		}, root)
	})

	t.Run("clamps negative levels to zero", func(t *testing.T) {
		root := pterm.NewTreeFromLeveledList(pterm.LeveledList{
			{Level: 0, Text: "a"},
			{Level: -5, Text: "b"},
		})

		assert.Equal(t, pterm.TreeNode{
			Text: "a",
			Children: []pterm.TreeNode{
				leaf("a"),
				leaf("b"),
			},
		}, root)
	})

	t.Run("empty list yields an empty root", func(t *testing.T) {
		assert.Zero(t, pterm.NewTreeFromLeveledList(pterm.LeveledList{}))
	})
}

func TestTreePrinter_SrenderIsPure(t *testing.T) {
	root := pterm.TreeNode{
		Text: "root",
		Children: []pterm.TreeNode{
			{Text: "a", Children: []pterm.TreeNode{{Text: "a1"}}},
		},
	}
	printer := pterm.DefaultTree.WithRoot(root)

	first, err := printer.Srender()
	require.NoError(t, err)

	second, err := printer.Srender()
	require.NoError(t, err)

	assert.Equal(t, first, second, "rendering twice must yield identical output")
	assert.Equal(t, pterm.TreeNode{
		Text: "root",
		Children: []pterm.TreeNode{
			{Text: "a", Children: []pterm.TreeNode{{Text: "a1"}}},
		},
	}, root, "rendering must not modify the input tree")
}
