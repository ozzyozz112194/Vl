package multimedia

import "fmt"

type ContenidoWeb struct {
	Multimedia []Multimedia
}

func (ml ContenidoWeb) Mostrar() string {
	var temp string = ""
	for _, f := range ml.Multimedia {
		fmt.Println(f.Mostrar())
	}
	return temp
}

type Multimedia interface {
	Mostrar() string
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

func (i Imagen) Mostrar() string {
	return i.Titulo + " " + i.Formato + " " + i.Canales
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion string
}

func (a Audio) Mostrar() string {
	return a.Titulo + " " + a.Formato + " " + a.Duracion
}

type Video struct {
	Titulo  string
	Formato string
	Frames  string
}

func (v Video) Mostrar() string {
	return v.Titulo + " " + v.Formato + " " + v.Frames
}
