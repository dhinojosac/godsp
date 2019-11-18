set size 1,1
set title "Welch's power spectral density"
set xlabel "Frequency [Hz]"
set ylabel "Power [dB]"
plot 'psd.dat' using 1:2 with lines title "psd output"

