package fakers

// import (
//	"rest-project/src/adapters"
//	"sync"
//)
//
// type FakeMemoryRepository[T adapters.Model] struct {
//	data map[string]T
//	mu   *sync.Mutex
//}

// func (f *FakeMemoryRepository[T]) Create(entity *T, ctx context.Context) error {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	ent := *entity
// 	f.data[ent.GetID().String()] = *entity
// 	return nil
// }
//
// func (f *FakeMemoryRepository[T]) BulkCreate(entity *[]T, ctx context.Context) error {
//	f.mu.Lock()
//	defer f.mu.Unlock()
//	for i, i2 := range *entity {
//		f.data[i2.GetID().String()] = (*entity)[i]
//	}
//	return nil
// }
//
// func (f *FakeMemoryRepository[T]) GetById(id uuid.UUID, ctx context.Context) (*T, error) {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	entity, ok := f.data[id.String()]
// 	if !ok {
// 		notFound := adapters.NotFoundError{Name: "Article"}
// 		return nil, notFound.Error()
// 	}
// 	return &entity, nil
// }
//
// func (f *FakeMemoryRepository[T]) Get(params *T, ctx context.Context) (*T, error) {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	for _, entity := range f.data {
// 		ent := entity
// 		param := *params
// 		if ent.GetID().String() == param.GetID().String() {
// 			return &entity, nil
// 		}
// 	}
// 	notFound := adapters.NotFoundError{Name: "Article"}
// 	return nil, notFound.Error()
// }
//
// func (f *FakeMemoryRepository[T]) GetAll(ctx context.Context) (*[]T, error) {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	if f.data != nil {
// 		var entities []T
// 		for _, entity := range f.data {
// 			entities = append(entities, entity)
// 		}
// 		return &entities, nil
// 	}
// 	return &[]T{}, nil
// }
//
// func (f *FakeMemoryRepository[T]) Update(entity *T, ctx context.Context) error {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	ent := *entity
// 	_, ok := f.data[ent.GetID().String()]
// 	if !ok {
// 		t := reflect.TypeOf(entity)
// 		notFound := adapters.NotFoundError{Name: t.Name()}
// 		return notFound.Error()
// 	}
// 	f.data[ent.GetID().String()] = ent
// 	return nil
// }
//
// func (f *FakeMemoryRepository[T]) UpdateAll(entity *[]T, ctx context.Context) error {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	for _, ent := range *entity {
// 		_, ok := f.data[ent.GetID().String()]
// 		if !ok {
// 			t := reflect.TypeOf(entity)
// 			notFound := adapters.NotFoundError{Name: t.Name()}
// 			return notFound.Error()
// 		}
// 		f.data[ent.GetID().String()] = ent
// 	}
// 	return nil
// }
//
// func (f *FakeMemoryRepository[T]) DeleteById(id uuid.UUID, ctx context.Context) error {
// 	f.mu.Lock()
// 	defer f.mu.Unlock()
// 	_, ok := f.data[id.String()]
// 	if !ok {
// 		t := reflect.TypeOf("TestEntity")
// 		notFound := adapters.NotFoundError{Name: t.Name()}
// 		return notFound.Error()
// 	}
// 	delete(f.data, id.String())
// 	return nil
// }
//
// func NewFakeMemoryRepository[T adapters.Model]() *FakeMemoryRepository[T] {
// 	return &FakeMemoryRepository[T]{
// 		data: make(map[string]T),
// 		mu:   &sync.Mutex{},
// 	}
// }
