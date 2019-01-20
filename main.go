package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

var uiEntryArray [18]*ui.Entry
var uiResultLbl *ui.Label

var uiMotResLbl *ui.Label
var uiCogResLbl *ui.Label
var uiFimResLbl *ui.Label
var uiMotPercResLbl *ui.Label
var uiCogPercResLbl *ui.Label
var uiResultPercentLbl *ui.Label

func showNewWindows() {
	errwin := ui.NewWindow("Alerta", 320, 160, true)
	vbox := ui.NewVerticalBox()
	acceptBtn := ui.NewButton("Aceptar")
	errwin.OnClosing(func(*ui.Window) bool {
		errwin.Hide()
		return true
	})
	errwin.SetChild(vbox)
	vbox.Append(ui.NewLabel("Realmente desea cerrar la aplicación?"), true)
	vbox.Append(acceptBtn, true)
	errwin.Show()
}

func makeMainForm() ui.Control {
	rootBox := ui.NewVerticalBox()
	rootBox.SetPadded(true)

	fullBox := ui.NewHorizontalBox()
	fullBox.SetPadded(true)
	rootBox.Append(fullBox, true)

	//fullBox.Append(ui.NewHorizontalSeparator(), false)

	/* DOMINIO MOTOR */
	group := ui.NewGroup("DOMINIO MOTOR")
	group.SetMargined(true)
	fullBox.Append(group, true)

	vBox := ui.NewVerticalBox()
	vBox.SetPadded(true)
	group.SetChild(vBox)

	group = ui.NewGroup("Cuidado personal")
	group.SetMargined(true)
	vBox.Append(group, false)
	entryForm1 := ui.NewForm()
	entryForm1.SetPadded(true)
	group.SetChild(entryForm1)
	uiEntryArray[0] = ui.NewEntry()
	uiEntryArray[1] = ui.NewEntry()
	uiEntryArray[2] = ui.NewEntry()
	uiEntryArray[3] = ui.NewEntry()
	uiEntryArray[4] = ui.NewEntry()
	uiEntryArray[5] = ui.NewEntry()
	entryForm1.Append("Comida", uiEntryArray[0], false)
	entryForm1.Append("Aseo", uiEntryArray[1], true)
	entryForm1.Append("Baño", uiEntryArray[2], true)
	entryForm1.Append("Vestimenta Sup", uiEntryArray[3], true)
	entryForm1.Append("Vestimenta Inf", uiEntryArray[4], true)
	entryForm1.Append("Toilet", uiEntryArray[5], true)

	group = ui.NewGroup("Control de esfínter")
	group.SetMargined(true)
	vBox.Append(group, false)
	entryForm2 := ui.NewForm()
	entryForm2.SetPadded(true)
	group.SetChild(entryForm2)
	uiEntryArray[6] = ui.NewEntry()
	uiEntryArray[7] = ui.NewEntry()
	entryForm2.Append("Manejo de vejiga", uiEntryArray[6], true)
	entryForm2.Append("Manejo de intestino", uiEntryArray[7], true)

	group = ui.NewGroup("Movilidad y transferencia")
	group.SetMargined(true)
	vBox.Append(group, false)
	entryForm3 := ui.NewForm()
	entryForm3.SetPadded(true)
	group.SetChild(entryForm3)
	uiEntryArray[8] = ui.NewEntry()
	uiEntryArray[9] = ui.NewEntry()
	uiEntryArray[10] = ui.NewEntry()
	entryForm3.Append("Cama, silla, silla de ruedas", uiEntryArray[8], true)
	entryForm3.Append("Toilet", uiEntryArray[9], true)
	entryForm3.Append("Ducha", uiEntryArray[10], true)

	group = ui.NewGroup("Locomoción")
	group.SetMargined(true)
	vBox.Append(group, false)
	entryForm4 := ui.NewForm()
	entryForm4.SetPadded(true)
	group.SetChild(entryForm4)
	uiEntryArray[11] = ui.NewEntry()
	uiEntryArray[12] = ui.NewEntry()
	entryForm4.Append("Camilla - Silla de ruedas", uiEntryArray[11], true)
	entryForm4.Append("Escaleras", uiEntryArray[12], true)

	/* DOMINIO COGNITIVO */
	group = ui.NewGroup("DOMINIO COGNITIVO")
	group.SetMargined(true)
	fullBox.Append(group, false)

	vBox = ui.NewVerticalBox()
	vBox.SetPadded(true)
	group.SetChild(vBox)

	group = ui.NewGroup("Comunicación")
	group.SetMargined(true)
	vBox.Append(group, false)
	entryForm5 := ui.NewForm()
	entryForm5.SetPadded(true)
	group.SetChild(entryForm5)
	uiEntryArray[13] = ui.NewEntry()
	uiEntryArray[14] = ui.NewEntry()
	entryForm5.Append("Comprensión", uiEntryArray[13], true)
	entryForm5.Append("Expresión", uiEntryArray[14], true)

	group = ui.NewGroup("Conexión social")
	group.SetMargined(true)
	vBox.Append(group, false)
	entryForm6 := ui.NewForm()
	entryForm6.SetPadded(true)
	group.SetChild(entryForm6)
	uiEntryArray[15] = ui.NewEntry()
	uiEntryArray[16] = ui.NewEntry()
	uiEntryArray[17] = ui.NewEntry()
	entryForm6.Append("Interacción social", uiEntryArray[15], true)
	entryForm6.Append("Resolución de problemas", uiEntryArray[16], true)
	entryForm6.Append("Memoria", uiEntryArray[17], true)

	vBox.Append(ui.NewHorizontalSeparator(), false)

	group = ui.NewGroup("Resultados")
	group.SetMargined(true)
	vBox.Append(group, false)

	sumBtn := ui.NewButton("Calcular")
	sumBtn.OnClicked(func(*ui.Button) {
		fmt.Println("Button clicked!")
		res, _ := checkAllEntries()
		if res == true {
			fmt.Println("Faltan completar algunos campos")
			ui.MsgBoxError(mainwin, "Formulario incompleto", "Debe completar todos los campos del formulario llenando de 1 a 7 según el nivel de funcionalidad.")
		} else {
			showResults()
			showPercentageResults()
		}

	})
	group.SetChild(sumBtn)
	uiResultLbl = ui.NewLabel("")
	vBox.Append(uiResultLbl, false)
	uiMotResLbl = ui.NewLabel("")
	vBox.Append(uiMotResLbl, false)
	uiCogResLbl = ui.NewLabel("")
	vBox.Append(uiCogResLbl, false)
	uiFimResLbl = ui.NewLabel("")
	vBox.Append(uiFimResLbl, false)
	uiResultPercentLbl = ui.NewLabel("")
	vBox.Append(uiResultPercentLbl, false)
	uiMotPercResLbl = ui.NewLabel("")
	vBox.Append(uiMotPercResLbl, false)
	uiCogPercResLbl = ui.NewLabel("")
	vBox.Append(uiCogPercResLbl, false)

	rootBox.Append(ui.NewLabel("Programa en desarrollo, más info a d.hinojosa.cordova@gmail.com						16/01/2019"), false)

	return rootBox
}

