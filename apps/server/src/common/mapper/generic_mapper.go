package mapper

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
)

// MapperError represents mapping-related errors
type MapperError struct {
	Operation string
	Err       error
}

func (e *MapperError) Error() string {
	return fmt.Sprintf("mapper %s error: %v", e.Operation, e.Err)
}

// MappingOptions provides configuration for mapping operations
type MappingOptions struct {
	IgnoreEmpty  bool
	DeepCopy     bool
	IgnoreFields []string
}

// DefaultMappingOptions returns sensible defaults for mapping
func DefaultMappingOptions() MappingOptions {
	return MappingOptions{
		IgnoreEmpty:  true,
		DeepCopy:     true,
		IgnoreFields: []string{},
	}
}

// UpdateMappingOptions returns options optimized for updates
func UpdateMappingOptions() MappingOptions {
	return MappingOptions{
		IgnoreEmpty:  true,
		DeepCopy:     false,
		IgnoreFields: []string{"ID", "CreatedAt"},
	}
}

// Compare 2 objects if they are same type
func IsSameType[T1, T2 any](a T1, b T2) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

// IsNil checks if an interface is nil or points to a nil value
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return val.IsNil()
	default:
		return false
	}
}

// Map from source to destination using copier library with default options
func DefaultMap(toValue interface{}, fromValue interface{}) error {
	if IsNil(fromValue) {
		return &MapperError{
			Operation: "default_map",
			Err:       fmt.Errorf("source value is nil"),
		}
	}

	if IsNil(toValue) {
		return &MapperError{
			Operation: "default_map",
			Err:       fmt.Errorf("destination value is nil"),
		}
	}

	err := copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})

	if err != nil {
		return &MapperError{
			Operation: "default_map",
			Err:       err,
		}
	}

	return nil
}

// MapWithOptions maps with custom options
func MapWithOptions(toValue interface{}, fromValue interface{}, opts MappingOptions) error {
	if IsNil(fromValue) {
		return &MapperError{
			Operation: "map_with_options",
			Err:       fmt.Errorf("source value is nil"),
		}
	}

	if IsNil(toValue) {
		return &MapperError{
			Operation: "map_with_options",
			Err:       fmt.Errorf("destination value is nil"),
		}
	}

	option := copier.Option{
		IgnoreEmpty: opts.IgnoreEmpty,
		DeepCopy:    opts.DeepCopy,
	}

	err := copier.CopyWithOption(toValue, fromValue, option)
	if err != nil {
		return &MapperError{
			Operation: "map_with_options",
			Err:       err,
		}
	}

	return nil
}

// MapForUpdate maps for update operations (ignoring certain fields)
func MapForUpdate(toValue interface{}, fromValue interface{}) error {
	return MapWithOptions(toValue, fromValue, UpdateMappingOptions())
}

// GenericMapper defines the interface for a generic mapping service with separate Request and Response DTOs.
type GenericMapper[E any, Req any, Res any] interface {
	// Response DTO operations
	ToResponseDTO(entity *E) (*Res, error)
	ToResponseDTOList(entities []E) ([]Res, error)
	ToResponseDTOSafe(entity *E) *Res // Returns nil on error instead of error

	// Entity operations
	ToEntity(requestDTO Req) (*E, error)
	ToEntityList(requestDTOs []Req) ([]*E, error)
	ToEntitySafe(requestDTO Req) *E // Returns nil on error instead of error

	// Update operations
	UpdateEntity(entity *E, requestDTO Req) error
	UpdateEntityPartial(entity *E, requestDTO Req, ignoreFields []string) error

	// Validation
	ValidateMapping(requestDTO Req) error

	// Batch operations
	BatchToResponseDTO(entities []*E) ([]Res, error)
	BatchToEntity(requestDTOs []Req) ([]*E, error)
}

type genericMapper[E any, Req any, Res any] struct {
	options MappingOptions
}

// NewGenericMapper creates a new instance of GenericMapper with default options.
func NewGenericMapper[E any, Req any, Res any]() GenericMapper[E, Req, Res] {
	return &genericMapper[E, Req, Res]{
		options: DefaultMappingOptions(),
	}
}

// NewGenericMapperWithOptions creates a new instance of GenericMapper with custom options.
func NewGenericMapperWithOptions[E any, Req any, Res any](opts MappingOptions) GenericMapper[E, Req, Res] {
	return &genericMapper[E, Req, Res]{
		options: opts,
	}
}

// ToResponseDTO maps an entity to a Response DTO.
func (m *genericMapper[E, Req, Res]) ToResponseDTO(entity *E) (*Res, error) {
	if entity == nil {
		return nil, nil
	}

	var dto Res
	if err := MapWithOptions(&dto, entity, m.options); err != nil {
		return nil, &MapperError{
			Operation: "to_response_dto",
			Err:       err,
		}
	}

	return &dto, nil
}

// ToResponseDTOSafe maps an entity to a Response DTO, returning nil on error
func (m *genericMapper[E, Req, Res]) ToResponseDTOSafe(entity *E) *Res {
	dto, err := m.ToResponseDTO(entity)
	if err != nil {
		return nil
	}
	return dto
}

// ToResponseDTOList maps a list of entities to a list of Response DTOs.
func (m *genericMapper[E, Req, Res]) ToResponseDTOList(entities []E) ([]Res, error) {
	if entities == nil {
		return nil, nil
	}

	dtos := make([]Res, 0, len(entities))
	for i, entity := range entities {
		dto, err := m.ToResponseDTO(&entity)
		if err != nil {
			return nil, &MapperError{
				Operation: fmt.Sprintf("to_response_dto_list[%d]", i),
				Err:       err,
			}
		}
		if dto != nil {
			dtos = append(dtos, *dto)
		}
	}

	return dtos, nil
}

