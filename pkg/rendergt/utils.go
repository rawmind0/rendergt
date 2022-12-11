package rendergt

func mergeMaps(a, b map[string]interface{}) {
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := a[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					mergeMaps(bv, v)
					continue
				}
			}
		}
		a[k] = v
	}
}
