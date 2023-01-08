# 算法框架

## 前缀和数组

一维数组

```java
class PrefixSum {
    // 前缀和数组
    private int[] prefix;
    
    /* 输入一个数组，构造前缀和 */
    public PrefixSum(int[] nums) {
        prefix = new int[nums.length + 1];
        // 计算 nums 的累加和
        for (int i = 1; i < prefix.length; i++) {
            prefix[i] = prefix[i - 1] + nums[i - 1];
        }
    }

    /* 查询闭区间 [i, j] 的累加和 */
    public int query(int i, int j) {
        return prefix[j + 1] - prefix[i];
    }
}

```

二维数组

```java
class NumMatrix {
    // 定义：preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
    private int[][] preSum;
    
    public NumMatrix(int[][] matrix) {
        int m = matrix.length, n = matrix[0].length;
        if (m == 0 || n == 0) return;
        // 构造前缀和矩阵
        preSum = new int[m + 1][n + 1];
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                // 计算每个矩阵 [0, 0, i, j] 的元素和
                preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i - 1][j - 1] - preSum[i-1][j-1];
            }
        }
    }
    
    // 计算子矩阵 [x1, y1, x2, y2] 的元素和
    public int sumRegion(int x1, int y1, int x2, int y2) {
        // 目标矩阵之和由四个相邻矩阵运算获得
        return preSum[x2+1][y2+1] - preSum[x1][y2+1] - preSum[x2+1][y1] + preSum[x1][y1];
    }
}
```

## 差分数组
```java
// 差分数组工具类
class Difference {
    // 差分数组
    private int[] diff;
    
    /* 输入一个初始数组，区间操作将在这个数组上进行 */
    public Difference(int[] nums) {
        assert nums.length > 0;
        diff = new int[nums.length];
        // 根据初始数组构造差分数组
        diff[0] = nums[0];
        for (int i = 1; i < nums.length; i++) {
            diff[i] = nums[i] - nums[i - 1];
        }
    }

    /* 给闭区间 [i, j] 增加 val（可以是负数）*/
    public void increment(int i, int j, int val) {
        diff[i] += val;
        if (j + 1 < diff.length) {
            diff[j + 1] -= val;
        }
    }

    /* 返回结果数组 */
    public int[] result() {
        int[] res = new int[diff.length];
        // 根据差分数组构造结果数组
        res[0] = diff[0];
        for (int i = 1; i < diff.length; i++) {
            res[i] = res[i - 1] + diff[i];
        }
        return res;
    }
}

```

## 滑动窗口
```c++
void slidingWindow(string s) {
    unordered_map<char, int> window;
    
    int left = 0, right = 0;
    while (right < s.size()) {
        // c 是将移入窗口的字符
        char c = s[right];
        // 增大窗口
        right++;
        // 进行窗口内数据的一系列更新
        ...

        /*** debug 输出的位置 ***/
        printf("window: [%d, %d)\n", left, right);
        /********************/
        
        // 判断左侧窗口是否要收缩
        while (window needs shrink) {
            // d 是将移出窗口的字符
            char d = s[left];
            // 缩小窗口
            left++;
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```

## Rabin-Karp 算法
```java
// Rabin-Karp 指纹字符串查找算法
int rabinKarp(String txt, String pat) {
    // 位数
    int L = pat.length();
    // 进制（只考虑 ASCII 编码）
    int R = 256;
    // 取一个比较大的素数作为求模的除数
    long Q = 1658598167;
    // R^(L - 1) 的结果
    long RL = 1;
    for (int i = 1; i <= L - 1; i++) {
        // 计算过程中不断求模，避免溢出
        RL = (RL * R) % Q;
    }
    // 计算模式串的哈希值，时间 O(L)
    long patHash = 0;
    for (int i = 0; i < pat.length(); i++) {
        patHash = (R * patHash + pat.charAt(i)) % Q;
    }
    
    // 滑动窗口中子字符串的哈希值
    long windowHash = 0;
    
    // 滑动窗口代码框架，时间 O(N)
    int left = 0, right = 0;
    while (right < txt.length()) {
        // 扩大窗口，移入字符
        windowHash = ((R * windowHash) % Q + txt.charAt(right)) % Q;
        right++;

        // 当子串的长度达到要求
        if (right - left == L) {
            // 根据哈希值判断是否匹配模式串
            if (windowHash == patHash) {
                // 当前窗口中的子串哈希值等于模式串的哈希值
                // 还需进一步确认窗口子串是否真的和模式串相同，避免哈希冲突
                if (pat.equals(txt.substring(left, right))) {
                    return left;
                }
            }
            // 缩小窗口，移出字符
            windowHash = (windowHash - (txt.charAt(left) * RL) % Q + Q) % Q;
            // X % Q == (X + Q) % Q 是一个模运算法则
            // 因为 windowHash - (txt[left] * RL) % Q 可能是负数
            // 所以额外再加一个 Q，保证 windowHash 不会是负数

            left++;
        }
    }
    // 没有找到模式串
    return -1;
}
```

