package functions

import "search-engine-extended/src/model"

func swap(num1 *float32, num2 *float32) (error){
	temp := *num1
	*num1 = *num2
	*num2 = temp

	return nil
}

//fungsi rekursif untuk membantu membentuk array dalam bentuk min-heap binary tree 
func heapify(arr []float32, n int, idxNode int) (error){
	lowest := idxNode
	left := 2*idxNode+1
	right := 2*idxNode+2

	//cari index dengan nilai terbesar
	if (left<n && arr[lowest]>arr[left]){
		lowest = left
	}

	if (right<n && arr[lowest]>arr[right]){
		lowest = right
	}

	//index dengan nilai terbesar bukan "root"
	if (lowest!=idxNode){	
		swap(&arr[idxNode],&arr[lowest])	//tuker
		heapify(arr,n,lowest)		//"rapihin" node yang ketuker
	}

	return nil
}

//sementara array of numerik dulu
func HeapSort(arr []float32) (error){
	n := len(arr)		//banyak elemen di arr
	
	//inisiasi awal kita bentuk array menjadi max-heap binary tree buat dapetin nilai paling kecil
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr,n,i)
	}

	//mulai susun nilai min di belakang
	for i := n-1; i>0; i--{
		swap(&arr[0],&arr[i])

		heapify(arr,i,0)
	}

	return nil
}

func swapAdapted(num1 *model.Page, num2 *model.Page) (error){
	temp := *num1
	*num1 = *num2
	*num2 = temp

	return nil
}


func heapifyAdapted(arr []model.Page, n int, idxNode int) (error){
	lowest := idxNode
	left := 2*idxNode+1
	right := 2*idxNode+2

	valueLowest := arr[lowest].Relevancy + arr[lowest].Importance*0.1

	//cari index dengan nilai terbesar
	if (left<n){
		valueLeft := arr[left].Relevancy + arr[left].Importance*0.1
		if (valueLowest>valueLeft){
			lowest = left
			valueLowest = valueLeft
		}
		
	}

	if (right<n){
		valueRight := arr[right].Relevancy + arr[right].Importance*0.1
		if (valueLowest>valueRight){
			lowest = right
			valueLowest = valueRight
		}
		
	}

	//index dengan nilai terbesar bukan "root"
	if (lowest!=idxNode){	
		swapAdapted(&arr[idxNode],&arr[lowest])	//tuker
		heapifyAdapted(arr,n,lowest)			//"rapihin" node yang ketuker
	}

	return nil
}

func HeapSortAdapted(arr []model.Page) error {
	n := len(arr)		//banyak elemen di arr
	
	//inisiasi awal kita bentuk array menjadi max-heap binary tree buat dapetin nilai paling kecil
	for i := n/2 - 1; i >= 0; i-- {
		heapifyAdapted(arr,n,i)
	}

	//mulai susun nilai min di belakang
	for i := n-1; i>0; i--{
		swapAdapted(&arr[0],&arr[i])

		heapifyAdapted(arr,i,0)
	}

	return nil
}
