package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarLista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	fmt.Println("Lista de tareas:")
	fmt.Println("--------------------------")
	for _, tl := range t.tasks {
		fmt.Println("Nombre: ", tl.nombre)
		fmt.Println("Descripción: ", tl.descripcion)
		fmt.Println("--------------------------")
	}
}

func (t *taskList) imprimirListaCompletados() {
	fmt.Println("Lista de tareas completadas:")
	fmt.Println("--------------------------")
	for _, tl := range t.tasks {
		if !tl.completado {
			continue
		}
		fmt.Println("Nombre: ", tl.nombre)
		fmt.Println("Descripción: ", tl.descripcion)
		fmt.Println("--------------------------")
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar curso de go",
		descripcion: "Completar curso de go de platzi esta semana",
	}
	t2 := &task{
		nombre:      "Completar curso de go 2",
		descripcion: "Completar curso de go 2 de platzi esta semana",
		completado:  true,
	}
	t3 := &task{
		nombre:      "Completar curso de go 3",
		descripcion: "Completar curso de go 3 de platzi esta semana",
		completado:  true,
	}

	lista := &taskList{
		tasks: []*task{t1, t2, t3},
	}

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Nestor"] = lista

	t4 := &task{
		nombre:      "Completar curso de go 4",
		descripcion: "Completar curso de go 4 de platzi esta semana",
	}
	t5 := &task{
		nombre:      "Completar curso de go 5",
		descripcion: "Completar curso de go 5 de platzi esta semana",
		completado:  true,
	}
	listaSecundaria := &taskList{
		tasks: []*task{t4, t5},
	}

	mapaTareas["Ricardo"] = listaSecundaria

	mapaTareas["Nestor"].imprimirLista()
	mapaTareas["Ricardo"].imprimirLista()
}
