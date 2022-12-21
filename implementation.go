package grm

// CreateObject ...
func (r Repository[T]) Create(data *T) error {
	return r.orm.Create(data).Error
}

// Update ...
func (r Repository[T]) Update(data *T) error {
	return r.orm.Save(data).Error
}

// Delete ...
func (r Repository[T]) Delete(query *T) error {
	return r.orm.Delete(query).Error
}

// GetOne ...
func (r Repository[T]) GetOne(query *T) (*T, error) {
	var data T
	err := r.orm.First(&data, query).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetMany ...
func (r Repository[T]) GetMany(query *T) ([]*T, error) {
	var data []*T
	err := r.orm.Find(&data, query).Error
	return data, err
}
