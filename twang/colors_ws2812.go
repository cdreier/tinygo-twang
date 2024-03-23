//go:build rp2040

package twang

import "image/color"

var colorPlayer = color.RGBA{0x00, 0xff, 0x00, 0xff}
var colorPlayerAttack = color.RGBA{0xff, 0x00, 0xff, 0xff}
var colorGoal = color.RGBA{0xff, 0xff, 0xff, 0xff}

var colorEnemy = color.RGBA{0xff, 0x00, 0x00, 0xff}

var colorOff = color.RGBA{0x00, 0x00, 0x00, 0xff}

var colorFireActive1 = color.RGBA{0xff, 0x00, 0x00, 0xff}
var colorFireActive2 = color.RGBA{0xff, 0x33, 0x33, 0xff}
var colorFireActive3 = color.RGBA{0xe0, 0x00, 0x00, 0xff}

var colorFireInactive1 = color.RGBA{0x14, 0x00, 0x00, 0xff}
var colorFireInactive2 = color.RGBA{0x29, 0x00, 0x00, 0xff}
var colorFireInactive3 = color.RGBA{0x3d, 0x00, 0x00, 0xff}

var colorWater = color.RGBA{0x00, 0x0a, 0x52, 0xff}
var colorWaterDark = color.RGBA{0x00, 0x02, 0x14, 0xff}