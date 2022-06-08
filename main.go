package main

import (
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/mkorman9/flatbuffers/monsters"
)

func main() {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("Chochoł")

	monsters.MonsterStart(builder)
	monsters.MonsterAddPos(builder, monsters.CreateVec3(builder, 10, 10, 10))
	monsters.MonsterAddHp(builder, 100)
	monsters.MonsterAddName(builder, name)
	monsters.MonsterAddColor(builder, monsters.ColorRed)
	monsters.MonsterAddEquippedType(builder, monsters.EquipmentNONE)
	builder.Finish(monsters.MonsterEnd(builder))
	buf := builder.FinishedBytes()

	monster := monsters.GetRootAsMonster(buf, 0)
	fmt.Println(string(monster.Name()))
	fmt.Println(monster.Hp())
}