func showResults() {
	sum := sumAllMotorEntries()
	uiMotResLbl.SetText("MOTOR:		" + strconv.Itoa(sum))
	sum = sumAllCognitiveEntries()
	uiCogResLbl.SetText("COGNITIVAS:	" + strconv.Itoa(sum))
	sum = sumAllEntries()
	uiFimResLbl.SetText("TOTAL FIM :	" + strconv.Itoa(sum))
}

func showPercentageResults() {
	sumMotor := sumAllMotorEntries()
	sumCognit := sumAllCognitiveEntries()
	sumMotor = mappingMotorScore(sumMotor)
	uiResultPercentLbl.SetText("PORCENTAJE DE INDEPENDENCIA")
	uiMotPercResLbl.SetText("MOTOR		:" + strconv.Itoa(sumMotor) + "%")
	sumCognit = mappingCognitiveScore(sumCognit)
	uiCogPercResLbl.SetText("COGNITIVAS	:" + strconv.Itoa(sumCognit) + "%")

}

func mappingMotorScore(p int) int {
	switch {
	case p == 13:
		return 0
	case p >= 77 && p <= 82:
		return p - 14
	case (p >= 73 && p <= 76) || p == 13 || p == 83 || p == 84:
		return p - 13
	case (p >= 70 && p <= 72) || p == 85 || p == 86:
		return p - 12
	case (p >= 67 && p <= 69):
		return p - 11
	case (p >= 64 && p <= 66) || p == 87:
		return p - 10
	case (p >= 62 && p <= 63) || p == 88:
		return p - 9
	case p == 61 || p == 60:
		return p - 8
	case p == 58 || p == 59:
		return p - 7
	case p == 56 || p == 57:
		return p - 6
	case (p >= 53 && p <= 55) || p == 89:
		return p - 5
	case p == 14 || p == 51 || p == 52:
		return p - 4
	case p == 49 || p == 5:
		return p - 3
	case p == 47 || p == 48:
		return p - 2
	case p == 45 || p == 46:
		return p - 1
	case p == 43 || p == 44:
		return p
	case p == 15 || p == 41 || p == 42:
		return p + 1
	case p == 39 || p == 40:
		return p + 2
	case p == 37 || p == 38:
		return p + 3
	case p == 16 || p == 34 || p == 5 || p == 36:
		return p + 4
	case p == 17 || p == 32 || p == 33:
		return p + 5
	case p == 18 || p == 29 || p == 30 || p == 31:
		return p + 6
	case p >= 19 && p <= 28:
		return p + 7
	default:
		return p
	}

}

