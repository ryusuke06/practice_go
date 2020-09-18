package main

import "fmt"

type Vertex struct {
	x,y int
}

func (v Vertex)Area()int{
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex3D struct { // Vertexをembeddedした（pythonとかのクラス継承）
	Vertex // Vertexにすでにx, y intが定義されている
	z int // つまりvertexにzを追加したものを3Dとするという意味
}

func (v Vertex3D)Area3D()int{
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D{
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale(10)
	fmt.Println(v.Area3D())
	fmt.Println(v.Area()) // Vertex3Dでembedded元のVertexのメソッドも使える！！しゅごいぃ！
}