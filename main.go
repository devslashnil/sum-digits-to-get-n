package main

/*
	Комментарий к решению:

	Смысл алгоритма заключается в переборе всех вариантов с сохранением промежуточного результата в виде дерева решений,
	где вершины это сумма всех вершин предков, таким образом решение будет найдено когда сумма в ноде будет равна sum-digits-to-get-n, а
	глубина учитывать все цифры из условия.

	Т.к. алгоритм решения представляет собой полный перебор, это худший возможный вариант по асимптотике, также худший из
	аналогов по потреблению памяти т.к. требуется хранить всё дерево расчётов, но так как не стояло условие по
	оптимизации, то я выбрал самое читаемое и понятное решение, ибо уважаю время проверяющих.

*/

import (
	"fmt"
	"strconv"
	"strings"
)

// Node хранит число и промежуточную сумму, и указатели на пути решения на основе цифры с которой начинается подсчёт,
// а также комбинаций в виде цифры "склеенной" со следующей по порядку и их негативные варианты
type Node struct {
	// само число хранить не обязательно, при построении решения его можно определить из названия всех лепестков нод,
	// это сэкономит память, но потребует чуть больше кода, а оптимизация решения не требуется по условию
	Num    int16
	Sum    int16
	Parent *Node
	// нода с суммой цифры
	Digit *Node
	// нода с разницей цифры
	NegDigit *Node
	// нода с суммой числа, состоящего из цифры "склеенной" со следующей по порядку
	Pair *Node
	// нода с разницей числа, состоящего из цифры "склеенной" со следующей по порядку
	NegPair *Node
}

func main() {
	FillNode(&Node{}, 9)
}

func FillNode(node *Node, digit int16) {
	if digit <= 0 {
		if node.Sum == 200 {
			PrintSolution(*node)
		}
		return
	}
	pair := digit*10 + digit - 1
	node.Pair = CreateNode(node, pair)
	node.NegPair = CreateNode(node, -pair)
	node.Digit = CreateNode(node, digit)
	node.NegDigit = CreateNode(node, -digit)
	FillNode(node.Pair, digit-2)
	FillNode(node.NegPair, digit-2)
	FillNode(node.Digit, digit-1)
	FillNode(node.NegDigit, digit-1)
}

func CreateNode(parent *Node, digit int16) *Node {
	return &Node{
		digit,
		parent.Sum + digit,
		parent,
		nil,
		nil,
		nil,
		nil,
	}
}

// PrintSolution восстанавливает решения за несколько шагов просто итерирую вверх по дереву вплоть корня
func PrintSolution(node Node) {
	var nodes []Node
	if node.Num == -1 {
		nodes = append(nodes, Node{})
	}
	for node.Parent != nil {
		nodes = append(nodes, node)
		node = *node.Parent
	}
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(int(nodes[len(nodes)-1].Num)))
	for i := len(nodes) - 2; i >= 0; i-- {
		if nodes[i].Num >= 0 {
			sb.WriteByte('+')
		}
		sb.WriteString(strconv.Itoa(int(nodes[i].Num)))
	}
	fmt.Println(sb.String())
}
