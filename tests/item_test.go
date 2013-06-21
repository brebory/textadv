package textadv

import(
  "testing"
  "textadv"
)

var tests = []string{
  "Describer",
  "EditedDescriptionEditor",
}

func TestItem(t *testing.T) {
  // Should implement Describer and DescriptionEditor
  var(
    test, actual string
    i *textadv.Item
    d textadv.Describer
    de textadv.DescriptionEditor
    n int
  )
  defer func() {
    msg := recover()
    if msg != nil {
      t.Error("Item didn't implement" , test)
    }
  }()
  test = "Describer"
  i = textadv.CreateItem(test, "DescriptionTest", false)
  d = i

  n = 0
  actual = d.Describe()
  if actual != tests[n] {
    t.Error(
      "For", i,
      "Expected", tests[n],
      "Got", actual,
    )
  }

  test = "DescriptionEditor"
  i = textadv.CreateItem(test, "DescriptionEditorTest", false)
  de = i

  n = 1
  de.EditDescription(tests[n])
}