## 排序

### 冒泡排序

```go
func sortArray(nums []int) []int {
    // 冒泡排序，比较交换，稳定算法，时间O(n^2), 空间O(1)
	// 每一轮遍历，将该轮最大值放到后面，同时小的往前冒
	// 从而形成后部是有序区
	// compare and swap
	for i:=0;i<len(nums);i++ {
		// 适当剪枝，len()-i到最后的部分都是有序区，避免再排
		for j:=1;j<len(nums)-i;j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}
```

### 选择排序

```go
func sortArray(nums []int) []int {
	// 选择排序，比较交换，不稳定算法，时间O(n^2)，空间O(1)
	// 每一轮遍历，该轮的最小值前挪，从而形成前面部分是有序区
	// compare and swap
	for i:=0;i<len(nums);i++ {
		// 剪枝前面部分，比较后面部分
		for j:=i+1;j<len(nums);j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
```

### 插入排序

```go
func sortArray(nums []int) []int {
	// 插入排序，比较交换，稳定算法，时间O(n^2)，空间O(1)
	// 0->len方向，每轮从后往前比较，相当于找到合适位置，插入进去
	// 数据规模小的时候，或基本有序，效率高
	n := len(nums)
	for i := 1; i < n; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp { //左边比右边大
			nums[j+1] = nums[j] //右移1位
			j--                 //扫描前一个数
		}
		nums[j+1] = tmp //添加到小于它的数的右边
	}
	return nums
}
```

### 希尔排序

```go
func sortArray(nums []int) []int {
	// 希尔排序，比较交换，不稳定算法，时间O(nlog2n)最坏O(n^2), 空间O(1)
	// 改进插入算法
	// 每一轮按照间隔插入排序，间隔依次减小，最后一次一定是1
	/*
	主要思想：
	设增量序列个数为k，则进行k轮排序。每一轮中，
	按照某个增量将数据分割成较小的若干组，
	每一组内部进行插入排序；各组排序完毕后，
	减小增量，进行下一轮的内部排序。
	*/
	gap := len(nums)/2
	for gap > 0 {
		for i:=gap;i<len(nums);i++ {
			j := i
			for j-gap >= 0 && nums[j-gap] > nums[j] {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
				j -= gap
			}
		}
		gap /= 2
	}
	return nums
}
```

### 归并排序

