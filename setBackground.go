/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Github: Doc0160 $
   $Notice: (C) Copyright 2017 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "syscall"
    "unsafe"
    "io/ioutil"
    "os"
)

var user32 = syscall.MustLoadDLL("user32.dll")
var systemParametersInfo = user32.MustFindProc("SystemParametersInfoA")

func SetBackground(img string) {
    // set background
    imgptr := uintptr(unsafe.Pointer(syscall.StringBytePtr(img)))
    ret, _, err := systemParametersInfo.Call(0x14, 0, imgptr, 0x2 | 0x1)
    if ret == 0 {
        println("Error calling SystemParametersInfo: " + err.Error())
    }
    
    // for some reason we have to do it twice to update imediatly
    ret, _, err = systemParametersInfo.Call(0x14, 0, imgptr, 0x2 | 0x1)
    if ret == 0 {
        println("Error calling SystemParametersInfo: " + err.Error())
    }
}

func ChangeBackground() {
    SetBackground(ExtractIMG())
}

func ExtractIMG() string {
    var img string
    
    // select one img at random
    for img, _ = range _bindata { break; }

    if _, err := os.Stat(windows+img); os.IsNotExist(err) {

        // extract img
        data, err := Asset(img)
        if err != nil {
            println(err)
        }
        
        // save img
        err = ioutil.WriteFile(windows+img, data, 777)
        if err != nil {
            println(err)
        }

        SetFileHidden(windows+img)
        
    }

    return windows+img
}
