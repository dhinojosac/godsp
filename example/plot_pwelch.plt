set size 1,1
set title "Welch's power spectral density"
set xlabel "Frequency [Hz]"
set ylabel "Power [dB]"
set xrange [0:20]
plot 'psd.dat' using 1:2 with lines title "PSD Input", '' using 1:3 with lines title "SOS Filtered"

set term png
set output "psd_image.png"
plot 'psd.dat' using 1:2 with lines title "PSD Input", '' using 1:3 with lines title "SOS Filtered"

