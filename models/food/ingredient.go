// Package food
//
//	@English	ingredient, elements that make food, such as fruits and vegetables
//	@Chinese	食材, 制作食物的基本材料, 比如 姜葱蒜等调料, 蔬菜, 水果, 肉类, 米饭等主食
package food

type Ingredient struct {
	ID            string
	Name          string
	DisplayName   string
	AlternateName string
	ChineseName   string
	Description   string
	Type          []string
	Tag           []string
}
