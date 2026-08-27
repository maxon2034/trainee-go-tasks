package duplicate

func Remove(s []string) []string {
	m := make(map[string]struct{})
	var res []string
	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		res = append(res, v)
	}
	return res
}
