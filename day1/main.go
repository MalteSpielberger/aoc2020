package main

import (
	"fmt"
	"strconv"
	"strings"
)

const values = "2010\n1856\n1905\n1786\n1557\n1995\n1830\n1971\n1909\n1500\n1806\n1846\n2003\n1839\n1943\n1977\n1537\n689\n1861\n1886\n1815\n1763\n1834\n1881\n1952\n1853\n1775\n1835\n1874\n1948\n1978\n347\n1672\n885\n1709\n1826\n1911\n1644\n1064\n1561\n1966\n1352\n1347\n1928\n1756\n615\n1513\n1932\n1968\n1762\n1842\n1475\n1921\n1716\n1533\n1975\n1924\n1850\n1456\n1783\n1587\n1913\n1908\n1502\n1993\n1635\n1691\n1706\n1871\n1857\n1915\n1604\n1618\n1902\n1860\n1648\n1933\n1999\n1960\n1389\n1858\n1793\n1609\n1484\n1735\n1535\n1891\n1879\n1517\n1766\n1926\n1668\n1495\n1585\n1831\n1308\n1767\n1479\n1638\n1600\n710\n1685\n1818\n1859\n1822\n1844\n1550\n1872\n1719\n1863\n1987\n199\n1840\n1817\n1752\n1612\n1983\n1838\n1504\n1997\n716\n1862\n1931\n1356\n1645\n1962\n1574\n1914\n1869\n1919\n1487\n1961\n1728\n1867\n1177\n1757\n1316\n1875\n1991\n1646\n700\n1972\n2004\n1577\n118\n1954\n1483\n1516\n2007\n1506\n1588\n1698\n1725\n2006\n179\n1849\n1894\n1695\n1399\n1726\n1658\n1920\n1825\n1837\n1878\n1591\n1611\n1409\n1553\n1705\n1845\n1718\n1732\n1639\n1885\n1929\n1887\n1787\n1541\n1946\n1391\n1884\n1938\n1496\n1720\n1669\n1965\n1967\n1890\n1743\n1889\n1970\n1866\n1912\n1785\n1998\n1708\n1810\n1939\n2005"

type Values struct {
	givenValues string
	nums []int
}

func (v *Values) ConvertValues() {
	values := strings.Split(v.givenValues, "\n")

	for _, r := range values {
		num, _ := strconv.Atoi(r)

		v.nums = append(v.nums, num)
	}
}

func (v *Values) Sum2() {
	for _, i := range v.nums {
		for _, j := range v.nums {
			if i+j == 2020 {
				fmt.Printf("The values are %v and %v!", i, j)
				fmt.Print("\n")
				fmt.Printf("The result of %v * %v = %v", i, j, i*j)
				return
			}
		}
	}
}

func (v *Values) Sum3() {
	for _, i := range v.nums {
		for _, j := range v.nums {
			for _, y := range v.nums {
				if i+j+y == 2020 {
					fmt.Printf("The values are %v, %v and %v!", i, j, y)
					fmt.Print("\n")
					fmt.Printf("The result of %v * %v * %v = %v", i, j, y, i*j*y)
					return
				}
			}
		}
	}
}

func main() {
	v := &Values{
		givenValues: values,
		nums: []int{},
	}

	v.ConvertValues()

	v.Sum3()
}