func mappingCognitiveScore(p int) int {
	switch {
	case p == 5:
		return 0
	case p == 6:
		return p + 6
	case p == 7:
		return p + 7
	case p == 8:
		return p + 16
	case p == 9:
		return p + 19
	case p == 10:
		return p + 20
	case p == 11:
		return p + 21
	case p == 12:
		return p + 22
	case p == 13 || p == 14:
		return p + 23
	case p == 15 || p == 16:
		return p + 24
	case p == 17 || p == 18:
		return p + 25
	case p == 19 || p == 20 || p == 21:
		return p + 26
	case p == 22 || p == 23:
		return p + 27
	case p == 24 || p == 25:
		return p + 28
	case p == 26:
		return p + 29
	case p == 27:
		return p + 30
	case p == 28:
		return p + 31
	case p == 29:
		return p + 33
	case p == 30:
		return p + 34
	case p == 31:
		return p + 37
	case p == 32:
		return p + 40
	case p == 33:
		return p + 44
	case p == 34:
		return p + 53
	case p == 35:
		return p + 65
	default:
		return p

	}
}

func sumAllEntries() int {
	sum := 0
	for i := range uiEntryArray {
		val, err := strconv.Atoi(uiEntryArray[i].Text())
		if err != nil {
			log.Panicf("Error atoi")
		}
		sum += val
	}
	return sum
}

func sumAllMotorEntries() int {
	sum := 0
	for i := 0; i < 13; i++ {
		val, err := strconv.Atoi(uiEntryArray[i].Text())
		if err != nil {
			log.Panicf("Error atoi")
		}
		sum += val
	}
	return sum
}

func sumAllCognitiveEntries() int {
	sum := 0
	for i := 13; i < 18; i++ {
		val, err := strconv.Atoi(uiEntryArray[i].Text())
		if err != nil {
			log.Panicf("Error atoi")
		}
		sum += val
	}
	return sum
}
func checkAllEntries() (bool, int) {
	res := false
	i := 0
	for i = range uiEntryArray {
		if len(uiEntryArray[i].Text()) <= 0 {
			res = true
		}
	}
	return res, i
}

func setupUI() {
	//Create main windows
	mainwin = ui.NewWindow("Calculadora FIM  v0.1 ", 630, 480, true)

	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	//tab := ui.NewTab()
	//mainwin.SetChild(tab)
	mainwin.SetMargined(true)
	mainwin.SetChild(makeMainForm())

	//tab.Append("Agregar usuario", makeMainForm())
	//tab.SetMargined(0, true)
	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
