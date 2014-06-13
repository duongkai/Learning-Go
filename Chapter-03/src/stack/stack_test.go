package stack

import "testing"

func TestPushPop (t *testing.T) {
    //c := new (Stack)
    var myStack Stack
    myStack.Push (5)
    if myStack.Pop() != 5 {
        t.Log ("Pop fails")
        t.Fail()
    }
}
