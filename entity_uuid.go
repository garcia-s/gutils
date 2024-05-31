package utils

type entityUUID struct {
	value string
}

func New() entityUUID {
	return entityUUID{
		value: "asdsasdsa",
	}
}

func Reconsitute(value string) entityUUID {
	return entityUUID{
		value: value,
	}
}
