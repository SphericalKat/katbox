// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/SphericalKat/katbox/ent/file"
	"github.com/SphericalKat/katbox/ent/predicate"
	"github.com/SphericalKat/katbox/ent/user"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFile = "File"
	TypeUser = "User"
)

// FileMutation represents an operation that mutates the File nodes in the graph.
type FileMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	storage_key   *string
	expires_at    *time.Time
	clearedFields map[string]struct{}
	user          *uuid.UUID
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*File, error)
	predicates    []predicate.File
}

var _ ent.Mutation = (*FileMutation)(nil)

// fileOption allows management of the mutation configuration using functional options.
type fileOption func(*FileMutation)

// newFileMutation creates new mutation for the File entity.
func newFileMutation(c config, op Op, opts ...fileOption) *FileMutation {
	m := &FileMutation{
		config:        c,
		op:            op,
		typ:           TypeFile,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFileID sets the ID field of the mutation.
func withFileID(id uuid.UUID) fileOption {
	return func(m *FileMutation) {
		var (
			err   error
			once  sync.Once
			value *File
		)
		m.oldValue = func(ctx context.Context) (*File, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().File.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFile sets the old File of the mutation.
func withFile(node *File) fileOption {
	return func(m *FileMutation) {
		m.oldValue = func(context.Context) (*File, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of File entities.
func (m *FileMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FileMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *FileMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().File.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetStorageKey sets the "storage_key" field.
func (m *FileMutation) SetStorageKey(s string) {
	m.storage_key = &s
}

// StorageKey returns the value of the "storage_key" field in the mutation.
func (m *FileMutation) StorageKey() (r string, exists bool) {
	v := m.storage_key
	if v == nil {
		return
	}
	return *v, true
}

// OldStorageKey returns the old "storage_key" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldStorageKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStorageKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStorageKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStorageKey: %w", err)
	}
	return oldValue.StorageKey, nil
}

// ResetStorageKey resets all changes to the "storage_key" field.
func (m *FileMutation) ResetStorageKey() {
	m.storage_key = nil
}

// SetExpiresAt sets the "expires_at" field.
func (m *FileMutation) SetExpiresAt(t time.Time) {
	m.expires_at = &t
}

// ExpiresAt returns the value of the "expires_at" field in the mutation.
func (m *FileMutation) ExpiresAt() (r time.Time, exists bool) {
	v := m.expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiresAt returns the old "expires_at" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldExpiresAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExpiresAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiresAt: %w", err)
	}
	return oldValue.ExpiresAt, nil
}

// ResetExpiresAt resets all changes to the "expires_at" field.
func (m *FileMutation) ResetExpiresAt() {
	m.expires_at = nil
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *FileMutation) SetUserID(id uuid.UUID) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *FileMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *FileMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *FileMutation) UserID() (id uuid.UUID, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *FileMutation) UserIDs() (ids []uuid.UUID) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *FileMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the FileMutation builder.
func (m *FileMutation) Where(ps ...predicate.File) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *FileMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (File).
func (m *FileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FileMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.storage_key != nil {
		fields = append(fields, file.FieldStorageKey)
	}
	if m.expires_at != nil {
		fields = append(fields, file.FieldExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case file.FieldStorageKey:
		return m.StorageKey()
	case file.FieldExpiresAt:
		return m.ExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FileMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case file.FieldStorageKey:
		return m.OldStorageKey(ctx)
	case file.FieldExpiresAt:
		return m.OldExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown File field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case file.FieldStorageKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStorageKey(v)
		return nil
	case file.FieldExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FileMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FileMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FileMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown File numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FileMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FileMutation) ClearField(name string) error {
	return fmt.Errorf("unknown File nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FileMutation) ResetField(name string) error {
	switch name {
	case file.FieldStorageKey:
		m.ResetStorageKey()
		return nil
	case file.FieldExpiresAt:
		m.ResetExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FileMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, file.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FileMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case file.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FileMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, file.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FileMutation) EdgeCleared(name string) bool {
	switch name {
	case file.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FileMutation) ClearEdge(name string) error {
	switch name {
	case file.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown File unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FileMutation) ResetEdge(name string) error {
	switch name {
	case file.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown File edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	email         *string
	password      *string
	clearedFields map[string]struct{}
	files         map[uuid.UUID]struct{}
	removedfiles  map[uuid.UUID]struct{}
	clearedfiles  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// AddFileIDs adds the "files" edge to the File entity by ids.
func (m *UserMutation) AddFileIDs(ids ...uuid.UUID) {
	if m.files == nil {
		m.files = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.files[ids[i]] = struct{}{}
	}
}

// ClearFiles clears the "files" edge to the File entity.
func (m *UserMutation) ClearFiles() {
	m.clearedfiles = true
}

// FilesCleared reports if the "files" edge to the File entity was cleared.
func (m *UserMutation) FilesCleared() bool {
	return m.clearedfiles
}

// RemoveFileIDs removes the "files" edge to the File entity by IDs.
func (m *UserMutation) RemoveFileIDs(ids ...uuid.UUID) {
	if m.removedfiles == nil {
		m.removedfiles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.files, ids[i])
		m.removedfiles[ids[i]] = struct{}{}
	}
}

// RemovedFiles returns the removed IDs of the "files" edge to the File entity.
func (m *UserMutation) RemovedFilesIDs() (ids []uuid.UUID) {
	for id := range m.removedfiles {
		ids = append(ids, id)
	}
	return
}

// FilesIDs returns the "files" edge IDs in the mutation.
func (m *UserMutation) FilesIDs() (ids []uuid.UUID) {
	for id := range m.files {
		ids = append(ids, id)
	}
	return
}

// ResetFiles resets all changes to the "files" edge.
func (m *UserMutation) ResetFiles() {
	m.files = nil
	m.clearedfiles = false
	m.removedfiles = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.files != nil {
		edges = append(edges, user.EdgeFiles)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFiles:
		ids := make([]ent.Value, 0, len(m.files))
		for id := range m.files {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedfiles != nil {
		edges = append(edges, user.EdgeFiles)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFiles:
		ids := make([]ent.Value, 0, len(m.removedfiles))
		for id := range m.removedfiles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedfiles {
		edges = append(edges, user.EdgeFiles)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeFiles:
		return m.clearedfiles
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeFiles:
		m.ResetFiles()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}