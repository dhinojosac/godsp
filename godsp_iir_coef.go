package godsp

/*****************************************************************************/
/*                                                                           */
/* An second infinite impulse response (IIR) filter can be represented by    */
/* the following equation:                                                   */
/*                                                                           */
/* H(z) = b0 + b1.z^-1 + b2.z^-2                                             */
/*        ----------------------                                             */
/*        a0 + a1.z^-1 + a2.z^-2                                             */
/*                                                                           */
/* where H(z) is the transfer function. a0 is always 1.000                   */
/*****************************************************************************/

// Numerator coefficients
var IIR_B_COEF = []float64{
	0.001266188, -0.017765229, 0.118262374, -0.496695948, 1.475820182, -3.295489709, 5.730328712, -7.925616894, 8.819780649, -7.925616894, 5.730328712, -3.295489709, 1.475820182, -0.496695948, 0.118262374, -0.017765229, 0.001266188,
}

// Denominator coefficeints
var IIR_A_COEF = []float64{
	1.000000000, -15.185452913, 108.386826969, -482.677212168, 1501.084046102, -3456.820826551, 6097.872071830, -8405.093482948, 9148.751971980, -7890.095327386, 5373.539353729, -2859.611006096, 1165.712984279, -351.893997513, 74.184923781, -9.758226777, 0.603353683,
}

// sos design filter
var G float64 = 0.0012662

var B_Coef = [][]float64{
	{-6.625455e-01, 1.000000e+00},
	{-1.705328e+00, 1},
	{-1.999908e+00, 1.000000e+00},
	{-1.819789e+00, 1},
	{-1.999416e+00, 1.000000e+00},
	{-1.845643e+00, 1.000000e+00},
	{-1.999016e+00, 1.000000e+00},
	{-1.998843e+00, 1},
}

var A_Coef = [][]float64{
	{-1.800924e+00, 8.311287e-01},
	{-1.800524e+00, 8.743608e-01},
	{-1.921320e+00, 9.266921e-01},
	{-1.835910e+00, 9.409341e-01},
	{-1.974803e+00, 9.770868e-01},
	{-1.864612e+00, 9.838206e-01},
	{-1.990704e+00, 9.923671e-01},
	{-1.996655e+00, 9.981525e-01},
}

// end sos design filter

