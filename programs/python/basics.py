import math

def find_minimum(nums):
    if len(nums) ==0:
        return None
    minn = float("inf")
    for i in nums:
        if minn > i:
            minn=i
    return minn

def sum(nums):
    if len(nums)==0:
        return 0
    s = 0
    for i in nums:
        s += i
    return s

def average_audience_followers(nums):
    n = len(nums)
    sum = 0
    for i in nums:
        sum += i
    avg = sum/n
    return avg * (n ** 1.2)

# exponentiation
def get_follower_prediction(follower_count, influencer_type, num_months):
    factor = 2
    if influencer_type == "fitness":
        factor = 4
    elif influencer_type == "cosmetic":
        factor = 6

    return follower_count * ( factor ** num_months )

# logarithm
def get_influencer_score(num_followers, average_engagement_percentage):
    return average_engagement_percentage * math.log(num_followers, 2)

# factorial
def num_possible_orders(num_posts):
    p = 1
    for i in range(1,num_posts + 1):
        p = p * i
    return p


# Big O
# O(1) < O(log(n)) < O(n) < O(n**2) < O(n**3) < O(2**n) < O(n!)

def find_max(nums):
    maxx = -float("inf")
    for num in nums:
        if maxx > num:
            maxx = num
    return maxx

# O(n**2)
def does_name_exist(first_names,  last_names , full_name):
    for fname in first_names:
        for lname in last_names:
            if f'{fname} {lname}' == full_name:
                return True
    return False

# O(nm)
def get_avg_brand_followers(all_handles, brand_name):
    cnt = 0
    for l1 in all_handles:
        for l2 in l1:
            if "cosmo" in l2:
                cnt++
    lists = len(all_handles)

    return cnt/lists

# O(1) -> fetching an item in dictionary (or array)


# O(log(n)) - Binary search - nums needs to be sorted
def binary_search(nums, val):
    low, high = 0, len(nums)-1

    while low <= high:
        mid = (low + high) // 2
        if val == nums[mid]:
            return mid
        elif val > nums[mid]:
            low = mid + 1
        else:
            high = mid -1
    return -1

if __name__ == "__main__":
    print(find_minimum([1,2,3]))
    print(find_minimum([]))
    print(sum([1,2,3]))
    print(sum([]))
    print(average_audience_followers([2,5,77,9]))
    print(get_follower_prediction(10, "fitness", 10))
    print(get_influencer_score(2, 75))
    print(num_possible_orders(10))
