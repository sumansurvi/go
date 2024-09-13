package forLoop

import "fmt"

func Forloop(){
	fmt.Println("this function will have several kind of for loop.")

	i := 1
	for i <= 5 {
		fmt.Println(i);
		i = i+1;
	}

	for j:=0; j<=5; j++{
		fmt.Println(j);
	}

	for i:= range 3{
		fmt.Println("range",i);
	}

	//infinite loop
	for{
		fmt.Println("infinite loop");
		break;
	}

	for n:= range 6 {
		if n%2==0{
			fmt.Println("even - ", n);
			continue;
		}
		fmt.Println("odd - ", n);
	}
}

func init(){
	fmt.Println("init")
}