// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testUploads(t *testing.T) {
	t.Parallel()

	query := Uploads(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUploadsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upload.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUploadsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Uploads(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUploadsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UploadSlice{upload}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUploadsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UploadExists(tx, upload.Token)
	if err != nil {
		t.Errorf("Unable to check if Upload exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UploadExistsG to return true, but got false.")
	}
}
func testUploadsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	uploadFound, err := FindUpload(tx, upload.Token)
	if err != nil {
		t.Error(err)
	}

	if uploadFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUploadsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Uploads(tx).Bind(upload); err != nil {
		t.Error(err)
	}
}

func testUploadsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Uploads(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUploadsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	uploadOne := &Upload{}
	uploadTwo := &Upload{}
	if err = randomize.Struct(seed, uploadOne, uploadDBTypes, false, uploadColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Upload struct: %s", err)
	}
	if err = randomize.Struct(seed, uploadTwo, uploadDBTypes, false, uploadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = uploadOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = uploadTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Uploads(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUploadsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	uploadOne := &Upload{}
	uploadTwo := &Upload{}
	if err = randomize.Struct(seed, uploadOne, uploadDBTypes, false, uploadColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Upload struct: %s", err)
	}
	if err = randomize.Struct(seed, uploadTwo, uploadDBTypes, false, uploadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = uploadOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = uploadTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func uploadBeforeInsertHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadAfterInsertHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadAfterSelectHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadBeforeUpdateHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadAfterUpdateHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadBeforeDeleteHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadAfterDeleteHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadBeforeUpsertHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func uploadAfterUpsertHook(e boil.Executor, o *Upload) error {
	*o = Upload{}
	return nil
}

func testUploadsHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	empty := &Upload{}

	AddUploadHook(boil.BeforeInsertHook, uploadBeforeInsertHook)
	if err = upload.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", upload)
	}
	uploadBeforeInsertHooks = []UploadHook{}

	AddUploadHook(boil.AfterInsertHook, uploadAfterInsertHook)
	if err = upload.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", upload)
	}
	uploadAfterInsertHooks = []UploadHook{}

	AddUploadHook(boil.AfterSelectHook, uploadAfterSelectHook)
	if err = upload.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", upload)
	}
	uploadAfterSelectHooks = []UploadHook{}

	AddUploadHook(boil.BeforeUpdateHook, uploadBeforeUpdateHook)
	if err = upload.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", upload)
	}
	uploadBeforeUpdateHooks = []UploadHook{}

	AddUploadHook(boil.AfterUpdateHook, uploadAfterUpdateHook)
	if err = upload.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", upload)
	}
	uploadAfterUpdateHooks = []UploadHook{}

	AddUploadHook(boil.BeforeDeleteHook, uploadBeforeDeleteHook)
	if err = upload.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", upload)
	}
	uploadBeforeDeleteHooks = []UploadHook{}

	AddUploadHook(boil.AfterDeleteHook, uploadAfterDeleteHook)
	if err = upload.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", upload)
	}
	uploadAfterDeleteHooks = []UploadHook{}

	AddUploadHook(boil.BeforeUpsertHook, uploadBeforeUpsertHook)
	if err = upload.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", upload)
	}
	uploadBeforeUpsertHooks = []UploadHook{}

	AddUploadHook(boil.AfterUpsertHook, uploadAfterUpsertHook)
	if err = upload.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(upload, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", upload)
	}
	uploadAfterUpsertHooks = []UploadHook{}
}
func testUploadsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUploadsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx, uploadColumns...); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUploadToOneVolumeUsingVolume(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Volume
	var local Upload

	foreignBlacklist := volumeColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, volumeDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	localBlacklist := uploadColumnsWithDefault
	if err := randomize.Struct(seed, &local, uploadDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Volume = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.VolumeByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UploadSlice{&local}
	if err = local.L.LoadVolume(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Volume = nil
	if err = local.L.LoadVolume(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUploadToOneAccountUsingAccount(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Account
	var local Upload

	foreignBlacklist := accountColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, accountDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	localBlacklist := uploadColumnsWithDefault
	if err := randomize.Struct(seed, &local, uploadDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Account = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AccountByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UploadSlice{&local}
	if err = local.L.LoadAccount(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Account == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Account = nil
	if err = local.L.LoadAccount(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Account == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUploadToOneSetOpVolumeUsingVolume(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Upload
	var b, c Volume

	foreignBlacklist := strmangle.SetComplement(volumePrimaryKeyColumns, volumeColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, volumeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, volumeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(uploadPrimaryKeyColumns, uploadColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, uploadDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Volume{&b, &c} {
		err = a.SetVolume(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Volume != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Uploads[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Volume != x.ID {
			t.Error("foreign key was wrong value", a.Volume)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Volume))
		reflect.Indirect(reflect.ValueOf(&a.Volume)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Volume != x.ID {
			t.Error("foreign key was wrong value", a.Volume, x.ID)
		}
	}
}
func testUploadToOneSetOpAccountUsingAccount(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Upload
	var b, c Account

	foreignBlacklist := strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, accountDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, accountDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(uploadPrimaryKeyColumns, uploadColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, uploadDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Account{&b, &c} {
		err = a.SetAccount(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Account != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Uploads[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Account != x.ID {
			t.Error("foreign key was wrong value", a.Account)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Account))
		reflect.Indirect(reflect.ValueOf(&a.Account)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Account != x.ID {
			t.Error("foreign key was wrong value", a.Account, x.ID)
		}
	}
}
func testUploadsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upload.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUploadsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UploadSlice{upload}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUploadsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Uploads(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	uploadDBTypes = map[string]string{`Account`: `integer`, `Expires`: `timestamp with time zone`, `Filename`: `text`, `Size`: `bigint`, `Token`: `character varying`, `Volume`: `integer`}
	_             = bytes.MinRead
)

func testUploadsUpdate(t *testing.T) {
	t.Parallel()

	if len(uploadColumns) == len(uploadPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := uploadColumnsWithDefault

	if err = randomize.Struct(seed, upload, uploadDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	if err = upload.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUploadsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(uploadColumns) == len(uploadPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := uploadPrimaryKeyColumns

	if err = randomize.Struct(seed, upload, uploadDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(uploadColumns, uploadPrimaryKeyColumns) {
		fields = uploadColumns
	} else {
		fields = strmangle.SetComplement(
			uploadColumns,
			uploadPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(upload))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UploadSlice{upload}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUploadsUpsert(t *testing.T) {
	t.Parallel()

	if len(uploadColumns) == len(uploadPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	upload := &Upload{}
	if err = randomize.Struct(seed, upload, uploadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upload.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Upload: %s", err)
	}

	count, err := Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := uploadPrimaryKeyColumns

	if err = randomize.Struct(seed, upload, uploadDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize Upload struct: %s", err)
	}

	if err = upload.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Upload: %s", err)
	}

	count, err = Uploads(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
