package types

import (
	"testing"
)

func TestCompatibleTable(t *testing.T) {
	t0Entity := Table0Default()
	t1Entity := Table1Default()
	t2Entity := Table2Default()
	t3Entity := Table3Default()
	t4Entity := Table4Default()
	t5Entity := Table5Default()
	t6Entity := Table6Default()

	t0 := t0Entity.AsSlice()
	t1 := t1Entity.AsSlice()
	t2 := t2Entity.AsSlice()
	t3 := t3Entity.AsSlice()
	t4 := t4Entity.AsSlice()
	t5 := t5Entity.AsSlice()
	t6 := t6Entity.AsSlice()

	var err error

	/* test strict mode */

	// Table0
	_, err = Table0FromSlice(t0, false)
	if err != nil {
		t.Fatal("table0 strict from slice fail on t0")
	}

	_, err = Table0FromSlice(t1, false)
	if err == nil {
		t.Fatal("table0 strict from slice success on t1")
	}

	_, err = Table0FromSlice(t2, false)
	if err == nil {
		t.Fatal("table0 strict from slice success on t2")
	}

	_, err = Table0FromSlice(t3, false)
	if err == nil {
		t.Fatal("table0 strict from slice success on t3")
	}

	_, err = Table0FromSlice(t4, false)
	if err == nil {
		t.Fatal("table0 strict from slice success on t4")
	}
	_, err = Table0FromSlice(t5, false)
	if err == nil {
		t.Fatal("table0 strict from slice success on t5")
	}
	_, err = Table0FromSlice(t6, false)
	if err == nil {
		t.Fatal("table0 strict from slice success on t6")
	}

	// Table2
	_, err = Table2FromSlice(t0, false)
	if err == nil {
		t.Fatal("table2 strict from slice success on t0")
	}

	_, err = Table2FromSlice(t1, false)
	if err == nil {
		t.Fatal("table2 strict from slice success on t1")
	}

	_, err = Table2FromSlice(t2, false)
	if err != nil {
		t.Fatal("table2 strict from slice fail on t2")
	}

	_, err = Table2FromSlice(t3, false)
	if err == nil {
		t.Fatal("table2 strict from slice success on t3")
	}

	_, err = Table2FromSlice(t4, false)
	if err == nil {
		t.Fatal("table2 strict from slice success on t4")
	}
	_, err = Table2FromSlice(t5, false)
	if err == nil {
		t.Fatal("table2 strict from slice success on t5")
	}
	_, err = Table2FromSlice(t6, false)
	if err == nil {
		t.Fatal("table2 strict from slice fail on t6")
	}

	// Table3
	_, err = Table3FromSlice(t0, false)
	if err == nil {
		t.Fatal("table3 strict from slice success on t0")
	}

	_, err = Table3FromSlice(t1, false)
	if err == nil {
		t.Fatal("table3 strict from slice success on t1")
	}

	_, err = Table3FromSlice(t2, false)
	if err == nil {
		t.Fatal("table3 strict from slice success on t2")
	}

	_, err = Table3FromSlice(t3, false)
	if err != nil {
		t.Fatal("table3 strict from slice fail on t3")
	}

	_, err = Table3FromSlice(t4, false)
	if err == nil {
		t.Fatal("table3 strict from slice success on t4")
	}
	_, err = Table3FromSlice(t5, false)
	if err == nil {
		t.Fatal("table3 strict from slice success on t5")
	}
	_, err = Table3FromSlice(t6, false)
	if err == nil {
		t.Fatal("table3 strict from slice success on t6")
	}

	// Table6
	_, err = Table6FromSlice(t0, false)
	if err == nil {
		t.Fatal("table6 strict from slice success on t0")
	}

	_, err = Table6FromSlice(t1, false)
	if err == nil {
		t.Fatal("table6 strict from slice success on t1")
	}

	_, err = Table6FromSlice(t2, false)
	if err == nil {
		t.Fatal("table6 strict from slice success on t2")
	}

	_, err = Table6FromSlice(t3, false)
	if err == nil {
		t.Fatal("table6 strict from slice success on t3")
	}

	_, err = Table6FromSlice(t4, false)
	if err == nil {
		t.Fatal("table6 strict from slice success on t4")
	}
	_, err = Table6FromSlice(t5, false)
	if err == nil {
		t.Fatal("table6 strict from slice success on t5")
	}
	_, err = Table6FromSlice(t6, false)
	if err != nil {
		t.Fatal("table6 strict from slice fail on t6")
	}

	/* test compatible mode */
	var i uint

	// Table0
	var res0 *Table0

	i = 0
	res0, err = Table0FromSlice(t0, true)
	if err != nil {
		t.Fatal("table0 compatible from slice fail on t0")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t0, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t0 has extra fields")
	}

	i++
	res0, err = Table0FromSlice(t1, true)
	if err != nil {
		t.Fatal("table0 compatible from slice fail on t1")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t1, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if !res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t1 doesn't has extra fields")
	}

	i++
	res0, err = Table0FromSlice(t2, true)
	if err != nil {
		t.Fatal("table0 strict from slice fail on t2")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t2, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if !res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t2 doesn't has extra fields")
	}

	i++
	res0, err = Table0FromSlice(t3, true)
	if err != nil {
		t.Fatal("table0 compatible from slice fail on t3")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t3, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if !res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t3 doesn't has extra fields")
	}

	i++
	res0, err = Table0FromSlice(t4, true)
	if err != nil {
		t.Fatal("table0 compatible from slice fail on t4")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t4, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if !res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t4 doesn't has extra fields")
	}

	i++
	res0, err = Table0FromSlice(t5, true)
	if err != nil {
		t.Fatal("table0 compatible from slice fail on t5")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t4, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if !res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t5 doesn't has extra fields")
	}

	i++
	res0, err = Table0FromSlice(t6, true)
	if err != nil {
		t.Fatal("table0 compatible from slice fail on t6")
	}
	if res0.CountExtraFields() != i {
		t.Fatal("table0 compatible from slice on t4, CountExtraFields not ", i, " but ", res0.CountExtraFields())
	}
	if !res0.HasExtraFields() {
		t.Fatal("table0 compatible from slice on t6 doesn't has extra fields")
	}

	// Table2
	var res2 *Table2

	_, err = Table2FromSlice(t0, true)
	if err == nil {
		t.Fatal("table2 compatible from slice success on t0")
	}

	_, err = Table2FromSlice(t1, true)
	if err == nil {
		t.Fatal("table2 compatible from slice success on t1")
	}

	i = 0
	res2, err = Table2FromSlice(t2, true)
	if err != nil {
		t.Fatal("table2 compatible from slice fail on t2")
	}
	if res2.CountExtraFields() != i {
		t.Fatal("table2 compatible from slice on t3, CountExtraFields not ", i, " but ", res2.CountExtraFields())
	}
	if res2.HasExtraFields() {
		t.Fatal("table2 compatible from slice on t3 has extra fields")
	}

	i++
	res2, err = Table2FromSlice(t3, true)
	if err != nil {
		t.Fatal("table2 compatible from slice fail on t3")
	}
	if res2.CountExtraFields() != i {
		t.Fatal("table3 compatible from slice on t3, CountExtraFields not ", i, " but ", res2.CountExtraFields())
	}
	if !res2.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t3 has extra fields")
	}

	i++
	res2, err = Table2FromSlice(t4, true)
	if err != nil {
		t.Fatal("table2 compatible from slice fail on t4", err)
	}
	if res2.CountExtraFields() != i {
		t.Fatal("table2 compatible from slice on t4, CountExtraFields not ", i, " but ", res2.CountExtraFields())
	}
	if !res2.HasExtraFields() {
		t.Fatal("table2 compatible from slice on t4 doesn't has extra fields")
	}

	i++
	res2, err = Table2FromSlice(t5, true)
	if err != nil {
		t.Fatal("table2 compatible from slice fail on t5")
	}
	if res2.CountExtraFields() != i {
		t.Fatal("table2 compatible from slice on t4, CountExtraFields not ", i, " but ", res2.CountExtraFields())
	}
	if !res2.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t5 doesn't has extra fields")
	}

	i++
	res2, err = Table2FromSlice(t6, true)
	if err != nil {
		t.Fatal("table2 compatible from slice fail on t6")
	}
	if res2.CountExtraFields() != i {
		t.Fatal("table2 compatible from slice on t4, CountExtraFields not ", i, " but ", res2.CountExtraFields())
	}
	if !res2.HasExtraFields() {
		t.Fatal("table2 compatible from slice on t6 doesn't has extra fields")
	}

	// Table3
	var res3 *Table3

	_, err = Table3FromSlice(t0, true)
	if err == nil {
		t.Fatal("table3 compatible from slice success on t0")
	}

	_, err = Table3FromSlice(t1, true)
	if err == nil {
		t.Fatal("table3 compatible from slice success on t1")
	}

	_, err = Table3FromSlice(t2, true)
	if err == nil {
		t.Fatal("table3 compatible from slice fail on t2")
	}

	i = 0
	res3, err = Table3FromSlice(t3, true)
	if err != nil {
		t.Fatal("table3 compatible from slice fail on t3")
	}
	if res3.CountExtraFields() != i {
		t.Fatal("table3 compatible from slice on t3, CountExtraFields not ", i, " but ", res3.CountExtraFields())
	}
	if res3.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t3 has extra fields")
	}

	i++
	res3, err = Table3FromSlice(t4, true)
	if err != nil {
		t.Fatal("table3 compatible from slice fail on t4", err)
	}
	if res3.CountExtraFields() != i {
		t.Fatal("table3 compatible from slice on t4, CountExtraFields not ", i, " but ", res3.CountExtraFields())
	}
	if !res3.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t4 doesn't has extra fields")
	}

	i++
	res3, err = Table3FromSlice(t5, true)
	if err != nil {
		t.Fatal("table3 compatible from slice fail on t5")
	}
	if res3.CountExtraFields() != i {
		t.Fatal("table3 compatible from slice on t4, CountExtraFields not ", i, " but ", res3.CountExtraFields())
	}
	if !res3.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t5 doesn't has extra fields")
	}

	i++
	res3, err = Table3FromSlice(t6, true)
	if err != nil {
		t.Fatal("table3 compatible from slice fail on t6")
	}
	if res3.CountExtraFields() != i {
		t.Fatal("table3 compatible from slice on t4, CountExtraFields not ", i, " but ", res3.CountExtraFields())
	}
	if !res3.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t6 doesn't has extra fields")
	}

	// Table6
	var res6 *Table6

	_, err = Table6FromSlice(t0, true)
	if err == nil {
		t.Fatal("table6 compatible from slice success on t0")
	}

	_, err = Table6FromSlice(t1, true)
	if err == nil {
		t.Fatal("table6 compatible from slice success on t1")
	}

	_, err = Table6FromSlice(t2, true)
	if err == nil {
		t.Fatal("table6 compatible from slice success on t2")
	}

	_, err = Table6FromSlice(t3, true)
	if err == nil {
		t.Fatal("table6 compatible from slice success on t3")
	}

	_, err = Table6FromSlice(t4, true)
	if err == nil {
		t.Fatal("table6 compatible from slice fail on t4")
	}

	_, err = Table6FromSlice(t5, true)
	if err == nil {
		t.Fatal("table6 compatible from slice success on t5")
	}

	i = 0
	res6, err = Table6FromSlice(t6, true)
	if err != nil {
		t.Fatal("table3 compatible from slice fail on t6")
	}
	if res6.CountExtraFields() != i {
		t.Fatal("table3 compatible from slice on t4, CountExtraFields not ", i, " but ", res3.CountExtraFields())
	}
	if res6.HasExtraFields() {
		t.Fatal("table3 compatible from slice on t6 doesn't has extra fields")
	}
}