```go
// 递归实现归并算法
func sortArray(nums []int) []int {
	// 归并排序，基于比较，稳定算法，时间O(nlogn)，空间O(logn) | O(n)
	// 基于递归的归并-自上而下的合并，另有非递归法的归并(自下而上的合并)
	// 都需要开辟一个大小为n的数组中转
	// 将数组分为左右两部分，递归左右两块，最后合并，即归并
	// 如在一个合并中，将两块部分的元素，遍历取较小值填入结果集
	// 类似两个有序链表的合并，每次两两合并相邻的两个有序序列，直到整个序列有序
	n := len(nums)
	temp := make([]int, n)

	var merge func([]int, int, int, int)
	merge = func(nums []int, lo, mid, hi int) {
		for i := lo; i <= hi; i++ {
			temp[i] = nums[i]
		}
		i, j := lo, mid+1
		for p := lo; p <= hi; p++ {
			if i == mid+1 {
				// 左半边数组已全部被合并
				nums[p] = temp[j]
				j++
			} else if j == hi+1 {
				// 右半边数组已全部被合并
				nums[p] = temp[i]
				i++
			} else if temp[i] > temp[j] {
				nums[p] = temp[j]
				j++
			} else {
				nums[p] = temp[i]
				i++
			}
		}
	}

	var sort func([]int, int, int)
	sort = func(nums []int, lo, hi int) {
		if lo == hi {
			// 单个元素不用排序
			return
		}
		mid := lo + (hi-lo)/2
		sort(nums, lo, mid)
		sort(nums, mid+1, hi)
		merge(nums, lo, mid, hi)
	}

	sort(nums, 0, n-1)	
	return nums
}

// 非递归实现归并算法
func sortArray(nums []int) []int {
	// 归并排序-非递归实现，利用变量，自下而上的方式合并
	// 时间O(nlogn)，空间O(n)
	if len(nums) <= 1 {return nums}
	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		var l,r,i int
		// 通过遍历完成比较填入res中
		for l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
			i++
		}
		// 如果left或者right还有剩余元素，添加到结果集的尾部
		copy(res[i:], left[l:])
		copy(res[i+len(left)-l:], right[r:])
		return res
	}
	i := 1 //子序列大小初始1
	res := make([]int, 0)
	// i控制每次划分的序列长度
	for i < len(nums) {
		// j根据i值执行具体的合并
		j := 0
		// 按顺序两两合并，j用来定位起始点
		// 随着序列翻倍，每次两两合并的数组大小也翻倍
		for j < len(nums) {
			if j+2*i > len(nums) {
				res = merge(nums[j:j+i], nums[j+i:])
			} else {
				res = merge(nums[j:j+i], nums[j+i:j+2*i])
			}
			// 通过index控制每次将合并的数据填入nums中
			// 重填入的次数和合并及二叉树的高度相关
			index := j
			for _, v := range res {
				nums[index] = v
				index++
			}
			j = j + 2*i
		}
		i *= 2
	}
	return nums
}
```

### 快速排序

```go
func sortArray(nums []int) []int {
    // 快速排序，基于比较，不稳定算法，时间平均O(nlogn)，最坏O(n^2)，空间O(logn)
	// 分治思想，选主元，依次将剩余元素的小于主元放其左侧，大的放右侧
	// 然后取主元的前半部分和后半部分进行同样处理，直至各子序列剩余一个元素结束，排序完成
	// 注意：
	// 小规模数据(n<100)，由于快排用到递归，性能不如插排
	// 进行排序时，可定义阈值，小规模数据用插排，往后用快排
	// golang的sort包用到了快排
	// (小数，主元，大数)
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		// 递归终止条件
		if left > right {
			return nil
		}
		// 左右指针及主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			// 寻找小于主元的右边元素
			for i < j && nums[j] >= pivot {
				j--
			}
			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}
			// 交换i/j下标元素
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换元素
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)
		quick(nums, i+1, right)
		return nums
	}
	return quick(nums, 0, len(nums)-1)
}
```

### 堆排序

```go
func sortArray(nums []int) []int {
    // 堆排序-大根堆，升序排序，基于比较交换的不稳定算法，时间O(nlogn)，空间O(1)-迭代建堆
	// 遍历元素时间O(n)，堆化时间O(logn)，开始建堆次数多些，后面次数少 
	// 主要思路：
	// 1.建堆，从非叶子节点开始依次堆化，注意逆序，从下往上堆化
	// 建堆流程：父节点与子节点比较，子节点大则交换父子节点，父节点索引更新为子节点，循环操作
	// 2.尾部遍历操作，弹出元素，再次堆化
	// 弹出元素排序流程：从最后节点开始，交换头尾元素，由于弹出，end--，再次对剩余数组元素建堆，循环操作
	// 建堆函数，堆化
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		// 大顶堆堆化，堆顶值小一直下沉
		for {
			// 左孩子节点索引
			child := root*2 + 1
			// 越界跳出
			if child > end {
				return
			}
			// 比较左右孩子，取大值，否则child不用++
			if child < end && nums[child] <= nums[child+1] {
				child++
			}
			// 如果父节点已经大于左右孩子大值，已堆化
			if nums[root] > nums[child] {
				return
			}
			// 孩子节点大值上冒
			nums[root], nums[child] = nums[child], nums[root]
			// 更新父节点到子节点，继续往下比较，不断下沉
			root = child
		}
	}
	end := len(nums)-1
	// 从最后一个非叶子节点开始堆化
	for i:=end/2;i>=0;i-- {
		heapify(nums, i, end)
	}
	// 依次弹出元素，然后再堆化，相当于依次把最大值放入尾部
	for i:=end;i>=0;i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}
```

### 基数排序

