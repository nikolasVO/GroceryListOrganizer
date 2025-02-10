package service

import (
	"sort"
	"strings"
)

// productOrder задаёт порядок расположения продуктов в магазине.
// Чем меньше число – тем раньше продукт должен идти в отсортированном списке.
var productOrder = map[string]int{
	"хлеб":               1,
	"молоко":             2,
	"яйца":               3,
	"сыр":                4,
	"масло":              5,
	"фрукты":             6,
	"овощи":              7,
	"мясо":               8,
	"рыба":               9,
	"специи":             10,
	"сладости":           11,
	"напитки":            12,
	"уходовая косметика": 13,
	"снеки":              14,
	"алкгольно":          15,
}

// normalizeProduct приводит название продукта к каноническому виду.
// Если слово найдено в словаре синонимов, возвращается каноническое название.
func normalizeProduct(product string) string {
	normalized := strings.ToLower(strings.TrimSpace(product))
	if canon, ok := SynonymsMap[normalized]; ok {
		return canon
	}
	return normalized
}

// OrganizeProducts нормализует и сортирует список продуктов согласно productOrder.
// Продукты, отсутствующие в productOrder, перемещаются в конец.
func OrganizeProducts(products []string) []string {
	normalizedProducts := make([]string, len(products))
	for i, p := range products {
		normalizedProducts[i] = normalizeProduct(p)
	}

	sorted := make([]string, len(normalizedProducts))
	copy(sorted, normalizedProducts)
	sort.SliceStable(sorted, func(i, j int) bool {
		orderI, okI := productOrder[sorted[i]]
		orderJ, okJ := productOrder[sorted[j]]
		if !okI {
			orderI = 100
		}
		if !okJ {
			orderJ = 100
		}
		return orderI < orderJ
	})
	return sorted
}
