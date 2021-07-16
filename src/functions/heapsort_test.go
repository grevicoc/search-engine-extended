package functions

import (
	"testing"
	"reflect"
)

func TestHitungNormal1(t *testing.T) {
    rawArray := []float32{0.27,11.1,5.4,6.02,3.01}
	resultArray := []float32{11.1,6.02,5.4,3.01,0.27}
	t.Logf("Urutan Awal: %v", rawArray)

	HeapSort(rawArray)

    if !reflect.DeepEqual(rawArray,resultArray) {
        t.Errorf(`
		Urutan Akhir: %v
		Urutan Seharusnya: %v
		`, rawArray,resultArray)
    }
}

func TestHitungNormal2(t *testing.T) {
    rawArray := []float32{0.11,4.03,5.4,6.02,7.190003,9.1}
	resultArray := []float32{9.1,7.190003,6.02,5.4,4.03,0.11}
	t.Logf("Urutan Awal: %v", rawArray)

	HeapSort(rawArray)

    if !reflect.DeepEqual(rawArray,resultArray) {
        t.Errorf(`
		Urutan Akhir: %v
		Urutan Seharusnya: %v
		`, rawArray,resultArray)
    }
}