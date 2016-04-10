//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package roman

var sorted_data []int = []int{
    1000, 500, 100, 50, 10, 5, 1,
}

var sorted_substract []int = []int{
    1000, 100, 10, 1,
}

var data map[int]string = map[int]string{
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
                return true, 0, data[cmp_value] + data[value]
            }
        }
    }
    return false, nbr, ""
}

func process (nbr int) (int, string) {
    direct, nbr, roman := check_direct(nbr)
    if direct {
        return nbr, roman
    }
    order := get_order(nbr)
    direct, _, roman = check_direct(order)
    if direct {
        return nbr - order, roman
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