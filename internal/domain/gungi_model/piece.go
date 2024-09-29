package gungi_model

type PieceType uint

const (
	Marshal           PieceType = iota // 帥
	General                            // 大
	LieutenantGeneral                  // 中
	MajorGeneral                       // 小
	Samurai                            // 侍
	Lancer                             // 槍
	Knight                             // 馬
	Spy                                // 忍
	Fortress                           // 砦
	Pawn                               // 兵
	Cannon                             // 砲
	Musketeer                          // 筒
	Archer                             // 弓
	Captain                            // 謀
)

func PieceTypes() []PieceType {
	return []PieceType{
		Marshal,
		General,
		LieutenantGeneral,
		MajorGeneral,
		Samurai,
		Lancer,
		Knight,
		Spy,
		Fortress,
		Pawn,
		Cannon,
		Musketeer,
		Archer,
		Captain,
	}
}

type Piece interface {
	// Type returns the type of the piece.
	Type() PieceType
	// Color returns the color of the piece.
	Color() Color
}

type piece struct {
	pieceType PieceType
	color     Color
}

func (p piece) Type() PieceType {
	return p.pieceType
}

func (p piece) Color() Color {
	return p.color
}

func NewPiece(pieceType PieceType, color Color) Piece {
	return &piece{
		pieceType: pieceType,
		color:     color,
	}
}
