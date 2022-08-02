// Installation
//	 Terminal
//		go get "github.com/ahopo/eztmpl"
//	 To your project
// 		import "github.com/ahopo/eztmpl"

// Usage
//
//		import "github.com/ahopo/eztmpl"
//
//		type Person struct{
//			Name string `tmpl:"PERSON_NAME"`
//			Age string `tmpl:"PERSON_AGE"`
//		}
//		//IN
//		/* ./path/sample.txt
//			my name is {{PERSON_NAME}}
//			my age is {{Age}}
//		*/
//		func main(){
//			person:=new(Person)
//			person.Name="Jhon"
//			person.Age=12
//
//			tmpl:=eztmpl.New() // Initialiaze
//			tmpl.File("./path/sample.txt").Struct(person).SaveAs("./path/newfile.txt")
//			output_string:=tmpl.File("./path/sample.txt").Struct(person).String()// string
//			fmt.Println(output_string)
//		}
//		//OUT
//		/* ./path/sample.txt
//			my name is Jhon
//			my age is 12
//		*/
package eztmpl

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/ahopo/ezs"
)

type easyTemplate struct {
	filename string
}

type _struct struct {
	_tmpl  easyTemplate
	_input interface{}
}
type out struct {
	_struct _struct
}

func (tmpl *easyTemplate) File(file string) *_struct {
	tmpl.filename = file
	_s := new(_struct)
	_s._tmpl = *tmpl
	return _s
}
func (_s *_struct) Struct(_struct interface{}) *out {
	_out := new(out)
	_s._input = _struct
	_out._struct = *_s

	return _out
}
func (_ss *out) SaveAs(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(build(_ss))

	if err2 != nil {
		log.Fatal(err2)
	}
}
func (_ss *out) String() string {
	return build(_ss)
}

func build(_ss *out) string {
	dat, err := os.ReadFile(_ss._struct._tmpl.filename)
	check(err)
	str := string(dat)
	field := ezs.Get(_ss._struct._input)
	for _, p := range field.Data {
		replace(&str, p.TagValue, fmt.Sprint(p.Value))
	}
	return str
}
func replace(text *string, tmpvar string, value string) {
	tmp := regexp.MustCompile(fmt.Sprintf(`{{%s}}`, tmpvar))
	*text = tmp.ReplaceAllString(*text, value)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func New() easyTemplate {
	return *new(easyTemplate)
}
