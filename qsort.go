// qsort
package gopractice

/**
 * QuickSort method, GO implementation
 *
 * @param []int elt -- the list to sort
 * @param int start -- where to start sorting
 * @param int end   -- where to end sorting
 *
 * @return void
 * @author Alex Lambson
 */
func QuickSort(elt []int, start int, end int) {
	if !(start < end) {
		return
	}
	divider_index := partition(elt, start, end)
	QuickSort(elt, start, divider_index-1)
	QuickSort(elt, divider_index+1, end)
	return
}

/**
 * swap method
 *
 * Swaps two ints
 * Usage example: elt[a], elt[b] = swap(elt[a], elt[b])
 * will swap the values of the two array indicies
 *
 * @param int a
 * @param int b
 *
 * @return int, int
 * @author Alex Lambson
 */
func swap(a int, b int) (int, int) {
	return b, a
}

/**
 * partition method
 *
 * used for QuickSort
 *
 * @param []int elt -- the array to sort and partition
 * @param int start -- where to start sorting
 * @param int end   -- where to stop sorting
 *
 * @return int      -- the index of where the next partition will be
 */
func partition(elt []int, start int, end int) (divider_index int) {
	pivot_value := elt[end]
	divider_index = start
	for current_index := start; current_index < end; current_index++ {
		if elt[current_index] <= pivot_value {
			elt[current_index], elt[divider_index] = swap(elt[current_index], elt[divider_index])
			divider_index++
		}
	}
	elt[divider_index], elt[end] = swap(elt[divider_index], elt[end])
	return
}
