package main

import (
	"fmt"
	
	//"./figuras"
	"./multimedia"
	
)

func main() {

	
		var opc string 
		var canti int

		var titulo string
		var formato string
		var canales string

		var tituloa string
		var formatoa string
		var duracion string
		

		var titulov string
		var formatov string
		var frames string



		fmt.Println("Cantidad")
		fmt.Scanln(&canti)


		for i:=0; i<canti; i++{
		fmt.Println("Que deseas ingresar:\nImagen 0)\nAudio 1)\nVideo 2) ")
		fmt.Scanln(&opc)

		switch opc {
		
		case "0":
			
			fmt.Println("Titulo :")
			
			fmt.Scanln(&titulo)
			
			fmt.Println("Formato :")
			
			fmt.Scanln(&formato)
			
			fmt.Println("Canales :")
			
			fmt.Scanln(&canales)

			
			case "1":
			//Audio
			fmt.Println("Titulo :")
			fmt.Scanln(&tituloa)
			
			fmt.Println("Formato :")
			
			fmt.Scanln(&formatoa)
			
			fmt.Println("Duracion :")
			
			fmt.Scanln(&duracion)

			
			
			

		case "2":
			//Video
			fmt.Println("Titulo :")
			
			fmt.Scanln(&titulov)
			
			fmt.Println("Formato :")
			
			fmt.Scanln(&formatov)
			
			fmt.Println("Frames :")
		
			fmt.Scanln(&frames)

			
			
		}
		
		ml := multimedia.ContenidoWeb{
			Multimedia: []multimedia.Multimedia{
				
				multimedia.Imagen{Titulo: titulo, Formato: formato, Canales: canales},
				multimedia.Audio{Titulo: tituloa, Formato: formatoa, Duracion: duracion},
				multimedia.Video{Titulo: titulov, Formato: formatov, Frames: frames},
			},
		}
		
		fmt.Println(ml.Mostrar())
		
		
		
	}
	

			


		

	
	

	//i01 := multimedia.Imagen{Titulo: "Imagen 1", Formato: "JPG", Canales: "Sonoro"}
	//fmt.Println(i01.Mostrar())

	//a01 := multimedia.Audio{Titulo: "Audio 1", Formato: "MP3", Duracion: "6"}
	//fmt.Println(a01.Mostrar())

	//v01 := multimedia.Video{Titulo: "Video 1", Formato: "MP4", Frames: "24"}
	//fmt.Println(v01.Mostrar())

	//fmt.Println(addMulti(i01, a01, v01))

	//c01 := figuras.Ciruclo{Radio: 6}
	//fmt.Println(c01.Area())

	//fr01 := figuras.Rectangulo{Altura: 10, Base: 2}
	//ffmt.Println(r01.Area())

	//ffmt.Println(sumArea(c01, r01))

	//fm := figuras.FiguraMultiple{
	//Figuras: []figuras.Figura{
	//fc01,
	//fr01,
	//figuras.Ciruclo{2},
	//f},
	//f}

	//fmt.Println(fm.Area())

	//c03 := Ciruclo{15}
	//c04 := new(Ciruclo)
	//c05 := &Ciruclo{100}

	//fmt.Println(c01)
	//fmt.Println(c02)
	//fmt.Println(c03)
	//fmt.Println(c04)
	//fmt.Println(c05)
}
