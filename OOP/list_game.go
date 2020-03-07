package main


type List []*game

func (l List) game_lst(){
	for _, el := range l{
		el.print()
	}
}
