package calculate

func ComputationalPages(size int, pageSize int) int {
	if size < pageSize {
		return 1
	}
	if size%pageSize == 0 {
		return size % pageSize
	} else {
		return size%pageSize + 1
	}
}
