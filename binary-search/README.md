# Binary search

**Binary Search** is a simple algorithm to find an item in a **sorted array,** and it’s usually referenced as a code sample to study when learning a new programming language.

### Explanation

Let's guess the sorted array is given:

``` 
[0, 1, 2, 3, 8, 11, 21]
```
And we wont to find the number **8**

**Iteration 1:**

low index is **0**

high index is **6**

median means **(low + high) / 2** ang equals **3**

**array[median]** is **3**

``` 
l=0      m=3        h=6
[0, 1, 2, 3, 8, 11, 21]
```
because of **array[median] < 8** we take **2nd half** for next iteration and set **low** as **median + 1**

**Iteration 2:**

low index is **4**

high index is **6**

median means **(low + high) / 2** ang equals **5**

**array[median]** is **8**


``` 
            l=4     h=6
[0, 1, 2, 3, 8, 11, 21]
                m=5     
```
because of **array[median] > 8** we take **1st half** for next iteration and set **high** as **median - 1**

**Iteration 3:**

low index is **4**

high index is **4**

median means **(low + high) / 2** ang equals **4**

``` 
            l=4        
            h=4
[0, 1, 2, 3, 8, 11, 21]
            m=4     
```

because of **low = high** we break the loop and check the result

**array[low] = 8**

Hooray! We found it.  

### Time Complexity

We found the number for 3 iterations.

At **1st** iteration length of array is **n**

At **2nd** iteration length of array is **n/2**

At **3rd** iteration length of array is **(n/2)/2**

Therefore, after iteration **k** length of array is **n⁄2k**

Also, we know that after **k** divisions, the length of array becomes **1**

That means length of array = **n⁄2k = 1** and **n = 2k**

Every iteration we divide **n** by **2** so we can calculate the complexity using the logarithm to the base **2** 

``` 
=> log2 (n) = log2 (2k)
=> log2 (n) = k log2 (2)
as (loga (a) = 1)
=> k = log2 (n)
```
Using **Big O** notation we can represent this complexity as **O(log n)**


