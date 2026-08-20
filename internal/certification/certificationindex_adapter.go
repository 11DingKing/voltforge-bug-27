package certification

func (i *CertificationIndexIndex) Load(bucket string) []string {
	if i == nil || i.buckets == nil {
		return nil
	}
	values := i.buckets[bucket]
	out := make([]string, len(values))
	copy(out, values)
	return out
}
func (i *CertificationIndexIndex) Ready() bool { return i != nil && i.buckets != nil }
