package main

import "fmt"

func main() {
	fmt.Println(minRecolors("WBBWWBBWBW", 7)) // Output: 3
	fmt.Println(minRecolors("WBWBBBW", 2))    // Output: 0
	fmt.Println(minRecolors("WWWW", 2))       // Output: 2
	fmt.Println(minRecolors("BBBB", 2))       // Output: 0
}

func minRecolors(blocks string, k int) int {
	count := 0
	// Count 'W' in the first window of size k
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			count++
		}
	}
	minCount := count // Store initial window count
	// Sliding window logic
	for i := k; i < len(blocks); i++ {
		// Remove the leftmost element from the previous window
		if blocks[i-k] == 'W' {
			count--
		}
		// Add the new element to the window
		if blocks[i] == 'W' {
			count++
		}
		// Update the minimum recoloring needed
		if count < minCount {
			minCount = count
		}
	}
	return minCount
}

/*

Example Walkthrough
Let's take blocks = "WBBWWBBWBW", k = 7.

Initial Window (k=7 first elements)
csharp
Copy
Edit
W B B W W B B W B W
[W B B W W B B] W B W  → Count of 'W' = 3
minOperations = 3 (need to change 3 'W' to 'B' to make k=7 continuous 'B')
Sliding Window Updates
Slide right (remove 'W', add 'W')
nginx
Copy
Edit
W B B W W B B W B W
  [B B W W B B W] B W  → Count of 'W' = 3 (no change)
Slide right (remove 'B', add 'B')
nginx
Copy
Edit
W B B W W B B W B W
    [B W W B B W B] W  → Count of 'W' = 3 (no change)
Slide right (remove 'B', add 'W')
nginx
Copy
Edit
W B B W W B B W B W
      [W W B B W B W]  → Count of 'W' = 3 (no change)
✅ Final Answer: 3 (Minimum operations to get 7 consecutive 'B').

🔹 1. What is the problem?
We are given a string blocks with 'W' (White) and 'B' (Black) characters. We need to find the minimum number of 'W' blocks to be changed to 'B' to get at least k consecutive 'B'.

🔹 2. Approach: Sliding Window Technique
We use a fixed-size sliding window of k to efficiently track the number of 'W' in each possible window of size k.

Why Sliding Window?
👉 Instead of checking all possible k-length substrings separately (O(n*k) time complexity), we use a sliding window to efficiently update the count in O(n).

🔹 3. Steps to Solve the Problem
Count the number of 'W' in the first k-length window.
This tells us how many changes are required for this window.
Slide the window one character at a time and update the 'W' count efficiently:
Remove the leftmost character from the previous window.
Add the new rightmost character in the new window.
Track the minimum number of changes required.
Return the minimum number of operations needed.
🔹 4. Step-by-Step Execution
Example:
go
Copy
Edit
blocks = "WBBWWBBWBW", k = 7
1️⃣ Initial window (first 7 characters):

css
Copy
Edit
W B B W W B B W B W
[W B B W W B B] W B W
Count of 'W' in [W B B W W B B] → 3
Store this as minCount = 3
2️⃣ Slide the window right:

Remove 'W' (leftmost) → Decrement count
Add 'W' (new character added at the right) → Increment count
New window:
css
Copy
Edit
W B B W W B B W B W
  [B B W W B B W] B W
Count remains 3 (no change).
3️⃣ Slide again:

css
Copy
Edit
W B B W W B B W B W
      [B W W B B W B] W
Still 3, so minCount = 3.
4️⃣ Final Output: 3 ✅

🔹 5. Time Complexity Analysis
Finding 'W' in the first k elements → O(k)
Sliding window (checking/removing/adding one element at a time) → O(n - k)
Total Complexity: O(n) ✅ (Efficient solution)
*/