```go
func radix_sort(li []int) {
    // 先以个位数的⼤小来对数据进⾏排序，接着以十位数的⼤小来对数据进⾏排序，接着以百位数的⼤小......
    // 排到最后，就是一组有序的元素了。不过，他在以某位数进行排序的时候，是⽤用“桶”来排序的。
    // 由于某位数(个位/⼗位....，不是整个数)的⼤小范围为0-9，所以我们需要10个桶，
    // 然后把具有相同 数值的数放进同⼀个桶⾥，之后再把桶里的数按照0号桶到9号桶的顺序取出来，
    // 这样一趟下来，按照某位数的排序就完成了。
    // 算法效率：时间复杂度: O(kn)，空间复杂度: O(n+k)，稳定

	// 一次遍历获取最大值
    max_num := li[0]
    for i := 0; i < len(li); i++ {
        if max_num < li[i] {
            max_num = li[i]
        }
    }
	// 根据最大值的位数确定分几轮基数排序，如234，需要3轮，9仅需要一轮排序
    for j := 0; j < len(strconv.Itoa(max_num)); j++ {
		// 1.每轮排序，先分桶，数据装桶
        bin := make([][]int, 10)
        for k := 0; k < len(li); k++ {
            n := li[k] / int(math.Pow(10, float64(j))) % 10
            bin[n] = append(bin[n], li[k])
        }
		// 2.每轮排序，装完桶后，依次遍历桶，重排数组
        m := 0
        for p := 0; p < len(bin); p++ {
            for q := 0; q < len(bin[p]); q++ {
                li[m] = bin[p][q]
                m++
            }
        }
    }
}
```

### 桶排序

```go
func sortArray(nums []int) []int {
    // 桶排序，基于哈希思想的外排稳定算法，空间换时间，时间O(n+k)
	// 相当于计数排序的改进版，服从均匀分布，先将数据分到有限数量的桶中，
	// 每个桶分别排序，最后将非空桶的数据拼接起来
	var bucket func(nums []int, bucketSize int) []int
	bucket = func(nums []int, bucketSize int) []int {
		if len(nums) < 2 {
			return nums
		}
		// 获取最大最小值
		minAndMax := func(nums []int) (min, max int) {
			minNum := math.MaxInt32
			maxNum := math.MinInt32
			for i:=0;i<len(nums);i++ {
				if nums[i] < minNum {
					minNum = nums[i]
				}
				if nums[i] > maxNum {
					maxNum = nums[i]
				}
			}
			return minNum, maxNum
		}
		min_, max_ := minAndMax(nums)
		// 定义桶
		// 构建计数桶
		bucketCount := (max_-min_)/bucketSize + 1
		buckets := make([][]int, bucketCount)
		for i:=0;i<bucketCount;i++ {
			buckets[i] = make([]int, 0)
		}
		// 装桶-排序过程
		for i:=0;i<len(nums);i++ {
			// 桶序号
			bucketNum := (nums[i]-min_) / bucketSize
			buckets[bucketNum] = append(buckets[bucketNum], nums[i])
		}
		// 桶中排序
		// 上述装桶完成，出桶填入元素组
		index := 0
		for _, bucket := range buckets {
			sort.Slice(bucket, func(i, j int) bool {
				return bucket[i] < bucket[j]
			})
			for _, num := range bucket {
				nums[index] = num
				index++
			}
		}
		return nums
	}
	// 定义桶中的数量
	var bucketSize int = 2
	return bucket(nums, bucketSize)
}
```

### 计数排序