// BatchToResponseDTO maps a list of entity pointers to Response DTOs
func (m *genericMapper[E, Req, Res]) BatchToResponseDTO(entities []*E) ([]Res, error) {
	if entities == nil {
		return nil, nil
	}

	dtos := make([]Res, 0, len(entities))
	for i, entity := range entities {
		if entity == nil {
			continue
		}
		dto, err := m.ToResponseDTO(entity)
		if err != nil {
			return nil, &MapperError{
				Operation: fmt.Sprintf("batch_to_response_dto[%d]", i),
				Err:       err,
			}
		}
		if dto != nil {
			dtos = append(dtos, *dto)
		}
	}

	return dtos, nil
}

// ToEntity maps a Request DTO to an entity.
func (m *genericMapper[E, Req, Res]) ToEntity(requestDTO Req) (*E, error) {
	var entity E
	if err := MapWithOptions(&entity, requestDTO, m.options); err != nil {
		return nil, &MapperError{
			Operation: "to_entity",
			Err:       err,
		}
	}

	return &entity, nil
}

// ToEntitySafe maps a Request DTO to an entity, returning nil on error
func (m *genericMapper[E, Req, Res]) ToEntitySafe(requestDTO Req) *E {
	entity, err := m.ToEntity(requestDTO)
	if err != nil {
		return nil
	}
	return entity
}

// ToEntityList maps a list of Request DTOs to entities
func (m *genericMapper[E, Req, Res]) ToEntityList(requestDTOs []Req) ([]*E, error) {
	if requestDTOs == nil {
		return nil, nil
	}

	entities := make([]*E, 0, len(requestDTOs))
	for i, dto := range requestDTOs {
		entity, err := m.ToEntity(dto)
		if err != nil {
			return nil, &MapperError{
				Operation: fmt.Sprintf("to_entity_list[%d]", i),
				Err:       err,
			}
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// BatchToEntity is an alias for ToEntityList for consistency
func (m *genericMapper[E, Req, Res]) BatchToEntity(requestDTOs []Req) ([]*E, error) {
	return m.ToEntityList(requestDTOs)
}

// UpdateEntity updates an existing entity with data from Request DTO
func (m *genericMapper[E, Req, Res]) UpdateEntity(entity *E, requestDTO Req) error {
	if entity == nil {
		return &MapperError{
			Operation: "update_entity",
			Err:       fmt.Errorf("entity is nil"),
		}
	}

	if err := MapForUpdate(entity, requestDTO); err != nil {
		return &MapperError{
			Operation: "update_entity",
			Err:       err,
		}
	}

	return nil
}

// UpdateEntityPartial updates an entity while ignoring specific fields
func (m *genericMapper[E, Req, Res]) UpdateEntityPartial(entity *E, requestDTO Req, ignoreFields []string) error {
	if entity == nil {
		return &MapperError{
			Operation: "update_entity_partial",
			Err:       fmt.Errorf("entity is nil"),
		}
	}

	opts := m.options
	opts.IgnoreFields = ignoreFields

	if err := MapWithOptions(entity, requestDTO, opts); err != nil {
		return &MapperError{
			Operation: "update_entity_partial",
			Err:       err,
		}
	}

	return nil
}

// ValidateMapping validates if a Request DTO can be mapped (basic validation)
func (m *genericMapper[E, Req, Res]) ValidateMapping(requestDTO Req) error {
	// Basic validation - check if the DTO is not nil (for pointer types)
	val := reflect.ValueOf(requestDTO)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return &MapperError{
			Operation: "validate_mapping",
			Err:       fmt.Errorf("request DTO is nil"),
		}
	}

	// Additional custom validation can be added here
	// For now, we'll just try to create a temporary entity
	_, err := m.ToEntity(requestDTO)
	if err != nil {
		return &MapperError{
			Operation: "validate_mapping",
			Err:       fmt.Errorf("mapping validation failed: %w", err),
		}
	}

	return nil
}

// Helper functions for common mapping patterns

// MapSlice maps a slice of one type to another using a transform function
func MapSlice[T, U any](slice []T, transform func(T) (U, error)) ([]U, error) {
	if slice == nil {
		return nil, nil
	}

	result := make([]U, 0, len(slice))
	for i, item := range slice {
		transformed, err := transform(item)
		if err != nil {
			return nil, &MapperError{
				Operation: fmt.Sprintf("map_slice[%d]", i),
				Err:       err,
			}
		}
		result = append(result, transformed)
	}

	return result, nil
}

// MapSliceSafe maps a slice, skipping items that fail to transform
func MapSliceSafe[T, U any](slice []T, transform func(T) (U, error)) []U {
	if slice == nil {
		return nil
	}

	result := make([]U, 0, len(slice))
	for _, item := range slice {
		if transformed, err := transform(item); err == nil {
			result = append(result, transformed)
		}
	}

	return result
}

// Clone creates a deep copy of an object
func Clone[T any](source T) (T, error) {
	var destination T
	if err := DefaultMap(&destination, source); err != nil {
		return destination, &MapperError{
			Operation: "clone",
			Err:       err,
		}
	}
	return destination, nil
}

// CloneSafe creates a deep copy, returning zero value on error
func CloneSafe[T any](source T) T {
	cloned, err := Clone(source)
	if err != nil {
		var zero T
		return zero
	}
	return cloned
}
