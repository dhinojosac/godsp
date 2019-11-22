set size 1,1
set title "IIR FILTER"
plot 'out.dat' using 0:1 with lines title "input signal",'' using 0:2 with lines title "output iir filter"

set term png
set output "filtered_image.png"
plot 'out.dat' using 0:1 with lines title "input signal", '' using 0:2 with lines title "output iir filter"



