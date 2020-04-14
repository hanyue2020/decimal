package decimal

func (d *Decimal) FromDB(data []byte) error {
	v, err := NewFromString(string(data))
	*d = v
	return err
}
func (d *Decimal) ToDB() ([]byte, error) {
	return []byte(d.String()), nil
}
