package Tool

type Comparable interface{
	CompareTo(a Comparable) bool
}