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
    "github.com/marcsauter/single"
)

func init() {
    // only one instance of this
    s := single.New("YLP")
    s.Lock()
}
