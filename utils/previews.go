package utils

import (
	"encoding/json"
	"fmt"
)

// Просмотр в ыекштпку структуру
// На вход подается любая структура
// На выходе JSON в стринге
// Полезна для контроля прихода структуры
func PreviewStruct(v any) {
	m, _ := json.Marshal(v)
	fmt.Println(string(m))
}
