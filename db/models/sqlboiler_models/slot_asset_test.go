// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/databrary/databrary/db/models/custom_types"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testSlotAssets(t *testing.T) {
	t.Parallel()

	query := SlotAssets(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSlotAssetsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = slotAsset.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotAssetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SlotAssets(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotAssetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SlotAssetSlice{slotAsset}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSlotAssetsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SlotAssetExists(tx, slotAsset.Asset)
	if err != nil {
		t.Errorf("Unable to check if SlotAsset exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SlotAssetExistsG to return true, but got false.")
	}
}
func testSlotAssetsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slotAssetFound, err := FindSlotAsset(tx, slotAsset.Asset)
	if err != nil {
		t.Error(err)
	}

	if slotAssetFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSlotAssetsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SlotAssets(tx).Bind(slotAsset); err != nil {
		t.Error(err)
	}
}

func testSlotAssetsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := SlotAssets(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSlotAssetsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAssetOne := &SlotAsset{}
	slotAssetTwo := &SlotAsset{}
	if err = randomize.Struct(seed, slotAssetOne, slotAssetDBTypes, false, slotAssetColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	if err = randomize.Struct(seed, slotAssetTwo, slotAssetDBTypes, false, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAssetOne.Segment = custom_types.SegmentRandom()
	slotAssetTwo.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAssetOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = slotAssetTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SlotAssets(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSlotAssetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAssetOne := &SlotAsset{}
	slotAssetTwo := &SlotAsset{}
	if err = randomize.Struct(seed, slotAssetOne, slotAssetDBTypes, false, slotAssetColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	if err = randomize.Struct(seed, slotAssetTwo, slotAssetDBTypes, false, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAssetOne.Segment = custom_types.SegmentRandom()
	slotAssetTwo.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAssetOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = slotAssetTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func slotAssetBeforeInsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterInsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterSelectHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetBeforeUpdateHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterUpdateHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetBeforeDeleteHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterDeleteHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetBeforeUpsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func slotAssetAfterUpsertHook(e boil.Executor, o *SlotAsset) error {
	*o = SlotAsset{}
	return nil
}

func testSlotAssetsHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	empty := &SlotAsset{}

	AddSlotAssetHook(boil.BeforeInsertHook, slotAssetBeforeInsertHook)
	if err = slotAsset.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetBeforeInsertHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterInsertHook, slotAssetAfterInsertHook)
	if err = slotAsset.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetAfterInsertHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterSelectHook, slotAssetAfterSelectHook)
	if err = slotAsset.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetAfterSelectHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.BeforeUpdateHook, slotAssetBeforeUpdateHook)
	if err = slotAsset.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetBeforeUpdateHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterUpdateHook, slotAssetAfterUpdateHook)
	if err = slotAsset.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetAfterUpdateHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.BeforeDeleteHook, slotAssetBeforeDeleteHook)
	if err = slotAsset.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetBeforeDeleteHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterDeleteHook, slotAssetAfterDeleteHook)
	if err = slotAsset.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetAfterDeleteHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.BeforeUpsertHook, slotAssetBeforeUpsertHook)
	if err = slotAsset.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetBeforeUpsertHooks = []SlotAssetHook{}

	AddSlotAssetHook(boil.AfterUpsertHook, slotAssetAfterUpsertHook)
	if err = slotAsset.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotAsset, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", slotAsset)
	}
	slotAssetAfterUpsertHooks = []SlotAssetHook{}
}
func testSlotAssetsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlotAssetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx, slotAssetColumns...); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlotAssetToManyAssetExcerpts(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a SlotAsset
	var b, c Excerpt

	foreignBlacklist := excerptColumnsWithDefault
	foreignBlacklist = append(foreignBlacklist, excerptColumnsWithCustom...)

	if err := randomize.Struct(seed, &b, excerptDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Excerpt struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, excerptDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Excerpt struct: %s", err)
	}
	b.Segment = custom_types.SegmentRandom()
	c.Segment = custom_types.SegmentRandom()
	b.Release = custom_types.NullReleaseRandom()
	c.Release = custom_types.NullReleaseRandom()

	localBlacklist := slotAssetColumnsWithDefault
	localBlacklist = append(localBlacklist, slotAssetColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, slotAssetDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	b.Asset = a.Asset
	c.Asset = a.Asset
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	excerpt, err := a.AssetExcerptsByFk(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range excerpt {
		if v.Asset == b.Asset {
			bFound = true
		}
		if v.Asset == c.Asset {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SlotAssetSlice{&a}
	if err = a.L.LoadAssetExcerpts(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssetExcerpts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AssetExcerpts = nil
	if err = a.L.LoadAssetExcerpts(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssetExcerpts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", excerpt)
	}
}

func testSlotAssetToManyAddOpAssetExcerpts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SlotAsset
	var b, c, d, e Excerpt

	seed := randomize.NewSeed()
	localComplelementList := strmangle.SetComplement(slotAssetPrimaryKeyColumns, slotAssetColumnsWithoutDefault)
	localComplelementList = append(localComplelementList, slotAssetColumnsWithCustom...)

	if err = randomize.Struct(seed, &a, slotAssetDBTypes, false, localComplelementList...); err != nil {
		t.Fatal(err)
	}
	a.Segment = custom_types.SegmentRandom()

	foreignComplementList := strmangle.SetComplement(excerptPrimaryKeyColumns, excerptColumnsWithoutDefault)
	foreignComplementList = append(foreignComplementList, excerptColumnsWithCustom...)

	foreigners := []*Excerpt{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, excerptDBTypes, false, foreignComplementList...); err != nil {
			t.Fatal(err)
		}
		x.Segment = custom_types.SegmentRandom()
		x.Release = custom_types.NullReleaseRandom()

	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Excerpt{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAssetExcerpts(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Asset != first.Asset {
			t.Error("foreign key was wrong value", a.Asset, first.Asset)
		}
		if a.Asset != second.Asset {
			t.Error("foreign key was wrong value", a.Asset, second.Asset)
		}

		if first.R.Asset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Asset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AssetExcerpts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AssetExcerpts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AssetExcerptsByFk(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSlotAssetToOneContainerUsingContainer(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Container
	var local SlotAsset

	foreignBlacklist := containerColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, containerDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	localBlacklist := slotAssetColumnsWithDefault
	localBlacklist = append(localBlacklist, slotAssetColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, slotAssetDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	local.Segment = custom_types.SegmentRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Container = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContainerByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SlotAssetSlice{&local}
	if err = local.L.LoadContainer(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Container == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Container = nil
	if err = local.L.LoadContainer(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Container == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSlotAssetToOneAssetUsingAsset(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Asset
	var local SlotAsset

	foreignBlacklist := assetColumnsWithDefault
	foreignBlacklist = append(foreignBlacklist, assetColumnsWithCustom...)

	if err := randomize.Struct(seed, &foreign, assetDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}
	foreign.Release = custom_types.NullReleaseRandom()

	localBlacklist := slotAssetColumnsWithDefault
	localBlacklist = append(localBlacklist, slotAssetColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, slotAssetDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	local.Segment = custom_types.SegmentRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Asset = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AssetByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SlotAssetSlice{&local}
	if err = local.L.LoadAsset(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Asset == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Asset = nil
	if err = local.L.LoadAsset(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Asset == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSlotAssetToOneSetOpContainerUsingContainer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a SlotAsset
	var b, c Container

	foreignBlacklist := strmangle.SetComplement(containerPrimaryKeyColumns, containerColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, containerDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, containerDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(slotAssetPrimaryKeyColumns, slotAssetColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, slotAssetColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, slotAssetDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Container{&b, &c} {
		err = a.SetContainer(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Container != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SlotAssets[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Container != x.ID {
			t.Error("foreign key was wrong value", a.Container)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Container))
		reflect.Indirect(reflect.ValueOf(&a.Container)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Container != x.ID {
			t.Error("foreign key was wrong value", a.Container, x.ID)
		}
	}
}
func testSlotAssetToOneSetOpAssetUsingAsset(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a SlotAsset
	var b, c Asset

	foreignBlacklist := strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)
	foreignBlacklist = append(foreignBlacklist, assetColumnsWithCustom...)

	if err := randomize.Struct(seed, &b, assetDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, assetDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}
	b.Release = custom_types.NullReleaseRandom()
	c.Release = custom_types.NullReleaseRandom()

	localBlacklist := strmangle.SetComplement(slotAssetPrimaryKeyColumns, slotAssetColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, slotAssetColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, slotAssetDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Asset{&b, &c} {
		err = a.SetAsset(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Asset != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SlotAsset != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Asset != x.ID {
			t.Error("foreign key was wrong value", a.Asset)
		}

		if exists, err := SlotAssetExists(tx, a.Asset); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testSlotAssetsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = slotAsset.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSlotAssetsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SlotAssetSlice{slotAsset}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSlotAssetsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SlotAssets(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	slotAssetDBTypes = map[string]string{`Asset`: `integer`, `Container`: `integer`, `Segment`: `USER-DEFINED`}
	_                = bytes.MinRead
)

func testSlotAssetsUpdate(t *testing.T) {
	t.Parallel()

	if len(slotAssetColumns) == len(slotAssetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := slotAssetColumnsWithDefault
	blacklist = append(blacklist, slotAssetColumnsWithCustom...)

	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	if err = slotAsset.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSlotAssetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(slotAssetColumns) == len(slotAssetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := slotAssetPrimaryKeyColumns
	blacklist = append(blacklist, slotAssetColumnsWithCustom...)

	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(slotAssetColumns, slotAssetPrimaryKeyColumns) {
		fields = slotAssetColumns
	} else {
		fields = strmangle.SetComplement(
			slotAssetColumns,
			slotAssetPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(slotAsset))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SlotAssetSlice{slotAsset}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSlotAssetsUpsert(t *testing.T) {
	t.Parallel()

	if len(slotAssetColumns) == len(slotAssetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	slotAsset := &SlotAsset{}
	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, true, slotAssetColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotAsset.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert SlotAsset: %s", err)
	}

	count, err := SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := slotAssetPrimaryKeyColumns

	blacklist = append(blacklist, slotAssetColumnsWithCustom...)

	if err = randomize.Struct(seed, slotAsset, slotAssetDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize SlotAsset struct: %s", err)
	}

	slotAsset.Segment = custom_types.SegmentRandom()

	if err = slotAsset.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert SlotAsset: %s", err)
	}

	count, err = SlotAssets(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
