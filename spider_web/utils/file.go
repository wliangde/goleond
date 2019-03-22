/**********************************************
** @Author: liangde.wld
** @Desc:

***********************************************/

package utils

import "io/ioutil"

func WriteFile(strFileName string, slcContent []byte) error {
	return ioutil.WriteFile(strFileName, slcContent, 0644)
}

func LoadXmlFromFile() {

}
