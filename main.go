package main

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	f := figure.NewFigure(strings.ToUpper("fooled ya!"), "", true)
	f.Print()
	banana := `
                              ';.
                            ;00dO0o.
                          'O0;...;Wo
                         dXl.....k0
                       ;Kk.......Xo
                      dX:........Wl
                    .0O.........'Wc
                   ;Xo...........Wc
                  lN:............Nl
                 oN;.............Xd
                lN;..........,:c:Xk.';,..
               ;Wc........'do:'..;dXl..;ld'
              .Xd........,0.      .;K,   .0:
              dX.........xc      .WMMk  ;WMK
             .Wl.........lx       ;c0o  .lKk
             oX...'.......ld,     'dl   .oo
             Kk...oKo,......:ollloXXllll;
             No...'NNXXOxc;.......dX
            .Wo....:WKOO0KXXK0kdl:dM.
             Xd.....;KN0kkkkkOO0KWNc
             k0.......lKNKOkkkkkkNN
             ,W;........;dKNX0OOkXM,
              O0............:ok0KNMO:.
    .'',col:. .Xd.................';0O       .'',col:.
   .NNKxl:dXWO;'Nd..................;W;     ,NNXxlcxXWk,
  .dW;      'OWl'Xk..................OO    .dW,      '0Wc
  xWc        cWO .OK,................;W,   kW:        lMk
  OWo        cNO   :Xd................O0   0Wl        lWk
   0Wd,    'dKM'    .kNl..............;W:   KNo'    ,dKW.
   ,OONNXXXKMMk.     KMOKo.............OMKl.,0MMXXXXXXo.
        ,:ol'lKWO:. kMd .oNk;..........,WKNWolM0.,:oc.
               'oXNKWd  ;0Wkx0x:........k0 :KMMc
                  ;l' ,0Mx.   ,d0kc.....;MOc,:,
                    ,KWO;        'oO0ddO0dlOXWK:.
           .'cokOOOO0KKOxl;.        .'..,cdkOKXK0OOOxl,.
         .xXx:'.       ..:dKO'       .dXxc'.        .;dKO'
         kN.                0K       xW.                kX
         ,k0OOkxxxkkkOOOOOOO0l       .x0OOkxxxkkkOOOOOOO0d
             ........   .                ...........
`

	fmt.Println(banana)
}
