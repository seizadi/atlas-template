// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/seizadi/cmdb/pkg/pb/cmdb.proto

// Generated with protoc-gen-gorm version: master
// Anticipating compatibility with atlas-app-toolkit version: master

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/seizadi/cmdb/pkg/pb/cmdb.proto

It has these top-level messages:
	Region
	CreateRegionRequest
	CreateRegionResponse
	ReadRegionRequest
	ReadRegionResponse
	UpdateRegionRequest
	UpdateRegionResponse
	DeleteRegionRequest
	DeleteRegionResponse
	ListRegionRequest
	ListRegionsResponse
	VersionResponse
*/
package pb

import context "context"
import errors "errors"

import auth1 "github.com/infobloxopen/atlas-app-toolkit/auth"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gateway1 "github.com/infobloxopen/atlas-app-toolkit/gateway"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import query1 "github.com/infobloxopen/atlas-app-toolkit/query"
import resource1 "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"

import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
import _ "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import _ "github.com/infobloxopen/protoc-gen-atlas-validate/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type RegionORM struct {
	AccountID   string
	Description string
	Id          int64 `gorm:"type:serial;primary_key"`
	Name        string
}

// TableName overrides the default tablename generated by GORM
func (RegionORM) TableName() string {
	return "regions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Region) ToORM(ctx context.Context) (RegionORM, error) {
	to := RegionORM{}
	var err error
	if prehook, ok := interface{}(m).(RegionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.DecodeInt64(&Region{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Name = m.Name
	to.Description = m.Description
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(RegionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *RegionORM) ToPB(ctx context.Context) (Region, error) {
	to := Region{}
	var err error
	if prehook, ok := interface{}(m).(RegionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Encode(&Region{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Name = m.Name
	to.Description = m.Description
	if posthook, ok := interface{}(m).(RegionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Region the arg will be the target, the caller the one being converted from

// RegionBeforeToORM called before default ToORM code
type RegionWithBeforeToORM interface {
	BeforeToORM(context.Context, *RegionORM) error
}

// RegionAfterToORM called after default ToORM code
type RegionWithAfterToORM interface {
	AfterToORM(context.Context, *RegionORM) error
}

// RegionBeforeToPB called before default ToPB code
type RegionWithBeforeToPB interface {
	BeforeToPB(context.Context, *Region) error
}

// RegionAfterToPB called after default ToPB code
type RegionWithAfterToPB interface {
	AfterToPB(context.Context, *Region) error
}

// DefaultCreateRegion executes a basic gorm create call
func DefaultCreateRegion(ctx context.Context, in *Region, db *gorm1.DB) (*Region, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateRegion")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeCreate); ok {
		if db, err = hook.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithAfterCreate); ok {
		if err = hook.AfterCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type RegionORMWithBeforeCreate interface {
	BeforeCreate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type RegionORMWithAfterCreate interface {
	AfterCreate(context.Context, *gorm1.DB) error
}

// DefaultReadRegion executes a basic gorm read call
func DefaultReadRegion(ctx context.Context, in *Region, db *gorm1.DB, fs *query1.FieldSelection) (*Region, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadRegion")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.New("DefaultReadRegion requires a non-zero primary key")
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db, fs); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, fs, &RegionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db, fs); err != nil {
			return nil, err
		}
	}
	ormResponse := RegionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(RegionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db, fs); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type RegionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB, *query1.FieldSelection) (*gorm1.DB, error)
}
type RegionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB, *query1.FieldSelection) (*gorm1.DB, error)
}
type RegionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB, *query1.FieldSelection) error
}

func DefaultDeleteRegion(ctx context.Context, in *Region, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteRegion")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeDelete); ok {
		if db, err = hook.BeforeDelete(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&RegionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithAfterDelete); ok {
		err = hook.AfterDelete(ctx, db)
	}
	return err
}

type RegionORMWithBeforeDelete interface {
	BeforeDelete(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type RegionORMWithAfterDelete interface {
	AfterDelete(context.Context, *gorm1.DB) error
}

// DefaultStrictUpdateRegion clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateRegion(ctx context.Context, in *Region, db *gorm1.DB) (*Region, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateRegion")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(map[string]interface{}{"account_id": accountID})
	var count int64
	lockedRow := &RegionORM{}
	count = db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow).RowsAffected
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

type RegionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type RegionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type RegionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchRegion executes a basic gorm update call with patch behavior
func DefaultPatchRegion(ctx context.Context, in *Region, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Region, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultPatchRegion")
	}
	var pbObj Region
	var err error
	if hook, ok := interface{}(&pbObj).(RegionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadRegion(ctx, &Region{Id: in.GetId()}, db, nil)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(RegionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskRegion(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(RegionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateRegion(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(RegionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type RegionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Region, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type RegionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Region, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type RegionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Region, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type RegionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Region, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultApplyFieldMaskRegion patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskRegion(ctx context.Context, patchee *Region, patcher *Region, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Region, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.New("Patchee inputs to DefaultApplyFieldMaskRegion must be non-nil")
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Description" {
			patchee.Description = patcher.Description
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListRegion executes a gorm list call
func DefaultListRegion(ctx context.Context, db *gorm1.DB, f *query1.Filtering, s *query1.Sorting, p *query1.Pagination, fs *query1.FieldSelection) ([]*Region, error) {
	in := Region{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &RegionORM{}, &Region{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []RegionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RegionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Region{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type RegionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type RegionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type RegionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]RegionORM, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) error
}
type RegionsDefaultServer struct {
}

// Create ...
func (m *RegionsDefaultServer) Create(ctx context.Context, in *CreateRegionRequest) (*CreateRegionResponse, error) {
	txn, ok := gorm2.FromContext(ctx)
	if !ok {
		return nil, errors.New("Database Transaction For Request Missing")
	}
	db := txn.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	if custom, ok := interface{}(in).(RegionsRegionWithBeforeCreate); ok {
		var err error
		if db, err = custom.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateRegion(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateRegionResponse{Result: res}
	if custom, ok := interface{}(in).(RegionsRegionWithAfterCreate); ok {
		var err error
		if err = custom.AfterCreate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RegionsRegionWithBeforeCreate called before DefaultCreateRegion in the default Create handler
type RegionsRegionWithBeforeCreate interface {
	BeforeCreate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// RegionsRegionWithAfterCreate called before DefaultCreateRegion in the default Create handler
type RegionsRegionWithAfterCreate interface {
	AfterCreate(context.Context, *CreateRegionResponse, *gorm1.DB) error
}

// Read ...
func (m *RegionsDefaultServer) Read(ctx context.Context, in *ReadRegionRequest) (*ReadRegionResponse, error) {
	txn, ok := gorm2.FromContext(ctx)
	if !ok {
		return nil, errors.New("Database Transaction For Request Missing")
	}
	db := txn.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	if custom, ok := interface{}(in).(RegionsRegionWithBeforeRead); ok {
		var err error
		if db, err = custom.BeforeRead(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadRegion(ctx, &Region{Id: in.GetId()}, db, in.Fields)
	if err != nil {
		return nil, err
	}
	out := &ReadRegionResponse{Result: res}
	if custom, ok := interface{}(in).(RegionsRegionWithAfterRead); ok {
		var err error
		if err = custom.AfterRead(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RegionsRegionWithBeforeRead called before DefaultReadRegion in the default Read handler
type RegionsRegionWithBeforeRead interface {
	BeforeRead(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// RegionsRegionWithAfterRead called before DefaultReadRegion in the default Read handler
type RegionsRegionWithAfterRead interface {
	AfterRead(context.Context, *ReadRegionResponse, *gorm1.DB) error
}

// Update ...
func (m *RegionsDefaultServer) Update(ctx context.Context, in *UpdateRegionRequest) (*UpdateRegionResponse, error) {
	var err error
	var res *Region
	txn, ok := gorm2.FromContext(ctx)
	if !ok {
		return nil, errors.New("Database Transaction For Request Missing")
	}
	db := txn.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	if custom, ok := interface{}(in).(RegionsRegionWithBeforeUpdate); ok {
		var err error
		if db, err = custom.BeforeUpdate(ctx, db); err != nil {
			return nil, err
		}
	}
	if in.GetFields() == nil {
		res, err = DefaultStrictUpdateRegion(ctx, in.GetPayload(), db)
	} else {
		res, err = DefaultPatchRegion(ctx, in.GetPayload(), in.GetFields(), db)
	}
	if err != nil {
		return nil, err
	}
	out := &UpdateRegionResponse{Result: res}
	if custom, ok := interface{}(in).(RegionsRegionWithAfterUpdate); ok {
		var err error
		if err = custom.AfterUpdate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RegionsRegionWithBeforeUpdate called before DefaultUpdateRegion in the default Update handler
type RegionsRegionWithBeforeUpdate interface {
	BeforeUpdate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// RegionsRegionWithAfterUpdate called before DefaultUpdateRegion in the default Update handler
type RegionsRegionWithAfterUpdate interface {
	AfterUpdate(context.Context, *UpdateRegionResponse, *gorm1.DB) error
}

// Delete ...
func (m *RegionsDefaultServer) Delete(ctx context.Context, in *DeleteRegionRequest) (*DeleteRegionResponse, error) {
	txn, ok := gorm2.FromContext(ctx)
	if !ok {
		return nil, errors.New("Database Transaction For Request Missing")
	}
	db := txn.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	if custom, ok := interface{}(in).(RegionsRegionWithBeforeDelete); ok {
		var err error
		if db, err = custom.BeforeDelete(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteRegion(ctx, &Region{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteRegionResponse{}
	if custom, ok := interface{}(in).(RegionsRegionWithAfterDelete); ok {
		var err error
		if err = custom.AfterDelete(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RegionsRegionWithBeforeDelete called before DefaultDeleteRegion in the default Delete handler
type RegionsRegionWithBeforeDelete interface {
	BeforeDelete(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// RegionsRegionWithAfterDelete called before DefaultDeleteRegion in the default Delete handler
type RegionsRegionWithAfterDelete interface {
	AfterDelete(context.Context, *DeleteRegionResponse, *gorm1.DB) error
}

// List ...
func (m *RegionsDefaultServer) List(ctx context.Context, in *ListRegionRequest) (*ListRegionsResponse, error) {
	txn, ok := gorm2.FromContext(ctx)
	if !ok {
		return nil, errors.New("Database Transaction For Request Missing")
	}
	db := txn.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	if custom, ok := interface{}(in).(RegionsRegionWithBeforeList); ok {
		var err error
		if db, err = custom.BeforeList(ctx, db); err != nil {
			return nil, err
		}
	}
	pagedRequest := false
	if in.GetPaging().GetLimit() >= 1 {
		in.Paging.Limit++
		pagedRequest = true
	}
	res, err := DefaultListRegion(ctx, db, in.Filter, in.OrderBy, in.Paging, in.Fields)
	if err != nil {
		return nil, err
	}
	var resPaging *query1.PageInfo
	if pagedRequest {
		var offset int32
		var size int32 = int32(len(res))
		if size == in.GetPaging().GetLimit() {
			size--
			res = res[:size]
			offset = in.GetPaging().GetOffset() + size
		}
		resPaging = &query1.PageInfo{Offset: offset}
	}
	out := &ListRegionsResponse{Results: res, Page: resPaging}
	if custom, ok := interface{}(in).(RegionsRegionWithAfterList); ok {
		var err error
		if err = custom.AfterList(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// RegionsRegionWithBeforeList called before DefaultListRegion in the default List handler
type RegionsRegionWithBeforeList interface {
	BeforeList(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// RegionsRegionWithAfterList called before DefaultListRegion in the default List handler
type RegionsRegionWithAfterList interface {
	AfterList(context.Context, *ListRegionsResponse, *gorm1.DB) error
}
