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


if __name__ == "__main__":
    print(find_minimum([1,2,3]))
    print(find_minimum([]))
    print(sum([1,2,3]))
    print(sum([]))
    print(average_audience_followers([2,5,77,9]))
    print(get_follower_prediction(10, "fitness", 10))
    print(get_influencer_score(2, 75))
    print(num_possible_orders(10))
