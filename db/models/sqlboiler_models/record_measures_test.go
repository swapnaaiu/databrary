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

func testRecordMeasures(t *testing.T) {
	t.Parallel()

	query := RecordMeasures(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testRecordMeasuresDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = recordMeasure.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordMeasuresQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = RecordMeasures(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordMeasuresSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := RecordMeasureSlice{recordMeasure}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testRecordMeasuresExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := RecordMeasureExists(tx, recordMeasure.ID)
	if err != nil {
		t.Errorf("Unable to check if RecordMeasure exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RecordMeasureExistsG to return true, but got false.")
	}
}
func testRecordMeasuresFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	recordMeasureFound, err := FindRecordMeasure(tx, recordMeasure.ID)
	if err != nil {
		t.Error(err)
	}

	if recordMeasureFound == nil {
		t.Error("want a record, got nil")
	}
}
func testRecordMeasuresBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = RecordMeasures(tx).Bind(recordMeasure); err != nil {
		t.Error(err)
	}
}

func testRecordMeasuresOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := RecordMeasures(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRecordMeasuresAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasureOne := &RecordMeasure{}
	recordMeasureTwo := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasureOne, recordMeasureDBTypes, false, recordMeasureColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}
	if err = randomize.Struct(seed, recordMeasureTwo, recordMeasureDBTypes, false, recordMeasureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = recordMeasureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := RecordMeasures(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRecordMeasuresCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasureOne := &RecordMeasure{}
	recordMeasureTwo := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasureOne, recordMeasureDBTypes, false, recordMeasureColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}
	if err = randomize.Struct(seed, recordMeasureTwo, recordMeasureDBTypes, false, recordMeasureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = recordMeasureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func recordMeasureBeforeInsertHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureAfterInsertHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureAfterSelectHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureBeforeUpdateHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureAfterUpdateHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureBeforeDeleteHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureAfterDeleteHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureBeforeUpsertHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func recordMeasureAfterUpsertHook(e boil.Executor, o *RecordMeasure) error {
	*o = RecordMeasure{}
	return nil
}

func testRecordMeasuresHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	empty := &RecordMeasure{}

	AddRecordMeasureHook(boil.BeforeInsertHook, recordMeasureBeforeInsertHook)
	if err = recordMeasure.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureBeforeInsertHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.AfterInsertHook, recordMeasureAfterInsertHook)
	if err = recordMeasure.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureAfterInsertHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.AfterSelectHook, recordMeasureAfterSelectHook)
	if err = recordMeasure.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureAfterSelectHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.BeforeUpdateHook, recordMeasureBeforeUpdateHook)
	if err = recordMeasure.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureBeforeUpdateHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.AfterUpdateHook, recordMeasureAfterUpdateHook)
	if err = recordMeasure.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureAfterUpdateHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.BeforeDeleteHook, recordMeasureBeforeDeleteHook)
	if err = recordMeasure.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureBeforeDeleteHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.AfterDeleteHook, recordMeasureAfterDeleteHook)
	if err = recordMeasure.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureAfterDeleteHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.BeforeUpsertHook, recordMeasureBeforeUpsertHook)
	if err = recordMeasure.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureBeforeUpsertHooks = []RecordMeasureHook{}

	AddRecordMeasureHook(boil.AfterUpsertHook, recordMeasureAfterUpsertHook)
	if err = recordMeasure.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(recordMeasure, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", recordMeasure)
	}
	recordMeasureAfterUpsertHooks = []RecordMeasureHook{}
}
func testRecordMeasuresInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecordMeasuresInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx, recordMeasureColumns...); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecordMeasureToOneRecordUsingID(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Record
	var local RecordMeasure

	foreignBlacklist := recordColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, recordDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := recordMeasureColumnsWithDefault
	if err := randomize.Struct(seed, &local, recordMeasureDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.IDByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RecordMeasureSlice{&local}
	if err = local.L.LoadID(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ID == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ID = nil
	if err = local.L.LoadID(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ID == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRecordMeasureToOneSetOpRecordUsingID(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a RecordMeasure
	var b, c Record

	foreignBlacklist := strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(recordMeasurePrimaryKeyColumns, recordMeasureColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, recordMeasureDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Record{&b, &c} {
		err = a.SetID(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ID != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.IDRecordMeasure != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ID != x.ID {
			t.Error("foreign key was wrong value", a.ID)
		}

		if exists, err := RecordMeasureExists(tx, a.ID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testRecordMeasuresReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = recordMeasure.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testRecordMeasuresReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := RecordMeasureSlice{recordMeasure}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testRecordMeasuresSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := RecordMeasures(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	recordMeasureDBTypes = map[string]string{`Category`: `smallint`, `ID`: `integer`, `Measures`: `ARRAYtext`, `Volume`: `integer`}
	_                    = bytes.MinRead
)

func testRecordMeasuresUpdate(t *testing.T) {
	t.Parallel()

	if len(recordMeasureColumns) == len(recordMeasurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := recordMeasureColumnsWithDefault

	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	if err = recordMeasure.Update(tx); err != nil {
		t.Error(err)
	}
}

func testRecordMeasuresSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(recordMeasureColumns) == len(recordMeasurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := recordMeasurePrimaryKeyColumns

	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(recordMeasureColumns, recordMeasurePrimaryKeyColumns) {
		fields = recordMeasureColumns
	} else {
		fields = strmangle.SetComplement(
			recordMeasureColumns,
			recordMeasurePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(recordMeasure))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := RecordMeasureSlice{recordMeasure}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testRecordMeasuresUpsert(t *testing.T) {
	t.Parallel()

	if len(recordMeasureColumns) == len(recordMeasurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	recordMeasure := &RecordMeasure{}
	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = recordMeasure.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert RecordMeasure: %s", err)
	}

	count, err := RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := recordMeasurePrimaryKeyColumns

	if err = randomize.Struct(seed, recordMeasure, recordMeasureDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize RecordMeasure struct: %s", err)
	}

	if err = recordMeasure.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert RecordMeasure: %s", err)
	}

	count, err = RecordMeasures(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
