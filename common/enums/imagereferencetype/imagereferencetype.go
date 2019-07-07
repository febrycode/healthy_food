package imagereferencetype

type ImageReferenceType int

const (
	USER ImageReferenceType = 1
	FOOD ImageReferenceType = 2
)

var imageReferenceTypeList = []ImageReferenceType{
	USER,
	FOOD,
}

func (enumVal ImageReferenceType) String() string {
	switch enumVal {
	case 1:
		return "User"
	case 2:
		return "Food"
	default:
		return "undefined"
	}
}

func (enumVal ImageReferenceType) IsValid() bool {
	return enumVal == 1 ||
		enumVal == 2
}

func GetAll() []ImageReferenceType {
	return imageReferenceTypeList
}
