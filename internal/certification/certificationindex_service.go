package certification

type CertificationIndexIndex struct{ buckets map[string][]string }

func NewCertificationIndexIndex() *CertificationIndexIndex { return &CertificationIndexIndex{} }
func (i *CertificationIndexIndex) Restore(bucket, value string) {
	i.buckets[bucket] = append(i.buckets[bucket], value)
}