```go
func sortArray(nums []int) []int {
    // 计数排序，基于哈希思想的稳定外排序算法，空间换时间，时间O(n)，空间O(n)
	// 数据量大时，空间占用大
	// 空间换时间，通过开辟额外数据空间存储索引号记录数组的值和数组额个数
	// 思路：
	// 1.找出待排序的数组的最大值和最小值
	// 2.创建数组存放各元素的出现次数，先于[min, max]之间
	// 3.统计数组值的个数
	// 4.反向填充数组，填充时注意,num[i]=j+min，
	// j-前面需要略过的数的个数，两个维度，依次递增的数j++，一个是重复的数的计数j-不变
	if len(nums) == 0 {
		return nums
	}
	// 获取最大最小值
	minAndMax := func(nums []int) (min,max int) {
		minNum := math.MaxInt32
		maxNum := math.MinInt32
		for i:=0;i<len(nums);i++ {
			if nums[i] < minNum {
				minNum = nums[i]
			}
			if nums[i] > maxNum {
				maxNum = nums[i]
			}
		}
		return minNum, maxNum
	}
	min_, max_ := minAndMax(nums)
	// 中转数组存放遍历元素
	// 空间只需要min-max
	tmpNums := make([]int, max_-min_+1)
	// 遍历原数组
	for i:=0;i<len(nums);i++ {
		tmpNums[nums[i]-min_]++
	}
	// 遍历中转数组填入原数组
	j := 0
	for i:=0;i<len(nums);i++ {
		// 如果对应数字cnt=0，说明可以计入下一位数字
		for tmpNums[j] == 0 {
			j++
		}
		// 填入数字
		nums[i] = j + min_
		// 填一个数字，对应数字cnt--
		tmpNums[j]--
	}
	return nums
}
```

## 二叉树

### 二叉搜索树 BST

定义：
1. 对于 BST 的每一个节点 node，左子树节点的值都比 node 的值要小，右子树节点的值都比 node 的值大。
2. 对于 BST 的每一个节点 node，它的左侧子树和右侧子树都是 BST。

定义一个二叉树节点：
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

#### 验证二叉树是否是二叉搜索树
```go
func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}

func traverse(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return traverse(root.Left, min, root) && traverse(root.Right, root, max)
}
```

#### 搜索值
```go
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
```

#### 插入节点
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
```

#### 删除节点
```go
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// found it and delete
		// 左右子树为空，则直接用右左子树替换此节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 左右子树都有节点，则将右子树最小值替换此节点(也可以用左子树最大值替换此节点，选一种实现即可)
		rightMin := findRightMinInBST(root.Right)
		// 删除rightMin节点
		root.Right = deleteNode(root.Right, rightMin.Val)
		// 再替换rightMin与root
		rightMin.Left = root.Left
		rightMin.Right = root.Right
		root = rightMin
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func findRightMinInBST(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
```

## 二叉堆/优先级队列

小顶堆 golang 实现（基于heap包）（大顶堆只用改写Less方法，或在读写元素时取相反数）：

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// heap.Pop(h)执行时，会先执行swap(0, len(h)-1)，将堆顶元素交换到数组末尾
// 所以这里Pop应返回数组末尾元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func main() {
	h := &IntHeap{}
	heap.Init(h)

	heap.Push(h, 3) // [3]
	heap.Push(h, 1) // [1, 3]
	heap.Push(h, 5) // [1, 3, 5]

	val := heap.Pop(h) // 1
}
```


## 动态规划

动态规划解题套路框架:

明确 base case -> 明确「状态」 -> 明确「选择」 -> 定义 dp 数组/函数的含义。

基本思路:
1. 确定 base case。
2. 确定「状态」，也就是原问题和子问题中会变化的变量。
3. 确定「选择」，也就是导致「状态」产生变化的行为。
4. 明确 dp 函数/数组的定义。

实现：
- 自顶向下的带备忘录的递归解法
- 自底向上的dp table递推解法

优化：
- 自顶向下，时间优化，观察状态偏移是否存在重复值，增加备忘录
- 自底向上，空间优化，关注dp数组递推关系，是否存在某些状态不变，进行降维

[四键键盘](https://labuladong.github.io/algo/3/28/94/)


### 背包问题

#### 0-1 背包问题

给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少？

```
输入：N = 3, W = 4, wt = [2, 1, 3], val = [4, 2, 3]

输出：6
解释：选择前两件物品装进背包，总重量 3 小于 W，可以获得最大价值 6
```

思路：
```
// dp[i][w] 的定义: 对于前 i 个物品，当前背包的容量为 w，这种情况下可以装的最大价值是 dp[i][w]
int[][] dp[N+1][W+1]
dp[0][..] = 0
dp[..][0] = 0

for i in [1..N]:
    for w in [1..W]:
        dp[i][w] = max(
            把物品 i 装进背包,
            不把物品 i 装进背包
        )
return dp[N][W]
```

实现：
```go
func knapsack(n, w int, wt []int, val []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] < 0 {
				// 背包容量不足，只能选择不装入背包
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxInt(
					dp[i-1][j],                  // 不装入背包
					dp[i-1][j-wt[i-1]]+val[i-1], // 装入背包
				)
			}
		}
	}
	return dp[n][w]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 完全背包问题

有一个背包，最大容量为 amount，有一系列物品 items，每个物品的重量为 items[i]，每个物品的数量无限。请问有多少种方法，能够把背包恰好装满？

思路：
```
// dp[i][j] 的定义：若只使用前 i 个物品（可以重复使用），当背包容量为 j 时，有 dp[i][j] 种方法可以装满背包
int dp[N+1][amount+1]
dp[0][..] = 0
dp[..][0] = 1

for i in [1..N]:
    for j in [1..amount]:
        把物品 i 装进背包,
        不把物品 i 装进背包
return dp[N][amount]
```

实现：
```go
func knapsack(amount int, items []int) int {
	n := len(items)
	// dp[i][j] 的定义：若只使用前 i 个物品（可以重复使用），当背包容量为 j 时，有 dp[i][j] 种方法可以装满背包
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-items[i-1] < 0 {
				// 容量不够
				dp[i][j] = dp[i-1][j]
			} else {
				// 不装入+装入的方法的和
				// 注意，完全背包可以装入重复的值，所以这里转入的方法，是dp[i][j-items[i-1]]
				dp[i][j] = dp[i-1][j] + dp[i][j-items[i-1]]
			}
		}
	}

	return dp[n][amount]
}
```

## 贪心

[会议室问题](https://labuladong.github.io/algo/3/29/100/)

## 回溯

框架：

```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```

### N 皇后问题

给你一个 N×N 的棋盘，让你放置 N 个皇后，使得它们不能互相攻击。皇后可以攻击同一行、同一列、左上左下右上右下四个方向的任意单位。

```c++
vector<vector<string>> res;

