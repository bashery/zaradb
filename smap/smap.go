package smap

type Map struct {
	key string
	val string
}

type List struct {
	list []Map
}

func NewSmap() *List {
	return &List{
		list: []Map{},
	}
}

func (l *List) Set(k, v string) {
	found := false
	for i := range l.list {
		if l.list[i].key == k {
			l.list[i].val = v
			found = true
			break
		}
	}
	if !found {
		l.list = append(l.list, Map{k, v})
	}
}

/*
func (l *List) Set(k, v string) {
	kv := Map{k, v}
	if len(l.list) < 1 {
		l.list = append(l.list, kv)
		return
	}

	for i := 0; i < len(l.list); i++ {
		if l.list[i].key == k {
			l.list[i].val = v
			break
		} else if l.list[i].key == "" {
			l.list[i].val = v
			break
		} else {
			l.list = append(l.list, kv)
			break
		}
	}
}
*/

func (l *List) Get(k string) string {

	for i := 0; i < len(l.list); i++ {
		if l.list[i].key == k {
			return l.list[i].val
		}
	}
	return ""
}

func (l *List) Delete(k string) {
	for i := 0; i < len(l.list); i++ {
		if l.list[i].key == k {
			l.list[i].val = ""
		}
	}
}

func (l *List) Len() int {
	return len(l.list)
}
