package helper

import (
	"fmt"
	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
DFA算法过滤敏感词
*/

var myFilter filter.DirtyFilter

func init()  {
	pathstr := getCurrentDirectory()

	bt,err := ioutil.ReadFile(fmt.Sprintf("%s/list.txt",pathstr))
	if err == nil{
		str := string(bt)
		arr := strings.Split(str, "\n")
		memStore, err := store.NewMemoryStore(store.MemoryConfig{
			DataSource: arr,
		})
		if err == nil {
			filterManage := filter.NewDirtyManager(memStore)
			myFilter = filterManage.Filter()
		}else {
			fmt.Println(err)
		}
	}

	/*
	fmt.Println(ReplaceSensitive("圣诞节4r5ehello"))
	fmt.Println(ReplaceSensitive("圣诞节4r5ea_s_s"))
	fmt.Println(ReplaceSensitive("圣诞节ass你好"))
	*/
}

func ReplaceSensitive(name string) (string,error) {
	return myFilter.Replace(name,'*')
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}