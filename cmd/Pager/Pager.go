package pager

import "os"

var PAGESIZE = os.Getpagesize()
const TABLEMAXPAGES = 100