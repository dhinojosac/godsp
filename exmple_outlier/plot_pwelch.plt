set size 1,1
set title "Remove Outlier"
set xlabel "Samples"
set ylabel "Amplitude"

#set yrange [-30:80]
#set ytics 5
#set xtics 0.5
#set grid ytics lc "red"
#set grid xtics lc "red"

plot 'out.dat' using 0:1 with lines title "outlier output", "" using 0:2 with lines title "outlier output"

#set term png
#set output "psd_image.png"
#plot 'out.dat' using 0:1 with lines title "outlier output"



