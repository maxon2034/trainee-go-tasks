package repeat

func Remove(s []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range s {
		if m[v] == true {
			continue
		}
		m[v] = true
		res = append(res, v)
	}
	return res
}
