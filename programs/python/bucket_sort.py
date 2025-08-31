# use buckets to sort
# Sort on less signifacant but and go to left
# tow digits need to buckets
# https://www.youtube.com/watch?v=Y95a-8oNqps


def bucket_sort(nums):
    bucket = [[] for _ in range(10)]
    iter = len(str(nums[0]))

    for it in range(iter):
      for i in nums:
        if it == 0:
            l1 = i
        else:
            l1 = int(i/(10 ** it))
        l1 = l1%10
        bucket[l1].append(i)
      nums = []
      for i in bucket:
          for j in i:
              if j:
                  nums.append(j)
      bucket = [[] for _ in range(10)]
      print(it, " ", nums)

    print("sorted final", nums)

if __name__ ==  "__main__":
    bucket_sort([10,22,12,46,32,56,21])
    bucket_sort([100,328,124,110,876])