/* 输入棋盘边长 n，返回所有合法的放置 */
vector<vector<string>> solveNQueens(int n) {
    // vector<string> 代表一个棋盘
    // '.' 表示空，'Q' 表示皇后，初始化空棋盘
    vector<string> board(n, string(n, '.'));
    backtrack(board, 0);
    return res;
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
void backtrack(vector<string>& board, int row) {
    // 触发结束条件
    if (row == board.size()) {
        res.push_back(board);
        return;
    }
    
    int n = board[row].size();
    for (int col = 0; col < n; col++) {
        // 排除不合法选择
        if (!isValid(board, row, col)) {
            continue;
        }
        // 做选择
        board[row][col] = 'Q';
        // 进入下一行决策
        backtrack(board, row + 1);
        // 撤销选择
        board[row][col] = '.';
    }
}

/* 是否可以在 board[row][col] 放置皇后？ */
bool isValid(vector<string>& board, int row, int col) {
    int n = board.size();
    // 检查列是否有皇后互相冲突
    for (int i = 0; i <= row; i++) {
        if (board[i][col] == 'Q')
            return false;
    }
    // 检查右上方是否有皇后互相冲突
    for (int i = row - 1, j = col + 1; 
            i >= 0 && j < n; i--, j++) {
        if (board[i][j] == 'Q')
            return false;
    }
    // 检查左上方是否有皇后互相冲突
    for (int i = row - 1, j = col - 1;
            i >= 0 && j >= 0; i--, j--) {
        if (board[i][j] == 'Q')
            return false;
    }
    return true;
}
```

### 排列组合与子集

排列是一类问题，组合和子集是一类问题。

#### 元素无重不可复选

即 nums 中的元素都是唯一的，每个元素最多只能被使用一次，backtrack 核心代码如下：

```java
/* 组合/子集问题回溯算法框架 */
void backtrack(int[] nums, int start) {
    // 回溯算法标准框架
    for (int i = start; i < nums.length; i++) {
        // 做选择
        track.addLast(nums[i]);
        // 注意参数
        backtrack(nums, i + 1);
        // 撤销选择
        track.removeLast();
    }
}

/* 排列问题回溯算法框架 */
void backtrack(int[] nums) {
    for (int i = 0; i < nums.length; i++) {
        // 剪枝逻辑
        if (used[i]) {
            continue;
        }
        // 做选择
        used[i] = true;
        track.addLast(nums[i]);

        backtrack(nums);
        // 撤销选择
        track.removeLast();
        used[i] = false;
    }
}
```

#### 元素可重不可复选

即 nums 中的元素可以存在重复，每个元素最多只能被使用一次，其关键在于排序和剪枝，backtrack 核心代码如下：

```java
Arrays.sort(nums);
/* 组合/子集问题回溯算法框架 */
void backtrack(int[] nums, int start) {
    // 回溯算法标准框架
    for (int i = start; i < nums.length; i++) {
        // 剪枝逻辑，跳过值相同的相邻树枝
        if (i > start && nums[i] == nums[i - 1]) {
            continue;
        }
        // 做选择
        track.addLast(nums[i]);
        // 注意参数
        backtrack(nums, i + 1);
        // 撤销选择
        track.removeLast();
    }
}


Arrays.sort(nums);
/* 排列问题回溯算法框架 */
void backtrack(int[] nums) {
    for (int i = 0; i < nums.length; i++) {
        // 剪枝逻辑
        if (used[i]) {
            continue;
        }
        // 剪枝逻辑，固定相同的元素在排列中的相对位置
        if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1]) {
            continue;
        }
        // 做选择
        used[i] = true;
        track.addLast(nums[i]);

        backtrack(nums);
        // 撤销选择
        track.removeLast();
        used[i] = false;
    }
}
```

#### 元素无重可复选

即 nums 中的元素都是唯一的，每个元素可以被使用若干次，只要删掉去重逻辑即可，backtrack 核心代码如下：

```java
/* 组合/子集问题回溯算法框架 */
void backtrack(int[] nums, int start) {
    // 回溯算法标准框架
    for (int i = start; i < nums.length; i++) {
        // 做选择
        track.addLast(nums[i]);
        // 注意参数
        backtrack(nums, i);
        // 撤销选择
        track.removeLast();
    }
}


