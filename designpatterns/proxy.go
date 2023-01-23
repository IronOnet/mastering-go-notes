// The Proxy pattern forwards to a real object and acts as an interface 
// to others. The proxy pattern controls access to an object and provides 
// addtional functionality. The additional functionality can be related to 
// authentication, authorization, and providing rights to acces to the resource-sensitive 
// object. The real object need not to be modified while providing 
// additional logic.

// The proxy pattern comprises the subject interface, the RealSubject class 
// and the Proxy class 

// Subject is an interface for the RealObject and Proxy class 
// The RealSubject object is created and maintained as a reference in the Proxy 
// class. RealSubject object is resource sensitive, required to be protected, 
// and expensive to crete. RealObject is a class that implements the IRealObject interface.

package main 

// main package has examples show 
// In Hands-on Data structures and algorithms with Go book 
import (
	"fmt"
)

// IRealObject interface 
type IRealObject interface{
	performAction()
}

// RealObject struct 
type RealObject struct{} 

// RealObject class implements IRealObject interface. The class has method 
// performAction. 
// RealObject class method performAction 

func (realObject *RealObject) performAction(){
	fmt.Println("RealObject performAction()") 

}

// VirtualProxy struct 
type VirtualProxy struct{
	realObject *RealObject
}

// VirtualProxy class method performAction 
func (virtualProxy *VirtualProxy) performAction(){
	if virtualProxy.realObject == nil{
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()") 
	virtualProxy.realObject.performAction()
}

func main(){
	var object VirtualProxy = VirtualProxy{} 
	object.performAction()
}