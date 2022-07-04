package main

import (
	"fmt"
)

type FingerTable struct {
	var id int
	var c_bits int
	var table[] int{}

	func FingerTable(var id_node int, var node_ini ChordNode, cbit int, node_id int)
	{
		id := id_node
		c_bits := cbit
		pos = 0

		for i := 0; i < c_bits; i++ {
			valkey = id_node + 2^i

			node_possession = node_ini
			for node_temp.id < valkey, pos++{
				continue
			}

			table.append({valkey, node_possession})
		}


	}

	func findkey(key){
		for i := c_bits - 1, i >= 0, i-- {
			if table[i][0] > key
				return table[i][0]
		}
	}

	func refresc(chord chordNode){
		pos = 0
		i := 0

		for ; i < c_bits; i++ {
			if table[i][0] <  chord.id {
				break
			}
		}

		for ; i < c.c_bits; i++ {
			if table[i][1].id < chord.id {
				break
			}
			table[i][1] = chord
		}
	}
}

type Proto_Chord struct {
	var maxnode int
	var node_ini ChordNode

	func create(maxval int, chordNode_ini ChordNode){
		maxnode = maxval
		node_ini = chordNode_ini
	}

	func insert(chord ChordNode){
		var var_id := chord.id

		if node_ini != nil {
			node_succesor = node_ini.findsucesor(var_id)
			node_predeccesor = node_succesor.predecesor

			node_succesor.predecesor = chord
			chord.succesor = node_succesor

			chord.predecesor = node_predeccesor
			node_predeccesor.succesor = chord
		}

		else {
			node_ini = chord
		}

		node = node_ini
		while node != nil {
			node.refrescFinger(chord)
		}

	}

	func delete(chord ChordNode){
		var var_id := chord.id

		state = "disconnect"

		node_predeccesor = chord.predecesor
		node_succesor = chord.succesor

		node_predeccesor.succesor = node_succesor
		node_succesor.predecesor = node_predeccesor
		

		node = node_ini

		while node != nil {
			node.refrescFinger(chord)
		}
	}

}

type ChordNode struct {
	var server_ip string
	var state string

	var id byte
	var key byte
	var path Etiqueta

	var predecesor chordNode
	var succesor chordNode
	
	var fingertable FingerTable

	func lookup(key){
		return FingerTable.findkey(key)
	}

	func create(){
		predecesor = nil
		succesor = this
	}

	func join(n ChordNode){
		predecesor = nil
		succesor = n.findsucesor(this)
	}



	func notify(n ChordNode){
		if predecesor == nil || predecesor.id < n.id && n.id < id {
			predecesor = n
		} 
	}
	
	func findsucesor(id_node int){
		if id_node < id{
			return false
		}

		else if id < id_node {
			return succesor.findsucesor(id)
		}

		else return this
	}

	func find_predecesor(id int){
		if id_node > id{
			return false
		}

		else if id > id_node {
			return predecesor.find_predecesor(id)
		}

		else return this	
	}
	
	func hash(value int){
		return value % maxnode
	}

	func checkPredecesor(){
		if predecesor.state == "disconnect"{
			predecesor := nil
		} 
	}

	func stabilize(){
		x = succesor.predecesor
		if( n.id_node < x.id_node && x.id_node < n.succesor.id_node ){
			var succesor := x
		}
		succesor.notify(n)
	}

	func refrescFinger(chord ChordNode){
		fingertable.refresc(chord)
	}

	

}

func main {

}