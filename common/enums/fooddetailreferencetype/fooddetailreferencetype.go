package fooddetailreferencetype

type FoodDetailReferenceType int

const (
	BENEFIT      FoodDetailReferenceType = 1
	DISADVANTAGE FoodDetailReferenceType = 2
)

var foodDetailReferenceTypeList = []FoodDetailReferenceType{
	BENEFIT,
	DISADVANTAGE,
}

func (enumVal FoodDetailReferenceType) String() string {
	switch enumVal {
	case 1:
		return "Kelebihan"
	case 2:
		return "Efek Samping"
	default:
		return "undefined"
	}
}

func (enumVal FoodDetailReferenceType) IsValid() bool {
	return enumVal == 1 ||
		enumVal == 2
}

func GetAll() []FoodDetailReferenceType {
	return foodDetailReferenceTypeList
}
