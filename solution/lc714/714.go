package lc714

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	lowest := prices[0]
	highest := prices[0]
	sum := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if (highest-price) > fee && (highest-lowest) > fee {
			// sell
			sum += highest - lowest - fee
			lowest = price
			highest = price
		} else if price < lowest {
			lowest = price
			highest = price
		} else if price > highest {
			highest = price
		}
	}

	if (highest - lowest) > fee {
		sum += highest - lowest - fee
	}

	return sum
}
