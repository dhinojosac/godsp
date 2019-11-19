set size 1,1
set title "Welch's power spectral density"
set xlabel "Frequency [Hz]"
set ylabel "Power [dB]"
set yrange [-30:80]
set ytics 5
set xtics 0.5
set grid ytics lc "red"
set grid xtics lc "red"

plot 'psd.dat' using 1:2 with lines title "psd output"

set term png
set output "psd_image.png"
plot 'psd.dat' using 1:2 with lines title "psd output"



