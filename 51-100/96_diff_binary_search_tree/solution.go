package diff_binary_search_tree

// 算是暴力求解吧
// 遍历1····n，每一个值当一个根节点，然后将所有能变成搜索的方法求和
// 时间复杂度：O(n * n)
// 空间复杂度：O(n)
func numTrees(n int) int {
	numList := make([]int, n+1)
	numList[0] = 1
	numList[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 左右子树个数的乘积
			numList[i] += numList[j-1] * numList[i-j]
		}
	}
	return numList[n]
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
//  c0 = 1
// cn+1 = (2(2n+1)/n+2) * cn
//  其实拆分题，就是求 1····n的排列组合个数
// https://baike.baidu.com/item/catalan/7605685?fr=aladdin
func numTrees1(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}