var IIR_signal_test = []float64{
	7.19E+03,
	6.35E+03,
	7.14E+03,
	7.23E+03,
	6.33E+03,
	6.85E+03,
	7.33E+03,
	6.48E+03,
	6.77E+03,
	7.80E+03,
	7.05E+03,
	7.07E+03,
	8.17E+03,
	7.37E+03,
	7.22E+03,
	8.42E+03,
	7.74E+03,
	7.17E+03,
	8.24E+03,
	7.78E+03,
	7.91E+03,
	6.95E+03,
	7.55E+03,
	8.53E+03,
	8.09E+03,
	8.30E+03,
	9.12E+03,
	8.80E+03,
	8.61E+03,
	9.99E+03,
	9.64E+03,
	9.09E+03,
	1.02E+04,
	9.71E+03,
	9.17E+03,
	1.02E+04,
	9.44E+03,
	8.92E+03,
	9.54E+03,
	8.82E+03,
	8.22E+03,
	9.08E+03,
	8.09E+03,
	7.61E+03,
	8.62E+03,
	7.75E+03,
	7.75E+03,
	8.90E+03,
	8.06E+03,
	8.18E+03,
	9.07E+03,
	8.13E+03,
	8.46E+03,
	9.08E+03,
	8.18E+03,
	8.68E+03,
	8.95E+03,
	8.07E+03,
	8.30E+03,
	8.98E+03,
	7.86E+03,
	8.42E+03,
	8.64E+03,
	7.84E+03,
	8.67E+03,
	8.75E+03,
	8.09E+03,
	8.97E+03,
	8.87E+03,
	8.19E+03,
	9.13E+03,
	8.74E+03,
	8.05E+03,
	9.06E+03,
	8.48E+03,
	7.78E+03,
	8.84E+03,
	8.15E+03,
	7.59E+03,
	8.63E+03,
	7.93E+03,
	7.51E+03,
	8.54E+03,
	7.94E+03,
	7.40E+03,
	8.30E+03,
	8.12E+03,
	7.64E+03,
	8.47E+03,
	8.50E+03,
	7.87E+03,
	8.55E+03,
	8.70E+03,
	7.88E+03,
	8.27E+03,
	8.72E+03,
	7.72E+03,
	8.08E+03,
	8.72E+03,
	7.56E+03,
	7.95E+03,
	8.61E+03,
	8.36E+03,
	7.67E+03,
	7.67E+03,
	8.67E+03,
	8.23E+03,
	7.70E+03,
	8.62E+03,
	8.26E+03,
	7.68E+03,
	8.35E+03,
	8.22E+03,
	7.64E+03,
	8.68E+03,
	8.58E+03,
	8.37E+03,
	9.41E+03,
	9.02E+03,
	8.40E+03,
	9.74E+03,
	9.21E+03,
	8.63E+03,
	9.63E+03,
	9.31E+03,
	8.59E+03,
	9.78E+03,
	9.36E+03,
	8.93E+03,
	1.02E+04,
	9.87E+03,
	9.33E+03,
	1.03E+04,
	9.42E+03,
	8.95E+03,
	9.67E+03,
	8.59E+03,
	8.24E+03,
	8.90E+03,
	7.94E+03,
	7.78E+03,
	8.45E+03,
	7.51E+03,
	7.64E+03,
	8.45E+03,
	7.54E+03,
	7.96E+03,
	8.70E+03,
	8.15E+03,
	8.30E+03,
	9.03E+03,
	8.04E+03,
	8.21E+03,
	8.86E+03,
	7.67E+03,
	8.04E+03,
	8.55E+03,
	7.42E+03,
	7.80E+03,
	8.16E+03,
	7.16E+03,
	7.84E+03,
	8.02E+03,
	7.21E+03,
	8.06E+03,
	7.95E+03,
	7.15E+03,
	7.67E+03,
	7.94E+03,
	6.92E+03,
	7.25E+03,
	7.89E+03,
	6.75E+03,
	6.72E+03,
	7.48E+03,
	6.58E+03,
	6.66E+03,
	7.68E+03,
	6.63E+03,
	6.30E+03,
	7.43E+03,
	6.84E+03,
	6.36E+03,
	7.24E+03,
	7.08E+03,
	5.96E+03,
	5.36E+03,
	6.29E+03,
	5.64E+03,
	4.80E+03,
	5.94E+03,
	5.31E+03,
	4.88E+03,
	5.84E+03,
	5.14E+03,
	4.49E+03,
	5.46E+03,
	4.49E+03,
	4.07E+03,
	5.32E+03,
	4.96E+03,
	4.86E+03,
	6.10E+03,
	5.88E+03,
	5.94E+03,
	7.49E+03,
	7.16E+03,
	7.20E+03,
	8.48E+03,
	8.01E+03,
	7.97E+03,
	8.93E+03,
	8.16E+03,
	8.38E+03,
	9.35E+03,
	8.53E+03,
	8.46E+03,
	9.83E+03,
	8.90E+03,
	9.26E+03,
	9.84E+03,
	8.50E+03,
	8.75E+03,
	8.99E+03,
	7.82E+03,
	8.42E+03,
	8.73E+03,
	7.73E+03,
	8.02E+03,
	9.14E+03,
	9.03E+03,
	8.54E+03,
	9.76E+03,
	9.30E+03,
	8.66E+03,
	9.67E+03,
	8.85E+03,
	8.09E+03,
	9.00E+03,
	8.11E+03,
	7.56E+03,
	8.47E+03,
	7.91E+03,
	7.32E+03,
	8.50E+03,
	7.95E+03,
	7.45E+03,
	8.28E+03,
	8.22E+03,
	7.63E+03,
	8.23E+03,
	8.39E+03,
	7.52E+03,
	7.74E+03,
	8.18E+03,
	7.03E+03,
	7.18E+03,
	7.73E+03,
	6.75E+03,
	7.70E+03,
	6.90E+03,
	6.77E+03,
	7.92E+03,
	7.38E+03,
	8.11E+03,
	7.63E+03,
	7.02E+03,
	7.79E+03,
	7.57E+03,
	6.78E+03,
	7.58E+03,
	7.43E+03,
	6.58E+03,
	7.24E+03,
	6.89E+03,
	5.99E+03,
	6.77E+03,
	6.28E+03,
	5.44E+03,
	6.48E+03,
	5.73E+03,
	4.90E+03,
	5.78E+03,
	5.60E+03,
	4.94E+03,
	6.16E+03,
	5.90E+03,
	5.47E+03,
	6.81E+03,
	6.33E+03,
	5.96E+03,
	7.30E+03,
	6.79E+03,
	6.44E+03,
	7.68E+03,
	6.93E+03,
	6.59E+03,
	7.64E+03,
	6.73E+03,
	6.64E+03,
	7.55E+03,
	6.73E+03,
	6.69E+03,
	7.78E+03,
	7.04E+03,
	6.91E+03,
	7.98E+03,
	7.22E+03,
	7.17E+03,
	8.12E+03,
	7.31E+03,
	7.38E+03,
	8.26E+03,
	7.33E+03,
	7.60E+03,
	8.54E+03,
	7.49E+03,
	8.01E+03,
	8.85E+03,
	7.87E+03,
	8.41E+03,
	9.00E+03,
	7.99E+03,
	8.62E+03,
	9.00E+03,
	7.80E+03,
	8.14E+03,
	8.88E+03,
	7.71E+03,
	7.75E+03,
	8.67E+03,
	7.77E+03,
	7.62E+03,
	8.56E+03,
	7.75E+03,
	7.42E+03,
	8.63E+03,
	8.02E+03,
	7.41E+03,
	8.56E+03,
	8.14E+03,
	7.36E+03,
	8.21E+03,
	8.27E+03,
	7.17E+03,
	7.96E+03,
	8.13E+03,
	7.33E+03,
	7.55E+03,
	8.25E+03,
	7.65E+03,
	7.47E+03,
	8.60E+03,
	8.01E+03,
	7.25E+03,
	8.26E+03,
	7.80E+03,
	7.11E+03,
	8.30E+03,
	7.73E+03,
	7.16E+03,
	8.34E+03,
	7.71E+03,
	7.20E+03,
	8.34E+03,
	7.63E+03,
	7.17E+03,
	8.24E+03,
	7.41E+03,
	7.18E+03,
	8.13E+03,
	7.60E+03,
	7.16E+03,
	8.34E+03,
	7.66E+03,
	7.46E+03,
	8.59E+03,
	8.22E+03,
	7.73E+03,
	8.95E+03,
	8.26E+03,
	7.85E+03,
	8.93E+03,
	8.10E+03,
	7.86E+03,
	8.82E+03,
	7.99E+03,
	7.89E+03,
	8.89E+03,
	8.20E+03,
	7.94E+03,
	9.03E+03,
	8.23E+03,
	8.13E+03,
	9.02E+03,
	8.18E+03,
	8.09E+03,
	8.90E+03,
	7.89E+03,
	7.90E+03,
	8.70E+03,
	7.55E+03,
	7.80E+03,
	8.54E+03,
	7.39E+03,
	7.73E+03,
	8.33E+03,
	7.31E+03,
	7.83E+03,
	8.40E+03,
	7.21E+03,
	7.58E+03,
	8.54E+03,
	7.57E+03,
	7.72E+03,
	8.71E+03,
	7.86E+03,
	7.75E+03,
	8.71E+03,
	7.95E+03,
	7.59E+03,
	8.70E+03,
	8.05E+03,
	7.45E+03,
	8.44E+03,
	8.07E+03,
	7.30E+03,
	8.23E+03,
	7.85E+03,
	8.47E+03,
	7.59E+03,
	8.04E+03,
	9.01E+03,
	8.41E+03,
	8.10E+03,
	9.17E+03,
	8.66E+03,
	8.07E+03,
	9.19E+03,
	8.45E+03,
	7.94E+03,
	9.00E+03,
	8.47E+03,
	7.83E+03,
	9.00E+03,
	8.35E+03,
	7.73E+03,
	8.85E+03,
	8.03E+03,
	7.71E+03,
	8.73E+03,
	7.93E+03,
	7.79E+03,
	8.73E+03,
	7.98E+03,
	7.92E+03,
	8.83E+03,
	7.89E+03,
	7.96E+03,
	8.72E+03,
	7.59E+03,
	7.81E+03,
	8.52E+03,
	7.40E+03,
	7.80E+03,
	8.39E+03,
	7.39E+03,
	7.89E+03,
	8.34E+03,
	7.38E+03,
	8.08E+03,
	8.36E+03,
	7.24E+03,
	7.69E+03,
	8.20E+03,
	7.22E+03,
	7.86E+03,
	8.17E+03,
	7.34E+03,
	8.20E+03,
	8.34E+03,
	7.68E+03,
	8.62E+03,
	8.65E+03,
	8.04E+03,
	9.12E+03,
	8.96E+03,
	8.34E+03,
	9.51E+03,
	9.09E+03,
	8.42E+03,
	9.56E+03,
	8.91E+03,
	8.03E+03,
	9.01E+03,
	8.58E+03,
	7.74E+03,
	8.52E+03,
	8.35E+03,
	7.54E+03,
	8.23E+03,
	8.55E+03,
	7.50E+03,
	7.96E+03,
	8.74E+03,
	7.68E+03,
	7.94E+03,
	8.94E+03,
	8.09E+03,
	8.04E+03,
	8.02E+03,
	8.95E+03,
	8.27E+03,
	7.92E+03,
	9.00E+03,
	8.40E+03,
	7.66E+03,
	8.80E+03,
	8.41E+03,
	7.75E+03,
	8.92E+03,
	8.35E+03,
	7.77E+03,
	8.64E+03,
	8.31E+03,
	7.61E+03,
	8.80E+03,
	8.27E+03,
	7.67E+03,
	8.86E+03,
	8.30E+03,
	7.69E+03,
	8.83E+03,
	8.19E+03,
	7.79E+03,
	9.09E+03,
	8.62E+03,
	8.15E+03,
	9.48E+03,
	8.92E+03,
	8.50E+03,
	9.67E+03,
	8.93E+03,
	8.62E+03,
	9.63E+03,
	8.74E+03,
	8.59E+03,
	9.51E+03,
	8.60E+03,
	8.45E+03,
	9.24E+03,
	8.20E+03,
	8.18E+03,
	8.96E+03,
	7.78E+03,
	8.04E+03,
	8.76E+03,
	7.65E+03,
	8.07E+03,
	8.92E+03,
	7.89E+03,
	8.05E+03,
	8.78E+03,
	7.65E+03,
	8.02E+03,
	8.72E+03,
	7.67E+03,
	8.05E+03,
	8.58E+03,
	7.60E+03,
	8.22E+03,
	8.61E+03,
	7.46E+03,
	7.85E+03,
	8.68E+03,
	7.62E+03,
	7.80E+03,
	8.77E+03,
	7.95E+03,
	7.86E+03,
	8.89E+03,
	8.10E+03,
	7.77E+03,
	8.91E+03,
	8.29E+03,
	7.67E+03,
	8.85E+03,
	8.37E+03,
	7.68E+03,
	8.64E+03,
	8.57E+03,
	7.94E+03,
	8.87E+03,
	9.13E+03,
	8.18E+03,
	8.66E+03,
	9.48E+03,
	8.90E+03,
	8.79E+03,
	9.75E+03,
	9.20E+03,
	8.55E+03,
	9.66E+03,
	8.94E+03,
	8.37E+03,
	9.43E+03,
	8.58E+03,
	8.22E+03,
	9.24E+03,
	8.36E+03,
	8.12E+03,
	9.24E+03,
	8.55E+03,
	8.12E+03,
	9.19E+03,
	8.37E+03,
	8.16E+03,
	9.12E+03,
	8.31E+03,
	8.16E+03,
	9.00E+03,
	8.12E+03,
	8.09E+03,
	8.87E+03,
	7.84E+03,
	7.96E+03,
	8.69E+03,
	7.52E+03,
	7.87E+03,
	8.55E+03,
	7.51E+03,
	7.95E+03,
	8.47E+03,
	7.46E+03,
	8.10E+03,
	8.36E+03,
	7.49E+03,
	8.29E+03,
	8.53E+03,
	7.50E+03,
	8.03E+03,
	8.50E+03,
	7.62E+03,
	8.39E+03,
	8.62E+03,
	7.91E+03,
	8.88E+03,
	8.92E+03,
	8.29E+03,
	9.30E+03,
	9.21E+03,
	8.60E+03,
	9.64E+03,
	9.32E+03,
	8.49E+03,
	9.26E+03,
	9.20E+03,
	8.14E+03,
	8.75E+03,
	8.99E+03,
	7.86E+03,
	8.21E+03,
	8.97E+03,
	7.88E+03,
	8.13E+03,
	9.07E+03,
	8.25E+03,
	8.23E+03,
	9.23E+03,
	8.41E+03,
	8.27E+03,
	9.34E+03,
	8.69E+03,
	8.20E+03,
	8.03E+03,
	9.07E+03,
	8.60E+03,
	7.77E+03,
	8.20E+03,
	8.56E+03,
	7.56E+03,
	7.96E+03,
	8.78E+03,
	7.81E+03,
	7.88E+03,
	8.86E+03,
	8.03E+03,
	7.91E+03,
	8.80E+03,
	7.93E+03,
	7.86E+03,
	8.67E+03,
	7.66E+03,
	7.80E+03,
	8.58E+03,
	7.50E+03,
	7.88E+03,
	8.63E+03,
	7.68E+03,
	8.26E+03,
	8.94E+03,
	8.11E+03,
	8.83E+03,
	9.28E+03,
	8.47E+03,
	9.31E+03,
	9.35E+03,
	8.54E+03,
	9.32E+03,
	9.13E+03,
	8.23E+03,
	8.81E+03,
	8.92E+03,
	7.95E+03,
	8.71E+03,
	8.60E+03,
	7.84E+03,
	8.63E+03,
	8.50E+03,
	7.81E+03,
	8.82E+03,
	8.59E+03,
	7.93E+03,
	9.10E+03,
	8.69E+03,
	8.03E+03,
	9.26E+03,
	8.67E+03,
	8.04E+03,
	9.17E+03,
	8.48E+03,
	7.95E+03,
	9.02E+03,
	8.24E+03,
	7.91E+03,
	8.89E+03,
	8.04E+03,
	7.89E+03,
	8.84E+03,
	8.15E+03,
	7.68E+03,
	8.85E+03,
	8.28E+03,
	7.68E+03,
	8.44E+03,
	8.11E+03,
	7.34E+03,
	8.23E+03,
	8.34E+03,
	7.37E+03,
	8.05E+03,
	8.52E+03,
	7.46E+03,
	7.98E+03,
	8.91E+03,
	8.02E+03,
	8.42E+03,
	9.54E+03,
	8.87E+03,
	8.76E+03,
	9.95E+03,
	9.44E+03,
	8.79E+03,
	9.86E+03,
	9.18E+03,
	8.55E+03,
	9.59E+03,
	8.73E+03,
	8.24E+03,
	9.19E+03,
	8.60E+03,
	7.99E+03,
	9.10E+03,
	8.46E+03,
	8.01E+03,
	9.11E+03,
	8.39E+03,
	8.11E+03,
	9.12E+03,
	8.29E+03,
	8.19E+03,
	9.08E+03,
	8.29E+03,
	8.21E+03,
	8.99E+03,
	8.05E+03,
	8.06E+03,
	8.79E+03,
	7.64E+03,
	7.86E+03,
	8.52E+03,
	7.37E+03,
	7.75E+03,
	8.38E+03,
	7.35E+03,
	7.84E+03,
	8.34E+03,
	7.41E+03,
	7.75E+03,
	8.49E+03,
	7.42E+03,
	7.87E+03,
	8.49E+03,
	7.52E+03,
	8.09E+03,
	8.55E+03,
	7.72E+03,
	8.56E+03,
	8.78E+03,
	8.11E+03,
	9.11E+03,
	9.14E+03,
	8.53E+03,
	9.45E+03,
	9.35E+03,
	8.58E+03,
	9.31E+03,
	9.24E+03,
	8.31E+03,
	8.74E+03,
	9.28E+03,
	8.07E+03,
	8.35E+03,
	9.16E+03,
	8.23E+03,
	8.18E+03,
	9.12E+03,
	8.24E+03,
	8.17E+03,
	9.19E+03,
	8.50E+03,
	8.13E+03,
	9.29E+03,
	8.76E+03,
	8.19E+03,
	9.06E+03,
	8.74E+03,
	8.13E+03,
	8.81E+03,
	8.65E+03,
	7.81E+03,
	8.44E+03,
	8.78E+03,
	7.61E+03,
	8.00E+03,
	8.85E+03,
	7.95E+03,
	7.78E+03,
	8.74E+03,
	8.00E+03,
	7.53E+03,
	8.55E+03,
	7.73E+03,
	7.53E+03,
	8.48E+03,
	7.68E+03,
	7.64E+03,
	8.55E+03,
	7.71E+03,
	7.70E+03,
	8.58E+03,
	7.69E+03,
	7.93E+03,
	8.82E+03,
	7.94E+03,
	8.39E+03,
	9.24E+03,
	8.29E+03,
	8.83E+03,
	9.54E+03,
	8.59E+03,
	9.11E+03,
	9.50E+03,
	8.48E+03,
	9.13E+03,
	9.10E+03,
	8.12E+03,
	8.43E+03,
	8.85E+03,
	7.80E+03,
	8.39E+03,
	8.59E+03,
	7.72E+03,
	8.05E+03,
	8.69E+03,
	7.68E+03,
	8.22E+03,
	8.63E+03,
	7.74E+03,
	8.45E+03,
	8.59E+03,
	7.75E+03,
	8.55E+03,
	8.44E+03,
	7.69E+03,
	8.46E+03,
	8.26E+03,
	7.60E+03,
	8.47E+03,
	8.18E+03,
	7.31E+03,
	8.19E+03,
	8.17E+03,
	7.45E+03,
	7.99E+03,
	8.42E+03,
	7.37E+03,
	7.74E+03,
	8.56E+03,
	7.58E+03,
	7.74E+03,
	8.65E+03,
	7.80E+03,
	7.77E+03,
	8.74E+03,
	7.95E+03,
	7.73E+03,
	8.83E+03,
	8.25E+03,
	7.82E+03,
	8.94E+03,
	8.22E+03,
	9.02E+03,
	8.46E+03,
	8.72E+03,
	9.78E+03,
	9.05E+03,
	8.87E+03,
	9.92E+03,
	9.39E+03,
	8.70E+03,
	9.80E+03,
	9.07E+03,
	8.46E+03,
	9.53E+03,
	8.72E+03,
	8.21E+03,
	9.22E+03,
	8.62E+03,
	7.98E+03,
	9.10E+03,
	8.46E+03,
	7.97E+03,
	9.03E+03,
	8.29E+03,
	7.94E+03,
	8.95E+03,
	8.14E+03,
	8.00E+03,
	8.86E+03,
	8.02E+03,
	7.92E+03,
	8.70E+03,
	7.79E+03,
	7.83E+03,
	8.52E+03,
	7.49E+03,
	7.71E+03,
	8.60E+03,
	7.78E+03,
	7.77E+03,
	8.56E+03,
	7.64E+03,
	7.78E+03,
	8.52E+03,
	7.40E+03,
	7.76E+03,
	8.42E+03,
	7.34E+03,
	7.72E+03,
	8.27E+03,
	7.26E+03,
	7.80E+03,
	8.19E+03,
	7.38E+03,
	8.09E+03,
	8.21E+03,
	7.46E+03,
	8.39E+03,
	8.39E+03,
	7.83E+03,
	8.87E+03,
	8.86E+03,
	8.31E+03,
	9.23E+03,
	9.31E+03,
	8.54E+03,
	9.10E+03,
	9.52E+03,
	8.54E+03,
	8.79E+03,
	9.30E+03,
	8.51E+03,
	8.47E+03,
	9.17E+03,
	8.45E+03,
	8.32E+03,
	9.12E+03,
	8.52E+03,
	8.05E+03,
	9.06E+03,
	8.54E+03,
	7.96E+03,
	8.01E+03,
	8.97E+03,
	8.45E+03,
	7.88E+03,
	8.79E+03,
	8.33E+03,
	7.59E+03,
	8.03E+03,
	8.54E+03,
	7.38E+03,
	7.76E+03,
	8.36E+03,
	7.34E+03,
	7.75E+03,
	8.23E+03,
	7.25E+03,
	7.85E+03,
	8.06E+03,
	7.22E+03,
	7.65E+03,
	8.25E+03,
	7.28E+03,
	7.83E+03,
	8.25E+03,
	7.34E+03,
	8.06E+03,
	8.15E+03,
	7.34E+03,
	8.14E+03,
	8.03E+03,
	7.33E+03,
	8.14E+03,
	7.97E+03,
	7.36E+03,
	8.41E+03,
	8.10E+03,
	7.57E+03,
	8.86E+03,
	8.52E+03,
	8.10E+03,
	9.49E+03,
	9.04E+03,
	8.67E+03,
	9.91E+03,
	9.28E+03,
	8.93E+03,
	9.97E+03,
	9.13E+03,
	8.68E+03,
	9.72E+03,
	8.97E+03,
	8.36E+03,
	9.31E+03,
	8.47E+03,
	8.10E+03,
	8.97E+03,
	8.09E+03,
	7.96E+03,
	8.78E+03,
	7.97E+03,
	7.91E+03,
	8.70E+03,
	7.77E+03,
	7.89E+03,
	8.59E+03,
	7.97E+03,
	7.90E+03,
	8.68E+03,
	8.04E+03,
	7.72E+03,
	8.67E+03,
	8.11E+03,
	7.54E+03,
	8.60E+03,
	7.98E+03,
	7.38E+03,
	8.58E+03,
	7.92E+03,
	7.29E+03,
	8.27E+03,
	7.99E+03,
	7.25E+03,
	7.65E+03,
	8.09E+03,
	7.16E+03,
	7.56E+03,
	8.27E+03,
	7.65E+03,
	7.67E+03,
	8.63E+03,
	7.81E+03,
	7.72E+03,
	8.58E+03,
	7.99E+03,
	7.61E+03,
	8.64E+03,
	7.87E+03,
	7.74E+03,
	8.77E+03,
	8.12E+03,
	8.21E+03,
	9.17E+03,
	8.49E+03,
	8.70E+03,
	9.55E+03,
	8.98E+03,
	8.92E+03,
	9.74E+03,
	8.90E+03,
	8.77E+03,
	9.42E+03,
	8.35E+03,
	8.36E+03,
	8.94E+03,
	7.73E+03,
	7.93E+03,
	8.49E+03,
	7.37E+03,
	7.74E+03,
	8.27E+03,
	7.31E+03,
	7.91E+03,
	8.25E+03,
	7.39E+03,
	8.16E+03,
	8.18E+03,
	7.43E+03,
	8.24E+03,
	8.07E+03,
	7.38E+03,
	8.28E+03,
	8.00E+03,
	7.36E+03,
	8.39E+03,
	7.96E+03,
	7.19E+03,
	8.02E+03,
	7.84E+03,
	7.17E+03,
	8.19E+03,
	7.80E+03,
	7.24E+03,
	7.89E+03,
	7.71E+03,
	7.07E+03,
	7.74E+03,
	7.84E+03,
	7.02E+03,
	7.43E+03,
	7.92E+03,
	6.96E+03,
	7.36E+03,
	8.03E+03,
	6.94E+03,
	7.48E+03,
	8.22E+03,
	7.20E+03,
	7.43E+03,
	8.30E+03,
	7.48E+03,
	7.40E+03,
	8.51E+03,
	7.82E+03,
	7.71E+03,
	9.06E+03,
	8.66E+03,
	8.19E+03,
	8.62E+03,
	9.19E+03,
	8.35E+03,
	8.67E+03,
	9.59E+03,
	8.57E+03,
	8.51E+03,
	9.15E+03,
	8.03E+03,
	8.14E+03,
	8.78E+03,
	7.64E+03,
	7.57E+03,
	8.43E+03,
	7.66E+03,
	7.70E+03,
	8.47E+03,
	7.36E+03,
	7.73E+03,
	8.40E+03,
	7.32E+03,
	7.77E+03,
	8.27E+03,
	7.27E+03,
	7.86E+03,
	8.03E+03,
	7.13E+03,
	7.90E+03,
	7.87E+03,
	7.09E+03,
	7.95E+03,
	7.79E+03,
	7.15E+03,
	8.07E+03,
	7.81E+03,
	7.16E+03,
	8.27E+03,
	7.79E+03,
	7.21E+03,
	8.40E+03,
	7.74E+03,
	7.22E+03,
	8.32E+03,
	7.61E+03,
	7.13E+03,
	8.27E+03,
	7.73E+03,
	7.15E+03,
	8.35E+03,
	7.74E+03,
	7.28E+03,
	8.38E+03,
	7.66E+03,
	7.33E+03,
	8.28E+03,
	7.46E+03,
	7.30E+03,
	8.16E+03,
	7.33E+03,
	7.17E+03,
	8.23E+03,
	7.73E+03,
	7.20E+03,
	8.47E+03,
	7.84E+03,
	7.23E+03,
	8.49E+03,
	8.02E+03,
	7.54E+03,
	8.55E+03,
	8.61E+03,
	7.98E+03,
	9.02E+03,
	9.27E+03,
	8.37E+03,
	9.06E+03,
	9.45E+03,
	8.30E+03,
	8.63E+03,
	8.10E+03,
	8.47E+03,
	8.75E+03,
	7.72E+03,
	8.00E+03,
	8.86E+03,
	7.84E+03,
	7.88E+03,
	8.81E+03,
	8.23E+03,
	7.73E+03,
	8.77E+03,
	7.95E+03,
	7.72E+03,
	8.69E+03,
	7.84E+03,
	7.72E+03,
	8.58E+03,
	7.78E+03,
	7.69E+03,
	8.49E+03,
	7.49E+03,
	7.65E+03,
	8.30E+03,
	7.21E+03,
	7.52E+03,
	8.14E+03,
	7.15E+03,
	7.48E+03,
	7.98E+03,
	7.02E+03,
	7.67E+03,
}