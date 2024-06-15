package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcBspDecal struct {
	Pos               []vectorCoord `json:"pos"`
	DecalTextureIndex int16         `json:"decal_texture_index"`
	EntityIndex       uint16        `json:"entity_index"`
	ModelIndex        uint16        `json:"model_index"`
	LowPriority       bool          `json:"low_priority"`
}

type vectorCoord struct {
	Value float32 `json:"value"`
	Valid bool    `json:"valid"`
}

func ParseSvcBspDecal(reader *bitreader.Reader, demo *types.Demo) SvcBspDecal {
	svcBspDecal := SvcBspDecal{
		Pos:               readVectorCoords(reader),
		DecalTextureIndex: int16(reader.TryReadBits(9)),
	}
	if reader.TryReadBool() {
		svcBspDecal.EntityIndex = uint16(reader.TryReadBits(11))
		svcBspDecal.ModelIndex = uint16(reader.TryReadBits(11))
	}
	svcBspDecal.LowPriority = reader.TryReadBool()
	demo.Writer.TempAppendLine("\t\tPosition: %v", svcBspDecal.Pos)
	demo.Writer.TempAppendLine("\t\tDecal Texture Index: %d", svcBspDecal.DecalTextureIndex)
	demo.Writer.TempAppendLine("\t\tEntity Index: %d", svcBspDecal.EntityIndex)
	demo.Writer.TempAppendLine("\t\tModel Index: %d", svcBspDecal.ModelIndex)
	demo.Writer.TempAppendLine("\t\tLow Priority: %t", svcBspDecal.LowPriority)
	return svcBspDecal
}

func readVectorCoords(reader *bitreader.Reader) []vectorCoord {
	const COORD_INTEGER_BITS uint8 = 14
	const COORD_FRACTIONAL_BITS uint8 = 5
	const COORD_DENOMINATOR uint8 = 1 << COORD_FRACTIONAL_BITS
	const COORD_RESOLUTION float32 = 1.0 / float32(COORD_DENOMINATOR)
	readVectorCoord := func() float32 {
		value := float32(0)
		integer := reader.TryReadBits(1)
		fraction := reader.TryReadBits(1)
		if integer != 0 || fraction != 0 {
			sign := reader.TryReadBits(1)
			if integer != 0 {
				integer = reader.TryReadBits(uint64(COORD_INTEGER_BITS)) + 1
			}
			if fraction != 0 {
				fraction = reader.TryReadBits(uint64(COORD_FRACTIONAL_BITS))
			}
			value = float32(integer) + float32(fraction)*COORD_RESOLUTION
			if sign != 0 {
				value = -value
			}
		}
		return value
	}
	x := reader.TryReadBits(1)
	y := reader.TryReadBits(1)
	z := reader.TryReadBits(1)
	return []vectorCoord{{Value: readVectorCoord(), Valid: x != 0}, {Value: readVectorCoord(), Valid: y != 0}, {Value: readVectorCoord(), Valid: z != 0}}
}
