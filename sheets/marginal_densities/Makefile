# generated automatically by sheet.go

make:
	make marginal_densities.pdf

marginal_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > marginal_densities.tex

marginal_densities.pdf: ../../*.tex ../../trademark.pdf *.tex marginal_densities.tex
	pdflatex marginal_densities.tex
	make defs


open: marginal_densities.pdf
	open marginal_densities.pdf

o:
	make open

defs: marginal_densities.tex
	defs marginal_densities.tex

spell:
	aspell -c marginal_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile