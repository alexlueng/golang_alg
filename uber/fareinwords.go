package uber

// feature 4 convet the calculated fare to English words for our system to recognize

func ones(fare int) string {
	switch fare {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	}
	return ""
}

func teens(fare int) string {
	switch fare {
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	}
	return ""
}

func tens(fare int) string {
	switch fare {
	case 2:
		return "Twenty"
	case 3:
		return "Thirty"
	case 4:
		return "Forty"
	case 5:
		return "Fifty"
	case 6:
		return "Sixty"
	case 7:
		return "Seventy"
	case 8:
		return "Eighty"
	case 9:
		return "Ninety"
	}
	return ""
}

func two(fare int) string {
	if fare == 0 {
		return ""
	} else if fare < 10 {
		return ones(fare)
	} else if fare < 20 {
		return teens(fare)
	} else {
		tenner := fare / 10
		rest := fare - tenner*10
		if rest != 0 {
			return tens(tenner) + " " + ones(rest)
		} else {
			return tens(tenner)
		}
	}
}

func three(fare int) string {
	hundred := fare / 100
	rest := fare - hundred*100
	res := ""

	if hundred*rest != 0 {
		res = ones(hundred) + " Hundred " + two(rest)
	} else if (hundred == 0) && (rest != 0) {
		res = two(rest)
	} else if hundred != 0 && rest == 0 {
		res = ones(hundred) + " Hundred"
	}
	return res
}

func FareInWords(fare int) string {
	if fare == 0 {
		return "Zero"
	}
	billion := fare / 1000000000
	million := (fare - billion*1000000000) / 1000000
	thousand := (fare - billion*1000000000 - million*1000000) / 1000
	rest := fare - billion*1000000000 - million*100000 - thousand*1000
	result := ""
	if billion != 0 {
		return three(billion) + " Billion"
	}
	if million != 0 {
		if len(result) != 0 {
			result += " "
		}
		result += three(million) + " Million"
	}
	if thousand != 0 {
		if len(result) != 0 {
			result += " "
		}
		result += three(thousand) + " Thousand"
	}
	if rest != 0 {
		if len(result) != 0 {
			result += " "
		}
		result += three(rest)
	}
	return result
}
