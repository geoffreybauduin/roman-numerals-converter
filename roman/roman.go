//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package roman

var sorted_data []int = []int{
    1000000, 500000, 100000, 50000, 10000, 5000, 1000, 500, 100, 50, 10, 5, 1,
}

var sorted_substract []int = []int{
    1000000, 100000, 10000, 1000, 100, 10, 1,
}

var substract_data map[int]string = map[int]string{
    1000000: "M̅",
    100000: "C̅",
    10000: "X̅",
    1000: "I̅",
    100: "C",
    10: "X",
    1: "I",
}

var data map[int]string = map[int]string{
    1000000: "M̅",
    500000: "D̅",
    100000: "C̅",
    50000: "L̅",
    10000: "X̅",
    5000: "V̅",
    1000: "M",
    500: "D",
    100: "C",
    50: "L",
    10: "X",
    5: "V",
    1: "I",
}

func get_order (nbr int) (int) {
    for _, value := range sorted_substract {
        if nbr >= value {
            return (nbr / value) * value
        }
    }
    return nbr
}

func check_direct (nbr int) (bool, int, string) {
    for _, value := range sorted_data {
        for _, cmp_value := range sorted_substract {
            if value - cmp_value == nbr {
                return true, 0, substract_data[cmp_value] + data[value]
            }
        }
    }
    return false, nbr, ""
}

func check_direct_order_depth (nbr, order, it int, current_roman string) (bool, int, string) {
    order_depth := get_order(nbr - order)
    direct, _, roman_depth := check_direct(order + order_depth)
    if direct {
        return check_direct_order_depth(nbr, order + order_depth, it + 1, roman_depth)
    } else if it > 0 {
        return true, nbr - order, current_roman       
    } else {
        return false, nbr, current_roman
    }
}

func check_direct_order (nbr int) (bool, int, string) {
    order := get_order(nbr)
    direct, _, roman := check_direct(order)
    if direct {
        in_depth, nbr, roman_depth := check_direct_order_depth(nbr, order, 0, roman)
        if in_depth {
            return true, nbr, roman_depth
        }
        return true, nbr - order, roman
    }
    return false, nbr, ""
}

func process (nbr int) (int, string) {
    direct, nbr, roman := check_direct(nbr)
    if direct {
        return nbr, roman
    }
    direct, nbr, roman = check_direct_order(nbr)
    if direct {
        return nbr, roman
    }
    for _, value := range sorted_data {
        if nbr >= value {
            return nbr - value, data[value]
        }
    }
    return 0, ""    
}

func IntToRoman (nbr int) string {
    roman_nbr := ""
    var roman string
    for nbr > 0 {
        nbr, roman = process(nbr)
        roman_nbr = roman_nbr + roman 
    }
    return roman_nbr
}