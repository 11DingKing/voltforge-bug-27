package certification

type CertificationIndexIndex struct{ buckets map[string][]string }

func NewCertificationIndexIndex() *CertificationIndexIndex {
	return &CertificationIndexIndex{buckets: make(map[string][]string)}
}
func (i *CertificationIndexIndex) Restore(bucket, value string) {
	if i.buckets == nil {
		i.buckets = make(map[string][]string)
	}
	i.buckets[bucket] = append(i.buckets[bucket], value)
}