/* 排列问题回溯算法框架 */
void backtrack(int[] nums) {
    for (int i = 0; i < nums.length; i++) {
        // 做选择
        track.addLast(nums[i]);
        backtrack(nums);
        // 撤销选择
        track.removeLast();
    }
}
```

### 有限状态机与 KMP 字符匹配

KMP 算法是在 txt 中查找子串 pat，如果存在，返回这个子串的起始索引，否则返回 -1。

思路：
```
空间换时间，先依据pat构建dp数组，再对txt搜索。

pat串长度为m，匹配字符的个数有256个（ascii码）

dp[j][c] = next
0 <= j < m，代表当前的状态，即当前匹配到的pat串的字符索引
0 <= c < 256，代表遇到的字符（ASCII 码）
0 <= next <= M，代表下一个状态

比如，对于 pat = "ABABC"
dp[4]['A'] = 3 表示：
当前是状态 4，如果遇到字符 A，
pat 应该转移到状态 3

dp[1]['B'] = 2 表示：
当前是状态 1，如果遇到字符 B，
pat 应该转移到状态 2
```

实现：

```go
type kmp struct {
	dp  [][]int
	pat string
}

func construct(pat string) kmp {
	m, n := len(pat), 256 // 256个ascii码
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	k := kmp{dp, pat}

	// base case
	dp[0][int(pat[0])] = 1
	// 影子状态初始为0
	x := 0
	for i := 1; i < m; i++ {
		// 遍历256个ascii码
		for j := 0; j < n; j++ {
			if int(pat[i]) == j {
				dp[i][j] = i + 1
			} else {
				dp[i][j] = dp[x][j]
			}
		}
		// 更新x为上一个与pat[i]有公共前缀的下标位置
		x = dp[x][int(pat[i])]
	}

	return k
}

