
build:
	@echo "Building executable"
	go build -o main.exe

run: build
	@echo "Running "
	./main.exe < input_signal.dat

plot: build run 
	@echo "Plot output"
	gnuplot plot_pwelch.plt -p;


