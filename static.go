//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a map containing all embedded resources. The map key is the
// file name.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES = map[string]EmbeddedResource{

	"data/markdown/demo.md": {
		Filename: "data/markdown/demo.md",
		Contents: "H4sIAAAAAAAC/0zNMarDMBCE4X5PMeDmvcY+Q8oUaYIhRUix2GMkYklBu47I7YNdpf0G5u8w0lxkDNEQDQqnOV5qTrQQp7BrYCW8YGYq2byqEx6IqaRUsixU3yqtF+k6XEolZrrG1UTOCz5lQ4sWkH6WQyfNuL8j2/F2Pd2QtD7n0vLjb6jahj34jyWu7OUbAAD//7WwbserAAAA",
		Length:   171,
	},

	"data/markdown/konami.md": {
		Filename: "data/markdown/konami.md",
		Contents: "H4sIAAAAAAAC/3xV3W7bRhO911PMF198diGLQC9ZwEAbpEGQIg3apEBhBNaKOyInXM4s9kcyCz98MUtRpi2gvLHP6uzM2Zmds1fwVrgNJmVnEgnH1eoO/pYMnTkgxNw0GOM+OzdCwAGHHQa0kDqE+4/CZiB4Kxa/XXcp+bqqkDdH6smjJbOR0FaKqon5oMybWhPcwf/uX+X9dl3R0Fa48dzCm1c/vrnRXde/BxglwxEDQksHZDDgiHtIAqmjCN60CPWRuK/hFlpkDEalj5L/b6HFBJ3u3Y2AnDAQt+UojVgEYTA8wsE4slMg4SlqpISbm9XqCn4RzhF+RZNyQKjh3SDfSaV92JccB4S9ZLYLNSo4dpKdhZ7lCKkzqSQdTOitHBn2EgaTEgaI2XsJKZbfc0SQPaCmWEOkwbuxLG5rNgPWWzBzoiM5BzuEgN6ZBi0cKXUliPE+iA9kEgINpsWNqv3SIWgMGHJMutGcjl2SQcIwrMFEyBywkZbpH7RlwzIV23IVMjed4RZtifxXCTNRiRuXLRYdsjuQ5KjNv4MnrVtPET6phif4Y471QRXCU+HUt89f/QKVhcLZamK3hfI9QV1gDadv5njvcMFR+Jqzo36mKEfhmXLmyLiFBUfGZ8rMacyLOArrS05axmlMuoxjpV1yrLSXnD2xM2y3J84J1i84LQW30KPwQk9PMS44Ci844iyGh8HwtnDOsL7gHEVZZ06B9ZnjkdtMPGs+wZea40DLfhX4ul+538KiPrlflAeeVqvV6uoKft5JTvBFB+SzaXG6+PNY6t+AETmpefggB5xG8/7PhAeEjzj4ydViXVVR1zY9Dn6zp+pmDRYP6MSjXcMxSMJ1GUaL3smI81TirjiHJv7Kewkps0noRogyIHgU73AaKOJIMZ29gYKuJONO9qc+MIyq8js2KYIJag17asi4kvihkwFvB2PxASxGT+l0Ghad77PPmTgbAEUtQfEZjl56ZFAzimSL64hHvo2SQ6Oj7tBEjD+9kBxwFzT1KDnAUUK/hjR6aorfqv9o8bMOdkyGrUqgBBIAH81AfMLEWrTUbeCdiWqB2LYRHPU41VBPmiPus9ZoOlIjnPAxbaYmT+2dCvPc4QnrCU/i7SSlHFjtrZRUwftPX+H99E7A57xz1MBv1CBHXMMBQyRh+HGtup1JGKY+0356hyh25Z/G8Kn3UznU1ZcNLG/eD3A/X6eWUpd3m0aGKvaP1fwSxM4ErJ5v3X/SbjarfwMAAP//vmZ5RLwHAAA=",
		Length:   1980,
	},

	"data/static/api/index.html": {
		Filename: "data/static/api/index.html",
		Contents: "H4sIAAAAAAAC/5xWbXPUthN/70+x6M/wT4bYyh1kShLbDFxISQskkGMKrzo6eX1WTpaMJN/F09LP3pHsu1zCQ2nzIufV/vbpt+u10qfXtYQlGiu0ysgo2SeAiutCqHlG3k9P4yfkaR6l907OJ9OPFy+gcrWEi/fPX51NgMSU/vZoQunJ9AQ+vJy+fgWjZETpizcESOVcc0TparVKVo8SbeZ0+o5ee/PRyBusn5PCFSSP0uD4upbKZl+xHR0eHvYWxIOOJPP5oQqWyIo8Smt0DLxljJ9ascwI18qhcrHrGiQwSBlxeO2o93QMvGLGostaV8ZPCNA8SqVQC/AWGRE1myO9jgXXioBBmZH+sTJYZoSWbOnlRHDd24YMFKsxIwVabkTjhMdvQr9tBV8AUwUgsx3Yihmh5qBLqJlZFHqlwGf3hbcFdittCrvlam2wBw2zDmdC7W2cDH73NgHWit6xdZ3EPPqfUFy2BYJ1zAlOubU0qBJubZTSAdYzEqpnrau02dRftTVTNnHrhJ1wEvPXQ6zLihmEI3h2cZbSXhWldOjVTBddHqWFWIIoMrIyrCFbskehCb0d5SlbByS3nSdc1ylleUqrkYeO8x/kN6XV2CdTiGUe3UR1uqlRtT5sK0Pht2K/1DX24aS4q2WNoCQPpX4dwA0yh5Tkk/DwTViFzNkK0XmoFy698C14yT5Rkp8+e7sFoCH3u7UNc0PyCMCTertLIXFPIoS/dCY1X3xqtcPNUTXOf7k8fzNAx9+BAqRN/o65Cg24iilorW+CqxBWOLPCIXS6hZp1wKTVYNtZLVxoDMw6uDi/nHp8p1sDBXMMCmGQO9mB08HLXVZJHn6FVoCqiBstVGDsKKVNHqWNwTy6D7w1EuKXQJ5xjo07AtY0UvBgR6+sf0/jDyE6xHGIGzZF5v89nGu0Dys0SMKGsUeUbr9sYRTXPY5SGkJusXFufGUGWbFmYuiHhdLoGhiUQuJ/z/d0yPWBdMe0Ya6iTlN/ktTFv844pV9v/0QrqyXCRAqf+Y9MwdRTBn7CaqY64LqumSpiKRT2ayvmvTPQrfO8GEwAzkrf/P8vEebaQepcvl5xwqXUuRyEso5JiUWYJM4UWLbcHowvSk24ViXJXSVs4NrPB8xQIXNVCEFba2iA03W0pKB9OL9QWosQrMMAr0eyb9lNxZ5FW6HUM/gL7kPbOFEj/Ak3BUA8+05DouHL96VmKXBFy5+K0cHjfYzH4xmLH+P4UcyelCwezcoDLEeH5cH+wQ8089bZsCrCpuCSWZsRLpEZksO2yi+RUmsX1jLANlxiOSyWHn9ba8S86tUAE910QYYHXDfdMYz3R49T2zDVL6nWGFTu9y6ET6lX5P7N2TR2zZx1uMRkgXWTlH7zXnoZfsW68Y39Vip361ovyUG8F8dw0RqEK7Zk/ed7D5SGq7ctmg7i2H87w/FwQQg3iRuwL3IlVKFXiVZSswIyKFvF/Ru7swt/eAoohZ8xjDr4KhN/tmQGOgMZ7ChcwQlzuLO7m8zRnbZSfkRmdnaPo8H4fVMwh8G+0LytUbm1i0Bjtjn2Dl5I9I/Pu7Ni5za7u8feTJSw09vdy0C1UsKQJoTTxBc46TcVZEBiAg+hM8H0cwSf/SUhVP7PxNy9bFxZukiu7LYHOlwKwt0sj/4OAAD//zWOPLuaCgAA",
		Length:   2714,
	},

	"data/static/cheatsheet/index.html": {
		Filename: "data/static/cheatsheet/index.html",
		Contents: "H4sIAAAAAAAC/4xWbW/bNhD+rl9xYbEtySKxdoqhdSQVrdOuwfqSri62oBgKWjpbjClSJSkrwtL/PpCSbCdN0+kTebznuReRDxk/vSoFrFEbrmRCRtFDAigzlXO5TMjH2cvwMXmaBvHe6bvp7OL8BRS2FHD+8fnrsymQkNK/jqeUns5O4e9XszevYRSNKH3xlgAprK0mlDZNEzXHkdJLOvuTXjn4aOQAwzjKbU7SIPbEV6WQJrkDO3ry5EmHIM5pIpjLD6VHIsvTIC7RMnDIEL/UfJ2QTEmL0oa2rZBAP0uIxStLHdMJZAXTBm1S20X4mABNg1hwuQKHSAgv2RLpVcgzJQloFAnphoXGRULogq3dPOKZ6rA+A8lKTEiOJtO8stz5b0K/r3m2AiZzQGZaMAXTXC5BLaBkepWrRoLL7hu2FbaN0rnZoRoAR1AxY3HO5dGGpOc92gQYFjpiY1uBafCAy0zUOYKxzPKMZsZQvxRlxgQx7d26jvjqWW0LpTf1F3XJpInskLDlVmD6po/1oWAaYQLDHKYFMmsKRBvTzjWIaf/v5ipv0yDO+Rp4npBGs4rszJ0Xav+vR2nMhgTIzWBRpsqYsjSmxci5jtP/2e+YFmOXTM7XabCNalVVoqxd2Fr4RtyI/UqV2IUT/PYqqzgl6bPzs+85ZBqZRUrSqR98123TM+fqJh+6Bt7tvmBfKElfPnu/40B97rdr6/cRSQPwn2vsnX/KtbJ3mQuVrb7UyuJgAogtmwuEhue2SMhvD38i2zW3qtPY5ukDeIXMSQqMYmpzb3MRe6uP0i1Qq+8k2DKMdxjGOwzjHzFsKY53KI53KI7vpzg8nCuRHx5u0fPUWWI6vx/4mVsmePZ5C+RpZ4opvx/6qcR/9ns5NBbX6NQwqlf0YEs2/P673Eg6bNJvgsTU/7sdQ5WeOckz4E4uE0aBqatKaYv5EbDO7DEmggtVQ8YkLJgQ4ZxlK7AK/BXAF9CqGhpuiiNgwhaqXhag5muuaiNaaBAqjQvUYAtmgRv5i4XaYB7FtNrJphinMx+sO54b+13b0Cc/K7gBboCBYWUl+lwnnjWuNKbBNbybX2JmAeAangm+xqduGFzDJOw+uIZ+NPH2KbMd/zVcoOmHwTX4zDr7WwWDPaY+DNxI662yCLZAV6OTntjadBJTa1PXMncQtRLeIUNpcVAoZ/AFhBkKAdxiaW51iN55Iotx/xvv6VvXjr1PTFivgN0eMxNKl9wW9dxJKWU5K8OKDlIZFqiRatbQ0t04mhqd0UyVpZLUX5WGuuvw0eOokksgr9VSwczpPMzwysKIHAz9uZn47ZlXKi9UmWDGJCQTyDRJYXfJadhCKetvBYBdd4GLTtd6/5urmi+LQfamqmr9HH7OVNWewPjh6FFsKiY7jay1Rmk/tz58TN1CCvMWbpw4szlyKyyraOGE/4Obwx9YVu7sfS+V23UNGt1P98IQzmuNcMnWrHtNHIFUcPm+Rt1CGLqr3Jv794p/2GydXZENl7lqIiWFYjkksKhl5t4k+wfwr2sBpfA7Wr/XXJWRs62ZhlZDAvsSGzhlFvcPDqIl2pe1EBfI9P7BSdCDP1Y563d3rrK6RGkHCt/GZGN2BC8EuuHz9izfv9ndgxMH4wvY73B7CchaCOjTBG+NXIHT7uKCBEhI4FdotYd+DeCre7P4yn/cmNtvn0tDV9Gl2WWg/ZvEPxXT4L8AAAD//8gkbVApCwAA",
		Length:   2857,
	},

	"data/static/css/Makefile": {
		Filename: "data/static/css/Makefile",
		Contents: "H4sIAAAAAAAC/4yPsarDMAxFZ+srREiWB4738AreO/QbTKI2Acc2ljwY/PElSbt3OiDEvecO48yMEw7jFkBZIRbUhKZwNry6TCbVk4thv+075S/HVLE1pHmN2Lkk+kWCW2Bx3mOqssagP68dKHtdfg3u//Hw0joWSUVuvQVY6OmKlwlZqqfDG5S96h/3DmD25MIEyuYd/85ZraHkQvAOAAD//yhspiDlAAAA",
		Length:   229,
	},

	"data/static/css/style.css": {
		Filename: "data/static/css/style.css",
		Contents: "H4sIAAAAAAAC/6xX32+juBP/V/xNVOl7EkRAmm4LT3t7j919udXes8FD8BU8yJhNsoj//WTzIya47e1pZRWV8cx4+MxnZpwU2aXLUSg/pxUvL/FHyWnpNVQ0fgOS50lF5ZGLOEhqyhgXxzhIUpq9HCW2gsVbOOiVZFiijLf7XK/EOGz4D4jDh/qclFyAXwA/FiqO7utzT7tRPwgeexoX+B3kJHoMgn57krTuxpPvg/pMaKswOXGmivgpuFtEkOd5kuLZbwrK8BRrbf0XHuozkceU/j/wzNqFvyV+hT/8f618gvSFq/+gLxlIX1LG2yYOg/o8nbuSr0X9tgDKQHY33hTWfgm5Wrpd60iN8jtna0Xt661942cRpDOCN09eZOlJr5lEkUGSmNwG0zeTIuzWdCN67bU/Q6vTwKMUS/Yq6/aPN6x7OGgBKAXSb2qaacd+eEWbFCGhnv0ykHIp+s4broCNTE1Lmr0kCs7KZ5ChpIqjiAUKuHqNul9CirkIDRYauxkeTcQwugFHoKxomdhxXsGJIjcWCeNNXdJLzIWGrt8qrCsQbbdIY6jXfLoObohnDDpFpbCKw/pMGiw5I1sI9bKIst6cjyJt6SBAUvJG+Y26lOCrSw0jxpNNybtl5BZ1RpoFtjahnTNpuwaozIouL5Gq2LC53xoArHZ2r5G+KowtaR9oCEY26n40fsR+gmcMZPBHin3nyJYN8yHXy/ZoIz65D22nbfmLuOZC2zEFiDl9fsxRDAiPkYdR5qwQB10YKgWMaJOlr+V4yMLoDeMsjPpthkKBUGMidcuaJsfD3RK6of9EJvzRStesIzm35WP1lv1+Yb636fLBWZj9LtNRrQOMHtfZ3WWGad2adS7tbY6o1tNjgMo1HJx660GznhGDqjUmXtW5jpr3onknCrtCxp7vqJBoOZfsdmOaVz6BdKWpNv9Jmua55Wek6Bu8zLJ+d5PzfrdKrCEGlZ15ximqoq8ldKeCKzC9GuJagq9vR4ktNNBPO+R/vKpRKirUUsltipYcJTP/xakE+uLr99VVy9FJDmsiWTtv6c4XKy4aUGS6Xn24vV1F66vb+xY/pezqrwE5zL37cGWVzSjGWGJfnz+haLCkjbf5KBgtgXxGgeSvr5uFYONtntuMM0oGfbgK/qSiIV8vNZwkVyA33uYP+Jt+a4eN0fh3rholgVbkG0i62HrmKQwMniRfeJW2wz553nifUdAMvc0nbCUHSb7AaeONL16FAg0vrIYX7J6g6neKzmM2LTF7mcA4zGAcJVwWOPa6pqgE2g0NKwyCO+fviYEhg1LQM9Xd3vT6DNlrdeAgbf9PAAAA//93GN5E1wwAAA==",
		Length:   3287,
	},

	"data/static/css/style.in": {
		Filename: "data/static/css/style.in",
		Contents: "H4sIAAAAAAAC/6xXS4/bNhC+81dMdxGgSS1bsuONLZ3StIcCm/SQID1T4shiLZECScf2Gv7vBanHSrK8xrYF92FxHhx9880MHUt2hBNJpTBeSgueH0P4qDjNJ6Cp0J5GxdOIFFRtuAjBj0hJGeNi4z7HNNlulNwJFsI9Lu2KSCJzqUK4X6R2RZVvzZ8whOChPEQk5wK9DPkmMyHM39utMyEUTtCY+r7vr1YRnAkNM/kDVUe2WlmplZH7vaIlnNro3vvlAejOyIjsOTNZCGv/zSDKNE0jQmJ58HRGmdyHYI3sb7AsD6A2Mf3Zn0D9Mw3eRsQr5JP3Oos9xltu/q2RYqg8RRnf6RAC3wJUxzAiGd204GRImUVu6NjI0ssxNcMTLrWUTdHNQKyq9feyhqrS3Yv5SiQ3Ihjkc21Xh5ZzB3TFBd8i0QCRBR2q9Ijs1sI5d2Td19yMZc5e5PNidcHnh6XbQmNQebqkiTvCC6qsdGKhE+g+VTwf7P3gmhtkcCLQ8D/OabKNCIDBg/EYJlJRw6UIQUiBvTPml7n/T6RqsXNr7vZaEC2rg/kFhEKqguYtiHX0HQjn8+t4Ma7LnB5D4MJiXNHayLJAsYPTgAeBXZ2AbMx1kPXbxNIYWYQQlAfQMucM7jGwq0u5UfH5+dhdDid4ZhF02iHkXBtPm2OOnjmWWKekY5xzOMHwrTouGvL60DeyvfFqvslUI1VJBidIc0lNCK5qqg5ZQ9Xtwa7h9jSbbrnwHVwN4V2rbF510cBZx+hAcd6zRXPAIOnQS9AytSuCrnfoZStqcQ0GR1jQ/1cqk/FUjSe2Cqf90wurTk7zUsE8ia5kCkZYyKSx5V2ZDbxeDL4kmL/oxMltzhMpDArzzAfbYyNoRuLDGxgiXXVLVyznZ3vbQGA0scP6hV4LXCzsVtfRonFUU/CDVRj1bOmc2HivRD9fjREFzmSaOC5fFMEtO3KfSmnG5mSF8fgQHNUcG6qXs7BS7o3Dq1rdsXo7qpvR9DtmM8965d4bo1f7ouu1qSuEGr1uEbjKfn0ROLOOx7YAXuJ8klScGaPMmUzHOXF2HKOuuOx/O+lN5hwRUip08xZgn3GDbiZhCKVCz143I5i9g0RrbwHvZpdqLpWNLvzEi1IqQ4VxZp/lE89zOgHNRYIQrNfrcSe9w/4sUVF47z2M68ox7Q+trlTMCUOIFdKtZzec4h/CoBJo4PdDmUuFCpbT5S+NXYcpFTHs5kiTdbcdJxvtzq34plV7UeZCo2lvyx+eL8tuTedve7G81uz1Fk1FVCVkd+om4sOyHYdLv/+iYadYGGOVqPcl65MUWuZUT+Duo2A0R/gshYS/vt31d+zj4y7hjNYm2Nn5SoWGb8cS94obVFbyG/5Nv+8qSWP/KzfaKKQFfLfk6MkeeYxVibZbX3gR7yoNeLyb2A80kRO4+yR3iqOCL7i/m0D9NIFCCunY2HlN1+j96RoLd0UhU0PrqmqvP3Euq5tsC5q7RFaobRQe+/i3s9d2Fqqw8Vc3+MC3X/U66bH3h56932NibeYuGcxUvkYu/2dCZu/gKyJkxpQ6nM023GS7eJrIYqa3h1lB1ZbJvdAZVeh2udY71DNXrolkeNlM2oIdlGinQsmZ/BMAAP//W05YSZsPAAA=",
		Length:   3995,
	},

	"data/static/faq/index.html": {
		Filename: "data/static/faq/index.html",
		Contents: "H4sIAAAAAAAC/5xWW2/cNhN9318x4Qd8sNFdMZv0IbElBY4vjdFc7GSNNE8FVxqt6KVImaRWKxT978VIWlm+NUafxMucw5kzwxHDd9tCwQatk0ZHbB68ZIA6ManUq4hdLc5mb9i7eBK+OPlyvPhxcQq5LxRcXL3/eH4MbMb599fHnJ8sTuCPD4tPH2EezDk//cyA5d6XB5zXdR3UrwNjV3zxlW8JPp8TYDcOUp+yeBK2xNtCaRc9gp2/ffu2QzAyOlCC/EPN4glAmKNIaQAQFugFEH6GN5XcRCwx2qP2M9+UyKCfRczj1nPiO4QkF9ahjyqfzd4w4D2TknoNhIqYLMQK+XYmE6MZWFQR64a5xSxiPBMbmgcyMS0+dL5RGE/+J3WiqhTBeeFlwhPneLsVJM5NQt6b3R7XUovK58YO5HlVCO0Cv/W3vrVRalFgxFJ0iZWll+TPEN5lJZM1CJ0CCteAy4WVegUmg0LYdWpqDaTAo4xrbGpjUzei24GmUArncSn1dCDquafDIbuNW3IvvcL4U7/+LRcW4QDOLN5UqL1q4MitMYXLCh3F4ULeISi3fJfccGnSpidM5QZkGrHaipJ1a6NVQqBti2oeh2InJLvrQZCYIuQiDnk+J9NX8TNFC3n+Kp6EPJWbeDIZTvWmLFBXdGyl4kmo5J2zP5gCu+OUvL8rSslZfHRx/pRBYlF45Cw+bgdPmuUovMsRPZnS5BtNnjLPxA1n8dnR5ciAt773sd3TtS+HQXC6efP439JIyg62S2WS9U1lPN4uErsaT2nBx99z4UE68Ll04KRHyIx9F/LU37dN7y4AhGW8GFD0NQpVAyk6udKYgjcglDI1NKaiCQonVZdsvJvnIOTlI+yPZOWCLgVITZx2ICFVp+Bz1D27zxEsukp5Kqv2wtfS5x0qsxJ16qaQGKVQrCp0UzAWjM/RQommVPjQo5DfV4Dk+2BqcJhYbEUsGqhKZURK0ePWP1vHI+VzU63y1vGrrx8d1LlMcqBYVqjRCo9pO2NW6NQUDPyQOASjSVdMKit9A6VFh9oHAOcZCLBYGI9QObSQSZ06EJ0i3nR6kKdE00AiNGwk1iB9qwhpS+GAfDJFy/jzl8VpyJfxASxytG0laAOyLWBpNF1q3JCwwjnqdrPSGo9JmxlyflCsL3r3XO3b0nUm8zUJU5oa7a6O0W5kgv+hjG2lSR8ni1IhrAz9/WZL4Uj9slQyERTTtE8PFTZpRrJCbhwlqSqNhpX0ebWEoYDpN+kOOO/WqR9yt97ycWOn6n6WGRX7zyUK+fi2h/xhS+hbz6jzJEo4F7FEobAsfsSAWlNmjG9b/l2QwqxrVwNqvGvlKt91s2NTNu0c/p+YsjmEVy/nv4auFLprfZW1qP2fTetEyGkjhmXzUEzncYPBGosyyKirf6M5/I5FSRI95cq96MYNeBRw+GI2g4vKIlyLjej++1Oq6+vLCm0Ds1lv1231r5f2qXML2LXvWurU1IHp7lIEWaUTKqO9ffhrlyLO4TdsryFQ5MFufSMsNBYi2NNYw4nwuLe/H6zQn1VK/UBh9/YPJyOSqzIVvmuAqUmqArUfU7UyR8MWEZ0qpOH75jzdu6v+/uEOKjPY67AvItCVUjByHdqdgII/7q4wRMBmDH6Bxg4Uf09Gn5B3CtFf/ScK3n/ZXTu+Dq7bN11P0ZZ3+2IJ23dmPPknAAD//1ZYt7hsCwAA",
		Length:   2924,
	},

	"data/static/favicon.ico": {
		Filename: "data/static/favicon.ico",
		Contents: "H4sIAAAAAAAC/6TOMQoCQQyF4S9gYfnSWFvuMT2t9nsC2XFAmUbQH5IfAuE9SkminGzFBRuCK+N+cPMm04/77mz3K71iru9OvcxwjqLdVAzPqe58/K1xf/IMAAD//8Pv0P0+AQAA",
		Length:   318,
	},

	"data/static/humans.txt": {
		Filename: "data/static/humans.txt",
		Contents: "H4sIAAAAAAAC/yTIsQrCMBAA0Dn5iptLafdMSqlLcWqdRCQ2ZxpM7iS56O8Ldnu8voFlPJ6h6bUaNnyaWfCDMGF6azUwiV3FQPnn1cptZ8fZd/Wl1fINIpgNHIpwRLoXrrFodcqczOgCPWr2WwvzyhItuRYuFAQdTIG846R/AQAA//8PqwuogAAAAA==",
		Length:   128,
	},

	"data/static/img/e.png": {
		Filename: "data/static/img/e.png",
		Contents: "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",
		Length:   7351,
	},

	"data/static/index.html": {
		Filename: "data/static/index.html",
		Contents: "H4sIAAAAAAAC/5RWX2/cuBF/308xUYGejXrFc64Pl0TSIeckjdGkcc5rtHkquOJoRZsiFXK4slDcdy+Gkuy1k7Q5PfHP/P3NzE8sfrntDOzRB+1smZ3mP2aAtnZK212ZXW3erH/OfqlWxZNXH842ny5eQ0udgYurX9+dn0G2FuKfP50J8WrzCv71dvP+HZzmp0K8/kcGWUvUPxdiGIZ8+Cl3fic2v4lbVj89ZYVlnStSWbUqkuHbzthQfkX39NmzZ5NGxkLPjeT40CZNlKpaFR2SBNZc4+eo92VWO0toaU1jjxnMuzIjvCXBll5A3UofkMpIzfrnDES1Koy2N8AaZaY7uUNxu9a1sxl4NGU2LVuPTZmJRu55n+vaTbopAis7LDOFofa6J83yd64/Rl3fgLQKUIYRQiu9tjtwDXTS3yg3WODovrB2g+PgvAoHphaFE+hlINxqe3JnZLZ7cudguZgMBxoNVqs/aVubqBACSdK1qEMQ6SqvQ1gVYhabEEnZy0it83f5t7GTNuS0BEyaDFbvZ1+XrfQIz+FSd71BWI6XmAoxia8KMddv69RYrQql96BVmQ1e9tnBnqXQp3qfVoVcgsgeOsxr1xVCVoVoT1n0afWdmBeifcrBKL2vVvdeyfUd2shuo0lgPPD91nU4uTP68a3stciqlxfn3xKoPUpCkVVnafFNsRYlhRaRWJQ3l7z5lngjP4usevPy44GASLE/zm3upaxaQfoY2LsyXS5lYhzn+61x9c3n6AiXI4CirzatDhA0IUhj3BBgdBHIMd7aTIjf13/dON9JIlQJ9pNJJ1WEWvSBFfcaB6AWweoazQgerUKPCjyGaAjWcM595DwBsW8dQMLdRKxjYHPLWGQptLwQ/UHQjAfcfQzgpkWPbMm6RAs75EhiwBTIbILBfKj2yUVQzv5AYJFTcjBVFaQFWdcuWsoPVL5wvdg5bxJqnbxBkNDpQGllFcTeOKmWwU9SSiv22KG07JGPamlBoUFC0PRFpIuXD7ZGFv9hj7NdVLz397OwGPNI0VvQNuXfRIp+igeVpv/l42qxy+UFGcl1TC/SmBFaGRgYt21i4DPugqvf3nEScu+0gtBrhWk+uUW6nsJX/EztfLj/si5TObnxrINoPcrgrNwa5B4ir2tm5gDO3uMwI/xd5XppqHVx10KPjtltaB106LlZ53pdXrx8D4M2BhptFYOoPXcEhtSvHju3R/WHs3t4W/TV35CYwD1D+Zhasmphuwc1ZmLgGPbox0SKJ3AdA8FeB02p3F/hqLmt+WeQ9LkV5CDH1DA792i++qWhe48Nei5wiNtOEwyakSOYhlTC1rshsEQriYPqXQh6a/AE9loChx2mH8hjYk32Aj9bYOZYGFpdtxBi3ztPIWXCA+waeLvZXKwvPlxuLk/Aeahd10mr1kZbhNpotBQgxLoFGaAgqhby0FQIouowu0Ic0uDCqolUayNDKLPaoPRZBYdXzLeNc5T+YACH4gabiYNn+Ye3Xu/ahaLPXD+mPfy5dv34Ap7+ePrXIvTSTnwevUdL/x6T+0LwRQXb8R46fh2F50IEwj3mN9j1ecNYXvIe/o5dzzh+K5THeS3/k3n7ZL2GC2aJa7mX0+vnhKfv+mPkRluv+emRjuf3VXqI3QtzkoO2yg25s2mGSmiiTZN6dAz/YQiEAG54rixnmWZ1Lz2MHko4sjjAK0l4dHyc75DeRGM+ofRHxy9Ws/JVr7iLWV+5OnZoaTGRYCzvjtnAa4O8/HU8V0cP0T1+wWq6gaNJ70kJNhoDc5iQTnNO8Gzm7RKydQZ/gdEn1d9X8Du/sVLm/x+Yx2+16yBu8utwYGFViPkBld621eq/AQAA//8Uxk1d2gsAAA==",
		Length:   3034,
	},

	"data/static/js/k.js": {
		Filename: "data/static/js/k.js",
		Contents: "H4sIAAAAAAAC/5RTT2vcPhC951MoPkn8hPKDbdN0XRVK6aH0Two9LnuQpbE1sSMFW1530eq7F9nLbtKEQsEYz9N7mjfD80715Mvt9w/fPst6dDqgdxRYzHDNDW+44ppXZSOj9gbWm9UNX93wV//nZ/WGr94e39fX/Pr1lmsLKqzd2HUcOrhfT+iMn1JpZESHYX3qYVms5GbLtYypxDrXte9p7osEHWlYxJo2wqrhdnI/ev8AfdhTZDN+aTe4ZVFvcCubDW4TdAMsZT5JKR0R2SQjKnSGapEN8aKF/fhQcCNa79Q9ssQNDKH3+7M3Fo0Y3V9VZSW1zGMmnnln7R23HGePd0IZ82kHLnzFIYCDnsXnGLUcea26AdhieVGGoLSdibPoXNLCu+I/y08NWxaRtofDsmkBsyaxlBJfZnjZWw/3fgfP7L0Av+zQwBOHj8p/cbhs80kqlLSihf1Hb+BwsGKyqG0OiBY5gKID1wT7vhIP42CpYiz2EMbepT8p76rjB4uVGCzWgbJHpOB/hh5dQ9mllNWj8nShFnOYKUupNCKnlwIrl0NiUnlBSE7r0P4ikkQyk8manFNEIjmO3HmtMibUMGDjaHG1Q5iulukLRhI5XeerOyKJg+n4V9K5ASsvfgcAAP//xwls/qkDAAA=",
		Length:   937,
	},

	"data/static/markdownshare.com.conf": {
		Filename: "data/static/markdownshare.com.conf",
		Contents: "H4sIAAAAAAAC/2SNXQrDMAyD330Kn2B5H+RpxyhluKn7sy12iF3W448QBoM9SUIf0lDInKddRpjIWCgzRsxUn7O+xTaqfEmaofLKZ8GIm3uxawh/SAAYFq2ZfISk4iyOEZ1Ph0JrW20Cdkx5b003AMPMCx0vt/HLhVSZ/Ae99VzU/N4vMOLDVOATAAD//0WRvTS/AAAA",
		Length:   191,
	},

	"data/static/robots.txt": {
		Filename: "data/static/robots.txt",
		Contents: "H4sIAAAAAAAC/wotTi3STUxPzSuxUtDici5KLM/RTUnNSay0UjA04HLJLE7Myckvt1LQT07P1E3KzNPXwibIBQgAAP//q49ThkcAAAA=",
		Length:   71,
	},

	"data/templates/create.tmpl": {
		Filename: "data/templates/create.tmpl",
		Contents: "H4sIAAAAAAAC/5xVX0/kthd9z6e465/006AyMcPuwwJJVrsDdFGXFgqo5akyyc3EjGMH25kQjea7V3aSIUOXtupTcu1z7p/je+3o03MpYIXacCVjMgsPCKBMVcblIiZ3t+fTj+RTEkTvTn+Z395fnUFhSwFXd1++XcyBTCn97f2c0tPbU/j96+3lN5iFM0rPfiZACmurY0qbpgmb96HSC3r7K3129NnMEYb/MLMZSYLIO34uhTTxd7izo6OjjkEc6Fgwlx9Kz0SWJUFUomXgmFN8qvkqJqmSFqWd2rZCAr0VE4vPljpPJ5AWTBu0cW3z6UcCNAkiweUSHCMmvGQLpM9TnipJQKOISfdbaMxjQnO2cnbIU9VxfQaSlRiTDE2qeWW5w29DX9c8XQKTGSAzLZiCaS4XoHIomV5mqpHgsvuLtyW2jdKZGbkaCPtQMWPxgcv9rZPe7/42wLDROTa2FZgE/+MyFXWGYCyzPKWpMdRvhakxQUR7WKeIr57VtlB6W39Rl0ya0A4JW24FJpd9rJuCaYRj6L6tqvU2jYh20CCi/dk9qKxNgijjK+BZTBrNKjKyHQq1P+tZErEhAbIbLExVGVGWRLSYOehh8i/1jmhx6JLJ+CoJXqJaVZUoaxe2Fl6IndhfVYldOMFf77KKU5J8vrp4C5BqZBYpSeb+501YgcyaAtE6qDNunPEWPGdPlCTnn69HAOpzf11b30ckCQAA1mueQ+gGeLPxC07ne3dkg8BwpXHFsem09ZAHodLlU60sdgvOzY4P+hqxXqPMRhFGrXG5bY2/8R/lSpfAUuvvql5DAiXaQmUxqZSxBAYwgC82FcyYmFhGXjYAInfsTCMDrRoTkw8HbrSEicnhwQHpp86PYrJeh/NOrM0mogNvFKWTdmtWO3G4rGrb3yemfii5HbwP1oqJGmPS6+sn6b/w570YdJzZNpeIOumS75xLn/5Yq1Qg0ySB8ZbrmVwp66dwV1qBeddHWyXGu5oviqHN5qpqvQ3/T1XVnsDhwexDZComu56stUZp/2h9+Ii6jQQeWtg2uLvdzTGlxuIKwyWWVZi7QbtxNvyEZeUa/61UXtc1zERvvptO4arWCI9sxbrbex+kgsfrGnUL06m7Ov1yfx7+IXkBuyIbLjPVhEoKxTKIIa+l79bJHqydBJTCj2jBFgiuytCtrZiGVkMME4kNnDKLk729cIH2vBbiHpme7J0EPfmuyphFz89UWpco7eDCyxhvl52DM4Hu90t7kU121d07cTSew6TjvYtB1kJAnyb41dAV2Pc+xECmBH6AVnvqJoCNeyN85f8szOu35tHQZfhoxh5o/wb4pzkJ/gwAAP//PoBGSZkIAAA=",
		Length:   2201,
	},

	"data/templates/delete.tmpl": {
		Filename: "data/templates/delete.tmpl",
		Contents: "H4sIAAAAAAAC/4xVb1PbxhN+r0+xud/Mb8wU6eKkLwJIYhIbGqZJCwGm5VXnkFbW4dOduFvZaBi+e+dOsmNomPad9vZ59s+z63V6/NAoWKF10uiMTZO3DFAXppR6kbHrq9P4AzvOo/TN/PfZ1c35CdTUKDi//vTlbAYs5vyP9zPO51dz+PPz1dcvME2mnJ/8xoDVRO0h5+v1Olm/T4xd8Ktv/MHTp1NP2HwnJZUsj9IQ+KFR2mU/4E4PDg4GBvOgQyV8fagDE0WZR2mDJMAzY7zv5CpjhdGEmmLqW2QwWhkjfCDuIx1BUQvrkLKOqvgDA55HqZJ6CZ6RMdmIBfKHWBZGM7CoMjZ81harjPFKrLydyMIM3FCBFg1mrERXWNmS9Pht6otOFksQugQUrgdXCyv1AkwFjbDL0qw1+Or+EW2J/drY0u2E2hD2oRWO8Fbq/W2QMe7+NsHGMQR21CvMo/9JXaiuRHAkSBa8cI4HV1I4F6V8hA2KhO5FR7Wx2/7rrhHaJbQpmCQpzL+OuS5rYREOYY4KCaE3nd3WkfIBG6V8HN6tKfs8Sku5AllmbG1Fy3Zsj0Ibhj3NU7GpgD3PlhSmSbnIU15PPfRd/h8FT3n9zhdTylUefc9Kpm1Qdz5tp4ISz3J/Ng0O6ZR86RWt5Cz/eH72GqCwKAg5y2fh41VYjYJcjUge6o1Lb7wGr8Q9Z/npx4sdAA+1v+xtXCSWRwAAXtZvKJTqofzBvI4HRQPyVplied8ZwuEBIK2MbUAUFE4IHwLwx8fkbP70xKBBqk2ZsdY4YrAhQSp129H4U3PdbSOJjeu+sVZCdZixG3T74xqFRRv53Kcdi+K7VY29hlYLJZzLWKFQWJbDrsurUBlDYa8AduEKq0GZEf/ca+Wi3gg3M20fbPh/Ydr+CN69nf6culboQeXOWtT0Vx/Sp9w7crjtYTsyf7DcIeeOcIXJEps2qfzqXHobfsWm9aN8rZSXfW2mPJpv4hjOO4twJ1ZiOEj7oA3cXXRoe4hjfw3C8ziHcBu/g32Ta6lLs06MVkaUkEHV6TDpyR48egk4h1+QgGoE32Xi31bCQm8hg4nGNcwF4WRvL1kgnXZK3aCwk72jaCRft6UgDPzSFF2DmjYhgozZ9tkHOFHoPz/1Z+Xkubp7R54mK5gMvDcZ6E4pGMuE8Jr4BmfD6kMGLGbwE/Q2UJ8iePJnL3T+78K8PJ93ji+TO7cbgY9XLfzb5NHfAQAA//+9789EbAcAAA==",
		Length:   1900,
	},

	"data/templates/edit.tmpl": {
		Filename: "data/templates/edit.tmpl",
		Contents: "H4sIAAAAAAAC/5xV31PcNhB+91+xUWc6x5SzOJKHBGxnkgMaJqGFAtPy1BH2+ixOlowkn/Hc3P/ekfyDOxraTt+00n7f7n7alaKPT6WAFWrDlYzJLDwggDJVGZeLmNzenE3fk49JEL05+XV+c3d5CoUtBVzefv52PgcypfT3t3NKT25O4I8vNxffYBbOKD39hQAprK2OKG2aJmzehkov6M1v9MnBZzMHGNZhZjOSBJEnfiqFNPF3sLMPHz50COKcjgRz+aH0SGRZEkQlWgYOOcXHmq9ikippUdqpbSsk0FsxsfhkqWM6hrRg2qCNa5tP3xOgSRAJLpfgEDHhJVsgfZryVEkCGkVMumWhMY8JzdnK2SFPVYf1GUhWYkwyNKnmleXOfwx9VfN0CUxmgMy0YAqmuVyAyqFkepmpRoLL7m9sS2wbpTOzRTUA9qFixuI9l/sjSc+7PwYYDjpiY1uBSfADl6moMwRjmeUpTY2h/ihMjQki2rt1ivjqWW0Lpcf6i7pk0oR2SNhyKzC56GNdF0wjHMFpxi20qtZjFhHtPIOI9ld3r7I2CaKMr4BnMWk0q8iW7bxQ+6ueJREb4pPdWGGqyoiyJKLFzLkeJv9R7ogWhy6ZjK+S4DmqVVWJsnZha+F12In9RZXYhRP85SmrOCXJp8vz1xxSjcwiJcncL151K5BZUyBa5+qMa2e85p6zR0qSs09XWw7U5/6ytr6NSBIAAKzXPIfQze9m4zecznfuygaB4VLjimPTaetd7oVKl4+1sthtOJodDvrSY71GmW1FeO6Mi7Ez/oE+ypUugaXWv1QUM27peh1+xXazIVCiLVQWk0oZS2DAAPiSU8GMiYll5PkAIHKXzzQy0KoxMXl34OZLmJgcHhyQfvT8PCbrdTjvJNtsIjrgtqJ0Ao9mtROHy6q2/aNi6vuS24F9sFZM1BiTXmU/Tv8Hf1tlzOIOPKJjLhF1CibfuZ0+/W2tUoFMkwS2j1zn5EpZP4u70grMu24aldg+1XxRDM02V1XrbfgxVVV7DIcHs3eRqZjsOrPWGqX9s/XhI+oOErhvYWxz98SbI0qNxRWGSyyrMHfjdu1s+Ipl5dr/tVRe1jVMRm++mU7hstYID2zFuid8H6SCh6sadQvTqXs//XZ/H/43eXZ2RTZcZqoJlRSKZRBDXkvftJM9WDsJKIWf0YItEFyVodtbMQ2thhgmEhs4YRYne3vhAu1ZLcQdMj3ZOw56cHfLHp+ptC5R2oHCyxiP247gVKBbfm7Ps8muunvHDsZzmHS4NzHIWgjo0wS/G7oC+96HGMiUwE/Qag/dBLBxH4Wv/N+FefnhPBi6DB/MNgPtfwL/PyfBXwEAAP//1XCH3p4IAAA=",
		Length:   2206,
	},

	"data/templates/raw.tmpl": {
		Filename: "data/templates/raw.tmpl",
		Contents: "H4sIAAAAAAAC/6quVtBTqK3lAgQAAP//tfh1hQgAAAA=",
		Length:   8,
	},

	"data/templates/view.tmpl": {
		Filename: "data/templates/view.tmpl",
		Contents: "H4sIAAAAAAAC/4xVTXPbNhO+81ds8M68Y09NIkp6SGSSGceyG0+T1Emcpjl1IGIpwgIBBgBFcTT67x2AlCK5cVuduB/Pfjy7K6Sv1rWEFRortMrIJHlKAFWhuVCLjHy+u45fkFd5lD6Z/XZ59/X2CipXS7j9/PrtzSWQmNIvzy8pnd3N4I83d+/ewiSZUHr1ngCpnGumlHZdl3TPE20W9O4jXXv4ZOIBu++EO07yKA2B17VUNvsBdvLy5csBQbzTVDJfH6qARMbzKK3RMfDIGL+1YpWRQiuHysWub5DAKGXE4dpRH+kciooZiy5rXRm/IEDzKJVCLcEjMiJqtkC6jkWhFQGDMiPDZ2WwzAgt2crLiSj0gLWul5hH/xOqkC1HsI45UdDCWhpMSWFtlNLRbUgVwrLWVdrsA1dtzZRN3NoNcUNnitWYEY62MKJxwtexb+lDK4olMMUBme3BVswItQBdQs3MkutOge/6b9GW2HfacHsQagc4g4ZZh3OhzvZBxrhn+wQ7wxDYCScxfzfqPlXMIExhs0luZtttSgdzlNJxXnPN+zxKuViB4BnpDGvIgey90IT5TvKU7bghxwmSQtcpZXlKq4l3fZb/Ry5SWj3zxXCxOkjqdFOjan3WVoYRHaV+o2scsknx0MoaQUl+cXvzmENhkDmkJL8MH4+6VcicrRCdd/XCJy885l6yb5Tk1xcfDhxoqH1o7Xtv44hJHm02ooTkonXVdjvYw0pmZK4NRzOFSbMGq6XgYJCfezaa/K5CKLWUuvN0+t21YCvdSg5zBMtWyKHUBsrWtQahtQgxpHOP66ETUoLSLrhWfgZswYRK6TyfprT5Id0cJTqkm81QKmy3JJ8FHVwOrTzGCXLhjnFXXLgfoA6J2mxQ8e12JCf8mXl25lIXy2+tduhdvhsg/NJmR53fqZhJsVBTMGJRuXNyUJFhXSjoZhbK+V1gB65C+HjxZb+Wvq4kkBEi0+PMKC2Oaf1BvNcOrnWr+LD4QX0I2JUXxmZQcTTIYXc60OsWOjQIUuuln6efnLBhRBUa/Ic6AkmHd1NIZm1GConMkBwenlSptQt3DHDoLrF0g3LwP7YGBslQwaVu+iDD/wvd9Ofw7Onk59Q2TA1r3RqDyv3Zh/Qp9YYc5j3s2fdvgp1Sah2uMFli3SSlv9VPXoZfsW489Y+V8rCv3VmN4pM4hlu/8PdsxYb/5jNQGu4/tGh6iGP/LgT1+KqE5+e7s2+yE4rrLtFKasYhg7JVhf+HPzmFjaeAUvgFXVgY32XidStmoDeQwYnCDmbM4cnpabJAd91K+RWZOTk9j0bw54YzhwHPddHWqNwuRKAx26t9gCuJ/vN1f8NPjtk9PfcwUcLJgHuSgWqlhLFMCNrENzieGmRAYgI/QW8CdBuB352h838n5uFDem/pMrm3hxHo+IqEBz2P/goAAP//8o+tTs8IAAA=",
		Length:   2255,
	},

	"data/templates/view_raw.tmpl": {
		Filename: "data/templates/view_raw.tmpl",
		Contents: "H4sIAAAAAAAC/2yQsU7DMBCG9zzFcXt8irrQYmegqUSlAhUKAkbTuHWE45TkQlJVfXfkpGxs9n/f9+t08iZ7XuYf2xVYrhxsX+836yVgTPQ2WxJleQbvD/njBhKREK2eENAyHxdEfd+Lfibq5kD5Cw1BT5Ig/L1FwQWmkRyLh8r5Vv3jJvP5fDIwQAun/UGh8aNpdJFGsjKsIZix+e7KH4W72rPxHPPpaBCuP4VsBqbQdAc7q5vWsOp4H98iUBpJV/ovaIxTqDu2dYNgG7NXSLartG8FDzyBXLIz6fkMYp3B5SJpCiJJ130+6+I0zsfDBGJMAsCVS6PfAAAA//8zq5IMVAEAAA==",
		Length:   340,
	},
}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	if entry, ok := RESOURCES[path]; ok {
		var raw bytes.Buffer
		var err error

		// Decode the data.
		in, err := base64.StdEncoding.DecodeString(entry.Contents)
		if err != nil {
			return nil, err
		}

		// Gunzip the data to the client
		gr, err := gzip.NewReader(bytes.NewBuffer(in))
		if err != nil {
			return nil, err
		}
		defer gr.Close()
		data, err := ioutil.ReadAll(gr)
		if err != nil {
			return nil, err
		}
		_, err = raw.Write(data)
		if err != nil {
			return nil, err
		}

		// Return it.
		return raw.Bytes(), nil
	}
	return nil, fmt.Errorf("failed to find resource '%s'", path)
}

//
// Return the available resources in a slice.
//
func getResources() []EmbeddedResource {
	i := 0
	ret := make([]EmbeddedResource, len(RESOURCES))
	for _, v := range RESOURCES {
		ret[i] = v
		i++
	}
	return ret
}