func (k kmp) search(txt string) int {
	m, n := len(k.pat), len(txt)
	j := 0
	for i := 0; i < n; i++ {
		j = k.dp[j][int(txt[i])]
		if j == m {
			return i - j + 1
		}
	}
	return -1
}
```

## BFS

### 最短距离

从一个图中找到从起点 start 到终点 target 的最近距离。

算法框架：

```java
// 计算从起点 start 到终点 target 的最近距离
int BFS(Node start, Node target) {
    Queue<Node> q; // 核心数据结构
    Set<Node> visited; // 避免走回头路
    
    q.offer(start); // 将起点加入队列
    visited.add(start);
    int step = 0; // 记录扩散的步数

    while (q not empty) {
        int sz = q.size();
        /* 将当前队列中的所有节点向四周扩散 */
        for (int i = 0; i < sz; i++) {
            Node cur = q.poll();
            /* 划重点：这里判断是否到达终点 */
            if (cur is target)
                return step;
            /* 将 cur 的相邻节点加入队列 */
            for (Node x : cur.adj()) {
                if (x not in visited) {
                    q.offer(x);
                    visited.add(x);
                }
            }
        }
        /* 划重点：更新步数在这里 */
        step++;
    }
}
```

## 数学公式

### 最大公约数

```
func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
```

### 最小公倍数

```
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
```

## 系列问题

### nSum 问题

本质是排序+双指针问题。

```go
func nSumTarget(nums []int, n, start, target int) [][]int {
	ans := [][]int{}
	if n < 2 || len(nums) < n {
		return ans
	}
	if n == 2 {
		// 2Sum 是 base case
		lo, hi := start, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			l, r := nums[lo], nums[hi]
			if sum < target {
				// 过滤重复值
				for ;lo < hi && nums[lo] == l; lo++ {
				}
			} else if sum > target {
				// 过滤重复值
				for ;lo < hi && nums[hi] == r; hi-- {
				}
			} else {
				ans = append(ans, []int{l, r})
				// 过滤重复值
				for ;lo < hi && nums[lo] == l; lo++ {
				}
				// 过滤重复值
				for ;lo < hi && nums[hi] == r; hi-- {
				}
			}
		}
	} else {
		// n > 2 时，递归计算 (n-1)Sum 的结果
		for i := start; i < len(nums); i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				// (n-1)Sum 加上 nums[i] 就是 nSum
				arr = append(arr, nums[i])
				ans = append(ans, arr)
			}
			// 跳过重复值
			for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
			}
		}
	}
	return ans
}
```

### 岛屿问题

二维数组 grid，其中只包含 0 或者 1，0 代表海水，1 代表陆地，且假设该矩阵四周都是被海水包围着的。连成片的陆地形成岛屿，计算这个矩阵 grid 中岛屿的个数。

dfs遍历+Flood Fill算法。遍历过程修改岛屿为海水，可以避免使用visited数组，节省空间。

```go
func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, row, col int) {
	m, n := len(grid), len(grid[0])
	if row >= m || row < 0 || col >= n || col < 0 {
		return
	}

	if grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'

	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)
}
```

### 股票问题

输入股票价格数组 prices，你最多进行 max_k 次交易，每次交易需要额外消耗 fee 的手续费，而且每次交易之后需要经过 cooldown 天的冷冻期才能进行下一次交易，请你计算并返回可以获得的最大利润。

```go
// dp[i][k][0 or 1]
// 0 <= i <= n-1, 1 <= k <= K
// n 为天数，K 为交易数的上限，0 和 1 代表是否持有股票
// 此问题共有 n*K*2 种状态

// base case：
// dp[-1][...][0] = dp[...][0][0] = 0
// dp[-1][...][1] = dp[...][0][1] = -infifity

// 状态转移
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
func maxProfit(prices []int, k, cooldown, fee int) int {
	n := len(prices)
	// 确定k的情况下，如果k>n/2，即无论如何都用不完k的额度，等于不限制k
	if k > n/2 {
		return maxProfitWithInfinityK(prices, cooldown, fee)
	}
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			dp[i][0][0] = 0
			dp[i][0][1] = math.MinInt
		}
	}

	for i := 0; i < n; i++ {
		for j := k; i > 0; j-- {
			// base case 1
			if i-1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i] - fee
				continue
			}
			// base case 2
			if i-cooldown-1 < 0 {
				dp[i][j][0] = maxInt(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = maxInt(dp[i-1][j][1], -prices[i]-fee)
				continue
			}
			dp[i][j][0] = maxInt(dp[i-1][j][0], dp[i-1][j][1])
			// 同时考虑 cooldown 和 fee
			dp[i][j][1] = maxInt(dp[i-1][j][1], dp[i-cooldown-1][j-1][0]-prices[i]-fee)
		}
	}
	return dp[n-1][k][0]
}

func maxProfitWithInfinityK(prices []int, cooldown, fee int) int {
	// 相当于k正无穷，k和k-1相同
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		// base case 1
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		// base case 2
		if i-cooldown-1 < 0 {
			dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = maxInt(dp[i-1][1], -prices[i]-fee)
			continue
		}
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		// 同时考虑 cooldown 和 fee
		dp[i][1] = maxInt(dp[i-1][1], dp[i-cooldown-1][0]-prices[i]-fee)
	}
	return dp[n-1][0]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```