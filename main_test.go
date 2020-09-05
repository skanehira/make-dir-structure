package main

import (
	"encoding/json"
	"testing"

	"github.com/Cside/jsondiff"
)

func Test_makeItems(t *testing.T) {
	testData := `
[
  {"path":"a/b/d/c.json"},
  {"path":"a/b/d/b.json"},
  {"path":"a/b/d.json"},
  {"path":"b/c/f.json"},
  {"path":"b/c/b.json"},
  {"path":"b/f/g.json"}
]
`
	items := makeItems(testData)

	got, err := json.Marshal(items)
	if err != nil {
		t.Fatalf("marshal error: %s", err)
	}

	want := `[
  {
    "name": "a",
    "children": [
      {
        "name": "b",
        "children": [
          {
            "name": "d",
            "children": [
              {
                "name": "c.json"
              },
              {
                "name": "b.json"
              }
            ]
          },
          {
            "name": "d.json"
          }
        ]
      }
    ]
  },
  {
    "name": "b",
    "children": [
      {
        "name": "c",
        "children": [
          {
            "name": "f.json"
          },
          {
            "name": "b.json"
          }
        ]
      },
      {
        "name": "f",
        "children": [
          {
            "name": "g.json"
          }
        ]
      }
    ]
  }
]`

	if diff := jsondiff.Diff([]byte(want), got); diff != "" {
		t.Fatal(diff)
	}
}
